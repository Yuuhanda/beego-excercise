-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 30, 2024 at 02:48 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `stok_barang_bfg`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `GenerateRepairLogSamples` ()   BEGIN
    DECLARE month_counter INT DEFAULT 1;
    DECLARE day_counter INT;
    DECLARE unit_id INT;
    DECLARE start_time DATETIME;
    DECLARE close_time DATETIME;

    WHILE month_counter <= 12 DO
        SET day_counter = 1;

        WHILE day_counter <= 4 DO
            SET unit_id = FLOOR(1 + RAND() * 500); -- Random id_unit between 1 and 500
            SET start_time = CONCAT('2024-', LPAD(month_counter, 2, '0'), '-', LPAD(FLOOR(1 + RAND() * 28), 2, '0'), ' ', LPAD(FLOOR(RAND() * 24), 2, '0'), ':', LPAD(FLOOR(RAND() * 60), 2, '0'), ':00');
            SET close_time = DATE_ADD(start_time, INTERVAL FLOOR(RAND() * 5) DAY); -- Random close date 0-5 days after start

            -- Insert "repair started" log
            INSERT INTO `repair_log` (`id_unit`, `comment`, `rep_type`, `datetime`)
            VALUES (unit_id, CONCAT('Repair started for unit ', unit_id), 1, start_time);

            -- Insert "repair done" log
            INSERT INTO `repair_log` (`id_unit`, `comment`, `rep_type`, `datetime`)
            VALUES (unit_id, CONCAT('Repair completed for unit ', unit_id), 2, close_time);

            SET day_counter = day_counter + 1;
        END WHILE;

        SET month_counter = month_counter + 1;
    END WHILE;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `api_route`
--

CREATE TABLE `api_route` (
  `id` int(11) NOT NULL,
  `path` varchar(255) NOT NULL,
  `method` varchar(10) NOT NULL,
  `controller` varchar(100) NOT NULL,
  `action` varchar(100) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `api_route`
--

INSERT INTO `api_route` (`id`, `path`, `method`, `controller`, `action`, `description`, `created_at`, `updated_at`) VALUES
(290, '/auth/login', 'POST', 'AuthController', 'Login', 'API endpoint for POST /auth/login', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(291, '/api/routes/scan', 'POST', 'APIRouteController', 'ScanRoutes', 'API endpoint for POST /api/routes/scan', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(292, '/api/routes/list', 'GET', 'APIRouteController', 'ListRoutes', 'API endpoint for GET /api/routes/list', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(293, '/api/routes/:id', 'GET', 'APIRouteController', 'Get', 'API endpoint for GET /api/routes/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(294, '/api/routes/:id', 'DELETE', 'APIRouteController', 'DeleteRoute', 'API endpoint for DELETE /api/routes/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(295, '/api/roles', 'POST', 'AuthRolesController', 'Create', 'API endpoint for POST /api/roles', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(296, '/api/roles/:id', 'GET', 'AuthRolesController', 'Get', 'API endpoint for GET /api/roles/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(297, '/api/roles', 'GET', 'AuthRolesController', 'List', 'API endpoint for GET /api/roles', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(298, '/api/roles/:id', 'PUT', 'AuthRolesController', 'Update', 'API endpoint for PUT /api/roles/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(299, '/api/roles/:id', 'DELETE', 'AuthRolesController', 'Delete', 'API endpoint for DELETE /api/roles/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(300, '/api/user-roles', 'POST', 'AuthRolesUserController', 'Create', 'API endpoint for POST /api/user-roles', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(301, '/api/user-roles/user/:userId', 'GET', 'AuthRolesUserController', 'GetUserRoles', 'API endpoint for GET /api/user-roles/user/:userId', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(302, '/api/user-roles/role/:roleId', 'GET', 'AuthRolesUserController', 'GetRoleUsers', 'API endpoint for GET /api/user-roles/role/:roleId', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(303, '/api/user-roles/:userId/:roleId', 'DELETE', 'AuthRolesUserController', 'Delete', 'API endpoint for DELETE /api/user-roles/:userId/:roleId', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(304, '/api/auth-items', 'POST', 'AuthItemController', 'Create', 'API endpoint for POST /api/auth-items', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(305, '/api/auth-items/:id', 'GET', 'AuthItemController', 'Get', 'API endpoint for GET /api/auth-items/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(306, '/api/auth-items', 'GET', 'AuthItemController', 'List', 'API endpoint for GET /api/auth-items', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(307, '/api/auth-items/:id', 'PUT', 'AuthItemController', 'Update', 'API endpoint for PUT /api/auth-items/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(308, '/api/auth-items/:id', 'DELETE', 'AuthItemController', 'Delete', 'API endpoint for DELETE /api/auth-items/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(309, '/api/auth-items/bulk', 'POST', 'AuthItemController', 'CreateBulk', 'API endpoint for POST /api/auth-items/bulk', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(310, '/user', 'POST', 'UserController', 'CreateUser', 'API endpoint for POST /user', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(311, '/auth/logout', 'POST', 'AuthController', 'Logout', 'API endpoint for POST /auth/logout', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(312, '/user/:id', 'GET', 'UserController', 'GetUser', 'API endpoint for GET /user/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(313, '/user/:id', 'PUT', 'UserController', 'UpdateUser', 'API endpoint for PUT /user/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(314, '/user/:id', 'DELETE', 'UserController', 'DeleteUser', 'API endpoint for DELETE /user/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(315, '/users', 'GET', 'UserController', 'ListUsers', 'API endpoint for GET /users', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(316, '/user/:id/visits', 'GET', 'UserVisitLogController', 'GetUserVisits', 'API endpoint for GET /user/:id/visits', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(317, '/api/item-units', 'POST', 'ItemUnitController', 'Create', 'API endpoint for POST /api/item-units', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(318, '/api/item-units/:id', 'GET', 'ItemUnitController', 'Get', 'API endpoint for GET /api/item-units/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(319, '/api/item-units', 'GET', 'ItemUnitController', 'List', 'API endpoint for GET /api/item-units', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(320, '/api/item-units/:id', 'PUT', 'ItemUnitController', 'Update', 'API endpoint for PUT /api/item-units/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(321, '/api/item-units/:id', 'DELETE', 'ItemUnitController', 'Delete', 'API endpoint for DELETE /api/item-units/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(322, '/api/item-units/serial/:serialNumber', 'GET', 'ItemUnitController', 'GetBySerialNumber', 'API endpoint for GET /api/item-units/serial/:serialNumber', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(323, '/api/item-units/warehouse/:warehouseId', 'GET', 'ItemUnitController', 'GetByWarehouse', 'API endpoint for GET /api/item-units/warehouse/:warehouseId', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(324, '/api/item', 'POST', 'ItemController', 'CreateItem', 'API endpoint for POST /api/item', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(325, '/api/item/:id', 'GET', 'ItemController', 'GetItem', 'API endpoint for GET /api/item/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(326, '/api/item/:id', 'PUT', 'ItemController', 'UpdateItem', 'API endpoint for PUT /api/item/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(327, '/api/item/:id', 'DELETE', 'ItemController', 'DeleteItem', 'API endpoint for DELETE /api/item/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(328, '/api/items', 'GET', 'ItemController', 'ListItems', 'API endpoint for GET /api/items', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(329, '/api/items/dashboard', 'GET', 'ItemController', 'SearchDashboard', 'API endpoint for GET /api/items/dashboard', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(330, '/api/item/:id/image', 'GET', 'ItemController', 'GetItemImage', 'API endpoint for GET /api/item/:id/image', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(331, '/api/categories', 'GET', 'ItemCategoryController', 'List', 'API endpoint for GET /api/categories', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(332, '/api/categories/:id', 'GET', 'ItemCategoryController', 'Get', 'API endpoint for GET /api/categories/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(333, '/api/categories', 'POST', 'ItemCategoryController', 'Create', 'API endpoint for POST /api/categories', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(334, '/api/categories/:id', 'PUT', 'ItemCategoryController', 'Update', 'API endpoint for PUT /api/categories/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(335, '/api/categories/:id', 'DELETE', 'ItemCategoryController', 'Delete', 'API endpoint for DELETE /api/categories/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(336, '/api/employees', 'POST', 'EmployeeController', 'Create', 'API endpoint for POST /api/employees', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(337, '/api/employees/:id', 'GET', 'EmployeeController', 'Get', 'API endpoint for GET /api/employees/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(338, '/api/employees', 'GET', 'EmployeeController', 'List', 'API endpoint for GET /api/employees', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(339, '/api/employees/:id', 'PUT', 'EmployeeController', 'Update', 'API endpoint for PUT /api/employees/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(340, '/api/employees/:id', 'DELETE', 'EmployeeController', 'Delete', 'API endpoint for DELETE /api/employees/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(341, '/api/lendings', 'POST', 'LendingController', 'Create', 'API endpoint for POST /api/lendings', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(342, '/api/lendings/:id', 'GET', 'LendingController', 'Get', 'API endpoint for GET /api/lendings/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(343, '/api/lendings', 'GET', 'LendingController', 'List', 'API endpoint for GET /api/lendings', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(344, '/api/lendings/:id', 'PUT', 'LendingController', 'Update', 'API endpoint for PUT /api/lendings/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(345, '/api/lendings/:id', 'DELETE', 'LendingController', 'Delete', 'API endpoint for DELETE /api/lendings/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(346, '/api/lendings/report/items', 'GET', 'LendingController', 'SearchItemReport', 'API endpoint for GET /api/lendings/report/items', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(347, '/api/lendings/report/units', 'GET', 'LendingController', 'SearchUnitReport', 'API endpoint for GET /api/lendings/report/units', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(348, '/api/lendings/return/:id', 'PUT', 'LendingController', 'Return', 'API endpoint for PUT /api/lendings/return/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(349, '/api/lending/:id/loan-image', 'GET', 'LendingController', 'GetLoanImage', 'API endpoint for GET /api/lending/:id/loan-image', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(350, '/api/lending/:id/return-image', 'GET', 'LendingController', 'GetReturnImage', 'API endpoint for GET /api/lending/:id/return-image', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(351, '/api/lendings/active', 'GET', 'LendingController', 'GetActiveLoans', 'API endpoint for GET /api/lendings/active', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(352, '/api/lendings/returned', 'GET', 'LendingController', 'GetReturnedLoans', 'API endpoint for GET /api/lendings/returned', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(353, '/api/unit-logs/:id', 'GET', 'UnitLogController', 'Get', 'API endpoint for GET /api/unit-logs/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(354, '/api/unit-logs/:id', 'PUT', 'UnitLogController', 'Update', 'API endpoint for PUT /api/unit-logs/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(355, '/api/unit-logs/:id', 'DELETE', 'UnitLogController', 'Delete', 'API endpoint for DELETE /api/unit-logs/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(356, '/api/unit-logs/unit/:unitId', 'GET', 'UnitLogController', 'GetByUnit', 'API endpoint for GET /api/unit-logs/unit/:unitId', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(357, '/api/repair-logs/:id', 'GET', 'RepairLogController', 'Get', 'API endpoint for GET /api/repair-logs/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(358, '/api/repair-logs/:id', 'DELETE', 'RepairLogController', 'Delete', 'API endpoint for DELETE /api/repair-logs/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(359, '/api/repair-logs/:id', 'PUT', 'RepairLogController', 'Update', 'API endpoint for PUT /api/repair-logs/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(360, '/api/repair-logs/unit/:unitId', 'GET', 'RepairLogController', 'GetByUnit', 'API endpoint for GET /api/repair-logs/unit/:unitId', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(361, '/api/repair-logs/finish', 'POST', 'RepairLogController', 'Finish', 'API endpoint for POST /api/repair-logs/finish', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(362, '/api/docs', 'GET', 'DocUploadedController', 'List', 'API endpoint for GET /api/docs', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(363, '/api/docs/:id', 'GET', 'DocUploadedController', 'Get', 'API endpoint for GET /api/docs/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(364, '/api/docs/:id', 'DELETE', 'DocUploadedController', 'Delete', 'API endpoint for DELETE /api/docs/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(365, '/api/docs/template/download', 'GET', 'DocUploadedController', 'DownloadTemplate', 'API endpoint for GET /api/docs/template/download', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(366, '/api/docs/upload', 'POST', 'DocUploadedController', 'Upload', 'API endpoint for POST /api/docs/upload', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(367, '/api/warehouse', 'POST', 'WarehouseController', 'CreateWarehouse', 'API endpoint for POST /api/warehouse', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(368, '/api/warehouse/:id', 'GET', 'WarehouseController', 'GetWarehouse', 'API endpoint for GET /api/warehouse/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(369, '/api/warehouse/:id', 'PUT', 'WarehouseController', 'UpdateWarehouse', 'API endpoint for PUT /api/warehouse/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(370, '/api/warehouse/:id', 'DELETE', 'WarehouseController', 'DeleteWarehouse', 'API endpoint for DELETE /api/warehouse/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(371, '/api/warehouses', 'GET', 'WarehouseController', 'ListWarehouses', 'API endpoint for GET /api/warehouses', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(372, '/api/warehouse/:id/users', 'GET', 'WarehouseController', 'GetWarehouseUsers', 'API endpoint for GET /api/warehouse/:id/users', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(373, '/api/status/:id', 'GET', 'StatusLookupController', 'GetStatus', 'API endpoint for GET /api/status/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(374, '/api/statuses', 'GET', 'StatusLookupController', 'ListStatuses', 'API endpoint for GET /api/statuses', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(375, '/api/condition/:id', 'GET', 'ConditionLookupController', 'GetCondition', 'API endpoint for GET /api/condition/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(376, '/api/conditions', 'GET', 'ConditionLookupController', 'ListConditions', 'API endpoint for GET /api/conditions', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(377, '/test/condition/:id', 'GET', 'ConditionLookupController', 'GetCondition', 'API endpoint for GET /test/condition/:id', '2024-12-23 13:49:27', '2024-12-23 13:49:27'),
(379, '/test/conditions', 'GET', 'ConditionLookupController', 'ListConditions', 'API endpoint for GET /test/conditions', '2024-12-30 08:32:05', '2024-12-30 08:32:05');

-- --------------------------------------------------------

--
-- Table structure for table `auth_item`
--

CREATE TABLE `auth_item` (
  `id` int(11) NOT NULL,
  `role` varchar(32) NOT NULL,
  `path` varchar(255) NOT NULL,
  `method` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `auth_item`
--

INSERT INTO `auth_item` (`id`, `role`, `path`, `method`) VALUES
(243, 'appadmin', '/api/auth-items', 'GET'),
(244, 'appadmin', '/api/auth-items', 'POST'),
(245, 'appadmin', '/api/auth-items/:id', 'DELETE'),
(246, 'appadmin', '/api/auth-items/:id', 'GET'),
(247, 'appadmin', '/api/auth-items/:id', 'PUT'),
(248, 'appadmin', '/api/auth-items/bulk', 'POST'),
(249, 'appadmin', '/api/categories', 'GET'),
(250, 'appadmin', '/api/categories', 'POST'),
(251, 'appadmin', '/api/categories/:id', 'DELETE'),
(252, 'appadmin', '/api/categories/:id', 'GET'),
(253, 'appadmin', '/api/categories/:id', 'PUT'),
(254, 'appadmin', '/api/condition/:id', 'GET'),
(255, 'appadmin', '/api/conditions', 'GET'),
(256, 'appadmin', '/api/docs', 'GET'),
(257, 'appadmin', '/api/docs/:id', 'DELETE'),
(258, 'appadmin', '/api/docs/:id', 'GET'),
(259, 'appadmin', '/api/docs/template/download', 'GET'),
(260, 'appadmin', '/api/docs/upload', 'POST'),
(261, 'appadmin', '/api/employees', 'GET'),
(262, 'appadmin', '/api/employees', 'POST'),
(263, 'appadmin', '/api/employees/:id', 'DELETE'),
(264, 'appadmin', '/api/employees/:id', 'GET'),
(265, 'appadmin', '/api/employees/:id', 'PUT'),
(266, 'appadmin', '/api/item', 'POST'),
(267, 'appadmin', '/api/item-units', 'GET'),
(268, 'appadmin', '/api/item-units', 'POST'),
(269, 'appadmin', '/api/item-units/:id', 'DELETE'),
(270, 'appadmin', '/api/item-units/:id', 'GET'),
(271, 'appadmin', '/api/item-units/:id', 'PUT'),
(272, 'appadmin', '/api/item-units/serial/:serialNumber', 'GET'),
(273, 'appadmin', '/api/item-units/warehouse/:warehouseId', 'GET'),
(274, 'appadmin', '/api/item/:id', 'DELETE'),
(275, 'appadmin', '/api/item/:id', 'GET'),
(276, 'appadmin', '/api/item/:id', 'PUT'),
(277, 'appadmin', '/api/item/:id/image', 'GET'),
(278, 'appadmin', '/api/items', 'GET'),
(279, 'appadmin', '/api/items/dashboard', 'GET'),
(280, 'appadmin', '/api/lending/:id/loan-image', 'GET'),
(281, 'appadmin', '/api/lending/:id/return-image', 'GET'),
(282, 'appadmin', '/api/lendings', 'GET'),
(283, 'appadmin', '/api/lendings', 'POST'),
(284, 'appadmin', '/api/lendings/:id', 'DELETE'),
(285, 'appadmin', '/api/lendings/:id', 'GET'),
(286, 'appadmin', '/api/lendings/:id', 'PUT'),
(287, 'appadmin', '/api/lendings/active', 'GET'),
(288, 'appadmin', '/api/lendings/report/items', 'GET'),
(289, 'appadmin', '/api/lendings/report/units', 'GET'),
(290, 'appadmin', '/api/lendings/return/:id', 'PUT'),
(291, 'appadmin', '/api/lendings/returned', 'GET'),
(292, 'appadmin', '/api/repair-logs/:id', 'DELETE'),
(293, 'appadmin', '/api/repair-logs/:id', 'GET'),
(294, 'appadmin', '/api/repair-logs/:id', 'PUT'),
(295, 'appadmin', '/api/repair-logs/finish', 'POST'),
(296, 'appadmin', '/api/repair-logs/unit/:unitId', 'GET'),
(297, 'appadmin', '/api/roles', 'GET'),
(298, 'appadmin', '/api/roles', 'POST'),
(299, 'appadmin', '/api/roles/:id', 'DELETE'),
(300, 'appadmin', '/api/roles/:id', 'GET'),
(301, 'appadmin', '/api/roles/:id', 'PUT'),
(302, 'appadmin', '/api/routes/:id', 'DELETE'),
(303, 'appadmin', '/api/routes/:id', 'GET'),
(304, 'appadmin', '/api/routes/list', 'GET'),
(305, 'appadmin', '/api/routes/scan', 'POST'),
(306, 'appadmin', '/api/status/:id', 'GET'),
(307, 'appadmin', '/api/statuses', 'GET'),
(308, 'appadmin', '/api/unit-logs/:id', 'DELETE'),
(309, 'appadmin', '/api/unit-logs/:id', 'GET'),
(310, 'appadmin', '/api/unit-logs/:id', 'PUT'),
(311, 'appadmin', '/api/unit-logs/unit/:unitId', 'GET'),
(312, 'appadmin', '/api/user-roles', 'POST'),
(313, 'appadmin', '/api/user-roles/:userId/:roleId', 'DELETE'),
(314, 'appadmin', '/api/user-roles/role/:roleId', 'GET'),
(315, 'appadmin', '/api/user-roles/user/:userId', 'GET'),
(316, 'appadmin', '/api/warehouse', 'POST'),
(317, 'appadmin', '/api/warehouse/:id', 'DELETE'),
(318, 'appadmin', '/api/warehouse/:id', 'GET'),
(319, 'appadmin', '/api/warehouse/:id', 'PUT'),
(320, 'appadmin', '/api/warehouse/:id/users', 'GET'),
(321, 'appadmin', '/api/warehouses', 'GET'),
(322, 'appadmin', '/auth/login', 'POST'),
(323, 'appadmin', '/auth/logout', 'POST'),
(324, 'appadmin', '/test/condition/:id', 'GET'),
(326, 'appadmin', '/user', 'POST'),
(327, 'appadmin', '/user/:id', 'DELETE'),
(328, 'appadmin', '/user/:id', 'GET'),
(329, 'appadmin', '/user/:id', 'PUT'),
(330, 'appadmin', '/user/:id/visits', 'GET'),
(331, 'appadmin', '/users', 'GET'),
(332, 'appadmin', '/test/conditions', 'GET');

-- --------------------------------------------------------

--
-- Table structure for table `auth_roles`
--

CREATE TABLE `auth_roles` (
  `code` varchar(32) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` datetime(2) NOT NULL,
  `updated_at` datetime(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `auth_roles`
--

INSERT INTO `auth_roles` (`code`, `name`, `description`, `created_at`, `updated_at`) VALUES
('appadmin', 'App Admin', 'Can manage the app, users, employee, and warehouse', '2024-12-19 02:48:38.76', '2024-12-19 02:48:38.76'),
('mro', 'Maintenance & Repair Officer', 'Can manage damaged unit. Sending Unit to repairs and accepting completed repairs to be sent back to warehouses', '2024-12-19 02:50:09.88', '2024-12-19 02:50:09.88'),
('super', 'CHECKING 2', 'CHECKING 2', '0000-00-00 00:00:00.00', '2024-12-19 02:57:06.31'),
('whadmin', 'warehouse_admin', 'Can manage warehouse inventory, lent out items, accept returns, sending unit to repairs', '2024-12-19 02:47:28.56', '2024-12-19 02:47:28.56');

-- --------------------------------------------------------

--
-- Table structure for table `auth_roles_user`
--

CREATE TABLE `auth_roles_user` (
  `user_id` int(11) NOT NULL,
  `roles_code` varchar(32) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `auth_roles_user`
--

INSERT INTO `auth_roles_user` (`user_id`, `roles_code`) VALUES
(1, 'appadmin'),
(22, 'super');

-- --------------------------------------------------------

--
-- Table structure for table `condition_lookup`
--

CREATE TABLE `condition_lookup` (
  `id_condition` tinyint(3) UNSIGNED NOT NULL COMMENT 'primary key',
  `condition_name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `condition_lookup`
--

INSERT INTO `condition_lookup` (`id_condition`, `condition_name`) VALUES
(1, 'OK'),
(2, 'Light Damage'),
(3, 'Moderately damaged (missing a part or component)'),
(4, 'Major damage (inoperable, repair required)'),
(5, 'Lost or Destroyed');

-- --------------------------------------------------------

--
-- Table structure for table `doc_uploaded`
--

CREATE TABLE `doc_uploaded` (
  `id_doc` int(11) NOT NULL COMMENT 'primary key',
  `file_name` varchar(255) NOT NULL,
  `datetime` datetime(2) NOT NULL,
  `user_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `doc_uploaded`
--

INSERT INTO `doc_uploaded` (`id_doc`, `file_name`, `datetime`, `user_id`) VALUES
(43, 'bulk_unit613_1732757427.xlsx', '2024-11-28 08:30:27.00', 14),
(44, 'bulk_unit207_1732757517.xlsx', '2024-11-28 08:31:57.00', 14),
(45, 'bulk_unit127_1732757518.xlsx', '2024-11-28 08:31:58.00', 14),
(46, 'bulk_unit946_1732757518.xlsx', '2024-11-28 08:31:58.00', 14),
(47, 'bulk_unit873_1732757519.xlsx', '2024-11-28 08:31:59.00', 14),
(48, 'bulk_unit531_1732757520.xlsx', '2024-11-28 08:32:00.00', 14),
(49, 'bulk_unit121_1732757520.xlsx', '2024-11-28 08:32:00.00', 14),
(50, 'bulk_unit775_1732757521.xlsx', '2024-11-28 08:32:01.00', 14),
(51, 'bulk_unit573_1732757521.xlsx', '2024-11-28 08:32:01.00', 14),
(52, 'bulk_unit249_1732757522.xlsx', '2024-11-28 08:32:02.00', 14),
(53, 'bulk_unit707_1732757523.xlsx', '2024-11-28 08:32:03.00', 14),
(54, 'bulk_unit112_1732757523.xlsx', '2024-11-28 08:32:03.00', 14),
(55, 'bulk_unit991_1732757524.xlsx', '2024-11-28 08:32:04.00', 14),
(56, 'bulk_unit755_1732757524.xlsx', '2024-11-28 08:32:04.00', 14),
(57, 'bulk_unit511_1732757525.xlsx', '2024-11-28 08:32:05.00', 14),
(58, 'bulk_unit210_1732757525.xlsx', '2024-11-28 08:32:05.00', 14),
(59, 'bulk_unit940_1732757590.xlsx', '2024-11-28 08:33:10.00', 14),
(60, 'bulk_unit742_1732757590.xlsx', '2024-11-28 08:33:10.00', 14),
(61, 'bulk_unit493_1732757600.xlsx', '2024-11-28 08:33:20.00', 14),
(62, 'bulk_unit866_1732757612.xlsx', '2024-11-28 08:33:32.00', 14),
(63, 'bulk_unit700_1732757624.xlsx', '2024-11-28 08:33:44.00', 14),
(64, 'bulk_unit803_1732757636.xlsx', '2024-11-28 08:33:56.00', 14),
(65, 'bulk_unit503_1732757648.xlsx', '2024-11-28 08:34:08.00', 14),
(66, 'bulk_unit919_1732757658.xlsx', '2024-11-28 08:34:18.00', 14),
(67, 'bulk_unit118_1732757667.xlsx', '2024-11-28 08:34:27.00', 14),
(68, 'bulk_unit524_1732757679.xlsx', '2024-11-28 08:34:39.00', 14),
(69, 'bulk_unit658_1732757689.xlsx', '2024-11-28 08:34:49.00', 14),
(70, 'bulk_unit914_1732757699.xlsx', '2024-11-28 08:34:59.00', 14),
(71, 'bulk_unit131_1732757713.xlsx', '2024-11-28 08:35:13.00', 14),
(72, 'bulk_unit229_1732757729.xlsx', '2024-11-28 08:35:29.00', 14),
(74, 'bulk_unit528_1732765464.xlsx', '2024-11-28 10:44:24.00', 5),
(75, 'bulk_unit456_1732765579.xlsx', '2024-11-28 10:46:19.00', 5),
(76, 'bulk_unit313_1732765621.xlsx', '2024-11-28 10:47:01.00', 5),
(77, 'bulk_unit853_1732765648.xlsx', '2024-11-28 10:47:28.00', 5),
(78, 'bulk_unit649_1732765775.xlsx', '2024-11-28 10:49:35.00', 5),
(79, 'bulk_unit625_1732777053.xlsx', '2024-11-28 13:57:33.00', 5),
(80, 'bulk_unit714_1732868440.xlsx', '2024-11-29 15:20:40.00', 1),
(81, 'bulk_unit988_1733103037.xlsx', '2024-12-02 08:30:37.00', 1),
(82, 'bulk_unit732_1733189780.xlsx', '2024-12-03 08:36:20.00', 1),
(83, 'bulk_unit392_1733209215.xlsx', '2024-12-03 14:00:15.00', 1),
(84, 'bulk_unit974_1733449957.xlsx', '2024-12-06 08:52:37.00', 1),
(85, 'bulk_unit609_1733450484.xlsx', '2024-12-06 09:01:24.00', 1),
(86, 'bulk_unit191_1733467508.xlsx', '2024-12-06 12:45:08.00', 5),
(87, 'bulk_unit498_1733467548.xlsx', '2024-12-06 12:45:48.00', 1),
(89, 'maintenance_report_dec2023.pdf', '2024-12-13 07:27:51.88', 1),
(90, 'maintenance_report_dec2023.pdf', '2024-12-16 06:35:41.67', 1),
(91, 'maintenance_report_dec2023.pdf', '2024-12-17 03:41:31.98', 1),
(92, 'maintenance_report_dec2023.pdf', '2024-12-17 03:41:37.12', 1),
(93, 'maintenance_report_dec2023.pdf', '2024-12-17 03:41:42.87', 1),
(94, 'maintenance_report_dec2023.pdf', '2024-12-17 03:41:48.03', 1),
(95, 'maintenance_report_dec2023.pdf', '2024-12-17 03:41:53.53', 1),
(96, 'maintenance_report_dec2023.pdf', '2024-12-17 03:41:59.12', 1),
(97, 'maintenance_report_dec2023.pdf', '2024-12-17 03:42:04.44', 1),
(98, 'maintenance_report_dec2023.pdf', '2024-12-17 03:42:09.72', 1),
(99, 'maintenance_report_dec2023.pdf', '2024-12-17 03:42:15.26', 1),
(100, 'maintenance_report_dec2023.pdf', '2024-12-17 03:42:20.63', 1),
(101, '18-12-2024-090524.xlsx', '2024-12-18 02:05:24.97', 1),
(102, '18-12-2024-090726.xlsx', '2024-12-18 02:07:26.09', 1),
(103, '18-12-2024-091008.xlsx', '2024-12-18 02:10:08.53', 1),
(104, 'BULK_UPLOAD-18-12-2024-091137.xlsx', '2024-12-18 02:11:37.09', 1),
(105, 'maintenance_report_dec2023.pdf', '2024-12-18 02:24:08.86', 1),
(106, 'maintenance_report_dec2023.pdf', '2024-12-18 02:26:02.96', 1),
(107, 'maintenance_report_dec2023.pdf', '2024-12-18 02:28:40.80', 1),
(108, 'maintenance_report_dec2023.pdf', '2024-12-30 01:18:49.89', 1),
(109, 'maintenance_report_dec2023.pdf', '2024-12-30 01:19:15.07', 1),
(110, 'maintenance_report_dec2023.pdf', '2024-12-30 01:21:53.69', 1),
(111, 'maintenance_report_dec2023.pdf', '2024-12-30 01:30:19.98', 1),
(112, 'maintenance_report_dec2023.pdf', '2024-12-30 01:32:04.61', 1);

-- --------------------------------------------------------

--
-- Table structure for table `employee`
--

CREATE TABLE `employee` (
  `id_employee` int(10) UNSIGNED NOT NULL COMMENT 'primary key',
  `emp_name` varchar(80) NOT NULL COMMENT 'employee name',
  `phone` varchar(20) NOT NULL,
  `email` varchar(60) NOT NULL,
  `address` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `employee`
--

INSERT INTO `employee` (`id_employee`, `emp_name`, `phone`, `email`, `address`) VALUES
(1, 'Johan', '0888888', 'ayymail@lmao.com', 'jalan jalan aj'),
(2, 'Emma', '0812121211212', 'emma@mail.com', 'a random street name'),
(3, 'Oleg', '08080808', 'tumbal@mail.com', 'jalan gk tau'),
(4, 'Freddy', '0821321321', 'freddymail@mail.com', 'somewhere in a city'),
(5, 'Shiorin', '081684316487', 'shiorin@ayymail.com', 'idk somewhere a'),
(7, 'Mark', '080818123215', 'Mark@mail.com', 'idk somewhere in us continent'),
(8, 'Manfred Albrecht Freiherr von Richthofen', '0812354612324', 'redbaronrulerofthesky@mail.com', 'somewhere in German'),
(10, 'SPAM PREVENTION TEST', 'SPAM PREVENTION TEST', 'SPAM PREVENTION TEST', 'SPAM PREVENTION TEST'),
(12, 'Jane Smith API test', '087654321098', 'jane.smith@company.com', '456 Park Avenue API'),
(15, 'John Smith', '087654321098', 'John.smith@company.com', '456 Park Avenue API'),
(32, 'John Smith', '087654321018', 'John.smithy@company.com', '456 Park Avenue API');

-- --------------------------------------------------------

--
-- Table structure for table `item`
--

CREATE TABLE `item` (
  `id_item` int(10) UNSIGNED NOT NULL COMMENT 'primary key',
  `item_name` varchar(60) NOT NULL,
  `SKU` varchar(50) NOT NULL COMMENT 'Stock Keeping Unit Code',
  `imagefile` varchar(255) NOT NULL COMMENT 'item pic',
  `id_category` int(10) DEFAULT NULL COMMENT 'foreign key'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `item`
--

INSERT INTO `item` (`id_item`, `item_name`, `SKU`, `imagefile`, `id_category`) VALUES
(11, 'Logitech Signature Slim Keyboard K950', 'PA24-0001', '736_1731394912.webp', 1),
(12, 'Logitech MX Anywhere 3S', 'PA24-0002', '304_1731394932.webp', 1),
(13, 'ThinkPad T14s Gen 4 (14″ AMD)', 'PC24-0001', '859_1731395020.jpg', 3),
(14, 'Generic Mousepad Small', 'PA24-0003', '409_1731395120.jpeg', 1),
(15, 'Sennheisser HD600', 'VD24-0001', '959_1731395082.jpeg', 2),
(16, 'Shure SM7B', 'VD24-0002', '121_1731395169.jpeg', 2),
(18, 'TC-Helicon GOXLR', 'VD24-0003', '903_1731395204.jpeg', 2),
(25, 'USB-C Hub Multi Dongle', 'PA24-0004', '624_1731395239.jpeg', 1),
(26, 'Macbook Air M3', 'PC24-0002', '116_1731395279.jpeg', 3),
(27, 'Sony A7 iv', 'VD24-0004', '302_1731460840.webp', 2),
(28, 'Test FlakPz Gepard', 'TC24-0001', '945_1732075232.jpeg', 4),
(29, 'Test Item 1', 'TC24-0002', '382_1732155439.jpeg', 4),
(30, 'Test G1', 'TC24-0003', '744_1732155555.jpg', 4),
(31, 'Laptop Backpack M24', 'PU24-0001', '646_1732515992.jpeg', 7),
(32, 'Suffers Mk1', 'TC24-0004', '476_1732519661.jpeg', 4),
(33, 'Computer Desk Mk1 M2024', 'DK24-0001', '231_1732848148.jpeg', 9),
(36, 'SPAM TEST ', 'TC24-0005', '239_1733102912.jpeg', 4),
(37, 'Optimization Clean Up', 'TC24-0006', '414_1733470397.jpeg', 4),
(38, 'TEST ITEM AC', 'TC24-0007', '581_1733717455.gif', 4),
(39, 'Logitech G Pro X Superlight White api test', 'PA24-0009', 'mouse_white api test.webp', 1),
(42, 'Logitech G Pro X Superlight', 'PA24-0999', 'mouse.webp', 1),
(44, 'Logitech G Pro X Superlight White api test A1', 'PA24-0024', 'mouse_white api test.webp', 1),
(46, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0008', 'mouse.webp', 1),
(51, 'Logitech G Pro X Superlight Ultra Mk 2C', 'PA24-0005', 'mouse.webp', 1),
(52, 'Logitech G Pro X Superlight Ultra Mk 2C', 'PA24-0006', 'mouse.webp', 1),
(53, 'Logitech G Pro X Superlight Ultra Mk 2C', 'PA24-0007', 'mouse.webp', 1),
(54, 'Logitech G Pro X Superlight Ultra Mk 2C', 'PA24-0010', 'mouse.webp', 1),
(55, 'Logitech G Pro X Superlight Ultra Mk 2C', 'PA24-0011', 'mouse.webp', 1),
(56, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0013', 'mouse.webp', 1),
(57, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0014', 'mouse.webp', 1),
(58, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0012', 'mouse.webp', 1),
(59, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0015', 'mouse.webp', 1),
(60, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0016', 'mouse.webp', 1),
(61, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0017', 'mouse.webp', 1),
(62, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0018', 'mouse.webp', 1),
(63, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0019', 'mouse.webp', 1),
(64, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0020', 'mouse.webp', 1),
(65, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0021', 'mouse.webp', 1),
(66, 'Logitech G Pro X Superlight Ultra Mk 2', 'PA24-0022', 'mouse.webp', 1);

-- --------------------------------------------------------

--
-- Table structure for table `item_category`
--

CREATE TABLE `item_category` (
  `id_category` int(10) NOT NULL COMMENT 'primary key',
  `category_name` varchar(60) NOT NULL,
  `cat_code` varchar(6) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `item_category`
--

INSERT INTO `item_category` (`id_category`, `category_name`, `cat_code`) VALUES
(1, 'Computer Peripherals', 'PA'),
(2, 'Audio Visual Production', 'VD'),
(3, 'Personal Computer', 'PC'),
(4, 'Test Cat', 'TC'),
(6, 'Test C2', 'TS'),
(7, 'Personal Utility Item', 'PU'),
(9, 'Desk', 'DK'),
(14, 'Office Supplies API test', 'OF'),
(16, 'Pneumatic Power Tool', 'PT');

-- --------------------------------------------------------

--
-- Table structure for table `item_unit`
--

CREATE TABLE `item_unit` (
  `id_unit` int(10) UNSIGNED NOT NULL COMMENT 'primary key',
  `id_item` int(10) UNSIGNED NOT NULL COMMENT 'foreign key - item',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '1 = available\r\n2 = in-use\r\n3 = in-repair\r\n4 = lost or destroyed',
  `id_wh` int(10) UNSIGNED DEFAULT NULL COMMENT 'where is the unit stored',
  `comment` varchar(60) DEFAULT NULL COMMENT 'additional info on unit',
  `serial_number` varchar(60) NOT NULL,
  `condition` tinyint(3) UNSIGNED NOT NULL COMMENT 'refer to condition_lookup ',
  `updated_by` int(11) DEFAULT NULL COMMENT 'last user that interacted with the data'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `item_unit`
--

INSERT INTO `item_unit` (`id_unit`, `id_item`, `status`, `id_wh`, `comment`, `serial_number`, `condition`, `updated_by`) VALUES
(1, 11, 1, 5, 'optimization test', 'PA24-0001-0001', 3, 1),
(2, 11, 2, 6, 'new data after purge', 'PA24-0001-0002', 1, 1),
(3, 11, 1, 7, 'test ret', 'PA24-0001-0003', 1, 1),
(4, 11, 2, 8, 'new data after purge', 'PA24-0001-0004', 1, 1),
(5, 11, 1, 9, 'new data after purge', 'PA24-0001-0005', 1, 14),
(6, 11, 1, 10, 'new data after purge', 'PA24-0001-0006', 1, 14),
(7, 11, 1, 11, 'new data after purge', 'PA24-0001-0007', 1, 14),
(8, 11, 1, 13, 'new data after purge', 'PA24-0001-0008', 1, 14),
(9, 11, 1, 5, 'new data after purge', 'PA24-0001-0009', 1, 14),
(10, 11, 1, 6, 'new data after purge', 'PA24-0001-0010', 1, 14),
(11, 11, 1, 7, 'new data after purge', 'PA24-0001-0011', 1, 14),
(12, 11, 1, 8, 'new data after purge', 'PA24-0001-0012', 1, 14),
(13, 11, 1, 9, 'new data after purge', 'PA24-0001-0013', 1, 14),
(14, 11, 1, 10, 'new data after purge', 'PA24-0001-0014', 1, 14),
(15, 11, 1, 11, 'new data after purge', 'PA24-0001-0015', 1, 14),
(16, 11, 1, 13, 'new data after purge', 'PA24-0001-0016', 1, 14),
(17, 11, 1, 5, 'new data after purge', 'PA24-0001-0017', 1, 14),
(18, 11, 1, 6, 'new data after purge', 'PA24-0001-0018', 1, 14),
(19, 11, 1, 7, 'new data after purge', 'PA24-0001-0019', 1, 14),
(20, 11, 1, 8, 'new data after purge', 'PA24-0001-0020', 1, 14),
(21, 11, 1, 9, 'new data after purge', 'PA24-0001-0021', 1, 14),
(22, 11, 1, 10, 'new data after purge', 'PA24-0001-0022', 1, 14),
(23, 11, 1, 11, 'new data after purge', 'PA24-0001-0023', 1, 14),
(24, 11, 1, 13, 'new data after purge', 'PA24-0001-0024', 1, 14),
(25, 11, 1, 5, 'new data after purge', 'PA24-0001-0025', 1, 14),
(26, 11, 1, 6, 'new data after purge', 'PA24-0001-0026', 1, 14),
(27, 11, 1, 7, 'new data after purge', 'PA24-0001-0027', 1, 14),
(28, 11, 1, 8, 'new data after purge', 'PA24-0001-0028', 1, 14),
(29, 11, 1, 9, 'new data after purge', 'PA24-0001-0029', 1, 14),
(30, 11, 1, 10, 'new data after purge', 'PA24-0001-0030', 1, 14),
(31, 11, 1, 11, 'new data after purge', 'PA24-0001-0031', 1, 14),
(32, 11, 1, 13, 'new data after purge', 'PA24-0001-0032', 1, 14),
(33, 11, 1, 5, 'new data after purge', 'PA24-0001-0033', 1, 14),
(34, 11, 1, 6, 'new data after purge', 'PA24-0001-0034', 1, 14),
(35, 11, 1, 7, 'new data after purge', 'PA24-0001-0035', 1, 14),
(36, 11, 1, 8, 'new data after purge', 'PA24-0001-0036', 1, 14),
(37, 11, 1, 9, 'new data after purge', 'PA24-0001-0037', 1, 14),
(38, 11, 1, 10, 'new data after purge', 'PA24-0001-0038', 1, 14),
(39, 11, 1, 11, 'new data after purge', 'PA24-0001-0039', 1, 14),
(40, 11, 1, 13, 'new data after purge', 'PA24-0001-0040', 1, 14),
(41, 11, 1, 5, 'new data after purge', 'PA24-0001-0041', 1, 14),
(42, 11, 1, 6, 'new data after purge', 'PA24-0001-0042', 1, 14),
(43, 11, 1, 7, 'new data after purge', 'PA24-0001-0043', 1, 14),
(44, 11, 1, 8, 'new data after purge', 'PA24-0001-0044', 1, 14),
(45, 11, 1, 9, 'new data after purge', 'PA24-0001-0045', 1, 14),
(46, 11, 1, 10, 'new data after purge', 'PA24-0001-0046', 1, 14),
(47, 11, 1, 11, 'new data after purge', 'PA24-0001-0047', 1, 14),
(48, 11, 1, 13, 'new data after purge', 'PA24-0001-0048', 1, 14),
(49, 11, 1, 5, 'new data after purge', 'PA24-0001-0049', 1, 14),
(50, 11, 1, 6, 'new data after purge', 'PA24-0001-0050', 1, 14),
(51, 11, 1, 7, 'new data after purge', 'PA24-0001-0051', 1, 14),
(52, 11, 1, 8, 'new data after purge', 'PA24-0001-0052', 1, 14),
(53, 11, 1, 9, 'new data after purge', 'PA24-0001-0053', 1, 14),
(54, 11, 1, 10, 'new data after purge', 'PA24-0001-0054', 1, 14),
(55, 11, 1, 11, 'new data after purge', 'PA24-0001-0055', 1, 14),
(56, 11, 1, 13, 'new data after purge', 'PA24-0001-0056', 1, 14),
(57, 11, 1, 5, 'new data after purge', 'PA24-0001-0057', 1, 14),
(58, 11, 1, 6, 'new data after purge', 'PA24-0001-0058', 1, 14),
(59, 11, 1, 7, 'new data after purge', 'PA24-0001-0059', 1, 14),
(60, 11, 1, 8, 'new data after purge', 'PA24-0001-0060', 1, 14),
(61, 11, 1, 9, 'new data after purge', 'PA24-0001-0061', 1, 14),
(62, 11, 1, 10, 'new data after purge', 'PA24-0001-0062', 1, 14),
(63, 11, 1, 11, 'new data after purge', 'PA24-0001-0063', 1, 14),
(64, 11, 1, 13, 'new data after purge', 'PA24-0001-0064', 1, 14),
(65, 33, 1, 5, 'new data after purge', 'DK24-0001-0001', 1, 14),
(66, 33, 2, 6, 'new data after purge', 'DK24-0001-0002', 1, 14),
(67, 33, 1, 7, 'new data after purge', 'DK24-0001-0003', 1, 14),
(68, 33, 1, 8, 'new data after purge', 'DK24-0001-0004', 1, 14),
(69, 33, 1, 9, 'new data after purge', 'DK24-0001-0005', 1, 14),
(70, 33, 1, 10, 'new data after purge', 'DK24-0001-0006', 1, 14),
(71, 33, 1, 11, 'new data after purge', 'DK24-0001-0007', 1, 14),
(72, 33, 1, 13, 'new data after purge', 'DK24-0001-0008', 1, 14),
(73, 33, 1, 5, 'new data after purge', 'DK24-0001-0009', 1, 14),
(74, 33, 1, 6, 'new data after purge', 'DK24-0001-0010', 1, 14),
(75, 33, 1, 7, 'new data after purge', 'DK24-0001-0011', 1, 14),
(76, 33, 1, 8, 'new data after purge', 'DK24-0001-0012', 1, 14),
(77, 33, 1, 9, 'new data after purge', 'DK24-0001-0013', 1, 14),
(78, 33, 1, 10, 'new data after purge', 'DK24-0001-0014', 1, 14),
(79, 33, 1, 11, 'new data after purge', 'DK24-0001-0015', 1, 14),
(80, 33, 1, 13, 'new data after purge', 'DK24-0001-0016', 1, 14),
(81, 33, 1, 5, 'new data after purge', 'DK24-0001-0017', 1, 14),
(82, 33, 1, 6, 'new data after purge', 'DK24-0001-0018', 1, 14),
(83, 33, 1, 7, 'new data after purge', 'DK24-0001-0019', 1, 14),
(84, 33, 1, 8, 'new data after purge', 'DK24-0001-0020', 1, 14),
(85, 33, 1, 9, 'new data after purge', 'DK24-0001-0021', 1, 14),
(86, 33, 1, 10, 'new data after purge', 'DK24-0001-0022', 1, 14),
(87, 33, 1, 11, 'new data after purge', 'DK24-0001-0023', 1, 14),
(88, 33, 1, 13, 'new data after purge', 'DK24-0001-0024', 1, 14),
(89, 33, 1, 5, 'new data after purge', 'DK24-0001-0025', 1, 14),
(90, 33, 1, 6, 'new data after purge', 'DK24-0001-0026', 1, 14),
(91, 33, 1, 7, 'new data after purge', 'DK24-0001-0027', 1, 14),
(92, 33, 1, 8, 'new data after purge', 'DK24-0001-0028', 1, 14),
(93, 33, 1, 9, 'new data after purge', 'DK24-0001-0029', 1, 14),
(94, 33, 1, 10, 'new data after purge', 'DK24-0001-0030', 1, 14),
(95, 33, 1, 11, 'new data after purge', 'DK24-0001-0031', 1, 14),
(96, 33, 1, 13, 'new data after purge', 'DK24-0001-0032', 1, 14),
(97, 33, 1, 5, 'new data after purge', 'DK24-0001-0033', 1, 14),
(98, 33, 1, 6, 'new data after purge', 'DK24-0001-0034', 1, 14),
(99, 33, 1, 7, 'new data after purge', 'DK24-0001-0035', 1, 14),
(100, 33, 1, 8, 'new data after purge', 'DK24-0001-0036', 1, 14),
(101, 33, 1, 9, 'new data after purge', 'DK24-0001-0037', 1, 14),
(102, 33, 1, 10, 'new data after purge', 'DK24-0001-0038', 1, 14),
(103, 33, 1, 11, 'new data after purge', 'DK24-0001-0039', 1, 14),
(104, 33, 1, 13, 'new data after purge', 'DK24-0001-0040', 1, 14),
(105, 33, 1, 5, 'new data after purge', 'DK24-0001-0041', 1, 14),
(106, 33, 1, 6, 'new data after purge', 'DK24-0001-0042', 1, 14),
(107, 33, 1, 7, 'new data after purge', 'DK24-0001-0043', 1, 14),
(108, 33, 1, 8, 'new data after purge', 'DK24-0001-0044', 1, 14),
(109, 33, 1, 9, 'new data after purge', 'DK24-0001-0045', 1, 14),
(110, 33, 1, 10, 'new data after purge', 'DK24-0001-0046', 1, 14),
(111, 33, 1, 11, 'new data after purge', 'DK24-0001-0047', 1, 14),
(112, 33, 1, 13, 'new data after purge', 'DK24-0001-0048', 1, 14),
(113, 33, 1, 5, 'new data after purge', 'DK24-0001-0049', 1, 14),
(114, 33, 1, 6, 'new data after purge', 'DK24-0001-0050', 1, 14),
(115, 33, 1, 7, 'new data after purge', 'DK24-0001-0051', 1, 14),
(116, 33, 1, 8, 'new data after purge', 'DK24-0001-0052', 1, 14),
(117, 33, 1, 9, 'new data after purge', 'DK24-0001-0053', 1, 14),
(118, 33, 1, 10, 'new data after purge', 'DK24-0001-0054', 1, 14),
(119, 33, 1, 11, 'new data after purge', 'DK24-0001-0055', 1, 14),
(120, 33, 1, 13, 'new data after purge', 'DK24-0001-0056', 1, 14),
(121, 33, 1, 5, 'new data after purge', 'DK24-0001-0057', 1, 14),
(122, 33, 1, 6, 'new data after purge', 'DK24-0001-0058', 1, 14),
(123, 33, 1, 7, 'new data after purge', 'DK24-0001-0059', 1, 14),
(124, 33, 1, 8, 'new data after purge', 'DK24-0001-0060', 1, 14),
(125, 33, 1, 9, 'new data after purge', 'DK24-0001-0061', 1, 14),
(126, 33, 1, 10, 'new data after purge', 'DK24-0001-0062', 1, 14),
(127, 33, 1, 11, 'new data after purge', 'DK24-0001-0063', 1, 14),
(128, 33, 1, 13, 'new data after purge', 'DK24-0001-0064', 1, 14),
(129, 12, 1, 13, 'optomization test', 'PA24-0002-0001', 2, 1),
(130, 12, 2, 6, 'new data after purge', 'PA24-0002-0002', 1, 1),
(131, 12, 1, 7, 'optimization test', 'PA24-0002-0003', 1, 1),
(132, 12, 1, 8, 'new data after purge', 'PA24-0002-0004', 1, 14),
(133, 12, 1, 9, 'new data after purge', 'PA24-0002-0005', 1, 14),
(134, 12, 1, 10, 'new data after purge', 'PA24-0002-0006', 1, 14),
(135, 12, 1, 11, 'new data after purge', 'PA24-0002-0007', 1, 14),
(136, 12, 1, 13, 'new data after purge', 'PA24-0002-0008', 1, 14),
(137, 12, 1, 5, 'new data after purge', 'PA24-0002-0009', 1, 14),
(138, 12, 1, 6, 'new data after purge', 'PA24-0002-0010', 1, 14),
(139, 12, 1, 7, 'new data after purge', 'PA24-0002-0011', 1, 14),
(140, 12, 1, 8, 'new data after purge', 'PA24-0002-0012', 1, 14),
(141, 12, 1, 9, 'new data after purge', 'PA24-0002-0013', 1, 14),
(142, 12, 1, 10, 'new data after purge', 'PA24-0002-0014', 1, 14),
(143, 12, 1, 11, 'new data after purge', 'PA24-0002-0015', 1, 14),
(144, 12, 1, 13, 'new data after purge', 'PA24-0002-0016', 1, 14),
(145, 12, 1, 5, 'new data after purge', 'PA24-0002-0017', 1, 14),
(146, 12, 1, 6, 'new data after purge', 'PA24-0002-0018', 1, 14),
(147, 12, 1, 7, 'new data after purge', 'PA24-0002-0019', 1, 14),
(148, 12, 1, 8, 'new data after purge', 'PA24-0002-0020', 1, 14),
(149, 12, 1, 9, 'new data after purge', 'PA24-0002-0021', 1, 14),
(150, 12, 1, 10, 'new data after purge', 'PA24-0002-0022', 1, 14),
(151, 12, 1, 11, 'new data after purge', 'PA24-0002-0023', 1, 14),
(152, 12, 1, 13, 'new data after purge', 'PA24-0002-0024', 1, 14),
(153, 12, 1, 5, 'new data after purge', 'PA24-0002-0025', 1, 14),
(154, 12, 1, 6, 'new data after purge', 'PA24-0002-0026', 1, 14),
(155, 12, 1, 7, 'new data after purge', 'PA24-0002-0027', 1, 14),
(156, 12, 1, 8, 'new data after purge', 'PA24-0002-0028', 1, 14),
(157, 12, 1, 9, 'new data after purge', 'PA24-0002-0029', 1, 14),
(158, 12, 1, 10, 'new data after purge', 'PA24-0002-0030', 1, 14),
(159, 12, 1, 11, 'new data after purge', 'PA24-0002-0031', 1, 14),
(160, 12, 1, 13, 'new data after purge', 'PA24-0002-0032', 1, 14),
(161, 12, 1, 5, 'new data after purge', 'PA24-0002-0033', 1, 14),
(162, 12, 1, 6, 'new data after purge', 'PA24-0002-0034', 1, 14),
(163, 12, 1, 7, 'new data after purge', 'PA24-0002-0035', 1, 14),
(164, 12, 1, 8, 'new data after purge', 'PA24-0002-0036', 1, 14),
(165, 12, 1, 9, 'new data after purge', 'PA24-0002-0037', 1, 14),
(166, 12, 1, 10, 'new data after purge', 'PA24-0002-0038', 1, 14),
(167, 12, 1, 11, 'new data after purge', 'PA24-0002-0039', 1, 14),
(168, 12, 1, 13, 'new data after purge', 'PA24-0002-0040', 1, 14),
(169, 12, 1, 5, 'new data after purge', 'PA24-0002-0041', 1, 14),
(170, 12, 1, 6, 'new data after purge', 'PA24-0002-0042', 1, 14),
(171, 12, 1, 7, 'new data after purge', 'PA24-0002-0043', 1, 14),
(172, 12, 1, 8, 'new data after purge', 'PA24-0002-0044', 1, 14),
(173, 12, 1, 9, 'new data after purge', 'PA24-0002-0045', 1, 14),
(174, 12, 1, 10, 'new data after purge', 'PA24-0002-0046', 1, 14),
(175, 12, 1, 11, 'new data after purge', 'PA24-0002-0047', 1, 14),
(176, 12, 1, 13, 'new data after purge', 'PA24-0002-0048', 1, 14),
(177, 12, 1, 5, 'new data after purge', 'PA24-0002-0049', 1, 14),
(178, 12, 1, 6, 'new data after purge', 'PA24-0002-0050', 1, 14),
(179, 12, 1, 7, 'new data after purge', 'PA24-0002-0051', 1, 14),
(180, 12, 1, 8, 'new data after purge', 'PA24-0002-0052', 1, 14),
(181, 12, 1, 9, 'new data after purge', 'PA24-0002-0053', 1, 14),
(182, 12, 1, 10, 'new data after purge', 'PA24-0002-0054', 1, 14),
(183, 12, 1, 11, 'new data after purge', 'PA24-0002-0055', 1, 14),
(184, 12, 1, 13, 'new data after purge', 'PA24-0002-0056', 1, 14),
(185, 12, 1, 5, 'new data after purge', 'PA24-0002-0057', 1, 14),
(186, 12, 1, 6, 'new data after purge', 'PA24-0002-0058', 1, 14),
(187, 12, 1, 7, 'new data after purge', 'PA24-0002-0059', 1, 14),
(188, 12, 1, 8, 'new data after purge', 'PA24-0002-0060', 1, 14),
(189, 12, 1, 9, 'new data after purge', 'PA24-0002-0061', 1, 14),
(190, 12, 1, 10, 'new data after purge', 'PA24-0002-0062', 1, 14),
(191, 12, 1, 11, 'new data after purge', 'PA24-0002-0063', 1, 14),
(192, 12, 1, 13, 'new data after purge', 'PA24-0002-0064', 1, 14),
(193, 13, 1, 5, 'optimization test', 'PC24-0001-0001', 3, 1),
(194, 13, 1, 6, 'new data after purge', 'PC24-0001-0002', 1, 14),
(195, 13, 1, 7, 'new data after purge', 'PC24-0001-0003', 1, 14),
(196, 13, 1, 8, 'new data after purge', 'PC24-0001-0004', 1, 14),
(197, 13, 1, 9, 'new data after purge', 'PC24-0001-0005', 1, 14),
(198, 13, 1, 10, 'new data after purge', 'PC24-0001-0006', 1, 14),
(199, 13, 1, 11, 'new data after purge', 'PC24-0001-0007', 1, 14),
(200, 13, 1, 13, 'new data after purge', 'PC24-0001-0008', 1, 14),
(201, 13, 1, 5, 'new data after purge', 'PC24-0001-0009', 1, 14),
(202, 13, 1, 6, 'new data after purge', 'PC24-0001-0010', 1, 14),
(203, 13, 1, 7, 'new data after purge', 'PC24-0001-0011', 1, 14),
(204, 13, 1, 8, 'new data after purge', 'PC24-0001-0012', 1, 14),
(205, 13, 1, 9, 'new data after purge', 'PC24-0001-0013', 1, 14),
(206, 13, 1, 10, 'new data after purge', 'PC24-0001-0014', 1, 14),
(207, 13, 1, 11, 'new data after purge', 'PC24-0001-0015', 1, 14),
(208, 13, 1, 13, 'new data after purge', 'PC24-0001-0016', 1, 14),
(209, 13, 1, 5, 'new data after purge', 'PC24-0001-0017', 1, 14),
(210, 13, 1, 6, 'new data after purge', 'PC24-0001-0018', 1, 14),
(211, 13, 1, 7, 'new data after purge', 'PC24-0001-0019', 1, 14),
(212, 13, 1, 8, 'new data after purge', 'PC24-0001-0020', 1, 14),
(213, 13, 1, 9, 'new data after purge', 'PC24-0001-0021', 1, 14),
(214, 13, 1, 10, 'new data after purge', 'PC24-0001-0022', 1, 14),
(215, 13, 1, 11, 'new data after purge', 'PC24-0001-0023', 1, 14),
(216, 13, 1, 13, 'new data after purge', 'PC24-0001-0024', 1, 14),
(217, 13, 1, 5, 'new data after purge', 'PC24-0001-0025', 1, 14),
(218, 13, 1, 6, 'new data after purge', 'PC24-0001-0026', 1, 14),
(219, 13, 1, 7, 'new data after purge', 'PC24-0001-0027', 1, 14),
(220, 13, 1, 8, 'new data after purge', 'PC24-0001-0028', 1, 14),
(221, 13, 1, 9, 'new data after purge', 'PC24-0001-0029', 1, 14),
(222, 13, 1, 10, 'new data after purge', 'PC24-0001-0030', 1, 14),
(223, 13, 1, 11, 'new data after purge', 'PC24-0001-0031', 1, 14),
(224, 13, 1, 13, 'new data after purge', 'PC24-0001-0032', 1, 14),
(225, 13, 1, 5, 'new data after purge', 'PC24-0001-0033', 1, 14),
(226, 13, 1, 6, 'new data after purge', 'PC24-0001-0034', 1, 14),
(227, 13, 1, 7, 'new data after purge', 'PC24-0001-0035', 1, 14),
(228, 13, 1, 8, 'new data after purge', 'PC24-0001-0036', 1, 14),
(229, 13, 1, 9, 'new data after purge', 'PC24-0001-0037', 1, 14),
(230, 13, 1, 10, 'new data after purge', 'PC24-0001-0038', 1, 14),
(231, 13, 1, 11, 'new data after purge', 'PC24-0001-0039', 1, 14),
(232, 13, 1, 13, 'new data after purge', 'PC24-0001-0040', 1, 14),
(233, 13, 1, 5, 'new data after purge', 'PC24-0001-0041', 1, 14),
(234, 13, 1, 6, 'new data after purge', 'PC24-0001-0042', 1, 14),
(235, 13, 1, 7, 'new data after purge', 'PC24-0001-0043', 1, 14),
(236, 13, 1, 8, 'new data after purge', 'PC24-0001-0044', 1, 14),
(237, 13, 1, 9, 'new data after purge', 'PC24-0001-0045', 1, 14),
(238, 13, 1, 10, 'new data after purge', 'PC24-0001-0046', 1, 14),
(239, 13, 1, 11, 'new data after purge', 'PC24-0001-0047', 1, 14),
(240, 13, 1, 13, 'new data after purge', 'PC24-0001-0048', 1, 14),
(241, 13, 1, 5, 'new data after purge', 'PC24-0001-0049', 1, 14),
(242, 13, 1, 6, 'new data after purge', 'PC24-0001-0050', 1, 14),
(243, 13, 1, 7, 'new data after purge', 'PC24-0001-0051', 1, 14),
(244, 13, 1, 8, 'new data after purge', 'PC24-0001-0052', 1, 14),
(245, 13, 1, 9, 'new data after purge', 'PC24-0001-0053', 1, 14),
(246, 13, 1, 10, 'new data after purge', 'PC24-0001-0054', 1, 14),
(247, 13, 1, 11, 'new data after purge', 'PC24-0001-0055', 1, 14),
(248, 13, 1, 13, 'new data after purge', 'PC24-0001-0056', 1, 14),
(249, 13, 1, 5, 'new data after purge', 'PC24-0001-0057', 1, 14),
(250, 13, 1, 6, 'new data after purge', 'PC24-0001-0058', 1, 14),
(251, 13, 1, 7, 'new data after purge', 'PC24-0001-0059', 1, 14),
(252, 13, 1, 8, 'new data after purge', 'PC24-0001-0060', 1, 14),
(253, 13, 1, 9, 'new data after purge', 'PC24-0001-0061', 1, 14),
(254, 13, 1, 10, 'new data after purge', 'PC24-0001-0062', 1, 14),
(255, 13, 1, 11, 'new data after purge', 'PC24-0001-0063', 1, 14),
(256, 13, 1, 13, 'new data after purge', 'PC24-0001-0064', 1, 14),
(257, 14, 1, 13, 'test optimization', 'PA24-0003-0001', 2, 1),
(258, 14, 1, 6, 'new data after purge', 'PA24-0003-0002', 1, 14),
(259, 14, 2, 7, 'new data after purge', 'PA24-0003-0003', 1, 1),
(260, 14, 1, 8, 'new data after purge', 'PA24-0003-0004', 1, 14),
(261, 14, 1, 9, 'new data after purge', 'PA24-0003-0005', 1, 14),
(262, 14, 1, 10, 'new data after purge', 'PA24-0003-0006', 1, 14),
(263, 14, 1, 11, 'new data after purge', 'PA24-0003-0007', 1, 14),
(264, 14, 1, 13, 'new data after purge', 'PA24-0003-0008', 1, 14),
(265, 14, 1, 5, 'new data after purge', 'PA24-0003-0009', 1, 14),
(266, 14, 1, 6, 'new data after purge', 'PA24-0003-0010', 1, 14),
(267, 14, 1, 7, 'new data after purge', 'PA24-0003-0011', 1, 14),
(268, 14, 1, 8, 'new data after purge', 'PA24-0003-0012', 1, 14),
(269, 14, 1, 9, 'new data after purge', 'PA24-0003-0013', 1, 14),
(270, 14, 1, 10, 'new data after purge', 'PA24-0003-0014', 1, 14),
(271, 14, 1, 11, 'new data after purge', 'PA24-0003-0015', 1, 14),
(272, 14, 1, 13, 'new data after purge', 'PA24-0003-0016', 1, 14),
(273, 14, 1, 5, 'new data after purge', 'PA24-0003-0017', 1, 14),
(274, 14, 1, 6, 'new data after purge', 'PA24-0003-0018', 1, 14),
(275, 14, 1, 7, 'new data after purge', 'PA24-0003-0019', 1, 14),
(276, 14, 1, 8, 'new data after purge', 'PA24-0003-0020', 1, 14),
(277, 14, 1, 9, 'new data after purge', 'PA24-0003-0021', 1, 14),
(278, 14, 1, 10, 'new data after purge', 'PA24-0003-0022', 1, 14),
(279, 14, 1, 11, 'new data after purge', 'PA24-0003-0023', 1, 14),
(280, 14, 1, 13, 'new data after purge', 'PA24-0003-0024', 1, 14),
(281, 14, 1, 5, 'new data after purge', 'PA24-0003-0025', 1, 14),
(282, 14, 1, 6, 'new data after purge', 'PA24-0003-0026', 1, 14),
(283, 14, 1, 7, 'new data after purge', 'PA24-0003-0027', 1, 14),
(284, 14, 1, 8, 'new data after purge', 'PA24-0003-0028', 1, 14),
(285, 14, 1, 9, 'new data after purge', 'PA24-0003-0029', 1, 14),
(286, 14, 1, 10, 'new data after purge', 'PA24-0003-0030', 1, 14),
(287, 14, 1, 11, 'new data after purge', 'PA24-0003-0031', 1, 14),
(288, 14, 1, 13, 'new data after purge', 'PA24-0003-0032', 1, 14),
(289, 14, 1, 5, 'new data after purge', 'PA24-0003-0033', 1, 14),
(290, 14, 1, 6, 'new data after purge', 'PA24-0003-0034', 1, 14),
(291, 14, 1, 7, 'new data after purge', 'PA24-0003-0035', 1, 14),
(292, 14, 1, 8, 'new data after purge', 'PA24-0003-0036', 1, 14),
(293, 14, 1, 9, 'new data after purge', 'PA24-0003-0037', 1, 14),
(294, 14, 1, 10, 'new data after purge', 'PA24-0003-0038', 1, 14),
(295, 14, 1, 11, 'new data after purge', 'PA24-0003-0039', 1, 14),
(296, 14, 1, 13, 'new data after purge', 'PA24-0003-0040', 1, 14),
(297, 14, 1, 5, 'new data after purge', 'PA24-0003-0041', 1, 14),
(298, 14, 1, 6, 'new data after purge', 'PA24-0003-0042', 1, 14),
(299, 14, 1, 7, 'new data after purge', 'PA24-0003-0043', 1, 14),
(300, 14, 1, 8, 'new data after purge', 'PA24-0003-0044', 1, 14),
(301, 14, 1, 9, 'new data after purge', 'PA24-0003-0045', 1, 14),
(302, 14, 1, 10, 'new data after purge', 'PA24-0003-0046', 1, 14),
(303, 14, 1, 11, 'new data after purge', 'PA24-0003-0047', 1, 14),
(304, 14, 1, 13, 'new data after purge', 'PA24-0003-0048', 1, 14),
(305, 14, 1, 5, 'new data after purge', 'PA24-0003-0049', 1, 14),
(306, 14, 1, 6, 'new data after purge', 'PA24-0003-0050', 1, 14),
(307, 14, 1, 7, 'new data after purge', 'PA24-0003-0051', 1, 14),
(308, 14, 1, 8, 'new data after purge', 'PA24-0003-0052', 1, 14),
(309, 14, 1, 9, 'new data after purge', 'PA24-0003-0053', 1, 14),
(310, 14, 1, 10, 'new data after purge', 'PA24-0003-0054', 1, 14),
(311, 14, 1, 11, 'new data after purge', 'PA24-0003-0055', 1, 14),
(312, 14, 1, 13, 'new data after purge', 'PA24-0003-0056', 1, 14),
(313, 14, 1, 5, 'new data after purge', 'PA24-0003-0057', 1, 14),
(314, 14, 1, 6, 'new data after purge', 'PA24-0003-0058', 1, 14),
(315, 14, 1, 7, 'new data after purge', 'PA24-0003-0059', 1, 14),
(316, 14, 1, 8, 'new data after purge', 'PA24-0003-0060', 1, 14),
(317, 14, 1, 9, 'new data after purge', 'PA24-0003-0061', 1, 14),
(318, 14, 1, 10, 'new data after purge', 'PA24-0003-0062', 1, 14),
(319, 14, 1, 11, 'new data after purge', 'PA24-0003-0063', 1, 14),
(320, 14, 1, 13, 'new data after purge', 'PA24-0003-0064', 1, 14),
(321, 15, 1, 13, 'log check', 'VD24-0001-0001', 2, 1),
(322, 15, 1, 6, 'new data after purge', 'VD24-0001-0002', 1, 14),
(323, 15, 1, 7, 'new data after purge', 'VD24-0001-0003', 1, 14),
(324, 15, 1, 8, 'new data after purge', 'VD24-0001-0004', 1, 14),
(325, 15, 1, 9, 'new data after purge', 'VD24-0001-0005', 1, 14),
(326, 15, 1, 10, 'new data after purge', 'VD24-0001-0006', 1, 14),
(327, 15, 1, 11, 'new data after purge', 'VD24-0001-0007', 1, 14),
(328, 15, 1, 13, 'new data after purge', 'VD24-0001-0008', 1, 14),
(329, 15, 1, 5, 'new data after purge', 'VD24-0001-0009', 1, 14),
(330, 15, 1, 6, 'new data after purge', 'VD24-0001-0010', 1, 14),
(331, 15, 1, 7, 'new data after purge', 'VD24-0001-0011', 1, 14),
(332, 15, 1, 8, 'new data after purge', 'VD24-0001-0012', 1, 14),
(333, 15, 1, 9, 'new data after purge', 'VD24-0001-0013', 1, 14),
(334, 15, 1, 10, 'new data after purge', 'VD24-0001-0014', 1, 14),
(335, 15, 1, 11, 'new data after purge', 'VD24-0001-0015', 1, 14),
(336, 15, 1, 13, 'new data after purge', 'VD24-0001-0016', 1, 14),
(337, 15, 1, 5, 'new data after purge', 'VD24-0001-0017', 1, 14),
(338, 15, 1, 6, 'new data after purge', 'VD24-0001-0018', 1, 14),
(339, 15, 1, 7, 'new data after purge', 'VD24-0001-0019', 1, 14),
(340, 15, 1, 8, 'new data after purge', 'VD24-0001-0020', 1, 14),
(341, 15, 1, 9, 'new data after purge', 'VD24-0001-0021', 1, 14),
(342, 15, 1, 10, 'new data after purge', 'VD24-0001-0022', 1, 14),
(343, 15, 1, 11, 'new data after purge', 'VD24-0001-0023', 1, 14),
(344, 15, 1, 13, 'new data after purge', 'VD24-0001-0024', 1, 14),
(345, 15, 1, 5, 'new data after purge', 'VD24-0001-0025', 1, 14),
(346, 15, 1, 6, 'new data after purge', 'VD24-0001-0026', 1, 14),
(347, 15, 1, 7, 'new data after purge', 'VD24-0001-0027', 1, 14),
(348, 15, 1, 8, 'new data after purge', 'VD24-0001-0028', 1, 14),
(349, 15, 1, 9, 'new data after purge', 'VD24-0001-0029', 1, 14),
(350, 15, 1, 10, 'new data after purge', 'VD24-0001-0030', 1, 14),
(351, 15, 1, 11, 'new data after purge', 'VD24-0001-0031', 1, 14),
(352, 15, 1, 13, 'new data after purge', 'VD24-0001-0032', 1, 14),
(353, 15, 1, 5, 'new data after purge', 'VD24-0001-0033', 1, 14),
(354, 15, 1, 6, 'new data after purge', 'VD24-0001-0034', 1, 14),
(355, 15, 1, 7, 'new data after purge', 'VD24-0001-0035', 1, 14),
(356, 15, 1, 8, 'new data after purge', 'VD24-0001-0036', 1, 14),
(357, 15, 1, 9, 'new data after purge', 'VD24-0001-0037', 1, 14),
(358, 15, 1, 10, 'new data after purge', 'VD24-0001-0038', 1, 14),
(359, 15, 1, 11, 'new data after purge', 'VD24-0001-0039', 1, 14),
(360, 15, 1, 13, 'new data after purge', 'VD24-0001-0040', 1, 14),
(361, 15, 1, 5, 'new data after purge', 'VD24-0001-0041', 1, 14),
(362, 15, 1, 6, 'new data after purge', 'VD24-0001-0042', 1, 14),
(363, 15, 1, 7, 'new data after purge', 'VD24-0001-0043', 1, 14),
(364, 15, 1, 8, 'new data after purge', 'VD24-0001-0044', 1, 14),
(365, 15, 1, 9, 'new data after purge', 'VD24-0001-0045', 1, 14),
(366, 15, 1, 10, 'new data after purge', 'VD24-0001-0046', 1, 14),
(367, 15, 1, 11, 'new data after purge', 'VD24-0001-0047', 1, 14),
(368, 15, 1, 13, 'new data after purge', 'VD24-0001-0048', 1, 14),
(369, 15, 1, 5, 'new data after purge', 'VD24-0001-0049', 1, 14),
(370, 15, 1, 6, 'new data after purge', 'VD24-0001-0050', 1, 14),
(371, 15, 1, 7, 'new data after purge', 'VD24-0001-0051', 1, 14),
(372, 15, 1, 8, 'new data after purge', 'VD24-0001-0052', 1, 14),
(373, 15, 1, 9, 'new data after purge', 'VD24-0001-0053', 1, 14),
(374, 15, 1, 10, 'new data after purge', 'VD24-0001-0054', 1, 14),
(375, 15, 1, 11, 'new data after purge', 'VD24-0001-0055', 1, 14),
(376, 15, 1, 13, 'new data after purge', 'VD24-0001-0056', 1, 14),
(377, 15, 1, 5, 'new data after purge', 'VD24-0001-0057', 1, 14),
(378, 15, 1, 6, 'new data after purge', 'VD24-0001-0058', 1, 14),
(379, 15, 1, 7, 'new data after purge', 'VD24-0001-0059', 1, 14),
(380, 15, 1, 8, 'new data after purge', 'VD24-0001-0060', 1, 14),
(381, 15, 1, 9, 'new data after purge', 'VD24-0001-0061', 1, 14),
(382, 15, 1, 10, 'new data after purge', 'VD24-0001-0062', 1, 14),
(383, 15, 1, 11, 'new data after purge', 'VD24-0001-0063', 1, 14),
(384, 15, 1, 13, 'new data after purge', 'VD24-0001-0064', 1, 14),
(385, 16, 1, 13, 'optimization test', 'VD24-0002-0001', 2, 1),
(386, 16, 1, 6, 'new data after purge', 'VD24-0002-0002', 1, 14),
(387, 16, 1, 7, 'new data after purge', 'VD24-0002-0003', 1, 14),
(388, 16, 1, 8, 'new data after purge', 'VD24-0002-0004', 1, 14),
(389, 16, 1, 9, 'new data after purge', 'VD24-0002-0005', 1, 14),
(390, 16, 1, 10, 'new data after purge', 'VD24-0002-0006', 1, 14),
(391, 16, 1, 11, 'new data after purge', 'VD24-0002-0007', 1, 14),
(392, 16, 1, 13, 'new data after purge', 'VD24-0002-0008', 1, 14),
(393, 16, 1, 5, 'new data after purge', 'VD24-0002-0009', 1, 14),
(394, 16, 1, 6, 'new data after purge', 'VD24-0002-0010', 1, 14),
(395, 16, 1, 7, 'new data after purge', 'VD24-0002-0011', 1, 14),
(396, 16, 1, 8, 'new data after purge', 'VD24-0002-0012', 1, 14),
(397, 16, 1, 9, 'new data after purge', 'VD24-0002-0013', 1, 14),
(398, 16, 1, 10, 'new data after purge', 'VD24-0002-0014', 1, 14),
(399, 16, 1, 11, 'new data after purge', 'VD24-0002-0015', 1, 14),
(400, 16, 1, 13, 'new data after purge', 'VD24-0002-0016', 1, 14),
(401, 16, 1, 5, 'new data after purge', 'VD24-0002-0017', 1, 14),
(402, 16, 1, 6, 'new data after purge', 'VD24-0002-0018', 1, 14),
(403, 16, 1, 7, 'new data after purge', 'VD24-0002-0019', 1, 14),
(404, 16, 1, 8, 'new data after purge', 'VD24-0002-0020', 1, 14),
(405, 16, 1, 9, 'new data after purge', 'VD24-0002-0021', 1, 14),
(406, 16, 1, 10, 'new data after purge', 'VD24-0002-0022', 1, 14),
(407, 16, 1, 11, 'new data after purge', 'VD24-0002-0023', 1, 14),
(408, 16, 1, 13, 'new data after purge', 'VD24-0002-0024', 1, 14),
(409, 16, 1, 5, 'new data after purge', 'VD24-0002-0025', 1, 14),
(410, 16, 1, 6, 'new data after purge', 'VD24-0002-0026', 1, 14),
(411, 16, 1, 7, 'new data after purge', 'VD24-0002-0027', 1, 14),
(412, 16, 1, 8, 'new data after purge', 'VD24-0002-0028', 1, 14),
(413, 16, 1, 9, 'new data after purge', 'VD24-0002-0029', 1, 14),
(414, 16, 1, 10, 'new data after purge', 'VD24-0002-0030', 1, 14),
(415, 16, 1, 11, 'new data after purge', 'VD24-0002-0031', 1, 14),
(416, 16, 1, 13, 'new data after purge', 'VD24-0002-0032', 1, 14),
(417, 16, 1, 5, 'new data after purge', 'VD24-0002-0033', 1, 14),
(418, 16, 1, 6, 'new data after purge', 'VD24-0002-0034', 1, 14),
(419, 16, 1, 7, 'new data after purge', 'VD24-0002-0035', 1, 14),
(420, 16, 1, 8, 'new data after purge', 'VD24-0002-0036', 1, 14),
(421, 16, 1, 9, 'new data after purge', 'VD24-0002-0037', 1, 14),
(422, 16, 1, 10, 'new data after purge', 'VD24-0002-0038', 1, 14),
(423, 16, 1, 11, 'new data after purge', 'VD24-0002-0039', 1, 14),
(424, 16, 1, 13, 'new data after purge', 'VD24-0002-0040', 1, 14),
(425, 16, 1, 5, 'new data after purge', 'VD24-0002-0041', 1, 14),
(426, 16, 1, 6, 'new data after purge', 'VD24-0002-0042', 1, 14),
(427, 16, 1, 7, 'new data after purge', 'VD24-0002-0043', 1, 14),
(428, 16, 1, 8, 'new data after purge', 'VD24-0002-0044', 1, 14),
(429, 16, 1, 9, 'new data after purge', 'VD24-0002-0045', 1, 14),
(430, 16, 1, 10, 'new data after purge', 'VD24-0002-0046', 1, 14),
(431, 16, 1, 11, 'new data after purge', 'VD24-0002-0047', 1, 14),
(432, 16, 1, 13, 'new data after purge', 'VD24-0002-0048', 1, 14),
(433, 16, 1, 5, 'new data after purge', 'VD24-0002-0049', 1, 14),
(434, 16, 1, 6, 'new data after purge', 'VD24-0002-0050', 1, 14),
(435, 16, 1, 7, 'new data after purge', 'VD24-0002-0051', 1, 14),
(436, 16, 1, 8, 'new data after purge', 'VD24-0002-0052', 1, 14),
(437, 16, 1, 9, 'new data after purge', 'VD24-0002-0053', 1, 14),
(438, 16, 1, 10, 'new data after purge', 'VD24-0002-0054', 1, 14),
(439, 16, 1, 11, 'new data after purge', 'VD24-0002-0055', 1, 14),
(440, 16, 1, 13, 'new data after purge', 'VD24-0002-0056', 1, 14),
(441, 16, 1, 5, 'new data after purge', 'VD24-0002-0057', 1, 14),
(442, 16, 1, 6, 'new data after purge', 'VD24-0002-0058', 1, 14),
(443, 16, 1, 7, 'new data after purge', 'VD24-0002-0059', 1, 14),
(444, 16, 1, 8, 'new data after purge', 'VD24-0002-0060', 1, 14),
(445, 16, 1, 9, 'new data after purge', 'VD24-0002-0061', 1, 14),
(446, 16, 1, 10, 'new data after purge', 'VD24-0002-0062', 1, 14),
(447, 16, 1, 11, 'new data after purge', 'VD24-0002-0063', 1, 14),
(448, 16, 1, 13, 'new data after purge', 'VD24-0002-0064', 1, 14),
(449, 18, 1, 5, 'new data after purge', 'VD24-0003-0001', 1, 14),
(450, 18, 1, 7, 'SPAM PREVENTION TEST', 'VD24-0003-0002', 1, 1),
(451, 18, 1, 7, 'new data after purge', 'VD24-0003-0003', 1, 14),
(452, 18, 1, 8, 'new data after purge', 'VD24-0003-0004', 1, 14),
(453, 18, 1, 9, 'new data after purge', 'VD24-0003-0005', 1, 14),
(454, 18, 1, 10, 'new data after purge', 'VD24-0003-0006', 1, 14),
(455, 18, 1, 11, 'new data after purge', 'VD24-0003-0007', 1, 14),
(456, 18, 1, 13, 'new data after purge', 'VD24-0003-0008', 1, 14),
(457, 18, 1, 5, 'new data after purge', 'VD24-0003-0009', 1, 14),
(458, 18, 1, 6, 'new data after purge', 'VD24-0003-0010', 1, 14),
(459, 18, 1, 7, 'new data after purge', 'VD24-0003-0011', 1, 14),
(460, 18, 1, 8, 'new data after purge', 'VD24-0003-0012', 1, 14),
(461, 18, 1, 9, 'new data after purge', 'VD24-0003-0013', 1, 14),
(462, 18, 1, 10, 'new data after purge', 'VD24-0003-0014', 1, 14),
(463, 18, 1, 11, 'new data after purge', 'VD24-0003-0015', 1, 14),
(464, 18, 1, 13, 'new data after purge', 'VD24-0003-0016', 1, 14),
(465, 18, 1, 5, 'new data after purge', 'VD24-0003-0017', 1, 14),
(466, 18, 1, 6, 'new data after purge', 'VD24-0003-0018', 1, 14),
(467, 18, 1, 7, 'new data after purge', 'VD24-0003-0019', 1, 14),
(468, 18, 1, 8, 'new data after purge', 'VD24-0003-0020', 1, 14),
(469, 18, 1, 9, 'new data after purge', 'VD24-0003-0021', 1, 14),
(470, 18, 1, 10, 'new data after purge', 'VD24-0003-0022', 1, 14),
(471, 18, 1, 11, 'new data after purge', 'VD24-0003-0023', 1, 14),
(472, 18, 1, 13, 'new data after purge', 'VD24-0003-0024', 1, 14),
(473, 18, 1, 5, 'new data after purge', 'VD24-0003-0025', 1, 14),
(474, 18, 1, 6, 'new data after purge', 'VD24-0003-0026', 1, 14),
(475, 18, 1, 7, 'new data after purge', 'VD24-0003-0027', 1, 14),
(476, 18, 1, 8, 'new data after purge', 'VD24-0003-0028', 1, 14),
(477, 18, 1, 9, 'new data after purge', 'VD24-0003-0029', 1, 14),
(478, 18, 1, 10, 'new data after purge', 'VD24-0003-0030', 1, 14),
(479, 18, 1, 11, 'new data after purge', 'VD24-0003-0031', 1, 14),
(480, 18, 1, 13, 'new data after purge', 'VD24-0003-0032', 1, 14),
(481, 18, 1, 5, 'new data after purge', 'VD24-0003-0033', 1, 14),
(482, 18, 1, 6, 'new data after purge', 'VD24-0003-0034', 1, 14),
(483, 18, 1, 7, 'new data after purge', 'VD24-0003-0035', 1, 14),
(484, 18, 1, 8, 'new data after purge', 'VD24-0003-0036', 1, 14),
(485, 18, 1, 9, 'new data after purge', 'VD24-0003-0037', 1, 14),
(486, 18, 1, 10, 'new data after purge', 'VD24-0003-0038', 1, 14),
(487, 18, 1, 11, 'new data after purge', 'VD24-0003-0039', 1, 14),
(488, 18, 1, 13, 'new data after purge', 'VD24-0003-0040', 1, 14),
(489, 18, 1, 5, 'new data after purge', 'VD24-0003-0041', 1, 14),
(490, 18, 1, 6, 'new data after purge', 'VD24-0003-0042', 1, 14),
(491, 18, 1, 7, 'new data after purge', 'VD24-0003-0043', 1, 14),
(492, 18, 1, 8, 'new data after purge', 'VD24-0003-0044', 1, 14),
(493, 18, 1, 9, 'new data after purge', 'VD24-0003-0045', 1, 14),
(494, 18, 1, 10, 'new data after purge', 'VD24-0003-0046', 1, 14),
(495, 18, 1, 11, 'new data after purge', 'VD24-0003-0047', 1, 14),
(496, 18, 1, 13, 'new data after purge', 'VD24-0003-0048', 1, 14),
(497, 18, 1, 5, 'new data after purge', 'VD24-0003-0049', 1, 14),
(498, 18, 1, 6, 'new data after purge', 'VD24-0003-0050', 1, 14),
(499, 18, 1, 7, 'new data after purge', 'VD24-0003-0051', 1, 14),
(500, 18, 1, 8, 'new data after purge', 'VD24-0003-0052', 1, 14),
(501, 18, 1, 9, 'new data after purge', 'VD24-0003-0053', 1, 14),
(502, 18, 1, 10, 'new data after purge', 'VD24-0003-0054', 1, 14),
(503, 18, 1, 11, 'new data after purge', 'VD24-0003-0055', 1, 14),
(504, 18, 1, 13, 'new data after purge', 'VD24-0003-0056', 1, 14),
(505, 18, 1, 5, 'new data after purge', 'VD24-0003-0057', 1, 14),
(506, 18, 1, 6, 'new data after purge', 'VD24-0003-0058', 1, 14),
(507, 18, 1, 7, 'new data after purge', 'VD24-0003-0059', 1, 14),
(508, 18, 1, 8, 'new data after purge', 'VD24-0003-0060', 1, 14),
(509, 18, 1, 9, 'new data after purge', 'VD24-0003-0061', 1, 14),
(510, 18, 1, 10, 'new data after purge', 'VD24-0003-0062', 1, 14),
(511, 18, 1, 11, 'new data after purge', 'VD24-0003-0063', 1, 14),
(512, 18, 1, 13, 'new data after purge', 'VD24-0003-0064', 1, 14),
(513, 25, 1, 13, 'repairlog optimization test', 'PA24-0004-0001', 2, 1),
(514, 25, 1, 6, 'new data after purge', 'PA24-0004-0002', 1, 14),
(515, 25, 1, 7, 'new data after purge', 'PA24-0004-0003', 1, 14),
(516, 25, 1, 8, 'new data after purge', 'PA24-0004-0004', 1, 14),
(517, 25, 1, 9, 'new data after purge', 'PA24-0004-0005', 1, 14),
(518, 25, 1, 10, 'new data after purge', 'PA24-0004-0006', 1, 14),
(519, 25, 1, 11, 'new data after purge', 'PA24-0004-0007', 1, 14),
(520, 25, 1, 13, 'new data after purge', 'PA24-0004-0008', 1, 14),
(521, 25, 1, 5, 'new data after purge', 'PA24-0004-0009', 1, 14),
(522, 25, 1, 6, 'new data after purge', 'PA24-0004-0010', 1, 14),
(523, 25, 1, 7, 'new data after purge', 'PA24-0004-0011', 1, 14),
(524, 25, 1, 8, 'new data after purge', 'PA24-0004-0012', 1, 14),
(525, 25, 1, 9, 'new data after purge', 'PA24-0004-0013', 1, 14),
(526, 25, 1, 10, 'new data after purge', 'PA24-0004-0014', 1, 14),
(527, 25, 1, 11, 'new data after purge', 'PA24-0004-0015', 1, 14),
(528, 25, 1, 13, 'new data after purge', 'PA24-0004-0016', 1, 14),
(529, 25, 1, 5, 'new data after purge', 'PA24-0004-0017', 1, 14),
(530, 25, 1, 6, 'new data after purge', 'PA24-0004-0018', 1, 14),
(531, 25, 1, 7, 'new data after purge', 'PA24-0004-0019', 1, 14),
(532, 25, 1, 8, 'new data after purge', 'PA24-0004-0020', 1, 14),
(533, 25, 1, 9, 'new data after purge', 'PA24-0004-0021', 1, 14),
(534, 25, 1, 10, 'new data after purge', 'PA24-0004-0022', 1, 14),
(535, 25, 1, 11, 'new data after purge', 'PA24-0004-0023', 1, 14),
(536, 25, 1, 13, 'new data after purge', 'PA24-0004-0024', 1, 14),
(537, 25, 1, 5, 'new data after purge', 'PA24-0004-0025', 1, 14),
(538, 25, 1, 6, 'new data after purge', 'PA24-0004-0026', 1, 14),
(539, 25, 1, 7, 'new data after purge', 'PA24-0004-0027', 1, 14),
(540, 25, 1, 8, 'new data after purge', 'PA24-0004-0028', 1, 14),
(541, 25, 1, 9, 'new data after purge', 'PA24-0004-0029', 1, 14),
(542, 25, 1, 10, 'new data after purge', 'PA24-0004-0030', 1, 14),
(543, 25, 1, 11, 'new data after purge', 'PA24-0004-0031', 1, 14),
(544, 25, 1, 13, 'new data after purge', 'PA24-0004-0032', 1, 14),
(545, 25, 1, 5, 'new data after purge', 'PA24-0004-0033', 1, 14),
(546, 25, 1, 6, 'new data after purge', 'PA24-0004-0034', 1, 14),
(547, 25, 1, 7, 'new data after purge', 'PA24-0004-0035', 1, 14),
(548, 25, 1, 8, 'new data after purge', 'PA24-0004-0036', 1, 14),
(549, 25, 1, 9, 'new data after purge', 'PA24-0004-0037', 1, 14),
(550, 25, 1, 10, 'new data after purge', 'PA24-0004-0038', 1, 14),
(551, 25, 1, 11, 'new data after purge', 'PA24-0004-0039', 1, 14),
(552, 25, 1, 13, 'new data after purge', 'PA24-0004-0040', 1, 14),
(553, 25, 1, 5, 'new data after purge', 'PA24-0004-0041', 1, 14),
(554, 25, 1, 6, 'new data after purge', 'PA24-0004-0042', 1, 14),
(555, 25, 1, 7, 'new data after purge', 'PA24-0004-0043', 1, 14),
(556, 25, 1, 8, 'new data after purge', 'PA24-0004-0044', 1, 14),
(557, 25, 1, 9, 'new data after purge', 'PA24-0004-0045', 1, 14),
(558, 25, 1, 10, 'new data after purge', 'PA24-0004-0046', 1, 14),
(559, 25, 1, 11, 'new data after purge', 'PA24-0004-0047', 1, 14),
(560, 25, 1, 13, 'new data after purge', 'PA24-0004-0048', 1, 14),
(561, 25, 1, 5, 'new data after purge', 'PA24-0004-0049', 1, 14),
(562, 25, 1, 6, 'new data after purge', 'PA24-0004-0050', 1, 14),
(563, 25, 1, 7, 'new data after purge', 'PA24-0004-0051', 1, 14),
(564, 25, 1, 8, 'new data after purge', 'PA24-0004-0052', 1, 14),
(565, 25, 1, 9, 'new data after purge', 'PA24-0004-0053', 1, 14),
(566, 25, 1, 10, 'new data after purge', 'PA24-0004-0054', 1, 14),
(567, 25, 1, 11, 'new data after purge', 'PA24-0004-0055', 1, 14),
(568, 25, 1, 13, 'new data after purge', 'PA24-0004-0056', 1, 14),
(569, 25, 1, 5, 'new data after purge', 'PA24-0004-0057', 1, 14),
(570, 25, 1, 6, 'new data after purge', 'PA24-0004-0058', 1, 14),
(571, 25, 1, 7, 'new data after purge', 'PA24-0004-0059', 1, 14),
(572, 25, 1, 8, 'new data after purge', 'PA24-0004-0060', 1, 14),
(573, 25, 1, 9, 'new data after purge', 'PA24-0004-0061', 1, 14),
(574, 25, 1, 10, 'new data after purge', 'PA24-0004-0062', 1, 14),
(575, 25, 1, 11, 'new data after purge', 'PA24-0004-0063', 1, 14),
(576, 25, 1, 13, 'new data after purge', 'PA24-0004-0064', 1, 14),
(577, 26, 1, 5, 'optimization test', 'PC24-0002-0001', 1, 1),
(578, 26, 1, 6, 'new data after purge', 'PC24-0002-0002', 1, 14),
(579, 26, 1, 7, 'new data after purge', 'PC24-0002-0003', 1, 14),
(580, 26, 1, 8, 'new data after purge', 'PC24-0002-0004', 1, 14),
(581, 26, 1, 9, 'new data after purge', 'PC24-0002-0005', 1, 14),
(582, 26, 1, 10, 'new data after purge', 'PC24-0002-0006', 1, 14),
(583, 26, 1, 11, 'new data after purge', 'PC24-0002-0007', 1, 14),
(584, 26, 1, 13, 'new data after purge', 'PC24-0002-0008', 1, 14),
(585, 26, 1, 5, 'new data after purge', 'PC24-0002-0009', 1, 14),
(586, 26, 1, 6, 'new data after purge', 'PC24-0002-0010', 1, 14),
(587, 26, 1, 7, 'new data after purge', 'PC24-0002-0011', 1, 14),
(588, 26, 1, 8, 'new data after purge', 'PC24-0002-0012', 1, 14),
(589, 26, 1, 9, 'new data after purge', 'PC24-0002-0013', 1, 14),
(590, 26, 1, 10, 'new data after purge', 'PC24-0002-0014', 1, 14),
(591, 26, 1, 11, 'new data after purge', 'PC24-0002-0015', 1, 14),
(592, 26, 1, 13, 'new data after purge', 'PC24-0002-0016', 1, 14),
(593, 26, 1, 5, 'new data after purge', 'PC24-0002-0017', 1, 14),
(594, 26, 1, 6, 'new data after purge', 'PC24-0002-0018', 1, 14),
(595, 26, 1, 7, 'new data after purge', 'PC24-0002-0019', 1, 14),
(596, 26, 1, 8, 'new data after purge', 'PC24-0002-0020', 1, 14),
(597, 26, 1, 9, 'new data after purge', 'PC24-0002-0021', 1, 14),
(598, 26, 1, 10, 'new data after purge', 'PC24-0002-0022', 1, 14),
(599, 26, 1, 11, 'new data after purge', 'PC24-0002-0023', 1, 14),
(600, 26, 1, 13, 'new data after purge', 'PC24-0002-0024', 1, 14),
(601, 26, 1, 5, 'new data after purge', 'PC24-0002-0025', 1, 14),
(602, 26, 1, 6, 'new data after purge', 'PC24-0002-0026', 1, 14),
(603, 26, 1, 7, 'new data after purge', 'PC24-0002-0027', 1, 14),
(604, 26, 1, 8, 'new data after purge', 'PC24-0002-0028', 1, 14),
(605, 26, 1, 9, 'new data after purge', 'PC24-0002-0029', 1, 14),
(606, 26, 1, 10, 'new data after purge', 'PC24-0002-0030', 1, 14),
(607, 26, 1, 11, 'new data after purge', 'PC24-0002-0031', 1, 14),
(608, 26, 1, 13, 'new data after purge', 'PC24-0002-0032', 1, 14),
(609, 26, 1, 5, 'new data after purge', 'PC24-0002-0033', 1, 14),
(610, 26, 1, 6, 'new data after purge', 'PC24-0002-0034', 1, 14),
(611, 26, 1, 7, 'new data after purge', 'PC24-0002-0035', 1, 14),
(612, 26, 1, 8, 'new data after purge', 'PC24-0002-0036', 1, 14),
(613, 26, 1, 9, 'new data after purge', 'PC24-0002-0037', 1, 14),
(614, 26, 1, 10, 'new data after purge', 'PC24-0002-0038', 1, 14),
(615, 26, 1, 11, 'new data after purge', 'PC24-0002-0039', 1, 14),
(616, 26, 1, 13, 'new data after purge', 'PC24-0002-0040', 1, 14),
(617, 26, 1, 5, 'new data after purge', 'PC24-0002-0041', 1, 14),
(618, 26, 1, 6, 'new data after purge', 'PC24-0002-0042', 1, 14),
(619, 26, 1, 7, 'new data after purge', 'PC24-0002-0043', 1, 14),
(620, 26, 1, 8, 'new data after purge', 'PC24-0002-0044', 1, 14),
(621, 26, 1, 9, 'new data after purge', 'PC24-0002-0045', 1, 14),
(622, 26, 1, 10, 'new data after purge', 'PC24-0002-0046', 1, 14),
(623, 26, 1, 11, 'new data after purge', 'PC24-0002-0047', 1, 14),
(624, 26, 1, 13, 'new data after purge', 'PC24-0002-0048', 1, 14),
(625, 26, 1, 5, 'new data after purge', 'PC24-0002-0049', 1, 14),
(626, 26, 1, 6, 'new data after purge', 'PC24-0002-0050', 1, 14),
(627, 26, 1, 7, 'new data after purge', 'PC24-0002-0051', 1, 14),
(628, 26, 1, 8, 'new data after purge', 'PC24-0002-0052', 1, 14),
(629, 26, 1, 9, 'new data after purge', 'PC24-0002-0053', 1, 14),
(630, 26, 1, 10, 'new data after purge', 'PC24-0002-0054', 1, 14),
(631, 26, 1, 11, 'new data after purge', 'PC24-0002-0055', 1, 14),
(632, 26, 1, 13, 'new data after purge', 'PC24-0002-0056', 1, 14),
(633, 26, 1, 5, 'new data after purge', 'PC24-0002-0057', 1, 14),
(634, 26, 1, 6, 'new data after purge', 'PC24-0002-0058', 1, 14),
(635, 26, 1, 7, 'new data after purge', 'PC24-0002-0059', 1, 14),
(636, 26, 1, 8, 'new data after purge', 'PC24-0002-0060', 1, 14),
(637, 26, 1, 9, 'new data after purge', 'PC24-0002-0061', 1, 14),
(638, 26, 1, 10, 'new data after purge', 'PC24-0002-0062', 1, 14),
(639, 26, 1, 11, 'new data after purge', 'PC24-0002-0063', 1, 14),
(640, 26, 1, 13, 'new data after purge', 'PC24-0002-0064', 1, 14),
(641, 27, 1, 5, 'optimization test', 'VD24-0004-0001', 1, 1),
(642, 27, 1, 6, 'new data after purge', 'VD24-0004-0002', 1, 14),
(643, 27, 1, 7, 'new data after purge', 'VD24-0004-0003', 1, 14),
(644, 27, 1, 8, 'new data after purge', 'VD24-0004-0004', 1, 14),
(645, 27, 1, 9, 'new data after purge', 'VD24-0004-0005', 1, 14),
(646, 27, 1, 10, 'new data after purge', 'VD24-0004-0006', 1, 14),
(647, 27, 1, 11, 'new data after purge', 'VD24-0004-0007', 1, 14),
(648, 27, 1, 13, 'new data after purge', 'VD24-0004-0008', 1, 14),
(649, 27, 1, 5, 'new data after purge', 'VD24-0004-0009', 1, 14),
(650, 27, 1, 6, 'new data after purge', 'VD24-0004-0010', 1, 14),
(651, 27, 1, 7, 'new data after purge', 'VD24-0004-0011', 1, 14),
(652, 27, 1, 8, 'new data after purge', 'VD24-0004-0012', 1, 14),
(653, 27, 1, 9, 'new data after purge', 'VD24-0004-0013', 1, 14),
(654, 27, 1, 10, 'new data after purge', 'VD24-0004-0014', 1, 14),
(655, 27, 1, 11, 'new data after purge', 'VD24-0004-0015', 1, 14),
(656, 27, 1, 13, 'new data after purge', 'VD24-0004-0016', 1, 14),
(657, 27, 1, 5, 'new data after purge', 'VD24-0004-0017', 1, 14),
(658, 27, 1, 6, 'new data after purge', 'VD24-0004-0018', 1, 14),
(659, 27, 1, 7, 'new data after purge', 'VD24-0004-0019', 1, 14),
(660, 27, 1, 8, 'new data after purge', 'VD24-0004-0020', 1, 14),
(661, 27, 1, 9, 'new data after purge', 'VD24-0004-0021', 1, 14),
(662, 27, 1, 10, 'new data after purge', 'VD24-0004-0022', 1, 14),
(663, 27, 1, 11, 'new data after purge', 'VD24-0004-0023', 1, 14),
(664, 27, 1, 13, 'new data after purge', 'VD24-0004-0024', 1, 14),
(665, 27, 1, 5, 'new data after purge', 'VD24-0004-0025', 1, 14),
(666, 27, 1, 6, 'new data after purge', 'VD24-0004-0026', 1, 14),
(667, 27, 1, 7, 'new data after purge', 'VD24-0004-0027', 1, 14),
(668, 27, 1, 8, 'new data after purge', 'VD24-0004-0028', 1, 14),
(669, 27, 1, 9, 'new data after purge', 'VD24-0004-0029', 1, 14),
(670, 27, 1, 10, 'new data after purge', 'VD24-0004-0030', 1, 14),
(671, 27, 1, 11, 'new data after purge', 'VD24-0004-0031', 1, 14),
(672, 27, 1, 13, 'new data after purge', 'VD24-0004-0032', 1, 14),
(673, 27, 1, 5, 'new data after purge', 'VD24-0004-0033', 1, 14),
(674, 27, 1, 6, 'new data after purge', 'VD24-0004-0034', 1, 14),
(675, 27, 1, 7, 'new data after purge', 'VD24-0004-0035', 1, 14),
(676, 27, 1, 8, 'new data after purge', 'VD24-0004-0036', 1, 14),
(677, 27, 1, 9, 'new data after purge', 'VD24-0004-0037', 1, 14),
(678, 27, 1, 10, 'new data after purge', 'VD24-0004-0038', 1, 14),
(679, 27, 1, 11, 'new data after purge', 'VD24-0004-0039', 1, 14),
(680, 27, 1, 13, 'new data after purge', 'VD24-0004-0040', 1, 14),
(681, 27, 1, 5, 'new data after purge', 'VD24-0004-0041', 1, 14),
(682, 27, 1, 6, 'new data after purge', 'VD24-0004-0042', 1, 14),
(683, 27, 1, 7, 'new data after purge', 'VD24-0004-0043', 1, 14),
(684, 27, 1, 8, 'new data after purge', 'VD24-0004-0044', 1, 14),
(685, 27, 1, 9, 'new data after purge', 'VD24-0004-0045', 1, 14),
(686, 27, 1, 10, 'new data after purge', 'VD24-0004-0046', 1, 14),
(687, 27, 1, 11, 'new data after purge', 'VD24-0004-0047', 1, 14),
(688, 27, 1, 13, 'new data after purge', 'VD24-0004-0048', 1, 14),
(689, 27, 1, 5, 'new data after purge', 'VD24-0004-0049', 1, 14),
(690, 27, 1, 6, 'new data after purge', 'VD24-0004-0050', 1, 14),
(691, 27, 1, 7, 'new data after purge', 'VD24-0004-0051', 1, 14),
(692, 27, 1, 8, 'new data after purge', 'VD24-0004-0052', 1, 14),
(693, 27, 1, 9, 'new data after purge', 'VD24-0004-0053', 1, 14),
(694, 27, 1, 10, 'new data after purge', 'VD24-0004-0054', 1, 14),
(695, 27, 1, 11, 'new data after purge', 'VD24-0004-0055', 1, 14),
(696, 27, 1, 13, 'new data after purge', 'VD24-0004-0056', 1, 14),
(697, 27, 1, 5, 'new data after purge', 'VD24-0004-0057', 1, 14),
(698, 27, 1, 6, 'new data after purge', 'VD24-0004-0058', 1, 14),
(699, 27, 1, 7, 'new data after purge', 'VD24-0004-0059', 1, 14),
(700, 27, 1, 8, 'new data after purge', 'VD24-0004-0060', 1, 14),
(701, 27, 1, 9, 'new data after purge', 'VD24-0004-0061', 1, 14),
(702, 27, 1, 10, 'new data after purge', 'VD24-0004-0062', 1, 14),
(703, 27, 1, 11, 'new data after purge', 'VD24-0004-0063', 1, 14),
(704, 27, 1, 13, 'new data after purge', 'VD24-0004-0064', 1, 14),
(705, 28, 1, 5, 'optimization test', 'TC24-0001-0001', 1, 1),
(706, 28, 1, 6, 'new data after purge', 'TC24-0001-0002', 1, 14),
(707, 28, 1, 7, 'new data after purge', 'TC24-0001-0003', 1, 14),
(708, 28, 1, 8, 'new data after purge', 'TC24-0001-0004', 1, 14),
(709, 28, 1, 9, 'new data after purge', 'TC24-0001-0005', 1, 14),
(710, 28, 1, 10, 'new data after purge', 'TC24-0001-0006', 1, 14),
(711, 28, 1, 11, 'new data after purge', 'TC24-0001-0007', 1, 14),
(712, 28, 1, 13, 'new data after purge', 'TC24-0001-0008', 1, 14),
(713, 28, 1, 5, 'new data after purge', 'TC24-0001-0009', 1, 14),
(714, 28, 1, 6, 'new data after purge', 'TC24-0001-0010', 1, 14),
(715, 28, 1, 7, 'new data after purge', 'TC24-0001-0011', 1, 14),
(716, 28, 1, 8, 'new data after purge', 'TC24-0001-0012', 1, 14),
(717, 28, 1, 9, 'new data after purge', 'TC24-0001-0013', 1, 14),
(718, 28, 1, 10, 'new data after purge', 'TC24-0001-0014', 1, 14),
(719, 28, 1, 11, 'new data after purge', 'TC24-0001-0015', 1, 14),
(720, 28, 1, 13, 'new data after purge', 'TC24-0001-0016', 1, 14),
(721, 28, 1, 5, 'new data after purge', 'TC24-0001-0017', 1, 14),
(722, 28, 1, 6, 'new data after purge', 'TC24-0001-0018', 1, 14),
(723, 28, 1, 7, 'new data after purge', 'TC24-0001-0019', 1, 14),
(724, 28, 1, 8, 'new data after purge', 'TC24-0001-0020', 1, 14),
(725, 28, 1, 9, 'new data after purge', 'TC24-0001-0021', 1, 14),
(726, 28, 1, 10, 'new data after purge', 'TC24-0001-0022', 1, 14),
(727, 28, 1, 11, 'new data after purge', 'TC24-0001-0023', 1, 14),
(728, 28, 1, 13, 'new data after purge', 'TC24-0001-0024', 1, 14),
(729, 28, 1, 5, 'new data after purge', 'TC24-0001-0025', 1, 14),
(730, 28, 1, 6, 'new data after purge', 'TC24-0001-0026', 1, 14),
(731, 28, 1, 7, 'new data after purge', 'TC24-0001-0027', 1, 14),
(732, 28, 1, 8, 'new data after purge', 'TC24-0001-0028', 1, 14),
(733, 28, 1, 9, 'new data after purge', 'TC24-0001-0029', 1, 14),
(734, 28, 1, 10, 'new data after purge', 'TC24-0001-0030', 1, 14),
(735, 28, 1, 11, 'new data after purge', 'TC24-0001-0031', 1, 14),
(736, 28, 1, 13, 'new data after purge', 'TC24-0001-0032', 1, 14),
(737, 28, 1, 5, 'new data after purge', 'TC24-0001-0033', 1, 14),
(738, 28, 1, 6, 'new data after purge', 'TC24-0001-0034', 1, 14),
(739, 28, 1, 7, 'new data after purge', 'TC24-0001-0035', 1, 14),
(740, 28, 1, 8, 'new data after purge', 'TC24-0001-0036', 1, 14),
(741, 28, 1, 9, 'new data after purge', 'TC24-0001-0037', 1, 14),
(742, 28, 1, 10, 'new data after purge', 'TC24-0001-0038', 1, 14),
(743, 28, 1, 11, 'new data after purge', 'TC24-0001-0039', 1, 14),
(744, 28, 1, 13, 'new data after purge', 'TC24-0001-0040', 1, 14),
(745, 28, 1, 5, 'new data after purge', 'TC24-0001-0041', 1, 14),
(746, 28, 1, 6, 'new data after purge', 'TC24-0001-0042', 1, 14),
(747, 28, 1, 7, 'new data after purge', 'TC24-0001-0043', 1, 14),
(748, 28, 1, 8, 'new data after purge', 'TC24-0001-0044', 1, 14),
(749, 28, 1, 9, 'new data after purge', 'TC24-0001-0045', 1, 14),
(750, 28, 1, 10, 'new data after purge', 'TC24-0001-0046', 1, 14),
(751, 28, 1, 11, 'new data after purge', 'TC24-0001-0047', 1, 14),
(752, 28, 1, 13, 'new data after purge', 'TC24-0001-0048', 1, 14),
(753, 28, 1, 5, 'new data after purge', 'TC24-0001-0049', 1, 14),
(754, 28, 1, 6, 'new data after purge', 'TC24-0001-0050', 1, 14),
(755, 28, 1, 7, 'new data after purge', 'TC24-0001-0051', 1, 14),
(756, 28, 1, 8, 'new data after purge', 'TC24-0001-0052', 1, 14),
(757, 28, 1, 9, 'new data after purge', 'TC24-0001-0053', 1, 14),
(758, 28, 1, 10, 'new data after purge', 'TC24-0001-0054', 1, 14),
(759, 28, 1, 11, 'new data after purge', 'TC24-0001-0055', 1, 14),
(760, 28, 1, 13, 'new data after purge', 'TC24-0001-0056', 1, 14),
(761, 28, 1, 5, 'new data after purge', 'TC24-0001-0057', 1, 14),
(762, 28, 1, 6, 'new data after purge', 'TC24-0001-0058', 1, 14),
(763, 28, 1, 7, 'new data after purge', 'TC24-0001-0059', 1, 14),
(764, 28, 1, 8, 'new data after purge', 'TC24-0001-0060', 1, 14),
(765, 28, 1, 9, 'new data after purge', 'TC24-0001-0061', 1, 14),
(766, 28, 1, 10, 'new data after purge', 'TC24-0001-0062', 1, 14),
(767, 28, 1, 11, 'new data after purge', 'TC24-0001-0063', 1, 14),
(768, 28, 1, 13, 'new data after purge', 'TC24-0001-0064', 1, 14),
(769, 29, 1, 5, 'new data after purge', 'TC24-0002-0001', 1, 1),
(770, 29, 1, 6, 'new data after purge', 'TC24-0002-0002', 1, 14),
(771, 29, 1, 7, 'new data after purge', 'TC24-0002-0003', 1, 14),
(772, 29, 1, 8, 'new data after purge', 'TC24-0002-0004', 1, 14),
(773, 29, 1, 9, 'new data after purge', 'TC24-0002-0005', 1, 14),
(774, 29, 1, 10, 'new data after purge', 'TC24-0002-0006', 1, 14),
(775, 29, 1, 11, 'new data after purge', 'TC24-0002-0007', 1, 14),
(776, 29, 1, 13, 'new data after purge', 'TC24-0002-0008', 1, 14),
(777, 29, 1, 5, 'new data after purge', 'TC24-0002-0009', 1, 14);
INSERT INTO `item_unit` (`id_unit`, `id_item`, `status`, `id_wh`, `comment`, `serial_number`, `condition`, `updated_by`) VALUES
(778, 29, 1, 6, 'new data after purge', 'TC24-0002-0010', 1, 14),
(779, 29, 1, 7, 'new data after purge', 'TC24-0002-0011', 1, 14),
(780, 29, 1, 8, 'new data after purge', 'TC24-0002-0012', 1, 14),
(781, 29, 1, 9, 'new data after purge', 'TC24-0002-0013', 1, 14),
(782, 29, 1, 10, 'new data after purge', 'TC24-0002-0014', 1, 14),
(783, 29, 1, 11, 'new data after purge', 'TC24-0002-0015', 1, 14),
(784, 29, 1, 13, 'new data after purge', 'TC24-0002-0016', 1, 14),
(785, 29, 1, 5, 'new data after purge', 'TC24-0002-0017', 1, 14),
(786, 29, 1, 6, 'new data after purge', 'TC24-0002-0018', 1, 14),
(787, 29, 1, 7, 'new data after purge', 'TC24-0002-0019', 1, 14),
(788, 29, 1, 8, 'new data after purge', 'TC24-0002-0020', 1, 14),
(789, 29, 1, 9, 'new data after purge', 'TC24-0002-0021', 1, 14),
(790, 29, 1, 10, 'new data after purge', 'TC24-0002-0022', 1, 14),
(791, 29, 1, 11, 'new data after purge', 'TC24-0002-0023', 1, 14),
(792, 29, 1, 13, 'new data after purge', 'TC24-0002-0024', 1, 14),
(793, 29, 1, 5, 'new data after purge', 'TC24-0002-0025', 1, 14),
(794, 29, 1, 6, 'new data after purge', 'TC24-0002-0026', 1, 14),
(795, 29, 1, 7, 'new data after purge', 'TC24-0002-0027', 1, 14),
(796, 29, 1, 8, 'new data after purge', 'TC24-0002-0028', 1, 14),
(797, 29, 1, 9, 'new data after purge', 'TC24-0002-0029', 1, 14),
(798, 29, 1, 10, 'new data after purge', 'TC24-0002-0030', 1, 14),
(799, 29, 1, 11, 'new data after purge', 'TC24-0002-0031', 1, 14),
(800, 29, 1, 13, 'new data after purge', 'TC24-0002-0032', 1, 14),
(801, 29, 1, 5, 'new data after purge', 'TC24-0002-0033', 1, 14),
(802, 29, 1, 6, 'new data after purge', 'TC24-0002-0034', 1, 14),
(803, 29, 1, 7, 'new data after purge', 'TC24-0002-0035', 1, 14),
(804, 29, 1, 8, 'new data after purge', 'TC24-0002-0036', 1, 14),
(805, 29, 1, 9, 'new data after purge', 'TC24-0002-0037', 1, 14),
(806, 29, 1, 10, 'new data after purge', 'TC24-0002-0038', 1, 14),
(807, 29, 1, 11, 'new data after purge', 'TC24-0002-0039', 1, 14),
(808, 29, 1, 13, 'new data after purge', 'TC24-0002-0040', 1, 14),
(809, 29, 1, 5, 'new data after purge', 'TC24-0002-0041', 1, 14),
(810, 29, 1, 6, 'new data after purge', 'TC24-0002-0042', 1, 14),
(811, 29, 1, 7, 'new data after purge', 'TC24-0002-0043', 1, 14),
(812, 29, 1, 8, 'new data after purge', 'TC24-0002-0044', 1, 14),
(813, 29, 1, 9, 'new data after purge', 'TC24-0002-0045', 1, 14),
(814, 29, 1, 10, 'new data after purge', 'TC24-0002-0046', 1, 14),
(815, 29, 1, 11, 'new data after purge', 'TC24-0002-0047', 1, 14),
(816, 29, 1, 13, 'new data after purge', 'TC24-0002-0048', 1, 14),
(817, 29, 1, 5, 'new data after purge', 'TC24-0002-0049', 1, 14),
(818, 29, 1, 6, 'new data after purge', 'TC24-0002-0050', 1, 14),
(819, 29, 1, 7, 'new data after purge', 'TC24-0002-0051', 1, 14),
(820, 29, 1, 8, 'new data after purge', 'TC24-0002-0052', 1, 14),
(821, 29, 1, 9, 'new data after purge', 'TC24-0002-0053', 1, 14),
(822, 29, 1, 10, 'new data after purge', 'TC24-0002-0054', 1, 14),
(823, 29, 1, 11, 'new data after purge', 'TC24-0002-0055', 1, 14),
(824, 29, 1, 13, 'new data after purge', 'TC24-0002-0056', 1, 14),
(825, 29, 1, 5, 'new data after purge', 'TC24-0002-0057', 1, 14),
(826, 29, 1, 6, 'new data after purge', 'TC24-0002-0058', 1, 14),
(827, 29, 1, 7, 'new data after purge', 'TC24-0002-0059', 1, 14),
(828, 29, 1, 8, 'new data after purge', 'TC24-0002-0060', 1, 14),
(829, 29, 1, 9, 'new data after purge', 'TC24-0002-0061', 1, 14),
(830, 29, 1, 10, 'new data after purge', 'TC24-0002-0062', 1, 14),
(831, 29, 1, 11, 'new data after purge', 'TC24-0002-0063', 1, 14),
(832, 29, 1, 13, 'new data after purge', 'TC24-0002-0064', 1, 14),
(833, 30, 1, 5, 'opt test', 'TC24-0003-0001', 1, 1),
(834, 30, 1, 6, 'new data after purge', 'TC24-0003-0002', 1, 14),
(835, 30, 1, 7, 'new data after purge', 'TC24-0003-0003', 1, 14),
(836, 30, 1, 8, 'new data after purge', 'TC24-0003-0004', 1, 14),
(837, 30, 1, 9, 'new data after purge', 'TC24-0003-0005', 1, 14),
(838, 30, 1, 10, 'new data after purge', 'TC24-0003-0006', 1, 14),
(839, 30, 1, 11, 'new data after purge', 'TC24-0003-0007', 1, 14),
(840, 30, 1, 13, 'new data after purge', 'TC24-0003-0008', 1, 14),
(841, 30, 1, 5, 'new data after purge', 'TC24-0003-0009', 1, 14),
(842, 30, 1, 6, 'new data after purge', 'TC24-0003-0010', 1, 14),
(843, 30, 1, 7, 'new data after purge', 'TC24-0003-0011', 1, 14),
(844, 30, 1, 8, 'new data after purge', 'TC24-0003-0012', 1, 14),
(845, 30, 1, 9, 'new data after purge', 'TC24-0003-0013', 1, 14),
(846, 30, 1, 10, 'new data after purge', 'TC24-0003-0014', 1, 14),
(847, 30, 1, 11, 'new data after purge', 'TC24-0003-0015', 1, 14),
(848, 30, 1, 13, 'new data after purge', 'TC24-0003-0016', 1, 14),
(849, 30, 1, 5, 'new data after purge', 'TC24-0003-0017', 1, 14),
(850, 30, 1, 6, 'new data after purge', 'TC24-0003-0018', 1, 14),
(851, 30, 1, 7, 'new data after purge', 'TC24-0003-0019', 1, 14),
(852, 30, 1, 8, 'new data after purge', 'TC24-0003-0020', 1, 14),
(853, 30, 1, 9, 'new data after purge', 'TC24-0003-0021', 1, 14),
(854, 30, 1, 10, 'new data after purge', 'TC24-0003-0022', 1, 14),
(855, 30, 1, 11, 'new data after purge', 'TC24-0003-0023', 1, 14),
(856, 30, 1, 13, 'new data after purge', 'TC24-0003-0024', 1, 14),
(857, 30, 1, 5, 'new data after purge', 'TC24-0003-0025', 1, 14),
(858, 30, 1, 6, 'new data after purge', 'TC24-0003-0026', 1, 14),
(859, 30, 1, 7, 'new data after purge', 'TC24-0003-0027', 1, 14),
(860, 30, 1, 8, 'new data after purge', 'TC24-0003-0028', 1, 14),
(861, 30, 1, 9, 'new data after purge', 'TC24-0003-0029', 1, 14),
(862, 30, 1, 10, 'new data after purge', 'TC24-0003-0030', 1, 14),
(863, 30, 1, 11, 'new data after purge', 'TC24-0003-0031', 1, 14),
(864, 30, 1, 13, 'new data after purge', 'TC24-0003-0032', 1, 14),
(865, 30, 1, 5, 'new data after purge', 'TC24-0003-0033', 1, 14),
(866, 30, 1, 6, 'new data after purge', 'TC24-0003-0034', 1, 14),
(867, 30, 1, 7, 'new data after purge', 'TC24-0003-0035', 1, 14),
(868, 30, 1, 8, 'new data after purge', 'TC24-0003-0036', 1, 14),
(869, 30, 1, 9, 'new data after purge', 'TC24-0003-0037', 1, 14),
(870, 30, 1, 10, 'new data after purge', 'TC24-0003-0038', 1, 14),
(871, 30, 1, 11, 'new data after purge', 'TC24-0003-0039', 1, 14),
(872, 30, 1, 13, 'new data after purge', 'TC24-0003-0040', 1, 14),
(873, 30, 1, 5, 'new data after purge', 'TC24-0003-0041', 1, 14),
(874, 30, 1, 6, 'new data after purge', 'TC24-0003-0042', 1, 14),
(875, 30, 1, 7, 'new data after purge', 'TC24-0003-0043', 1, 14),
(876, 30, 1, 8, 'new data after purge', 'TC24-0003-0044', 1, 14),
(877, 30, 1, 9, 'new data after purge', 'TC24-0003-0045', 1, 14),
(878, 30, 1, 10, 'new data after purge', 'TC24-0003-0046', 1, 14),
(879, 30, 1, 11, 'new data after purge', 'TC24-0003-0047', 1, 14),
(880, 30, 1, 13, 'new data after purge', 'TC24-0003-0048', 1, 14),
(881, 30, 1, 5, 'new data after purge', 'TC24-0003-0049', 1, 14),
(882, 30, 1, 6, 'new data after purge', 'TC24-0003-0050', 1, 14),
(883, 30, 1, 7, 'new data after purge', 'TC24-0003-0051', 1, 14),
(884, 30, 1, 8, 'new data after purge', 'TC24-0003-0052', 1, 14),
(885, 30, 1, 9, 'new data after purge', 'TC24-0003-0053', 1, 14),
(886, 30, 1, 10, 'new data after purge', 'TC24-0003-0054', 1, 14),
(887, 30, 1, 11, 'new data after purge', 'TC24-0003-0055', 1, 14),
(888, 30, 1, 13, 'new data after purge', 'TC24-0003-0056', 1, 14),
(889, 30, 1, 5, 'new data after purge', 'TC24-0003-0057', 1, 14),
(890, 30, 1, 6, 'new data after purge', 'TC24-0003-0058', 1, 14),
(891, 30, 1, 7, 'new data after purge', 'TC24-0003-0059', 1, 14),
(892, 30, 1, 8, 'new data after purge', 'TC24-0003-0060', 1, 14),
(893, 30, 1, 9, 'new data after purge', 'TC24-0003-0061', 1, 14),
(894, 30, 1, 10, 'new data after purge', 'TC24-0003-0062', 1, 14),
(895, 30, 1, 11, 'new data after purge', 'TC24-0003-0063', 1, 14),
(896, 30, 1, 13, 'new data after purge', 'TC24-0003-0064', 1, 14),
(897, 31, 2, 5, 'new data after purge', 'PU24-0001-0001', 1, 14),
(898, 31, 1, 6, 'new data after purge', 'PU24-0001-0002', 1, 14),
(899, 31, 1, 7, 'new data after purge', 'PU24-0001-0003', 1, 14),
(900, 31, 1, 8, 'new data after purge', 'PU24-0001-0004', 1, 14),
(901, 31, 1, 9, 'new data after purge', 'PU24-0001-0005', 1, 14),
(902, 31, 1, 10, 'new data after purge', 'PU24-0001-0006', 1, 14),
(903, 31, 1, 11, 'new data after purge', 'PU24-0001-0007', 1, 14),
(904, 31, 1, 13, 'new data after purge', 'PU24-0001-0008', 1, 14),
(905, 31, 1, 5, 'new data after purge', 'PU24-0001-0009', 1, 14),
(906, 31, 1, 6, 'new data after purge', 'PU24-0001-0010', 1, 14),
(907, 31, 1, 7, 'new data after purge', 'PU24-0001-0011', 1, 14),
(908, 31, 1, 8, 'new data after purge', 'PU24-0001-0012', 1, 14),
(909, 31, 1, 9, 'new data after purge', 'PU24-0001-0013', 1, 14),
(910, 31, 1, 10, 'new data after purge', 'PU24-0001-0014', 1, 14),
(911, 31, 1, 11, 'new data after purge', 'PU24-0001-0015', 1, 14),
(912, 31, 1, 13, 'new data after purge', 'PU24-0001-0016', 1, 14),
(913, 31, 1, 5, 'new data after purge', 'PU24-0001-0017', 1, 14),
(914, 31, 1, 6, 'new data after purge', 'PU24-0001-0018', 1, 14),
(915, 31, 1, 7, 'new data after purge', 'PU24-0001-0019', 1, 14),
(916, 31, 1, 8, 'new data after purge', 'PU24-0001-0020', 1, 14),
(917, 31, 1, 9, 'new data after purge', 'PU24-0001-0021', 1, 14),
(918, 31, 1, 10, 'new data after purge', 'PU24-0001-0022', 1, 14),
(919, 31, 1, 11, 'new data after purge', 'PU24-0001-0023', 1, 14),
(920, 31, 1, 13, 'new data after purge', 'PU24-0001-0024', 1, 14),
(921, 31, 1, 5, 'new data after purge', 'PU24-0001-0025', 1, 14),
(922, 31, 1, 6, 'new data after purge', 'PU24-0001-0026', 1, 14),
(923, 31, 1, 7, 'new data after purge', 'PU24-0001-0027', 1, 14),
(924, 31, 1, 8, 'new data after purge', 'PU24-0001-0028', 1, 14),
(925, 31, 1, 9, 'new data after purge', 'PU24-0001-0029', 1, 14),
(926, 31, 1, 10, 'new data after purge', 'PU24-0001-0030', 1, 14),
(927, 31, 1, 11, 'new data after purge', 'PU24-0001-0031', 1, 14),
(928, 31, 1, 13, 'new data after purge', 'PU24-0001-0032', 1, 14),
(929, 31, 1, 5, 'new data after purge', 'PU24-0001-0033', 1, 14),
(930, 31, 1, 6, 'new data after purge', 'PU24-0001-0034', 1, 14),
(931, 31, 1, 7, 'new data after purge', 'PU24-0001-0035', 1, 14),
(932, 31, 1, 8, 'new data after purge', 'PU24-0001-0036', 1, 14),
(933, 31, 1, 9, 'new data after purge', 'PU24-0001-0037', 1, 14),
(934, 31, 1, 10, 'new data after purge', 'PU24-0001-0038', 1, 14),
(935, 31, 1, 11, 'new data after purge', 'PU24-0001-0039', 1, 14),
(936, 31, 1, 13, 'new data after purge', 'PU24-0001-0040', 1, 14),
(937, 31, 1, 5, 'new data after purge', 'PU24-0001-0041', 1, 14),
(938, 31, 1, 6, 'new data after purge', 'PU24-0001-0042', 1, 14),
(939, 31, 1, 7, 'new data after purge', 'PU24-0001-0043', 1, 14),
(940, 31, 1, 8, 'new data after purge', 'PU24-0001-0044', 1, 14),
(941, 31, 1, 9, 'new data after purge', 'PU24-0001-0045', 1, 14),
(942, 31, 1, 10, 'new data after purge', 'PU24-0001-0046', 1, 14),
(943, 31, 1, 11, 'new data after purge', 'PU24-0001-0047', 1, 14),
(944, 31, 1, 13, 'new data after purge', 'PU24-0001-0048', 1, 14),
(945, 31, 1, 5, 'new data after purge', 'PU24-0001-0049', 1, 14),
(946, 31, 1, 6, 'new data after purge', 'PU24-0001-0050', 1, 14),
(947, 31, 1, 7, 'new data after purge', 'PU24-0001-0051', 1, 14),
(948, 31, 1, 8, 'new data after purge', 'PU24-0001-0052', 1, 14),
(949, 31, 1, 9, 'new data after purge', 'PU24-0001-0053', 1, 14),
(950, 31, 1, 10, 'new data after purge', 'PU24-0001-0054', 1, 14),
(951, 31, 1, 11, 'new data after purge', 'PU24-0001-0055', 1, 14),
(952, 31, 1, 13, 'new data after purge', 'PU24-0001-0056', 1, 14),
(953, 31, 1, 5, 'new data after purge', 'PU24-0001-0057', 1, 14),
(954, 31, 1, 6, 'new data after purge', 'PU24-0001-0058', 1, 14),
(955, 31, 1, 7, 'new data after purge', 'PU24-0001-0059', 1, 14),
(956, 31, 1, 8, 'new data after purge', 'PU24-0001-0060', 1, 14),
(957, 31, 1, 9, 'new data after purge', 'PU24-0001-0061', 1, 14),
(958, 31, 1, 10, 'new data after purge', 'PU24-0001-0062', 1, 14),
(959, 31, 1, 11, 'new data after purge', 'PU24-0001-0063', 1, 14),
(960, 31, 1, 13, 'new data after purge', 'PU24-0001-0064', 1, 14),
(961, 32, 1, 5, 'new data after purge', 'TC24-0004-0001', 1, 1),
(962, 32, 1, 6, 'new data after purge', 'TC24-0004-0002', 1, 14),
(963, 32, 1, 7, 'new data after purge', 'TC24-0004-0003', 1, 14),
(964, 32, 1, 8, 'new data after purge', 'TC24-0004-0004', 1, 14),
(965, 32, 1, 9, 'new data after purge', 'TC24-0004-0005', 1, 14),
(966, 32, 1, 10, 'new data after purge', 'TC24-0004-0006', 1, 14),
(967, 32, 1, 11, 'new data after purge', 'TC24-0004-0007', 1, 14),
(968, 32, 1, 13, 'new data after purge', 'TC24-0004-0008', 1, 14),
(969, 32, 1, 5, 'new data after purge', 'TC24-0004-0009', 1, 14),
(970, 32, 1, 6, 'new data after purge', 'TC24-0004-0010', 1, 14),
(971, 32, 1, 7, 'new data after purge', 'TC24-0004-0011', 1, 14),
(972, 32, 1, 8, 'new data after purge', 'TC24-0004-0012', 1, 14),
(973, 32, 1, 9, 'new data after purge', 'TC24-0004-0013', 1, 14),
(974, 32, 1, 10, 'new data after purge', 'TC24-0004-0014', 1, 14),
(975, 32, 1, 11, 'new data after purge', 'TC24-0004-0015', 1, 14),
(976, 32, 1, 13, 'new data after purge', 'TC24-0004-0016', 1, 14),
(977, 32, 1, 5, 'new data after purge', 'TC24-0004-0017', 1, 14),
(978, 32, 1, 6, 'new data after purge', 'TC24-0004-0018', 1, 14),
(979, 32, 1, 7, 'new data after purge', 'TC24-0004-0019', 1, 14),
(980, 32, 1, 8, 'new data after purge', 'TC24-0004-0020', 1, 14),
(981, 32, 1, 9, 'new data after purge', 'TC24-0004-0021', 1, 14),
(982, 32, 1, 10, 'new data after purge', 'TC24-0004-0022', 1, 14),
(983, 32, 1, 11, 'new data after purge', 'TC24-0004-0023', 1, 14),
(984, 32, 1, 13, 'new data after purge', 'TC24-0004-0024', 1, 14),
(985, 32, 1, 5, 'new data after purge', 'TC24-0004-0025', 1, 14),
(986, 32, 1, 6, 'new data after purge', 'TC24-0004-0026', 1, 14),
(987, 32, 1, 7, 'new data after purge', 'TC24-0004-0027', 1, 14),
(988, 32, 1, 8, 'new data after purge', 'TC24-0004-0028', 1, 14),
(989, 32, 1, 9, 'new data after purge', 'TC24-0004-0029', 1, 14),
(990, 32, 1, 10, 'new data after purge', 'TC24-0004-0030', 1, 14),
(991, 32, 1, 11, 'new data after purge', 'TC24-0004-0031', 1, 14),
(992, 32, 1, 13, 'new data after purge', 'TC24-0004-0032', 1, 14),
(993, 32, 1, 5, 'new data after purge', 'TC24-0004-0033', 1, 14),
(994, 32, 1, 6, 'new data after purge', 'TC24-0004-0034', 1, 14),
(995, 32, 1, 7, 'new data after purge', 'TC24-0004-0035', 1, 14),
(996, 32, 1, 8, 'new data after purge', 'TC24-0004-0036', 1, 14),
(997, 32, 1, 9, 'new data after purge', 'TC24-0004-0037', 1, 14),
(998, 32, 1, 10, 'new data after purge', 'TC24-0004-0038', 1, 14),
(999, 32, 1, 11, 'new data after purge', 'TC24-0004-0039', 1, 14),
(1000, 32, 1, 13, 'new data after purge', 'TC24-0004-0040', 1, 14),
(1001, 32, 1, 5, 'new data after purge', 'TC24-0004-0041', 1, 14),
(1002, 32, 1, 6, 'new data after purge', 'TC24-0004-0042', 1, 14),
(1003, 32, 1, 7, 'new data after purge', 'TC24-0004-0043', 1, 14),
(1004, 32, 1, 8, 'new data after purge', 'TC24-0004-0044', 1, 14),
(1005, 32, 1, 9, 'new data after purge', 'TC24-0004-0045', 1, 14),
(1006, 32, 1, 10, 'new data after purge', 'TC24-0004-0046', 1, 14),
(1007, 32, 1, 11, 'new data after purge', 'TC24-0004-0047', 1, 14),
(1008, 32, 1, 13, 'new data after purge', 'TC24-0004-0048', 1, 14),
(1009, 32, 1, 5, 'new data after purge', 'TC24-0004-0049', 1, 14),
(1010, 32, 1, 6, 'new data after purge', 'TC24-0004-0050', 1, 14),
(1011, 32, 1, 7, 'new data after purge', 'TC24-0004-0051', 1, 14),
(1012, 32, 1, 8, 'new data after purge', 'TC24-0004-0052', 1, 14),
(1013, 32, 1, 9, 'new data after purge', 'TC24-0004-0053', 1, 14),
(1014, 32, 1, 10, 'new data after purge', 'TC24-0004-0054', 1, 14),
(1015, 32, 1, 11, 'new data after purge', 'TC24-0004-0055', 1, 14),
(1016, 32, 1, 13, 'new data after purge', 'TC24-0004-0056', 1, 14),
(1017, 32, 1, 5, 'new data after purge', 'TC24-0004-0057', 1, 14),
(1018, 32, 1, 6, 'new data after purge', 'TC24-0004-0058', 1, 14),
(1019, 32, 1, 7, 'new data after purge', 'TC24-0004-0059', 1, 14),
(1020, 32, 1, 8, 'new data after purge', 'TC24-0004-0060', 1, 14),
(1021, 32, 1, 9, 'new data after purge', 'TC24-0004-0061', 1, 14),
(1022, 32, 1, 10, 'new data after purge', 'TC24-0004-0062', 1, 14),
(1023, 32, 1, 11, 'new data after purge', 'TC24-0004-0063', 1, 14),
(1024, 32, 1, 13, 'new data after purge', 'TC24-0004-0064', 1, 14),
(1025, 11, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0001-0065', 1, 5),
(1026, 11, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0001-0066', 1, 5),
(1027, 11, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0001-0067', 1, 5),
(1028, 11, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0001-0068', 1, 5),
(1029, 11, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0001-0069', 1, 5),
(1030, 12, 1, 13, 'SPAM PREVENTION TEST', 'PA24-0002-0065', 1, 1),
(1031, 12, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0002-0066', 1, 1),
(1032, 12, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0002-0067', 1, 1),
(1033, 12, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0002-0068', 1, 1),
(1034, 12, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0002-0069', 1, 1),
(1035, 12, 1, 13, 'test lock. Wh admin tie to a wh bulk add unit', 'PA24-0002-0070', 1, 1),
(1036, 29, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0002-0065', 1, 1),
(1037, 29, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0002-0066', 1, 1),
(1038, 29, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0002-0067', 1, 1),
(1039, 29, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0002-0068', 1, 1),
(1040, 29, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0002-0069', 1, 1),
(1041, 29, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0002-0070', 1, 1),
(1042, 37, 1, 16, 'OPTIMIZATION TEST LOG', 'TC24-0006-0001', 1, 1),
(1043, 37, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0006-0002', 1, 1),
(1044, 37, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0006-0003', 1, 1),
(1045, 37, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0006-0004', 1, 1),
(1046, 37, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0006-0005', 1, 1),
(1047, 37, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0006-0006', 1, 1),
(1048, 37, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0006-0007', 1, 1),
(1049, 37, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0006-0008', 1, 1),
(1050, 37, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0006-0009', 1, 1),
(1051, 37, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0006-0010', 1, 1),
(1052, 37, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0006-0011', 1, 1),
(1053, 37, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0006-0012', 1, 1),
(1054, 37, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0006-0013', 1, 1),
(1055, 36, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0005-0001', 1, 1),
(1056, 36, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0005-0002', 1, 1),
(1057, 36, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0005-0003', 1, 1),
(1058, 36, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0005-0004', 1, 1),
(1059, 36, 1, 5, 'SPAM PREVENTION TEST', 'TC24-0005-0005', 1, 1),
(1060, 36, 1, 6, 'SPAM PREVENTION TEST', 'TC24-0005-0006', 1, 1),
(1063, 33, 2, 8, 'Updated desk unit comment api test', 'DK24-0001-9999', 2, 14),
(1074, 33, 1, 8, 'api test create new', 'DK24-0001-91111', 1, 14),
(1075, 39, 1, 7, 'Updated unit comment api test AA EDIT', 'PA24-0009-0020', 2, 14),
(1077, 39, 1, 8, 'api test create new', 'PA24-0009-0001', 1, 1),
(1078, 39, 1, 8, 'api test create new 2', 'PA24-0009-0002', 1, 1),
(1079, 39, 2, 8, 'Updated unit comment api test 3 EDIT', 'PA24-0009-0003', 2, 14),
(1080, 55, 1, 8, 'api test create new G11', 'PA24-0011-0001', 1, 1),
(1084, 55, 1, 8, 'api test create new G11', 'PA24-0011-0005', 1, 1),
(1085, 55, 1, 8, 'api test create new G11', 'PA24-0011-0006', 1, 1),
(1086, 55, 1, 8, 'api test create new G11', 'PA24-0011-0007', 1, 1),
(1088, 55, 1, 8, 'api test create new G11', 'PA24-0011-0008', 1, 1),
(1089, 55, 1, 8, 'api test create new G11', 'PA24-0011-0009', 1, 1),
(1090, 55, 1, 8, 'api test create new G11', 'PA24-0011-0010', 1, 1),
(1091, 55, 1, 8, 'api test create new G11', 'PA24-0011-0011', 1, 1),
(1092, 55, 1, 8, 'api test create new G11', 'PA24-0011-0012', 1, 1),
(1093, 55, 1, 8, 'api test create new G11', 'PA24-0011-0013', 1, 1),
(1094, 55, 1, 8, 'api test create new G11', 'PA24-0011-0014', 1, 1),
(1095, 55, 1, 8, 'api test create new G11', 'PA24-0011-0015', 1, 1),
(1096, 55, 1, 8, 'api test create new G11', 'PA24-0011-0016', 1, 1),
(1097, 55, 1, 8, 'api test create new G11', 'PA24-0011-0017', 1, 1),
(1098, 55, 1, 8, 'api test create new G11', 'PA24-0011-0018', 1, 1),
(1099, 55, 1, 8, 'api test create new G11', 'PA24-0011-0019', 1, 1),
(1100, 55, 1, 8, 'api test create new G11', 'PA24-0011-0020', 1, 1),
(1101, 55, 1, 8, 'api test create new G11', 'PA24-0011-0021', 1, 1),
(1102, 55, 1, 8, 'api test create new G11', 'PA24-0011-0022', 1, 1),
(1103, 55, 1, 8, 'api test create new G11', 'PA24-0011-0023', 1, 1),
(1104, 55, 1, 8, 'api test create new G11', 'PA24-0011-0024', 1, 1),
(1105, 55, 1, 8, 'api test create new G11', 'PA24-0011-0025', 1, 1),
(1106, 55, 1, 8, 'api test create new G11', 'PA24-0011-0026', 1, 1),
(1107, 55, 1, 8, 'api test create new G11', 'PA24-0011-0027', 1, 1),
(1108, 55, 1, 8, 'api test create new G11', 'PA24-0011-0028', 1, 1),
(1109, 55, 1, 8, 'api test create new G11', 'PA24-0011-0029', 1, 1),
(1110, 55, 1, 8, 'api test create new G11', 'PA24-0011-0030', 1, 1),
(1111, 55, 1, 8, 'api test create new G11', 'PA24-0011-0031', 1, 1),
(1112, 55, 1, 8, 'api test create new G11', 'PA24-0011-0032', 1, 1),
(1113, 55, 1, 8, 'api test create new G11', 'PA24-0011-0033', 1, 1),
(1114, 55, 1, 8, 'api test create new G11', 'PA24-0011-0034', 1, 1),
(1115, 55, 1, 8, 'api test create new G11', 'PA24-0011-0035', 1, 1),
(1116, 55, 1, 8, 'api test create new G11', 'PA24-0011-0036', 1, 1),
(1117, 55, 1, 8, 'api test create new G11', 'PA24-0011-0037', 1, 1),
(1118, 55, 1, 8, 'api test create new G11', 'PA24-0011-0038', 1, 1),
(1119, 55, 1, 8, 'api test create new G11', 'PA24-0011-0039', 1, 1),
(1120, 55, 1, 8, 'api test create new G11', 'PA24-0011-0040', 1, 1),
(1121, 55, 1, 8, 'api test create new G11', 'PA24-0011-0041', 1, 1),
(1122, 55, 1, 8, 'api test create new G11', 'PA24-0011-0042', 1, 1),
(1123, 55, 1, 8, 'api test create new G11', 'PA24-0011-0043', 1, 1),
(1124, 55, 1, 8, 'api test create new G11', 'PA24-0011-0044', 1, 1),
(1125, 55, 1, 8, 'api test create new G11', 'PA24-0011-0045', 1, 1),
(1126, 55, 1, 8, 'api test create new G11', 'PA24-0011-0046', 1, 1),
(1127, 55, 1, 8, 'api test create new G11', 'PA24-0011-0047', 1, 1),
(1128, 55, 1, 8, 'api test create new G11', 'PA24-0011-0048', 1, 1),
(1129, 55, 1, 8, 'api test create new G11', 'PA24-0011-0049', 1, 1),
(1130, 55, 1, 8, 'api test create new G11', 'PA24-0011-0050', 1, 1),
(1131, 55, 1, 8, 'api test create new G11', 'PA24-0011-0051', 1, 1),
(1132, 55, 1, 8, 'api test create new G11', 'PA24-0011-0052', 1, 1),
(1133, 55, 1, 8, 'api test create new G11', 'PA24-0011-0053', 1, 1),
(1134, 55, 1, 8, 'api test create new G11', 'PA24-0011-0054', 1, 1),
(1135, 55, 1, 8, 'api test create new G11', 'PA24-0011-0055', 1, 1),
(1136, 55, 1, 8, 'api test create new G11', 'PA24-0011-0056', 1, 1),
(1137, 55, 1, 8, 'api test create new G11', 'PA24-0011-0057', 1, 1),
(1138, 55, 1, 8, 'api test create new G11', 'PA24-0011-0058', 1, 1),
(1139, 55, 1, 8, 'api test create new G11', 'PA24-0011-0059', 1, 1),
(1140, 55, 1, 8, 'api test create new G11', 'PA24-0011-0060', 1, 1),
(1141, 55, 1, 8, 'api test create new G11', 'PA24-0011-0061', 1, 1),
(1142, 55, 1, 8, 'api test create new G11', 'PA24-0011-0062', 1, 1),
(1143, 55, 1, 8, 'api test create new G11', 'PA24-0011-0063', 1, 1),
(1144, 55, 1, 8, 'api test create new G11', 'PA24-0011-0064', 1, 1),
(1145, 55, 1, 8, 'api test create new G11', 'PA24-0011-0065', 1, 1),
(1146, 55, 1, 8, 'api test create new G11', 'PA24-0011-0066', 1, 1),
(1147, 55, 1, 8, 'api test create new G11', 'PA24-0011-0067', 1, 1),
(1148, 55, 1, 8, 'api test create new G11', 'PA24-0011-0068', 1, 1),
(1149, 55, 1, 8, 'api test create new G11', 'PA24-0011-0069', 1, 1),
(1150, 55, 1, 8, 'api test create new G11', 'PA24-0011-0070', 1, 1),
(1151, 55, 1, 8, 'api test create new G11', 'PA24-0011-0071', 1, 1),
(1152, 55, 1, 8, 'api test create new G11', 'PA24-0011-0072', 1, 1),
(1153, 55, 1, 8, 'api test create new G11', 'PA24-0011-0073', 1, 1),
(1154, 55, 1, 8, 'api test create new G11', 'PA24-0011-0074', 1, 1),
(1155, 55, 1, 8, 'api test create new G11', 'PA24-0011-0075', 1, 1),
(1156, 55, 1, 8, 'api test create new G11', 'PA24-0011-0076', 1, 1),
(1157, 55, 1, 8, 'api test create new G11', 'PA24-0011-0077', 1, 1),
(1158, 55, 1, 8, 'api test create new G11', 'PA24-0011-0078', 1, 1),
(1159, 55, 1, 8, 'api test create new G11', 'PA24-0011-0079', 1, 1),
(1160, 55, 1, 8, 'api test create new G11', 'PA24-0011-0080', 1, 1),
(1161, 55, 1, 8, 'api test create new G11', 'PA24-0011-0081', 1, 1),
(1162, 55, 1, 8, 'api test create new G11', 'PA24-0011-0082', 1, 1),
(1163, 55, 1, 8, 'api test create new G11', 'PA24-0011-0083', 1, 1),
(1164, 55, 1, 8, 'api test create new G11', 'PA24-0011-0084', 1, 1),
(1165, 55, 1, 8, 'api test create new G11', 'PA24-0011-0085', 1, 1),
(1166, 55, 1, 8, 'api test create new G11', 'PA24-0011-0086', 1, 1),
(1167, 55, 1, 8, 'api test create new G11', 'PA24-0011-0087', 1, 1),
(1168, 55, 1, 8, 'api test create new G12', 'PA24-0011-0088', 1, 1),
(1171, 55, 1, 8, 'test comment for return with pic', 'PA24-0011-0089', 2, 1);

-- --------------------------------------------------------

--
-- Table structure for table `lending`
--

CREATE TABLE `lending` (
  `id_lending` int(10) UNSIGNED NOT NULL COMMENT 'primary key',
  `id_unit` int(10) UNSIGNED NOT NULL COMMENT 'foreign key - item_unit',
  `user_id` int(11) NOT NULL COMMENT 'foreign key - user ''id''',
  `id_employee` int(10) UNSIGNED NOT NULL COMMENT 'foreign key - employee',
  `type` tinyint(3) UNSIGNED NOT NULL COMMENT '1 = in-use\r\n2 = returned',
  `date` date NOT NULL COMMENT 'date of data',
  `pic_loan` varchar(255) NOT NULL COMMENT 'pic before being used',
  `pic_return` varchar(255) NOT NULL COMMENT 'pic on return'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `lending`
--

INSERT INTO `lending` (`id_lending`, `id_unit`, `user_id`, `id_employee`, `type`, `date`, `pic_loan`, `pic_return`) VALUES
(1, 1, 14, 1, 2, '2024-11-28', '357_1732758351.jpeg', '858_1732758711.jpeg'),
(2, 129, 14, 2, 2, '2024-11-28', '203_1732758360.jpeg', '759_1732758720.jpeg'),
(3, 193, 14, 3, 2, '2024-11-28', '884_1732758368.jpeg', '135_1732758731.jpeg'),
(4, 257, 14, 4, 2, '2024-11-28', '803_1732758376.jpeg', '282_1732758740.jpeg'),
(5, 321, 14, 5, 2, '2024-11-28', '716_1732758386.jpeg', '281_1732758755.jpeg'),
(6, 385, 14, 7, 2, '2024-11-28', '742_1732758394.jpeg', '410_1732758763.jpeg'),
(7, 450, 14, 8, 2, '2024-11-29', '118_1732758403.jpeg', '624_1732868487.jpeg'),
(8, 513, 14, 8, 2, '2024-11-28', '375_1732758449.jpeg', '864_1732766351.jpeg'),
(9, 577, 14, 1, 2, '2024-12-06', '840_1732758459.jpeg', '398_1733448480.gif'),
(10, 641, 14, 2, 2, '2024-12-06', '789_1732758469.webp', '703_1733448563.gif'),
(11, 705, 14, 2, 2, '2024-12-06', '100_1732758484.jpeg', '134_1733449222.jpeg'),
(12, 769, 14, 4, 2, '2024-12-06', '139_1732758494.gif', '647_1733449400.gif'),
(13, 833, 14, 5, 2, '2024-12-06', '666_1732758504.gif', '316_1733449296.gif'),
(14, 897, 14, 8, 1, '2024-11-28', '682_1732758515.jpeg', ''),
(15, 961, 14, 4, 1, '2024-11-28', '792_1732758525.jpeg', ''),
(16, 961, 14, 4, 2, '2024-12-06', '472_1732758525.jpeg', '976_1733449470.gif'),
(17, 66, 14, 2, 1, '2024-11-28', '777_1732758537.jpeg', ''),
(18, 1, 1, 1, 1, '2024-11-29', '324_1732866705.jpeg', ''),
(19, 2, 1, 2, 1, '2024-11-29', '341_1732866727.jpeg', ''),
(20, 130, 1, 1, 1, '2024-11-29', '891_1732868467.jpeg', ''),
(21, 1, 1, 1, 2, '2024-12-02', '456_1733103106.jpeg', '246_1733103124.jpeg'),
(22, 259, 1, 2, 1, '2024-12-05', '562_1733384641.jpeg', ''),
(23, 131, 1, 8, 2, '2024-12-06', '978_1733448172.jpeg', '588_1733448413.gif'),
(24, 3, 1, 4, 2, '2024-12-09', '264_1733708513.gif', '611_1733708798.jpeg'),
(25, 4, 1, 3, 1, '2024-12-09', '731_1733728074.gif', ''),
(26, 2, 15, 2, 1, '2024-12-14', 'device_loan.jpg', 'device_return.jpg'),
(27, 2, 15, 2, 1, '2024-12-14', 'device_loan.jpg', 'device_return.jpg'),
(28, 651, 1, 2, 1, '2024-12-14', 'device_loan.jpg', 'device_return.jpg'),
(29, 22, 1, 2, 1, '2024-12-14', 'device_loan.jpg', 'device_return.jpg'),
(32, 102, 1, 2, 1, '2024-12-14', 'device_loan.jpg', 'device_return.jpg'),
(33, 102, 1, 2, 2, '2024-12-16', 'device_loadn.jpg', 'API test device_return.jpg'),
(34, 102, 1, 2, 2, '2024-12-16', 'device_loan.jpg', 'API test device_return.jpg'),
(40, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', '3d623ad2-40bb-48e8-884d-f297111dc6ae.jpeg'),
(41, 1075, 1, 2, 2, '2024-12-16', 'device_loan.jpg', ''),
(42, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(43, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(44, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(45, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(46, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(47, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(48, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(49, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(50, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(51, 1075, 1, 2, 2, '2024-12-17', 'device_loan.jpg', ''),
(52, 1171, 1, 3, 2, '2024-12-30', '0eef824d-cbd9-4c67-88b5-e900ccce86d2.jpeg', '277d4f7b-4733-4dc6-9bcc-dc7cf7b1c594.jpeg');

-- --------------------------------------------------------

--
-- Table structure for table `lending_type_lookup`
--

CREATE TABLE `lending_type_lookup` (
  `id_type` tinyint(3) UNSIGNED NOT NULL COMMENT 'primary key',
  `type_name` varchar(30) NOT NULL COMMENT 'type of lending data be it lending out or returned'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `lending_type_lookup`
--

INSERT INTO `lending_type_lookup` (`id_type`, `type_name`) VALUES
(1, 'Item going out from warehouse '),
(2, 'Item returned to warehouse');

-- --------------------------------------------------------

--
-- Table structure for table `migration`
--

CREATE TABLE `migration` (
  `version` varchar(180) NOT NULL,
  `apply_time` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `migration`
--

INSERT INTO `migration` (`version`, `apply_time`) VALUES
('m000000_000000_base', 1729233876),
('m140608_173539_create_user_table', 1729233879),
('m140611_133903_init_rbac', 1729233879),
('m140808_073114_create_auth_item_group_table', 1729233880),
('m140809_072112_insert_superadmin_to_user', 1729233880),
('m140809_073114_insert_common_permisison_to_auth_item', 1729233880),
('m141023_141535_create_user_visit_log', 1729233880),
('m141116_115804_add_bind_to_ip_and_registration_ip_to_user', 1729233880),
('m141121_194858_split_browser_and_os_column', 1729233880),
('m141201_220516_add_email_and_email_confirmed_to_user', 1729233880),
('m141207_001649_create_basic_user_permissions', 1729233881);

-- --------------------------------------------------------

--
-- Table structure for table `repair_log`
--

CREATE TABLE `repair_log` (
  `id_repair` int(10) NOT NULL COMMENT 'primary key',
  `id_unit` int(10) NOT NULL COMMENT 'foreign key - item_unit',
  `comment` varchar(120) NOT NULL COMMENT 'comments content',
  `rep_type` tinyint(3) NOT NULL COMMENT 'type of repair log',
  `datetime` datetime(2) NOT NULL COMMENT 'date & time of which the log is written'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `repair_log`
--

INSERT INTO `repair_log` (`id_repair`, `id_unit`, `comment`, `rep_type`, `datetime`) VALUES
(1, 129, 'PA24-0002-0001. rep sample', 1, '2024-11-28 08:52:58.00'),
(2, 193, 'PC24-0001-0001. rep sample', 1, '2024-11-28 08:53:02.00'),
(3, 257, 'PA24-0003-0001. rep sample', 1, '2024-11-28 08:53:06.00'),
(4, 321, 'VD24-0001-0001. rep sample', 1, '2024-11-28 08:53:11.00'),
(5, 385, 'VD24-0002-0001. rep sample', 1, '2024-11-28 08:53:16.00'),
(6, 193, 'PC24-0001-0001. rep sample', 2, '2024-11-28 08:53:29.00'),
(7, 257, 'PA24-0003-0001. rep sample', 2, '2024-11-28 08:53:39.00'),
(8, 487, 'Repair started for unit 487', 1, '2024-01-17 00:21:00.00'),
(9, 487, 'Repair completed for unit 487', 2, '2024-01-20 00:21:00.00'),
(10, 303, 'Repair started for unit 303', 1, '2024-01-23 05:44:00.00'),
(11, 303, 'Repair completed for unit 303', 2, '2024-01-27 05:44:00.00'),
(12, 348, 'Repair started for unit 348', 1, '2024-01-15 13:13:00.00'),
(13, 348, 'Repair completed for unit 348', 2, '2024-01-17 13:13:00.00'),
(14, 302, 'Repair started for unit 302', 1, '2024-01-18 09:04:00.00'),
(15, 302, 'Repair completed for unit 302', 2, '2024-01-18 09:04:00.00'),
(16, 271, 'Repair started for unit 271', 1, '2024-02-08 15:28:00.00'),
(17, 271, 'Repair completed for unit 271', 2, '2024-02-10 15:28:00.00'),
(18, 356, 'Repair started for unit 356', 1, '2024-02-08 05:20:00.00'),
(19, 356, 'Repair completed for unit 356', 2, '2024-02-08 05:20:00.00'),
(20, 96, 'Repair started for unit 96', 1, '2024-02-23 11:54:00.00'),
(21, 96, 'Repair completed for unit 96', 2, '2024-02-23 11:54:00.00'),
(22, 64, 'Repair started for unit 64', 1, '2024-02-04 04:34:00.00'),
(23, 64, 'Repair completed for unit 64', 2, '2024-02-05 04:34:00.00'),
(24, 495, 'Repair started for unit 495', 1, '2024-03-26 14:13:00.00'),
(25, 495, 'Repair completed for unit 495', 2, '2024-03-27 14:13:00.00'),
(26, 474, 'Repair started for unit 474', 1, '2024-03-22 00:50:00.00'),
(27, 474, 'Repair completed for unit 474', 2, '2024-03-22 00:50:00.00'),
(28, 70, 'Repair started for unit 70', 1, '2024-03-09 00:12:00.00'),
(29, 70, 'Repair completed for unit 70', 2, '2024-03-09 00:12:00.00'),
(30, 207, 'Repair started for unit 207', 1, '2024-03-02 23:44:00.00'),
(31, 207, 'Repair completed for unit 207', 2, '2024-03-05 23:44:00.00'),
(32, 351, 'Repair started for unit 351', 1, '2024-04-05 15:42:00.00'),
(33, 351, 'Repair completed for unit 351', 2, '2024-04-08 15:42:00.00'),
(34, 97, 'Repair started for unit 97', 1, '2024-04-28 06:30:00.00'),
(35, 97, 'Repair completed for unit 97', 2, '2024-05-01 06:30:00.00'),
(36, 383, 'Repair started for unit 383', 1, '2024-04-25 00:35:00.00'),
(37, 383, 'Repair completed for unit 383', 2, '2024-04-29 00:35:00.00'),
(38, 279, 'Repair started for unit 279', 1, '2024-04-05 05:30:00.00'),
(39, 279, 'Repair completed for unit 279', 2, '2024-04-09 05:30:00.00'),
(40, 113, 'Repair started for unit 113', 1, '2024-05-08 15:23:00.00'),
(41, 113, 'Repair completed for unit 113', 2, '2024-05-08 15:23:00.00'),
(42, 70, 'Repair started for unit 70', 1, '2024-05-15 02:59:00.00'),
(43, 70, 'Repair completed for unit 70', 2, '2024-05-18 02:59:00.00'),
(44, 221, 'Repair started for unit 221', 1, '2024-05-04 09:30:00.00'),
(45, 221, 'Repair completed for unit 221', 2, '2024-05-05 09:30:00.00'),
(46, 127, 'Repair started for unit 127', 1, '2024-05-06 07:51:00.00'),
(47, 127, 'Repair completed for unit 127', 2, '2024-05-07 07:51:00.00'),
(48, 161, 'Repair started for unit 161', 1, '2024-06-14 09:36:00.00'),
(49, 161, 'Repair completed for unit 161', 2, '2024-06-18 09:36:00.00'),
(50, 165, 'Repair started for unit 165', 1, '2024-06-05 18:18:00.00'),
(51, 165, 'Repair completed for unit 165', 2, '2024-06-06 18:18:00.00'),
(52, 289, 'Repair started for unit 289', 1, '2024-06-28 03:51:00.00'),
(53, 289, 'Repair completed for unit 289', 2, '2024-07-02 03:51:00.00'),
(54, 314, 'Repair started for unit 314', 1, '2024-06-17 03:49:00.00'),
(55, 314, 'Repair completed for unit 314', 2, '2024-06-20 03:49:00.00'),
(56, 124, 'Repair started for unit 124', 1, '2024-07-01 06:19:00.00'),
(57, 124, 'Repair completed for unit 124', 2, '2024-07-05 06:19:00.00'),
(58, 59, 'Repair started for unit 59', 1, '2024-07-04 08:17:00.00'),
(59, 59, 'Repair completed for unit 59', 2, '2024-07-06 08:17:00.00'),
(60, 86, 'Repair started for unit 86', 1, '2024-07-18 16:29:00.00'),
(61, 86, 'Repair completed for unit 86', 2, '2024-07-20 16:29:00.00'),
(62, 325, 'Repair started for unit 325', 1, '2024-07-28 21:33:00.00'),
(63, 325, 'Repair completed for unit 325', 2, '2024-07-28 21:33:00.00'),
(64, 15, 'Repair started for unit 15', 1, '2024-08-20 10:05:00.00'),
(65, 15, 'Repair completed for unit 15', 2, '2024-08-20 10:05:00.00'),
(66, 255, 'Repair started for unit 255', 1, '2024-08-02 18:43:00.00'),
(67, 255, 'Repair completed for unit 255', 2, '2024-08-03 18:43:00.00'),
(68, 116, 'Repair started for unit 116', 1, '2024-08-09 20:18:00.00'),
(69, 116, 'Repair completed for unit 116', 2, '2024-08-13 20:18:00.00'),
(70, 28, 'Repair started for unit 28', 1, '2024-08-09 06:35:00.00'),
(71, 28, 'Repair completed for unit 28', 2, '2024-08-09 06:35:00.00'),
(72, 234, 'Repair started for unit 234', 1, '2024-09-06 14:25:00.00'),
(73, 234, 'Repair completed for unit 234', 2, '2024-09-07 14:25:00.00'),
(74, 118, 'Repair started for unit 118', 1, '2024-09-08 15:23:00.00'),
(75, 118, 'Repair completed for unit 118', 2, '2024-09-08 15:23:00.00'),
(76, 494, 'Repair started for unit 494', 1, '2024-09-24 06:46:00.00'),
(77, 494, 'Repair completed for unit 494', 2, '2024-09-24 06:46:00.00'),
(78, 68, 'Repair started for unit 68', 1, '2024-09-12 13:41:00.00'),
(79, 68, 'Repair completed for unit 68', 2, '2024-09-15 13:41:00.00'),
(80, 270, 'Repair started for unit 270', 1, '2024-10-15 00:32:00.00'),
(81, 270, 'Repair completed for unit 270', 2, '2024-10-18 00:32:00.00'),
(82, 312, 'Repair started for unit 312', 1, '2024-10-05 21:05:00.00'),
(83, 312, 'Repair completed for unit 312', 2, '2024-10-08 21:05:00.00'),
(84, 173, 'Repair started for unit 173', 1, '2024-10-16 17:59:00.00'),
(85, 173, 'Repair completed for unit 173', 2, '2024-10-20 17:59:00.00'),
(86, 6, 'Repair started for unit 6', 1, '2024-10-19 05:09:00.00'),
(87, 6, 'Repair completed for unit 6', 2, '2024-10-19 05:09:00.00'),
(88, 1, 'Repair started for unit 1', 1, '2024-11-21 15:06:00.00'),
(89, 1, 'Repair completed for unit 1', 2, '2024-11-23 15:06:00.00'),
(90, 222, 'Repair started for unit 222', 1, '2024-11-16 11:37:00.00'),
(91, 222, 'Repair completed for unit 222', 2, '2024-11-19 11:37:00.00'),
(92, 465, 'Repair started for unit 465', 1, '2024-11-10 23:53:00.00'),
(93, 465, 'Repair completed for unit 465', 2, '2024-11-12 23:53:00.00'),
(94, 435, 'Repair started for unit 435', 1, '2024-11-23 10:44:00.00'),
(95, 435, 'Repair completed for unit 435', 2, '2024-11-25 10:44:00.00'),
(104, 513, 'PA24-0004-0001. new data after purge', 1, '2024-11-28 10:59:19.00'),
(105, 129, 'PA24-0002-0001. rep sample', 2, '2024-11-28 14:12:46.00'),
(106, 257, 'PA24-0003-0001. rep sample 1', 1, '2024-11-29 14:30:02.00'),
(107, 129, 'PA24-0002-0001. rep sample 2', 1, '2024-11-29 14:30:48.00'),
(108, 129, 'PA24-0002-0001. rep sample 2', 2, '2024-11-29 14:31:36.00'),
(109, 129, 'PA24-0002-0001. rep sample 3', 1, '2024-11-29 14:34:57.00'),
(110, 129, 'PA24-0002-0001. rep sample 3', 2, '2024-11-29 14:35:47.00'),
(111, 129, 'PA24-0002-0001. rep sample 4', 1, '2024-11-29 14:36:07.00'),
(112, 129, 'PA24-0002-0001. rep sample 4', 2, '2024-11-29 14:36:30.00'),
(113, 257, 'PA24-0003-0001. rep sample a', 2, '2024-11-29 14:37:50.00'),
(114, 321, 'VD24-0001-0001. rep sample b', 2, '2024-11-29 14:37:56.00'),
(116, 385, 'VD24-0002-0001. rep sample c', 2, '2024-11-29 14:38:02.00'),
(117, 513, 'PA24-0004-0001. new data after purge ', 2, '2024-11-29 14:38:08.00'),
(118, 385, 'VD24-0002-0001. rep sample d', 1, '2024-11-29 14:42:18.00'),
(119, 385, 'VD24-0002-0001. rep sample d', 2, '2024-11-29 14:42:29.00'),
(120, 257, 'PA24-0003-0001. rep sample e', 1, '2024-11-29 14:43:21.00'),
(123, 257, 'PA24-0003-0001. rep sample e', 2, '2024-11-29 14:43:29.00'),
(125, 257, 'PA24-0003-0001. rep sample aaa', 1, '2024-11-29 14:53:51.00'),
(126, 257, 'PA24-0003-0001. rep sample aaa', 2, '2024-11-29 15:02:43.00'),
(127, 129, 'PA24-0002-0001. SPAM PREVENTION TEST', 1, '2024-11-29 15:21:53.00'),
(128, 129, 'PA24-0002-0001. SPAM PREVENTION TEST', 2, '2024-11-29 15:22:02.00'),
(129, 129, 'PA24-0002-0001. SPAM PREVENTION TEST', 1, '2024-12-02 08:35:43.00'),
(130, 129, 'PA24-0002-0001. SPAM PREVENTION TEST', 2, '2024-12-02 08:35:51.00'),
(131, 129, 'PA24-0002-0001. SPAM PREVENTION TEST', 1, '2024-12-02 13:30:42.00'),
(132, 129, 'PA24-0002-0001. SPAM PREVENTION TEST', 2, '2024-12-02 13:30:52.00'),
(133, 193, 'PC24-0001-0001. rep sample', 1, '2024-12-03 09:43:27.00'),
(134, 321, 'VD24-0001-0001. rep sample b', 1, '2024-12-03 09:43:32.00'),
(135, 129, 'PA24-0002-0001. SPAM PREVENTION TEST', 1, '2024-12-05 08:40:43.00'),
(136, 257, 'PA24-0003-0001. rep sample aaa', 1, '2024-12-05 09:25:42.00'),
(137, 193, 'PC24-0001-0001. rep sample', 2, '2024-12-05 14:54:34.00'),
(138, 193, 'PC24-0001-0001. rep sample', 1, '2024-12-05 14:54:45.00'),
(139, 385, 'VD24-0002-0001. optimization test', 1, '2024-12-06 08:33:48.00'),
(140, 129, 'PA24-0002-0001. optomization test', 2, '2024-12-06 08:35:59.00'),
(141, 193, 'PC24-0001-0001. optimization test', 2, '2024-12-06 08:36:51.00'),
(142, 257, 'PA24-0003-0001. test optimization', 2, '2024-12-06 08:48:23.00'),
(143, 193, 'PC24-0001-0001. optimization test', 1, '2024-12-06 08:48:33.00'),
(144, 129, 'PA24-0002-0001. optomization test', 1, '2024-12-06 08:48:46.00'),
(145, 129, 'PA24-0002-0001. optomization test', 2, '2024-12-06 08:48:53.00'),
(146, 193, 'PC24-0001-0001. optimization test', 2, '2024-12-06 08:48:57.00'),
(147, 513, 'PA24-0004-0001. repairlog optimization test', 1, '2024-12-06 09:18:00.00'),
(148, 257, 'PA24-0003-0001. test optimization', 1, '2024-12-06 09:26:49.00'),
(149, 193, 'PC24-0001-0001. optimization test', 1, '2024-12-06 09:31:32.00'),
(150, 257, 'PA24-0003-0001. test optimization', 2, '2024-12-06 09:33:21.00'),
(151, 193, 'PC24-0001-0001. optimization test', 2, '2024-12-06 08:33:48.00'),
(152, 321, 'VD24-0001-0001. rep sample b', 2, '2024-12-06 08:34:15.00'),
(153, 321, 'VD24-0001-0001. log check', 1, '2024-12-06 09:39:43.00'),
(154, 321, 'VD24-0001-0001. log check', 2, '2024-12-09 08:00:50.00'),
(155, 385, 'VD24-0002-0001. optimization test', 2, '2024-12-09 08:00:54.00'),
(156, 513, 'PA24-0004-0001. repairlog optimization test', 2, '2024-12-09 08:00:58.00'),
(157, 513, 'Replaced faulty RAM module', 2, '2024-12-13 02:55:17.34'),
(160, 513, 'Replaced faulty RAM module', 2, '2024-12-13 07:27:51.53'),
(161, 1075, 'Replaced faulty RAM module', 1, '2024-12-16 04:40:42.60'),
(162, 1075, 'Replaced faulty RAM module', 1, '2024-12-16 04:41:13.85'),
(163, 1075, 'Replaced faulty RAM module API', 1, '2024-12-16 04:41:55.93'),
(164, 1075, 'Replaced faulty RAM module API', 1, '2024-12-16 06:03:57.01'),
(165, 1075, 'Replaced faulty RAM module API', 1, '2024-12-16 06:15:20.00'),
(166, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-16 06:17:07.23'),
(167, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-16 06:18:11.06'),
(168, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-16 06:29:23.94'),
(169, 1075, 'Replaced faulty RAM module API', 1, '2024-12-16 06:35:41.11'),
(170, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-16 06:35:41.22'),
(171, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:41:31.45'),
(172, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:41:31.55'),
(173, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:41:36.65'),
(174, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:41:36.76'),
(175, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:41:42.35'),
(176, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:41:42.47'),
(177, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:41:47.51'),
(178, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:41:47.62'),
(179, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:41:52.97'),
(180, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:41:53.09'),
(181, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:41:58.62'),
(182, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:41:58.70'),
(183, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:42:03.90'),
(184, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:42:04.02'),
(185, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:42:09.21'),
(186, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:42:09.30'),
(187, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:42:14.69'),
(188, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:42:14.79'),
(189, 1075, 'Replaced faulty RAM module API', 1, '2024-12-17 03:42:20.09'),
(190, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-17 03:42:20.23'),
(191, 1075, 'Replaced faulty RAM module API', 1, '2024-12-18 02:24:08.34'),
(192, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-18 02:24:08.46'),
(193, 1075, 'Replaced faulty RAM module API', 1, '2024-12-18 02:26:02.41'),
(194, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-18 02:26:02.54'),
(195, 1075, 'Replaced faulty RAM module API', 1, '2024-12-18 02:28:40.29'),
(196, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-18 02:28:40.40'),
(197, 1075, 'Replaced faulty RAM module API TEST', 1, '2024-12-18 03:03:36.94'),
(198, 1075, 'Replaced faulty RAM module API TEST', 1, '2024-12-18 03:03:54.32'),
(199, 1075, 'PA24-0009-0018Sent to repair by superadmin. Info : Replaced faulty RAM module API TEST', 1, '2024-12-18 03:05:03.23'),
(200, 1075, 'PA24-0009-0018Sent to repair by superadmin. Info : Replaced faulty RAM module API TEST', 1, '2024-12-18 03:05:13.53'),
(201, 1075, 'PA24-0009-0018 sent to repair by superadmin. Info : Replaced faulty RAM module API TEST', 1, '2024-12-18 03:05:25.80'),
(202, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-23 04:23:12.08'),
(203, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-23 06:08:39.95'),
(204, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-23 06:09:27.18'),
(205, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-23 06:09:46.38'),
(206, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-30 01:18:49.43'),
(207, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-30 01:19:14.67'),
(208, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-30 01:21:53.28'),
(209, 1075, 'PA24-0009-0019 sent to repair by superadmin. Info : Replaced faulty RAM module API TEST', 1, '2024-12-30 01:30:19.45'),
(210, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-30 01:30:19.56'),
(211, 1075, 'PA24-0009-0020 sent to repair by superadmin. Info : Replaced faulty RAM module API TEST', 1, '2024-12-30 01:32:04.07'),
(212, 1075, 'Replaced faulty RAM module API test new log', 1, '2024-12-30 01:32:04.19');

-- --------------------------------------------------------

--
-- Table structure for table `rep_type_lookup`
--

CREATE TABLE `rep_type_lookup` (
  `id_rep_t` tinyint(3) NOT NULL COMMENT 'primary key',
  `rep_type` varchar(255) NOT NULL COMMENT 'type of log be it repair initiated or closed'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `rep_type_lookup`
--

INSERT INTO `rep_type_lookup` (`id_rep_t`, `rep_type`) VALUES
(1, 'Repair Started. Repair opened'),
(2, 'Repair done. Repair Closed');

-- --------------------------------------------------------

--
-- Table structure for table `status_lookup`
--

CREATE TABLE `status_lookup` (
  `id_status` tinyint(3) UNSIGNED NOT NULL COMMENT 'primary key',
  `status_name` varchar(30) NOT NULL COMMENT 'status of the unit'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `status_lookup`
--

INSERT INTO `status_lookup` (`id_status`, `status_name`) VALUES
(1, 'Available in warehouse'),
(2, 'Borrowed/Lent'),
(3, 'In-repair'),
(4, 'Broken beyond repair or Lost');

-- --------------------------------------------------------

--
-- Table structure for table `unit_log`
--

CREATE TABLE `unit_log` (
  `id_log` int(11) NOT NULL COMMENT 'primary key',
  `id_unit` int(10) NOT NULL COMMENT 'foreign key - item_unit',
  `content` varchar(255) NOT NULL COMMENT 'content of the log, text. can be comment from itemunit',
  `update_at` datetime(6) NOT NULL COMMENT 'date time of which the log is written',
  `actors_action` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `unit_log`
--

INSERT INTO `unit_log` (`id_log`, `id_unit`, `content`, `update_at`, `actors_action`) VALUES
(1, 1, 'Unit PA24-0001-0001 lent to Johan by admin_en', '2024-11-28 08:45:51.000000', ''),
(2, 129, 'Unit PA24-0002-0001 lent to Emma by admin_en', '2024-11-28 08:46:00.000000', ''),
(3, 193, 'Unit PC24-0001-0001 lent to Oleg by admin_en', '2024-11-28 08:46:08.000000', ''),
(4, 257, 'Unit PA24-0003-0001 lent to Freddy by admin_en', '2024-11-28 08:46:16.000000', ''),
(5, 321, 'Unit VD24-0001-0001 lent to Shiorin by admin_en', '2024-11-28 08:46:26.000000', ''),
(6, 385, 'Unit VD24-0002-0001 lent to Mark by admin_en', '2024-11-28 08:46:34.000000', ''),
(7, 450, 'Unit VD24-0003-0002 lent to Manfred Albrecht Freiherr von Richthofen by admin_en', '2024-11-28 08:46:43.000000', ''),
(8, 513, 'Unit PA24-0004-0001 lent to Manfred Albrecht Freiherr von Richthofen by admin_en', '2024-11-28 08:47:29.000000', ''),
(9, 577, 'Unit PC24-0002-0001 lent to Johan by admin_en', '2024-11-28 08:47:39.000000', ''),
(10, 641, 'Unit VD24-0004-0001 lent to Emma by admin_en', '2024-11-28 08:47:49.000000', ''),
(11, 705, 'Unit TC24-0001-0001 lent to Emma by admin_en', '2024-11-28 08:48:04.000000', ''),
(12, 769, 'Unit TC24-0002-0001 lent to Freddy by admin_en', '2024-11-28 08:48:14.000000', ''),
(13, 833, 'Unit TC24-0003-0001 lent to Shiorin by admin_en', '2024-11-28 08:48:24.000000', ''),
(14, 897, 'Unit PU24-0001-0001 lent to Manfred Albrecht Freiherr von Richthofen by admin_en', '2024-11-28 08:48:35.000000', ''),
(15, 961, 'Unit TC24-0004-0001 lent to Freddy by admin_en', '2024-11-28 08:48:45.000000', ''),
(16, 961, 'Unit TC24-0004-0001 lent to Freddy by admin_en', '2024-11-28 08:48:45.000000', ''),
(17, 66, 'Unit DK24-0001-0002 lent to Emma by admin_en', '2024-11-28 08:48:57.000000', ''),
(18, 1, 'Unit PA24-0001-0001 returned by Johan, recieved by admin_en', '2024-11-28 08:51:51.000000', ''),
(19, 129, 'Unit PA24-0002-0001 returned by Emma, recieved by admin_en', '2024-11-28 08:52:00.000000', ''),
(20, 193, 'Unit PC24-0001-0001 returned by Oleg, recieved by admin_en', '2024-11-28 08:52:11.000000', ''),
(21, 257, 'Unit PA24-0003-0001 returned by Freddy, recieved by admin_en', '2024-11-28 08:52:20.000000', ''),
(22, 321, 'Unit VD24-0001-0001 returned by Shiorin, recieved by admin_en', '2024-11-28 08:52:35.000000', ''),
(23, 385, 'Unit VD24-0002-0001 returned by Mark, recieved by admin_en', '2024-11-28 08:52:43.000000', ''),
(24, 129, 'Unit PA24-0002-0001 sent to repair by admin_en', '2024-11-28 08:52:58.000000', ''),
(25, 193, 'Unit PC24-0001-0001 sent to repair by admin_en', '2024-11-28 08:53:02.000000', ''),
(26, 257, 'Unit PA24-0003-0001 sent to repair by admin_en', '2024-11-28 08:53:06.000000', ''),
(27, 321, 'Unit VD24-0001-0001 sent to repair by admin_en', '2024-11-28 08:53:11.000000', ''),
(28, 385, 'Unit VD24-0002-0001 sent to repair by admin_en', '2024-11-28 08:53:16.000000', ''),
(29, 193, 'Unit PC24-0001-0001 repaired. Taken to warehouse by admin_en', '2024-11-28 08:53:29.000000', ''),
(30, 257, 'Unit PA24-0003-0001 repaired. Taken to warehouse by admin_en', '2024-11-28 08:53:39.000000', ''),
(31, 513, 'Unit PA24-0004-0001 returned by Manfred Albrecht Freiherr von Richthofen, recieved by franzferdinand', '2024-11-28 10:59:11.000000', ''),
(32, 513, 'Unit PA24-0004-0001 sent to repair by franzferdinand', '2024-11-28 10:59:19.000000', ''),
(33, 129, 'Unit PA24-0002-0001 repaired. Taken to warehouse by superadmin', '2024-11-28 14:12:46.000000', ''),
(34, 257, 'Unit PA24-0003-0001 sent to repair by superadmin', '2024-11-29 14:30:02.000000', ''),
(35, 129, 'Unit PA24-0002-0001 sent to repair by superadmin', '2024-11-29 14:30:48.000000', ''),
(36, 129, 'Unit PA24-0002-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:31:36.000000', ''),
(37, 129, 'Unit PA24-0002-0001 sent to repair by bobtherepairman', '2024-11-29 14:34:57.000000', ''),
(38, 129, 'Unit PA24-0002-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:35:47.000000', ''),
(39, 129, 'Unit PA24-0002-0001 sent to repair by franzferdinand', '2024-11-29 14:36:07.000000', ''),
(40, 129, 'Unit PA24-0002-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:36:30.000000', ''),
(41, 257, 'Unit PA24-0003-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:37:50.000000', ''),
(42, 321, 'Unit VD24-0001-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:37:56.000000', ''),
(43, 385, 'Unit VD24-0002-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:38:02.000000', ''),
(44, 385, 'Unit VD24-0002-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:38:02.000000', ''),
(45, 513, 'Unit PA24-0004-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:38:08.000000', ''),
(46, 385, 'Unit VD24-0002-0001 sent to repair by bobtherepairman', '2024-11-29 14:42:18.000000', ''),
(47, 385, 'Unit VD24-0002-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:42:29.000000', ''),
(48, 257, 'Unit PA24-0003-0001 sent to repair by bobtherepairman', '2024-11-29 14:43:21.000000', ''),
(49, 257, 'Unit PA24-0003-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:43:28.000000', ''),
(50, 257, 'Unit PA24-0003-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:43:28.000000', ''),
(51, 257, 'Unit PA24-0003-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:43:29.000000', ''),
(52, 257, 'Unit PA24-0003-0001 repaired. Taken to warehouse by bobtherepairman', '2024-11-29 14:43:29.000000', ''),
(53, 1, 'Unit PA24-0001-0001 lent to Johan by superadmin', '2024-11-29 14:51:45.000000', ''),
(54, 2, 'Unit PA24-0001-0002 lent to Emma by superadmin', '2024-11-29 14:52:07.000000', ''),
(55, 257, 'Unit PA24-0003-0001 sent to repair by superadmin', '2024-11-29 14:53:51.000000', ''),
(56, 257, 'Unit PA24-0003-0001 repaired. Taken to warehouse by superadmin', '2024-11-29 15:02:43.000000', ''),
(57, 1030, 'New unit PA24-0002-0065 added by superadmin', '2024-11-29 15:20:09.000000', ''),
(58, 130, 'Unit PA24-0002-0002 lent to Johan by superadmin', '2024-11-29 15:21:07.000000', ''),
(59, 450, 'Unit VD24-0003-0002 returned by Manfred Albrecht Freiherr von Richthofen, recieved by superadmin', '2024-11-29 15:21:27.000000', ''),
(60, 129, 'Unit PA24-0002-0001 sent to repair by superadmin', '2024-11-29 15:21:53.000000', ''),
(61, 129, 'Unit PA24-0002-0001 repaired. Taken to warehouse by superadmin', '2024-11-29 15:22:02.000000', ''),
(62, 1, 'Unit PA24-0001-0001 updated by superadmin', '2024-11-29 15:35:39.000000', ''),
(63, 1, 'Unit PA24-0001-0001 lent to Johan by superadmin', '2024-12-02 08:31:46.000000', ''),
(64, 1, 'Unit PA24-0001-0001 returned by Johan, recieved by superadmin', '2024-12-02 08:32:04.000000', ''),
(65, 129, 'Unit PA24-0002-0001 sent to repair by superadmin', '2024-12-02 08:35:43.000000', ''),
(66, 129, 'Unit PA24-0002-0001 repaired. Taken to warehouse by superadmin', '2024-12-02 08:35:51.000000', ''),
(67, 129, 'Unit PA24-0002-0001 sent to repair by superadmin', '2024-12-02 13:30:42.000000', ''),
(68, 129, 'Unit PA24-0002-0001 repaired. Taken to warehouse by superadmin', '2024-12-02 13:30:52.000000', ''),
(69, 193, 'Unit PC24-0001-0001 sent to repair by superadmin', '2024-12-03 09:43:27.000000', ''),
(70, 321, 'Unit VD24-0001-0001 sent to repair by superadmin', '2024-12-03 09:43:32.000000', ''),
(71, 129, 'Unit PA24-0002-0001 sent to repair by superadmin', '2024-12-05 08:40:43.000000', ''),
(72, 257, 'Unit PA24-0003-0001 sent to repair by superadmin', '2024-12-05 09:25:42.000000', ''),
(73, 259, 'Unit PA24-0003-0003 lent to Emma by superadmin', '2024-12-05 14:44:01.000000', ''),
(74, 193, 'Unit PC24-0001-0001 repaired. Taken to warehouse by superadmin', '2024-12-05 14:54:34.000000', ''),
(75, 193, 'Unit PC24-0001-0001 sent to repair by superadmin', '2024-12-05 14:54:45.000000', ''),
(76, 131, 'Unit PA24-0002-0003 lent to Manfred Albrecht Freiherr von Richthofen by superadmin', '2024-12-06 08:22:52.000000', ''),
(77, 131, 'Unit PA24-0002-0003 returned by Manfred Albrecht Freiherr von Richthofen, recieved by superadmin', '2024-12-06 08:26:54.000000', ''),
(78, 577, 'Unit PC24-0002-0001 returned by Johan - received by superadmin', '2024-12-06 08:28:01.000000', ''),
(79, 641, 'Unit VD24-0004-0001 returned by Emma - received by superadmin', '2024-12-06 08:29:23.000000', ''),
(80, 385, 'Unit VD24-0002-0001 sent to repair by superadmin', '2024-12-06 08:33:48.000000', ''),
(81, 129, 'Unit PA24-0002-0001 repaired. Taken to warehouse by superadmin', '2024-12-06 08:35:59.000000', ''),
(82, 193, 'Unit PC24-0001-0001 repaired. Taken to warehouse by superadmin', '2024-12-06 08:36:51.000000', ''),
(83, 961, 'Unit TC24-0004-0001 returned by Freddy - received by superadmin', '2024-12-06 08:44:30.000000', ''),
(84, 1, 'Unit PA24-0001-0001 updated by superadmin', '2024-12-06 08:47:25.000000', ''),
(85, 257, 'Unit PA24-0003-0001 repaired and taken to warehouse by superadmin', '2024-12-06 08:48:23.000000', ''),
(86, 193, 'Unit PC24-0001-0001 sent to repair by superadmin', '2024-12-06 08:48:33.000000', ''),
(87, 129, 'Unit PA24-0002-0001 sent to repair by superadmin', '2024-12-06 08:48:46.000000', ''),
(88, 129, 'Unit PA24-0002-0001 repaired and taken to warehouse by superadmin', '2024-12-06 08:48:53.000000', ''),
(89, 193, 'Unit PC24-0001-0001 repaired and taken to warehouse by superadmin', '2024-12-06 08:48:57.000000', ''),
(90, 1042, 'New unit TC24-0006-0001 added by superadmin', '2024-12-06 08:52:20.000000', ''),
(91, 1049, 'New unit TC24-0006-0008 added by superadmin', '2024-12-06 09:01:25.000000', ''),
(92, 1050, 'New unit TC24-0006-0009 added by superadmin', '2024-12-06 09:01:25.000000', ''),
(93, 1051, 'New unit TC24-0006-0010 added by superadmin', '2024-12-06 09:01:25.000000', ''),
(94, 1052, 'New unit TC24-0006-0011 added by superadmin', '2024-12-06 09:01:25.000000', ''),
(95, 1053, 'New unit TC24-0006-0012 added by superadmin', '2024-12-06 09:01:25.000000', ''),
(96, 1054, 'New unit TC24-0006-0013 added by superadmin', '2024-12-06 09:01:26.000000', ''),
(97, 513, 'Unit PA24-0004-0001 sent to repair by superadmin', '2024-12-06 09:18:00.000000', ''),
(98, 257, 'Unit PA24-0003-0001 sent to repair by superadmin', '2024-12-06 09:26:49.000000', ''),
(99, 193, 'Unit PC24-0001-0001 sent to repair by superadmin', '2024-12-06 09:31:32.000000', ''),
(100, 257, 'Unit PA24-0003-0001 repaired and taken to warehouse by superadmin', '2024-12-06 09:33:21.000000', ''),
(101, 193, 'Unit PC24-0001-0001 repaired and taken to warehouse by superadmin', '2024-12-06 08:33:48.000000', ''),
(102, 321, 'Unit VD24-0001-0001 repaired and taken to warehouse by superadmin', '2024-12-06 08:34:15.000000', ''),
(103, 1042, 'Unit TC24-0006-0001 updated by superadmin', '2024-12-06 08:35:12.000000', ''),
(104, 1, 'Unit PA24-0001-0001 updated by superadmin', '2024-12-06 09:36:36.000000', ''),
(105, 321, 'Unit VD24-0001-0001 sent to repair by superadmin', '2024-12-06 09:39:43.000000', ''),
(106, 1055, 'New unit TC24-0005-0001 added by superadmin', '2024-12-06 12:45:51.000000', ''),
(107, 1056, 'New unit TC24-0005-0002 added by superadmin', '2024-12-06 12:45:51.000000', ''),
(108, 1057, 'New unit TC24-0005-0003 added by superadmin', '2024-12-06 12:45:51.000000', ''),
(109, 1058, 'New unit TC24-0005-0004 added by superadmin', '2024-12-06 12:45:51.000000', ''),
(110, 1059, 'New unit TC24-0005-0005 added by superadmin', '2024-12-06 12:45:51.000000', ''),
(111, 1060, 'New unit TC24-0005-0006 added by superadmin', '2024-12-06 12:45:51.000000', ''),
(112, 3, 'Unit PA24-0001-0003 lent to Freddy by superadmin', '2024-12-09 07:41:53.000000', ''),
(113, 3, 'Unit PA24-0001-0003 returned by Freddy - received by superadmin', '2024-12-09 07:46:38.000000', ''),
(114, 321, 'Unit VD24-0001-0001 repaired and taken to warehouse by superadmin', '2024-12-09 08:00:50.000000', ''),
(115, 385, 'Unit VD24-0002-0001 repaired and taken to warehouse by superadmin', '2024-12-09 08:00:54.000000', ''),
(116, 513, 'Unit PA24-0004-0001 repaired and taken to warehouse by superadmin', '2024-12-09 08:00:58.000000', ''),
(117, 4, 'Unit PA24-0001-0004 lent to Oleg by superadmin', '2024-12-09 13:07:54.000000', ''),
(119, 2, 'Battery replacement performed on laptop', '2024-12-13 07:27:51.158183', ''),
(120, 2, 'Battery replacement performed on laptop', '2024-12-16 01:54:20.333420', ''),
(121, 2, 'Battery replacement performed on laptop', '2024-12-16 01:55:09.449124', ''),
(122, 81, 'Battery replacement performed on laptop', '2024-12-16 01:55:15.614582', ''),
(123, 81, 'Battery replacement performed on laptop', '2024-12-16 01:56:17.304349', ''),
(124, 81, 'Battery replacement performed on laptop', '2024-12-16 01:57:48.223942', ''),
(125, 81, 'Battery replacement performed on laptop', '2024-12-16 02:00:28.588893', ''),
(126, 81, 'Battery replacement performed on laptop', '2024-12-16 02:11:31.571526', 'Unit Sent to repair by USER'),
(127, 1067, 'api test create new', '2024-12-16 02:33:19.994335', 'New Unit added by admin_en into Wonosobo Office Wh'),
(128, 1075, 'api test create new', '2024-12-16 03:02:11.818589', 'New Unit added by admin_en into Wonosobo Office Wh'),
(129, 1075, 'Updated unit comment api test 3', '2024-12-16 03:23:28.129390', 'Unit  updated by admin_en into Wonosobo Office Wh'),
(130, 1074, 'Unit DK24-0001-91111 lent out by superadmin', '2024-12-16 03:41:50.025435', 'Unit DK24-0001-91111 lent outEmmaby superadmin'),
(131, 1075, 'Unit DK24-0001-91112 lent out by superadmin', '2024-12-16 04:18:57.511099', 'Unit DK24-0001-91112 lent out Emma by superadmin'),
(132, 1075, 'Unit DK24-0001-91112 lent out by superadmin', '2024-12-16 04:20:36.778948', 'Unit DK24-0001-91112 lent out Emma by superadmin'),
(133, 1075, 'Unit DK24-0001-91112 lent out by superadmin', '2024-12-16 04:22:05.898450', 'Unit DK24-0001-91112 lent out to Emma returned toPWT Office Wh by superadmin'),
(134, 1075, 'Unit sent to repair', '2024-12-16 04:41:13.863603', 'Unit DK24-0001-91112 sent to repair by superadmin. Info : Replaced faulty RAM module'),
(135, 1075, 'Unit sent to repair', '2024-12-16 04:41:55.945825', 'Unit DK24-0001-91112 sent to repair by superadmin. Info : Replaced faulty RAM module API'),
(136, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-16 06:03:57.122127', 'Unit DK24-0001-91112 sent to repair by superadmin'),
(137, 1075, 'Unit sent repaired. Info : Replaced faulty RAM module API', '2024-12-16 06:15:20.010624', 'Unit DK24-0001-91112 sent to repair by superadmin'),
(138, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-16 06:17:07.249413', 'Unit DK24-0001-91112 returned to PWT Office Wh by superadmin'),
(139, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-16 06:18:11.073012', 'Unit DK24-0001-91112 returned to JKT Headquarter by superadmin'),
(140, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-16 06:29:23.983933', 'Unit DK24-0001-91112 returned to JKT Headquarter by superadmin'),
(141, 1075, 'Updated unit comment api test 3', '2024-12-16 06:35:38.033083', 'Unit  updated by admin_en in Wonosobo Office Wh'),
(142, 1075, 'Unit lent out', '2024-12-16 06:35:39.565082', 'Unit DK24-0001-91112 lent out Emma by superadmin'),
(143, 1075, 'api test return with update to item unit ', '2024-12-16 06:35:39.742810', 'Unit DK24-0001-91112 lent out to Emma returned toPWT Office Wh by superadmin'),
(144, 81, 'Battery replacement performed on laptop', '2024-12-16 06:35:40.635773', 'Unit Sent to repair by USER'),
(145, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-16 06:35:41.121965', 'Unit DK24-0001-91112 sent to repair by superadmin'),
(146, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-16 06:35:41.228793', 'Unit DK24-0001-91112 returned to JKT Headquarter by superadmin'),
(147, 1077, 'api test create new', '2024-12-16 07:40:38.206068', 'New Unit added by superadmin into Wonosobo Office Wh'),
(148, 1078, 'api test create new 2', '2024-12-16 07:40:50.854355', 'New Unit added by superadmin into Wonosobo Office Wh'),
(149, 1079, 'api test create new 3', '2024-12-16 07:40:56.098976', 'New Unit added by superadmin into Wonosobo Office Wh'),
(150, 1079, 'Updated unit comment api test 3 EDIT', '2024-12-16 07:44:23.007133', 'Unit  updated by admin_en in Wonosobo Office Wh'),
(151, 1075, 'Updated unit comment api test 3 EDIT', '2024-12-16 07:45:20.589135', 'Unit  updated by admin_en in Wonosobo Office Wh'),
(152, 1075, 'Updated unit comment api test AA EDIT', '2024-12-16 08:10:25.655831', 'Unit PA24-0009-0004 updated by admin_en in Wonosobo Office Wh'),
(153, 1080, 'api test create new G11', '2024-12-17 01:29:09.239547', 'New Unit added by superadmin into Wonosobo Office Wh'),
(154, 1081, 'api test create new G11', '2024-12-17 01:29:16.118893', 'New Unit added by superadmin into Wonosobo Office Wh'),
(155, 1082, 'api test create new G11', '2024-12-17 01:29:18.990037', 'New Unit added by superadmin into Wonosobo Office Wh'),
(156, 1083, 'api test create new G11', '2024-12-17 01:29:23.473513', 'New Unit added by superadmin into Wonosobo Office Wh'),
(157, 1084, 'api test create new G11', '2024-12-17 01:29:44.025840', 'New Unit added by superadmin into Wonosobo Office Wh'),
(158, 1085, 'api test create new G11', '2024-12-17 01:29:45.171156', 'New Unit added by superadmin into Wonosobo Office Wh'),
(159, 1086, 'api test create new G11', '2024-12-17 01:29:46.293913', 'New Unit added by superadmin into Wonosobo Office Wh'),
(160, 1087, 'api test create new G11', '2024-12-17 01:29:47.446089', 'New Unit added by superadmin into Wonosobo Office Wh'),
(161, 1088, 'api test create new G11', '2024-12-17 01:30:08.521124', 'New Unit added by superadmin into Wonosobo Office Wh'),
(162, 1089, 'api test create new G11', '2024-12-17 01:30:10.066275', 'New Unit added by superadmin into Wonosobo Office Wh'),
(163, 1090, 'api test create new G11', '2024-12-17 01:30:12.110093', 'New Unit added by superadmin into Wonosobo Office Wh'),
(164, 1091, 'api test create new G11', '2024-12-17 01:32:08.384168', 'New Unit added by superadmin into Wonosobo Office Wh'),
(165, 1092, 'api test create new G11', '2024-12-17 01:32:08.530860', 'New Unit added by superadmin into Wonosobo Office Wh'),
(166, 1093, 'api test create new G11', '2024-12-17 01:32:08.627567', 'New Unit added by superadmin into Wonosobo Office Wh'),
(167, 1094, 'api test create new G11', '2024-12-17 01:32:08.694160', 'New Unit added by superadmin into Wonosobo Office Wh'),
(168, 1095, 'api test create new G11', '2024-12-17 01:32:08.777626', 'New Unit added by superadmin into Wonosobo Office Wh'),
(169, 1096, 'api test create new G11', '2024-12-17 01:32:08.862161', 'New Unit added by superadmin into Wonosobo Office Wh'),
(170, 1097, 'api test create new G11', '2024-12-17 01:32:08.970848', 'New Unit added by superadmin into Wonosobo Office Wh'),
(171, 1098, 'api test create new G11', '2024-12-17 01:32:09.068881', 'New Unit added by superadmin into Wonosobo Office Wh'),
(172, 1099, 'api test create new G11', '2024-12-17 01:32:09.163801', 'New Unit added by superadmin into Wonosobo Office Wh'),
(173, 1100, 'api test create new G11', '2024-12-17 01:32:09.241000', 'New Unit added by superadmin into Wonosobo Office Wh'),
(174, 1101, 'api test create new G11', '2024-12-17 01:32:09.329051', 'New Unit added by superadmin into Wonosobo Office Wh'),
(175, 1102, 'api test create new G11', '2024-12-17 01:32:09.409372', 'New Unit added by superadmin into Wonosobo Office Wh'),
(176, 1103, 'api test create new G11', '2024-12-17 01:32:09.484354', 'New Unit added by superadmin into Wonosobo Office Wh'),
(177, 1104, 'api test create new G11', '2024-12-17 01:32:09.563863', 'New Unit added by superadmin into Wonosobo Office Wh'),
(178, 1105, 'api test create new G11', '2024-12-17 01:32:09.662512', 'New Unit added by superadmin into Wonosobo Office Wh'),
(179, 1106, 'api test create new G11', '2024-12-17 01:32:09.773050', 'New Unit added by superadmin into Wonosobo Office Wh'),
(180, 1107, 'api test create new G11', '2024-12-17 01:32:09.855200', 'New Unit added by superadmin into Wonosobo Office Wh'),
(181, 1108, 'api test create new G11', '2024-12-17 01:32:09.944220', 'New Unit added by superadmin into Wonosobo Office Wh'),
(182, 1109, 'api test create new G11', '2024-12-17 01:32:10.023448', 'New Unit added by superadmin into Wonosobo Office Wh'),
(183, 1110, 'api test create new G11', '2024-12-17 01:32:10.119163', 'New Unit added by superadmin into Wonosobo Office Wh'),
(184, 1111, 'api test create new G11', '2024-12-17 01:32:10.217549', 'New Unit added by superadmin into Wonosobo Office Wh'),
(185, 1112, 'api test create new G11', '2024-12-17 01:32:10.322880', 'New Unit added by superadmin into Wonosobo Office Wh'),
(186, 1113, 'api test create new G11', '2024-12-17 01:32:10.428801', 'New Unit added by superadmin into Wonosobo Office Wh'),
(187, 1114, 'api test create new G11', '2024-12-17 01:32:10.538139', 'New Unit added by superadmin into Wonosobo Office Wh'),
(188, 1115, 'api test create new G11', '2024-12-17 01:32:10.619302', 'New Unit added by superadmin into Wonosobo Office Wh'),
(189, 1116, 'api test create new G11', '2024-12-17 01:32:10.709329', 'New Unit added by superadmin into Wonosobo Office Wh'),
(190, 1117, 'api test create new G11', '2024-12-17 01:32:10.775327', 'New Unit added by superadmin into Wonosobo Office Wh'),
(191, 1118, 'api test create new G11', '2024-12-17 01:32:10.886500', 'New Unit added by superadmin into Wonosobo Office Wh'),
(192, 1119, 'api test create new G11', '2024-12-17 01:32:10.965952', 'New Unit added by superadmin into Wonosobo Office Wh'),
(193, 1120, 'api test create new G11', '2024-12-17 01:32:11.046136', 'New Unit added by superadmin into Wonosobo Office Wh'),
(194, 1121, 'api test create new G11', '2024-12-17 01:32:11.126681', 'New Unit added by superadmin into Wonosobo Office Wh'),
(195, 1122, 'api test create new G11', '2024-12-17 01:32:11.240266', 'New Unit added by superadmin into Wonosobo Office Wh'),
(196, 1123, 'api test create new G11', '2024-12-17 01:32:11.386866', 'New Unit added by superadmin into Wonosobo Office Wh'),
(197, 1124, 'api test create new G11', '2024-12-17 01:32:11.482447', 'New Unit added by superadmin into Wonosobo Office Wh'),
(198, 1125, 'api test create new G11', '2024-12-17 01:32:11.564431', 'New Unit added by superadmin into Wonosobo Office Wh'),
(199, 1126, 'api test create new G11', '2024-12-17 01:32:11.656684', 'New Unit added by superadmin into Wonosobo Office Wh'),
(200, 1127, 'api test create new G11', '2024-12-17 01:32:11.769854', 'New Unit added by superadmin into Wonosobo Office Wh'),
(201, 1128, 'api test create new G11', '2024-12-17 01:32:11.860827', 'New Unit added by superadmin into Wonosobo Office Wh'),
(202, 1129, 'api test create new G11', '2024-12-17 01:32:11.958112', 'New Unit added by superadmin into Wonosobo Office Wh'),
(203, 1130, 'api test create new G11', '2024-12-17 01:32:12.094976', 'New Unit added by superadmin into Wonosobo Office Wh'),
(204, 1131, 'api test create new G11', '2024-12-17 01:32:12.174088', 'New Unit added by superadmin into Wonosobo Office Wh'),
(205, 1132, 'api test create new G11', '2024-12-17 01:32:12.563013', 'New Unit added by superadmin into Wonosobo Office Wh'),
(206, 1133, 'api test create new G11', '2024-12-17 01:32:12.684845', 'New Unit added by superadmin into Wonosobo Office Wh'),
(207, 1134, 'api test create new G11', '2024-12-17 01:32:12.766667', 'New Unit added by superadmin into Wonosobo Office Wh'),
(208, 1135, 'api test create new G11', '2024-12-17 01:32:12.846537', 'New Unit added by superadmin into Wonosobo Office Wh'),
(209, 1136, 'api test create new G11', '2024-12-17 01:32:12.928337', 'New Unit added by superadmin into Wonosobo Office Wh'),
(210, 1137, 'api test create new G11', '2024-12-17 01:32:13.042240', 'New Unit added by superadmin into Wonosobo Office Wh'),
(211, 1138, 'api test create new G11', '2024-12-17 01:32:13.116154', 'New Unit added by superadmin into Wonosobo Office Wh'),
(212, 1139, 'api test create new G11', '2024-12-17 01:32:13.211509', 'New Unit added by superadmin into Wonosobo Office Wh'),
(213, 1140, 'api test create new G11', '2024-12-17 01:32:13.323683', 'New Unit added by superadmin into Wonosobo Office Wh'),
(214, 1141, 'api test create new G11', '2024-12-17 01:32:13.419654', 'New Unit added by superadmin into Wonosobo Office Wh'),
(215, 1142, 'api test create new G11', '2024-12-17 01:32:13.528537', 'New Unit added by superadmin into Wonosobo Office Wh'),
(216, 1143, 'api test create new G11', '2024-12-17 01:32:13.640012', 'New Unit added by superadmin into Wonosobo Office Wh'),
(217, 1144, 'api test create new G11', '2024-12-17 01:32:13.719977', 'New Unit added by superadmin into Wonosobo Office Wh'),
(218, 1145, 'api test create new G11', '2024-12-17 01:32:13.812048', 'New Unit added by superadmin into Wonosobo Office Wh'),
(219, 1146, 'api test create new G11', '2024-12-17 01:32:13.912881', 'New Unit added by superadmin into Wonosobo Office Wh'),
(220, 1147, 'api test create new G11', '2024-12-17 01:32:14.023312', 'New Unit added by superadmin into Wonosobo Office Wh'),
(221, 1148, 'api test create new G11', '2024-12-17 01:32:14.113131', 'New Unit added by superadmin into Wonosobo Office Wh'),
(222, 1149, 'api test create new G11', '2024-12-17 01:32:14.207210', 'New Unit added by superadmin into Wonosobo Office Wh'),
(223, 1150, 'api test create new G11', '2024-12-17 01:32:14.288936', 'New Unit added by superadmin into Wonosobo Office Wh'),
(224, 1151, 'api test create new G11', '2024-12-17 01:32:14.382633', 'New Unit added by superadmin into Wonosobo Office Wh'),
(225, 1152, 'api test create new G11', '2024-12-17 01:32:14.491748', 'New Unit added by superadmin into Wonosobo Office Wh'),
(226, 1153, 'api test create new G11', '2024-12-17 01:32:14.575773', 'New Unit added by superadmin into Wonosobo Office Wh'),
(227, 1154, 'api test create new G11', '2024-12-17 01:32:14.661404', 'New Unit added by superadmin into Wonosobo Office Wh'),
(228, 1155, 'api test create new G11', '2024-12-17 01:40:42.551915', 'New UnitPA24-0011-0075added by superadmin into Wonosobo Office Wh'),
(229, 1156, 'api test create new G11', '2024-12-17 01:41:08.275878', 'New UnitPA24-0011-0076added by superadmin into Wonosobo Office Wh'),
(230, 1157, 'api test create new G11', '2024-12-17 01:41:42.937959', 'New Unit PA24-0011-0077 added by superadmin into Wonosobo Office Wh'),
(231, 1158, 'api test create new G11', '2024-12-17 03:41:28.078096', 'New Unit PA24-0011-0078 added by superadmin into Wonosobo Office Wh'),
(232, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:41:28.266519', 'Unit PA24-0009-0005 updated by admin_en in Wonosobo Office Wh'),
(233, 1075, 'Unit lent out', '2024-12-17 03:41:29.759283', 'Unit PA24-0009-0005 lent out Emma by superadmin'),
(234, 1075, 'api test return with update to item unit ', '2024-12-17 03:41:29.934532', 'Unit PA24-0009-0005 lent out to Emma returned toPWT Office Wh by superadmin'),
(235, 81, 'Battery replacement performed on laptop', '2024-12-17 03:41:30.984150', 'Unit Sent to repair by USER'),
(236, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:41:31.462932', 'Unit PA24-0009-0005 sent to repair by superadmin'),
(237, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:41:31.565462', 'Unit PA24-0009-0005 returned to JKT Headquarter by superadmin'),
(238, 1159, 'api test create new G11', '2024-12-17 03:41:33.347631', 'New Unit PA24-0011-0079 added by superadmin into Wonosobo Office Wh'),
(239, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:41:33.496586', 'Unit PA24-0009-0006 updated by admin_en in Wonosobo Office Wh'),
(240, 1075, 'Unit lent out', '2024-12-17 03:41:35.163619', 'Unit PA24-0009-0006 lent out Emma by superadmin'),
(241, 1075, 'api test return with update to item unit ', '2024-12-17 03:41:35.331266', 'Unit PA24-0009-0006 lent out to Emma returned toPWT Office Wh by superadmin'),
(242, 81, 'Battery replacement performed on laptop', '2024-12-17 03:41:36.222682', 'Unit Sent to repair by USER'),
(243, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:41:36.663597', 'Unit PA24-0009-0006 sent to repair by superadmin'),
(244, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:41:36.773087', 'Unit PA24-0009-0006 returned to JKT Headquarter by superadmin'),
(245, 1160, 'api test create new G11', '2024-12-17 03:41:38.713613', 'New Unit PA24-0011-0080 added by superadmin into Wonosobo Office Wh'),
(246, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:41:38.898124', 'Unit PA24-0009-0007 updated by admin_en in Wonosobo Office Wh'),
(247, 1075, 'Unit lent out', '2024-12-17 03:41:40.480990', 'Unit PA24-0009-0007 lent out Emma by superadmin'),
(248, 1075, 'api test return with update to item unit ', '2024-12-17 03:41:40.655475', 'Unit PA24-0009-0007 lent out to Emma returned toPWT Office Wh by superadmin'),
(249, 81, 'Battery replacement performed on laptop', '2024-12-17 03:41:41.569857', 'Unit Sent to repair by USER'),
(250, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:41:42.373233', 'Unit PA24-0009-0007 sent to repair by superadmin'),
(251, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:41:42.485806', 'Unit PA24-0009-0007 returned to JKT Headquarter by superadmin'),
(252, 1161, 'api test create new G11', '2024-12-17 03:41:44.219709', 'New Unit PA24-0011-0081 added by superadmin into Wonosobo Office Wh'),
(253, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:41:44.375412', 'Unit PA24-0009-0008 updated by admin_en in Wonosobo Office Wh'),
(254, 1075, 'Unit lent out', '2024-12-17 03:41:45.908532', 'Unit PA24-0009-0008 lent out Emma by superadmin'),
(255, 1075, 'api test return with update to item unit ', '2024-12-17 03:41:46.117567', 'Unit PA24-0009-0008 lent out to Emma returned toPWT Office Wh by superadmin'),
(256, 81, 'Battery replacement performed on laptop', '2024-12-17 03:41:47.046612', 'Unit Sent to repair by USER'),
(257, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:41:47.524223', 'Unit PA24-0009-0008 sent to repair by superadmin'),
(258, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:41:47.637436', 'Unit PA24-0009-0008 returned to JKT Headquarter by superadmin'),
(259, 1162, 'api test create new G11', '2024-12-17 03:41:49.458886', 'New Unit PA24-0011-0082 added by superadmin into Wonosobo Office Wh'),
(260, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:41:49.654524', 'Unit PA24-0009-0009 updated by admin_en in Wonosobo Office Wh'),
(261, 1075, 'Unit lent out', '2024-12-17 03:41:51.361131', 'Unit PA24-0009-0009 lent out Emma by superadmin'),
(262, 1075, 'api test return with update to item unit ', '2024-12-17 03:41:51.561657', 'Unit PA24-0009-0009 lent out to Emma returned toPWT Office Wh by superadmin'),
(263, 81, 'Battery replacement performed on laptop', '2024-12-17 03:41:52.522108', 'Unit Sent to repair by USER'),
(264, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:41:52.994274', 'Unit PA24-0009-0009 sent to repair by superadmin'),
(265, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:41:53.106003', 'Unit PA24-0009-0009 returned to JKT Headquarter by superadmin'),
(266, 1163, 'api test create new G11', '2024-12-17 03:41:55.080358', 'New Unit PA24-0011-0083 added by superadmin into Wonosobo Office Wh'),
(267, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:41:55.279569', 'Unit PA24-0009-0010 updated by admin_en in Wonosobo Office Wh'),
(268, 1075, 'Unit lent out', '2024-12-17 03:41:57.080909', 'Unit PA24-0009-0010 lent out Emma by superadmin'),
(269, 1075, 'api test return with update to item unit ', '2024-12-17 03:41:57.286348', 'Unit PA24-0009-0010 lent out to Emma returned toPWT Office Wh by superadmin'),
(270, 81, 'Battery replacement performed on laptop', '2024-12-17 03:41:58.185735', 'Unit Sent to repair by USER'),
(271, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:41:58.633732', 'Unit PA24-0009-0010 sent to repair by superadmin'),
(272, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:41:58.715026', 'Unit PA24-0009-0010 returned to JKT Headquarter by superadmin'),
(273, 1164, 'api test create new G11', '2024-12-17 03:42:00.557127', 'New Unit PA24-0011-0084 added by superadmin into Wonosobo Office Wh'),
(274, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:42:00.760514', 'Unit PA24-0009-0011 updated by admin_en in Wonosobo Office Wh'),
(275, 1075, 'Unit lent out', '2024-12-17 03:42:02.277738', 'Unit PA24-0009-0011 lent out Emma by superadmin'),
(276, 1075, 'api test return with update to item unit ', '2024-12-17 03:42:02.498396', 'Unit PA24-0009-0011 lent out to Emma returned toPWT Office Wh by superadmin'),
(277, 81, 'Battery replacement performed on laptop', '2024-12-17 03:42:03.499477', 'Unit Sent to repair by USER'),
(278, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:42:03.919671', 'Unit PA24-0009-0011 sent to repair by superadmin'),
(279, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:42:04.032123', 'Unit PA24-0009-0011 returned to JKT Headquarter by superadmin'),
(280, 1165, 'api test create new G11', '2024-12-17 03:42:05.869330', 'New Unit PA24-0011-0085 added by superadmin into Wonosobo Office Wh'),
(281, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:42:06.056550', 'Unit PA24-0009-0012 updated by admin_en in Wonosobo Office Wh'),
(282, 1075, 'Unit lent out', '2024-12-17 03:42:07.678269', 'Unit PA24-0009-0012 lent out Emma by superadmin'),
(283, 1075, 'api test return with update to item unit ', '2024-12-17 03:42:07.818065', 'Unit PA24-0009-0012 lent out to Emma returned toPWT Office Wh by superadmin'),
(284, 81, 'Battery replacement performed on laptop', '2024-12-17 03:42:08.789986', 'Unit Sent to repair by USER'),
(285, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:42:09.224285', 'Unit PA24-0009-0012 sent to repair by superadmin'),
(286, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:42:09.317815', 'Unit PA24-0009-0012 returned to JKT Headquarter by superadmin'),
(287, 1166, 'api test create new G11', '2024-12-17 03:42:11.431146', 'New Unit PA24-0011-0086 added by superadmin into Wonosobo Office Wh'),
(288, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:42:11.632481', 'Unit PA24-0009-0013 updated by admin_en in Wonosobo Office Wh'),
(289, 1075, 'Unit lent out', '2024-12-17 03:42:13.169809', 'Unit PA24-0009-0013 lent out Emma by superadmin'),
(290, 1075, 'api test return with update to item unit ', '2024-12-17 03:42:13.327561', 'Unit PA24-0009-0013 lent out to Emma returned toPWT Office Wh by superadmin'),
(291, 81, 'Battery replacement performed on laptop', '2024-12-17 03:42:14.346217', 'Unit Sent to repair by USER'),
(292, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:42:14.706251', 'Unit PA24-0009-0013 sent to repair by superadmin'),
(293, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:42:14.805445', 'Unit PA24-0009-0013 returned to JKT Headquarter by superadmin'),
(294, 1167, 'api test create new G11', '2024-12-17 03:42:16.798216', 'New Unit PA24-0011-0087 added by superadmin into Wonosobo Office Wh'),
(295, 1075, 'Updated unit comment api test AA EDIT', '2024-12-17 03:42:16.967406', 'Unit PA24-0009-0014 updated by admin_en in Wonosobo Office Wh'),
(296, 1075, 'Unit lent out', '2024-12-17 03:42:18.475464', 'Unit PA24-0009-0014 lent out Emma by superadmin'),
(297, 1075, 'api test return with update to item unit ', '2024-12-17 03:42:18.659794', 'Unit PA24-0009-0014 lent out to Emma returned toPWT Office Wh by superadmin'),
(298, 81, 'Battery replacement performed on laptop', '2024-12-17 03:42:19.635793', 'Unit Sent to repair by USER'),
(299, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-17 03:42:20.128067', 'Unit PA24-0009-0014 sent to repair by superadmin'),
(300, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-17 03:42:20.251813', 'Unit PA24-0009-0014 returned to JKT Headquarter by superadmin'),
(301, 1168, 'api test create new G12', '2024-12-17 03:49:13.806481', 'New Unit PA24-0011-0088 added by superadmin into Wonosobo Office Wh'),
(302, 1171, 'api test create new G12', '2024-12-17 04:08:48.869992', 'New Unit PA24-0011-0089 added by superadmin into Wonosobo Office Wh'),
(303, 1171, 'Unit lent out', '2024-12-17 06:47:20.371190', 'Unit PA24-0011-0089 lent out Oleg by superadmin'),
(304, 1075, 'test comment for return with pic', '2024-12-17 07:04:42.483155', 'Unit PA24-0009-0014 lent out to Emma returned toWonosobo Office Wh by superadmin'),
(305, 1171, 'test comment for return with pic', '2024-12-17 07:05:20.873728', 'Unit PA24-0011-0089 lent out to Oleg returned toWonosobo Office Wh by superadmin'),
(306, 1075, 'Updated unit comment api test AA EDIT', '2024-12-18 02:24:01.669153', 'Unit PA24-0009-0015 updated by admin_en in Wonosobo Office Wh'),
(307, 1171, 'test comment for return with pic', '2024-12-18 02:24:06.748688', 'Unit PA24-0011-0089 lent out to Oleg returned to Wonosobo Office Wh by superadmin'),
(308, 81, 'Battery replacement performed on laptop', '2024-12-18 02:24:07.921321', 'Unit Sent to repair by USER'),
(309, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-18 02:24:08.362927', 'Unit PA24-0009-0015 sent to repair by superadmin'),
(310, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-18 02:24:08.474304', 'Unit PA24-0009-0015 returned to JKT Headquarter by superadmin'),
(311, 1075, 'Updated unit comment api test AA EDIT', '2024-12-18 02:25:58.891037', 'Unit PA24-0009-0016 updated by admin_en in Wonosobo Office Wh'),
(312, 1171, 'test comment for return with pic', '2024-12-18 02:26:00.899813', 'Unit PA24-0011-0089 lent out to Oleg returned to Wonosobo Office Wh by superadmin'),
(313, 81, 'Battery replacement performed on laptop', '2024-12-18 02:26:01.978196', 'Unit Sent to repair by USER'),
(314, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-18 02:26:02.425463', 'Unit PA24-0009-0016 sent to repair by superadmin'),
(315, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-18 02:26:02.558944', 'Unit PA24-0009-0016 returned to JKT Headquarter by superadmin'),
(316, 1075, 'Updated unit comment api test AA EDIT', '2024-12-18 02:28:20.064189', 'Unit PA24-0009-0017 updated by admin_en in Wonosobo Office Wh'),
(317, 1075, 'Updated unit comment api test AA EDIT', '2024-12-18 02:28:36.831720', 'Unit PA24-0009-0018 updated by admin_en in Wonosobo Office Wh'),
(318, 1171, 'test comment for return with pic', '2024-12-18 02:28:38.850831', 'Unit PA24-0011-0089 lent out to Oleg returned to Wonosobo Office Wh by superadmin'),
(319, 81, 'Battery replacement performed on laptop', '2024-12-18 02:28:39.850717', 'Unit Sent to repair by USER'),
(320, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API', '2024-12-18 02:28:40.307807', 'Unit PA24-0009-0018 sent to repair by superadmin'),
(321, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-18 02:28:40.415780', 'Unit PA24-0009-0018 returned to JKT Headquarter by superadmin'),
(322, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API TEST', '2024-12-18 03:03:37.060333', 'Unit PA24-0009-0018 sent to repair by superadmin'),
(323, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API TEST', '2024-12-18 03:03:54.339417', 'Unit PA24-0009-0018 sent to repair by superadmin'),
(324, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API TEST', '2024-12-18 03:05:03.247413', 'Unit PA24-0009-0018 sent to repair by superadmin'),
(325, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API TEST', '2024-12-18 03:05:13.873119', 'Unit PA24-0009-0018 sent to repair by superadmin'),
(326, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API TEST', '2024-12-18 03:05:25.819952', 'Unit PA24-0009-0018 sent to repair by superadmin'),
(327, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-23 04:23:12.100486', 'Unit PA24-0009-0018 returned to JKT Headquarter by superadmin'),
(328, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-23 06:08:39.964458', 'Unit PA24-0009-0018 returned to JKT Headquarter by superadmin'),
(329, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-23 06:09:27.197723', 'Unit PA24-0009-0018 returned to JKT Headquarter by superadmin'),
(330, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-23 06:09:46.396531', 'Unit PA24-0009-0018 returned to JKT Headquarter by superadmin'),
(331, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-30 01:18:49.448653', 'Unit PA24-0009-0018 returned to JKT Headquarter by superadmin'),
(332, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-30 01:19:14.683566', 'Unit PA24-0009-0018 returned to JKT Headquarter by superadmin'),
(333, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-30 01:21:53.289853', 'Unit PA24-0009-0018 returned to JKT Headquarter by superadmin'),
(334, 1075, 'Updated unit comment api test AA EDIT', '2024-12-30 01:30:15.880511', 'Unit PA24-0009-0019 updated by admin_en in Wonosobo Office Wh'),
(335, 1171, 'test comment for return with pic', '2024-12-30 01:30:17.935581', 'Unit PA24-0011-0089 lent out to Oleg returned to Wonosobo Office Wh by superadmin'),
(336, 81, 'Battery replacement performed on laptop', '2024-12-30 01:30:18.995509', 'Unit Sent to repair by USER'),
(337, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API TEST', '2024-12-30 01:30:19.465487', 'Unit PA24-0009-0019 sent to repair by superadmin'),
(338, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-30 01:30:19.583011', 'Unit PA24-0009-0019 returned to JKT Headquarter by superadmin'),
(339, 1075, 'Updated unit comment api test AA EDIT', '2024-12-30 01:32:00.286756', 'Unit PA24-0009-0020 updated by admin_en in Wonosobo Office Wh'),
(340, 1171, 'test comment for return with pic', '2024-12-30 01:32:02.557357', 'Unit PA24-0011-0089 lent out to Oleg returned to Wonosobo Office Wh by superadmin'),
(341, 81, 'Battery replacement performed on laptop', '2024-12-30 01:32:03.586393', 'Unit Sent to repair by USER'),
(342, 1075, 'Unit sent to repair. Info : Replaced faulty RAM module API TEST', '2024-12-30 01:32:04.084642', 'Unit PA24-0009-0020 sent to repair by superadmin'),
(343, 1075, 'Unit repaired. Info: Replaced faulty RAM module API test new log', '2024-12-30 01:32:04.204058', 'Unit PA24-0009-0020 returned to JKT Headquarter by superadmin');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `status` int(11) NOT NULL DEFAULT 1,
  `superadmin` smallint(6) DEFAULT 0,
  `created_at` datetime(2) NOT NULL,
  `updated_at` datetime(2) NOT NULL,
  `registration_ip` varchar(15) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `auth_key` varchar(255) DEFAULT NULL,
  `bind_to_ip` varchar(255) DEFAULT NULL,
  `email_confirmed` int(11) NOT NULL,
  `confirmation_token` varchar(255) DEFAULT NULL,
  `id_wh` int(10) UNSIGNED DEFAULT NULL COMMENT 'assigned warehouse if this account is a warehouse admin',
  `user_lang` varchar(3) DEFAULT NULL COMMENT 'preffered ui language'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `username`, `password_hash`, `status`, `superadmin`, `created_at`, `updated_at`, `registration_ip`, `email`, `auth_key`, `bind_to_ip`, `email_confirmed`, `confirmation_token`, `id_wh`, `user_lang`) VALUES
(1, 'superadmin', '$2y$13$bp2w2.mTeJ/ORRVlEjA.jOHw0o49vwAJ.A15RTPjnSyk05M.20ZyS', 1, 1, '0000-00-00 00:00:00.00', '2024-12-30 01:18:18.05', NULL, 'super@mail.com', '20241230081818_5Dl6dfh6w+sFfGAx+n+hb5Iva1kVKWaFbl+PPLxLmQQ=', '', 1, '', NULL, 'id'),
(4, 'appadmin', '$2y$13$.X94ue5lX8Yt10motmlym.HyhumhiBXBZ7leukSITV7e9sTgLjNrK', 1, 0, '2024-10-28 13:20:22.00', '0000-00-00 00:00:00.00', '::1', 'bogosbinted@mail.com', 'pUwr74uXpIAq5h1XQU-3y3PuplNbm2P8', '', 1, '', NULL, 'en'),
(5, 'franzferdinand', '$2y$13$QEUqv2hQRuKQ2uFtNfXbcuPoiz2pyiZsn1kpv3RfboEQEdA8MZC9e', 1, 0, '2024-10-28 14:32:42.00', '0000-00-00 00:00:00.00', '::1', 'ferdinand@mail.com', 'LZN0hVpdM-xAb6SA0AEALIcxiVeCAS5H', '', 1, '', 13, 'en'),
(8, 'warehouse@mail.com', '$2y$13$XZ6TofHa8d5cAMutTqSMs.QzdqeRCGMeRs3ZfqoMQCWveNyOjdTs2', 1, 0, '2024-10-30 08:40:40.00', '0000-00-00 00:00:00.00', '::1', 'fred@mail.com', 'JGU2pLcVye5PBUUqBlGuB7h1M8c6gkrx', '', 1, '', 6, NULL),
(9, 'bobtherepairman', '$2y$13$0FJ7ZV.5Th2sydV4mDEP2u81kx56ocrbVEFZomEzdjK6xPfRt0xGa', 1, 0, '0000-00-00 00:00:00.00', '0000-00-00 00:00:00.00', '::1', 'bobrepair@mail.com', 'MbIXLNVCI6OHZFAZgwzyyLXTeiaVd6At', '', 1, '', NULL, 'id'),
(10, 'bogosbinted', '$2y$13$byH/Orep9xl5ZWDL1IMtH.1NRxOm2IK/s9vDxODkiEp7o2gJ9TuZ2', 1, 0, '0000-00-00 00:00:00.00', '0000-00-00 00:00:00.00', '::1', 'bogosbinted@alien.com', 'LMStQ2EH_AMPtTfWSbwc9Li493FfLSTY', '', 1, '', NULL, NULL),
(11, 'manfred', '$2y$13$S.Pu6X0K3VWvc/kr2atfyuWdDK9xYc1BMFyvHBnjKWMjJCZuUwKRG', 1, 0, '0000-00-00 00:00:00.00', '0000-00-00 00:00:00.00', '::1', 'redbaron@mail.com', 'E8fnjVCzTMoHqvXiFnfAyEEw7aU80X5W', '', 1, '', NULL, 'de'),
(12, 'hugh', '$2y$13$iREFX85Oa9udbjkKDslyy.46GR1OLNOKYh6JaiZ03pfwS59xaCnUS', 1, 0, '0000-00-00 00:00:00.00', '0000-00-00 00:00:00.00', '::1', 'hugh@mail.com', '-R2PpAC6v4_aTMieMxphBxZ3l_tgVRez', '', 1, '', NULL, NULL),
(14, 'admin_en', '$2y$13$NHfn.C7OGhaRG5/7IIHy/.Wj7q0jNKBOznQAz9kjlnCtDNJP0G3IG', 1, 1, '0000-00-00 00:00:00.00', '0000-00-00 00:00:00.00', '::1', 'admin_en@mail.com', '9YMT5b8a1ZCf18Sa01GV9z_YC4o7tLzI', '', 1, NULL, NULL, 'en'),
(15, 'ayylmao', '$2y$13$x0QphLkVQfC3H0emEqg...BIYV/WcQ3BgnkfjBDH9EBxRw6D91sIK', 1, 0, '0000-00-00 00:00:00.00', '0000-00-00 00:00:00.00', '::1', 'ayylmao@mail.com', 'qCfi3ew0w3uCaGeSUCTXCGiWsnyHhiOI', '', 1, NULL, NULL, NULL),
(16, 'wilhelm', '$2y$13$blFp7hYsLyNBWP24NbxLxuSS8uVGcW/iSaKgnmfekmCy5rB6fYp3.', 1, 0, '0000-00-00 00:00:00.00', '0000-00-00 00:00:00.00', '::1', 'kaiser@mail.com', 'JZ1mfkN_QMtMh9X1Gwf3IHatXo1DqOVB', '', 1, NULL, NULL, 'id'),
(17, 'paul ayyy LMAO bogos binted3', '$2y$13$rxLl1fw.rhbgq/5DyP129OHNCBlpBBg3Equ1vSuNSb71G/pALVruW', 1, 0, '0000-00-00 00:00:00.00', '2024-12-30 01:31:59.26', '::1', 'paul@mail.com', 'xRQolWsjnmiVd1mWyRY32ppfye--PKNh', '', 1, '', NULL, ''),
(22, 'goadmin', '$2a$10$8XAA5AqN2xBGqLAEjIAcH.BJrzE99ay2ICBRhaGt9AviXeN3hGnzm', 1, 1, '2024-12-18 05:56:55.64', '2024-12-20 01:50:21.28', '::1', 'goadmin@example.com', '20241220085021_NwVXGfNh+nSI9jQj/uaEhAZk0GzwNnDpjrE7ooHxwSE=', '', 0, '', NULL, ''),
(30, '1gogolang', '$2a$10$XD03b3xbYiqzTqu3SGCimOyKOaga.ah5qtP/hMBTX51lytY01Fbzy', 1, 0, '2024-12-23 04:23:05.43', '2024-12-23 04:23:05.43', '::1', '1gogolang@example.com', '', '', 0, '', NULL, '');

-- --------------------------------------------------------

--
-- Table structure for table `user_visit_log`
--

CREATE TABLE `user_visit_log` (
  `id` int(11) NOT NULL,
  `token` varchar(255) NOT NULL,
  `ip` varchar(15) NOT NULL,
  `language` char(2) NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `visit_time` int(11) NOT NULL,
  `browser` varchar(30) DEFAULT NULL,
  `os` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `user_visit_log`
--

INSERT INTO `user_visit_log` (`id`, `token`, `ip`, `language`, `user_agent`, `user_id`, `visit_time`, `browser`, `os`) VALUES
(1, '672c16dc84d32', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1730942684, 'Chrome', 'Windows'),
(2, '672c18f5e7365', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1730943221, 'Chrome', 'Windows'),
(3, '672c21efda00e', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1730945519, 'Chrome', 'Windows'),
(4, '672c28a54750c', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1730947237, 'Chrome', 'Windows'),
(5, '672c2c60da2ca', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1730948192, 'Chrome', 'Windows'),
(6, '672c35658ae4b', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 9, 1730950501, 'Chrome', 'Windows'),
(7, '672c37b84241f', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 9, 1730951096, 'Chrome', 'Windows'),
(8, '672c39c88e63e', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 9, 1730951624, 'Chrome', 'Windows'),
(9, '672c3a385d005', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1730951736, 'Chrome', 'Windows'),
(10, '672c3afd75f13', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1730951933, 'Chrome', 'Windows'),
(11, '672c3b5ac8bc0', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1730952026, 'Chrome', 'Windows'),
(12, '672c3b61b9b5c', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1730952033, 'Chrome', 'Windows'),
(13, '672c3be089697', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1730952160, 'Chrome', 'Windows'),
(14, '672c3dd86522e', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 10, 1730952664, 'Chrome', 'Windows'),
(15, '672c3f7005e81', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 11, 1730953072, 'Chrome', 'Windows'),
(16, '672c4196049a8', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1730953622, 'Chrome', 'Windows'),
(17, '672c41cd410d8', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', NULL, 1730953677, 'Chrome', 'Windows'),
(18, '672c41f6c5272', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 11, 1730953718, 'Chrome', 'Windows'),
(19, '672c43edd6228', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 11, 1730954221, 'Chrome', 'Windows'),
(20, '672c43ee9b9ab', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 11, 1730954222, 'Chrome', 'Windows'),
(21, '672c5e257ecb2', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1730960933, 'Chrome', 'Windows'),
(22, '672c5e6b70cba', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 11, 1730961003, 'Chrome', 'Windows'),
(23, '672c607b9de46', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 9, 1730961531, 'Chrome', 'Windows'),
(24, '672c631a85584', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 11, 1730962202, 'Chrome', 'Windows'),
(25, '672c63a2a9199', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 9, 1730962338, 'Chrome', 'Windows'),
(26, '672c64e152848', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1730962657, 'Chrome', 'Windows'),
(27, '672c64ebcb9fb', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 10, 1730962667, 'Chrome', 'Windows'),
(28, '672c66f688a01', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 9, 1730963190, 'Chrome', 'Windows'),
(29, '672c66feeaa8a', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 10, 1730963198, 'Chrome', 'Windows'),
(30, '672d65a65750b', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731028390, 'Chrome', 'Windows'),
(31, '672d84ff293ca', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 10, 1731036415, 'Chrome', 'Windows'),
(32, '672d850843ba3', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1731036424, 'Chrome', 'Windows'),
(33, '672d86868f007', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 4, 1731036806, 'Chrome', 'Windows'),
(34, '672d8742e61f7', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 5, 1731036994, 'Chrome', 'Windows'),
(35, '672d8b9728c31', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1731038103, 'Chrome', 'Windows'),
(36, '672d8ce8234f1', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731038440, 'Chrome', 'Windows'),
(37, '672d8dfce99b4', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0', 10, 1731038716, 'Chrome', 'Windows'),
(38, '67315a5579f86', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731287637, 'Chrome', 'Windows'),
(39, '673161bb0d07e', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 10, 1731289531, 'Chrome', 'Windows'),
(40, '673164f703a23', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731290359, 'Chrome', 'Windows'),
(41, '6731687eb8890', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1731291262, 'Chrome', 'Windows'),
(42, '6731751f8a0aa', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731294495, 'Chrome', 'Windows'),
(43, '6731a3ad86b2c', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 10, 1731306413, 'Chrome', 'Windows'),
(44, '6731aa3f0e91b', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1731308095, 'Chrome', 'Windows'),
(45, '6732ab845e8c9', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731373956, 'Chrome', 'Windows'),
(46, '6732acf4ea9cd', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 10, 1731374324, 'Chrome', 'Windows'),
(47, '6732afa4450ca', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1731375012, 'Chrome', 'Windows'),
(48, '6732b0d9e8250', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 11, 1731375321, 'Chrome', 'Windows'),
(49, '6732b0dab1ec6', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 11, 1731375322, 'Chrome', 'Windows'),
(50, '6732b9632abe5', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 11, 1731377507, 'Chrome', 'Windows'),
(51, '6732bb292fbbe', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 11, 1731377961, 'Chrome', 'Windows'),
(52, '6732c64469c9b', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731380804, 'Chrome', 'Windows'),
(53, '6732d039c4535', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731383353, 'Chrome', 'Windows'),
(54, '6732d1633d2e9', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731383651, 'Chrome', 'Windows'),
(55, '6732f1799957b', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1731391865, 'Chrome', 'Windows'),
(56, '6733fe076394c', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1731460615, 'Chrome', 'Windows'),
(57, '6733fe136345e', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1731460627, 'Chrome', 'Windows'),
(58, '67340022e2456', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 9, 1731461154, 'Chrome', 'Windows'),
(59, '673400972f908', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1731461271, 'Chrome', 'Windows'),
(60, '673d45b403fb5', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1732068788, 'Chrome', 'Windows'),
(61, '673d4bca6ff9a', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1732070346, 'Chrome', 'Windows'),
(62, '673e88c959f1f', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1732151497, 'Chrome', 'Windows'),
(63, '673e88ca5f3d1', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1732151498, 'Chrome', 'Windows'),
(64, '673eaa293ad8c', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1732160041, 'Chrome', 'Windows'),
(65, '673eaa3839d9b', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1732160056, 'Chrome', 'Windows'),
(66, '673eaa58113dc', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1732160088, 'Chrome', 'Windows'),
(67, '673eae0ad8ae6', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 4, 1732161034, 'Chrome', 'Windows'),
(68, '673ee49b8cf5c', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1732175003, 'Chrome', 'Windows'),
(69, '673ee4a854360', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 9, 1732175016, 'Chrome', 'Windows'),
(70, '673ee5352558f', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1732175157, 'Chrome', 'Windows'),
(71, '673fdab05b216', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 1, 1732238000, 'Chrome', 'Windows'),
(72, '673fe051b99ca', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1732239441, 'Chrome', 'Windows'),
(73, '673fe416141fc', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 9, 1732240406, 'Chrome', 'Windows'),
(74, '673fe7ad45e4e', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0', 5, 1732241325, 'Chrome', 'Windows'),
(89, '6743d75373355', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 2147483647, 'Chrome', 'Windows'),
(90, '6743d79f0e2b0', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 2147483647, 'Chrome', 'Windows'),
(91, '6743d7fee5853', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 2147483647, 'Chrome', 'Windows'),
(92, '6743dab9d4a6d', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732500153, 'Chrome', 'Windows'),
(93, '6743dba65af1e', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732500390, 'Chrome', 'Windows'),
(94, '6743dc4cc1625', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732525756, 'Chrome', 'Windows'),
(95, '6743dc5e14174', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732522174, 'Chrome', 'Windows'),
(96, '6743dd7b5c1d6', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 11, 1732522459, 'Chrome', 'Windows'),
(97, '6743de79da2e9', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732522713, 'Chrome', 'Windows'),
(98, '6743df35cff4a', '::1', 'de', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 11, 1732522901, 'Chrome', 'Windows'),
(99, '6743e21beffcd', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1732523643, 'Chrome', 'Windows'),
(100, '6743e26ecabce', '::1', 'de', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 11, 1732523726, 'Chrome', 'Windows'),
(101, '6743e3def0791', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 14, 1732524094, 'Chrome', 'Windows'),
(102, '6743e3dfbe290', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 14, 1732524095, 'Chrome', 'Windows'),
(103, '6743e9cb3238c', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1732525611, 'Chrome', 'Windows'),
(104, '6743ec171820c', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732526199, 'Chrome', 'Windows'),
(105, '6743ee3618eea', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 4, 1732526742, 'Chrome', 'Windows'),
(106, '6743ef3d7a0cb', '::1', 'de', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 11, 1732527005, 'Chrome', 'Windows'),
(107, '6743ef3e529cd', '::1', 'de', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 11, 1732527006, 'Chrome', 'Windows'),
(108, '674416f3d5682', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 11, 1732537171, 'Chrome', 'Windows'),
(109, '67442c1be0bb4', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 17, 1732542587, 'Chrome', 'Windows'),
(110, '67442cc74dc33', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 17, 1732542759, 'Chrome', 'Windows'),
(111, '67442d50cb2b0', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1732542896, 'Chrome', 'Windows'),
(112, '67442d5e0676f', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 4, 1732542910, 'Chrome', 'Windows'),
(113, '6745215f850cb', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732605375, 'Chrome', 'Windows'),
(114, '67452160b8777', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732605376, 'Chrome', 'Windows'),
(115, '674521aab8ca5', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 17, 1732605450, 'Chrome', 'Windows'),
(116, '6745225402f54', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:126.0) Gecko/20100101 Firefox/126.0', 4, 1732605620, 'Firefox', 'Windows'),
(117, '674526a7d2dac', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:126.0) Gecko/20100101 Firefox/126.0', 5, 1732606727, 'Firefox', 'Windows'),
(118, '674533031726b', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', 1, 1732609891, 'Chrome', 'Windows'),
(119, '6745359f1ca3a', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 14, 1732610559, 'Chrome', 'Windows'),
(120, '6745385b07f9e', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 9, 1732611259, 'Chrome', 'Windows'),
(121, '674538636b512', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 14, 1732611267, 'Chrome', 'Windows'),
(122, '6745789e7baaf', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', 1, 1732627710, 'Chrome', 'Windows'),
(123, '67457eaace6b0', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', 16, 1732629258, 'Chrome', 'Windows'),
(124, '67457f3baf2db', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 14, 1732629403, 'Chrome', 'Windows'),
(125, '6747c66364a5e', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 14, 1732778691, 'Chrome', 'Windows'),
(126, '6747dba523bf7', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732784133, 'Chrome', 'Windows'),
(127, '6747dba60b4a2', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732784134, 'Chrome', 'Windows'),
(128, '6747e3caed14b', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1732786218, 'Chrome', 'Windows'),
(129, '674814093f8c7', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1732798569, 'Chrome', 'Windows'),
(130, '674914be85960', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732864286, 'Chrome', 'Windows'),
(131, '674923960ecec', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1732868086, 'Chrome', 'Windows'),
(132, '67496dc4cccfd', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 9, 1732887076, 'Chrome', 'Windows'),
(133, '67497135e8dc7', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1732887957, 'Chrome', 'Windows'),
(134, '674d0d19e689f', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733124473, 'Chrome', 'Windows'),
(135, '674d382a1a872', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1733135498, 'Chrome', 'Windows'),
(136, '674d384018eab', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733135520, 'Chrome', 'Windows'),
(137, '674d385c4998b', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733135548, 'Chrome', 'Windows'),
(138, '674d386ca5c3e', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 16, 1733135564, 'Chrome', 'Windows'),
(139, '674d38864e060', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733135590, 'Chrome', 'Windows'),
(140, '674e5d4cd6c62', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733210540, 'Chrome', 'Windows'),
(141, '674e7d2a49f56', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 4, 1733218698, 'Chrome', 'Windows'),
(142, '674e7ecd4fcbb', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1733219117, 'Chrome', 'Windows'),
(143, '674e81339382d', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 9, 1733219731, 'Chrome', 'Windows'),
(144, '674e8134691dc', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 9, 1733219732, 'Chrome', 'Windows'),
(145, '674e813536ab7', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 9, 1733219733, 'Chrome', 'Windows'),
(146, '674e9f3a737d7', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733227418, 'Chrome', 'Windows'),
(147, '674eb8fc991f0', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 9, 1733234012, 'Chrome', 'Windows'),
(148, '674eb93874e3a', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733234072, 'Chrome', 'Windows'),
(149, '674eb9394c177', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733234073, 'Chrome', 'Windows'),
(150, '674eb966e966b', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1733234118, 'Chrome', 'Windows'),
(151, '674eb975cd45e', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733234133, 'Chrome', 'Windows'),
(152, '674fb6618b915', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733298881, 'Chrome', 'Windows'),
(153, '6751048d92451', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733384429, 'Chrome', 'Windows'),
(154, '675251cceb0f3', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733469740, 'Chrome', 'Windows'),
(155, '67526560d540e', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 5, 1733474752, 'Chrome', 'Windows'),
(156, '67526c2706562', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', 1, 1733476487, 'Chrome', 'Windows'),
(157, '67526d97b9d5f', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', 5, 1733476855, 'Chrome', 'Windows'),
(158, '6752746206873', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', 5, 1733478594, 'Chrome', 'Windows'),
(159, '67527462ca07c', '::1', 'en', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36', 5, 1733478594, 'Chrome', 'Windows'),
(160, '675648840bb4f', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733729508, 'Chrome', 'Windows'),
(161, '675658acb8d90', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733733644, 'Chrome', 'Windows'),
(162, '675658b2483db', '::1', 'id', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0', 1, 1733733650, 'Chrome', 'Windows'),
(163, '241218105856', '::1', '', 'PostmanRuntime/7.43.0', NULL, 1734494336, 'Unknown', 'Unknown'),
(164, '241218110218', '::1', '', 'PostmanRuntime/7.43.0', NULL, 1734494538, 'Unknown', 'Unknown'),
(165, '241218110222', '::1', '', 'PostmanRuntime/7.43.0', NULL, 1734494542, 'Unknown', 'Unknown'),
(166, '241218110402', '::1', '', 'PostmanRuntime/7.43.0', NULL, 1734494642, 'Unknown', 'Unknown'),
(167, '241218111830', '::1', '', 'PostmanRuntime/7.43.0', NULL, 1734495510, 'Unknown', 'Unknown'),
(168, '241218114457', '::1', '', 'PostmanRuntime/7.43.0', NULL, 1734497097, 'Unknown', 'Unknown'),
(169, '241218114620', '::1', '', 'PostmanRuntime/7.43.0', NULL, 1734497180, 'Unknown', 'Unknown'),
(170, '241218130210', '::1', '', 'PostmanRuntime/7.43.0', 22, 1734501730, 'Unknown', 'Unknown'),
(171, '241218130924', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734502164, 'Unknown', 'Unknown'),
(172, '241218131554', '::1', '', 'PostmanRuntime/7.43.0', 22, 1734502554, 'Unknown', 'Unknown'),
(173, '241218133513', '::1', '', 'PostmanRuntime/7.43.0', 22, 1734503713, 'Unknown', 'Unknown'),
(174, '241218133604', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734503764, 'Unknown', 'Unknown'),
(175, '241218133829', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734503909, 'Unknown', 'Unknown'),
(176, '241218133839', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734503919, 'Unknown', 'Unknown'),
(177, '241219082517', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734571517, 'Unknown', 'Unknown'),
(178, '241219144637', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734594397, 'Unknown', 'Unknown'),
(179, '241219145023', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734594623, 'Unknown', 'Unknown'),
(180, '241219145027', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734594627, 'Unknown', 'Unknown'),
(181, '241219145901', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734595141, 'Unknown', 'Unknown'),
(182, '241220083913', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734658753, 'Unknown', 'Unknown'),
(183, '241220085021', '::1', '', 'PostmanRuntime/7.43.0', 22, 1734659421, 'Unknown', 'Unknown'),
(184, '241223083003', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734917403, 'Unknown', 'Unknown'),
(185, '241223110212', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734926532, 'Unknown', 'Unknown'),
(186, '241223130901', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734934141, 'Unknown', 'Unknown'),
(187, '241223141258', '::1', '', 'PostmanRuntime/7.43.0', 1, 1734937978, 'Unknown', 'Unknown'),
(188, '241230081818', '::1', '', 'PostmanRuntime/7.43.0', 1, 1735521498, 'Unknown', 'Unknown');

-- --------------------------------------------------------

--
-- Table structure for table `warehouse`
--

CREATE TABLE `warehouse` (
  `id_wh` int(10) UNSIGNED NOT NULL COMMENT 'primary key',
  `wh_name` varchar(255) NOT NULL COMMENT 'warehouse name',
  `wh_address` varchar(255) NOT NULL COMMENT 'warehouse address'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `warehouse`
--

INSERT INTO `warehouse` (`id_wh`, `wh_name`, `wh_address`) VALUES
(5, 'JOG Condongcatur', 'Jl Gk tau, Condongcatur, Sleman'),
(6, 'Semarang Office', 'somewhere in semarang'),
(7, 'JKT Headquarter', 'somewhere in Jaksel'),
(8, 'Wonosobo Office Wh', 'somewhere in wonosobo'),
(9, 'PWT Office Wh', 'somewhere in pwt'),
(10, 'CLP Office Wh', 'somewhere in CLP'),
(11, 'SRBY Office Wh', 'somewhere in surabaya'),
(13, 'BDG Office Wh', 'somewhere in a bandung'),
(15, 'SPAM PREVENTION TEST', 'SPAM PREVENTION TEST'),
(16, 'Brandenburg', 'somewhere in German'),
(17, 'Marsaille', 'somewhere in france'),
(18, 'West Warehouse', '123 West Street'),
(33, 'West Warehouse 2', '123 West Street');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `api_route`
--
ALTER TABLE `api_route`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `path_method` (`path`,`method`);

--
-- Indexes for table `auth_item`
--
ALTER TABLE `auth_item`
  ADD PRIMARY KEY (`id`),
  ADD KEY `route` (`path`),
  ADD KEY `role` (`role`);

--
-- Indexes for table `auth_roles`
--
ALTER TABLE `auth_roles`
  ADD PRIMARY KEY (`code`);

--
-- Indexes for table `auth_roles_user`
--
ALTER TABLE `auth_roles_user`
  ADD PRIMARY KEY (`user_id`),
  ADD KEY `roles_code` (`roles_code`);

--
-- Indexes for table `condition_lookup`
--
ALTER TABLE `condition_lookup`
  ADD PRIMARY KEY (`id_condition`);

--
-- Indexes for table `doc_uploaded`
--
ALTER TABLE `doc_uploaded`
  ADD PRIMARY KEY (`id_doc`),
  ADD KEY `doc_user` (`user_id`);

--
-- Indexes for table `employee`
--
ALTER TABLE `employee`
  ADD PRIMARY KEY (`id_employee`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `item`
--
ALTER TABLE `item`
  ADD PRIMARY KEY (`id_item`),
  ADD UNIQUE KEY `SKU` (`SKU`),
  ADD KEY `cat_item` (`id_category`);

--
-- Indexes for table `item_category`
--
ALTER TABLE `item_category`
  ADD PRIMARY KEY (`id_category`);

--
-- Indexes for table `item_unit`
--
ALTER TABLE `item_unit`
  ADD PRIMARY KEY (`id_unit`),
  ADD UNIQUE KEY `serial_number` (`serial_number`),
  ADD KEY `id_item` (`id_item`),
  ADD KEY `id_wh` (`id_wh`),
  ADD KEY `item_unit_ibfk_1` (`status`),
  ADD KEY `item_unit_ibfk_2` (`condition`),
  ADD KEY `item_unit_user_updated` (`updated_by`);

--
-- Indexes for table `lending`
--
ALTER TABLE `lending`
  ADD PRIMARY KEY (`id_lending`),
  ADD KEY `type` (`type`),
  ADD KEY `id_unit` (`id_unit`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `id_employee` (`id_employee`);

--
-- Indexes for table `lending_type_lookup`
--
ALTER TABLE `lending_type_lookup`
  ADD PRIMARY KEY (`id_type`);

--
-- Indexes for table `migration`
--
ALTER TABLE `migration`
  ADD PRIMARY KEY (`version`);

--
-- Indexes for table `repair_log`
--
ALTER TABLE `repair_log`
  ADD PRIMARY KEY (`id_repair`),
  ADD KEY `rep_type` (`rep_type`);

--
-- Indexes for table `rep_type_lookup`
--
ALTER TABLE `rep_type_lookup`
  ADD PRIMARY KEY (`id_rep_t`);

--
-- Indexes for table `status_lookup`
--
ALTER TABLE `status_lookup`
  ADD PRIMARY KEY (`id_status`);

--
-- Indexes for table `unit_log`
--
ALTER TABLE `unit_log`
  ADD PRIMARY KEY (`id_log`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD KEY `wh_usuer` (`id_wh`);

--
-- Indexes for table `user_visit_log`
--
ALTER TABLE `user_visit_log`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `warehouse`
--
ALTER TABLE `warehouse`
  ADD PRIMARY KEY (`id_wh`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `api_route`
--
ALTER TABLE `api_route`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=380;

--
-- AUTO_INCREMENT for table `auth_item`
--
ALTER TABLE `auth_item`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=333;

--
-- AUTO_INCREMENT for table `condition_lookup`
--
ALTER TABLE `condition_lookup`
  MODIFY `id_condition` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `doc_uploaded`
--
ALTER TABLE `doc_uploaded`
  MODIFY `id_doc` int(11) NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=113;

--
-- AUTO_INCREMENT for table `employee`
--
ALTER TABLE `employee`
  MODIFY `id_employee` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=33;

--
-- AUTO_INCREMENT for table `item`
--
ALTER TABLE `item`
  MODIFY `id_item` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=81;

--
-- AUTO_INCREMENT for table `item_category`
--
ALTER TABLE `item_category`
  MODIFY `id_category` int(10) NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=28;

--
-- AUTO_INCREMENT for table `item_unit`
--
ALTER TABLE `item_unit`
  MODIFY `id_unit` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=1172;

--
-- AUTO_INCREMENT for table `lending`
--
ALTER TABLE `lending`
  MODIFY `id_lending` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=53;

--
-- AUTO_INCREMENT for table `lending_type_lookup`
--
ALTER TABLE `lending_type_lookup`
  MODIFY `id_type` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `repair_log`
--
ALTER TABLE `repair_log`
  MODIFY `id_repair` int(10) NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=213;

--
-- AUTO_INCREMENT for table `rep_type_lookup`
--
ALTER TABLE `rep_type_lookup`
  MODIFY `id_rep_t` tinyint(3) NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `status_lookup`
--
ALTER TABLE `status_lookup`
  MODIFY `id_status` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `unit_log`
--
ALTER TABLE `unit_log`
  MODIFY `id_log` int(11) NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=344;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=31;

--
-- AUTO_INCREMENT for table `user_visit_log`
--
ALTER TABLE `user_visit_log`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=189;

--
-- AUTO_INCREMENT for table `warehouse`
--
ALTER TABLE `warehouse`
  MODIFY `id_wh` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key', AUTO_INCREMENT=34;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `auth_item`
--
ALTER TABLE `auth_item`
  ADD CONSTRAINT `role` FOREIGN KEY (`role`) REFERENCES `auth_roles` (`code`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `route` FOREIGN KEY (`path`) REFERENCES `api_route` (`path`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `auth_roles_user`
--
ALTER TABLE `auth_roles_user`
  ADD CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Constraints for table `doc_uploaded`
--
ALTER TABLE `doc_uploaded`
  ADD CONSTRAINT `doc_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `item`
--
ALTER TABLE `item`
  ADD CONSTRAINT `cat_item` FOREIGN KEY (`id_category`) REFERENCES `item_category` (`id_category`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `item_unit`
--
ALTER TABLE `item_unit`
  ADD CONSTRAINT `item_unit_ibfk_1` FOREIGN KEY (`status`) REFERENCES `status_lookup` (`id_status`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `item_unit_ibfk_2` FOREIGN KEY (`condition`) REFERENCES `condition_lookup` (`id_condition`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `item_unit_ibfk_3` FOREIGN KEY (`id_item`) REFERENCES `item` (`id_item`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `item_unit_ibfk_4` FOREIGN KEY (`id_wh`) REFERENCES `warehouse` (`id_wh`) ON DELETE SET NULL ON UPDATE NO ACTION,
  ADD CONSTRAINT `item_unit_user_updated` FOREIGN KEY (`updated_by`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `lending`
--
ALTER TABLE `lending`
  ADD CONSTRAINT `lending_ibfk_1` FOREIGN KEY (`id_unit`) REFERENCES `item_unit` (`id_unit`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `lending_ibfk_2` FOREIGN KEY (`type`) REFERENCES `lending_type_lookup` (`id_type`),
  ADD CONSTRAINT `lending_ibfk_3` FOREIGN KEY (`id_employee`) REFERENCES `employee` (`id_employee`) ON DELETE CASCADE,
  ADD CONSTRAINT `user_lending_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Constraints for table `repair_log`
--
ALTER TABLE `repair_log`
  ADD CONSTRAINT `rep_type` FOREIGN KEY (`rep_type`) REFERENCES `rep_type_lookup` (`id_rep_t`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `user`
--
ALTER TABLE `user`
  ADD CONSTRAINT `wh_usuer` FOREIGN KEY (`id_wh`) REFERENCES `warehouse` (`id_wh`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `user_visit_log`
--
ALTER TABLE `user_visit_log`
  ADD CONSTRAINT `user_visit_log_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
