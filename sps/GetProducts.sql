CREATE DEFINER=`root`@`localhost` PROCEDURE `GetProducts`(
	in categoryId int(11)
)
BEGIN
	select p.ID, p.Name, p.Alias, p.CategoryID
    from Products as p
    where CategoryID = categoryId;
END