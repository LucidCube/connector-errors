package connector_errors

import "github.com/chargehive/proto/golang/chargehive/chtype"

func(){
	switch code {
	case api.RCode000:
		setFailoverMerchantMessage(detail, "Approved - No action required.")
		detail.ErrorType = chtype.RESPONSE_ERROR_NONE
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode001:
		setFailoverMerchantMessage(detail, "Transaction Received - This is sent to acknowledge that the submitted transaction has been received. Note: This response applies only to V10.x versions of the API.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode010:
		setFailoverMerchantMessage(detail, "Partially Approved - The authorized amount is less than the requested amount.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode011:
		setFailoverMerchantMessage(detail, "Offline Approval - Offline approval issued while the terminal is unable to communicate with the issuer.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode013:
		setFailoverMerchantMessage(detail, "Offline Approval (unable to go online) - Offline approval issued while the terminal is unable to communicate with the issuer.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode100:
		setFailoverMerchantMessage(detail, "Processing Network Unavailable - There is a problem with the card or PINless Debit network. Contact the network for more information.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
	case api.RCode101:
		setFailoverMerchantMessage(detail, "Issuer Unavailable - There is a problem with the issuer network. Please contact the issuing bank.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
	case api.RCode102:
		setFailoverMerchantMessage(detail, "Re-submit Transaction - There is a temporary problem with your submission. Please re-submit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_RETRY
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode108:
		setFailoverMerchantMessage(detail, "Try again later - Returned if the eProtect token is not immediately available, when submitting an Auth or Sale transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode110:
		setFailoverMerchantMessage(detail, "Insufficient Funds - The card does not have enough funds to cover the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode111:
		setFailoverMerchantMessage(detail, "Authorization amount has already been depleted - The total amount of the original Authorization has been used. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode120:
		setFailoverMerchantMessage(detail, "Call Issuer - There is an unspecified problem, contact the issuing bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode121:
		setFailoverMerchantMessage(detail, "Call AMEX - There is an unspecified problem; contact AMEX.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode122:
		setFailoverMerchantMessage(detail, "Call Diners Club - There is an unspecified problem; contact Diners Club.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode123:
		setFailoverMerchantMessage(detail, "Call Discover - There is an unspecified problem; contact Discover.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode124:
		setFailoverMerchantMessage(detail, "Call JBS - There is an unspecified problem; contact JBS.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode125:
		setFailoverMerchantMessage(detail, "Call Visa/MasterCard - There is an unspecified problem; contact Visa or MasterCard.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode126:
		setFailoverMerchantMessage(detail, "Call Issuer - Update Cardholder Data - Some data is out of date; contact the issuer to update this information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode127:
		setFailoverMerchantMessage(detail, "Exceeds Approval Amount Limit - This transaction exceeds the daily approval limit for the card.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode130:
		setFailoverMerchantMessage(detail, "Call Indicated Number - There is an unspecified problem; contact the phone number provided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode140:
		setFailoverMerchantMessage(detail, "Update Cardholder Data - Cardholder data is incorrect; contact the issuing bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode150:
		setFailoverMerchantMessage(detail, "Original transaction found. - A Query transaction response indicating that the original transaction was found.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode151:
		setFailoverMerchantMessage(detail, "Original transaction not found. - A Query transaction response indicating that the original transaction was not found.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode152:
		setFailoverMerchantMessage(detail, "Original transaction found, but response not yet available. - A Query transaction response indicating that the original transaction was found, but the final response information is not yet available.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode153:
		setFailoverMerchantMessage(detail, "Query transaction not enabled. - A Query transaction response indicating that you are not enabled for use of the Query transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode154:
		setFailoverMerchantMessage(detail, "At least one of origId or origCnpTxnId is required - When submitting a Query transaction, you must include either the origId or origCnpTxnId.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode155:
		setFailoverMerchantMessage(detail, "origCnpTxnId is required when showStatusOnly is used - When submitting a Query transaction with the showStatusOnly flag set to Y, you must include the origCnpTxnId element.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode191:
		setFailoverMerchantMessage(detail, "The merchant is not registered in the update program. - This is an Account Updater response indicating a set-up problem that must be resolved prior to submitting another request file. Escalate this to your Relationship Manager.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode192:
		setFailoverMerchantMessage(detail, "Merchant not certified/enabled for IIAS - Your organization is not certified or enabled for IIAS/FSA transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode206:
		setFailoverMerchantMessage(detail, "Issuer Generated Error - An unspecified error was returned by the issuer. Please retry the transaction and if the problem persist, contact the issuing bank.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode207:
		setFailoverMerchantMessage(detail, "Pickup card - Other than Lost/Stolen - The issuer indicated that the gift card should be removed from use.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode209:
		setFailoverMerchantMessage(detail, "Invalid Amount - The specified amount is invalid for this transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode211:
		setFailoverMerchantMessage(detail, "Reversal Unsuccessful - The reversal transaction was unsuccessful.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode212:
		setFailoverMerchantMessage(detail, "Missing Data - Contact your Relationship Manager.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode213:
		setFailoverMerchantMessage(detail, "Pickup Card - Lost Card - The submitted card was reported as lost and should be removed from use.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode214:
		setFailoverMerchantMessage(detail, "Pickup Card - Stolen Card - The submitted card was reported as stolen and should be removed from use.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode215:
		setFailoverMerchantMessage(detail, "Restricted Card - The specified Gift Card is not available for use.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode216:
		setFailoverMerchantMessage(detail, "Invalid Deactivate - The Deactivate transaction is invalid for the specified card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode217:
		setFailoverMerchantMessage(detail, "Card Already Active - The submitted card is already active.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode218:
		setFailoverMerchantMessage(detail, "Card Not Active - The submitted card has not been activated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode219:
		setFailoverMerchantMessage(detail, "Card Already Deactivate - The submitted card has already been deactivated.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode221:
		setFailoverMerchantMessage(detail, "Over Max Balance - The activate or load amount exceeds the maximum allowed for the specified gift Card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode222:
		setFailoverMerchantMessage(detail, "Invalid Activate - The activate transaction is not valid or can no longer be reversed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode223:
		setFailoverMerchantMessage(detail, "No transaction Found for Reversal - The transaction referenced in the reversal transaction does not exist.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode226:
		setFailoverMerchantMessage(detail, "Incorrect CVV - The transaction was declined because it was submitted with the incorrect security code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode229:
		setFailoverMerchantMessage(detail, "Illegal Transaction - The transaction would violate the law.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode251:
		setFailoverMerchantMessage(detail, "Duplicate Transaction - The transaction is a duplicate of a previously submitted transaction. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode252:
		setFailoverMerchantMessage(detail, "System Error - Contact your Relationship Manager.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode253:
		setFailoverMerchantMessage(detail, "Deconverted BIN - The BIN is no longer valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode254:
		setFailoverMerchantMessage(detail, "Merchant Depleted - No balance remains on gift Card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode255:
		setFailoverMerchantMessage(detail, "Gift Card Escheated - The Gift Card has been seized by the government while resolving an estate.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode256:
		setFailoverMerchantMessage(detail, "Invalid Reversal Type for Credit Card Transaction - You attempted to use a Closed Loop Gift Card reversal transaction to reverse a credit card transaction. For example, you cannot use a Deposit Reversal transaction to reverse a Capture. To reverse a credit card Capture transaction, use a Credit transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode257:
		setFailoverMerchantMessage(detail, "System Error (message format error) - Issuer reported message format is incorrect. Contact your Relationship Manager.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode258:
		setFailoverMerchantMessage(detail, "System Error (cannot process) - Issuer reported transaction could not be processed. Contact your Relationship Manager.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode271:
		setFailoverMerchantMessage(detail, "Refund rejected due to pending deposit status - The refund is tied to a deposit that is still in pending state or the state is in doubt. You can retry the refund at a later time.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode272:
		setFailoverMerchantMessage(detail, "Refund rejected due to declined deposit status - The refund is tied to a deposit that failed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode273:
		setFailoverMerchantMessage(detail, "Refund rejected by the processing network - The refund is tied to a deposit that succeeded, but was declined by PayPro.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode284:
		setFailoverMerchantMessage(detail, "Capture, Credit and AuthReversal tags cannot be used for Gift Card Transactions - You must use the Gift Card version of these transactions for Gift Cards (i.e., giftCardCapture, giftCardCredit, and giftCardAuthReversal).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode301:
		setFailoverMerchantMessage(detail, "Invalid Account Number - The account number is not valid; contact the cardholder to confirm information or inquire about another form of payment.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode302:
		setFailoverMerchantMessage(detail, "Account Number Does Not Match Payment Type - The payment type was selected as one card type (e.g. Visa), but the card number indicates a different card type (e.g. MasterCard).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode303:
		setFailoverMerchantMessage(detail, "Pick Up Card - This is a card present response, but in a card not present environment. Do not process the transaction and contact the issuing bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode304:
		setFailoverMerchantMessage(detail, "Lost/Stolen Card - The card has been designated as lost or stolen; contact the issuing bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode305:
		setFailoverMerchantMessage(detail, "Expired Card - The card is expired.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode306:
		setFailoverMerchantMessage(detail, "Authorization has expired; no need to reverse - The original Authorization is no longer valid, because it has expired. You can not perform an Authorization Reversal for an expired Authorization. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode307:
		setFailoverMerchantMessage(detail, "Restricted Card - The card has a restriction preventing approval for this transaction. Please contact the issuing bank for a specific reason. You may also receive this code if the transaction was declined due to Prior Fraud Advice Filtering and you are using a schema version V8.10 or older.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode308:
		setFailoverMerchantMessage(detail, "Restricted Card - Chargeback - This transaction is being declined due the operation of the Prior Chargeback Card Filtering Service or the card has a restriction preventing approval if there are any chargebacks against it.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode309:
		setFailoverMerchantMessage(detail, "Restricted Card - Prepaid Card Filtering Service - This transaction is being declined due the operation of the Prepaid Card Filtering service.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode310:
		setFailoverMerchantMessage(detail, "Invalid track data - The track data is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode311:
		setFailoverMerchantMessage(detail, "Deposit is already referenced by a chargeback - The deposit is already referenced by a chargeback; therefore, a refund cannot be processed against the original transaction. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode312:
		setFailoverMerchantMessage(detail, "Restricted Card - International Card Filtering Service - This transaction is being declined due the operation of the International Card Filtering Service.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode313:
		setFailoverMerchantMessage(detail, "International filtering for issuing card country <country> (where <country> is the 3-character country code) - This is returned when the transaction involves a US based merchant processing Canadian transactions has a transaction that uses a US card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode315:
		setFailoverMerchantMessage(detail, "Restricted Card - Auth Fraud Velocity Filtering Service - This transaction is being declined due the operation of the Auth Fraud Velocity Filtering Service.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode316:
		setFailoverMerchantMessage(detail, "Automatic Refund Already Issued - This refund transaction is a duplicate for one already processed automatically by the Fraud Chargeback Prevention Service (FCPS). Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode318:
		setFailoverMerchantMessage(detail, "Restricted Card - Auth Fraud Advice Filtering Service - This transaction is being declined due the operation of the Auth Fraud Advice Filtering Service.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode319:
		setFailoverMerchantMessage(detail, "Restricted Card - Fraud AVS Filtering Service - This transaction is being declined due the operation of the Auth Fraud AVS Filtering Service.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode320:
		setFailoverMerchantMessage(detail, "Invalid Expiration Date - The expiration date is invalid")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode321:
		setFailoverMerchantMessage(detail, "Invalid Merchant - The card is not allowed to make purchases from this merchant (e.g. a Travel only card trying to purchase electronics).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode322:
		setFailoverMerchantMessage(detail, "Invalid Transaction Note: If you are enabled for Transaction Filtering, but have not upgraded to use schema version 8.3 or above, the system returns this code for transactions filtered by the Prepaid or International Card Filtering Service. If you are enabled for Velocity Fraud Filtering, but have not upgraded to V8.9, you will receive this code for filtered transactions. If you are enabled for AVS Fraud Filtering, but have not upgraded to V8.13, you will receive this code for filtered transactions. - The transaction is not permitted; contact the issuing bank. The system also returns this code if you attempt to use a Void transaction to cancel a Gift Card transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode323:
		setFailoverMerchantMessage(detail, "No such issuer - The card number references an issuer that does not exist. Do not process the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode324:
		setFailoverMerchantMessage(detail, "Invalid Pin - The PIN provided is invalid. Appears in Declined Transaction report")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode325:
		setFailoverMerchantMessage(detail, "Transaction not allowed at terminal - The transaction is not permitted; contact the issuing bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode326:
		setFailoverMerchantMessage(detail, "Exceeds number of PIN entries - (Referring to a debit card) The incorrect PIN has been entered excessively and the card is locked.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode327:
		setFailoverMerchantMessage(detail, "Cardholder transaction not permitted - Merchant does not allow that card type or specific transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode328:
		setFailoverMerchantMessage(detail, "Cardholder requested that recurring or installment payment be stopped - Recurring/Installment Payments no longer accepted by the card issuing bank.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode330:
		setFailoverMerchantMessage(detail, "Invalid Payment Type - This payment type is not accepted by the issuer.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode331:
		setFailoverMerchantMessage(detail, "Invalid POS Capability for Cardholder Authorized Terminal Transaction - For a Cardholder Authorized Terminal Transaction the POS capability must be set to magstripe.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode332:
		setFailoverMerchantMessage(detail, "Invalid POS Cardholder ID for Cardholder Authorized Terminal Transaction - For a Cardholder Authorized Terminal Transaction the POS Cardholder ID must be set to nopin.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode335:
		setFailoverMerchantMessage(detail, "This method of payment does not support authorization reversals - You can not perform an Authorization Reversal transaction for this payment type.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode336:
		setFailoverMerchantMessage(detail, "Reversal amount does not match Authorization amount. - For a merchant initiated reversal against an American Express authorization, the reversal amount must match the authorization amount exactly.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode337:
		setFailoverMerchantMessage(detail, "Transaction did not convert to Pinless - Retry the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_RETRY
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
	case api.RCode340:
		setFailoverMerchantMessage(detail, "Invalid Amount - The transaction amount is invalid (too high or too low). For example, less than 0 for an authorization, or less than .01 for other payment types.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode341:
		setFailoverMerchantMessage(detail, "Invalid Healthcare Amounts - The amount submitted with this FSA/Healthcare transaction is invalid. The FSA amount must be greater than 0, and cannot be greater than the transaction amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode346:
		setFailoverMerchantMessage(detail, "Invalid billing descriptor prefix - The billing descriptor prefix submitted is not valid. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode347:
		setFailoverMerchantMessage(detail, "Invalid billing descriptor - The billing descriptor is not valid because you are not authorized to send transactions with custom billing fields. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode348:
		setFailoverMerchantMessage(detail, "Invalid Report Group - The Report Group specified in the transaction is invalid, because it is either not in the defined list of acceptable Report Groups or there is a mis-match between the Report Group and the defined Billing Descriptor.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode349:
		setFailoverMerchantMessage(detail, "Do Not Honor - The issuing bank has put a temporary hold on the card.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode350:
		setFailoverMerchantMessage(detail, "Generic Decline - There is an unspecified problem; contact the issuing bank for more details. Note: This code can be a hard or soft decline, depending on the method of payment, and other variables.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode351:
		setFailoverMerchantMessage(detail, "Decline - Request Positive ID - Card Present transaction that requires a picture ID match.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode352:
		setFailoverMerchantMessage(detail, "Decline CVV2/CID Fail - The CVV2/CID is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode354:
		setFailoverMerchantMessage(detail, "3-D Secure transaction not supported by merchant - You are not certified to submit 3-D Secure transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode356:
		setFailoverMerchantMessage(detail, "Invalid purchase level III, the transaction contained bad or missing data - Submitted Level III data is bad or missing.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode357:
		setFailoverMerchantMessage(detail, "Missing healthcareIIAS tag for an FSA transaction - The FSA Transactions submitted does not contain the <healtcareIIAS> data element.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode358:
		setFailoverMerchantMessage(detail, "Restricted by Vantiv due to security code mismatch. - The transaction was declined due to the security code (CVV2, CID, etc) not matching.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode360:
		setFailoverMerchantMessage(detail, "No transaction found with specified Transaction Id - There were no transactions found with the specified Transaction Id. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode361:
		setFailoverMerchantMessage(detail, "Authorization no longer available - The authorization for this transaction is no longer available. Either the authorization has already been consumed by another capture, or the authorization has expired. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode362:
		setFailoverMerchantMessage(detail, "Transaction Not Voided - Already Settled - This transaction cannot be voided; it has already been delivered. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode363:
		setFailoverMerchantMessage(detail, "Auto-void on refund - This transaction (both capture and refund) has been voided. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode364:
		setFailoverMerchantMessage(detail, "Invalid Account Number - original or NOC updated eCheck account required - The submitted account number is invalid. Confirm the original account number or check NOC for new account number.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode365:
		setFailoverMerchantMessage(detail, "Total credit amount exceeds capture amount - The amount of the credit is greater than the capture, or the amount of this credit plus other credits already referencing this capture are greater than the capture amount. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode366:
		setFailoverMerchantMessage(detail, "Exceed the threshold for sending redeposits - NACHA rules allow two redeposit attempts within 180 days of the settlement date of the initial deposit attempt. This threshold has been exceeded. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode367:
		setFailoverMerchantMessage(detail, "Deposit has not been returned for insufficient/non-sufficient funds - NACHA rules only allow redeposit attempts against deposits returned for Insufficient or Uncollected Funds. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode368:
		setFailoverMerchantMessage(detail, "Invalid check number - The check number is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode369:
		setFailoverMerchantMessage(detail, "Redeposit against invalid transaction type - The redeposit attempted against an invalid transaction type. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode370:
		setFailoverMerchantMessage(detail, "Internal System Error - Call Vantiv - There is a problem with the system. Contact eCommerceSupport@vantiv.com.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode371:
		setFailoverMerchantMessage(detail, "Original Transaction has been Processed - Future Redeposits Canceled - Do not send additional redeposit transactions, since the original transaction was processed. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode372:
		setFailoverMerchantMessage(detail, "Soft Decline - Auto Recycling In Progress - The transaction was intercepted because it is being auto recycled by the Recycling Engine.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode373:
		setFailoverMerchantMessage(detail, "Hard Decline - Auto Recycling Complete - The transaction was intercepted because auto recycling has completed with a final decline.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode375:
		setFailoverMerchantMessage(detail, "Merchant is not enabled for surcharging - The submitted transaction contained a surcharge and the merchant is not enabled for surcharging.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode376:
		setFailoverMerchantMessage(detail, "This method of payment does not support surcharging - The use of a surcharge is only allowed for Visa and MasterCard methods of payment.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode377:
		setFailoverMerchantMessage(detail, "Surcharge is not valid for debit or prepaid cards - You cannot apply a surcharge to a transaction using a debit or prepaid card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode378:
		setFailoverMerchantMessage(detail, "Surcharge cannot exceed 4% of the sale amount - The surcharge in the submitted transaction exceeded 4% maximum allowed for a surcharge.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode379:
		setFailoverMerchantMessage(detail, "Transaction declined by the processing network - The SEPA Direct Debit processing network declined the transaction for unspecified reasons. Some possible reasons are: insufficient funds, IBAN/Name disagreement, red flag on account, etc.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode380:
		setFailoverMerchantMessage(detail, "Secondary amount cannot exceed the sale amount - The secondary amount exceeded the sale amount in the submitted transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode381:
		setFailoverMerchantMessage(detail, "This method of payment does not support secondary amount - The submitted method of payment does not allow the use of Convenience Fees.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode382:
		setFailoverMerchantMessage(detail, "Secondary amount cannot be less than zero - The secondary amount must be a positive integer.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode383:
		setFailoverMerchantMessage(detail, "Partial transaction is not supported when including a secondary amount - Transactions set to allow partial authorizations cannot include a secondary amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode384:
		setFailoverMerchantMessage(detail, "Secondary amount required on partial refund when used on deposit - If the associated sale or capture transaction included a secondary amount, an associated partial refund must include a secondary amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode385:
		setFailoverMerchantMessage(detail, "Secondary amount not allowed on refund if not included on deposit - If the associated sale or capture transaction did not included a secondary amount, you cannot include a secondary amount on an associated refund.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode386:
		setFailoverMerchantMessage(detail, "Processing Network Error - Worldpay is experiencing issues communicating with the SEPA Direct Debit network. Please retry the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_RETRY
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
	case api.RCode401:
		setFailoverMerchantMessage(detail, "Invalid E-mail - The e-mail address provided is not valid. Verify that it was entered correctly.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode469:
		setFailoverMerchantMessage(detail, "Invalid Recurring Request - See Recurring Response for Details - The Recurring Request was invalid, which invalidated the transaction. The Response Code and Message in the Recurring Response contains additional information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode470:
		setFailoverMerchantMessage(detail, "Approved - Recurring Subscription Created - The recurring request was processed successfully.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode471:
		setFailoverMerchantMessage(detail, "Parent Transaction Declined - Recurring Subscription Not Created - The original payment transaction was declined, so the recurring payments have not been scheduled. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode472:
		setFailoverMerchantMessage(detail, "Invalid Plan Code - The plan specified in the recurring request was invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode473:
		setFailoverMerchantMessage(detail, "Scheduled Recurring Payment Processed - The scheduled recurring payment has been processed successfully.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode475:
		setFailoverMerchantMessage(detail, "Invalid Subscription Id - The referenced subscription Id does not exist. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode476:
		setFailoverMerchantMessage(detail, "Add On Code Already Exists - The specified Add On code already exists. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode477:
		setFailoverMerchantMessage(detail, "Duplicate Add On Codes in Requests - Multiple createAddOn requests submitted with the same Add On Code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode478:
		setFailoverMerchantMessage(detail, "No Matching Add On Code for the Subscription - The Add On code specified does not exist. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode480:
		setFailoverMerchantMessage(detail, "No Matching Discount Code for the Subscription - The Discount Code supplied in the updateDiscount or deleteDiscount transaction does not exist. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode481:
		setFailoverMerchantMessage(detail, "Duplicate Discount Codes in Request - Multiple createDiscount requests submitted with the same Discount Code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode482:
		setFailoverMerchantMessage(detail, "Invalid Start Date - The supplied Start Date is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode483:
		setFailoverMerchantMessage(detail, "Merchant Not Registered for Recurring Engine - You are not registered for the use of the Recurring Engine.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode484:
		setFailoverMerchantMessage(detail, "Insufficient data to update subscription - The transaction did not include data needed for update operation.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode485:
		setFailoverMerchantMessage(detail, "Invalid Billing Date - The submitted billing date is either before the current date or otherwise invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode486:
		setFailoverMerchantMessage(detail, "Discount Code Already Exists - The specified Discount code already exists.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode487:
		setFailoverMerchantMessage(detail, "Plan Code already exists - The specified Plan Code already exists.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode500:
		setFailoverMerchantMessage(detail, "The account number was changed - An Account Updater response indicating the Account Number changed from the original number.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode501:
		setFailoverMerchantMessage(detail, "The account was closed - An Account Updater response indicating the account was closed. Contact the cardholder directly for updated information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode502:
		setFailoverMerchantMessage(detail, "The expiration date was changed - An Account Updater response indicating the Expiration date for the card has changed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode503:
		setFailoverMerchantMessage(detail, "The issuing bank does not participate in the update program - An Account Updater response indicating the issuing bank does not participate in the update program")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode504:
		setFailoverMerchantMessage(detail, "Contact the cardholder for updated information - An Account Updater response indicating you should contact the cardholder directly for updated information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode505:
		setFailoverMerchantMessage(detail, "No match found - An Account Updater response indicating no match was found in the updated information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode506:
		setFailoverMerchantMessage(detail, "No changes found - An Account Updater response indicating there have been no changes to the account information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode521:
		setFailoverMerchantMessage(detail, "Soft Decline - Card reader decryption service is not available - The connection to the decryption service is currently unavailable. Please retry the transaction and/or contact your Relationship Manager.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode523:
		setFailoverMerchantMessage(detail, "Soft Decline - Decryption failed - Our attempt to decrypt the card information failed. Please retry the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode524:
		setFailoverMerchantMessage(detail, "Hard Decline - Input data is invalid. - The submitted data is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode530:
		setFailoverMerchantMessage(detail, "Apple Pay Key Mismatch - The submitted publicKeyHash element does not match any configured entries. Contact your Implementation Consultant.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode531:
		setFailoverMerchantMessage(detail, "Apple Pay Decryption Failed - Worldpay was unable to decrypt the submitted information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode540:
		setFailoverMerchantMessage(detail, "Hard Decline - Decryption Failed - Worldpay was unable to decrypt the submitted card number and/or CVV.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode550:
		setFailoverMerchantMessage(detail, "Restricted Device or IP - ThreatMetrix Fraud Score Below Threshold - The transaction was declined because the resulting ThreatMetrix Fraud Score was below the acceptable threshold set in the merchant’s policy.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode601:
		setFailoverMerchantMessage(detail, "Soft Decline - Primary Funding Source Failed - A PayPal response indicating the transaction failed due to an issue with primary funding source (e.g. expired Card, insufficient funds, etc.).")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode602:
		setFailoverMerchantMessage(detail, "Soft Decline - Buyer has alternate funding source - ")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode610:
		setFailoverMerchantMessage(detail, "Hard Decline - Invalid Billing Agreement Id - A PayPal response indicating the Billing Agreement ID is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode611:
		setFailoverMerchantMessage(detail, "Hard Decline - Primary Funding Source Failed - A PayPal response indicating the issuer is unavailable.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode612:
		setFailoverMerchantMessage(detail, "Hard Decline - Issue with Paypal Account - A PayPal response indicating the transaction failed due to an issue with the buyer account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode613:
		setFailoverMerchantMessage(detail, "Hard Decline - PayPal authorization ID missing - A PayPal response indicating the need to correct the authorization ID before resubmitting.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode614:
		setFailoverMerchantMessage(detail, "Hard Decline - confirmed email address is not available - A PayPal response indicating your account is configured to decline transactions without a confirmed address. request another payment method or contact eCommerceSupport@vantiv.com to modify your account settings.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode615:
		setFailoverMerchantMessage(detail, "Hard Decline - PayPal buyer account denied - A PayPal response indicating account unauthorized payment risk.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode616:
		setFailoverMerchantMessage(detail, "Hard Decline - PayPal buyer account restricted - A PayPal response indicating PayPal is unable to process the payment. Buyer should contact PayPal with questions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode617:
		setFailoverMerchantMessage(detail, "Hard Decline - PayPal order has been voided, expired, or completed - A PayPal response indicating no further authorizations/captures can be processed against this order. A new order must be created.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode618:
		setFailoverMerchantMessage(detail, "Hard Decline - issue with PayPal refund - A PayPal response indicating one of these potential refund related issues: duplicate partial refund must be less than or equal to original or remaining amount, past time limit, not allowed for transaction type, consumer account locked/inactive, or complaint exists - only a full refund of total/remaining amount allowed. Contact eCommerceSupport@vantiv.com for specific details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode619:
		setFailoverMerchantMessage(detail, "Hard Decline - PayPal credentials issue - A PayPal response indicating you do not have permissions to make this API call.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode620:
		setFailoverMerchantMessage(detail, "Hard Decline - PayPal authorization voided or expired - A PayPal response indicating you cannot capture against this authorization. You need to perform a brand new authorization for the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode621:
		setFailoverMerchantMessage(detail, "Hard Decline - required PayPal parameter missing - A PayPal response indicating missing parameters are required. Contact eCommerceSupport@vantiv.com for specific details.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode622:
		setFailoverMerchantMessage(detail, "Hard Decline - PayPal transaction ID or auth ID is invalid - A PayPal response indicating the need to check the validity of the authorization ID prior to reattempting the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode623:
		setFailoverMerchantMessage(detail, "Hard Decline - Exceeded maximum number of PayPal authorization attempts - A PayPal response indicating you should capture against a previous authorization.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode624:
		setFailoverMerchantMessage(detail, "Hard Decline - Transaction amount exceeds merchant’s PayPal account limit. - A PayPal response indicating the transaction amount exceeds the merchant’s account limit. Contact eCommerceSupport@vantiv.com to modify your account settings.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode625:
		setFailoverMerchantMessage(detail, "Hard Decline - PayPal funding sources unavailable. - A PayPal response indicating the buyer needs to add another funding sources to their account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode626:
		setFailoverMerchantMessage(detail, "Hard Decline - issue with PayPal primary funding source. - A PayPal response indicating there are issues with the buyer’s primary funding source.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode627:
		setFailoverMerchantMessage(detail, "Hard Decline - PayPal profile does not allow this transaction type. - Contact your Relationship Manager to adjust your PayPal merchant profile preferences.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode628:
		setFailoverMerchantMessage(detail, "Internal System Error with PayPal - Contact Vantiv - There is a problem with your username and password. Contact eCommerceSupport@vantiv.com.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode629:
		setFailoverMerchantMessage(detail, "Hard Decline - Contact PayPal consumer for another payment method - A PayPal response indicating you should contact the consumer for another payment method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode637:
		setFailoverMerchantMessage(detail, "Invalid terminal Id - The terminal Id submitted with the POS transaction is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode640:
		setFailoverMerchantMessage(detail, "PINless Debit processing not supported for non-recurring transactions - At this time, we support PINless Debit transaction only for recurring transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode641:
		setFailoverMerchantMessage(detail, "PINless Debit processing not supported for partial auths - PINless Debit does not support partial authorizations. You can resubmit the transaction without using the partial auth flag.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode642:
		setFailoverMerchantMessage(detail, "Merchant not configured for PINless Debit processing - You are not enabled for PINless Debit processing. Please consult your Relationship Manager for additional information about this feature.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode701:
		setFailoverMerchantMessage(detail, "Under 18 years old - A PayPal Credit response indicating the customer is under 18 years of age based upon the date of birth.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode702:
		setFailoverMerchantMessage(detail, "Bill to outside USA - A PayPal Credit response indicating the billing address is outside the United States.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode703:
		setFailoverMerchantMessage(detail, "Bill to address is not equal to ship to address - A PayPal Credit response indicating that the billing address does not match the shipping address.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode704:
		setFailoverMerchantMessage(detail, "Declined, foreign currency, must be USD - A PayPal Credit or PINless Debit response indicating the transaction is declined, because it is not in US dollars.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode705:
		setFailoverMerchantMessage(detail, "On negative file - A PayPal Credit response indicating the account is on the negative file.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode706:
		setFailoverMerchantMessage(detail, "Blocked agreement - A PayPal Credit response indicating a blocked agreement account status.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode707:
		setFailoverMerchantMessage(detail, "Insufficient buying power - A PayPal Credit response indicating that the account holder does not have sufficient credit available for the transaction amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode708:
		setFailoverMerchantMessage(detail, "Invalid Data - A PayPal Credit response indicating that there are one or more problems with the submitted data.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode709:
		setFailoverMerchantMessage(detail, "Invalid Data - data elements missing - A PayPal Credit response indicating one or more required data elements are missing. Also, returned for an Direct Debit transaction that is missing a required data element. For example, failure to include the name element in an echeckSale orecheckCredittransaction would result in this code being returned.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode710:
		setFailoverMerchantMessage(detail, "Invalid Data - data format error - A PayPal Credit response indicating that some data was formatted incorrectly.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode711:
		setFailoverMerchantMessage(detail, "Invalid Data - Invalid T&C version - A PayPal Credit response indicating the T&C version is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode712:
		setFailoverMerchantMessage(detail, "Duplicate transaction - A PayPal Credit response indicating that the transaction is a duplicate.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode713:
		setFailoverMerchantMessage(detail, "Verify billing address - A PayPal Credit response indicating that you should verify the billing address.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode714:
		setFailoverMerchantMessage(detail, "Inactive Account - A PayPal Credit response indicating the customer account is inactive.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode716:
		setFailoverMerchantMessage(detail, "Invalid Auth - A PayPal Credit response indicating that the referenced authorization is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode717:
		setFailoverMerchantMessage(detail, "Authorization already exists for the order - A PayPal Credit response indicating that an authorization already exists for the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode730:
		setFailoverMerchantMessage(detail, "Lodging transactions are not allowed for this MCC - Your current MCC does not allow lodging transactions. Please consult your Relationship Manager.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode731:
		setFailoverMerchantMessage(detail, "Duration cannot be negative - You submitted a negative value for the <duration> element. Correct the error and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode732:
		setFailoverMerchantMessage(detail, "Hotel Folio Number cannot be blank - Although the schema does not require the submission of the <hotelFolioNumber> element, if you do include it, you must specify a value (i.e., null not allowed). Please either add a valid value or remove the element and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode733:
		setFailoverMerchantMessage(detail, "Invalid check in date - There is a problem with the submitted check in date (for example, 2018-02-32). Please correct the date and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode734:
		setFailoverMerchantMessage(detail, "Invalid check out date - There is a problem with the submitted check out date (for example, 2018-02-32). Please correct the date and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode735:
		setFailoverMerchantMessage(detail, "Invalid check in or check out date - There is a problem with the submitted check in or check out date (for example, 2018-02-32). Please correct the date(s) and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode736:
		setFailoverMerchantMessage(detail, "Check out date cannot be before check in date - The check out date you submitted was before the check in date. Correct the error and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode737:
		setFailoverMerchantMessage(detail, "Number of adults cannot be negative - You submitted a negative value for the <numAdult> element. Correct the error and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode738:
		setFailoverMerchantMessage(detail, "Room rate cannot be negative - You submitted a negative value for the <roomRate> element. Correct the error and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode739:
		setFailoverMerchantMessage(detail, "Room tax cannot be negative - You submitted a negative value for the <roomTax> element. Correct the error and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode801:
		setFailoverMerchantMessage(detail, "Account number was successfully registered - The card number was successfully registered and a token number was returned.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode802:
		setFailoverMerchantMessage(detail, "Account number was previously registered - The card number was previously registered for tokenization. Note: You also receive this response code when using a low value token in a transaction, because the system registers the PAN at the time it creates the LVT.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode803:
		setFailoverMerchantMessage(detail, "Valid Token - The token is valid.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode805:
		setFailoverMerchantMessage(detail, "Card Validation Number Updated - The stored value for CVV2/CVC2/CID has been successfully updated.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode820:
		setFailoverMerchantMessage(detail, "Credit card number was invalid - The card number submitted for tokenization is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode821:
		setFailoverMerchantMessage(detail, "Merchant is not authorized for tokens - Your organization is not authorized to use tokens.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode822:
		setFailoverMerchantMessage(detail, "Token was not found - The token number submitted with this transaction was not found.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode823:
		setFailoverMerchantMessage(detail, "Token Invalid - The submitted token is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode825:
		setFailoverMerchantMessage(detail, "Merchant not authorized for eCheck tokens - Your organization is not authorized for Direct Debit tokenization.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode826:
		setFailoverMerchantMessage(detail, "Checkout Id was invalid - The submitted checkoutId is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode827:
		setFailoverMerchantMessage(detail, "Checkout Id was not found - The submitted checkoutId was not found. The low value token is only good for 24 hours and may have expired.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode828:
		setFailoverMerchantMessage(detail, "Generic Checkout Id error - An unknown error caused the use of checkoutId to fail.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode835:
		setFailoverMerchantMessage(detail, "Capture amount can not be more than authorized amount - The amount in the submitted Capture exceeds 115% of the authorized amount. Appears in Declined Transaction report.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode850:
		setFailoverMerchantMessage(detail, "Tax Billing only allowed for MCC 9311 - Tax Billing elements are allowed only for MCC 9311.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode851:
		setFailoverMerchantMessage(detail, "MCC 9311 requires taxType element - Missing taxType element")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode852:
		setFailoverMerchantMessage(detail, "Debt Repayment only allowed for VI transactions on MCCs 6012 and 6051 - You must be either MCC 6012 or 6051 to designate a Visa transaction as Debt Repayment (debtRepayment element set to true).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode861:
		setFailoverMerchantMessage(detail, "Routing Number did not match one on file for token - The routing number submitted does not match the number submitted when the token was created. Verify the routing number and resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode877:
		setFailoverMerchantMessage(detail, "Invalid Pay Page Registration Id - An eProtect response indicating that the Registration ID submitted is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode878:
		setFailoverMerchantMessage(detail, "Expired Pay Page Registration Id - An eProtect response indicating that the Registration ID has expired (Registration IDs expire 24 hours after being issued).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode879:
		setFailoverMerchantMessage(detail, "Merchant is not authorized for Pay Page - Your organization is not authorized to use eProtect.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode898:
		setFailoverMerchantMessage(detail, "Generic token registration error - There is an unspecified token registration error; contact your Relationship Manager")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode899:
		setFailoverMerchantMessage(detail, "Generic token use error - There is an unspecified token use error; contact your Relationship Manager.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode900:
		setFailoverMerchantMessage(detail, "Invalid Bank Routing Number - The Direct Debit routing number submitted with this transaction has failed validation.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode901:
		setFailoverMerchantMessage(detail, "Missing Name - The customer name is required for SEPA transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode902:
		setFailoverMerchantMessage(detail, "Invalid Name - The customer name must be a minimum of two characters for SEPA transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode903:
		setFailoverMerchantMessage(detail, "Missing Billing Country Code - The Billing Country code is required for SEPA transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode904:
		setFailoverMerchantMessage(detail, "Invalid IBAN - The submitted International Bank Account number is invalid. Please correct the number and resubmit the traction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode905:
		setFailoverMerchantMessage(detail, "Missing Email Address - The customer email address is required for SEPA transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode906:
		setFailoverMerchantMessage(detail, "Missing mandate reference - You must provide a Mandate reference for standard and BYOM recurring SEPA deposit transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode907:
		setFailoverMerchantMessage(detail, "Invalid mandate reference - The Mandate reference is invalid. It must conform to the following format: 1 to 35 characters consisting of alphanumeric, colon, question mark, forward slash, plus,parenthesis, comas, period, space, and dash. The applicable regular expression is: ^[A-Za-z0-9:?/+(),. -]{1,35}$")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode908:
		setFailoverMerchantMessage(detail, "Missing mandate URL - You must provide a Mandate URL for SEPA Bring Your Own Mandate deposit transactions (both one-time and recurring).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode909:
		setFailoverMerchantMessage(detail, "Invalid mandate URL - The Mandate URL must start with https and be followed by 5 to 120 characters adhering to the following regular expression: ^https://.{5,120}$")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode911:
		setFailoverMerchantMessage(detail, "Missing mandate signature date - You must provide a Mandate signature date for SEPA Bring Your Own Mandate deposit transactions (both one-time and recurring).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode912:
		setFailoverMerchantMessage(detail, "Invalid mandate signature date - You must provide a Mandate signature date earlier than or the same as the current date with the following format: YYYY-MM-DD.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode913:
		setFailoverMerchantMessage(detail, "Recurring mandate already exists - Worldpay returns this message when you submit multiple first Bring Your Own Mandate recurring transactions with the same mandate reference. The mandate references among the recurring transactions for a single merchant must be unique.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode914:
		setFailoverMerchantMessage(detail, "Recurring mandate was not found - Worldpay returns this message when you submit a subsequent or final standard recurring transaction before a first standard recurring is processed and confirmed.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode915:
		setFailoverMerchantMessage(detail, "Final recurring was already received using this mandate - Worldpay returns this message when you submit a first or subsequent recurring transaction after we received a final recurring. The life cycle of a recurring mandate ends after the processing of a final recurring deposit transaction. This applies to both standard and Bring Your Own Mandate recurring transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode916:
		setFailoverMerchantMessage(detail, "IBAN did not match one on file for mandate - Worldpay returns this message when you submit a subsequent or final recurring with a different IBAN than the IBAN used in the first recurring transaction. This applies toboth standard and Bring Your Own Mandate recurring transactions. Note: If a customer wants to use a different IBAN for a recurring SEPA mandate, they must use a request for another mandate reference using this new IBAN.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CHARGE
	case api.RCode917:
		setFailoverMerchantMessage(detail, "Invalid Billing Country - Some of the alternative payment methods restrict the allowed consumer's billing country codes. For example, iDEAL only allows country code NL.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode940:
		setFailoverMerchantMessage(detail, "This Funding Instruction results in a negative account balance - There are insufficient funds in the FBO Settlement Account to cover the Funding Instruction. Wait for additional funds to settle to the account and then resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode941:
		setFailoverMerchantMessage(detail, "Account balance information unavailable at this time. - Typically, this response occurs only for new FBO Settlement accounts that do not yet have any settled transactions and were not pre-funded (i.e., no balance yet recorded for the account). Wait for additional funds to settle to the account or pre-fund the account and then resubmit the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
	case api.RCode942:
		setFailoverMerchantMessage(detail, "The submitted card is not eligible for Fast Access Funding. - The card you submitted in the Fast Access Funding instruction cannot receive funds using this method.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode943:
		setFailoverMerchantMessage(detail, "Transaction cannot use both ccdPaymentInformation and ctxPaymentInformation - A transaction can not contain both the ccdPaymentInformation and ctxPaymentInformation elements.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode944:
		setFailoverMerchantMessage(detail, "Processing Error - Please retry the transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_RETRY
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode950:
		setFailoverMerchantMessage(detail, "Decline - Negative Information on File - An Direct Debit response indicating the account is on the negative file.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode951:
		setFailoverMerchantMessage(detail, "Absolute Decline - An Direct Debit response indicating that this transaction was declined.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode952:
		setFailoverMerchantMessage(detail, "The Merchant Profile does not allow the requested operation - An Direct Debit response indicating that your Merchant Profile does not allow the requested operation. Contact your Relationship Manager for additional information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode953:
		setFailoverMerchantMessage(detail, "The account cannot accept ACH transactions - An Direct Debit response indicating the customer’s checking account does not accept ACH transactions.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode954:
		setFailoverMerchantMessage(detail, "The account cannot accept ACH transactions or site drafts - An Direct Debit response indicating the customer’s checking account does not accept ACH transactions or site drafts.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode955:
		setFailoverMerchantMessage(detail, "Amount greater than limit specified in the Merchant Profile - A Direct Debit response indicating that the dollar amount of this transaction exceeds the maximum amount specified in your Merchant Profile. Contact your Relationship Manager for additional information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode956:
		setFailoverMerchantMessage(detail, "Merchant is not authorized to perform eCheck Verification transactions - A Direct Debit response indicating that your organization is not authorized to perform eCheck verifications. Contact your Relationship Manager for additional information.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode957:
		setFailoverMerchantMessage(detail, "First Name and Last Name required for eCheck Verifications - A Direct Debit response indicating that the first and last name of the customer is required for eCheck verifications.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode958:
		setFailoverMerchantMessage(detail, "Company Name required for corporate account for eCheck Verifications - A Direct Debit response indicating that the company name is required for verifications on corporate accounts.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode959:
		setFailoverMerchantMessage(detail, "Phone number required for eCheck Verifications - A Direct Debit response indicating that the phone number of the customer is required for echeck verifications.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode961:
		setFailoverMerchantMessage(detail, "Card Brand token not supported - This code is returned if the merchant submits a Visa generated token.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode962:
		setFailoverMerchantMessage(detail, "Private Label Card not supported - This code is returned if the transaction involves a Visa Private Label card.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	default:
		setFailoverMerchantMessage(detail, "Unknown error response")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
		logger.Instance().Error("Unknown error code from worldpay connector", zap.String("Error code", string(code)))
	}
}
