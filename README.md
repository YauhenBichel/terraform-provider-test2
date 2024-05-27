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
