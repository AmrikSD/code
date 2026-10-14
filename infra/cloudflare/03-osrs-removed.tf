# The osrs.amrik.co.uk Pages project was deleted manually in Cloudflare;
# forget the orphaned state entries instead of trying to refresh them.
removed {
  from = cloudflare_pages_project.osrs_amrik_co_uk

  lifecycle {
    destroy = false
  }
}

removed {
  from = cloudflare_pages_domain.osrs_subdomain

  lifecycle {
    destroy = false
  }
}

removed {
  from = cloudflare_record.osrs_amrik_co_uk

  lifecycle {
    destroy = false
  }
}
