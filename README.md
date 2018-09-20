# WooliesX Code Challenge

The WooliesX code challenge is set to interact with and return content from specific API's with the intent to use this as an intermediate step between a website and the back end systems.

This is a very basic implementation of an API. Tests have not been added to the code currently but in any production deployment this would be a necessity.
Logging is currently only done to to the console but in reality would be consumed and pushed to a logging system such as Sumo Logic.0

## Notes

### Ways this API needs to be improved

- The user endpoint is hardcoded to a single user and a token. This should be read from some form of session request (cookie, server session etc) and returned. Tokens should never be hardcoded in an codebase.
- The specials should not be coded into the application such as they are. There should be a further endpoint or access to a DBMS of some sort used to manage available specials.
- Current special applies 2 for the price of one logic on every item in the cart.
- The use of PACT or other contracting systems would allow for any ingest API's and output endpoints to be held to account for data cleanliness.
- I would use a split layout of data: error: fields in the API response from the start so as to allow for error information to be passed back to the requesting server.
- Personally I would send the token as a http header so as to keep it from any URL caching that may occour.

### Issues with the WooliesX endpoints

- The product listing is coming back with 0 items in the quantity. 
- - This is a 200 response implying that the content returned is correct.
- - In reality one would assume that an item of zero quantity should not be included in the product listing.
- The Specials endpoint returns different content depending on response.
- - If there is a valid response the resulting data is returned with a single integer.
- - If there is an error the response is a JSON object.
- - To improve this, I would return the data under a data: tag and errors under error: so as to provide a unified response format.
- - The Specials endpoint requires one product, and one quantity object even tho there is multiple objects in the trolley.
- - The Specials endpoint requires one and only one special for any transaction. Any more and a No Content response is returned.
- I am not 100% certian but I believe that in some instances prices have been returned as int's and some as floats. This should be kept consistant.