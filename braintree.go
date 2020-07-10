package connector_errors

import "github.com/chargehive/proto/golang/chargehive/chtype"

func(){
	switch code {
	case api.RCode403:
		setFailoverMerchantMessage(detail, "Unauthorized")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode1000:
		setFailoverMerchantMessage(detail, "Approved")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NONE
	case api.RCode1001:
		setFailoverMerchantMessage(detail, "Approved, check customer ID")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NONE
	case api.RCode1002:
		setFailoverMerchantMessage(detail, "Processed  - This code will be assigned to all refunds, credits, and voice authorizations. These types of transactions do not need to be authorized; they are immediately submitted for settlement.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NONE
	case api.RCode1003:
		setFailoverMerchantMessage(detail, "Approved with Risk  - The bank account can be used for transactions, but some risk has been identified (e.g. some customer information does not exactly match the bank's records).")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NONE
	case api.RCode2000:
		setFailoverMerchantMessage(detail, "Do Not Honor - The customer's bank is unwilling to accept the transaction. The customer will need to contact their bank for more details regarding this generic decline.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2001:
		setFailoverMerchantMessage(detail, "Insufficient Funds - The account did not have sufficient funds to cover the transaction amount at the time of the transaction – subsequent attempts at a later date may be successful.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_AVAILABLE_FUNDS
	case api.RCode2002:
		setFailoverMerchantMessage(detail, "Limit Exceeded - The attempted transaction exceeds the withdrawal limit of the account. The customer will need to contact their bank to change the account limits or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode2003:
		setFailoverMerchantMessage(detail, "Cardholder's Activity Limit Exceeded - The attempted transaction exceeds the activity limit of the account. The customer will need to contact their bank to change the account limits or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode2004:
		setFailoverMerchantMessage(detail, "Expired Card - Card is expired. The customer will need to use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_EXPIRED
	case api.RCode2005:
		setFailoverMerchantMessage(detail, "Invalid Credit Card Number - The customer entered an invalid payment method or made a typo in their credit card information. Have the customer correct their payment information and attempt the transaction again – if the decline persists, they will need to contact their bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_USER_INPUT
	case api.RCode2006:
		setFailoverMerchantMessage(detail, "Invalid Expiration Date - The customer entered an invalid payment method or made a typo in their card expiration date. Have the customer correct their payment information and attempt the transaction again – if the decline persists, they will need to contact their bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_USER_INPUT
	case api.RCode2007:
		setFailoverMerchantMessage(detail, "No Account - The submitted card number is not on file with the card-issuing bank. The customer will need to contact their bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNAVAILABLE
	case api.RCode2008:
		setFailoverMerchantMessage(detail, "Card Account Length Error - The submitted card number does not include the proper number of digits. Have the customer attempt the transaction again – if the decline persists, the customer will need to contact their bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_USER_INPUT
	case api.RCode2009:
		setFailoverMerchantMessage(detail, "No Such Issuer - This decline code could indicate that the submitted card number does not correlate to an existing card-issuing bank or that there is a connectivity error with the issuer. The customer will need to contact their bank for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2010:
		setFailoverMerchantMessage(detail, "Card Issuer Declined CVV - The customer entered in an invalid security code or made a typo in their card information. Have the customer attempt the transaction again – if the decline persists, the customer will need to contact their bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_CVV
	case api.RCode2011:
		setFailoverMerchantMessage(detail, "Voice Authorization Required - The customer’s bank is requesting that the merchant (you) call to obtain a special authorization code in order to complete this transaction. This can result in a lengthy process – we recommend obtaining a new payment method instead. Contact us for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2012:
		setFailoverMerchantMessage(detail, "Processor Declined – Possible Lost Card - The card used has likely been reported as lost. The customer will need to contact their bank for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_LOST
	case api.RCode2013:
		setFailoverMerchantMessage(detail, "Processor Declined – Possible Stolen Card - The card used has likely been reported as stolen. The customer will need to contact their bank for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_LOST
	case api.RCode2014:
		setFailoverMerchantMessage(detail, "Processor Declined – Fraud Suspected - The customer’s bank suspects fraud – they will need to contact their bank for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
		detail.ErrorType = chtype.RESPONSE_ERROR_FRAUD
	case api.RCode2015:
		setFailoverMerchantMessage(detail, "Transaction Not Allowed - The customer's bank is declining the transaction for unspecified reasons, possibly due to an issue with the card itself. They will need to contact their bank or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2016:
		setFailoverMerchantMessage(detail, "Duplicate Transaction - The submitted transaction appears to be a duplicate of a previously submitted transaction and was declined to prevent charging the same card twice for the same service.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_DUPLICATE
	case api.RCode2017:
		setFailoverMerchantMessage(detail, "Cardholder Stopped Billing - The customer requested a cancellation of a single transaction – reach out to them for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2018:
		setFailoverMerchantMessage(detail, "Cardholder Stopped All Billing - The customer requested the cancellation of a recurring transaction or subscription – reach out to them for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2019:
		setFailoverMerchantMessage(detail, "Invalid Transaction - The customer’s bank declined the transaction, typically because the card in question does not support this type of transaction – for example, the customer used an FSA debit card for a non-healthcare related purchase. They will need to contact their bank for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2020:
		setFailoverMerchantMessage(detail, "Violation - The customer will need to contact their bank for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode2021:
		setFailoverMerchantMessage(detail, "Security Violation - The customer's bank is declining the transaction, possibly due to a fraud concern. They will need to contact their bank or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2022:
		setFailoverMerchantMessage(detail, "Declined – Updated Cardholder Available - The submitted card has expired or been reported lost and a new card has been issued. Reach out to your customer to obtain updated card information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_EXPIRED
	case api.RCode2023:
		setFailoverMerchantMessage(detail, "Processor Does Not Support This Feature - Your account can't process transactions with the intended feature – for example, 3D Secure or Level 2/Level 3 data. If you believe your merchant account should be set up to accept this type of transaction, contact us.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2024:
		setFailoverMerchantMessage(detail, "Card Type Not Enabled - Your account can't process the attempted card type. If you believe your merchant account should be set up to accept this type of card, contact us for assistance.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2025:
		setFailoverMerchantMessage(detail, "Set Up Error – Merchant - Depending on your region, this response could indicate a connectivity or setup issue. Contact us for more information regarding this error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2026:
		setFailoverMerchantMessage(detail, "Invalid Merchant ID - The customer’s bank declined the transaction, typically because the card in question does not support this type of transaction. If this response persists across transactions for multiple customers, it could indicate a connectivity or setup issue. Contact us for more information regarding this error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2027:
		setFailoverMerchantMessage(detail, "Set Up Error – Amount - This rare decline code indicates an issue with processing the amount of the transaction. The customer will need to contact their bank for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2028:
		setFailoverMerchantMessage(detail, "Set Up Error – Hierarchy - There is a setup issue with your account. Contact us for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2029:
		setFailoverMerchantMessage(detail, "Set Up Error – Card - This response generally indicates that there is a problem with the submitted card. The customer will need to use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2030:
		setFailoverMerchantMessage(detail, "Set Up Error – Terminal - There is a setup issue with your account. Contact us for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2031:
		setFailoverMerchantMessage(detail, "Encryption Error - The cardholder’s bank does not support $0.00 card verifications. Enable the Retry All Failed $0option to resolve this error. Contact us with questions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2032:
		setFailoverMerchantMessage(detail, "Surcharge Not Permitted - Surcharge amount not permitted on this card. The customer will need to use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2033:
		setFailoverMerchantMessage(detail, "Inconsistent Data - An error occurred when communicating with the processor. The customer will need to contact their bank for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
		detail.ErrorType = chtype.RESPONSE_ERROR_UNAVAILABLE
	case api.RCode2034:
		setFailoverMerchantMessage(detail, "No Action Taken - An error occurred and the intended transaction was not completed. Attempt the transaction again.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode2035:
		setFailoverMerchantMessage(detail, "Partial Approval For Amount In Group III Version - The customer's bank approved the transaction for less than the requested amount. Have the customer attempt the transaction again – if the decline persists, the customer will need to use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode2036:
		setFailoverMerchantMessage(detail, "Authorization could not be found - An error occurred when trying to process the authorization. This response could indicate an issue with the customer’s card or that the processor doesn't allow this action – contact us for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2037:
		setFailoverMerchantMessage(detail, "Already Reversed - The indicated authorization has already been reversed. If you believe this to be false, contact usfor more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_DUPLICATE
	case api.RCode2038:
		setFailoverMerchantMessage(detail, "Processor Declined - The customer's bank is unwilling to accept the transaction. The reasons for this response can vary – customer will need to contact their bank for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2039:
		setFailoverMerchantMessage(detail, "Invalid Authorization Code - The authorization code was not found or not provided. Have the customer attempt the transaction again – if the decline persists, they will need to contact their bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode2040:
		setFailoverMerchantMessage(detail, "Invalid Store - There may be an issue with the configuration of your account. Have the customer attempt the transaction again – if the decline persists, contact us for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2041:
		setFailoverMerchantMessage(detail, "Declined – Call For Approval - The card used for this transaction requires customer approval – they will need to contact their bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2042:
		setFailoverMerchantMessage(detail, "Invalid Client ID - There may be an issue with the configuration of your account. Have the customer attempt the transaction again – if the decline persists, contact us for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2043:
		setFailoverMerchantMessage(detail, "Error – Do Not Retry, Call Issuer - The card-issuing bank will not allow this transaction. The customer will need to contact their bank for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2044:
		setFailoverMerchantMessage(detail, "Declined – Call Issuer - The card-issuing bank has declined this transaction. Have the customer attempt the transaction again – if the decline persists, they will need to contact their bank for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2045:
		setFailoverMerchantMessage(detail, "Invalid Merchant Number - There is a setup issue with your account. Contact us for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2046:
		setFailoverMerchantMessage(detail, "Declined - The customer's bank is unwilling to accept the transaction. For credit/debit card transactions, the customer will need to contact their bank for more details regarding this generic decline; if this is a PayPal transaction, the customer will need to contact PayPal.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2047:
		setFailoverMerchantMessage(detail, "Call Issuer. Pick Up Card - The customer’s card has been reported as lost or stolen by the cardholder and the card-issuing bank has requested that merchants keep the card and call the number on the back to report it. As an online merchant, you don’t have the physical card and can't complete this request – obtain a different payment method from the customer.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PICKUP
	case api.RCode2048:
		setFailoverMerchantMessage(detail, "Invalid Amount - The authorized amount is set to zero, is unreadable, or exceeds the allowable amount. Make sure the amount is greater than zero and in a suitable format.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2049:
		setFailoverMerchantMessage(detail, "Invalid SKU Number - A non-numeric value was sent with the attempted transaction. Fix errors and resubmit with the transaction with the proper SKU Number.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2050:
		setFailoverMerchantMessage(detail, "Invalid Credit Plan - There may be an issue with the customer’s card or a temporary issue at the card-issuing bank. The customer will need to contact their bank for more information or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2051:
		setFailoverMerchantMessage(detail, "Credit Card Number does not match method of payment - There may be an issue with the customer’s credit card or a temporary issue at the card-issuing bank. Have the customer attempt the transaction again – if the decline persists, ask for a different card or payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2053:
		setFailoverMerchantMessage(detail, "Card reported as lost or stolen - The card used was reported lost or stolen. The customer will need to contact their bank for more information or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_LOST
	case api.RCode2054:
		setFailoverMerchantMessage(detail, "Reversal amount does not match authorization amount - Either the refund amount is greater than the original transaction or the card-issuing bank does not allow partial refunds. The customer will need to contact their bank for more information or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2055:
		setFailoverMerchantMessage(detail, "Invalid Transaction Division Number - Contact us for more information regarding this error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2056:
		setFailoverMerchantMessage(detail, "Transaction amount exceeds the transaction division limit - Contact us for more information regarding this error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2057:
		setFailoverMerchantMessage(detail, "Issuer or Cardholder has put a restriction on the card - The customer will need to contact their issuing bank for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_LOST
	case api.RCode2058:
		setFailoverMerchantMessage(detail, "Merchant not Mastercard SecureCode enabled - The attempted card can't be processed without enabling 3D Secure for your account. Contact us for more information regarding this feature or contact the customer for a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2059:
		setFailoverMerchantMessage(detail, "Address Verification Failed - PayPal was unable to verify that the transaction qualifies for Seller Protection because the address was improperly formatted. The customer should contact PayPal for more information or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_ADDRESS
	case api.RCode2060:
		setFailoverMerchantMessage(detail, "Address Verification and Card Security Code Failed - Both the AVS and CVV checks failed for this transaction. The customer should contact PayPal for more information or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_ADDRESS
	case api.RCode2061:
		setFailoverMerchantMessage(detail, "Invalid Transaction Data - There may be an issue with the customer’s card or a temporary issue at the card-issuing bank. Have the customer attempt the transaction again – if the decline persists, ask for a different card or payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2062:
		setFailoverMerchantMessage(detail, "Invalid Tax Amount - There may be an issue with the customer’s card or a temporary issue at the card-issuing bank. Have the customer attempt the transaction again – if the decline persists, ask for a different card or payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2063:
		setFailoverMerchantMessage(detail, "PayPal Business Account preference resulted in the transaction failing - You can't process this transaction because your account is set to block certain payment types, such as eChecks or foreign currencies. If you believe you have received this decline in error, contact us.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode2064:
		setFailoverMerchantMessage(detail, "Invalid Currency Code - There may be an issue with the configuration of your account for the currency specified. Contact usfor more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_ADDRESS
	case api.RCode2065:
		setFailoverMerchantMessage(detail, "Refund Time Limit Exceeded - PayPal requires that refunds are issued within 180 days of the sale. This refund can't be successfully processed.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode2066:
		setFailoverMerchantMessage(detail, "PayPal Business Account Restricted - Contact PayPal’s Support team to resolve this issue with your account. Then, you can attempt the transaction again.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode2067:
		setFailoverMerchantMessage(detail, "Authorization Expired - The PayPal authorization is no longer valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_EXPIRED
	case api.RCode2068:
		setFailoverMerchantMessage(detail, "PayPal Business Account Locked or Closed - You'll need to contact PayPal’s Support team to resolve an issue with your account. Once resolved, you can attempt to process the transaction again.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNAVAILABLE
	case api.RCode2069:
		setFailoverMerchantMessage(detail, "PayPal Blocking Duplicate Order IDs - The submitted PayPal transaction appears to be a duplicate of a previously submitted transaction. This decline code indicates an attempt to prevent charging the same PayPal account twice for the same service.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PROCESSING
		detail.ErrorType = chtype.RESPONSE_ERROR_DUPLICATE
	case api.RCode2070:
		setFailoverMerchantMessage(detail, "PayPal Buyer Revoked Pre-Approved Payment Authorization - The customer revoked authorization for this payment method. Reach out to the customer for more information or a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2071:
		setFailoverMerchantMessage(detail, "PayPal Payee Account Invalid Or Does Not Have a Confirmed Email - Customer has not finalized setup of their PayPal account. Reach out to the customer for more information or a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_USER_INPUT
	case api.RCode2072:
		setFailoverMerchantMessage(detail, "PayPal Payee Email Incorrectly Formatted - Customer made a typo or is attempting to use an invalid PayPal account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_USER_INPUT
	case api.RCode2073:
		setFailoverMerchantMessage(detail, "PayPal Validation Error - PayPal can't validate this transaction. Contact us for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2074:
		setFailoverMerchantMessage(detail, "Funding Instrument In The PayPal Account Was Declined By The Processor Or Bank, Or It Can't Be Used For This Payment - The customer’s payment method associated with their PayPal account was declined. Reach out to the customer for more information or a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2075:
		setFailoverMerchantMessage(detail, "Payer Account Is Locked Or Closed - The customer’s PayPal account can't be used for transactions at this time. The customer will need to contact PayPal for more information or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNAVAILABLE
	case api.RCode2076:
		setFailoverMerchantMessage(detail, "Payer Cannot Pay For This Transaction With PayPal - The customer should contact PayPal for more information or use a different payment method. You may also receive this response if you create transactions using the email address registered with your PayPal Business Account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2077:
		setFailoverMerchantMessage(detail, "Transaction Refused Due To PayPal Risk Model - PayPal has declined this transaction due to risk limitations. You'll need to contact PayPal’s Support team to resolve this issue.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2079:
		setFailoverMerchantMessage(detail, "PayPal Merchant Account Configuration Error - You'll need to contact us to resolve an issue with your account. Once resolved, you can attempt to process the transaction again.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2081:
		setFailoverMerchantMessage(detail, "PayPal pending payments are not supported - Braintree received an unsupported Pending Verification response from PayPal. Contact Braintree’s Support team to resolve a potential issue with your account settings. If there is no issue with your account, have the customer reach out to PayPal for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2082:
		setFailoverMerchantMessage(detail, "PayPal Domestic Transaction Required - This transaction requires the customer to be a resident of the same country as the merchant. Reach out to the customer for a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2083:
		setFailoverMerchantMessage(detail, "PayPal Phone Number Required - This transaction requires the payer to provide a valid phone number. The customer should contact PayPal for more information or use a different payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2084:
		setFailoverMerchantMessage(detail, "PayPal Tax Info Required - The customer must complete their PayPal account information, including submitting their phone number and all required tax information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2085:
		setFailoverMerchantMessage(detail, "PayPal Payee Blocked Transaction - Fraud settings on your PayPal business account are blocking payments from this customer. These can be adjusted in the Block Payments section of your PayPal business account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2086:
		setFailoverMerchantMessage(detail, "PayPal Transaction Limit Exceeded - The settings on the customer's account do not allow a transaction amount this large. They will need to contact PayPal to resolve this issue.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2087:
		setFailoverMerchantMessage(detail, "PayPal reference transactions not enabled for your account - PayPal API permissions are not set up to allow reference transactions. You'll need to contact PayPal’s Support team to resolve an issue with your account. Once resolved, you can attempt to process the transaction again.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2088:
		setFailoverMerchantMessage(detail, "Currency not enabled for your PayPal seller account - This currency is not currently supported by your PayPal account. You can accept additional currencies by updating your PayPal profile.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2089:
		setFailoverMerchantMessage(detail, "PayPal payee email permission denied for this request - PayPal API permissions are not set up between your PayPal business accounts. Contact us for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2090:
		setFailoverMerchantMessage(detail, "PayPal account not configured to refund more than settled amount - Your PayPal account is not set up to refund amounts higher than the original transaction amount. Contact PayPal's Support team for information on how to enable this.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2091:
		setFailoverMerchantMessage(detail, "Currency of this transaction must match currency of your PayPal account - Your PayPal account can only process transactions in the currency of your home country. Contact PayPal's Support team for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2092:
		setFailoverMerchantMessage(detail, "No Data Found - Try Another Verification Method - The processor is unable to provide a definitive answer about the customer's bank account. Please try a different US bank account verification method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2093:
		setFailoverMerchantMessage(detail, "PayPal payment method is invalid - The PayPal payment method has either expired or been canceled.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2094:
		setFailoverMerchantMessage(detail, "PayPal payment has already been completed - Your integration is likely making PayPal calls out of sequence. Check the error response for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_ALREADY_DONE
	case api.RCode2095:
		setFailoverMerchantMessage(detail, "PayPal refund is not allowed after partial refund - Once a PayPal transaction is partially refunded, all subsequent refunds must also be partial refunds for the remaining amount or less. Full refunds are not allowed after a PayPal transaction has been partially refunded.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2096:
		setFailoverMerchantMessage(detail, "PayPal buyer account can't be the same as the seller account - PayPal buyer account can't be the same as the seller account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode2097:
		setFailoverMerchantMessage(detail, "PayPal authorization amount limit exceeded - PayPal authorization amount is greater than the allowed limit on the order.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode2098:
		setFailoverMerchantMessage(detail, "PayPal authorization count limit exceeded - The number of PayPal authorizations is greater than the allowed number on the order.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode2099:
		setFailoverMerchantMessage(detail, "Cardholder Authentication Required - The customer's bank declined the transaction because a 3D Secure authentication was not performed during checkout. Have the customer authenticate using 3D Secure, then attempt the authorization again.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode2100:
		setFailoverMerchantMessage(detail, "PayPal channel initiated billing not enabled for your account - Your PayPal permissions are not set up to allow channel initiated billing transactions. Contact PayPal's Support team for information on how to enable this. Once resolved, you can attempt to process the transaction again.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode3000:
		setFailoverMerchantMessage(detail, "Processor Network Unavailable – Try Again - This response could indicate a problem with the back-end processing network, not necessarily a problem with the payment method. Have the customer attempt the transaction again – if the decline persists, contact us for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
		detail.ErrorType = chtype.RESPONSE_ERROR_UNAVAILABLE
	case api.RCode4001:
		setFailoverMerchantMessage(detail, "Settlement Declined - The processor declined to settle the sale or refund request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4003:
		setFailoverMerchantMessage(detail, "Already Captured - The transaction has already been fully captured.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_ALREADY_DONE
	case api.RCode4004:
		setFailoverMerchantMessage(detail, "Already Refunded - The transaction has already been fully refunded.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_ALREADY_DONE
	case api.RCode4005:
		setFailoverMerchantMessage(detail, "PayPal Risk Rejected - The sale request was rejected by PayPal risk.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4006:
		setFailoverMerchantMessage(detail, "Capture Amount Exceeded Allowable Limit - The specified capture amount exceeded the amount allowed by the processor.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4018:
		setFailoverMerchantMessage(detail, "PayPal Pending Payments Not Supported - PayPal returned a pending sale or refund response which is disallowed by Braintree. This failure is likely due to a misconfiguration in your PayPal account. Further details may be found in the transaction details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4019:
		setFailoverMerchantMessage(detail, "PayPal Refund Transaction with an Open Case Not Allowed - PayPal declined to settle the refund request as there is an open dispute against the transaction. If you have enabled PayPal disputes within Braintree, you may resolve the dispute within the Braintree disputes dashboard. Otherwise, you may do so via your PayPal account's Resolution Center.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4020:
		setFailoverMerchantMessage(detail, "PayPal Refund Attempt Limit Reached - PayPal's maximum number of refund attempts for this transaction has been exceeded.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode4021:
		setFailoverMerchantMessage(detail, "PayPal Refund Transaction Not Allowed - PayPal does not allow you to refund this type of transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4022:
		setFailoverMerchantMessage(detail, "PayPal Refund Invalid Partial Amount - The partial refund amount is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4023:
		setFailoverMerchantMessage(detail, "PayPal Refund Merchant Account Missing ACH - Your PayPal account does not have an associated verified bank account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81501:
		setFailoverMerchantMessage(detail, "Amount cannot be negative. - Even if creating a credit transaction, the amount should be given as x.xx, not -x.xx.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81502:
		setFailoverMerchantMessage(detail, "Amount is required. - You'll get this error if amount is not given or is blank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81503:
		setFailoverMerchantMessage(detail, "Amount is an invalid format. - Amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the amount cannot include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81509:
		setFailoverMerchantMessage(detail, "Credit card type is not accepted by this merchant account. - The credit card card type must be accepted by your merchant account. Note that there is a different error code when you get this error when creating transactions using tokens (91517).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode81520:
		setFailoverMerchantMessage(detail, "Processor authorization code must be 6 characters. - Processor authorization code must be 6 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81527:
		setFailoverMerchantMessage(detail, "Custom field is too long: - Custom field values must be less than 255 characters. The error message for this validation error will contain a list of the custom fields that were too long.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81528:
		setFailoverMerchantMessage(detail, "Amount is too large. - The amount cannot be greater than the maximum allowed by the processor. Contact us for your specific limit.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81531:
		setFailoverMerchantMessage(detail, "Amount must be greater than zero. - The amount of a transaction cannot be zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81534:
		setFailoverMerchantMessage(detail, "Tax amount cannot be negative. - The tax amount cannot be less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81535:
		setFailoverMerchantMessage(detail, "Tax amount is an invalid format. - Tax amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the amount cannot include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81536:
		setFailoverMerchantMessage(detail, "Tax amount is too large. - The tax amount cannot be longer than 9 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81571:
		setFailoverMerchantMessage(detail, "Failed to authenticate, please try a different form of payment. - Failed to authenticate, please try a different form of payment.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode81601:
		setFailoverMerchantMessage(detail, "Company is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81603:
		setFailoverMerchantMessage(detail, "Custom field is too long: - Custom field values must be less than or equal to 255 characters. The error message for this validation error will contain a list of the custom fields that were too long.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81604:
		setFailoverMerchantMessage(detail, "Email is an invalid format. - Email must be a well-formed email address. If you are migrating from a system that does not have this constraint and want to record the email address in the Vault, you can use custom_fields")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81605:
		setFailoverMerchantMessage(detail, "Email is too long. - Maximum 254 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81606:
		setFailoverMerchantMessage(detail, "Email is required if sending a receipt. - This only applies when creating a transaction. If you specify that you want to send a receipt then the customer email will be required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81607:
		setFailoverMerchantMessage(detail, "Fax is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81608:
		setFailoverMerchantMessage(detail, "First name is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81613:
		setFailoverMerchantMessage(detail, "Last name is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81614:
		setFailoverMerchantMessage(detail, "Phone is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81615:
		setFailoverMerchantMessage(detail, "Website is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81616:
		setFailoverMerchantMessage(detail, "Website is an invalid format. - Website must be well-formed. The http:// at the beginning is optional. If you want to provide websites that may be not well-formed you can use custom_fields")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81703:
		setFailoverMerchantMessage(detail, "Credit card type is not accepted by this merchant account. - Applies when specifying a credit card in a sale or verification request. Not applicable when only storing in the Vault, since Vault records are not associated to specific merchant accounts.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode81706:
		setFailoverMerchantMessage(detail, "CVV is required. - CVV will only be required if CVV processing rules are configured to require it. If the rules are configured to require it, then CVV is required when storing a card in the Vault and performing card verification or when creating transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_CVV
	case api.RCode81707:
		setFailoverMerchantMessage(detail, "CVV must be 4 digits for American Express and 3 digits for other card types. - CVV must be 4 digits for American Express and 3 digits for other card types.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_CVV
	case api.RCode81709:
		setFailoverMerchantMessage(detail, "Expiration date is required. - You must provide the expiration date either as a single field or as month and year separately.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81710:
		setFailoverMerchantMessage(detail, "Expiration date is invalid. - Valid formats are M/YY, M/YYYY, MM/YY, and MM/YYYY. The month must be 1-12 or 01-12.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81711:
		setFailoverMerchantMessage(detail, "Expiration year is invalid. It must be between 1975 and 2201. - The expiration year must be greater than 1975 and less than 2201.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81712:
		setFailoverMerchantMessage(detail, "Expiration month is invalid. - It must be 1-12 or 01-12.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81713:
		setFailoverMerchantMessage(detail, "Expiration year is invalid. - It must be between 1975 and 2201.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81714:
		setFailoverMerchantMessage(detail, "Credit card number is required. - You'll get this error if number is omitted or if it is an empty string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81715:
		setFailoverMerchantMessage(detail, "Credit card number is invalid. - The credit card number must pass a Luhn-10 check.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81716:
		setFailoverMerchantMessage(detail, "Credit card number must be 12-19 digits. - Inclusive.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81717:
		setFailoverMerchantMessage(detail, "Credit card number is not an accepted test number. - Only test numbers can be used in the sandbox.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81718:
		setFailoverMerchantMessage(detail, "Credit card number cannot be updated to an unsupported card type when it is associated to subscriptions. - Only applies when using recurring billing. If a credit card is being used for recurring billing subscriptions, the card can only be updated to a card type that is accepted by the merchant account that is being used for the subscriptions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81723:
		setFailoverMerchantMessage(detail, "Cardholder name is too long. - Maximum 175 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81724:
		setFailoverMerchantMessage(detail, "Duplicate card exists in the vault. - Duplicate card exists in the vault.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_DUPLICATE
	case api.RCode81725:
		setFailoverMerchantMessage(detail, "Credit card must include number, payment_method_nonce, or venmo_sdk_payment_method_code. - Credit card must include number, payment_method_nonce, or venmo_sdk_payment_method_code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81736:
		setFailoverMerchantMessage(detail, "CVV verification failed. - CVV was incorrect or not supplied.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_CVV
	case api.RCode81737:
		setFailoverMerchantMessage(detail, "Postal code verification failed. - Postal code was incorrect or not supplied.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_ADDRESS
	case api.RCode81750:
		setFailoverMerchantMessage(detail, "Credit card number is prohibited. - Cannot transact with an issuer on OFAC's prohibited list.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode81801:
		setFailoverMerchantMessage(detail, "Addresses must have at least one field filled in. - At least one of the address attributes must be present, but it doesn't matter which one. This doesn't apply when creating transactions—billing and shipping address can be blank unless AVS processing rules are configured to require billing street and postal.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81802:
		setFailoverMerchantMessage(detail, "Company is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81804:
		setFailoverMerchantMessage(detail, "Extended address is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81805:
		setFailoverMerchantMessage(detail, "First name is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81806:
		setFailoverMerchantMessage(detail, "Last name is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81807:
		setFailoverMerchantMessage(detail, "Locality is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81808:
		setFailoverMerchantMessage(detail, "Postal code is required. - Applies when AVS rules are configured to require postal code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81809:
		setFailoverMerchantMessage(detail, "Postal code may contain no more than 9 letter or number characters. - The length only applies to letters or numbers; it ignores spaces, hyphens, and all other special characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81810:
		setFailoverMerchantMessage(detail, "Region is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81811:
		setFailoverMerchantMessage(detail, "Street address is required. - Applies when creating a transaction or performing card verification when AVS rules are configured to require street address.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81812:
		setFailoverMerchantMessage(detail, "Street address is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81813:
		setFailoverMerchantMessage(detail, "Postal code can only contain letters, numbers, spaces, and hyphens. - Postal code must begin with a letter or number, and can only contain letters, numbers, spaces, and hyphens.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81827:
		setFailoverMerchantMessage(detail, "US state codes must be two characters to meet PayPal Seller Protection requirements. - US state codes must be two characters to meet PayPal Seller Protection requirements.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81828:
		setFailoverMerchantMessage(detail, "Postal code is required for the card type and processor. - The processor you are using requires a postal code to be included for transactions on this card brand.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81901:
		setFailoverMerchantMessage(detail, "Cannot edit a canceled subscription. - After a subscription has been canceled it cannot be updated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode81902:
		setFailoverMerchantMessage(detail, "ID has already been taken. - Subscription IDs need to be unique.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_DUPLICATE
	case api.RCode81903:
		setFailoverMerchantMessage(detail, "Price cannot be blank. - If you provide a price, it can't be an empty string. If you omit the price, then the subscription will inherit the price from the plan.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81904:
		setFailoverMerchantMessage(detail, "Price is an invalid format. - Price must be formatted like '10' or '10.00'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81905:
		setFailoverMerchantMessage(detail, "Subscription has already been canceled. - You can't cancel subscriptions that have already been canceled.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81906:
		setFailoverMerchantMessage(detail, "ID is invalid (use only letters, numbers, '-', and ''). - If specifying the ID for the subscription, you can only use letters, numbers, and -.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81907:
		setFailoverMerchantMessage(detail, "Trial Duration is an invalid format. - It must be 1-3 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81908:
		setFailoverMerchantMessage(detail, "Trial Duration is required. - It's required if trial period is set to true.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81909:
		setFailoverMerchantMessage(detail, "Trial Duration Unit is invalid. - Valid values are \"day\" and \"month\".")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode81910:
		setFailoverMerchantMessage(detail, "Cannot edit an expired subscription. - You cannot edit a subscription with Expired status.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode81923:
		setFailoverMerchantMessage(detail, "Price is too large. - The subscription price cannot be greater than the maximum allowed by the processor. Contact us for your specific limit.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82301:
		setFailoverMerchantMessage(detail, "Settlement Date is required - You must provide a settlement date as the first argument.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82302:
		setFailoverMerchantMessage(detail, "Settlement Date is invalid - The settlement date provided must be a valid date.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82303:
		setFailoverMerchantMessage(detail, "Group By Custom Field is not a valid custom field - The custom field provided must be valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82602:
		setFailoverMerchantMessage(detail, "Applicant merchant id is too long. - The merchant account ID cannot be longer than 32 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82603:
		setFailoverMerchantMessage(detail, "Applicant merchant id format is invalid. - You can only use letters, numbers, _ and - for the merchant account ID.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82604:
		setFailoverMerchantMessage(detail, "Applicant merchant id is in use. - Each merchant account ID needs to be unique.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82605:
		setFailoverMerchantMessage(detail, "Applicant merchant id is not allowed. - A merchant account ID may not be named 'all' or 'new.'")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82606:
		setFailoverMerchantMessage(detail, "Master merchant account ID is required. - You must provide a master merchant account ID when creating a merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82607:
		setFailoverMerchantMessage(detail, "Master merchant account ID is invalid. - You'll get this error if we cannot find a master merchant account with the id specified.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82608:
		setFailoverMerchantMessage(detail, "Master merchant account must be active. - You'll get this error if the supplied master merchant account ID is not active.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82609:
		setFailoverMerchantMessage(detail, "Applicant first name is required. - You must provide the first name of the applicant.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82610:
		setFailoverMerchantMessage(detail, "Terms Of Service needs to be accepted. Applicant tos_accepted required. - You must indicate that the terms of service are accepted.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode82611:
		setFailoverMerchantMessage(detail, "Applicant last name is required. - You must provide the last name of the applicant.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82612:
		setFailoverMerchantMessage(detail, "Applicant date of birth is required. - You must provide the applicant's date of birth.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82613:
		setFailoverMerchantMessage(detail, "Applicant routing number is required. - You must provide the applicant's bank routing number.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82614:
		setFailoverMerchantMessage(detail, "Applicant account number is required. - You must provide the applicant's bank account number.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82615:
		setFailoverMerchantMessage(detail, "Applicant SSN must be blank, last 4 digits, or full 9 digits. - The applicant's social security number must be valid (full 9 digits, with or without dashes, or last 4 digits).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82616:
		setFailoverMerchantMessage(detail, "Applicant email is invalid. - The applicant's email must be valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82617:
		setFailoverMerchantMessage(detail, "Applicant street address is required. - You must provide the applicant's street address.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82618:
		setFailoverMerchantMessage(detail, "Applicant locality is required. - You must provide the applicant's city, town, or municipality.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82619:
		setFailoverMerchantMessage(detail, "Applicant postal code is required. - You must provide the applicant's postal code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82620:
		setFailoverMerchantMessage(detail, "Applicant region is required. - You must provide the applicant's region.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82621:
		setFailoverMerchantMessage(detail, "Applicant declined due to OFAC. - The applicant has failed an OFAC check. The OFAC search confirms whether a sub-merchant is on the criminal and terrorists watch lists collected from databases around the world.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82622:
		setFailoverMerchantMessage(detail, "Applicant declined due to MasterCard MATCH. - The applicant has failed a Mastercard MATCH check. The Mastercard MATCH File is a database file used by payment processing banks to identify specific merchants and principals who may been terminated for reasons like fraud or violation(s) of Visa and/or Mastercard rules.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode82623:
		setFailoverMerchantMessage(detail, "Applicant declined due to failed KYC. - The applicant has failed a Know Your Customer check.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode82624:
		setFailoverMerchantMessage(detail, "Applicant declined due to invalid SSN. - The applicant's social security number is invalid. If you provide a social security number, you must provide either the entire number or the last four digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode82625:
		setFailoverMerchantMessage(detail, "Applicant declined due to SSN matching that of a deceased person. - The applicant has been declined because the social security number provided appears in a database of social security numbers belonging to deceased persons.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode82626:
		setFailoverMerchantMessage(detail, "Applicant declined after review. - After review, the applicant has been declined.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode82627:
		setFailoverMerchantMessage(detail, "Applicant first name is invalid. - The applicant's first name must not contain '/', '\\', '&', '<', '>' or any control characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82628:
		setFailoverMerchantMessage(detail, "Applicant last name is invalid. - The applicant's last name must not contain '/', '\\', '&', '<', '>' or any control characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82629:
		setFailoverMerchantMessage(detail, "Applicant street address is invalid. - You must provide a valid street address for the applicant that includes at least one digit.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82630:
		setFailoverMerchantMessage(detail, "Applicant postal code is invalid. - You must provide a valid postal code for the applicant.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82631:
		setFailoverMerchantMessage(detail, "Applicant company name is invalid. - The applicant's company name must contain only letters, numbers, and these characters: &-!@#$()'./+,\". The maximum length is 40 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82632:
		setFailoverMerchantMessage(detail, "Applicant tax ID is invalid. - The applicant's tax id must be 9 digits long.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82633:
		setFailoverMerchantMessage(detail, "Applicant company name is required with tax ID. - If the applicant's tax id is provided then the company name must be provided as well.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82634:
		setFailoverMerchantMessage(detail, "Applicant tax ID is required with company name. - If the applicant's company name is provided then the tax id must be provided as well.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82635:
		setFailoverMerchantMessage(detail, "Applicant routing number is invalid. - The applicant's bank routing number must be valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82636:
		setFailoverMerchantMessage(detail, "Applicant phone is invalid. - The provided phone is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82637:
		setFailoverMerchantMessage(detail, "Individual first name is required. - Applicant first name is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82638:
		setFailoverMerchantMessage(detail, "Individual last name is required. - Applicant last name is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82639:
		setFailoverMerchantMessage(detail, "Individual date of birth is required. - Individual date of birth is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82640:
		setFailoverMerchantMessage(detail, "Funding routing number is required. - Funding routing number is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82641:
		setFailoverMerchantMessage(detail, "Funding account number is required. - Funding account number is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82642:
		setFailoverMerchantMessage(detail, "Individual SSN must be blank, last 4 digits, or full 9 digits. - Individual SSN must be blank, last 4 digits, or full 9 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82643:
		setFailoverMerchantMessage(detail, "Individual email is invalid. - Individual email is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82644:
		setFailoverMerchantMessage(detail, "Individual first name is invalid. - Individual first name is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82645:
		setFailoverMerchantMessage(detail, "Individual last name is invalid. - Applicant last name is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82646:
		setFailoverMerchantMessage(detail, "Business DBA name is invalid. - The provided DBA name is not valid. The applicant's company name must contain only letters, numbers, and these characters: &-!@#$()'./+,\". The maximum length is 40 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82647:
		setFailoverMerchantMessage(detail, "Business tax ID is invalid. - Business tax ID is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82648:
		setFailoverMerchantMessage(detail, "Business tax ID is required with business legal name. - You must provide a tax id if a legal name has been provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82649:
		setFailoverMerchantMessage(detail, "Funding routing number is invalid. - Funding routing number is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82650:
		setFailoverMerchantMessage(detail, "An unexpected error occurred trying to save the merchant account; support has been notified and is looking into the issue. You may safely retry this request - An unexpected error occurred trying to save the merchant account; Support has been notified and is looking into the issue. You may safely retry this request.")
		detail.FailureType = chtype.FAILURE_TYPE_RETRY
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode82656:
		setFailoverMerchantMessage(detail, "Individual phone is invalid. - Individual phone is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82657:
		setFailoverMerchantMessage(detail, "Individual street address is required. - Individual street address is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82658:
		setFailoverMerchantMessage(detail, "Individual locality is required. - Individual locality is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82659:
		setFailoverMerchantMessage(detail, "Individual postal code is required. - Individual postal code is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82660:
		setFailoverMerchantMessage(detail, "Individual region is required. - Individual region is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82661:
		setFailoverMerchantMessage(detail, "Individual street address is invalid. - Individual street address is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82662:
		setFailoverMerchantMessage(detail, "Individual postal code is invalid. - Individual postal code is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82663:
		setFailoverMerchantMessage(detail, "Applicant date of birth is invalid - You must provide a valid date of birth.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82664:
		setFailoverMerchantMessage(detail, "Applicant region is invalid. - You must provide a valid region for the applicant. Only two-letter abbreviations are accepted, e.g. 'CA' but not 'California.'")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82665:
		setFailoverMerchantMessage(detail, "Applicant email is required. - You must provide an email address.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82666:
		setFailoverMerchantMessage(detail, "Individual date of birth is invalid. - Individual date of birth is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82667:
		setFailoverMerchantMessage(detail, "Individual email is required. - Individual email is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82668:
		setFailoverMerchantMessage(detail, "Individual region is invalid. - Individual region is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82669:
		setFailoverMerchantMessage(detail, "Business legal name is required with tax ID. - You must provide a legal name if a tax id has been provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82670:
		setFailoverMerchantMessage(detail, "Applicant account number is invalid. - The provided bank account number is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82671:
		setFailoverMerchantMessage(detail, "Funding account number is invalid. - Funding account number is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82672:
		setFailoverMerchantMessage(detail, "Business tax ID must be blank unless business legal name is present. - The tax id must be blank if no company name/legal name is provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82673:
		setFailoverMerchantMessage(detail, "Applicant tax ID must be blank unless company name present. - Applicant tax ID must be blank unless company name present.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82674:
		setFailoverMerchantMessage(detail, "Merchant accounts with a status of pending or suspended cannot be updated. - The merchant account cannot be updated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode82675:
		setFailoverMerchantMessage(detail, "Merchant account id can not be updated. - You'll get this error if the ID cannot be updated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode82676:
		setFailoverMerchantMessage(detail, "Master merchant account id can not be updated. - You'll get this error if the merchant account ID cannot be updated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode82677:
		setFailoverMerchantMessage(detail, "Business legal name is invalid. - The provided legal name is not valid. The applicant's company name must contain only letters, numbers, and these characters: &-!@#$()'./+,\". The maximum length is 40 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82678:
		setFailoverMerchantMessage(detail, "Funding destination is required. - You must provide a funding destination.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82679:
		setFailoverMerchantMessage(detail, "Funding destination is invalid. - You must provide a valid funding destination. Deprecated funding destinations are treated as invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82680:
		setFailoverMerchantMessage(detail, "Funding email is required when destination is email. - You must provide a funding email address when your funding destination is email.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82681:
		setFailoverMerchantMessage(detail, "Funding email is invalid. - The provided funding email address is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82682:
		setFailoverMerchantMessage(detail, "Funding mobile phone is required when destination is mobile phone. - You must provide a funding mobile phone when your funding destination is mobile phone.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82683:
		setFailoverMerchantMessage(detail, "Funding mobile phone is invalid. - The provided funding mobile phone is not valid. Phone must be 10 - 14 characters and can only contain numbers, parentheses, and periods.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82684:
		setFailoverMerchantMessage(detail, "Business region is invalid. - The provided business region is not valid. Only two-letter abbreviations are accepted, e.g. 'CA' but not 'California.'")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82685:
		setFailoverMerchantMessage(detail, "Business street address is invalid. - The provided business street address is not valid. It must contain at least one digit.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82686:
		setFailoverMerchantMessage(detail, "Business postal code is invalid. - The provided business zip code is not valid. It must be 5 digits followed by an optional hyphen, space, and an additional 4 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82687:
		setFailoverMerchantMessage(detail, "Individual params provided in an invalid format. - You must provide the attributes for Individual in the correct format.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82688:
		setFailoverMerchantMessage(detail, "Business params provided in an invalid format. - You must provide the attributes for Business in the correct format.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82689:
		setFailoverMerchantMessage(detail, "Business locality is invalid. - The provided business address is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82690:
		setFailoverMerchantMessage(detail, "Individual locality is invalid. - The provided individual address is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82691:
		setFailoverMerchantMessage(detail, "Applicant locality is invalid. - The provided applicant address is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode82901:
		setFailoverMerchantMessage(detail, "Incomplete PayPal account information. - You must specify an access token or a consent code for this operation.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode82902:
		setFailoverMerchantMessage(detail, "Pre-Approved Payment enabled PayPal account required for vaulting. - When you vault a PayPal account, you must provide a payment method noncethat was retrieved via the Vault flow.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode82903:
		setFailoverMerchantMessage(detail, "Invalid PayPal account information. - You cannot specify both an access token and a consent code for this operation.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode82904:
		setFailoverMerchantMessage(detail, "PayPal Accounts are not accepted by this merchant account. - Your account has not been enabled to accept PayPal.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode82905:
		setFailoverMerchantMessage(detail, "A customer ID is required to vault a PayPal Account. - When adding a PayPal account to an existing customer, the customer ID is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode83501:
		setFailoverMerchantMessage(detail, "Apple Pay cards are not accepted by this merchant account. - Apple Pay cards are not accepted by this merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode83502:
		setFailoverMerchantMessage(detail, "A customer ID is required to vault an Apple Pay Card. - When storing an Apple Pay card in the vault, you must provide the customer ID of a customer already stored in the vault to whom the card will belong.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode83512:
		setFailoverMerchantMessage(detail, "Apple Pay payment data decryption failed - The Apple Pay PKPaymentTokenpayment data could not be decrypted. This occurs when (a) the Apple Pay merchant id used in your iOS App entitlements does not match the values provided to Braintree; (b) the provisioning profile you used to sign your iOS App does not correspond to the iOS Developer Account with which the Apple Pay certificate was generated; (c) the payment data was not valid as it was received by the Gateway. Contact usif you cannot resolve this error")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode83518:
		setFailoverMerchantMessage(detail, "Credit card type is not accepted by this merchant account. - The specified merchant account is not configured to accept cards from this payment network. Please specify the correct payment networks when initializing a PKPaymentRequest.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode83520:
		setFailoverMerchantMessage(detail, "Payment data is malformed - Payment data is malformed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode84101:
		setFailoverMerchantMessage(detail, "Common ID is required. - Common ID is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode84102:
		setFailoverMerchantMessage(detail, "Username is required. - Username is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode84103:
		setFailoverMerchantMessage(detail, "Venmo user ID is required. - Venmo user ID is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode84104:
		setFailoverMerchantMessage(detail, "Customer ID is required. - Customer ID is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode84105:
		setFailoverMerchantMessage(detail, "Venmo accounts are not accepted by this merchant account. - Venmo accounts are not accepted by this merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode84106:
		setFailoverMerchantMessage(detail, "Customer ID is invalid. - Customer ID is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91501:
		setFailoverMerchantMessage(detail, "Order ID is too long. - Order ID must be less than 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91502:
		setFailoverMerchantMessage(detail, "Order ID is required when the payment_method_nonce is an iDEAL payment. - A transaction for an iDEAL payment must have an order ID. This error still exists in the gateway but is no longer returned. It will be removed from the docs once the error has been removed from gateway.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91504:
		setFailoverMerchantMessage(detail, "Transaction can only be voided if status is authorized, submitted_for_settlement, or - for PayPal - settlement_pending. - Transaction can only be voided if status is authorized, submitted_for_settlement, or - for PayPal - settlement_pending.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91505:
		setFailoverMerchantMessage(detail, "Credit transactions cannot be refunded. - Only sale transactions can be refunded.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91506:
		setFailoverMerchantMessage(detail, "Cannot refund transaction unless it is settled. - Transaction status must be settled to refund it.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91507:
		setFailoverMerchantMessage(detail, "Cannot submit for settlement unless status is authorized. - Transaction status must be authorized to submit the transaction for settlement.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91508:
		setFailoverMerchantMessage(detail, "Cannot determine payment method. - You must specify the payment method to charge, either directly (by payment_method_nonce, payment_method_token, credit_card, paypal_account, etc.) or indirectly (by customer_id, subscription_id, etc.).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode91510:
		setFailoverMerchantMessage(detail, "Customer ID is invalid. - You'll get this error if you create a transaction using a customer ID and the customer ID isn't in your Vault.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91511:
		setFailoverMerchantMessage(detail, "Customer does not have any credit cards. - When creating a transaction using a customer ID, we'll use the customer's default credit card. If the customer does not have any credit cards associated to it, you'll get this error.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91512:
		setFailoverMerchantMessage(detail, "Transaction has already been fully refunded. - You'll get this error if you create a refund on a transaction that has been fully refunded.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_DUPLICATE
	case api.RCode91513:
		setFailoverMerchantMessage(detail, "Merchant account ID is invalid. - If you specify the merchant account ID to use to process a transaction and it does not match any of your merchant accounts, you'll get this error.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91514:
		setFailoverMerchantMessage(detail, "Merchant account is suspended. - You'll get this error if you try to create a transaction using a suspended merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode91515:
		setFailoverMerchantMessage(detail, "Cannot provide both payment_method_token and credit_card attributes. - If you specify both a Vault token and a full credit card number you'll get this error.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91516:
		setFailoverMerchantMessage(detail, "Cannot provide both payment_method_token and customer_id unless the payment_method belongs to the customer. - If you specify both a customer ID and a payment method token when creating a transaction, the payment_method_token must belong to the customer ID.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91517:
		setFailoverMerchantMessage(detail, "Payment instrument type is not accepted by this merchant account. - When providing a payment method token, your merchant account must be configured to accept the payment method type represented by the token.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91518:
		setFailoverMerchantMessage(detail, "Payment method token is invalid. - You'll get this error if the payment method token isn't in the Vault.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode91519:
		setFailoverMerchantMessage(detail, "Processor authorization code cannot be set unless for a voice authorization. - You can only set the processor authorization code for voice authorization transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode91521:
		setFailoverMerchantMessage(detail, "Refund amount is too large. - You cannot refund more than the amount submitted for settlement.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91522:
		setFailoverMerchantMessage(detail, "Settlement amount is too large. - You can't settle more than the authorized amount unless your industry and processor support settlement adjustment (settling a certain percentage over the authorized amount); contact us for details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91523:
		setFailoverMerchantMessage(detail, "Transaction type is invalid. - Valid transaction types are \"sale\" and \"credit\".")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91524:
		setFailoverMerchantMessage(detail, "Transaction type is required. - We need to know if you want to create a sale or a credit.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91525:
		setFailoverMerchantMessage(detail, "Vault is disabled. - If you set the option to store in Vault, your Vault needs to be enabled.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91526:
		setFailoverMerchantMessage(detail, "Custom field is invalid: - Custom field keys must match the API name of a custom field configured in the Control Panel. The error message for this validation error will contain a list of the invalid keys.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91527:
		setFailoverMerchantMessage(detail, "Cannot provide both payment_method_token and subscription_id unless the payment_method belongs to the subscription. - If you specify both a payment method token and a subscription ID the subscription ID must be associated to the token given.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91528:
		setFailoverMerchantMessage(detail, "Subscription ID is invalid. - You'll get this error if the subscription ID given isn't one of your subscriptions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91529:
		setFailoverMerchantMessage(detail, "Cannot provide both subscription_id and customer_id unless the subscription belongs to the customer. - If you give both a customer ID and a subscription ID the subscription must be associated to the customer.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91530:
		setFailoverMerchantMessage(detail, "Cannot provide a billing address unless also providing a credit card. - If you're creating a transaction using a credit card token, then we will use the billing address associated to that token in the Vault. You'll get this error if creating a transaction using a token and specifying a billing address.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91531:
		setFailoverMerchantMessage(detail, "Subscription status must be Past Due in order to retry. - A subscription must be in past due status in order to manually retry the charge.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91537:
		setFailoverMerchantMessage(detail, "Purchase order number is too long. - The purchase order number cannot be larger than 12 characters for AIB and 17 characters for other all other processors.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91538:
		setFailoverMerchantMessage(detail, "Cannot refund a transaction with a suspended merchant account. - You cannot refund a transaction associated with a suspended merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91539:
		setFailoverMerchantMessage(detail, "Voice Authorization is not allowed for this card type - The specified card type does not support voice authorization codes.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91540:
		setFailoverMerchantMessage(detail, "Transaction cannot be cloned if payment method is stored in vault. - Instead, create a new transaction using the payment method's token.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91541:
		setFailoverMerchantMessage(detail, "Cannot clone voice authorization transactions. - Cloning voice authorizations is currently unsupported.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91542:
		setFailoverMerchantMessage(detail, "Unsuccessful transaction cannot be cloned. - Only transactions that were authorized or settled are cloneable.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91543:
		setFailoverMerchantMessage(detail, "Credits cannot be cloned. - You may only clone sale transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91544:
		setFailoverMerchantMessage(detail, "Cannot clone transaction without submit_for_settlement flag. - You must specify whether or not to submit the cloned transaction for settlement upon creation.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91545:
		setFailoverMerchantMessage(detail, "Voice Authorizations are not supported for this processor. - Your processor does not support voice authorizations.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91546:
		setFailoverMerchantMessage(detail, "Credits are not supported by this processor. - Your processor does not support credits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91547:
		setFailoverMerchantMessage(detail, "Merchant account does not support refunds. - The merchant account account does not support refunds.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91548:
		setFailoverMerchantMessage(detail, "Purchase order number is invalid. - The purchase order number must be printable ASCII characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91549:
		setFailoverMerchantMessage(detail, "Cannot provide more than one of payment_method_token, payment_method_nonce, credit_card, and venmo_sdk_payment_method_code attributes. - Cannot provide more than one of payment_method_token, payment_method_nonce, credit_card, and venmo_sdk_payment_method_code attributes.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91550:
		setFailoverMerchantMessage(detail, "Channel is too long. - Channel is too long. This field is for use by partners and shopping cart providers only.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91551:
		setFailoverMerchantMessage(detail, "Settlement amount cannot be less than the service fee amount. - The settlement amount must be greater than or equal to the service fee amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91552:
		setFailoverMerchantMessage(detail, "Credits not allowed with service fee. - Service fees are allowed on sale transactions only.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91553:
		setFailoverMerchantMessage(detail, "Sub-merchant account requires a service fee. - Transactions for sub-merchant accounts need a service fee amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91554:
		setFailoverMerchantMessage(detail, "Amount cannot be negative. - Service fee amount must be greater than or equal to zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91555:
		setFailoverMerchantMessage(detail, "Amount is an invalid format. - Service fee amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the amount cannot include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91556:
		setFailoverMerchantMessage(detail, "Service fee amount is larger than transaction amount. - Service fee amount must be less than the transaction amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91557:
		setFailoverMerchantMessage(detail, "Service fee not supported on master merchant account. - Transactions for a master merchant account cannot have a service fee amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91558:
		setFailoverMerchantMessage(detail, "Merchant account does not support MOTO transactions unless configured by processor. - This merchant account cannot be used for Mail Order/Telephone Order.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91559:
		setFailoverMerchantMessage(detail, "Cannot refund a transaction with a pending merchant account. - A Merchant Account must be Active in order to refund a transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91560:
		setFailoverMerchantMessage(detail, "Transaction could not be held in escrow. - The Transaction cannot be held in escrow.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode91561:
		setFailoverMerchantMessage(detail, "Cannot release a transaction that is not escrowed. - Cannot release a transaction that is not escrowed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91562:
		setFailoverMerchantMessage(detail, "Release can only be cancelled if the transaction is submitted for release. - Release can only be canceled if the transaction is submitted for release.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91563:
		setFailoverMerchantMessage(detail, "Escrowed transactions cannot be partially refunded. - The Transaction must be fully refunded after being held in escrow.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91564:
		setFailoverMerchantMessage(detail, "Cannot use a payment_method_nonce more than once. - The payment method nonce has already been consumed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91565:
		setFailoverMerchantMessage(detail, "Unknown or expired payment_method_nonce. - The payment method nonce has either expired or never existed. Nonces are deleted upon expiration (3 hours after generation).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91567:
		setFailoverMerchantMessage(detail, "Payment instrument type is not accepted by this merchant account. - Payment instrument type is not accepted by this merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91568:
		setFailoverMerchantMessage(detail, "Three D Secure Token is invalid. - 3D Secure token is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91569:
		setFailoverMerchantMessage(detail, "payment_method_nonce does not contain a valid payment instrument type. - payment_method_nonce does not contain a valid payment instrument type.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91570:
		setFailoverMerchantMessage(detail, "Transaction data does not match data from Three D Secure verify call. - The credit card number and expiration date used for 3D Secure verification must match the values used to create the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode91572:
		setFailoverMerchantMessage(detail, "Current payment method does not support use_billing_for_shipping flag. - Current payment method does not support use_billing_for_shipping flag.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91573:
		setFailoverMerchantMessage(detail, "Transaction cannot be cloned if payment method is a PayPal account. - Transaction cannot be cloned if payment method is a PayPal account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91574:
		setFailoverMerchantMessage(detail, "Cannot refund a transaction transaction in settling status on this merchant account. Try again after the transaction has settled. - Cannot refund a transaction transaction in settling status on this merchant account. Try again after the transaction has settled.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91575:
		setFailoverMerchantMessage(detail, "Cannot transition transaction to settled, settlement_confirmed, or settlement_declined - Cannot transition transaction to settled or settlement_declined.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91576:
		setFailoverMerchantMessage(detail, "PayPal is not enabled for your merchant account. - PayPal is not enabled for your merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91577:
		setFailoverMerchantMessage(detail, "Merchant account does not support payment instrument. - Merchant account does not support payment instrument.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode91578:
		setFailoverMerchantMessage(detail, "Service fee can not be applied on PayPal transactions. - Service fee can not be applied on PayPal transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91580:
		setFailoverMerchantMessage(detail, "PayPal custom field must be less than 256 characters in length. - PayPal custom field must be less than 256 characters in length.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91581:
		setFailoverMerchantMessage(detail, "Shipping address customer does not match customer in request. - Shipping address customer does not match customer in request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91582:
		setFailoverMerchantMessage(detail, "PayPal unilateral transactions must also be submitted for settlement. - A unilateral payment is a payment that is made to a receiver who does not have a PayPal account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91583:
		setFailoverMerchantMessage(detail, "This PayPal account was not vaulted with the required data - This PayPal account was not vaulted with the required data.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91584:
		setFailoverMerchantMessage(detail, "Merchant account must match the 3D Secure authorization merchant account. - The merchant account specified when creating the transaction must match the merchant account specified when generating the client token. If you do not specify a merchant account ID, your default merchant account will be used.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode91585:
		setFailoverMerchantMessage(detail, "Amount must match the 3D Secure authorization amount. - The amount used for 3D Secure verification must match the amount used to create the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode91586:
		setFailoverMerchantMessage(detail, "Shared billing address ID cannot be used in the same call as a standard billing address ID - Shared billing address ID cannot be used in the same call as a standard billing address ID.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91587:
		setFailoverMerchantMessage(detail, "Shared customer ID cannot be used in the same call as a standard customer ID - Shared customer ID cannot be used in the same call as a standard customer ID.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91588:
		setFailoverMerchantMessage(detail, "Shared payment method token cannot be used in the same call as a standard payment method token - Shared payment method token cannot be used in the same call as a standard payment method token.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode91589:
		setFailoverMerchantMessage(detail, "Shared payment method token cannot be used in the same call as a non-shared identifier param - Shared payment method token cannot be used in the same call as a non-shared identifier param.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode91590:
		setFailoverMerchantMessage(detail, "Shared identifier param cannot be used with non-shared payment method token - Shared identifier param cannot be used with non-shared payment method token.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91591:
		setFailoverMerchantMessage(detail, "Shared shipping address ID cannot be used in the same call as a standard shipping address ID - Shared shipping address ID cannot be used in the same call as a standard shipping address ID.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91592:
		setFailoverMerchantMessage(detail, "Shared payment methods cannot be vaulted - Shared payment methods cannot be vaulted.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91593:
		setFailoverMerchantMessage(detail, "Shared payment methods cannot be vaulted - Shared payment methods cannot be vaulted.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91594:
		setFailoverMerchantMessage(detail, "Shared shipping addresses cannot be vaulted - Shared shipping addresses cannot be vaulted.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91595:
		setFailoverMerchantMessage(detail, "Shared payment methods cannot be updated - Shared payment methods cannot be updated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91596:
		setFailoverMerchantMessage(detail, "Shared payment method token is invalid. - Shared payment method token is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode91597:
		setFailoverMerchantMessage(detail, "Cannot provide both shared_payment_method_token and shared_customer_id unless the payment_method belongs to the customer. - Cannot provide both shared_payment_method_token and shared_customer_id unless the payment_method belongs to the customer.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91598:
		setFailoverMerchantMessage(detail, "Payment instrument type is not accepted by this merchant account. - Payment instrument type is not accepted by this merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91599:
		setFailoverMerchantMessage(detail, "Shared Shipping address customer does not match customer in request. - Shared Shipping address customer does not match customer in request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91602:
		setFailoverMerchantMessage(detail, "Custom field is invalid: - Custom field keys must match the API name of a custom field configured in the Control Panel. The error message for this validation error will contain a list of the invalid keys.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91609:
		setFailoverMerchantMessage(detail, "Customer ID has already been taken. - Customer IDs have to be unique.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91610:
		setFailoverMerchantMessage(detail, "Customer ID is invalid (use only letters, numbers, '-', and '_'). - Valid characters are letters, numbers, dashes, and underscores.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91611:
		setFailoverMerchantMessage(detail, "Customer ID is not an allowed ID. - We reserve a few words that can't be used as IDs. 'all' and 'new' currently cannot be used.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91612:
		setFailoverMerchantMessage(detail, "Customer ID is too long. - Maximum 36 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91613:
		setFailoverMerchantMessage(detail, "Customer ID is required. - Customer ID is required when performing updates.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91617:
		setFailoverMerchantMessage(detail, "Nonce references a vaulted payment instrument - cannot be transferred between customers - Nonce references a vaulted payment instrument - cannot be transferred between customers.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode91618:
		setFailoverMerchantMessage(detail, "Customer attribute must be a map of keys and values representing a customer. - Customer must be a well-formed object, not a string or integer.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91619:
		setFailoverMerchantMessage(detail, "Ambiguous usage of default payment method token. - Cannot set default_payment_method_token and credit_card.options.make_defaultin the same customer update request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91620:
		setFailoverMerchantMessage(detail, "PayPal custom field must be less than 256 characters in length. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91621:
		setFailoverMerchantMessage(detail, "PayPal description must be less than 256 characters in length. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91622:
		setFailoverMerchantMessage(detail, "Order ID must be less than 256 characters in length. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91623:
		setFailoverMerchantMessage(detail, "Amount format is invalid. - Amount must be formatted like '10' or '10.00'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91701:
		setFailoverMerchantMessage(detail, "Cannot provide both a billing address and a billing address ID. - When you create or update a credit card you can set the billing address using full billing address details, or you can set it to a billing address ID of an address already associated to the customer, but not both.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91702:
		setFailoverMerchantMessage(detail, "Billing address ID is invalid. - If setting the billing address on a credit card using an ID, the ID must be an ID of an address associated to the customer.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91704:
		setFailoverMerchantMessage(detail, "Customer ID is required. - When adding a credit card to an existing customer, the customer ID is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91705:
		setFailoverMerchantMessage(detail, "Customer ID is invalid. - When specifying the customer ID to add a credit card to an existing customer, the ID must be the ID a customer stored in the Vault.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91708:
		setFailoverMerchantMessage(detail, "Cannot provide expirationdate if you are also providing expiration_month and expiration_year. - You can provide the credit card expiration date as a single field, or as month and year separately, but not all 3.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91718:
		setFailoverMerchantMessage(detail, "Token is invalid (use only letters, numbers, '-', and ''). - If you're specifying the credit card token, you can use letters, numbers, '-', and '_'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91719:
		setFailoverMerchantMessage(detail, "Credit card token is taken. - Credit card tokens have to be unique.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91720:
		setFailoverMerchantMessage(detail, "Credit card token is too long. - Maximum 36 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91721:
		setFailoverMerchantMessage(detail, "Token is not an allowed token. - We reserve a few tokens: 'new' and 'all'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91722:
		setFailoverMerchantMessage(detail, "Payment Method token is required. - When updating a credit card you can omit the token if you don't want to change it, but you can't set it to an empty string. If set to an empty string on creation, the gateway will generate a random token.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91723:
		setFailoverMerchantMessage(detail, "Update Existing Token is invalid. - Applies when updating a customer and credit card at the same time and specifying the token of the credit card to update. You'll get this error if the token specified is for a credit card that does not exist, or references a credit card that does not belong to the customer that is being updated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91726:
		setFailoverMerchantMessage(detail, "Credit card type is not accepted by this merchant account. - Credit card type is not accepted by this merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91727:
		setFailoverMerchantMessage(detail, "Invalid VenmoSDK payment method code - Invalid VenmoSDK payment method code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91728:
		setFailoverMerchantMessage(detail, "Verification Merchant Account ID is invalid. - There must be a merchant account with this ID.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91729:
		setFailoverMerchantMessage(detail, "Update Existing Token is not allowed when creating a customer. - Update Existing Token is not allowed when creating a customer.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91730:
		setFailoverMerchantMessage(detail, "Verifications are not supported on this merchant account - This error occurs when a merchant account does not support credit card verification. This can also occur if your API user does not have access to the merchant account used for the verification.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91731:
		setFailoverMerchantMessage(detail, "Cannot use a payment_method_nonce more than once. - The payment method nonce has already been used once.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91732:
		setFailoverMerchantMessage(detail, "Unknown or expired payment_method_nonce. - The payment method nonce has either expired or never existed. Nonces are deleted upon expiration (3 hours after generation).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91733:
		setFailoverMerchantMessage(detail, "Payment method nonce locked. - Deprecated The payment method nonce must be unlocked before it is used.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91734:
		setFailoverMerchantMessage(detail, "Credit card type is not accepted by this merchant account. - Applies when specifying a credit card when creating a transaction, but not when only storing in the Vault since Vault records are not associated to specific merchant accounts.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode91735:
		setFailoverMerchantMessage(detail, "Payment method nonces cannot be used to update an existing card. - A payment method nonce cannot be used to update an existing credit card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91738:
		setFailoverMerchantMessage(detail, "Payment method is not a credit card payment method. - This operation requires a credit card, and the payment method you specified is not a credit card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91739:
		setFailoverMerchantMessage(detail, "Verification amount cannot be negative. - The amount you specified for verification was less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91740:
		setFailoverMerchantMessage(detail, "Verification amount is invalid. - The amount you specified for verification had an invalid format.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91741:
		setFailoverMerchantMessage(detail, "Verification amount not supported by processor. - The processor you are using for verification does not allow the verification amount you specified.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91742:
		setFailoverMerchantMessage(detail, "Verification Merchant Account is suspended. - Verification Merchant Account is suspended.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91743:
		setFailoverMerchantMessage(detail, "The current user does not have access to the specified verification_merchant_account_id - The current user does not have access to the specified verification_merchant_account_id.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91744:
		setFailoverMerchantMessage(detail, "Billing address format is invalid. - Billing address format is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91745:
		setFailoverMerchantMessage(detail, "Payment method params supplied are not valid for updating a credit card. - The payment method params you supplied are not valid for updating a credit card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91752:
		setFailoverMerchantMessage(detail, "Verification amount is too large. - Verification amount is too large.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91755:
		setFailoverMerchantMessage(detail, "Verification Merchant Account ID cannot be a sub-merchant account. - Verifications cannot be created using sub-merchant accounts.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91756:
		setFailoverMerchantMessage(detail, "Customer ID cannot be updated. - Customer ID cannot be updated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91803:
		setFailoverMerchantMessage(detail, "Country name is not an accepted country. - We only accept specific country names.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91814:
		setFailoverMerchantMessage(detail, "Country code (alpha2) is not an accepted country. - We only accept specific alpha-2 values.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91815:
		setFailoverMerchantMessage(detail, "Provided country information is inconsistent. - You can only specify one of country name, country code alpha2, country code alpha3 and country code numeric.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91816:
		setFailoverMerchantMessage(detail, "Country code (alpha3) is not an accepted country. - We only accept specific alpha-3 values.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91817:
		setFailoverMerchantMessage(detail, "Country code (numeric) is not an accepted country. - We only accept specific numeric values.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91818:
		setFailoverMerchantMessage(detail, "Customer has already reached the maximum of 50 addresses. - You will get this validation error when trying to add an address to a customer which has already reached the maximum of 50 addresses.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91819:
		setFailoverMerchantMessage(detail, "First name must be a string. - First name must be a string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode91820:
		setFailoverMerchantMessage(detail, "Last name must be a string. - Last name must be a string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode91821:
		setFailoverMerchantMessage(detail, "Company must be a string. - Company must be a string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode91822:
		setFailoverMerchantMessage(detail, "Street address must be a string. - Street address must be a string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode91823:
		setFailoverMerchantMessage(detail, "Extended address must be a string. - Extended address must be a string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode91824:
		setFailoverMerchantMessage(detail, "Locality must be a string. - Locality must be a string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode91825:
		setFailoverMerchantMessage(detail, "Region must be a string. - Region must be a string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode91826:
		setFailoverMerchantMessage(detail, "Postal code must be a string. - Postal code must be a string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode91828:
		setFailoverMerchantMessage(detail, "Address is invalid. - Address must provided in a valid format.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode91901:
		setFailoverMerchantMessage(detail, "Merchant Account ID is invalid. - If specifying the merchant account to use to process transactions for this subscription it needs to be the ID for one of your merchant accounts.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91902:
		setFailoverMerchantMessage(detail, "Payment method token payment instrument type is not accepted by this merchant account. - When providing a payment method token, your merchant account must be configured to accept the payment method type represented by the token.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode91903:
		setFailoverMerchantMessage(detail, "Payment method token is invalid. - You'll get this error if we can't find a payment method with the token specified.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91904:
		setFailoverMerchantMessage(detail, "Plan ID is invalid. - You'll get this error if we can't find a plan with the given ID.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91905:
		setFailoverMerchantMessage(detail, "Payment method token does not belong to the subscription's customer. - When updating a subscription and changing the payment method token, you can only use tokens associated to the same customer that the subscription is currently associated to.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91906:
		setFailoverMerchantMessage(detail, "Number Of Billing Cycles must be numeric. - It must be a number.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91907:
		setFailoverMerchantMessage(detail, "Number Of Billing Cycles must be greater than zero. - It must be greater than 0.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91908:
		setFailoverMerchantMessage(detail, "Cannot specify both number of billing cycles and never expires as true. - You cannot specify both number of billing cycles and never expires as true.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91909:
		setFailoverMerchantMessage(detail, "Number Of Billing Cycles is less than the current billing cycle. - You cannot edit a subscription and change the number of billing cycles to be below the current count of billing cycles.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91911:
		setFailoverMerchantMessage(detail, "Cannot add duplicate add-on or discount. - Add-Ons and Discounts must be unique, but you can change the quantity.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91912:
		setFailoverMerchantMessage(detail, "Number Of Billing Cycles cannot be blank if the subscription expires. - Blanks are not allowed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91913:
		setFailoverMerchantMessage(detail, "Billing Day of Month must be numeric. - The billing date must be a number.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91914:
		setFailoverMerchantMessage(detail, "Billing Day of Month must be between 1 and 28, or 31. - The billing date must be 1-28 or 31 (for the last day of every month).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91915:
		setFailoverMerchantMessage(detail, "First Billing Date is invalid. - The first billing date is an incorrect format.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91916:
		setFailoverMerchantMessage(detail, "First Billing Date cannot be in the past. - The first billing date cannot be in the past.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91917:
		setFailoverMerchantMessage(detail, "Cannot specify more than one type of start date. - The type of start date can be only one of the following: start immediately, after a trial period, or on a specific day of the month.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91918:
		setFailoverMerchantMessage(detail, "Billing Day of Month cannot be updated. - The billing date cannot be updated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91919:
		setFailoverMerchantMessage(detail, "First Billing Date cannot be updated. - First billing date cannot be updated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91920:
		setFailoverMerchantMessage(detail, "Can only edit id, merchant account id, payment method token, and descriptor on a past due subscription. - You cannot edit any fields which could change the price on a past due subscription.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91921:
		setFailoverMerchantMessage(detail, "Invalid request format. - The add-ons and discounts are in an invalid format.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91922:
		setFailoverMerchantMessage(detail, "Cannot update subscription to a plan with a different billing frequency. - You will get this validation if you try to update the plan on a subscription and the billing cycle for the new plan is not the same as the billing cycle of the old plan.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91923:
		setFailoverMerchantMessage(detail, "Subscription Plan currency must be the same as the merchant account's currency. - Subscription Plan currency must be the same as the merchant account's currency.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91924:
		setFailoverMerchantMessage(detail, "Payment method nonce payment instrument type is not accepted by this merchant account. - The supplied payment method nonce represents a payment method of a type that is not accepted by this merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91925:
		setFailoverMerchantMessage(detail, "Payment method nonce is invalid. - The supplied payment method nonce was not in a valid format or is unknown.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91926:
		setFailoverMerchantMessage(detail, "Payment method nonce does not belong to the subscription's customer. - The payment method nonce used to create a subscription must be vaulted and must belong to the customer owning the subscription.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91927:
		setFailoverMerchantMessage(detail, "Payment method nonce represents an un-vaulted payment instrument. - You cannot create a subscription with a nonce representing an unvaulted payment instrument. Use the payment method nonce to create a vaulted payment method first.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91928:
		setFailoverMerchantMessage(detail, "Payment instrument type is not valid for subscriptions. - Payment instrument type is not valid for subscriptions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91929:
		setFailoverMerchantMessage(detail, "Payment instrument type is not valid for subscriptions. - Payment instrument type is not valid for subscriptions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91930:
		setFailoverMerchantMessage(detail, "Merchant Account does not support the given payment instrument type. - Merchant Account does not support the given payment instrument type.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode91931:
		setFailoverMerchantMessage(detail, "Too many concurrent attempts to update this subscription. Try again later. - Too many concurrent attempts to update this subscription. Try again later.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_PROCESSING
	case api.RCode92001:
		setFailoverMerchantMessage(detail, "Quantity is invalid. - Quantity must be a number.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92002:
		setFailoverMerchantMessage(detail, "Amount is invalid. - Amount must be formatted like '10' or '10.00'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92003:
		setFailoverMerchantMessage(detail, "Amount cannot be blank. - Blanks are not allowed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92004:
		setFailoverMerchantMessage(detail, "Quantity cannot be blank. - Blanks are not allowed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92005:
		setFailoverMerchantMessage(detail, "Number of billing cycles is invalid. - Number of billing cycles must be numeric.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92010:
		setFailoverMerchantMessage(detail, "Quantity must be greater than zero. - Quantity must be greater than 0.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92011:
		setFailoverMerchantMessage(detail, "Existing ID is invalid. - Modification ID must be associated with the subscription.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92012:
		setFailoverMerchantMessage(detail, "Existing ID is required. - Modification ID is required to update a modification.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92013:
		setFailoverMerchantMessage(detail, "Inherited From ID is invalid. - Modification with that inherited from ID is not available.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92014:
		setFailoverMerchantMessage(detail, "Inherited From ID is required. - Must provide an inherited from ID.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92015:
		setFailoverMerchantMessage(detail, "Cannot update a removed add-on or discount. - Cannot update and remove a modification at the same time.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92016:
		setFailoverMerchantMessage(detail, "Cannot remove add-on or discount if not already associated with subscription. - Cannot remove add-on or discount if not already associated with subscription.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92017:
		setFailoverMerchantMessage(detail, "Number of billing cycles cannot be blank. - Number Of Billing Cycles cannot be blank if the subscription expires.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92018:
		setFailoverMerchantMessage(detail, "Cannot specify both number of billing cycles and never expires as true. - Number of billing cycles is blank and never expires is not set to true.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92019:
		setFailoverMerchantMessage(detail, "Number of billing cycles must be greater than zero. - Number of billing cycles must be greater than 0.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92020:
		setFailoverMerchantMessage(detail, "Existing ID is not of the correct kind. - Existing ID must be of the type of modification that is being edited.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92021:
		setFailoverMerchantMessage(detail, "ID to remove is incorrect kind. - Existing ID must be of the type of modification that is being removed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92022:
		setFailoverMerchantMessage(detail, "Cannot edit add-on or discount on a past due subscription. - Unable to edit modifications on subscriptions that are past due.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92023:
		setFailoverMerchantMessage(detail, "Amount is too large. - The amount cannot be greater than the maximum allowed by the processor. Contact us for your specific limit.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92024:
		setFailoverMerchantMessage(detail, "Cannot pass null modification. - Modification is missing from the API call.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92025:
		setFailoverMerchantMessage(detail, "ID to remove is invalid. - ID to remove is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92201:
		setFailoverMerchantMessage(detail, "Company name/DBA section is invalid. - The company/DBA name passed in your descriptor does not meet the requirements set by your processor.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92202:
		setFailoverMerchantMessage(detail, "Phone number is invalid. - Typically, phone must be 10 - 14 characters and can only contain numbers, dashes, parentheses and periods. Some example phone numbers are: 3125556666 312-555-6666 (312)555-6666 However, your processor may impact the number of characters or format accepted. Contact us for more information on your processor's specific requirements.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode92203:
		setFailoverMerchantMessage(detail, "Dynamic descriptors have not been enabled for this account. Please contact support at https://help.braintreepayments.com. - Dynamic descriptors have not been enabled for this account. Contact us.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode92204:
		setFailoverMerchantMessage(detail, "Descriptor format is invalid. - Descriptor name must be less than or equal to 15 characters and can only contain letters and numbers. This will be prefixed by the preset company name for this account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92205:
		setFailoverMerchantMessage(detail, "International phone number is invalid. - Deprecated Phone can only contain numbers, dashes and periods and must be less than or equal to 13 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode92206:
		setFailoverMerchantMessage(detail, "URL must be 13 characters or shorter. - The url/web address to a customer's statement must be less than or equal to 13 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92801:
		setFailoverMerchantMessage(detail, "Cannot specify make_default without a customer_id - Cannot specify make_default without a customer_id")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92802:
		setFailoverMerchantMessage(detail, "Cannot specify verify_card without a customer_id - Cannot specify verify_card without a customer_id")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92803:
		setFailoverMerchantMessage(detail, "Cannot specify fail_on_duplicate_payment_method without a customer_id - Cannot specify fail_on_duplicate_payment_method without a customer_id")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92804:
		setFailoverMerchantMessage(detail, "Customer specified by customer_id does not exist - Customer specified by customer_id does not exist")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92806:
		setFailoverMerchantMessage(detail, "Unsupported client token version - Unsupported client token version")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92807:
		setFailoverMerchantMessage(detail, "Merchant Account specified by merchant_account_id does not exist - Merchant Account specified by merchant_account_id does not exist")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92906:
		setFailoverMerchantMessage(detail, "PayPal Account token is taken. - PayPal account tokens have to be unique.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92907:
		setFailoverMerchantMessage(detail, "Cannot use a payment_method_nonce more than once. - A payment method nonce may only be consumed once.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92908:
		setFailoverMerchantMessage(detail, "Unknown or expired payment_method_nonce. - The payment method nonce has either expired or never existed. Nonces are deleted upon expiration (3 hours after generation).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92909:
		setFailoverMerchantMessage(detail, "Payment method nonce locked. - Deprecated The payment method nonce must be unlocked before it is used.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92910:
		setFailoverMerchantMessage(detail, "Error communicating with PayPal. - There was an error communicating with PayPal.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92911:
		setFailoverMerchantMessage(detail, "PayPal authentication expired. - Deprecated The authentication you received from your user has expired.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92912:
		setFailoverMerchantMessage(detail, "Funding source selection was given without an access token. - You cannot specify a funding source without also specifying an access token.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92913:
		setFailoverMerchantMessage(detail, "Funding source object is invalid or missing required fields. - You sent an invalid or incomplete funding source specification.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92914:
		setFailoverMerchantMessage(detail, "Payment method nonces cannot be used to update an existing PayPal account. - A payment method nonce cannot be used to update an existing PayPal account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92915:
		setFailoverMerchantMessage(detail, "Payment method params supplied are not valid for updating a PayPal account. - The payment method params you supplied are not valid for updating a PayPal account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode92916:
		setFailoverMerchantMessage(detail, "Error executing PayPal order. - There was an error while executing the PayPal order.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode92917:
		setFailoverMerchantMessage(detail, "Error executing PayPal billing agreement. - There was an error while executing the PayPal billing agreement.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode92918:
		setFailoverMerchantMessage(detail, "Merchant setup must be completed before vaulting PayPal payment methods. - The merchant account used is not fully set up; PayPal payment methods cannot be vaulted until setup is completed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode92919:
		setFailoverMerchantMessage(detail, "This merchant account does not allow PayPal payments using the old Vault flow. - Your account does not have the old Vault flow enabled for PayPal payments. Contact us to enable the old Vault flow.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode92920:
		setFailoverMerchantMessage(detail, "Using a payment method nonce for PayPal intent=order is not permitted. - Create a payment method from the nonce and then use the payment method token in the transaction sale request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93101:
		setFailoverMerchantMessage(detail, "Payment method params are required. - A top level payment method parameter is missing.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93102:
		setFailoverMerchantMessage(detail, "Nonce is invalid. - The nonce that was received is not a valid nonce.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93103:
		setFailoverMerchantMessage(detail, "Nonce is required. - A nonce was not provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93104:
		setFailoverMerchantMessage(detail, "Customer ID is required. - A customer id was not provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93105:
		setFailoverMerchantMessage(detail, "Customer ID is invalid. - The customer id does not reference a customer.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93106:
		setFailoverMerchantMessage(detail, "Cannot forward a payment method of this type. - Only credit cards may be forwarded.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93107:
		setFailoverMerchantMessage(detail, "Cannot use a payment_method_nonce more than once. - A payment method nonce may only be consumed once.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93108:
		setFailoverMerchantMessage(detail, "Unknown or expired payment_method_nonce. - The payment method nonce has either expired or never existed. Nonces are deleted upon expiration (3 hours after generation).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93109:
		setFailoverMerchantMessage(detail, "Nonce is not vaultable. - Nonce is not vaultable.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93113:
		setFailoverMerchantMessage(detail, "PayPal custom field must be less than 256 characters in length. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93114:
		setFailoverMerchantMessage(detail, "PayPal description must be less than 256 characters in length. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93115:
		setFailoverMerchantMessage(detail, "Order ID must be less than 256 characters in length. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93116:
		setFailoverMerchantMessage(detail, "Amount format is invalid. - Amount must be formatted like '10' or '10.00'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93117:
		setFailoverMerchantMessage(detail, "This payment method is no longer supported. - This payment method is no longer supported.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93118:
		setFailoverMerchantMessage(detail, "PayPal refresh token is invalid. - PayPal has reported this refresh token is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93119:
		setFailoverMerchantMessage(detail, "Nonce is not a valid parameter when PayPal refresh token is provided. - Nonce is not a valid parameter when PayPal refresh token is provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93120:
		setFailoverMerchantMessage(detail, "PayPal merchant account required to vault refresh token. - A PayPal merchant account required to vault a refresh token. Enable PayPal in a new or existing merchant account and retry.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93121:
		setFailoverMerchantMessage(detail, "US bank account verification method is invalid. - Valid US bank account verification methods include independent_check, network_check, or micro_transfers.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93122:
		setFailoverMerchantMessage(detail, "US bank account is not accepted by merchant account. - A merchant account must support US bank accounts before a US bank account payment method can be created.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode93401:
		setFailoverMerchantMessage(detail, "Industry type is invalid. - Industry type is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93402:
		setFailoverMerchantMessage(detail, "Lodging data is empty. - Lodging data is empty.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93403:
		setFailoverMerchantMessage(detail, "Folio number is invalid. - Folio number is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93404:
		setFailoverMerchantMessage(detail, "Check in date is invalid. - Check in date is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93405:
		setFailoverMerchantMessage(detail, "Check out date is invalid. - Check out date is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93406:
		setFailoverMerchantMessage(detail, "Check out date must occur after the check in date. - Check out date must occur after the check in date.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93407:
		setFailoverMerchantMessage(detail, "Data fields are unknown. - Data fields are unknown.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93408:
		setFailoverMerchantMessage(detail, "Travel and Cruise data is empty. - Travel and Cruise data is empty.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93409:
		setFailoverMerchantMessage(detail, "Data fields are unknown. - Data fields are unknown.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93410:
		setFailoverMerchantMessage(detail, "Travel Package is invalid. - Travel Package is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode93411:
		setFailoverMerchantMessage(detail, "Departure date is invalid. - Departure date is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode93412:
		setFailoverMerchantMessage(detail, "Lodging check in date is invalid. - Lodging check in date is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode93413:
		setFailoverMerchantMessage(detail, "Lodging check out date is invalid. - Lodging check out date is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode93414:
		setFailoverMerchantMessage(detail, "Travel and Flight data is empty. - Travel and Flight data is empty.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode93415:
		setFailoverMerchantMessage(detail, "Data fields are unknown. - Data fields are unknown.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode93416:
		setFailoverMerchantMessage(detail, "Customer code is too long. - Customer code can't be longer than 17 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93417:
		setFailoverMerchantMessage(detail, "Fare amount cannot be negative. - Fare amount can't be less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93418:
		setFailoverMerchantMessage(detail, "Fare amount is an invalid format. - Fare amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the fare amount can't include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93419:
		setFailoverMerchantMessage(detail, "Fare amount is too large. - Fare amount can't be longer than 9 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93420:
		setFailoverMerchantMessage(detail, "Fee amount cannot be negative. - Fare amount can't be less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93421:
		setFailoverMerchantMessage(detail, "Fee amount is an invalid format. - Fee amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the fee amount can't include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93422:
		setFailoverMerchantMessage(detail, "Fee amount is too large. - Fee amount can't be longer than 9 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93423:
		setFailoverMerchantMessage(detail, "Issued date is an invalid format. - Issued date is an invalid format.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93424:
		setFailoverMerchantMessage(detail, "Issuing carrier code is too long. - Issuing carrier code can't be longer than 4 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93425:
		setFailoverMerchantMessage(detail, "Passenger middle initial is too long. - Passenger middle initial can't be longer than 1 character.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93426:
		setFailoverMerchantMessage(detail, "Restricted ticket is required. - You must specify if the ticket is restricted.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode93427:
		setFailoverMerchantMessage(detail, "Tax amount cannot be negative. - Tax amount can't be less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93428:
		setFailoverMerchantMessage(detail, "Tax amount is an invalid format. - Tax amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the tax amount can't include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93429:
		setFailoverMerchantMessage(detail, "Tax amount is too large. - Tax amount can't be longer than 9 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93430:
		setFailoverMerchantMessage(detail, "Ticket number is too long. - Ticket number can't be longer than 15 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93431:
		setFailoverMerchantMessage(detail, "Expected a collection of legs but none provided. - Expected a collection of legs but unrecognized data provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode93432:
		setFailoverMerchantMessage(detail, "Too many legs. - A maximum of 4 legs can be provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode93503:
		setFailoverMerchantMessage(detail, "Apple Pay token is taken. - Payment method tokens must be unique across all payment method types.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93504:
		setFailoverMerchantMessage(detail, "Cannot use a payment_method_nonce more than once. - The payment method nonce has already been used.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93505:
		setFailoverMerchantMessage(detail, "Unknown or expired payment_method_nonce. - The payment method nonce has either expired or never existed. Nonces are deleted upon expiration (3 hours after generation).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93506:
		setFailoverMerchantMessage(detail, "Payment method nonce locked. - Deprecated The payment method nonce must be unlocked before it is used.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93507:
		setFailoverMerchantMessage(detail, "Payment method nonces cannot be used to update an existing Apple Pay Card. - A vaulted payment method cannot be updated with an Apple Pay nonce. Create a new payment method instead.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93508:
		setFailoverMerchantMessage(detail, "Number is required for Apple Pay Card - The Apple Pay PKPaymentTokenpayment data was malformed (did not contain a card number).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93509:
		setFailoverMerchantMessage(detail, "Expiration Month is required for Apple Pay Card - The Apple Pay PKPaymentTokenpayment data was malformed (did not contain an expiration month).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93510:
		setFailoverMerchantMessage(detail, "Expiration Year is required for Apple Pay Card - The Apple Pay PKPaymentTokenpayment data was malformed (did not contain an expiration year).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93511:
		setFailoverMerchantMessage(detail, "Cryptogram is required for Apple Pay Card - The Apple Pay PKPaymentTokenpayment data was malformed (did not contain a cryptogram). This is sometimes caused by misconfigured merchantCapabilities in the PKPaymentRequest. See our recommendations in the Apple Pay guide.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93513:
		setFailoverMerchantMessage(detail, "Apple Pay is disabled for this merchant - Your merchant account is not configured for Apple Pay support. Contact us to configure and enable Apple Pay.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93514:
		setFailoverMerchantMessage(detail, "Apple Pay certificate, private key or merchant ID not configured - Your merchant account is not configured for Apple Pay support. Contact us to configure and enable Apple Pay.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93517:
		setFailoverMerchantMessage(detail, "Certificate provided is not valid - Certificate provided is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93519:
		setFailoverMerchantMessage(detail, "Public key used to sign payment data does not match stored certificate - Public key used to sign payment data does not match stored certificate.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93521:
		setFailoverMerchantMessage(detail, "Private key stored does not match private key used to encrypt payment data - Private key stored does not match private key used to encrypt payment data.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode93522:
		setFailoverMerchantMessage(detail, "Certificate does not match stored key pair - The Apple Pay certificate you uploaded does not match the key pair we have stored for your account. Please download a new CSR from the Control Panel and create a new certificate for your Apple Pay merchant ID using this CSR.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode93523:
		setFailoverMerchantMessage(detail, "Domain name is required. - You must specify a domain name.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93524:
		setFailoverMerchantMessage(detail, "An error occurred when validating your domain with Apple. - An error occurred when validating your domain with Apple. Please try again.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode93525:
		setFailoverMerchantMessage(detail, "Domain verification with Apple failed. Please verify the file is available at the verification path and try again. - Domain verification with Apple failed. Please verify the file is available at the verification path and try again.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode93526:
		setFailoverMerchantMessage(detail, "There was an issue contacting Apple, please try again later. - There was an issue contacting Apple, please try again later.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode93801:
		setFailoverMerchantMessage(detail, "Invalid grant: - The provided authorization code or refresh token is invalid due to one of the reasons described in the error message: code not found code has been used code has expired refresh token not found refresh token is expired refresh token is revokedA revoked refresh token might also mean that it was already used.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93802:
		setFailoverMerchantMessage(detail, "Invalid credentials: - Client authentication failed (e.g. unknown client, no client authentication included, or unsupported authentication method).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93803:
		setFailoverMerchantMessage(detail, "Invalid scope: - The requested scope is invalid, unknown, malformed, or exceeds the scope granted by the resource owner.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93804:
		setFailoverMerchantMessage(detail, "Invalid request: - The request is missing a required parameter, includes an unsupported parameter value (other than grant type), repeats a parameter, includes multiple credentials, utilizes more than one mechanism for authenticating the client, or is otherwise malformed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode93805:
		setFailoverMerchantMessage(detail, "Unsupported grant type: - The authorization grant type is not supported by the authorization server.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94201:
		setFailoverMerchantMessage(detail, "Verification amount cannot be negative. - Verification amount cannot be negative.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94202:
		setFailoverMerchantMessage(detail, "Verification amount is invalid. - Verification amount is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94203:
		setFailoverMerchantMessage(detail, "Verification amount not supported by processor. - Verification amount not supported by processor.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94204:
		setFailoverMerchantMessage(detail, "Verification Merchant Account ID is invalid. - Verification merchant account ID is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94205:
		setFailoverMerchantMessage(detail, "Verification Merchant Account is suspended. - Verification Merchant Account is suspended.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94206:
		setFailoverMerchantMessage(detail, "The current user does not have access to the specified merchant_account_id - The current user does not have access to the specified merchant_account_id.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94207:
		setFailoverMerchantMessage(detail, "Verification amount is too large. - Verification amount is too large.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94208:
		setFailoverMerchantMessage(detail, "Verification Merchant Account ID cannot be a sub-merchant account. - Verifications cannot be created using sub-merchant accounts.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94501:
		setFailoverMerchantMessage(detail, "Payment method nonce not found - We cannot find the original nonce you are attempting to upgrade for 3D Secure.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94502:
		setFailoverMerchantMessage(detail, "Nonce is already consumed - This nonce was previously consumed and cannot be used to upgrade 3D Secure.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94503:
		setFailoverMerchantMessage(detail, "Cannot 3D Secure a non-credit card payment instrument - 3D Secure is not applicable because the transaction was created with an alternative payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode94504:
		setFailoverMerchantMessage(detail, "Unsupported card type - 3D Secure was attempted using a card brand that is not supported by your merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode94505:
		setFailoverMerchantMessage(detail, "Amount must be a number - The amount you specified must be a number greater than 0.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94506:
		setFailoverMerchantMessage(detail, "Amount too large - The amount exceeds the processor's transaction limit.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94507:
		setFailoverMerchantMessage(detail, "Nonce is already 3D Secure - Nonce is already 3D Secure.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94508:
		setFailoverMerchantMessage(detail, "An unexpected error occurred - Unexpected failure during Cardinal API operation.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode94509:
		setFailoverMerchantMessage(detail, "Could not locate 3D Secure verification - Cannot find a 3D Secure verification matching the ID passed to the server.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94510:
		setFailoverMerchantMessage(detail, "Could not locate payment method nonce - An invalid payment method nonce was sent.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94511:
		setFailoverMerchantMessage(detail, "Payment method nonce has not been upgraded to a 3D Secure nonce - Payment method nonce has not been upgraded to a 3D Secure nonce.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94512:
		setFailoverMerchantMessage(detail, "Cardinal response missing required field: cavv - Cardinal response missing required field: cavv")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode94513:
		setFailoverMerchantMessage(detail, "Cardinal response missing required field: xid - Cardinal response missing required field: xid")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode94514:
		setFailoverMerchantMessage(detail, "Cardinal response missing required field: eci_flag - Cardinal response missing required field: eci_flag")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode94515:
		setFailoverMerchantMessage(detail, "Merchant account not 3D Secure enabled - Merchant account is not 3D Secure enabled.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode94516:
		setFailoverMerchantMessage(detail, "Billing address first name is too long. - Deprecated Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94517:
		setFailoverMerchantMessage(detail, "Billing address last name is too long. - Deprecated Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94518:
		setFailoverMerchantMessage(detail, "Billing address line one is too long. - Deprecated Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94519:
		setFailoverMerchantMessage(detail, "Billing address line two is too long. - Deprecated Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94520:
		setFailoverMerchantMessage(detail, "Billing address city is too long. - Deprecated Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94521:
		setFailoverMerchantMessage(detail, "Customer email format is invalid. - Deprecated Customer email format is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94522:
		setFailoverMerchantMessage(detail, "Customer email is too long. - Deprecated Maximum 254 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94523:
		setFailoverMerchantMessage(detail, "Customer mobile phone number format is invalid. - Deprecated Phone must be 2 - 20 characters and can only contain numbers.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94524:
		setFailoverMerchantMessage(detail, "Billing phone number format is invalid. - Deprecated Phone must be 2 - 20 characters and can only contain numbers.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94525:
		setFailoverMerchantMessage(detail, "Customer shipping method code is invalid. - Deprecated Valid customer shipping method codes include '01', '02', '03', '04', '05', and '06'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94526:
		setFailoverMerchantMessage(detail, "Billing Alpha2 country code is not accepted. - Deprecated Must be valid two-letter country codeand formatted as 'xx'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94527:
		setFailoverMerchantMessage(detail, "Billing postal code format is invalid. - Deprecated Billing postal code format is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94528:
		setFailoverMerchantMessage(detail, "Billing state code is invalid. - Deprecated Billing code must be two alphabetical characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94529:
		setFailoverMerchantMessage(detail, "Billing given name format is invalid. - Must be a string. ASCII printable characters only. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94530:
		setFailoverMerchantMessage(detail, "Billing surname format is invalid. - Must be a string. ASCII printable characters only. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94531:
		setFailoverMerchantMessage(detail, "Billing line1 format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94532:
		setFailoverMerchantMessage(detail, "Billing line2 format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94533:
		setFailoverMerchantMessage(detail, "The authenticate result JWT signature did not verify. - There was an error when decoding the JSON Web Token (JWT). Either the signature did not verify, or an invalid JWT was passed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode94534:
		setFailoverMerchantMessage(detail, "Billing line3 format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94535:
		setFailoverMerchantMessage(detail, "Billing city format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94536:
		setFailoverMerchantMessage(detail, "Billing state format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94537:
		setFailoverMerchantMessage(detail, "Billing postal code format is invalid. - Must be a string. Maximum 10 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94538:
		setFailoverMerchantMessage(detail, "Billing country code format is invalid. - Must be a valid two-letter country code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94539:
		setFailoverMerchantMessage(detail, "Billing phone number format is invalid. - Must be a number. Maximum 20 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94540:
		setFailoverMerchantMessage(detail, "Work phone number format is invalid. - Must be a number. Maximum 25 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94541:
		setFailoverMerchantMessage(detail, "Mobile phone number format is invalid. - Must be a number. Maximum 25 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94542:
		setFailoverMerchantMessage(detail, "Email format is invalid. - Must be a string. Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94543:
		setFailoverMerchantMessage(detail, "Shipping method format is invalid. - Possible values: 01, 02, 03, 04, 05, 06.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94544:
		setFailoverMerchantMessage(detail, "Shipping line1 format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94545:
		setFailoverMerchantMessage(detail, "Shipping line2 format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94546:
		setFailoverMerchantMessage(detail, "Shipping line3 format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94547:
		setFailoverMerchantMessage(detail, "Shipping state format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94548:
		setFailoverMerchantMessage(detail, "Shipping city format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94549:
		setFailoverMerchantMessage(detail, "Shipping country code format is invalid. - Must be a valid two-letter country code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94550:
		setFailoverMerchantMessage(detail, "Shipping postal code format is invalid. - Must be a string. Maximum 10 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94551:
		setFailoverMerchantMessage(detail, "Shipping method indicator format is invalid. - Possible values: 01, 02, 03, 04, 05, 06, 07.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94552:
		setFailoverMerchantMessage(detail, "Authentication indicator format is invalid. - Possible values: 01, 02, 03, 04, 05, 06.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94553:
		setFailoverMerchantMessage(detail, "Product code format is invalid. - Possible values: AIR, GEN, DIG, SVC, RES, TRA, DSP, REN, GAS, LUX, ACC, TBD.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94554:
		setFailoverMerchantMessage(detail, "Delivery timeframe format is invalid. - Possible values: 01, 02, 03, 04.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94555:
		setFailoverMerchantMessage(detail, "Delivery email format is invalid. - Must be a string. Maximum 254 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94556:
		setFailoverMerchantMessage(detail, "Reorder indicator format is invalid. - Possible values: 01, 02.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94557:
		setFailoverMerchantMessage(detail, "Preorder indicator format is invalid. - Possible values: 01, 02.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94558:
		setFailoverMerchantMessage(detail, "Preorder date format is invalid. - Must be a date in format 20190131.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94559:
		setFailoverMerchantMessage(detail, "Gift card amount format is invalid. - Must be a number. Maximum 15 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94560:
		setFailoverMerchantMessage(detail, "Gift card currency code format is invalid. - Must be a valid currency code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94561:
		setFailoverMerchantMessage(detail, "Gift card count format is invalid. - Must be a number. Maximum 2 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94562:
		setFailoverMerchantMessage(detail, "Account age indicator format is invalid. - Possible values: 01, 02, 03, 04, 05.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94563:
		setFailoverMerchantMessage(detail, "Account create date format is invalid. - Must be a date in format 20190131.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94564:
		setFailoverMerchantMessage(detail, "Account change indicator format is invalid. - Possible values: 01, 02, 03, 04.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94565:
		setFailoverMerchantMessage(detail, "Account change date format is invalid. - Must be a date in format 20190131.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94566:
		setFailoverMerchantMessage(detail, "Account pwd change indicator format is invalid. - Possible values: 01, 02, 03, 04, 05.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94567:
		setFailoverMerchantMessage(detail, "Account pwd change date format is invalid. - Must be a date in format 20190131.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94568:
		setFailoverMerchantMessage(detail, "Shipping address usage indicator format is invalid. - Possible values: 01, 02, 03, 04.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94569:
		setFailoverMerchantMessage(detail, "Shipping address usage date format is invalid. - Must be a date in format 20190131.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94570:
		setFailoverMerchantMessage(detail, "Transaction count day format is invalid. - Must be a number. Maximum 3 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94571:
		setFailoverMerchantMessage(detail, "Transaction count year format is invalid. - Must be a number. Maximum 3 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94572:
		setFailoverMerchantMessage(detail, "Account purchases format is invalid. - Must be a number. Maximum 4 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94573:
		setFailoverMerchantMessage(detail, "Fraud activity format is invalid. - Possible values: 01, 02.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94574:
		setFailoverMerchantMessage(detail, "Shipping name indicator format is invalid. - Possible values: 01, 02.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94575:
		setFailoverMerchantMessage(detail, "Payment account indicator format is invalid. - Possible values: 01, 02, 03, 04, 05.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94576:
		setFailoverMerchantMessage(detail, "Payment account age format is invalid. - Must be a date in format 20190131.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94578:
		setFailoverMerchantMessage(detail, "ACS window size format is invalid. - Possible values: 01, 02, 03, 04, 05.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94579:
		setFailoverMerchantMessage(detail, "SDK max timeout format is invalid. - Must be a number. Maximum 2 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94581:
		setFailoverMerchantMessage(detail, "Address match format is invalid. - Possible values: Y, N.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94582:
		setFailoverMerchantMessage(detail, "Account ID format is invalid. - Must be a string. Maximum 64 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94584:
		setFailoverMerchantMessage(detail, "IP address format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94586:
		setFailoverMerchantMessage(detail, "Order description format is invalid. - Must be a string. Maximum 256 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94587:
		setFailoverMerchantMessage(detail, "Shipping given name format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94588:
		setFailoverMerchantMessage(detail, "Shipping surname format is invalid. - Must be a string. Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94589:
		setFailoverMerchantMessage(detail, "Shipping phone format is invalid. - Must be a number. Maximum 20 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94590:
		setFailoverMerchantMessage(detail, "Tax amount format is invalid. - Must be a number. Maximum 20 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94591:
		setFailoverMerchantMessage(detail, "Override payment method format is invalid. - Possible values: NA, CR, DB, VSAVR, VSAVA.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94592:
		setFailoverMerchantMessage(detail, "User agent format is invalid. - Must be a string. Maximum 500 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94598:
		setFailoverMerchantMessage(detail, "Installment format is invalid. - Must be a number. Maximum 3 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94599:
		setFailoverMerchantMessage(detail, "Purchase date format is invalid. - Must be a date in format YYYYMMDDHHMMSS.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode94701:
		setFailoverMerchantMessage(detail, "Customer Browser is too long. - Maximum 255 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95601:
		setFailoverMerchantMessage(detail, "Transaction must be submitted for settlement for the authorized amount. - Transaction cannot be adjusted and must be submitted for settlement for the authorized amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95602:
		setFailoverMerchantMessage(detail, "Failed to submit for settlement for an amount different than the authorized amount. Please submit for the authorized amount. - By submitting an amount other than the authorized amount, an adjustment needed to be performed on the transaction. This adjustment resulted in a hard decline and should not be retried.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95603:
		setFailoverMerchantMessage(detail, "Failed to submit for settlement for an amount different than the authorized amount. Please try again at a later time. - By submitting an amount other than the authorized amount, an adjustment needed to be performed on the transaction. This adjustment resulted in a soft decline and can be retried.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95604:
		setFailoverMerchantMessage(detail, "Failed to submit for settlement for an amount different than the authorized amount too many times. Please submit for the authorized amount. - Deprecated By submitting an amount other than the authorized amount, an adjustment needed to be performed on the transaction. This adjustment has resulted in too many failures. Please try again later or submit for the authorized amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95701:
		setFailoverMerchantMessage(detail, "Evidence can only be attached to disputes that are in an Open state - Evidence can only be attached to disputes that are in an Open state.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95702:
		setFailoverMerchantMessage(detail, "Evidence can only be removed from disputes that are in an Open state - Evidence can only be removed from disputes that are in an Open state.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95703:
		setFailoverMerchantMessage(detail, "A document with kind other than Braintree::DocumentUpload::Kind::EvidenceDocument cannot be added to the dispute - A document with kind other than Braintree::DocumentUpload::Kind::EvidenceDocument cannot be added to the dispute.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95704:
		setFailoverMerchantMessage(detail, "Disputes can only be accepted when they are in an Open state - Disputes can only be accepted when they are in an Open state.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95705:
		setFailoverMerchantMessage(detail, "Disputes can only be finalized when they are in an Open state - Disputes can only be finalized when they are in an Open state.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95706:
		setFailoverMerchantMessage(detail, "The category you supplied on the evidence record is not valid - Invalid category provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95707:
		setFailoverMerchantMessage(detail, "Categorized evidence for date time categories must be in the ISO 8601format (e.g. YYYY-MM-DDTHH:MM:SSZ). - Invalid categorized evidence date provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95708:
		setFailoverMerchantMessage(detail, "Categorized text evidence must be 50 characters or less - Maximum 50 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95709:
		setFailoverMerchantMessage(detail, "ARN text evidence must be 30 characters or less - Acquirer reference numbers/trace IDs must be a maximum of 30 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95710:
		setFailoverMerchantMessage(detail, "Prior transaction phone number must be 20 characters or less - Maximum 20 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95711:
		setFailoverMerchantMessage(detail, "Only text evidence can be provided for this category - Only text evidence can be provided for this category.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95712:
		setFailoverMerchantMessage(detail, "Only document evidence can be provided for this category - Only document evidence can be provided for this category.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode95713:
		setFailoverMerchantMessage(detail, "Category cannot be provided for this dispute's reason code - Category cannot be provided for this dispute's reason code; see Evidence Requirements in the Disputes guide for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode95714:
		setFailoverMerchantMessage(detail, "Only one evidence record with this category may exist on the dispute - Only one evidence record with this category may exist on the dispute.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode95715:
		setFailoverMerchantMessage(detail, "The email provided is not valid - The email provided is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95716:
		setFailoverMerchantMessage(detail, "This file already has been used on the dispute - This file has already been used as evidence on this dispute.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95717:
		setFailoverMerchantMessage(detail, "Submitting proof of delivery requires a valid AVS match - Submitting proof of delivery requires that the disputed transaction has a valid AVS postal code match (M); see AVS and CVV Response Codes for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode95720:
		setFailoverMerchantMessage(detail, "Digital goods responses require two types of additional proof - At least two types of evidence categories groups are required when responding with a DOWNLOAD_DATE_TIME; see Evidence Requirements in the Disputes guide for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode95721:
		setFailoverMerchantMessage(detail, "Digital goods responses require a download date - DOWNLOAD_DATE_TIME required when responding with any digital goods validation categories.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode95722:
		setFailoverMerchantMessage(detail, "ARN required for prior transaction responses - When providing any prior transaction category responses you must provide the PRIOR_NON_DISPUTED_TRANSACTION_ID or PRIOR_NON_DISPUTED_TRANSACTION_ARN.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95723:
		setFailoverMerchantMessage(detail, "Date time required for prior transaction responses - When providing any prior transaction category responses you must provide the PRIOR_NON_DISPUTED_TRANSACTION_DATE_TIME.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95724:
		setFailoverMerchantMessage(detail, "Date time required for recurring transaction responses - RECURRING_TRANSACTION_DATE_TIMEcategorized evidence required when providing a RECURRING_TRANSACTION_ID or RECURRING_TRANSACTION_ARN.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95725:
		setFailoverMerchantMessage(detail, "ARN required for recurring transaction responses - RECURRING_TRANSACTION_ID or RECURRING_TRANSACTION_ARNcategorized evidence required when providing a RECURRING_TRANSACTION_DATE_TIME.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95726:
		setFailoverMerchantMessage(detail, "Valid evidence must be added before finalizing the dispute - You must provide valid evidence for this dispute; see Evidence Requirements in the Disputes guide for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95727:
		setFailoverMerchantMessage(detail, "Prior undisputed transaction responses require one or more optional evidence categories - You must provide at least one optional type of evidence when responding with a PRIOR_NON_DISPUTED_TRANSACTION_ID (or PRIOR_NON_DISPUTED_TRANSACTION_ARN) and PRIOR_NON_DISPUTED_TRANSACTION_DATE_TIME; see Evidence Requirements in the Disputes guide for more details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95801:
		setFailoverMerchantMessage(detail, "Commodity code is too long. - Commodity code can't be longer than 12 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95803:
		setFailoverMerchantMessage(detail, "Description is too long. - Description can't be longer than 127 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95804:
		setFailoverMerchantMessage(detail, "Discount amount is an invalid format. - Discount amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the amount can't include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95805:
		setFailoverMerchantMessage(detail, "Discount amount is too large. - Discount amount is too large.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95806:
		setFailoverMerchantMessage(detail, "Discount amount cannot be negative. - Discount amount can't be less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95807:
		setFailoverMerchantMessage(detail, "Kind must be 'debit' or 'credit'. - Kind must be 'debit' or 'credit'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95808:
		setFailoverMerchantMessage(detail, "Kind is required. - Kind is required, and must be 'debit' or 'credit'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95809:
		setFailoverMerchantMessage(detail, "Product code is too long. - Product code can't be longer than 12 characters, or 127 characters for PayPal transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95810:
		setFailoverMerchantMessage(detail, "Quantity is in an invalid format. - Quantity must be formatted like '10' or '10.0000'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95811:
		setFailoverMerchantMessage(detail, "Quantity is required. - Quantity is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95812:
		setFailoverMerchantMessage(detail, "Quantity is too large. - Quantity can't be longer than 12 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95813:
		setFailoverMerchantMessage(detail, "Total amount is an invalid format. - Total amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the amount can't include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95814:
		setFailoverMerchantMessage(detail, "Total amount is required. - Total amount is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95815:
		setFailoverMerchantMessage(detail, "Total amount is too large. - Total amount can't be longer than 9 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95816:
		setFailoverMerchantMessage(detail, "Total amount must be greater than zero. - Total amount can't be less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95817:
		setFailoverMerchantMessage(detail, "Unit amount is an invalid format. - Unit amount must be formatted like '10', '10.00', or '10.0000'.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95818:
		setFailoverMerchantMessage(detail, "Unit amount is required. - Unit amount is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95819:
		setFailoverMerchantMessage(detail, "Unit amount is too large. - Unit amount can't be longer than 12 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95820:
		setFailoverMerchantMessage(detail, "Unit amount must be greater than zero. - Unit amount can't be less than or equal to zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95821:
		setFailoverMerchantMessage(detail, "Unit of measure is too long. - Unit of measure cannot be longer than 12 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95822:
		setFailoverMerchantMessage(detail, "Name is required. - Name is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95823:
		setFailoverMerchantMessage(detail, "Name is too long. - Maximum 35 characters or 127 characters for PayPal transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95824:
		setFailoverMerchantMessage(detail, "Unit tax amount is an invalid format. - Unit tax amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the unit tax amount can't include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95825:
		setFailoverMerchantMessage(detail, "Unit tax amount is too large. - Unit tax amount is too large.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95826:
		setFailoverMerchantMessage(detail, "Unit tax amount cannot be negative. - Unit tax amount can't be less than or equal to zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95827:
		setFailoverMerchantMessage(detail, "Tax amount is an invalid format. - Tax amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the unit tax amount can't include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95828:
		setFailoverMerchantMessage(detail, "Tax amount is too large. - Tax amount can't be longer than 9 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode95829:
		setFailoverMerchantMessage(detail, "Tax amount cannot be negative. - Tax amount can't be less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode815141:
		setFailoverMerchantMessage(detail, "iDEAL payment is not complete - An iDEAL payment needs to enter a 'COMPLETE' state before creating a transaction. This error still exists in the gateway but is no longer returned. It will be removed from the docs once the error has been removed from gateway.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode815142:
		setFailoverMerchantMessage(detail, "iDEAL payment is already consumed - Deprecated in favor of 91564 – Cannot use a payment_method_nonce more than once. This error still exists in the gateway but is no longer returned. It will be removed from the docs once the error has been removed from gateway.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915100:
		setFailoverMerchantMessage(detail, "Shared Customer ID is invalid. - Shared Customer ID is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915101:
		setFailoverMerchantMessage(detail, "Payment instrument type is not accepted. - Payment instrument type is not accepted.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915102:
		setFailoverMerchantMessage(detail, "Partial settlements are not supported by this processor. - Partial settlements are not supported by this processor.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode915103:
		setFailoverMerchantMessage(detail, "Cannot submit for partial settlement. - Cannot submit for partial settlement.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915104:
		setFailoverMerchantMessage(detail, "Delayed settlements are not supported for this processor. The submit for settlement option is required. - Delayed settlements are not supported for this processor. The submit for settlement option is required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode915105:
		setFailoverMerchantMessage(detail, "Merchant account does not support Amex rewards. - Merchant account does not support Amex rewards.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode915106:
		setFailoverMerchantMessage(detail, "Points amount is too large. - Amex Pay with Points amount is too large.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915107:
		setFailoverMerchantMessage(detail, "Updating order_id on submit_for_settlement is not supported by this processor. - Updating order_id on submit_for_settlement is not supported by this processor.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915108:
		setFailoverMerchantMessage(detail, "Updating descriptor on submit_for_settlement is not supported by this processor. - Updating descriptor on submit_for_settlement is not supported by this processor.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915109:
		setFailoverMerchantMessage(detail, "PayPal supplementary data fields must be less than 4001 characters in length: - PayPal supplementary data fields must be less than 4001 characters in length.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915110:
		setFailoverMerchantMessage(detail, "Cannot clone facilitated transactions. - Cannot clone facilitated transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915111:
		setFailoverMerchantMessage(detail, "PayPal supplementary data field count must be less than 101. - PayPal supplementary data field count must be less than 101.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915112:
		setFailoverMerchantMessage(detail, "Shared payment method token originated from another merchant and is not allowed to be shared - Shared payment method token originated from another merchant and is not allowed to be shared.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915113:
		setFailoverMerchantMessage(detail, "EciFlag is required. - ECI flag is required for 3D Secure Pass Thru transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode915114:
		setFailoverMerchantMessage(detail, "EciFlag is invalid. - The provided ECI flag is invalid. Valid ECI Flags are: 00 01 02 05 06 07")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode915116:
		setFailoverMerchantMessage(detail, "Cavv is required for specified EciFlag. - CAVV is required for the specified ECI flag.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode915119:
		setFailoverMerchantMessage(detail, "The version of 3D Secure authentication must be composed only of digits and separated by periods (e.g. 1.0.2). - The version of 3D Secure authentication must be composed only of digits and separated by periods (e.g. 1.0.2).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode915131:
		setFailoverMerchantMessage(detail, "Merchant account does not support 3D Secure transactions for card type. - Merchant account has no active processor settings that support 3D Secure for this card type.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode915133:
		setFailoverMerchantMessage(detail, "Transaction source must be either moto, recurring_first, recurring, or unscheduled. - The source passed with the transaction must be moto, recurring_first, recurring, or unscheduled.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915134:
		setFailoverMerchantMessage(detail, "submit_for_settlement is required and must be true. - The payment method for this transaction does not support authorization with delayed settlement. Funds must be submitted for settlement upon creation.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915135:
		setFailoverMerchantMessage(detail, "shared_payment_method_nonce does not contain valid payment instrument type. - shared_payment_method_nonce does not contain valid payment instrument type.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915136:
		setFailoverMerchantMessage(detail, "Payment instrument type is not accepted by this merchant. - Payment instrument type is not accepted by this merchant.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915137:
		setFailoverMerchantMessage(detail, "Cannot clone Braintree Marketplace transactions via the API. - Braintree Marketplace transactions cannot be cloned via the API. If you wish to clone a Braintree Marketplace transaction, you must use the Braintree Control Panel.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915143:
		setFailoverMerchantMessage(detail, "Merchant account must match the iDEAL payment merchant account. - A transaction for an iDEAL payment must have a merchant account matching the iDEAL payment. This error still exists in the gateway but is no longer returned. It will be removed from the docs once the error has been removed from gateway.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915144:
		setFailoverMerchantMessage(detail, "Amount must match the iDEAL payment amount. - A transaction for an iDEAL payment must have an amount matching the iDEAL payment. This error still exists in the gateway but is no longer returned. It will be removed from the docs once the error has been removed from gateway.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915147:
		setFailoverMerchantMessage(detail, "Vaulted cards from this payment method can only be used for recurring transactions. - This payment method only allows recurring transactions from vaulted cards. All other transactions are not allowed with vaulted cards.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915148:
		setFailoverMerchantMessage(detail, "Transaction can not be voided if status of a PayPal partial settlement child transaction is settlement_pending. - Transaction can not be voided if status of a PayPal partial settlement child transaction is settlement_pending.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915149:
		setFailoverMerchantMessage(detail, "Too many concurrent attempts to refund this transaction. Try again later. - Too many concurrent attempts to refund this transaction. Try again later.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PROCESSING
	case api.RCode915150:
		setFailoverMerchantMessage(detail, "iDEAL payments cannot be vaulted. - iDEAL payments cannot be vaulted. This error still exists in the gateway but is no longer returned. It will be removed from the docs once the error has been removed from gateway.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915151:
		setFailoverMerchantMessage(detail, "Too many concurrent attempts to void this transaction. Try again later. - Too many concurrent attempts to void this transaction. Try again later.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PROCESSING
	case api.RCode915152:
		setFailoverMerchantMessage(detail, "request_id is required in a rewards payment request. - request_id is required in a rewards payment request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915153:
		setFailoverMerchantMessage(detail, "points is required in a rewards payments required. - points is required in a rewards payments required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915154:
		setFailoverMerchantMessage(detail, "currency_amount is required in a rewards payments required. - currency_amount is required in a rewards payments required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915155:
		setFailoverMerchantMessage(detail, "currency_iso_code is required in a rewards payments required. - currency_iso_code is required in a rewards payments required.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915156:
		setFailoverMerchantMessage(detail, "Third party vault status must be either vaulted or will_vault. - The vault status indicator passed with the transaction must be vaulted or will_vault.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915157:
		setFailoverMerchantMessage(detail, "Too many line items. - A maximum of 249 line items can be provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915158:
		setFailoverMerchantMessage(detail, "Expected a collection of line items but none provided. - Expected a collection of line items but unrecognized data provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915159:
		setFailoverMerchantMessage(detail, "Discount amount is an invalid format. - Discount amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the amount can't include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915160:
		setFailoverMerchantMessage(detail, "Discount amount cannot be negative. - Discount amount can't be less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915161:
		setFailoverMerchantMessage(detail, "Discount amount is too large. - Discount amount can't be longer than 9 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915162:
		setFailoverMerchantMessage(detail, "Shipping amount is an invalid format. - Shipping amount must be formatted like '10' or '10.00'. If the currency does not use decimal places, the amount can't include decimal places.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915163:
		setFailoverMerchantMessage(detail, "Shipping amount cannot be negative. - Shipping amount can't be less than zero.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915164:
		setFailoverMerchantMessage(detail, "Shipping amount is too large. - Shipping amount can't be longer than 9 digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915165:
		setFailoverMerchantMessage(detail, "Ships from postal code is too long. - Ships from postal code must not contain more than 9 letters and numbers.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915166:
		setFailoverMerchantMessage(detail, "Ships from postal code must be a string. - Ships from postal code must be a string.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915167:
		setFailoverMerchantMessage(detail, "Ships from Postal code can only contain letters, numbers, spaces, and hyphens. - Ships from Postal code can only contain letters, numbers, spaces, and hyphens.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915168:
		setFailoverMerchantMessage(detail, "Device session id must be less than 256 characters. - The device session id provided is too long. It must be less than 256 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915169:
		setFailoverMerchantMessage(detail, "Merchant account does not support 3D Secure. - The transacting merchant account must be configured to support 3D Secure. Either specify a merchant account that has 3D Secure enabled, or contact us.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode915170:
		setFailoverMerchantMessage(detail, "Venmo Profile ID does not exist. - The Venmo Profile associated with the profile_id does not exist.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915171:
		setFailoverMerchantMessage(detail, "US bank account payment_method_nonce must originate from Plaid or be vaulted prior to transaction. - US bank accounts must be verified before they can be charged. This can be done by either tokenizing with Plaid or vaulting the nonce and providing verification options.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915172:
		setFailoverMerchantMessage(detail, "US bank account payment method must be verified prior to transaction. - US bank accounts must be verified before they can be charged. This can be done by updating the payment method and providing verification options.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915173:
		setFailoverMerchantMessage(detail, "PayPal payee ID is not permitted for an un-vaulted PayPal nonce. - Use payee email for un-vaulted PayPal nonces.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915174:
		setFailoverMerchantMessage(detail, "PayPal payee ID is not permitted for a PayPal Order payment method. - Use payee email for PayPal Order payment methods.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915175:
		setFailoverMerchantMessage(detail, "Status must be either vaulted or will_vault. - See external_vault.statusdescription for more information on the status.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915176:
		setFailoverMerchantMessage(detail, "External vault may only be used with credit card transactions. - An external vault may only be used with credit card transactions. The external vault is used to communicate how a credit card is being stored externally from Braintree.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915177:
		setFailoverMerchantMessage(detail, "Status must be vaulted when used with previous network transaction ID. - The network transaction ID is only needed when conducting a transaction using a stored credit card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915178:
		setFailoverMerchantMessage(detail, "Previous network transaction ID can only be used with Visa cards. - Network transaction IDs are currently only applicable to Visa credit cards.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915179:
		setFailoverMerchantMessage(detail, "Previous network transaction ID must be a 15 digit number. - Previous network transaction ID must be a 15 digit number.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915180:
		setFailoverMerchantMessage(detail, "Merchant account id must match the subscription's merchant account. - Merchant account ID must match the subscription's merchant account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915181:
		setFailoverMerchantMessage(detail, "Shared payment method nonce originated from another merchant and is not allowed to be shared - Shared payment method nonce originated from another merchant and is not allowed to be shared.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915183:
		setFailoverMerchantMessage(detail, "Cannot void a transaction that has already been captured. - PayPal has already captured funds for this transaction, so it can't be voided in the Braintree gateway.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915189:
		setFailoverMerchantMessage(detail, "Merchant account must match the merchant account used to create the payment. - Merchant account provided in the Transaction.sale() call must match the merchant account used to create the payment by the client.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode915190:
		setFailoverMerchantMessage(detail, "Fake payment method nonce cannot be used in production. - Test nonces can only be used in the sandbox.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915191:
		setFailoverMerchantMessage(detail, "Settlement failed with a soft processor decline. - The processor failed to capture funds for this transaction, but the authorization is still open. You may safely retry this request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
	case api.RCode915192:
		setFailoverMerchantMessage(detail, "Local payment transaction has already been processed. - Local payment transaction has already been processed for the parameters provided in the Transaction.sale()call.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode915996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode915999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode916996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode916997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode916998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode916999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode917996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode917997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode917998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode917999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode918996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode918997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode918998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode918999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode919996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode919997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode919998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode919999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode920996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode920997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode920998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode920999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode922996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode922997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode922998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode922999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode923996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode923997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode923998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode923999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode926996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode926997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode926998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode926999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode928996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode928997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode928998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode928999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode929996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode929997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode929998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode929999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode931996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode931997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode931998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode931999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode934996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode934997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode934998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode934999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode935996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode935997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode935998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode935999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode938996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode938997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode938998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode938999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode941996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode941997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode941998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode941999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode942996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode942997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode942998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode942999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945100:
		setFailoverMerchantMessage(detail, "Recurring end format is invalid. - Must be a date in format 20190131.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945101:
		setFailoverMerchantMessage(detail, "Recurring frequency format is invalid. - Must be a number. Maximum 3 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945102:
		setFailoverMerchantMessage(detail, "Add card attempts format is invalid. - Must be a number. Maximum 3 characters.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945103:
		setFailoverMerchantMessage(detail, "One of the additional info fields sent is invalid. Check the field-level docs in the latest client SDK references for correct field names. - For field-level requirements, see the latest Android, iOS, or JavaScriptreference docs for the SDK version you're using.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945104:
		setFailoverMerchantMessage(detail, "The additional info sent is invalid. Check the field-level docs in the latest client SDK references for format requirements. - For field-level requirements, see the latest Android, iOS, or JavaScriptreference docs for the SDK version you're using.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945105:
		setFailoverMerchantMessage(detail, "The customer information sent is invalid. Check the field-level docs in the latest client SDK references for format requirements. - For field-level requirements, see the latest Android, iOS, or JavaScriptreference docs for the SDK version you're using.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945106:
		setFailoverMerchantMessage(detail, "The region cannot be provided without a corresponding country code. - A country code is necessary to disambiguate the provided region name.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945107:
		setFailoverMerchantMessage(detail, "Requesting a challenge and requesting an exemption are mutually exclusive options. - You cannot request a challenge and an exemption on the same request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode945999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode947996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode947997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode947998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode947999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode956996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode956997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode956998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode956999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode957996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode957997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode957998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode957999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode958996:
		setFailoverMerchantMessage(detail, "Required attribute is missing - The missing attribute is provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode958997:
		setFailoverMerchantMessage(detail, "Attribute is not in the required format - The attribute and the expected format are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode958998:
		setFailoverMerchantMessage(detail, "Attribute is not in the list of expected values - The attribute and the expected values are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode958999:
		setFailoverMerchantMessage(detail, "Attribute is the wrong type - The attribute and the expected types are provided in the error message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCodeRangeStart2101:
		setFailoverMerchantMessage(detail, "Processor Declined - The customer's bank is unwilling to accept the transaction. The customer will need to contact their bank for more details regarding this generic decline.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCodeRangeEnd2999:
		setFailoverMerchantMessage(detail, "Processor Declined - The customer's bank is unwilling to accept the transaction. The customer will need to contact their bank for more details regarding this generic decline.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	}
}