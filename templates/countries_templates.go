package templates

const GetBrazilQuery = `
query GetCountry {
  country(code: "BR") {
	name
	native
	capital
	emoji
	currency
	languages {
	  code
	  name
	}
  }
}`