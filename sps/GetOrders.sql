CREATE DEFINER=`root`@`localhost` PROCEDURE `GetOrders`(
	in ip_FromDate datetime,
    in ip_ToDate datetime
)
BEGIN
	SELECT o.ID as OrderId
	, o.CreatedDate as OrderDate
    , o.CreatedBy
    , o.PaymentMethod
    , o.Status as OrderStatus
    , o.PaymentStatus
    , sum(od.Price) as Price
FROM ustoradb.orders as o
LEFT JOIN ustoradb.orderdetails as od on o.ID = od.OrderID
WHERE CreatedDate between ip_FromDate and ip_ToDate
group by o.ID;
END