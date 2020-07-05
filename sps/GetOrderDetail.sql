CREATE DEFINER=`root`@`localhost` PROCEDURE `GetOrderDetail`(
	in ip_OrderId int
)
BEGIN
	select 
		od.OrderID
		, od.ProductID
		, p.Name as ProductName
		, p.Alias as ProductAlias
		, od.Price
		, od.Quantity        
	from ustoradb.orderdetails as od
	left join ustoradb.products as p on p.ID = od.ProductID
	where od.OrderID = ip_OrderId;
END