- go mod edit -module terraform-provider-test2
- go mod tidy
- mkdir docker_compose
- go run main.go
- terraform init
- go install .
- terraform plan
- go get github.com/hashicorp-demoapp/hashicups-client-go@v0.1.0
- go mod tidy
- go install .
- cd docker_compose
- docker-compose up
- curl localhost:19090/health/readyz
    kill postgres if it is running locally
- curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}'
- {"UserID":1,"Username":"education","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY5MTMxNDMsInRva2VuX2lkIjoxLCJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImVkdWNhdGlvbiJ9.6yVNw_hYe1oPTiK5e0EmddkzTrxan0q-dMD4F070xSM"}
    export HASHICUPS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY5MTMxNDMsInRva2VuX2lkIjoxLCJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImVkdWNhdGlvbiJ9.6yVNw_hYe1oPTiK5e0EmddkzTrxan0q-dMD4F070xSM

- Build and install the updated provider.
$ go install .

- terraform plan

- HASHICUPS_HOST=http://localhost:19090 \
  HASHICUPS_USERNAME=education \
  HASHICUPS_PASSWORD=test123 \
  terraform plan

    │ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - hashicorp.com/edu/hashicups in /Users/yauhenbichel/go/bin
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the state to become
│ incompatible with published releases.
╵
data.hashicups_coffees.example: Reading...
data.hashicups_coffees.example: Read complete after 0s

No changes. Your infrastructure matches the configuration.

Terraform has compared your real infrastructure against your configuration and found no differences, so no changes are needed.

- block data read from from server using client and updated state
terraform plan
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - hashicorp.com/edu/hashicups in /Users/yauhenbichel/go/bin
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the state to become
│ incompatible with published releases.
╵
data.hashicups_coffees.edu: Reading...
data.hashicups_coffees.edu: Read complete after 0s

Changes to Outputs:
  + edu_coffees = {
      + coffees = [
          + {
              + description = ""
              + id          = 1
              + image       = "/hashicorp.png"
              + ingredients = [
                  + {
                      + id = 6
                    },
                ]
              + name        = "HCP Aeropress"
              + price       = 200
              + teaser      = "Automation in a cup"
            },
          + {
              + description = ""
              + id          = 2
              + image       = "/packer.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 2
                    },
                  + {
                      + id = 4
                    },
                ]
              + name        = "Packer Spiced Latte"
              + price       = 350
              + teaser      = "Packed with goodness to spice up your images"
            },
          + {
              + description = ""
              + id          = 3
              + image       = "/vault.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 2
                    },
                ]
              + name        = "Vaulatte"
              + price       = 200
              + teaser      = "Nothing gives you a safe and secure feeling like a Vaulatte"
            },
          + {
              + description = ""
              + id          = 4
              + image       = "/nomad.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 3
                    },
                ]
              + name        = "Nomadicano"
              + price       = 150
              + teaser      = "Drink one today and you will want to schedule another"
            },
          + {
              + description = ""
              + id          = 5
              + image       = "/terraform.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                ]
              + name        = "Terraspresso"
              + price       = 150
              + teaser      = "Nothing kickstarts your day like a provision of Terraspresso"
            },
          + {
              + description = ""
              + id          = 6
              + image       = "/vagrant.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                ]
              + name        = "Vagrante espresso"
              + price       = 200
              + teaser      = "Stdin is not a tty"
            },
          + {
              + description = ""
              + id          = 7
              + image       = "/consul.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 5
                    },
                ]
              + name        = "Connectaccino"
              + price       = 250
              + teaser      = "Discover the wonders of our meshy service"
            },
          + {
              + description = ""
              + id          = 8
              + image       = "/boundary.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 6
                    },
                ]
              + name        = "Boundary Red Eye"
              + price       = 200
              + teaser      = "Perk up and watch out for your access management"
            },
          + {
              + description = ""
              + id          = 9
              + image       = "/waypoint.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 2
                    },
                ]
              + name        = "Waypointiato"
              + price       = 250
              + teaser      = "Deploy with a little foam"
            },
        ]
    }

- TF_LOG or TF_LOG_
>TF_LOG=TRACE terraform plan
>TF_LOG=TRACE TF_LOG_PATH=trace.txt terraform plan
  2024-05-27T19:20:51.722+0100 [DEBUG] provider.terraform-provider-hashicups: Creating HashiCups client: @module=hashicups tf_provider_addr=hashicorp.com/edu/hashicups tf_req_id=7d48ea4a-ba06-53d7-35a9-287e4ef5436e tf_rpc=ConfigureProvider @caller=/Users/yauhenbichel/DevBox/terraform/terraform-provider-test2/internal/provider/provider.go:171 hashicups_host=http://localhost:19090 hashicups_password="***" hashicups_username=education timestamp="2024-05-27T19:20:51.721+0100"

>TF_LOG=INFO terraform plan
>TF_LOG_PROVIDER=INFO terraform plan


>terraform apply
Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

hashicups_order.edu: Creating...
hashicups_order.edu: Creation complete after 0s [id=1]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.

Outputs:

edu_coffees = {
  "coffees" = tolist([
    {
      "description" = ""
      "id" = 1
      "image" = "/hashicorp.png"
      "ingredients" = tolist([
        {
          "id" = 6
        },
      ])
      "name" = "HCP Aeropress"
      "price" = 200
      "teaser" = "Automation in a cup"
    },
    {
      "description" = ""
      "id" = 2
      "image" = "/packer.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 2
        },
        {
          "id" = 4
        },
      ])
      "name" = "Packer Spiced Latte"
      "price" = 350
      "teaser" = "Packed with goodness to spice up your images"
    },
    {
      "description" = ""
      "id" = 3
      "image" = "/vault.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 2
        },
      ])
      "name" = "Vaulatte"
      "price" = 200
      "teaser" = "Nothing gives you a safe and secure feeling like a Vaulatte"
    },
    {
      "description" = ""
      "id" = 4
      "image" = "/nomad.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 3
        },
      ])
      "name" = "Nomadicano"
      "price" = 150
      "teaser" = "Drink one today and you will want to schedule another"
    },
    {
      "description" = ""
      "id" = 5
      "image" = "/terraform.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
      ])
      "name" = "Terraspresso"
      "price" = 150
      "teaser" = "Nothing kickstarts your day like a provision of Terraspresso"
    },
    {
      "description" = ""
      "id" = 6
      "image" = "/vagrant.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
      ])
      "name" = "Vagrante espresso"
      "price" = 200
      "teaser" = "Stdin is not a tty"
    },
    {
      "description" = ""
      "id" = 7
      "image" = "/consul.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 5
        },
      ])
      "name" = "Connectaccino"
      "price" = 250
      "teaser" = "Discover the wonders of our meshy service"
    },
    {
      "description" = ""
      "id" = 8
      "image" = "/boundary.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 6
        },
      ])
      "name" = "Boundary Red Eye"
      "price" = 200
      "teaser" = "Perk up and watch out for your access management"
    },
    {
      "description" = ""
      "id" = 9
      "image" = "/waypoint.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 2
        },
      ])
      "name" = "Waypointiato"
      "price" = 250
      "teaser" = "Deploy with a little foam"
    },
  ])
}
edu_order = {
  "id" = "1"
  "items" = tolist([
    {
      "coffee" = {
        "description" = ""
        "id" = 3
        "image" = "/vault.png"
        "name" = "Vaulatte"
        "price" = 200
        "teaser" = "Nothing gives you a safe and secure feeling like a Vaulatte"
      }
      "quantity" = 2
    },
    {
      "coffee" = {
        "description" = ""
        "id" = 1
        "image" = "/hashicorp.png"
        "name" = "HCP Aeropress"
        "price" = 200
        "teaser" = "Automation in a cup"
      }
      "quantity" = 2
    },
  ])
  "last_updated" = "Monday, 27-May-24 19:44:30 BST"
}

- terraform state show hashicups_order.edu
# hashicups_order.edu:
resource "hashicups_order" "edu" {
    id           = "1"
    items        = [
        {
            coffee   = {
                description = null
                id          = 3
                image       = "/vault.png"
                name        = "Vaulatte"
                price       = 200
                teaser      = "Nothing gives you a safe and secure feeling like a Vaulatte"
            }
            quantity = 2
        },
        {
            coffee   = {
                description = null
                id          = 1
                image       = "/hashicorp.png"
                name        = "HCP Aeropress"
                price       = 200
                teaser      = "Automation in a cup"
            }
            quantity = 2
        },
    ]
    last_updated = "Monday, 27-May-24 19:44:30 BST"
}

- curl -X GET  -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/1
{"id":1,"items":[{"coffee":{"id":3,"name":"Vaulatte","teaser":"Nothing gives you a safe and secure feeling like a Vaulatte","collection":"Foundations","origin":"Spring 2015","color":"#FFD814","description":"","price":200,"image":"/vault.png","ingredients":[{"ingredient_id":1},{"ingredient_id":2}]},"quantity":2},{"coffee":{"id":1,"name":"HCP Aeropress","teaser":"Automation in a cup","collection":"Foundations","origin":"Summer 2020","color":"#444","description":"","price":200,"image":"/hashicorp.png","ingredients":[{"ingredient_id":6}]},"quantity":2}]}%                                                                                                                   

- terraform apply -auto-approve
> curl -X GET -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/1
  {"id":1,"items":[{"coffee":{"id":3,"name":"Vaulatte","teaser":"Nothing gives you a safe and secure feeling like a Vaulatte","collection":"Foundations","origin":"Spring 2015","color":"#FFD814","description":"","price":200,"image":"/vault.png","ingredients":[{"ingredient_id":1},{"ingredient_id":2}]},"quantity":2},{"coffee":{"id":2,"name":"Packer Spiced Latte","teaser":"Packed with goodness to spice up your images","collection":"Origins","origin":"Summer 2013","color":"#1FA7EE","description":"","price":350,"image":"/packer.png","ingredients":[{"ingredient_id":1},{"ingredient_id":2},{"ingredient_id":4}]},"quantity":3}]}% 

- terraform destroy -au
to-approve
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides
│ are set in the CLI configuration:
│  - hashicorp.com/edu/hashicups in /Users/yauhenbichel/go/bin
│ 
│ The behavior may therefore not match any
│ released version of the provider and applying
│ changes may cause the state to become
│ incompatible with published releases.
╵
data.hashicups_coffees.edu: Reading...
hashicups_order.edu: Refreshing state... [id=1]
data.hashicups_coffees.edu: Read complete after 0s

Terraform used the selected providers to
generate the following execution plan. Resource
actions are indicated with the following
symbols:
  - destroy

Terraform will perform the following actions:

  # hashicups_order.edu will be destroyed
  - resource "hashicups_order" "edu" {
      - id           = "1" -> null
      - items        = [
          - {
              - coffee   = {
                  - id          = 3 -> null
                  - image       = "/vault.png" -> null
                  - name        = "Vaulatte" -> null
                  - price       = 200 -> null
                  - teaser      = "Nothing gives you a safe and secure feeling like a Vaulatte" -> null
                    # (1 unchanged attribute hidden)
                } -> null
              - quantity = 2 -> null
            },
          - {
              - coffee   = {
                  - id          = 2 -> null
                  - image       = "/packer.png" -> null
                  - name        = "Packer Spiced Latte" -> null
                  - price       = 350 -> null
                  - teaser      = "Packed with goodness to spice up your images" -> null
                    # (1 unchanged attribute hidden)
                } -> null
              - quantity = 3 -> null
            },
        ] -> null
      - last_updated = "Monday, 27-May-24 22:21:32 BST" -> null
    }

Plan: 0 to add, 0 to change, 1 to destroy.

Changes to Outputs:
  - edu_coffees = {
      - coffees = [
          - {
              - description = ""
              - id          = 1
              - image       = "/hashicorp.png"
              - ingredients = [
                  - {
                      - id = 6
                    },
                ]
              - name        = "HCP Aeropress"
              - price       = 200
              - teaser      = "Automation in a cup"
            },
          - {
              - description = ""
              - id          = 2
              - image       = "/packer.png"
              - ingredients = [
                  - {
                      - id = 1
                    },
                  - {
                      - id = 2
                    },
                  - {
                      - id = 4
                    },
                ]
              - name        = "Packer Spiced Latte"
              - price       = 350
              - teaser      = "Packed with goodness to spice up your images"
            },
          - {
              - description = ""
              - id          = 3
              - image       = "/vault.png"
              - ingredients = [
                  - {
                      - id = 1
                    },
                  - {
                      - id = 2
                    },
                ]
              - name        = "Vaulatte"
              - price       = 200
              - teaser      = "Nothing gives you a safe and secure feeling like a Vaulatte"
            },
          - {
              - description = ""
              - id          = 4
              - image       = "/nomad.png"
              - ingredients = [
                  - {
                      - id = 1
                    },
                  - {
                      - id = 3
                    },
                ]
              - name        = "Nomadicano"
              - price       = 150
              - teaser      = "Drink one today and you will want to schedule another"
            },
          - {
              - description = ""
              - id          = 5
              - image       = "/terraform.png"
              - ingredients = [
                  - {
                      - id = 1
                    },
                ]
              - name        = "Terraspresso"
              - price       = 150
              - teaser      = "Nothing kickstarts your day like a provision of Terraspresso"
            },
          - {
              - description = ""
              - id          = 6
              - image       = "/vagrant.png"
              - ingredients = [
                  - {
                      - id = 1
                    },
                ]
              - name        = "Vagrante espresso"
              - price       = 200
              - teaser      = "Stdin is not a tty"
            },
          - {
              - description = ""
              - id          = 7
              - image       = "/consul.png"
              - ingredients = [
                  - {
                      - id = 1
                    },
                  - {
                      - id = 5
                    },
                ]
              - name        = "Connectaccino"
              - price       = 250
              - teaser      = "Discover the wonders of our meshy service"
            },
          - {
              - description = ""
              - id          = 8
              - image       = "/boundary.png"
              - ingredients = [
                  - {
                      - id = 1
                    },
                  - {
                      - id = 6
                    },
                ]
              - name        = "Boundary Red Eye"
              - price       = 200
              - teaser      = "Perk up and watch out for your access management"
            },
          - {
              - description = ""
              - id          = 9
              - image       = "/waypoint.png"
              - ingredients = [
                  - {
                      - id = 1
                    },
                  - {
                      - id = 2
                    },
                ]
              - name        = "Waypointiato"
              - price       = 250
              - teaser      = "Deploy with a little foam"
            },
        ]
    } -> null
  - edu_order   = {
      - id           = "1"
      - items        = [
          - {
              - coffee   = {
                  - description = ""
                  - id          = 3
                  - image       = "/vault.png"
                  - name        = "Vaulatte"
                  - price       = 200
                  - teaser      = "Nothing gives you a safe and secure feeling like a Vaulatte"
                }
              - quantity = 2
            },
          - {
              - coffee   = {
                  - description = ""
                  - id          = 2
                  - image       = "/packer.png"
                  - name        = "Packer Spiced Latte"
                  - price       = 350
                  - teaser      = "Packed with goodness to spice up your images"
                }
              - quantity = 3
            },
        ]
      - last_updated = "Monday, 27-May-24 22:21:32 BST"
    } -> null
hashicups_order.edu: Destroying... [id=1]
hashicups_order.edu: Destruction complete after 0s

Destroy complete! Resources: 1 destroyed.

- terraform apply -auto-approve
terraform apply -auto-approve
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides
│ are set in the CLI configuration:
│  - hashicorp.com/edu/hashicups in /Users/yauhenbichel/go/bin
│ 
│ The behavior may therefore not match any
│ released version of the provider and applying
│ changes may cause the state to become
│ incompatible with published releases.
╵
data.hashicups_coffees.edu: Reading...
data.hashicups_coffees.edu: Read complete after 0s

Terraform used the selected providers to
generate the following execution plan. Resource
actions are indicated with the following
symbols:
  + create

Terraform will perform the following actions:

  # hashicups_order.edu will be created
  + resource "hashicups_order" "edu" {
      + id           = (known after apply)
      + items        = [
          + {
              + coffee   = {
                  + description = (known after apply)
                  + id          = 3
                  + image       = (known after apply)
                  + name        = (known after apply)
                  + price       = (known after apply)
                  + teaser      = (known after apply)
                }
              + quantity = 2
            },
          + {
              + coffee   = {
                  + description = (known after apply)
                  + id          = 2
                  + image       = (known after apply)
                  + name        = (known after apply)
                  + price       = (known after apply)
                  + teaser      = (known after apply)
                }
              + quantity = 3
            },
        ]
      + last_updated = (known after apply)
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + edu_coffees = {
      + coffees = [
          + {
              + description = ""
              + id          = 1
              + image       = "/hashicorp.png"
              + ingredients = [
                  + {
                      + id = 6
                    },
                ]
              + name        = "HCP Aeropress"
              + price       = 200
              + teaser      = "Automation in a cup"
            },
          + {
              + description = ""
              + id          = 2
              + image       = "/packer.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 2
                    },
                  + {
                      + id = 4
                    },
                ]
              + name        = "Packer Spiced Latte"
              + price       = 350
              + teaser      = "Packed with goodness to spice up your images"
            },
          + {
              + description = ""
              + id          = 3
              + image       = "/vault.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 2
                    },
                ]
              + name        = "Vaulatte"
              + price       = 200
              + teaser      = "Nothing gives you a safe and secure feeling like a Vaulatte"
            },
          + {
              + description = ""
              + id          = 4
              + image       = "/nomad.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 3
                    },
                ]
              + name        = "Nomadicano"
              + price       = 150
              + teaser      = "Drink one today and you will want to schedule another"
            },
          + {
              + description = ""
              + id          = 5
              + image       = "/terraform.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                ]
              + name        = "Terraspresso"
              + price       = 150
              + teaser      = "Nothing kickstarts your day like a provision of Terraspresso"
            },
          + {
              + description = ""
              + id          = 6
              + image       = "/vagrant.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                ]
              + name        = "Vagrante espresso"
              + price       = 200
              + teaser      = "Stdin is not a tty"
            },
          + {
              + description = ""
              + id          = 7
              + image       = "/consul.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 5
                    },
                ]
              + name        = "Connectaccino"
              + price       = 250
              + teaser      = "Discover the wonders of our meshy service"
            },
          + {
              + description = ""
              + id          = 8
              + image       = "/boundary.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 6
                    },
                ]
              + name        = "Boundary Red Eye"
              + price       = 200
              + teaser      = "Perk up and watch out for your access management"
            },
          + {
              + description = ""
              + id          = 9
              + image       = "/waypoint.png"
              + ingredients = [
                  + {
                      + id = 1
                    },
                  + {
                      + id = 2
                    },
                ]
              + name        = "Waypointiato"
              + price       = 250
              + teaser      = "Deploy with a little foam"
            },
        ]
    }
  + edu_order   = {
      + id           = (known after apply)
      + items        = [
          + {
              + coffee   = {
                  + description = (known after apply)
                  + id          = 3
                  + image       = (known after apply)
                  + name        = (known after apply)
                  + price       = (known after apply)
                  + teaser      = (known after apply)
                }
              + quantity = 2
            },
          + {
              + coffee   = {
                  + description = (known after apply)
                  + id          = 2
                  + image       = (known after apply)
                  + name        = (known after apply)
                  + price       = (known after apply)
                  + teaser      = (known after apply)
                }
              + quantity = 3
            },
        ]
      + last_updated = (known after apply)
    }
hashicups_order.edu: Creating...
hashicups_order.edu: Creation complete after 0s [id=2]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.

Outputs:

edu_coffees = {
  "coffees" = tolist([
    {
      "description" = ""
      "id" = 1
      "image" = "/hashicorp.png"
      "ingredients" = tolist([
        {
          "id" = 6
        },
      ])
      "name" = "HCP Aeropress"
      "price" = 200
      "teaser" = "Automation in a cup"
    },
    {
      "description" = ""
      "id" = 2
      "image" = "/packer.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 2
        },
        {
          "id" = 4
        },
      ])
      "name" = "Packer Spiced Latte"
      "price" = 350
      "teaser" = "Packed with goodness to spice up your images"
    },
    {
      "description" = ""
      "id" = 3
      "image" = "/vault.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 2
        },
      ])
      "name" = "Vaulatte"
      "price" = 200
      "teaser" = "Nothing gives you a safe and secure feeling like a Vaulatte"
    },
    {
      "description" = ""
      "id" = 4
      "image" = "/nomad.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 3
        },
      ])
      "name" = "Nomadicano"
      "price" = 150
      "teaser" = "Drink one today and you will want to schedule another"
    },
    {
      "description" = ""
      "id" = 5
      "image" = "/terraform.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
      ])
      "name" = "Terraspresso"
      "price" = 150
      "teaser" = "Nothing kickstarts your day like a provision of Terraspresso"
    },
    {
      "description" = ""
      "id" = 6
      "image" = "/vagrant.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
      ])
      "name" = "Vagrante espresso"
      "price" = 200
      "teaser" = "Stdin is not a tty"
    },
    {
      "description" = ""
      "id" = 7
      "image" = "/consul.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 5
        },
      ])
      "name" = "Connectaccino"
      "price" = 250
      "teaser" = "Discover the wonders of our meshy service"
    },
    {
      "description" = ""
      "id" = 8
      "image" = "/boundary.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 6
        },
      ])
      "name" = "Boundary Red Eye"
      "price" = 200
      "teaser" = "Perk up and watch out for your access management"
    },
    {
      "description" = ""
      "id" = 9
      "image" = "/waypoint.png"
      "ingredients" = tolist([
        {
          "id" = 1
        },
        {
          "id" = 2
        },
      ])
      "name" = "Waypointiato"
      "price" = 250
      "teaser" = "Deploy with a little foam"
    },
  ])
}
edu_order = {
  "id" = "2"
  "items" = tolist([
    {
      "coffee" = {
        "description" = ""
        "id" = 3
        "image" = "/vault.png"
        "name" = "Vaulatte"
        "price" = 200
        "teaser" = "Nothing gives you a safe and secure feeling like a Vaulatte"
      }
      "quantity" = 2
    },
    {
      "coffee" = {
        "description" = ""
        "id" = 2
        "image" = "/packer.png"
        "name" = "Packer Spiced Latte"
        "price" = 350
        "teaser" = "Packed with goodness to spice up your images"
      }
      "quantity" = 3
    },
  ])
  "last_updated" = "Monday, 27-May-24 22:28:03 BST"
}

- terraform show

- terraform state rm hashicups_order.edu
terraform state rm hashicups_order.edu
Removed hashicups_order.edu
Successfully removed 1 resource instance(s).
yauhenbichel@192 example % 

- terraform show 
- curl -X GET -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/2

- The interface requires the following:
A Metadata method that sets the function name. Unlike resources and data sources, function names do not start with the provider name.
A Definition method that defines the parameters, return value, and documentation for the function.
A Run method that executes the function code.

- go install .
- terraform apply -auto-approve
total_price = 5.43