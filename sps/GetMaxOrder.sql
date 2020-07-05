CREATE DEFINER=`root`@`localhost` PROCEDURE `GetMaxOrder`()
BEGIN
	select max(ID)
	from orders;
END