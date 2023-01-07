terraform {
  required_providers {
    infakt = {
      source  = "sequring/infakt"
      version = "0.2.3"
    }
  }
}

provider "infakt" {
  # Configuration options
  token = "123-123-123-123"
}
