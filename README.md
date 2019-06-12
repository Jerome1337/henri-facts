# Henri Facts

[![OpenFaaS](https://img.shields.io/badge/openfaas-cloud-blue.svg)](https://www.openfaas.com)

Simple golang FAAS powered by OpenFaas used as a Slack command.

This function return a random quote from [@henri9813](https://github.com/henri9813)

## Usage

Use https://jerome1337.o6s.io/henri-facts

For example

```bash
$ curl https://jerome1337.o6s.io/henri-facts
```

#### Response

```json
{
  "response_type":"in_channel",
  "text":"Il est blazing nické"
}
```

> The `response_type` key is necessary when this FAAS is used as a Slack command
