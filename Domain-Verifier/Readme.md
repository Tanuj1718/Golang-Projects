# JARGONS In this project

* Domain: You start with a domain (e.g., example.com).

* hasMX: You check if the domain has valid MX records, ensuring it can receive emails.

* hasSPF: You check if the domain has a valid SPF record, ensuring emails from the domain are coming from authorized mail servers.

* spfRecord: You retrieve the actual SPF record, which lists the authorized IP addresses or servers for the domain.

* hasDMARC: You check if the domain has a valid DMARC record, which adds an additional layer of security by defining what to do when emails fail SPF or DKIM checks.

* dmarcRecord: You retrieve the actual DMARC record, which tells receiving mail servers how to handle email failures and where to send reports.
