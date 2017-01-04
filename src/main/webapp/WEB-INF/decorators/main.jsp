<%@ page language="java" import="java.util.*"  contentType="text/html;charset=utf-8" pageEncoding="UTF-8"%>
<%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core"%>
<%@taglib prefix="decorator" uri="http://www.opensymphony.com/sitemesh/decorator"%>
<%
String path = request.getContextPath();
String basePath = request.getScheme()+"://"+request.getServerName()+":"+request.getServerPort()+path;
%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
	<head>
		<meta http-equiv="pragma" content="no-cache"/>
		<meta http-equiv="cache-control" content="no-cache"/>
		<meta http-equiv="expires" content="0"/>    
		<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
<%-- 		<script type="text/javascript" src="<%= basePath%>/libs/js/jquery.js"></script>
		<script type="text/javascript" src="<%=basePath%>/libs/js/language/cn.js"></script>
		<script type="text/javascript" src="<%= basePath%>/libs/js/framework.js"></script>
		<link href="<%= basePath%>/libs/css/import_basic.css" rel="stylesheet" type="text/css"/>
		<link rel="stylesheet" type="text/css" id="skin" prePath="<%= path%>/"/>
		<link rel="stylesheet" type="text/css" id="customSkin"/>
		框架必需end
		<script type="text/javascript" src="<%=basePath%>/libs/js/nav/pageNumber.js"></script> --%>
		<decorator:head />
	</head>
	<body <decorator:getProperty property="body.style" writeEntireProperty="true" /> <decorator:getProperty property="body.onload" writeEntireProperty="true" />>
		<decorator:body />
	</body>

</html>