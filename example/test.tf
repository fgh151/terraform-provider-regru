terraform {
  required_providers {
    regru = {
      source = "openitstudio.ru/dns/regru"
      version = "0.0.1"
    }
  }
}

provider "regru" {
  password = "passwd"
  username = "username"
}

resource "regru_dns_zone" "test_com" {
  domain  = "test.com"
}

resource "regru_dns_zone_record" "a_a" {
  zone = regru_dns_zone.test_com.domain
  host            = "A"
  type            = "TXT"
  value           = "11.22.33.44"
  ttl             = 10
  external_id     = ""
  additional_info = ""
}
