terraform {
  backend "s3" {
    // config passed in through backend-config init args, see apply.sh
  }
}

provider "aws" {
  region = var.region
}

provider "kubernetes" {
  host                   = module.cluster.eks_cluster_endpoint
  cluster_ca_certificate = base64decode(module.cluster.eks_cluster_certificate_authority_data)
  exec {
    api_version = "client.authentication.k8s.io/v1beta1"
    command     = "aws"
    args        = ["eks", "get-token", "--cluster-name", module.cluster.eks_cluster_id]
  }
}

provider "helm" {
  kubernetes {
    host                   = module.cluster.eks_cluster_endpoint
    cluster_ca_certificate = base64decode(module.cluster.eks_cluster_certificate_authority_data)
    exec {
      api_version = "client.authentication.k8s.io/v1beta1"
      command     = "aws"
      args        = ["eks", "get-token", "--cluster-name", module.cluster.eks_cluster_id]
    }
  }
}

provider "cloudflare" {
  api_token = var.cloudflare_api_token
}

locals {
  node_pool_label_key = "node-pool"
  system_node_pool    = "xmtp-system"
  nodes_node_pool     = "xmtp-nodes"
  ingress_class_name  = "traefik"
  node_api_http_port  = 5001

  argocd_hostnames     = [for hostname in var.hostnames : "argo.${hostname}"]
  chat_app_hostnames   = [for hostname in var.hostnames : "chat.${hostname}"]
  grafana_hostnames    = [for hostname in var.hostnames : "grafana.${hostname}"]
  jaeger_hostnames     = [for hostname in var.hostnames : "jaeger.${hostname}"]
  prometheus_hostnames = [for hostname in var.hostnames : "prometheus.${hostname}"]
}

module "ecr_node_repo" {
  source  = "cloudposse/ecr/aws"
  version = "0.35.0"

  namespace  = "xmtp"
  stage      = "devnet"
  name       = "aws"
  attributes = ["node"]

  force_delete = true
}

module "cluster" {
  source = "../../modules/clusters/aws"

  namespace = "xmtp"
  stage     = "devnet"
  name      = "aws"

  region                       = var.region
  availability_zones           = var.availability_zones
  vpc_cidr_block               = var.vpc_cidr_block
  kubernetes_version           = var.kubernetes_version
  enabled_cluster_log_types    = var.enabled_cluster_log_types
  cluster_log_retention_period = var.cluster_log_retention_period

  node_pools = [
    {
      name           = local.system_node_pool
      instance_types = ["t3.small"]
      desired_size   = 2
      max_size       = 3
      min_size       = 2
      labels = {
        (local.node_pool_label_key) = local.system_node_pool
      }
    },
    {
      name           = local.nodes_node_pool
      instance_types = ["t3.small"]
      desired_size   = 2
      max_size       = 3
      min_size       = 2
      labels = {
        (local.node_pool_label_key) = local.nodes_node_pool
      }
    }
  ]
}

module "system" {
  source     = "../../modules/cluster-system"
  depends_on = [module.cluster]

  namespace            = "xmtp-system"
  node_pool_label_key  = local.node_pool_label_key
  node_pool            = local.system_node_pool
  argocd_project       = "xmtp-system"
  argocd_hostnames     = local.argocd_hostnames
  ingress_class_name   = local.ingress_class_name
  ingress_service_type = "LoadBalancer"
}

module "tools" {
  source     = "../../modules/cluster-tools"
  depends_on = [module.system]

  namespace            = "xmtp-tools"
  node_pool_label_key  = local.node_pool_label_key
  node_pool            = local.system_node_pool
  argocd_namespace     = module.system.namespace
  argocd_project       = "xmtp-tools"
  ingress_class_name   = local.ingress_class_name
  wait_for_ready       = false
  enable_chat_app      = var.enable_chat_app
  enable_monitoring    = var.enable_monitoring
  public_api_url       = "http://${var.hostnames[0]}"
  chat_app_hostnames   = local.chat_app_hostnames
  grafana_hostnames    = local.grafana_hostnames
  jaeger_hostnames     = local.jaeger_hostnames
  prometheus_hostnames = local.prometheus_hostnames
}

module "nodes" {
  source     = "../../modules/cluster-nodes"
  depends_on = [module.system]

  namespace                 = "xmtp-nodes"
  container_image           = var.node_container_image
  node_pool_label_key       = local.node_pool_label_key
  node_pool                 = local.nodes_node_pool
  argocd_namespace          = module.system.namespace
  argocd_project            = "xmtp-nodes"
  nodes                     = var.nodes
  node_keys                 = var.node_keys
  ingress_class_name        = local.ingress_class_name
  hostnames                 = var.hostnames
  node_api_http_port        = local.node_api_http_port
  storage_class_name        = "standard"
  container_storage_request = "1Gi"
  container_cpu_request     = "10m"
  debug                     = true
  wait_for_ready            = false
  one_instance_per_k8s_node = false
}

resource "cloudflare_record" "hostnames" {
  count   = length(var.hostnames)
  zone_id = var.cloudflare_zone_id
  name    = var.hostnames[count.index]
  value   = module.system.ingress_public_hostname
  type    = "CNAME"
  tags    = []
  proxied = true
}

resource "cloudflare_record" "argocd_hostnames" {
  count   = length(local.argocd_hostnames)
  zone_id = var.cloudflare_zone_id
  name    = local.argocd_hostnames[count.index]
  value   = module.system.ingress_public_hostname
  type    = "CNAME"
  tags    = []
  proxied = true
}

resource "cloudflare_record" "grafana_hostnames" {
  count   = length(local.grafana_hostnames)
  zone_id = var.cloudflare_zone_id
  name    = local.grafana_hostnames[count.index]
  value   = module.system.ingress_public_hostname
  type    = "CNAME"
  tags    = []
  proxied = true
}

resource "cloudflare_record" "jaeger_hostnames" {
  count   = length(local.jaeger_hostnames)
  zone_id = var.cloudflare_zone_id
  name    = local.jaeger_hostnames[count.index]
  value   = module.system.ingress_public_hostname
  type    = "CNAME"
  tags    = []
  proxied = true
}

resource "cloudflare_record" "prometheus_hostnames" {
  count   = length(local.prometheus_hostnames)
  zone_id = var.cloudflare_zone_id
  name    = local.prometheus_hostnames[count.index]
  value   = module.system.ingress_public_hostname
  type    = "CNAME"
  tags    = []
  proxied = true
}
