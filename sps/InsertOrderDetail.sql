CREATE DEFINER=`root`@`localhost` PROCEDURE `InsertOrderDetail`(
		in ip_OrderID		int(11)
    ,	in ip_ProductID		bigint(20)
    ,	in ip_Price			decimal(18,2)
    ,	in ip_Quantity		int(11)
)
BEGIN
	SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;  
    insert into orderdetails(OrderID
		, ProductID
        , Price
        , Quantity)
	values(ip_OrderID
		, ip_ProductID
        , ip_Price
        , ip_Quantity);
    
    SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;
END