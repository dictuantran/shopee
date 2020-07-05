CREATE DEFINER=`root`@`localhost` PROCEDURE `GetRevenuesStatistic`(
	in fromDate varchar(255)
    , in toDate varchar(255))
BEGIN
select o.CreatedDate as Date,
    SUM(od.Price* od.Quantity) as Revenues,
    SUM((od.Price* od.Quantity)-(p.OriginalPrice*od.Quantity)) as Benefit
    from Orders o 
    inner join OrderDetails od
    on o.ID = od.OrderID
    inner join Products p
    on od.ProductID = p.ID
    inner join TrackOrders tro on tro.OrderId = o.ID
    where cast(o.CreatedDate as date) between cast(fromDate as date) and cast(toDate as date)
    and tro.Status = 0
    group by o.CreatedDate;
END