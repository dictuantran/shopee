CREATE DEFINER=`root`@`localhost` PROCEDURE `GetShoppingCarts`(
	in userId varchar(255)
)
BEGIN
	select p.Image,p.Name,odt.Quantity,p.ID as ProductId,p.Alias,odt.Price,od.PaymentStatus 
    from OrderDetails odt
    inner join Products p on p.ID = odt.ProductID
    inner join Orders od on od.ID = odt.OrderID
    inner join ApplicationUsers u on u.Id = od.CustomerId
    where (od.PaymentStatus = 0 or od.PaymentStatus = 1)
    and od.CustomerId = userId;
END