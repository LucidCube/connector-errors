package connector_errors

import "github.com/chargehive/proto/golang/chargehive/chtype"

func(){

	switch code {
	case api.RCode5001000:
		setFailoverMerchantMessage(detail, "An internal error occurred.")
		detail.FailureType = chtype.FAILURE_TYPE_INTERNAL
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode5021001:
		setFailoverMerchantMessage(detail, "An error occurred with the external processing gateway.")
		detail.FailureType = chtype.FAILURE_TYPE_INTERNAL
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode500503:
		setFailoverMerchantMessage(detail, "An unknown error has occurred.")
		detail.FailureType = chtype.FAILURE_TYPE_INTERNAL
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode5001002:
		setFailoverMerchantMessage(detail, "An internal error occurred.")
		detail.FailureType = chtype.FAILURE_TYPE_INTERNAL
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode5001003:
		setFailoverMerchantMessage(detail, "An internal error occurred.")
		detail.FailureType = chtype.FAILURE_TYPE_INTERNAL
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode5001007:
		setFailoverMerchantMessage(detail, "An internal error occurred.")
		detail.FailureType = chtype.FAILURE_TYPE_INTERNAL
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode5001008:
		setFailoverMerchantMessage(detail, "An internal error occurred.")
		detail.FailureType = chtype.FAILURE_TYPE_INTERNAL
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode4291200:
		setFailoverMerchantMessage(detail, "The API call has been denied as it has exceeded the permissible call rate limit.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode4003001:
		setFailoverMerchantMessage(detail, "You submitted an unsupported card type with your request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4003002:
		setFailoverMerchantMessage(detail, "You submitted an invalid card number or brand or combination of card number and brand with your request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
		detail.ErrorProperties["card.number"] = "card.number"
	case api.RCode4003004:
		setFailoverMerchantMessage(detail, "The zip/postal code must be provided for an AVS check request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
		detail.ErrorProperties["postal"] = "postal"
	case api.RCode4003005:
		setFailoverMerchantMessage(detail, "You submitted an incorrect CVV value with your request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_CVV
	case api.RCode4003006:
		setFailoverMerchantMessage(detail, "You submitted an expired credit card number with your request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4023007:
		setFailoverMerchantMessage(detail, "Your request has failed the AVS check. Note that the amount has still been reserved on the customer's card and will be released in 3–5 business days. "+
			"Please ensure the billing address is accurate before retrying the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_ADDRESS
	case api.RCode4003008:
		setFailoverMerchantMessage(detail, "You submitted a card type for which the merchant account is not configured.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4023009:
		setFailoverMerchantMessage(detail, "Your request has been declined by the issuing bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4023011:
		setFailoverMerchantMessage(detail, "Your request has been declined by the issuing bank because the card used is a restricted card."+
			"| Contact the cardholder's credit card company for further investigation.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4023012:
		setFailoverMerchantMessage(detail, "Your request has been declined by the issuing bank because the credit card expiry date submitted is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_EXPIRED
		detail.ErrorProperties["card.expiry"] = "card.expiry"
	case api.RCode4023013:
		setFailoverMerchantMessage(detail, "Your request has been declined by the issuing bank due to problems with the credit card account.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4023014:
		setFailoverMerchantMessage(detail, "Your request has been declined - the issuing bank has returned an unknown response."+
			"| Contact the card holder's credit card company for further investigation.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4023015:
		setFailoverMerchantMessage(detail, "The bank has requested that you process the transaction manually by calling the cardholder's credit card company.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PERMISSION
	case api.RCode4023016:
		setFailoverMerchantMessage(detail, "The bank has requested that you retrieve the card from the cardholder – it may be a lost or stolen card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
		detail.ErrorType = chtype.RESPONSE_ERROR_PICKUP
	case api.RCode4023017:
		setFailoverMerchantMessage(detail, "You submitted an invalid credit card number with your request.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
		detail.ErrorProperties["card.number"] = "card.number"
	case api.RCode4023018:
		setFailoverMerchantMessage(detail, "The bank has requested that you retry the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_RETRY
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_RETRY
	case api.RCode4023019:
		setFailoverMerchantMessage(detail, "Your request has failed the CVV check. Please note that the amount may still have been reserved on the customer's card, "+
			"in which case it will be released in 3–5 business days.")
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
		detail.SpecificErrors = append(detail.SpecificErrors, &chtype.ResponseDetail{Category: chtype.RESPONSE_CATEGORY_METHOD, ErrorType: chtype.RESPONSE_ERROR_CVV})
	case api.RCode4023020:
		setFailoverMerchantMessage(detail, "The bank has requested that you retry the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_RETRY
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_RETRY
	case api.RCode4043021:
		setFailoverMerchantMessage(detail, "The confirmation number included in this request could not be found.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4023022:
		setFailoverMerchantMessage(detail, "The card has been declined due to insufficient funds.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_AVAILABLE_FUNDS
	case api.RCode4023023:
		setFailoverMerchantMessage(detail, "Your request has been declined by the issuing bank due to its proprietary card activity regulations.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_VELOCITY
	case api.RCode4023024:
		setFailoverMerchantMessage(detail, "Your request has been declined because the issuing bank does not permit the transaction for this card.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PERMISSION
	case api.RCode4003025:
		setFailoverMerchantMessage(detail, "The external processing gateway has reported invalid data.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4003026:
		setFailoverMerchantMessage(detail, "The external processing gateway has reported the account type is invalid.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4023027:
		setFailoverMerchantMessage(detail, "The external processing gateway has reported a limit has been exceeded.")
		detail.Category = chtype.RESPONSE_CATEGORY_PROCESSING
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode5023028:
		setFailoverMerchantMessage(detail, "The external processing gateway has reported a system error.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
		detail.ErrorType = chtype.RESPONSE_ERROR_SYSTEM
	case api.RCode4023029:
		setFailoverMerchantMessage(detail, "The external processing gateway has rejected the transaction.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4023030:
		setFailoverMerchantMessage(detail, "The external processing gateway has reported the transaction is unauthorized.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PERMISSION
	case api.RCode4003031:
		setFailoverMerchantMessage(detail, "The confirmation number you submitted with your request references a transaction that is not on hold.")
		detail.Category = chtype.RESPONSE_CATEGORY_PROCESSING
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4023032:
		setFailoverMerchantMessage(detail, "Your request has been declined by the issuing bank or external gateway because the card is probably in one of their negative databases.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_FRAUD
	case api.RCode4023033:
		setFailoverMerchantMessage(detail, "Declined due to Authroizaion Revocation.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4023035:
		setFailoverMerchantMessage(detail, "Your request has been declined due to exceeded PIN attempts.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_VELOCITY
	case api.RCode4023036:
		setFailoverMerchantMessage(detail, "Your request has been declined due to an invalid issuer.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4023037:
		setFailoverMerchantMessage(detail, "Your request has been declined because it is invalid.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4023038:
		setFailoverMerchantMessage(detail, "Your request has been declined due to customer cancellation.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4023039:
		setFailoverMerchantMessage(detail, "Your request has been declined due to an invalid authentication value.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4023040:
		setFailoverMerchantMessage(detail, "Your request has been declined because the request type is not permitted on the card.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4023041:
		setFailoverMerchantMessage(detail, "Your request has been declined due to a timeout.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
		detail.ErrorType = chtype.RESPONSE_ERROR_TIMEOUT
	case api.RCode4023042:
		setFailoverMerchantMessage(detail, "Your request has been declined due to a cryptographic error.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
		detail.ErrorType = chtype.RESPONSE_ERROR_SYSTEM
	case api.RCode4023044:
		setFailoverMerchantMessage(detail, "You have submitted a duplicate request.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DUPLICATE
	case api.RCode4023045:
		setFailoverMerchantMessage(detail, "You submitted an invalid date format for this request.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4023046:
		setFailoverMerchantMessage(detail, "The transaction was declined because the amount was set to zero.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4023047:
		setFailoverMerchantMessage(detail, "The transaction was declined because the amount exceeds the floor limit.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode4023048:
		setFailoverMerchantMessage(detail, "The transaction was declined because the amount is less than the floor limit.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode4023049:
		setFailoverMerchantMessage(detail, "The bank has requested that you retrieve the card from the cardholder – the credit card has expired.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PICKUP
		detail.SpecificErrors = append(detail.SpecificErrors, &chtype.ResponseDetail{Category: chtype.RESPONSE_CATEGORY_METHOD, ErrorType: chtype.RESPONSE_ERROR_EXPIRED})
	case api.RCode4023050:
		setFailoverMerchantMessage(detail, "The bank has requested that you retrieve the card from the cardholder – fraudulent activity is suspected.")
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
		detail.ErrorType = chtype.RESPONSE_ERROR_FRAUD
	case api.RCode4023051:
		setFailoverMerchantMessage(detail, "The bank has requested that you retrieve the card from the cardholder – contact the acquirer for more information.")
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
		detail.ErrorType = chtype.RESPONSE_ERROR_PICKUP
	case api.RCode4023052:
		setFailoverMerchantMessage(detail, "The bank has requested that you retrieve the card from the cardholder – the credit card is restricted.")
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4023053:
		setFailoverMerchantMessage(detail, "The bank has requested that you retrieve the card from the cardholder – please call the acquirer.")
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
		detail.ErrorType = chtype.RESPONSE_ERROR_PICKUP
	case api.RCode4023054:
		setFailoverMerchantMessage(detail, "The transaction was declined due to suspected fraud.")
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4003200:
		setFailoverMerchantMessage(detail, "You have submitted an invalidly formatted Authorization ID for this Settlement.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4043201:
		setFailoverMerchantMessage(detail, "The Authorization ID included in this Settlement request could not be found.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4023202:
		setFailoverMerchantMessage(detail, "You have exceeded the maximum number of Settlements allowed.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode4023203:
		setFailoverMerchantMessage(detail, "The Authorization is either fully settled or cancelled.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_ALREADY_DONE
	case api.RCode4023204:
		setFailoverMerchantMessage(detail, "The requested Settlement amount exceeds the remaining Authorization amount.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4023205:
		setFailoverMerchantMessage(detail, "The Authorization you are attempting to settle has expired.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_EXPIRED
	case api.RCode4023206:
		setFailoverMerchantMessage(detail, "The external processing gateway has rejected the transaction.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4023402:
		setFailoverMerchantMessage(detail, "The requested Refund amount exceeds the remaining Settlement amount.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
		detail.FailureType = chtype.FAILURE_TYPE_HARD
	case api.RCode4023403:
		setFailoverMerchantMessage(detail, "You have already processed the maximum number of refunds allowed for this Settlement.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode4023404:
		setFailoverMerchantMessage(detail, "The Settlement has already been fully refunded.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_ALREADY_DONE
		detail.FailureType = chtype.FAILURE_TYPE_HARD
	case api.RCode4023405:
		setFailoverMerchantMessage(detail, "The Settlement you are attempting to Refund has expired.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_EXPIRED
	case api.RCode4023406:
		setFailoverMerchantMessage(detail, "The Settlement you are attempting to Refund has not been batched yet. | There are no settled funds available to Refund.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_UNAVAILABLE
	case api.RCode4003407:
		setFailoverMerchantMessage(detail, "The Settlement referred to by the transaction response ID you provided cannot be found.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4003408:
		setFailoverMerchantMessage(detail, "You have submitted an invalidly formatted response ID for the original Purchase or Settlement.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4023412:
		setFailoverMerchantMessage(detail, "The Refund transaction you attempted was not permitted because your merchant account is in overdraft.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_AVAILABLE_FUNDS
	case api.RCode4023413:
		setFailoverMerchantMessage(detail, "The requested Refund amount exceeds the permissible Visa credit ratio.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode4003414:
		setFailoverMerchantMessage(detail, "The Refund referred to by the transaction response ID you provided cannot be found.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4023415:
		setFailoverMerchantMessage(detail, "You cannot cancel this transaction as it is no longer in a pending state.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4023416:
		setFailoverMerchantMessage(detail, "The external processing gateway for which your merchant account is configured does not support partial Settlements.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4023417:
		setFailoverMerchantMessage(detail, "There is already another request being processed on the transaction referenced for this request.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_DUPLICATE
	case api.RCode4023418:
		setFailoverMerchantMessage(detail, "The external processing gateway for which your merchant account is configured does not support partial Credits.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4003500:
		setFailoverMerchantMessage(detail, "The confirmation number included in this request could not be found.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4023501:
		setFailoverMerchantMessage(detail, "The requested Authorization Reversal amount exceeds the remaining Authorization amount.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode4023502:
		setFailoverMerchantMessage(detail, "You cannot process an Authorization Reversal transaction against an Authorization that has been settled.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_ALREADY_DONE
	case api.RCode4023503:
		setFailoverMerchantMessage(detail, "The Authorization Reversal transaction is not supported for the card type used for the Authorization you are attempting to reverse.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4023504:
		setFailoverMerchantMessage(detail, "The external processing gateway for which your merchant account is configured does not support partial Authorization Reversals.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode5003505:
		setFailoverMerchantMessage(detail, "The Authorization Reversal could not be completed.")
		detail.Category = chtype.RESPONSE_CATEGORY_PROCESSING
		detail.ErrorType = chtype.RESPONSE_ERROR_RETRY
	case api.RCode4023506:
		setFailoverMerchantMessage(detail, "The reversal amount exceeds the remaining amount of the Authorization.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_LIMIT
	case api.RCode4023703:
		setFailoverMerchantMessage(detail, "The external gateway has reported that you have submitted an invalid amount with your request.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4024001:
		setFailoverMerchantMessage(detail, "The card number or email address associated with this transaction is in our negative database.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_FRAUD
	case api.RCode4024002:
		setFailoverMerchantMessage(detail, "The transaction was declined by our Risk Management department.")
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
		detail.ErrorType = chtype.RESPONSE_ERROR_SYSTEM
	case api.RCode4015000:
		setFailoverMerchantMessage(detail, "Your merchant account authentication failed."+
			"| Either your store ID/password are invalid or the IP address from which you are sending the transaction has not been authorized.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4005001:
		setFailoverMerchantMessage(detail, "The submitted currency code is invalid or your account does not support this currency.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4005003:
		setFailoverMerchantMessage(detail, "You submitted an invalid amount with your request.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4005004:
		setFailoverMerchantMessage(detail, "You submitted an invalid account type with your request.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4005005:
		setFailoverMerchantMessage(detail, "You submitted an invalid operation type with your request.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4005010:
		setFailoverMerchantMessage(detail, "The submitted country code is invalid.")
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4005016:
		setFailoverMerchantMessage(detail, "The merchant account you provided cannot be found.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4005017:
		setFailoverMerchantMessage(detail, "The merchant account you provided is disabled.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DISABLED
	case api.RCode4025021:
		setFailoverMerchantMessage(detail, "Your transaction request has been declined.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_DECLINE
	case api.RCode4005023:
		setFailoverMerchantMessage(detail, "The request is not parseable.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4095031:
		setFailoverMerchantMessage(detail, "The transaction you have submitted has already been processed.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_DUPLICATE
	case api.RCode4015040:
		setFailoverMerchantMessage(detail, "Your merchant account is not configured for the transaction you attempted.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_UNSUPPORTED
	case api.RCode4005042:
		setFailoverMerchantMessage(detail, "The merchant reference number is missing or invalid or it exceeds the maximum permissible length.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4005050:
		setFailoverMerchantMessage(detail, "An error occurred with your merchant account configuration")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode4005068:
		setFailoverMerchantMessage(detail, "Either you submitted a request that is missing a mandatory field or the value of a field does not match the format expected.")
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4045269:
		setFailoverMerchantMessage(detail, "The ID(s) specified in the URL do not correspond to the values in the system.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4035270:
		setFailoverMerchantMessage(detail, "The credentials provided with the request do not have permission to access the requested data.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_PERMISSION
	case api.RCode4065271:
		setFailoverMerchantMessage(detail, "You requested a response in the 'Accept' header that is in an unsupported format.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4065272:
		setFailoverMerchantMessage(detail, "The 'Content-Type' you specified in request header was submitted in an unsupported format.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4045273:
		setFailoverMerchantMessage(detail, "Your client reached our application but we were unable to service your request due to an invalid URL.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4015275:
		setFailoverMerchantMessage(detail, "The authentication credentials provided with the request have expired.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_EXPIRED
	case api.RCode4015276:
		setFailoverMerchantMessage(detail, "The authentication credentials provided with the request provided have been disabled.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_DISABLED
	case api.RCode4015277:
		setFailoverMerchantMessage(detail, "The authentication credentials provided with the request have been locked due to multiple authentication failures.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_PERMISSION
	case api.RCode4015278:
		setFailoverMerchantMessage(detail, "The authentication credentials provided with the request were not accepted for an unknown reason.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_UNKNOWN
	case api.RCode4015279:
		setFailoverMerchantMessage(detail, "The authentication credentials are invalid.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4015280:
		setFailoverMerchantMessage(detail, "The required authentication credentials were not provided.")
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4055281:
		setFailoverMerchantMessage(detail, "The request uses an action (e.g., GET, POST, or PUT) that is not supported by the resource.")
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
		detail.ErrorType = chtype.RESPONSE_ERROR_PAYLOAD
	case api.RCode4005501:
		setFailoverMerchantMessage(detail, "The profile does not have an active credit card.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	case api.RCode4005502:
		setFailoverMerchantMessage(detail, "Either the payment token is invalid or the corresponding profile or bank account is not active.")
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
		detail.ErrorType = chtype.RESPONSE_ERROR_NOT_FOUND
	}
}
