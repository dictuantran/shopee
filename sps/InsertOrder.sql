CREATE DEFINER=`root`@`localhost` PROCEDURE `InsertOrder`(
		in ip_OrderID			int(11)
	,	in ip_CustomerName 		varchar(100)
    ,	in ip_CustomerAddress	varchar(250)
	,	in ip_CustomerEmail		varchar(100)
    ,	in ip_CustomerMobile	varchar(50)
    ,	in ip_CustomerMessage	longtext
    ,	in ip_CreatedDate		datetime(6)
    ,	in ip_CreatedBy			varchar(250)
    ,	in ip_PaymentMethod		varchar(250)
    ,	in ip_PaymentStatus		int(11)
    ,	in ip_Status			tinyint(1)
    , 	in CustomerId			varchar(128)
)
BEGIN
	SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
    insert into orders(ID
		, CustomerName
		, CustomerAddress
		, CustomerEmail
		, CustomerMobile
		, CustomerMessage
		, CreatedDate
		, CreatedBy
		, PaymentMethod
		, PaymentStatus
		, Status
		, CustomerId)
    values(ip_OrderID
		, ip_CustomerName
		, ip_CustomerAddress
		, ip_CustomerEmail
		, ip_CustomerMobile
		, ip_CustomerMessage
		, CURRENT_TIME
		, ip_CreatedBy
		, ip_PaymentMethod
		, ip_PaymentStatus
		, ip_Status
		, CustomerId);
    
    SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;
END