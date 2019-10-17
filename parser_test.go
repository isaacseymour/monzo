package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var homeLinks = []string{
	"https://gocardless.com/", "https://gocardless.com/features/invoicing/", "https://gocardless.com/features/recurring-payments/", "https://gocardless.com/enterprise", "https://gocardless.com/partners/", "https://gocardless.com/schemes/", "https://gocardless.com/pricing/", "https://gocardless.com/stories/", "https://gocardless.com/industries/accountants", "https://gocardless.com/industries/charities", "https://gocardless.com/industries/agencies", "https://gocardless.com/industries/finance", "https://gocardless.com/industries/gyms-and-sports-clubs", "https://gocardless.com/industries/local-government", "https://gocardless.com/industries/membership-organisations", "https://gocardless.com/industries/saas", "https://gocardless.com/industries/telcos", "https://gocardless.com/industries/utilities", "https://gocardless.com/developers/", "https://support.gocardless.com/hc/en-gb", "https://gocardless.com/guides", "https://gocardless.com/faq/merchants/", "https://gocardless.com/about/", "https://gocardless.com/about/careers/", "https://gocardless.com/blog", "https://gocardless.com/partner-with-us/", "https://gocardless.com/contact-sales/", "https://manage.gocardless.com/", "https://manage.gocardless.com/signup", "https://manage.gocardless.com/signup", "https://gocardless.com/features/invoicing/", "https://gocardless.com/features/recurring-payments/", "https://gocardless.com/partners/", "https://gocardless.com/guides/posts/direct-debit-guarantee", "https://gocardless.com/stories/fd-works/", "https://gocardless.com/pricing/", "https://manage.gocardless.com/signup", "https://gocardless.com/contact-sales/", "https://www.linkedin.com/company/gocardless", "https://twitter.com/gocardless", "https://www.facebook.com/GoCardless", "https://www.youtube.com/gocardless", "https://gocardless.com/features/invoicing/", "https://gocardless.com/features/recurring-payments/", "https://gocardless.com/enterprise/", "https://gocardless.com/partners/", "https://gocardless.com/direct-debit-api/", "https://gocardless.com/pricing/", "https://support.gocardless.com/hc/en-gb", "https://gocardless.com/guides/", "https://gocardless.com/security/", "https://gocardless.com/faq/merchants/", "https://developer.gocardless.com/api-reference/", "https://gocardless.com/customer-offers/", "https://gocardless.com/about/", "https://gocardless.com/about/careers/", "https://gocardless.com/blog/", "https://gocardless.com/about/press/", "https://gocardless.com/privacy/", "https://gocardless.com/legal/", "https://gocardless.com/cookies/settings/", "https://gocardless.com/partner-with-us/", "https://gocardless.com/contact-sales/", "https://support.gocardless.com/hc/en-gb", "https://gocardless.com/payment-lookup/", "https://gocardless.com/xero/", "https://gocardless.com/en-au/", "https://gocardless.com/en-ca/", "https://gocardless.com/da-dk/", "https://gocardless.com/de-de/", "https://gocardless.com/es-es/", "https://gocardless.com/en-eu/", "https://gocardless.com/fr-fr/", "https://gocardless.com/en-ie/", "https://gocardless.com/en-nz/", "https://gocardless.com/no-no/", "https://gocardless.com/sv-se/", "https://gocardless.com/en-us/", "https://gocardless.com/",
}

func TestFindLinks(t *testing.T) {
	f, err := os.Open("fixtures/home.html")
	if err != nil {
		t.Fatalf("error opening file: %s", err)
	}

	reader := bufio.NewReader(f)
	results, err := FindLinks(reader)
	if err != nil {
		t.Fatalf("error opening file: %s", err)
	}

	assert.ElementsMatch(t, homeLinks, results)
}
