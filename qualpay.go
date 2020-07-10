package connector_errors

import "github.com/chargehive/proto/golang/chargehive/chtype"

func(){

	switch code {
	case api.RCode000:
		setFailoverMerchantMessage(detail, "Success - Approved and completed successfully, For authorization and sale messages this indicates that the transaction was approved by the card authorization system.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode001:
		setFailoverMerchantMessage(detail, "Refer to issuer - Refer to card issuer")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode002:
		setFailoverMerchantMessage(detail, "Refer to issuer - Refer to card issuer, special condition")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode003:
		setFailoverMerchantMessage(detail, "Invalid merchant - Invalid merchant")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode004:
		setFailoverMerchantMessage(detail, "Success - Approved and completed successfully, For authorization and sale messages this indicates that the transaction was approved by the card authorization system.")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode005:
		setFailoverMerchantMessage(detail, "Decline - Do not honor")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode006:
		setFailoverMerchantMessage(detail, "Error - Error")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode007:
		setFailoverMerchantMessage(detail, "Pick up card fraud - Pick up card, special condition (fraud account)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode008:
		setFailoverMerchantMessage(detail, "Honor with ID - Honor with ID")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode010:
		setFailoverMerchantMessage(detail, "Partial Approval - Partial approval")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode011:
		setFailoverMerchantMessage(detail, "Approved VIP - Approved (V.I.P)")
		detail.FailureType = chtype.FAILURE_TYPE_NONE
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode012:
		setFailoverMerchantMessage(detail, "Invalid transaction - Invalid transaction")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode013:
		setFailoverMerchantMessage(detail, "Invalid amount - Invalid amount or currency conversion field overflow")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode014:
		setFailoverMerchantMessage(detail, "Bad card number - Invalid account number (no such number)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode015:
		setFailoverMerchantMessage(detail, "No such issuer - No such issuer")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode019:
		setFailoverMerchantMessage(detail, "Re-enter - Re-enter transaction")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode021:
		setFailoverMerchantMessage(detail, "No action taken - No action taken")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode025:
		setFailoverMerchantMessage(detail, "Unable to locate - Unable to locate record in file")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode028:
		setFailoverMerchantMessage(detail, "File temporarily unavailable - File temporarily not available for update or inquiry")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode030:
		setFailoverMerchantMessage(detail, "Format error - Format Error - Decline (MasterCard, Discover and PayPal)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode032:
		setFailoverMerchantMessage(detail, "Partial reversal - Valid for MasterCard Reversal Requests Only - Used in a reversal message to indicate that the reversal request is for an amount that is less than the original transaction.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode034:
		setFailoverMerchantMessage(detail, "Fraud reversal - Used for MasterCard reversal requests only, Suspect Fraud, indicating an approved e-Commerce transaction is cancelled by the merchant")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode039:
		setFailoverMerchantMessage(detail, "No credit account - No credit account")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode041:
		setFailoverMerchantMessage(detail, "Pick up - Lost - Lost card, pick up (fraud account)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode043:
		setFailoverMerchantMessage(detail, "Pick up - Stolen - Stolen card, pick up (fraud account)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode051:
		setFailoverMerchantMessage(detail, "Insufficient funds - Not sufficient funds")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode052:
		setFailoverMerchantMessage(detail, "No checking account - No checking account")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode053:
		setFailoverMerchantMessage(detail, "No savings account - No savings account")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode054:
		setFailoverMerchantMessage(detail, "Expired card - Expired card or expiration date is missing")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode055:
		setFailoverMerchantMessage(detail, "Invalid PIN - Incorrect PIN or PIN missing")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode057:
		setFailoverMerchantMessage(detail, "Not permitted - Transaction not permitted to cardholder")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode058:
		setFailoverMerchantMessage(detail, "Not allowed at terminal - Transaction not allowed at terminal")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode059:
		setFailoverMerchantMessage(detail, "Suspected fraud - Suspected fraud")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_FRAUD
	case api.RCode061:
		setFailoverMerchantMessage(detail, "Exceeds amount limit - Exceeds approval amount limit")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode062:
		setFailoverMerchantMessage(detail, "Restricted card - Restricted card (card invalid in this region or country)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode063:
		setFailoverMerchantMessage(detail, "Security violation - Security violation (source is not correct issuer)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode064:
		setFailoverMerchantMessage(detail, "AML requirement not fulfilled - Transaction does not fulfill AML requirement")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode065:
		setFailoverMerchantMessage(detail, "Activity limit exceeded - Exceeds withdrawal frequency limit")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode068:
		setFailoverMerchantMessage(detail, "Reversal - Used in MasterCard Reversal Requests and Discover and PayPal Responses")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode075:
		setFailoverMerchantMessage(detail, "PIN entry attempts exceeded - Allowable number of PIN entry tries exceeded")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode076:
		setFailoverMerchantMessage(detail, "RRN not found - Reversal: Unable to locate previous message (no match on Retrieval Reference number)")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode077:
		setFailoverMerchantMessage(detail, "Invalid reversal data - Previous message located for a repeat or reversal, but repeat or reversal data are inconsistent with original message")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode078:
		setFailoverMerchantMessage(detail, "Decline - Invalid/nonexistent account - Decline (MasterCard specific)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode079:
		setFailoverMerchantMessage(detail, "Already reversed - Already reversed (by Switch)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode080:
		setFailoverMerchantMessage(detail, "No financial impact - No financial impact (Reversal for declined debit)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode081:
		setFailoverMerchantMessage(detail, "PIN cryptographic error - Cryptographic error found in PIN")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode082:
		setFailoverMerchantMessage(detail, "Incorrect CVV - Negative CAM, dCVV, iCVV, or CVV results")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode083:
		setFailoverMerchantMessage(detail, "Unable to verify PIN - Unable to verify PIN")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode084:
		setFailoverMerchantMessage(detail, "Decline - Invalid Authorization Life Cycle - Decline (MasterCard) Duplicate Transaction Detected (Visa)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode085:
		setFailoverMerchantMessage(detail, "No reason to decline - No reason to decline a request for address verification, CVV2 verification, or a credit voucher or merchandise return")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode086:
		setFailoverMerchantMessage(detail, "Cannot verify PIN - Cannot verify PIN; for example, no PVV")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode089:
		setFailoverMerchantMessage(detail, "Ineligible - Ineligible to receive financial position information (GIV)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode091:
		setFailoverMerchantMessage(detail, "Issuer unavailable - Issuer or switch inoperative and STIP not applicable or not available for this transaction; Time-out when no stand-in; POS Check Service: Destination unavailable; Credit Voucher and Merchandise Return Authorizations: V.I.P. sent the transaction to the issuer, but the issuer was unavailable.")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
	case api.RCode092:
		setFailoverMerchantMessage(detail, "Destination not found - Financial institution or intermediate network facility cannot be found for routing (receiving institution ID is invalid)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode093:
		setFailoverMerchantMessage(detail, "Transaction cannot complete - Transaction cannot be completed - violation of law")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode094:
		setFailoverMerchantMessage(detail, "Duplicate transmission - Duplicate Transmission Detected (Integrated Debit and MasterCard)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode096:
		setFailoverMerchantMessage(detail, "System malfunction - System malfunction")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0B1:
		setFailoverMerchantMessage(detail, "Surcharge not permitted - Surcharge amount not permitted on Visa cards or EBT food stamps (U.S. acquirers only)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode0B2:
		setFailoverMerchantMessage(detail, "Surcharge not supported - Surcharge amount not supported by debit network issuer")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode0N0:
		setFailoverMerchantMessage(detail, "Force STIP - Force STIP")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode0N3:
		setFailoverMerchantMessage(detail, "Not available - Cash service not available")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode0N4:
		setFailoverMerchantMessage(detail, "Exceeds issuer limit - Cash request exceeds issuer or approved limit")
		detail.FailureType = chtype.FAILURE_TYPE_SOFT
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode0N5:
		setFailoverMerchantMessage(detail, "Ineligible for resubmission - Ineligible for resubmission")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0N7:
		setFailoverMerchantMessage(detail, "Decline CVV2 failure - Decline for CVV2 failure")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode0N8:
		setFailoverMerchantMessage(detail, "Preauthorized amount exceeded - Transaction amount exceeds preauthorized approval amount")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode0Q1:
		setFailoverMerchantMessage(detail, "Card authentication failed - Card Authentication failed")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode0R0:
		setFailoverMerchantMessage(detail, "Stop payment order - The transaction was declined or returned because the cardholder requested that payment of a specific recurring or installment payment transaction be stopped.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode0R1:
		setFailoverMerchantMessage(detail, "Revocation of authorization - The transaction was declined or returned because the cardholder has requested that payment of all recurring or installment payment transactions for a specific merchant account be stopped.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0R2:
		setFailoverMerchantMessage(detail, "Transaction not qualified - The transaction does not qualify for Visa PIN")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode0R3:
		setFailoverMerchantMessage(detail, "Revocation of authorization - Revocation of All Authorizations Order")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode0T0:
		setFailoverMerchantMessage(detail, "First time check - First time check")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_VERIFICATION
	case api.RCode0T1:
		setFailoverMerchantMessage(detail, "Check valid - not converted - Check is OK but cannot be converted")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0T2:
		setFailoverMerchantMessage(detail, "Invalid transit routing - Invalid Routing Transit Number or check belongs to a category that is not eligible for conversion; Transaction failed ABA check digit validation.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode0T3:
		setFailoverMerchantMessage(detail, "Amount exceeds service limit - Amount greater than established service limit")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode0T4:
		setFailoverMerchantMessage(detail, "Unpaid items - Unpaid items, failed negative file check")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0T5:
		setFailoverMerchantMessage(detail, "Duplicate check number - Duplicate check number")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode0T6:
		setFailoverMerchantMessage(detail, "MICR error - MICR error")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0T7:
		setFailoverMerchantMessage(detail, "Check limit exceeded - Too many checks (over merchant or bank limit)")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0XA:
		setFailoverMerchantMessage(detail, "Forward to issuer - Forward to issuer")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0XD:
		setFailoverMerchantMessage(detail, "Forward to issuer - Forward to issuer")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0Y1:
		setFailoverMerchantMessage(detail, "Offline approved - Offline-approved")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0Y3:
		setFailoverMerchantMessage(detail, "Offline approved - Unable to go online; offline-approved")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0Z1:
		setFailoverMerchantMessage(detail, "Offline declined - Offline-declined")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode0Z3:
		setFailoverMerchantMessage(detail, "Offline-declined - Unable to go online; offline-declined")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode100:
		setFailoverMerchantMessage(detail, "Bad request - The message was invalid. The field rmsg will contain detail about the invalid or missing field(s).")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode101:
		setFailoverMerchantMessage(detail, "Invalid credentials - The merchant_id and security_key provided do not match an account with Qualpay.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode102:
		setFailoverMerchantMessage(detail, "Invalid pg_id - The pg_id value could not be linked to a valid transaction. Used for Payment Gateway Refunds.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode103:
		setFailoverMerchantMessage(detail, "Missing cardholder data - The request was missing valid cardholder data. Requests requiring cardholder data need only one of the following combinations: [card_number (only allowed for force credit and tokenize messages)], [card_number, exp_date], [card_id], [card_id, exp_date], [card_swipe], [Customer_id]")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode104:
		setFailoverMerchantMessage(detail, "Invalid transaction amount - The request was either missing the amt_tran or the value provided was invalid. Verify requests require the amt_tran to be zero or not present in the request message. Other messages require the amt_tran field to be numeric and greater than zero. Negative amounts are not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode105:
		setFailoverMerchantMessage(detail, "Missing auth_code - Force transactions require the field auth_code in the request message.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode106:
		setFailoverMerchantMessage(detail, "Invalid AVS (Address Verification Service) data - If the field avs_address is provided in the request message, then the message must also contain the _avs_zip_ field.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode107:
		setFailoverMerchantMessage(detail, "Invalid expiration date - The exp_date provided in the request was not properly formatted.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode108:
		setFailoverMerchantMessage(detail, "Invalid card number - The card_number field in the request message was non-numeric or contained either too few or too many digits.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode109:
		setFailoverMerchantMessage(detail, "Field length validation failed - Returned when any field exceeds the maximum allowed length. The rmsg field will contain the name of the first field that failed validation.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode110:
		setFailoverMerchantMessage(detail, "Dynamic DBA not allowed - Returned when the request message contained any of the dynamic DBA fields and the account is not configured for dynamic DBA.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode111:
		setFailoverMerchantMessage(detail, "Credits not allowed - Returned when an unreferenced credit to a previous transaction is submitted and the merchant is not currently authorized to process credits. For security reasons, we recommend performing a Refund rather than a Credit. Please email support@qualpay.com to discuss enabling Credits on your account.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode112:
		setFailoverMerchantMessage(detail, "Invalid customer data - Returned when the message requested a customer to be created in the Qualpay Customer Vault, but the customer_id already exists or required customer fields are not included.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode113:
		setFailoverMerchantMessage(detail, "Declined due to merchant risk settings - Declined because of AVS and/or CVV2 result")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONFIGURATION
	case api.RCode114:
		setFailoverMerchantMessage(detail, "Invalid currency code - The tran_currency provided in the request should be the ISO numeric currency code.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode115:
		setFailoverMerchantMessage(detail, "Invalid surcharge data - Returned when surcharge is not allowed for the merchant or on the card or if the surcharge amount exceeds 4% of the transaction amount.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode116:
		setFailoverMerchantMessage(detail, "Invalid email address - Returned when the email_address included in the request is invalid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode117:
		setFailoverMerchantMessage(detail, "Email address is required - The request is missing email_address.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_PERSON
	case api.RCode118:
		setFailoverMerchantMessage(detail, "Invalid merchant logo URL - The logo_url provided in the request is not valid.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode119:
		setFailoverMerchantMessage(detail, "Invalid ACH payment - Invalid ACH payment")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode401:
		setFailoverMerchantMessage(detail, "Void failed - Returned if the transaction has already been captured or voided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode402:
		setFailoverMerchantMessage(detail, "Refund failed - Returned if the transaction has already been refunded, the original transaction has not been captured, the total amount of all refunds exceeds the original transaction amount or the original transaction was not a sale.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode403:
		setFailoverMerchantMessage(detail, "Capture failed - Returned if the amount exceeds the authorized amount (except when the merchant category code allows tips), the transaction has already been captured or the authorization has been voided.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_METHOD
	case api.RCode404:
		setFailoverMerchantMessage(detail, "Batch close failed. - Returned if the Batch Close transaction fails.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	case api.RCode405:
		setFailoverMerchantMessage(detail, "Tokenization failed - Returned in the Tokenize transaction request fails.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_REQUEST
	case api.RCode998:
		setFailoverMerchantMessage(detail, "Timeout - Returned if the authorization request timed out without returning a response. Timeouts occur when the authorization system does not receive a response from the host within 10 seconds.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_CONNECTIVITY
	case api.RCode999:
		setFailoverMerchantMessage(detail, "Internal error - This rcode is returned when the payment gateway application encountered an unexpected error while processing the request.")
		detail.FailureType = chtype.FAILURE_TYPE_HARD
		detail.Category = chtype.RESPONSE_CATEGORY_UNKNOWN
	}
}
