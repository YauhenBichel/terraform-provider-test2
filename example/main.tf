terraform {
  required_providers {
    hashicups = {
      source  = "hashicorp.com/edu/hashicups"
    }
  }
  required_version = ">= 1.8.0"
}

provider "hashicups" {
  username = "education"
  password = "test123"
  host     = "http://localhost:19090"
}

data "hashicups_coffees" "edu" {}

output "edu_coffees" {
    value = data.hashicups_coffees.edu
}

resource "hashicups_order" "edu" {
  items = [{
    coffee = {
      id = 3
    }
    quantity = 2
    },
    {
      coffee = {
        id = 2
      }
      quantity = 3
  }]
}


output "edu_order" {
  value = hashicups_order.edu
}

output "total_price" {
  value = provider::hashicups::compute_tax(5.00, 0.085)
}