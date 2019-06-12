--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

-- Started on 2019-06-12 11:15:20 CST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 48628)
-- Name: zhparser; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS zhparser WITH SCHEMA public;


--
-- TOC entry 4764 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION zhparser; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION zhparser IS 'a parser for full-text search of Chinese';


--
-- TOC entry 420 (class 1255 OID 46434)
-- Name: _date_before(integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public._date_before(p_day_numb integer) RETURNS timestamp without time zone
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying;   
	var_date_before timestamp;
BEGIN  
    SELECT 'SELECT (now() + ''-' || (p_day_numb-1) || ' day'')::date::timestamp' INTO var_sql;
	RAISE NOTICE 'in _date_before, var_sql = %', var_sql;
	
	EXECUTE var_sql INTO var_date_before;
	RAISE NOTICE 'in _date_before, var_date_before = %', var_date_before;
	
	RETURN var_date_before::timestamp;
	
END;  
$$;


ALTER FUNCTION public._date_before(p_day_numb integer) OWNER TO dbuser;

--
-- TOC entry 422 (class 1255 OID 46435)
-- Name: _date_list(timestamp without time zone); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public._date_list(p_time_before timestamp without time zone) RETURNS TABLE(dt date)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying;   
	var_time_before timestamp;
BEGIN  
	RETURN QUERY
		SELECT col::date FROM generate_series(p_time_before, now(), '1 day'::interval) col;
	
END;  
$$;


ALTER FUNCTION public._date_list(p_time_before timestamp without time zone) OWNER TO dbuser;

--
-- TOC entry 423 (class 1255 OID 46436)
-- Name: _project_list(integer, character varying); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public._project_list(p_project_id integer, p_project_type character varying) RETURNS TABLE(id integer)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying;   
BEGIN 
	SELECT 
		'select prj.id from "TstProject" prj'
	INTO var_sql;
	
	IF p_project_type='project' THEN
    	SELECT var_sql || ' where prj.id = ' || p_project_id
		INTO var_sql;
	ELSIF p_project_type='group' THEN
		SELECT var_sql || ' left join "TstProject" grp on prj."parentId"=grp.id' ||
			' where grp.id = ' || p_project_id 
		INTO var_sql; 
	ELSIF p_project_type='org' THEN
		SELECT var_sql || ' where prj."orgId" = ' || p_project_id 
		INTO var_sql;
	END IF;
	
	SELECT var_sql ||
		  ' AND prj.deleted != true AND prj.disabled != true'
	INTO var_sql;
	
	RAISE NOTICE 'var_sql = %', var_sql;

	RETURN QUERY EXECUTE var_sql;
END;  
$$;


ALTER FUNCTION public._project_list(p_project_id integer, p_project_type character varying) OWNER TO dbuser;

--
-- TOC entry 424 (class 1255 OID 46437)
-- Name: _project_user(integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public._project_user(p_project_id integer) RETURNS TABLE("userId" integer)
    LANGUAGE plpgsql
    AS $$  
declare  
	
BEGIN
	RETURN QUERY
		select relation1."entityId" from "TstProjectRoleEntityRelation" relation1
			where relation1."type" = 'user' AND relation1."projectId" = p_project_id
		UNION
		select relta."userId" from "TstOrgGroupUserRelation" relta
			where relta."orgGroupId" in
			(
				select relation2."entityId" from "TstProjectRoleEntityRelation" relation2
					where relation2.type = 'group' AND relation2."projectId" = p_project_id
			);
END;  
$$;


ALTER FUNCTION public._project_user(p_project_id integer) OWNER TO dbuser;

--
-- TOC entry 425 (class 1255 OID 46438)
-- Name: _user_org_role(integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public._user_org_role(p_user_id integer) RETURNS TABLE("orgRoleId" integer)
    LANGUAGE plpgsql
    AS $$  
declare  
	
BEGIN
	RETURN QUERY
		SELECT r_role_user."orgRoleId"
                  from "TstOrgRoleUserRelation" r_role_user
                where r_role_user."userId" = p_user_id

		UNION

		SELECT r_role_group."orgRoleId"
		  from "TstOrgRoleGroupRelation" r_role_group
		  JOIN "TstOrgGroupUserRelation" r_group_user 
		  		ON r_group_user."orgGroupId" = r_role_group."orgGroupId"
		WHERE r_group_user."userId" = p_user_id;

END;  
$$;


ALTER FUNCTION public._user_org_role(p_user_id integer) OWNER TO dbuser;

--
-- TOC entry 426 (class 1255 OID 46439)
-- Name: _user_org_role(integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public._user_org_role(p_user_id integer, p_org_id integer) RETURNS TABLE("orgRoleId" integer)
    LANGUAGE plpgsql
    AS $$  
declare  
	
BEGIN
	RETURN QUERY
		SELECT r_role_user."orgRoleId"
                  from "TstOrgRoleUserRelation" r_role_user
                where r_role_user."userId" = p_user_id
                  AND r_role_user."orgId" = p_org_id

		UNION

		SELECT r_role_group."orgRoleId"
		  from "TstOrgRoleGroupRelation" r_role_group
		  JOIN "TstOrgGroupUserRelation" r_group_user 
		  		ON r_group_user."orgGroupId" = r_role_group."orgGroupId"
		WHERE r_group_user."userId" = p_user_id
		  AND r_group_user."orgId" = p_org_id;

END;  
$$;


ALTER FUNCTION public._user_org_role(p_user_id integer, p_org_id integer) OWNER TO dbuser;

--
-- TOC entry 419 (class 1255 OID 46440)
-- Name: _user_project_role(integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public._user_project_role(p_user_id integer) RETURNS TABLE("projectId" integer, "projectRoleId" integer)
    LANGUAGE plpgsql
    AS $$  
declare  
	
BEGIN
	RETURN QUERY
		select relation."projectId", relation."projectRoleId" 
	    from "TstProjectRoleEntityRelation" relation
      	where (
          (relation.type = 'user' AND relation."entityId" = p_user_id)
          or (relation.type = 'group' AND
              relation."entityId" in (
                select grp.id from "TstOrgGroup" grp
                  left join "TstOrgGroupUserRelation" relat on relat."orgGroupId" = grp.id
                  left join "TstUser" userr on relat."userId" = userr.id
                where userr.id = p_user_id
			  )
          )
        )
	    order by relation."projectId",  relation."projectRoleId";

END;  
$$;


ALTER FUNCTION public._user_project_role(p_user_id integer) OWNER TO dbuser;

--
-- TOC entry 421 (class 1255 OID 46441)
-- Name: _user_project_role(integer, integer, character varying); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public._user_project_role(p_user_id integer, p_project_id integer, p_project_type character varying) RETURNS TABLE("projectId" integer, "projectRoleId" integer)
    LANGUAGE plpgsql
    AS $$  
declare  
	
BEGIN
	RETURN QUERY
		select relation."projectId", relation."projectRoleId" 
	    from "TstProjectRoleEntityRelation" relation
      	where (
          (relation.type = 'user' AND relation."entityId" = p_user_id)
          or (relation.type = 'group' AND
              relation."entityId" in (
                select grp.id from "TstOrgGroup" grp
                  left join "TstOrgGroupUserRelation" relat on relat."orgGroupId" = grp.id
                  left join "TstUser" userr on relat."userId" = userr.id
                where userr.id = p_user_id
                -- UNION
                -- select grp.id from TstOrgGroup grp
                -- where grp.name = '所有人' and grp.orgId = orgId
			  )
          )
        )
        and relation."projectId" = ANY (select * from _project_list(p_project_id,p_project_type)) 
	    order by relation."projectId",  relation."projectRoleId";

END;  
$$;


ALTER FUNCTION public._user_project_role(p_user_id integer, p_project_id integer, p_project_type character varying) OWNER TO dbuser;

--
-- TOC entry 452 (class 1255 OID 46442)
-- Name: add_case_to_suite(integer, integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.add_case_to_suite(p_case_id integer, p_suite_id integer, p_project_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  

DECLARE
   var_count integer = 0;

BEGIN
	--IF NOT EXISTS 
	--	(select csin.* from "TstCaseInSuite" csin 
	--		where csin."suiteId"=p_suite_id and csin."caseId"=p_case_id and csin.deleted != true)  
	--THEN
		INSERT INTO "TstCaseInSuite"
			("projectId", "suiteId",  
			 	"pId", "caseId", "isParent", ordr, disabled, deleted, "createTime")
		  SELECT p_project_id, p_suite_id, 
		  		cs."pId", cs.id, cs."isParent", cs.ordr, false, false, NOW()
          FROM "TstCase" cs WHERE cs.id=p_case_id;
		
		GET DIAGNOSTICS var_count = ROW_COUNT;
	--END IF;
	
	RAISE NOTICE 'var_count = %', var_count;
	RETURN var_count;
END;

$$;


ALTER FUNCTION public.add_case_to_suite(p_case_id integer, p_suite_id integer, p_project_id integer) OWNER TO dbuser;

--
-- TOC entry 431 (class 1255 OID 46443)
-- Name: add_case_to_task(integer, integer, integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.add_case_to_task(p_case_id integer, p_task_id integer, p_plan_id integer, p_project_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  

DECLARE
   var_count integer = 0;

BEGIN

	--IF NOT EXISTS 
	--	(select csin.* from "TstCaseInTask" csin 
	--		where csin."taskId"=p_task_id and csin."caseId"=p_case_id and csin.deleted != true)  
	--THEN
		INSERT INTO "TstCaseInTask"
			("projectId", "planId", "taskId", 
			 	"pId", "caseId", "isParent", ordr, status, disabled, deleted, "createTime")
		  SELECT p_project_id, p_plan_id, p_task_id, 
		  		cs."pId", cs.id, cs."isParent", cs.ordr, 'untest', false, false, NOW()
          FROM "TstCase" cs WHERE cs.id=p_case_id;
		
		GET DIAGNOSTICS var_count = ROW_COUNT;
	--END IF;
	
	RAISE NOTICE 'var_count = %', var_count;
	RETURN var_count;
END;

$$;


ALTER FUNCTION public.add_case_to_task(p_case_id integer, p_task_id integer, p_plan_id integer, p_project_id integer) OWNER TO dbuser;

--
-- TOC entry 450 (class 1255 OID 46444)
-- Name: add_cases_to_suite(character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.add_cases_to_suite(p_case_ids character varying, p_suite_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  

DECLARE
   var_total integer = 0;
   var_project_id integer;
   arr_case_id integer[];
   var_case_id integer;
BEGIN
    SELECT "projectId" from "TstSuite" where id=p_suite_id INTO var_project_id;
	
    SELECT string_to_array(p_case_ids, ',') into arr_case_id;
	RAISE NOTICE 'arr_case_id = %', arr_case_id;
	
	delete from "TstCaseInSuite" where "suiteId" = p_suite_id;
														 
   FOREACH var_case_id IN ARRAY arr_case_id
   LOOP
      RAISE NOTICE 'var_case_id = %', var_case_id;
	  SELECT var_total + add_case_to_suite(var_case_id, p_suite_id, var_project_id) INTO var_total;
   END LOOP;
	
	RETURN var_total;
END;

$$;


ALTER FUNCTION public.add_cases_to_suite(p_case_ids character varying, p_suite_id integer) OWNER TO dbuser;

--
-- TOC entry 433 (class 1255 OID 46445)
-- Name: add_cases_to_task(character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.add_cases_to_task(p_case_ids character varying, p_task_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  

DECLARE
   var_total integer = 0;
   
   var_plan_id integer;
   var_project_id integer;
   
   arr_case_id integer[];
   var_case_id integer;
BEGIN
    SELECT "planId", "projectId" from "TstTask" where id=p_task_id INTO var_plan_id, var_project_id;
	
    SELECT string_to_array(p_case_ids, ',') into arr_case_id;
	RAISE NOTICE 'arr_case_id = %', arr_case_id;
	
	update "TstCaseInTask" set deleted=true where "taskId" = p_task_id;
														 
   FOREACH var_case_id IN ARRAY arr_case_id
   LOOP
      RAISE NOTICE 'var_case_id = %', var_case_id;
	  SELECT var_total + add_case_to_task(var_case_id, p_task_id, var_plan_id, var_project_id) INTO var_total;
   END LOOP;
	
	RETURN var_total;
	
END;

$$;


ALTER FUNCTION public.add_cases_to_task(p_case_ids character varying, p_task_id integer) OWNER TO dbuser;

--
-- TOC entry 432 (class 1255 OID 46446)
-- Name: add_cases_to_task_by_suites(character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.add_cases_to_task_by_suites(p_suite_ids character varying, p_task_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  

DECLARE
   var_total integer = 0;
   
   arr_suite_id integer[];
   var_suite_id integer;
   
   var_case_ids character varying;
BEGIN

    SELECT string_to_array(p_suite_ids, ',') into arr_suite_id;
	RAISE NOTICE 'arr_suite_id = %', arr_suite_id;
	
	update "TstCaseInTask" set deleted=true where "taskId" = p_task_id;
														 
   FOREACH var_suite_id IN ARRAY arr_suite_id
   LOOP
      RAISE NOTICE 'var_suite_id = %', var_suite_id;
	  
	  SELECT array_to_string(ARRAY(SELECT unnest(array_agg("caseId"))),',') 
			 FROM "TstCaseInSuite" where "suiteId"=var_suite_id INTO var_case_ids;
	  
	  SELECT var_total + add_cases_to_task(var_case_ids, p_task_id) INTO var_total;
   END LOOP;
	
	RETURN var_total;
	
END;

$$;


ALTER FUNCTION public.add_cases_to_task_by_suites(p_suite_ids character varying, p_task_id integer) OWNER TO dbuser;

--
-- TOC entry 429 (class 1255 OID 46447)
-- Name: chart_issue_age(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_issue_age(p_project_id integer, p_project_type character varying, p_day_numb integer) RETURNS TABLE(category text, priority character varying, numb bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying;   
	
	var_project_ids integer[];
BEGIN  
	select array(select _project_list(p_project_id,p_project_type)) INTO var_project_ids;

	RETURN QUERY

	select 
		case
			when temp.days > p_day_numb then '>' || p_day_numb
			else '' || temp.days
		end category, prio.label priority, temp.count
	from (
		select date_part('day', NOW() - isu."createTime") days,
			isu."priorityId", count(isu.id) count 
		from "IsuIssue" isu 
		LEFT JOIN "IsuStatus" sta ON sta.id = isu."statusId"
		where isu."projectId" = ANY (var_project_ids)  
			  AND sta."finalVal" != true
		GROUP BY days, isu."priorityId"
	) temp

	LEFT JOIN "IsuPriority" prio ON prio.id = temp."priorityId"

	ORDER BY temp.days, temp."priorityId";
	
END;  
$$;


ALTER FUNCTION public.chart_issue_age(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO dbuser;

--
-- TOC entry 427 (class 1255 OID 46448)
-- Name: chart_issue_distrib_by_priority(integer, character varying); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_issue_distrib_by_priority(p_project_id integer, p_project_type character varying) RETURNS TABLE(label character varying, count bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	  
BEGIN 
    RETURN QUERY
	select prio.label, count(isu.id) count
	from "IsuIssue" isu
		left join "IsuPriority" prio on isu."priorityId"=prio.id
	where isu."projectId" = any (select _project_list(p_project_id,p_project_type)) 
	     AND isu.deleted != true AND isu.disabled != true
	     group by prio.label;
END;  
$$;


ALTER FUNCTION public.chart_issue_distrib_by_priority(p_project_id integer, p_project_type character varying) OWNER TO dbuser;

--
-- TOC entry 428 (class 1255 OID 46449)
-- Name: chart_issue_distrib_by_status(integer, character varying); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_issue_distrib_by_status(p_project_id integer, p_project_type character varying) RETURNS TABLE(label character varying, count bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	  
BEGIN 
    RETURN QUERY
	select sta.label, count(isu.id) count from "IsuIssue" isu
	      left join "IsuStatus" sta on isu."statusId"=sta.id
	where isu."projectId" = any (select _project_list(p_project_id,p_project_type)) 
	     AND isu.deleted != true AND isu.disabled != true
	     group by sta.label;
END;  
$$;


ALTER FUNCTION public.chart_issue_distrib_by_status(p_project_id integer, p_project_type character varying) OWNER TO dbuser;

--
-- TOC entry 430 (class 1255 OID 46450)
-- Name: chart_issue_trend_create(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_issue_trend_create(p_project_id integer, p_project_type character varying, p_day_numb integer) RETURNS TABLE(date date, sum numeric, numb bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying; 
	
	var_project_ids integer[];
	
    var_date_before timestamp;
    var_numb_before bigint;
BEGIN  
	SELECT _date_before(p_day_numb) INTO var_date_before;
	select array(select _project_list(p_project_id,p_project_type)) INTO var_project_ids;
	
    SELECT COUNT(id) from "IsuIssue" isu WHERE isu."createTime" < var_date_before 
        and isu."projectId" = ANY (var_project_ids)
		INTO var_numb_before;
	RAISE NOTICE 'var_numb_before = %', var_numb_before;
	
	RETURN QUERY
	 SELECT days.dt2, 
	 		(sum(COALESCE(count_by_day.numb, 0)) over (order by days.dt2)) + var_numb_before, 
								   COALESCE(count_by_day.numb, 0)
	 FROM 
	 	(SELECT dt dt2 from _date_list(var_date_before)) days
	 LEFT JOIN (
		SELECT COUNT(isu.id) numb, isu."createTime"::date dt
    	FROM "IsuIssue" isu
        WHERE isu."projectId" = ANY (var_project_ids)  
		    AND isu."createTime" >= var_date_before
			AND isu.deleted != true AND isu.disabled != true
        GROUP BY dt) count_by_day
     
	ON count_by_day.dt = days.dt2;
	
END;  
$$;


ALTER FUNCTION public.chart_issue_trend_create(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO dbuser;

--
-- TOC entry 434 (class 1255 OID 46451)
-- Name: chart_issue_trend_final(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_issue_trend_final(p_project_id integer, p_project_type character varying, p_day_numb integer) RETURNS TABLE(date date, sum numeric, numb bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying;  
	
	var_project_ids integer[];
	
    var_date_before timestamp;
    var_numb_before bigint;
BEGIN  
	SELECT _date_before(p_day_numb) INTO var_date_before;
	select array(select _project_list(p_project_id,p_project_type)) INTO var_project_ids;
	
    SELECT COUNT(id) from "IsuIssue" isu WHERE isu."setFinalTime" < var_date_before 
		and isu."projectId" = ANY (var_project_ids)
	INTO var_numb_before;
	RAISE NOTICE 'var_numb_before = %', var_numb_before;
	
	RETURN QUERY
	 SELECT days.dt2, 
	 		(sum(COALESCE(count_by_day.numb, 0)) over (order by days.dt2)) + var_numb_before, 
								   COALESCE(count_by_day.numb, 0)
	 FROM 
	 	(SELECT dt dt2 from _date_list(var_date_before)) days
	 LEFT JOIN (
		SELECT COUNT(isu.id) numb, isu."setFinalTime"::date dt
    	FROM "IsuIssue" isu
        WHERE isu."projectId" = ANY (var_project_ids)  
		    AND isu."setFinalTime" >= var_date_before
			AND isu.deleted != true AND isu.disabled != true
        GROUP BY dt) count_by_day
     
	ON count_by_day.dt = days.dt2;
	
END;  
$$;


ALTER FUNCTION public.chart_issue_trend_final(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO dbuser;

--
-- TOC entry 435 (class 1255 OID 46452)
-- Name: chart_test_design_progress_by_project(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_test_design_progress_by_project(p_project_id integer, p_project_type character varying, p_day_numb integer) RETURNS TABLE(date date, sum numeric, numb bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying;  
	
	var_project_ids integer[];
	
    var_date_before timestamp;
    var_numb_before bigint;
BEGIN  
	SELECT _date_before(p_day_numb) INTO var_date_before;
	select array(select _project_list(p_project_id,p_project_type)) INTO var_project_ids;
	
    SELECT COUNT(id) from "TstCase" cs 
		WHERE cs."createTime" < var_date_before 
		and cs."projectId" = ANY (var_project_ids)
		INTO var_numb_before;
	RAISE NOTICE 'var_numb_before = %', var_numb_before;
	
	RETURN QUERY
	 SELECT days.dt2, 
	 		(sum(COALESCE(count_by_day.numb, 0)) over (order by days.dt2)) + var_numb_before, 
								   COALESCE(count_by_day.numb, 0)
	 FROM 
	 	(SELECT dt dt2 from _date_list(var_date_before)) days
	 LEFT JOIN (
		SELECT COUNT(cs.id) numb, cs."createTime"::date dt
    	FROM "TstCase" cs
        WHERE cs."projectId" = ANY (var_project_ids)  
			AND cs."isParent"=false
		    AND cs."createTime" >= var_date_before
			AND cs.deleted != true AND cs.disabled != true
        GROUP BY dt) count_by_day
     
	ON count_by_day.dt = days.dt2;
	
END;  
$$;


ALTER FUNCTION public.chart_test_design_progress_by_project(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO dbuser;

--
-- TOC entry 436 (class 1255 OID 46453)
-- Name: chart_test_execution_process_by_plan(integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_test_execution_process_by_plan(p_plan_id integer, p_day_numb integer) RETURNS TABLE(date date, status character varying, sum numeric, numb bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying; 
	
    var_date_before timestamp;
BEGIN  
	SELECT _date_before(p_day_numb) INTO var_date_before;

	RETURN QUERY
	 SELECT days.dt2, count_by_day_and_status.status,
	 		(sum(COALESCE(count_by_day_and_status.numb, 0)) over (order by days.dt2)), 
								   COALESCE(count_by_day_and_status.numb, 0)
	 FROM 
	 	(SELECT dt dt2 from _date_list(var_date_before)) days
	 LEFT JOIN (
		 SELECT COUNT(csr.id) numb, csr."exeTime"::date dt, csr."status"
			FROM "TstCaseInTask" csr
			  left join "TstTask" task on csr."taskId"=task.id
			WHERE csr."planId"=p_plan_id 
		 		  and task.deleted != true AND task.disabled != true
				  AND csr."isParent"=false AND csr.deleted != true AND csr.disabled != TRUE
				  AND csr."status" != 'untest'
			GROUP BY dt, csr."status") count_by_day_and_status
     
	ON count_by_day_and_status.dt = days.dt2;
	
END;  
$$;


ALTER FUNCTION public.chart_test_execution_process_by_plan(p_plan_id integer, p_day_numb integer) OWNER TO dbuser;

--
-- TOC entry 453 (class 1255 OID 46454)
-- Name: chart_test_execution_process_by_plan_user(integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_test_execution_process_by_plan_user(p_plan_id integer, p_day_numb integer) RETURNS TABLE(date date, name character varying, sum numeric, numb bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying; 
    var_date_before timestamp;
BEGIN  
	SELECT _date_before(p_day_numb) INTO var_date_before;
	
	RETURN QUERY
	 SELECT days.dt2, usr.nickname "name",
	 		(sum(COALESCE(count_by_day_and_user.numb, 0)) over (order by days.dt2)), 
								   COALESCE(count_by_day_and_user.numb, 0)
	 FROM 
	 	(SELECT dt dt2 from _date_list(var_date_before)) days
	 LEFT JOIN (
		 SELECT COUNT(csr.id) numb, csr."exeTime"::date dt, csr."exeBy"
			FROM "TstCaseInTask" csr
			  left join "TstTask" task on csr."taskId"=task.id
			WHERE csr."planId"=p_plan_id 
		 		  and task.deleted != true AND task.disabled != true
				  AND csr."isParent"=false AND csr.deleted != true AND csr.disabled != TRUE
				  AND csr."status" != 'untest'
			GROUP BY dt, csr."exeBy") count_by_day_and_user
	 LEFT JOIN "TstUser" usr on count_by_day_and_user."exeBy" = usr.id
     
	ON count_by_day_and_user.dt = days.dt2
	ORDER BY days.dt2, "name";
	
END;  
$$;


ALTER FUNCTION public.chart_test_execution_process_by_plan_user(p_plan_id integer, p_day_numb integer) OWNER TO dbuser;

--
-- TOC entry 437 (class 1255 OID 46455)
-- Name: chart_test_execution_process_by_project(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_test_execution_process_by_project(p_project_id integer, p_project_type character varying, p_day_numb integer) RETURNS TABLE(date date, status character varying, sum numeric, numb bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying; 
	
	var_project_ids integer[];
	
    var_date_before timestamp;
    var_numb_before bigint;
BEGIN  
	SELECT _date_before(p_day_numb) INTO var_date_before;
	select array(select _project_list(p_project_id,p_project_type)) INTO var_project_ids;
	
	RETURN QUERY
	 SELECT days.dt2, count_by_day_and_status.status,
	 		(sum(COALESCE(count_by_day_and_status.numb, 0)) over (order by days.dt2)) + var_numb_before, 
								   COALESCE(count_by_day_and_status.numb, 0)
	 FROM 
	 	(SELECT dt dt2 from _date_list(var_date_before)) days
	 LEFT JOIN (
		 SELECT COUNT(csr.id) numb, csr."exeTime"::date dt, csr."status"
			FROM "TstCaseInTask" csr
		      left JOIN "TstPlan" plan on csr."planId" = plan.id
			  left join "TstTask" task on csr."taskId"=task.id
			WHERE csr."projectId" = ANY (var_project_ids)  
		          and plan.deleted != true AND plan.disabled != true
		 		  and task.deleted != true AND task.disabled != true
				  AND csr."isParent"=false AND csr.deleted != true AND csr.disabled != TRUE
				  AND csr."status" != 'untest'
			GROUP BY dt, csr."status") count_by_day_and_status
     
	ON count_by_day_and_status.dt = days.dt2
	order by days.dt2;
	
END;  
$$;


ALTER FUNCTION public.chart_test_execution_process_by_project(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO dbuser;

--
-- TOC entry 439 (class 1255 OID 46456)
-- Name: chart_test_execution_progress_by_plan(integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_test_execution_progress_by_plan(p_plan_id integer, p_day_numb integer) RETURNS TABLE(date date, "left" numeric, numb bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	var_sql character varying;   
	
    var_date_before timestamp;
    var_total_numb bigint;
BEGIN  
	SELECT _date_before(p_day_numb) INTO var_date_before;
	
	SELECT COUNT(csr.id) numb
    FROM "TstCaseInTask" csr
      left join "TstTask" task on csr."taskId"=task.id
    WHERE csr."planId"=p_plan_id and task.deleted != true AND task.disabled != true
          AND csr."isParent"=false AND csr.deleted != true AND csr.disabled != TRUE
    into var_total_numb;
	
	RETURN QUERY
	 SELECT days.dt2, 
	 		var_total_numb - ( sum(COALESCE(count_by_day.numb, 0)) over (order by days.dt2) ), 
			COALESCE(count_by_day.numb, 0)
	 FROM 
	 	(SELECT dt dt2 from _date_list(var_date_before)) days
	 LEFT JOIN (
		SELECT COUNT(csr.id) numb, csr."exeTime"::date dt
    	FROM "TstCaseInTask" csr
			left join "TstTask" task on csr."taskId"=task.id
        WHERE csr."planId"=p_plan_id and task.deleted != true AND task.disabled != true
			AND csr.status <> 'untest' AND csr."isParent"=false 
		 		AND csr.deleted != true AND csr.disabled != TRUE
		    AND csr."exeTime" >= var_date_before
        GROUP BY dt) count_by_day
     
	ON count_by_day.dt = days.dt2;
	
END;  
$$;


ALTER FUNCTION public.chart_test_execution_progress_by_plan(p_plan_id integer, p_day_numb integer) OWNER TO dbuser;

--
-- TOC entry 440 (class 1255 OID 46457)
-- Name: chart_test_execution_result_by_plan(integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.chart_test_execution_result_by_plan(p_plan_id integer) RETURNS TABLE(status character varying, count bigint)
    LANGUAGE plpgsql
    AS $$  
declare  
	  
BEGIN 
    RETURN QUERY
    select tcin.status status, count(tcin.id) count
    from "TstCaseInTask" tcin
      left join "TstTask" task on tcin."taskId"=task.id
    where tcin."planId"  = p_plan_id and task.deleted != true AND task.disabled != true
          AND tcin.deleted != true AND tcin.disabled != true  AND tcin."isParent"=false
    group by tcin.status;
END;  
$$;


ALTER FUNCTION public.chart_test_execution_result_by_plan(p_plan_id integer) OWNER TO dbuser;

--
-- TOC entry 441 (class 1255 OID 46458)
-- Name: close_plan_if_all_task_closed(integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.close_plan_if_all_task_closed(p_plan_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  
declare  
    v_cnt bigint = 0;  
	v_count bigint;
BEGIN
    select count(task.id) from "TstTask" task
    where task."planId" = p_plan_id
          and task.status <> 'end' and task.deleted!=true and task.disabled!=true 
    into v_count;
	RAISE NOTICE 'v_count = %', v_count;

    IF (v_count = 0) THEN
      update "TstPlan" set status='end' where id=p_plan_id;
	  GET DIAGNOSTICS v_cnt = ROW_COUNT;
    END IF;
	
	RAISE NOTICE 'v_cnt = %', v_cnt;
	
    RETURN v_cnt;  
END;  
$$;


ALTER FUNCTION public.close_plan_if_all_task_closed(p_plan_id integer) OWNER TO dbuser;

--
-- TOC entry 442 (class 1255 OID 46459)
-- Name: gen_project_access_history(integer, integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.gen_project_access_history(p_org_id integer, p_project_id integer, p_project_name character varying, p_user_id integer) RETURNS void
    LANGUAGE plpgsql
    AS $$  
declare  
    v_cnt bigint = 0;  
	v_project_id integer;
	v_project_name character varying;
BEGIN
    select his."prjId", his."prjName" from "TstProjectAccessHistory" his
    where his."orgId" = p_org_id and his."userId" = p_user_id and his."prjId" = p_project_id
    into v_project_id, v_project_name;

    IF (v_project_id is null) THEN
      insert into "TstProjectAccessHistory"
      ("orgId", "userId", "prjId", "prjName", "lastAccessTime")
      values
        (p_org_id, p_user_id, p_project_id, p_project_name, NOW());
	   GET DIAGNOSTICS v_cnt = ROW_COUNT;
	   RAISE NOTICE 'insert v_cnt = %', v_cnt;
    ELSif (v_project_id is not null AND v_project_name <> p_project_name) THEN
      update "TstProjectAccessHistory"
      set "prjName" = p_project_name, "lastAccessTime" = NOW()
      WHERE "prjId" = p_project_id;
	  GET DIAGNOSTICS v_cnt = ROW_COUNT;
	  RAISE NOTICE 'update v_cnt = %', v_cnt;
    END IF;
 
END;  
$$;


ALTER FUNCTION public.gen_project_access_history(p_org_id integer, p_project_id integer, p_project_name character varying, p_user_id integer) OWNER TO dbuser;

--
-- TOC entry 451 (class 1255 OID 64108)
-- Name: get_org_privilege_for_user(integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.get_org_privilege_for_user(p_user_id integer, p_org_id integer) RETURNS TABLE(code character varying, action character varying, name character varying)
    LANGUAGE plpgsql
    AS $$  
declare  
	
BEGIN
	RETURN QUERY
	
	SELECT priv."code", priv."action", priv.name 
	from "TstOrgPrivilegeDefine" priv
	join "TstOrgRolePrivilegeRelation" r_role_priv ON r_role_priv."orgPrivilegeId"=priv.id
	where r_role_priv."orgRoleId" = any (select "orgRoleId" from _user_org_role(p_user_id, p_org_id))
		AND NOT priv.deleted and NOT priv.disabled
	order by priv.id asc;

END;  
$$;


ALTER FUNCTION public.get_org_privilege_for_user(p_user_id integer, p_org_id integer) OWNER TO dbuser;

--
-- TOC entry 443 (class 1255 OID 46461)
-- Name: get_project_privilege_for_user(integer, integer, character varying); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.get_project_privilege_for_user(p_user_id integer, p_project_id integer, p_project_type character varying) RETURNS TABLE("projectId" text, code character varying, action character varying)
    LANGUAGE plpgsql
    AS $$  
declare  
	
BEGIN
	RETURN QUERY
	
   select '' || tmp."projectId", define.code, define.action
   from "TstProjectPrivilegeDefine" define
   left join "TstProjectRolePriviledgeRelation" r on r."projectPrivilegeDefineId" = define.id
   INNER join (select tp."projectId", tp."projectRoleId" 
			   			from _user_project_role(p_user_id,p_project_id,p_project_type) tp ) tmp
       
   on r."projectRoleId" = tmp."projectRoleId"

    where TRUE
    order by tmp."projectId",  define.code;

END;  
$$;


ALTER FUNCTION public.get_project_privilege_for_user(p_user_id integer, p_project_id integer, p_project_type character varying) OWNER TO dbuser;

--
-- TOC entry 446 (class 1255 OID 46462)
-- Name: init_org(integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.init_org(p_org_id integer, p_user_id integer) RETURNS void
    LANGUAGE plpgsql
    AS $$  
declare
    i integer;

    user_name character varying;
    org_role_id integer;
    org_group_id integer;
    project_role_id integer;
    project_role_leader_id integer;
    project_id integer;
	
	plan_id integer;
	task_id integer;

    p_case_id integer;
    case_id integer;
	case_default_exe_status_id integer;
    case_default_priority_id integer;
    case_default_type_id integer;
	
	issue_type_id integer; 
	issue_status_id integer;
	issue_priority_id integer;
	issue_resolution_id integer;

    issue_type_solution_id integer;
    issue_priority_solution_id integer;
    issue_workflow_solution_id integer;

    issue_page_id integer;
    issue_page_solution_id integer;

    issue_workflow_id integer;

    record_gap_to_define_table integer;
    
BEGIN  
    select usr.nickname from "TstUser" usr where id=p_user_id into user_name;

    insert into "TstOrgUserRelation" ("orgId", "userId") values(p_org_id, p_user_id);

    insert into "TstOrgRole" (code, name, "orgId", "buildIn", disabled, deleted, "createTime") 
    values('org_admin', '组织管理员', p_org_id, true, false, false, NOW());
    select max(id) from "TstOrgRole" into org_role_id;
    insert into "TstOrgRolePrivilegeRelation" ("orgId", "orgRoleId", "orgPrivilegeId") values(p_org_id, org_role_id, 1);
    insert into "TstOrgRolePrivilegeRelation" ("orgId", "orgRoleId", "orgPrivilegeId") values(p_org_id, org_role_id, 3);

    insert into "TstOrgRoleUserRelation" ("orgId", "orgRoleId", "userId") values(p_org_id, org_role_id, p_user_id);
		   
    /* insert into "TstOrgRole" (code, name, "orgId", disabled, deleted, "createTime") values('site_admin', '站点管理员', p_org_id, false, false, NOW());
    select max(id) from "TstOrgRole" into org_role_id;
    insert into "TstOrgRolePrivilegeRelation" ("orgId", "orgRoleId", "orgPrivilegeId") values(p_org_id, org_role_id, 2); */

    insert into "TstOrgRole" (code, name, "orgId", "buildIn", disabled, deleted, "createTime") 
    values('project_admin', '项目管理员', p_org_id, true, false, false, NOW());
    select max(id) from "TstOrgRole" into org_role_id;
    insert into "TstOrgRolePrivilegeRelation" ("orgId", "orgRoleId", "orgPrivilegeId") values(p_org_id, org_role_id, 3);

    insert into "TstOrgGroup" (name, "orgId", "buildIn", disabled, deleted, "createTime") 
    values('所有人', p_org_id, true, false, false, NOW());
    select max(id) from "TstOrgGroup" into org_group_id;
		   
    INSERT INTO public."TstOrgGroupUserRelation"("orgId", "orgGroupId", "userId")
	VALUES (p_org_id, org_group_id, p_user_id);

    insert into "TstCaseExeStatus" 
		   ("value", "label", ordr, "buildIn", "finalVal", "orgId", disabled, deleted, "createTime")
    select "value", "label", ordr, true, "finalVal", p_org_id, disabled, deleted, now()
		   from "TstCaseExeStatusDefine";  
	select id from "TstCaseExeStatus" where value='untest' AND "orgId"=p_org_id 
		   into case_default_exe_status_id;

    insert into "TstCasePriority" ("value", "label", ordr, "buildIn", "defaultVal", "orgId", disabled, deleted, "createTime")
	select "value", "label", ordr, true, "defaultVal", p_org_id, disabled, deleted, now()
		   from "TstCasePriorityDefine";
    select id from "TstCasePriority" where value='medium' AND "orgId"=p_org_id 
		   into case_default_priority_id;

    insert into "TstCaseType" ("value", "label", ordr, "buildIn", "defaultVal", "orgId", disabled, deleted, "createTime")
		select "value", "label", ordr, true, "defaultVal", p_org_id, disabled, deleted, now()
		   from "TstCaseTypeDefine";
    select id from "TstCaseType" where value='functional' AND "orgId"=p_org_id 
		   into case_default_type_id;

    insert into "TstProjectRole" (code, name, "buildIn", "orgId", disabled, deleted, "createTime")
    values('test_leader', '测试主管', false, p_org_id, false, false, NOW());
    select max(id) from "TstProjectRole" into project_role_id;
    SELECT project_role_id INTO project_role_leader_id;

    insert into "TstProjectRolePriviledgeRelation" ( "projectPrivilegeDefineId",   "projectRoleId", "orgId" )
    select d.id,project_role_id, p_org_id from "TstProjectPrivilegeDefine" d;

    insert into "TstProjectRole" (code, name, "buildIn", "orgId", disabled, deleted, "createTime")
    values('test_designer', '测试设计', false, p_org_id, false, false, NOW());
    select max(id) from "TstProjectRole" into project_role_id;

    insert into "TstProjectRolePriviledgeRelation" ( "projectPrivilegeDefineId",   "projectRoleId" )
    select d.id,project_role_id from "TstProjectPrivilegeDefine" d 
		   where d.id != 12000 and d.id != 12400;

    insert into "TstProjectRole" (code, name, "buildIn", "orgId", disabled, deleted, "createTime")
    values('tester', '测试执行', false, p_org_id, false, false, NOW());
    select max(id) from "TstProjectRole" into project_role_id;

    insert into "TstProjectRolePriviledgeRelation" ( "projectPrivilegeDefineId",   "projectRoleId" )
    select d.id,project_role_id from "TstProjectPrivilegeDefine" d 
		   where d.id != 12000 and d.id != 12200 and d.id != 12400;

    insert into "TstProjectRole" (code, name, "buildIn", "orgId", disabled, deleted, "createTime")
    values('readonly', '只读用户', false, p_org_id, false, false, NOW());
    select max(id) from "TstProjectRole" into project_role_id;

    insert into "TstProjectRolePriviledgeRelation" ( "projectPrivilegeDefineId",   "projectRoleId" )
    select d.id,project_role_id from "TstProjectPrivilegeDefine" d where d.action = 'view';

    insert into "TstProject" (name, "type", "parentId", "orgId", disabled, deleted, "createTime")
    values('默认项目组', 'group', NULL, p_org_id, false, false, NOW());
    select max(id) from "TstProject" into project_id;

    insert into "TstProject" (name, "type", "parentId", "orgId", disabled, deleted, "createTime")
    values('默认项目', 'project', project_id, p_org_id, false, false, NOW());
    select max(id) from "TstProject" into project_id;

    insert into "TstHistory" ("projectId", "entityId",  "entityType", "userId", disabled, deleted, "createTime", title)
    values(project_id, project_id, 'project', p_user_id, false, false, NOW(),
           CONCAT('用户<span class="dict">',user_name,'</span>初始化项目<span class="dict">','默认项目','</span>'));

    insert into "TstProjectRoleEntityRelation" ("orgId", "projectId", "projectRoleId", "entityId", "type")
    values(p_org_id, project_id, project_role_leader_id, p_user_id, 'user');

    insert into "TstProjectAccessHistory" ("orgId", "prjId", "userId", "prjName", "lastAccessTime" , "createTime")
    values(p_org_id, project_id, p_user_id, '默认项目', NOW(), NOW());
    update "TstUser" set "defaultPrjId" = project_id, "defaultPrjName" = '默认项目' where id = p_user_id;
		   
	-- 初始化执行计划和任务
    INSERT INTO public."TstPlan"(name, status, "projectId", "userId", disabled, deleted, "createTime")
		VALUES ('示例计划', 'not_start', project_id, p_user_id, false, false, now());
    select max(id) from "TstPlan" into plan_id;
		
    INSERT INTO public."TstTask"(
			name, status, "projectId", "caseProjectId", "planId", "userId", 
			disabled, deleted, "createTime")
		VALUES ('示例任务', 'not_start', project_id, project_id, plan_id, p_user_id, 
			false, false, now());
	select max(id) from "TstTask" into task_id;
				
	INSERT INTO public."TstTaskAssigneeRelation"("taskId", "assigneeId")
	VALUES (task_id, p_user_id);

    insert into "TstCase" (name, "projectId", "pId", estimate, "isParent", ordr, "createById", "contentType", disabled, deleted, "createTime")
    values('测试用例', project_id, null, 10, true, 0, p_user_id, 'steps', false, false, NOW());
    select max(id) from "TstCase" into case_id;
    select case_id into p_case_id;
		   
	INSERT INTO "TstCaseInTask"(
		"caseId", "isParent", "pId", ordr, "exeBy", "exeTime", status, "projectId", "planId", "taskId", 
			disabled, deleted, "createBy", "createTime")
	VALUES (case_id, true, null, 1, null, null, 'untest', project_id, plan_id, task_id, 
			false, false, p_user_id, now());
		   
    insert into "TstCase" (name, "projectId", "pId", estimate, "isParent", ordr, "createById", "contentType", disabled, deleted, "createTime")
    values('新特性', project_id, case_id, 10, true, 0, p_user_id, 'steps', false, false, NOW());
    select max(id) from "TstCase" into case_id;
	INSERT INTO "TstCaseInTask"(
		"caseId", "isParent", "pId", ordr, "exeBy", "exeTime", status, "projectId", "planId", "taskId", 
			disabled, deleted, "createBy", "createTime")
	VALUES (case_id, true, p_case_id, 1, null, null, 'untest', project_id, plan_id, task_id, 
			false, false, p_user_id, now());
	select case_id into p_case_id;
		   
    insert into "TstCase" (name, "projectId", "pId", estimate, "priorityId", "typeId", "isParent", ordr, "createById", "contentType", disabled, deleted, "createTime")
    values('新用例', project_id, case_id, 10, case_default_priority_id, case_default_type_id, false, 0, p_user_id, 'steps', false, false, NOW());
    select max(id) from "TstCase" into case_id;
	INSERT INTO "TstCaseInTask"(
		"caseId", "isParent", "pId", ordr, "exeBy", "exeTime", status, "projectId", "planId", "taskId", 
			disabled, deleted, "createBy", "createTime")
	VALUES (case_id, false, p_case_id, 1, null, null, 'untest', project_id, plan_id, task_id, 
			false, false, p_user_id, now());

    insert into "TstCaseStep" (opt, expect, "caseId", ordr, disabled, deleted, "createTime")
    values('操作步骤1', '期待结果1', case_id, 1, false, false, NOW());
    insert into "TstCaseStep" (opt, expect, "caseId", ordr, disabled, deleted, "createTime")
    values('操作步骤2', '期待结果2', case_id, 2, false, false, NOW());
    insert into "TstCaseStep" (opt, expect, "caseId", ordr, disabled, deleted, "createTime")
    values('操作步骤3', '期待结果3', case_id, 3, false, false, NOW());

    -- 初始化问题类型
    insert into "IsuType"("value",label,ordr,"orgId","defaultVal","buildIn",disabled,deleted,"createTime") 
        select d."value",d.label,d.ordr,p_org_id,d."defaultVal",true,d.disabled,d.deleted,NOW() from "IsuTypeDefine" d;
    select id from "IsuType" where "defaultVal" = true and "orgId"=p_org_id into issue_type_id;
		   
    insert into "IsuTypeSolution" (name, "orgId","defaultVal","buildIn", disabled, deleted, "createTime")
    values('默认问题类型方案', p_org_id, true, true, false, false, NOW());
    select max(id) from "IsuTypeSolution" into issue_type_solution_id;

    insert into "IsuTypeSolutionItem" ("typeId", "solutionId", "orgId")
    select d.id,issue_type_solution_id,p_org_id from "IsuType" d where d."orgId"=p_org_id;

    -- 初始化问题优先级
    insert into "IsuPriority"("value",label,ordr,"orgId","defaultVal","buildIn",disabled,deleted,"createTime") 
        select d."value",d.label,d.ordr,p_org_id,d."defaultVal",true,d.disabled,d.deleted,NOW() from "IsuPriorityDefine" d;
	select id from "IsuPriority" where "defaultVal" = true and "orgId"=p_org_id into issue_priority_id;
		   
    insert into "IsuPrioritySolution" (name, "orgId","defaultVal","buildIn", disabled, deleted, "createTime")
    values('默认问题优先级方案', p_org_id, true, true, false, false, NOW());
    select max(id) from "IsuPrioritySolution" into issue_priority_solution_id;

    insert into "IsuPrioritySolutionItem" ("priorityId", "solutionId", "orgId")
    select d.id,issue_priority_solution_id,p_org_id from "IsuPriority" d where d."orgId"=p_org_id;

    -- 初始化其他问题属性
    insert into "IsuStatus"("value",label,"categoryId",ordr,"orgId","defaultVal","finalVal","buildIn",disabled,deleted,"createTime") 
        select d."value",d.label,"categoryId",d.ordr,p_org_id,d."defaultVal",d."finalVal",true,d.disabled,d.deleted,NOW() from "IsuStatusDefine" d;
	select id from "IsuStatus" where "defaultVal" = true and "orgId"=p_org_id into issue_status_id;
		   
    insert into "IsuResolution"("value",label,ordr,"orgId","defaultVal","buildIn",disabled,deleted,"createTime") 
        select d."value",d.label,d.ordr,p_org_id,d."defaultVal",true,d.disabled,d.deleted,NOW() from "IsuResolutionDefine" d;
	select id from "IsuResolution" where "defaultVal" = true and "orgId"=p_org_id into issue_resolution_id;
		   
    insert into "IsuField"("colCode",label,"type",input, 
            "defaultShowInFilters","filterOrdr", 
            "defaultShowInColumns","columnOrdr",
            "defaultShowInPage","elemOrdr",
                readonly,"fullLine",required,"orgId",disabled,deleted,"createTime") 
    select d."colCode",d.label,d."type",d.input,
            d."defaultShowInFilters",d."filterOrdr", 
            d."defaultShowInColumns",d."columnOrdr",
            d."defaultShowInPage",d."elemOrdr",
                d.readonly,d."fullLine",d.required,p_org_id,d.disabled,d.deleted,NOW() 
        from "IsuFieldDefine" d where d."defaultShowInPage" IS NOT NULL;

    -- 初始化问题自定义属性
    insert into "CustomField"("colCode",label,"type",input,"textFormat","applyTo",rows,required,
        ordr,"orgId",readonly,"fullLine",disabled,deleted,"createTime") 
    select d."colCode",d.label,d."type",d.input,d."textFormat",d."applyTo",d.rows,d.required,
        d.ordr,p_org_id,readonly,"fullLine",d.disabled,d.deleted,NOW() from "CustomFieldDefine" d;

    PERFORM init_org_custom_field_option(p_org_id);
		   
    -- 创建示例缺陷
    INSERT INTO public."IsuIssue"(
		title, "orgId", "projectId", 
		"typeId", "statusId", "priorityId", "resolutionId", 
		"assigneeId", "creatorId", "reporterId", 
		"createTime", disabled, deleted, uuid, "extProp")
	VALUES ('示例缺陷', p_org_id, project_id, 
		   	issue_type_id, issue_status_id, issue_priority_id, issue_resolution_id,
		   p_user_id, p_user_id, p_user_id,
		   now(), false, false, 'uuid', '{}');

    -- 初始化问题页面
    insert into "IsuPage"(name,"orgId","defaultVal","buildIn",disabled,deleted,"createTime") 
            values ('默认界面', p_org_id, true,true,FALSE,FALSE,NOW());
    select max(id) from "IsuPage" into issue_page_id;
    
    insert into "IsuPageElement"("colCode",label,"type",input,"fullLine",required,
        ordr,readonly,"buildIn", "key","fieldId","pageId","orgId",
        disabled,deleted,"createTime")
    SELECT f."colCode",f.label,f."type",f.input,f."fullLine",f.required,
        f."elemOrdr",f.readonly,true, CONCAT('sys-', f.id),f.id,issue_page_id,p_org_id,
        false,false,NOW()
        from "IsuField" f where f."orgId" = p_org_id and f."defaultShowInPage" ORDER BY f."elemOrdr";

    insert into "IsuPageSolution"(name,"orgId","defaultVal",disabled,deleted,"createTime") 
        values ('默认界面方案', p_org_id,TRUE,FALSE,FALSE,NOW());
    select max(id) from "IsuPageSolution" into issue_page_solution_id;

    PERFORM init_org_issue_page_solution_item(issue_page_id, issue_page_solution_id, p_org_id);

    -- 初始化默认问题解决界面
    insert into "IsuPage"(name,"orgId","defaultVal","buildIn",disabled,deleted,"createTime") 
        values ('默认问题解决界面', p_org_id, false, true, FALSE,FALSE,NOW());
    select max(id) from "IsuPage" into issue_page_id;
    
    insert into "IsuPageElement"("colCode",label,"type",input,"fullLine",required,
        ordr,readonly,"buildIn","key","fieldId","pageId","orgId",
        disabled,deleted,"createTime")
    SELECT f."colCode",f.label,f."type",f.input,f."fullLine",f.required,
        f."elemOrdr",f.readonly,true,CONCAT('sys-', f.id),f.id,issue_page_id,p_org_id,
        false,false,NOW()
        from "IsuField" f where f."orgId" = p_org_id and f."colCode" LIKE 'resolution%' ORDER BY f."elemOrdr";

    -- 初始化问题工作流
    insert into "IsuWorkflow"(name,"orgId","defaultVal","buildIn",disabled,deleted,"createTime") 
    values ('默认工作流', p_org_id, true,true,FALSE,FALSE,NOW());
    select max(id) from "IsuWorkflow" into issue_workflow_id;

    insert into "IsuWorkflowSolution" (name, "orgId","defaultVal","buildIn", disabled, deleted, "createTime")
    values('默认工作流方案', p_org_id, true, true, false, false, NOW());
    select max(id) from "IsuWorkflowSolution" into issue_workflow_solution_id;

    insert into "IsuWorkflowSolutionItem" ("typeId", "workflowId", "solutionId", "orgId")
    select tp.id, wf.id, issue_workflow_solution_id, p_org_id from "IsuWorkflow" wf, "IsuType" tp
        where wf."orgId"=p_org_id and tp."orgId"=p_org_id
		order by tp.id;

    -- 工作流配置
    select ((select max(id) from "IsuStatus") - (select max(id) from "IsuStatusDefine")) into record_gap_to_define_table;
    -- 
    insert into "IsuWorkflowStatusRelation"("workflowId","statusId","orgId")
    SELECT issue_workflow_id,d."statusId"+record_gap_to_define_table,p_org_id
        from "IsuWorkflowStatusRelationDefine" d ORDER BY d.id;

    -- 
    insert into "IsuWorkflowTransition"(name,"srcStatusId","dictStatusId",
        "actionPageId",
        "workflowId","orgId",disabled,deleted,"createTime")
    SELECT d.name,d."srcStatusId"+record_gap_to_define_table,d."dictStatusId"+record_gap_to_define_table,
        case "isSolveIssue" 
            when true then issue_page_id
            else NULL
        end,
        issue_workflow_id, p_org_id,false,false,NOW()
        from "IsuWorkflowTransitionDefine" d ORDER BY d.id;

    insert into "IsuWorkflowTransitionProjectRoleRelation"(
        "projectRoleId","workflowTransitionId","workflowId","orgId")
    SELECT role.id, tran.id,issue_workflow_id,p_org_id from "TstProjectRole" role, "IsuWorkflowTransition" tran
        where role."orgId"=p_org_id AND tran."orgId"=p_org_id;

   -- 更新项目配置
   update "TstProject" set "issueTypeSolutionId"=issue_type_solution_id, 
             "issuePrioritySolutionId"=issue_priority_solution_id, 
             "issuePageSolutionId"=issue_page_solution_id, 
             "issueWorkflowSolutionId"=issue_workflow_solution_id
    WHERE id = project_id;
												
END;  
$$;


ALTER FUNCTION public.init_org(p_org_id integer, p_user_id integer) OWNER TO dbuser;

--
-- TOC entry 444 (class 1255 OID 46464)
-- Name: init_org_custom_field_option(integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.init_org_custom_field_option(p_org_id integer) RETURNS void
    LANGUAGE plpgsql
    AS $$  
declare  
    custom_field_define_id integer;
	
	rec_custom_field RECORD;
	cur_custom_fields CURSOR(p_org_id INTEGER) 
		FOR SELECT "id", "colCode", "input" from "CustomField" where "orgId" = p_org_id;
BEGIN  

	OPEN cur_custom_fields(p_org_id);
	
	LOOP
      FETCH cur_custom_fields INTO rec_custom_field;
      EXIT WHEN NOT FOUND;
	  
	  RAISE NOTICE 'rec_custom_field.input=%', rec_custom_field."input";
	  
	   if (rec_custom_field."input"='dropdown' OR rec_custom_field."input"='radio'
			OR rec_custom_field."input"='checkbox' OR rec_custom_field."input"='multi_select') then 

		   select id from "CustomFieldDefine" WHERE "colCode"=rec_custom_field."colCode" 
			into custom_field_define_id;
	    
	       insert into "CustomFieldOption"("label",ordr,"fieldId","orgId",
				"defaultVal","buildIn",disabled,deleted,"createTime") 
		     select d.label,d.ordr,rec_custom_field.id,p_org_id,
				  "defaultVal",true,d.disabled,d.deleted,NOW() from "CustomFieldOptionDefine" d
			    where d."fieldId" = custom_field_define_id;

	   end if;
	  
   END LOOP;

   CLOSE cur_custom_fields;
 
END;  
$$;


ALTER FUNCTION public.init_org_custom_field_option(p_org_id integer) OWNER TO dbuser;

--
-- TOC entry 438 (class 1255 OID 46465)
-- Name: init_org_issue_page_solution_item(integer, integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.init_org_issue_page_solution_item(p_issue_page_id integer, p_issue_page_solution_id integer, p_org_id integer) RETURNS void
    LANGUAGE plpgsql
    AS $$  
declare  
    issue_type_id integer;  
	cur_issue_type_ids CURSOR(p_org_id INTEGER) 
		FOR SELECT id from "IsuType" where "orgId" = p_org_id;
BEGIN  
	RAISE NOTICE 'v_cnt';
	
	OPEN cur_issue_type_ids(p_org_id);
	
	LOOP
      FETCH cur_issue_type_ids INTO issue_type_id;
      EXIT WHEN NOT FOUND;
	  
	  RAISE NOTICE 'issue_type_id=%', issue_type_id;
	  
		insert into "IsuPageSolutionItem"("typeId",opt,"pageId","solutionId","orgId") 
			values (issue_type_id, 'create',p_issue_page_id,p_issue_page_solution_id,p_org_id);

		insert into "IsuPageSolutionItem"("typeId",opt,"pageId","solutionId","orgId") 
			values (issue_type_id, 'edit',p_issue_page_id, p_issue_page_solution_id,p_org_id);

		insert into "IsuPageSolutionItem"("typeId",opt,"pageId","solutionId","orgId") 
			values (issue_type_id, 'view',p_issue_page_id, p_issue_page_solution_id,p_org_id);
	  
   END LOOP;

   CLOSE cur_issue_type_ids;
END;  
$$;


ALTER FUNCTION public.init_org_issue_page_solution_item(p_issue_page_id integer, p_issue_page_solution_id integer, p_org_id integer) OWNER TO dbuser;

--
-- TOC entry 445 (class 1255 OID 46466)
-- Name: init_user(integer, character varying); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.init_user(p_user_id integer, p_org_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$  
declare  
    var_org_id integer;  
BEGIN  
	 
    insert into "TstOrg" (name, disabled, deleted, "createTime") 
		values(p_org_name, false, false, NOW());
    select max(id) from "TstOrg" into var_org_id;

    update "TstUser" set "defaultOrgId" = var_org_id, 
						 "defaultOrgName" = p_org_name where id=p_user_id;

    RAISE NOTICE 'var_org_id = %', var_org_id;
			   
    PERFORM init_org(var_org_id, p_user_id);
																	
END;  
$$;


ALTER FUNCTION public.init_user(p_user_id integer, p_org_name character varying) OWNER TO dbuser;

--
-- TOC entry 411 (class 1255 OID 46467)
-- Name: remove_all(); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.remove_all() RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE
   tmp VARCHAR(512);
   DECLARE names CURSOR FOR 
    select tablename from pg_tables where schemaname='public';
BEGIN
FOR stmt IN names LOOP
     tmp := 'DROP TABLE '|| quote_ident(stmt.tablename) || ' CASCADE;';
	RAISE NOTICE 'notice: %', tmp;EXECUTE tmp;
	
END LOOP;
END;

$$;


ALTER FUNCTION public.remove_all() OWNER TO dbuser;

--
-- TOC entry 412 (class 1255 OID 46468)
-- Name: remove_all_tables(); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.remove_all_tables() RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE
   tmp VARCHAR(512);
   DECLARE names CURSOR FOR 
    select tablename from pg_tables where schemaname='public';
	
	DECLARE funcs CURSOR FOR 
		SELECT ns.nspname || '.' || proname 
			   || '(' || oidvectortypes(proargtypes) || ')' func
		FROM pg_proc INNER JOIN pg_namespace ns ON (pg_proc.pronamespace = ns.oid) 
		WHERE ns.nspname = 'public' and  proname!='remove_all_tables' order by proname  ;
	
BEGIN
FOR stmt IN names LOOP
     tmp := 'DROP TABLE '|| quote_ident(stmt.tablename) || ' CASCADE;';
	RAISE NOTICE 'notice: %', tmp;
	EXECUTE tmp;
END LOOP;

FOR stmt IN funcs LOOP
     tmp := 'DROP FUNCTION ' || stmt.func || ' CASCADE;';
	RAISE NOTICE 'notice: %', tmp;
	EXECUTE tmp;
END LOOP;

END;

$$;


ALTER FUNCTION public.remove_all_tables() OWNER TO dbuser;

--
-- TOC entry 418 (class 1255 OID 64111)
-- Name: remove_case_and_its_children(integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.remove_case_and_its_children(p_case_id integer, p_project_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  
declare  
    v_cnt integer;  
BEGIN  
	update "TstCase" set deleted = true where id = any (
		WITH RECURSIVE cs AS ( 
			 SELECT parent.* FROM "TstCase" parent WHERE parent.id = p_case_id
			 union  all 
			 SELECT child.* FROM "TstCase" child, cs WHERE child."pId" = cs.id 
		) 
		SELECT cs.id FROM cs where cs."projectId"=p_project_id ORDER BY cs.id
	);
	
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	RAISE NOTICE 'v_cnt = %', v_cnt;
	
    RETURN v_cnt;  
END;  
$$;


ALTER FUNCTION public.remove_case_and_its_children(p_case_id integer, p_project_id integer) OWNER TO dbuser;

--
-- TOC entry 410 (class 1255 OID 46470)
-- Name: remove_user_from_org(integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.remove_user_from_org(p_user_id integer, p_org_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  
declare  
	v_cnt integer;
BEGIN

	delete from "TstOrgUserRelation"
	where "userId"=p_user_id and "orgId"=p_org_id;
	
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	RAISE NOTICE 'v_cnt = %', v_cnt;
    RETURN v_cnt; 
	
    delete from "TstOrgRoleUserRelation"
	where "userId"=p_user_id 
		and "orgRoleId" in (select tmp.id from "TstOrgRole" tmp where tmp."orgId"=p_org_id);
	   
    delete from "TstOrgGroupUserRelation"
	where "userId"=p_user_id 
		and "orgGroupId" in (select tmp.id from "TstOrgGroup" tmp where tmp."orgId"=p_org_id); 

END;  
$$;


ALTER FUNCTION public.remove_user_from_org(p_user_id integer, p_org_id integer) OWNER TO dbuser;

--
-- TOC entry 413 (class 1255 OID 46471)
-- Name: test(integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.test(_p integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  
declare  
    v_cnt integer;  
BEGIN  
	INSERT INTO "Test" (name) 
		VALUES ('XYZ');
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	RAISE NOTICE 'v_cnt = %', v_cnt;
	
    RETURN v_cnt;  
END;  
$$;


ALTER FUNCTION public.test(_p integer) OWNER TO dbuser;

--
-- TOC entry 454 (class 1255 OID 48642)
-- Name: update_issue_tsv_content(); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.update_issue_tsv_content() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
   p_str VARCHAR;
BEGIN							   
	select 
	   string_agg(distinct(CASE WHEN fld.type = 'string' 
	            THEN replace(temp.val::TEXT, '"', '')
            --WHEN (fld.input = 'dropdown' OR fld.input = 'radio')
			--  THEN (select opt.label from "CustomFieldOption" opt
			--    where opt.id = temp.val::TEXT::integer)::text
            --WHEN fld.type = 'muti_select' OR fld.type = 'checkbox'
			--  THEN (select string_agg(opt.label, '') from "CustomFieldOption" opt
			--    where opt.id = any(array(SELECT * FROM regexp_split_to_array(temp.val::text, ','))::TEXT::integer[]))
            --ELSE ''
       END), ' ')		   
	from "IsuIssue" b, jsonb_each(coalesce(b."extProp",'{}')) as temp(key,val)
	join "CustomField" fld on fld."colCode" = temp.key::text			   
	WHERE b.id = NEW.id AND (fld.type = 'string' or fld.type = 'integer')
	into p_str;
	
	RAISE NOTICE 'NEW_ID: %', NEW.id;
	RAISE NOTICE 'p_str: %', p_str;
	
	update "IsuIssue" a
	set tsv_content = 
           setweight(to_tsvector('chinese_zh', coalesce(a.tag,'')), 'A') 
		|| setweight(to_tsvector('chinese_zh', coalesce(a.title,'')), 'B') 
		|| setweight(to_tsvector('chinese_zh', coalesce(p_str,'')), 'C')
		|| setweight(to_tsvector('chinese_zh', coalesce(a.descr,'')), 'D')
	WHERE a.id = NEW.id;
									 
	RETURN null;
END;

$$;


ALTER FUNCTION public.update_issue_tsv_content() OWNER TO dbuser;

--
-- TOC entry 447 (class 1255 OID 46472)
-- Name: update_workflow_statuses(integer, character varying); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.update_workflow_statuses(p_workflow_id integer, p_status_ids character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$  
declare  
	var_status_ids_for_insert integer[];
    var_status_ids_for_deleted integer[];
	var_trans_ids_for_deleted integer[];
BEGIN

	select array(SELECT rla."statusId" FROM "IsuWorkflowStatusRelation" rla
        WHERE rla."workflowId"=p_workflow_id
			AND rla."statusId" != ALL ( string_to_array( p_status_ids, ',')::int[] ) )
    INTO var_status_ids_for_deleted;
	
	RAISE NOTICE 'var_status_ids_for_deleted = %', var_status_ids_for_deleted;
	
	select array(SELECT tran.id FROM "IsuWorkflowTransition" tran
        WHERE tran."workflowId"=p_workflow_id
          AND ( tran."srcStatusId" = any (var_status_ids_for_deleted) 
				OR tran."dictStatusId" = any (var_status_ids_for_deleted) ) )
    INTO var_trans_ids_for_deleted;
	
	RAISE NOTICE 'var_trans_ids_for_deleted = %', var_trans_ids_for_deleted;
	
	--DELETE FROM "IsuWorkflowTransitionProjectRoleRelation" 
	--	WHERE "workflowTransitionId" = any (var_trans_ids_for_deleted);
	--DELETE FROM "IsuWorkflowTransition" WHERE id = any (var_trans_ids_for_deleted);
	
	DELETE FROM "IsuWorkflowStatusRelation" 
		WHERE "workflowId" = p_workflow_id;
						   
END;  
$$;


ALTER FUNCTION public.update_workflow_statuses(p_workflow_id integer, p_status_ids character varying) OWNER TO dbuser;

--
-- TOC entry 448 (class 1255 OID 46473)
-- Name: update_workflow_statuses(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.update_workflow_statuses(p_workflow_id integer, p_status_ids character varying, p_org_id integer) RETURNS integer
    LANGUAGE plpgsql
    AS $$  
declare  
    v_cnt integer;  
BEGIN  
	INSERT INTO "Test" (name) 
		VALUES ('XYZ');
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	RAISE NOTICE 'v_cnt = %', v_cnt;
	
    RETURN v_cnt;  
END;  
$$;


ALTER FUNCTION public.update_workflow_statuses(p_workflow_id integer, p_status_ids character varying, p_org_id integer) OWNER TO dbuser;

--
-- TOC entry 449 (class 1255 OID 46474)
-- Name: user_not_in_project(integer, integer); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.user_not_in_project(p_user_id integer, p_project_id integer) RETURNS boolean
    LANGUAGE plpgsql
    AS $$  
declare  
	v_cnt bigint;
BEGIN

    select count(u.id)
    from "TstUser" u
    where u.id = p_user_id and u.id in
	 (
	   select relation1."entityId" from "TstProjectRoleEntityRelation" relation1
	   where relation1.type = 'user' and relation1."projectId" = p_project_id
	   UNION
	   select relta."userId" from "TstOrgGroupUserRelation" relta
	   where relta."orgGroupId" in
			 (
			   select relation2."entityId" from "TstProjectRoleEntityRelation" relation2
			   where relation2.type = 'group' and relation2."projectId" = p_project_id
			 )
	 )
	 into v_cnt;
	 
	RAISE NOTICE 'v_cnt = %', v_cnt;
    RETURN (v_cnt < 1); 

END;  
$$;


ALTER FUNCTION public.user_not_in_project(p_user_id integer, p_project_id integer) OWNER TO dbuser;

--
-- TOC entry 2419 (class 3602 OID 48646)
-- Name: chinese_zh; Type: TEXT SEARCH CONFIGURATION; Schema: public; Owner: dbuser
--

CREATE TEXT SEARCH CONFIGURATION public.chinese_zh (
    PARSER = public.zhparser );

ALTER TEXT SEARCH CONFIGURATION public.chinese_zh
    ADD MAPPING FOR a WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.chinese_zh
    ADD MAPPING FOR e WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.chinese_zh
    ADD MAPPING FOR i WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.chinese_zh
    ADD MAPPING FOR l WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.chinese_zh
    ADD MAPPING FOR n WITH simple;

ALTER TEXT SEARCH CONFIGURATION public.chinese_zh
    ADD MAPPING FOR v WITH simple;


ALTER TEXT SEARCH CONFIGURATION public.chinese_zh OWNER TO dbuser;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 197 (class 1259 OID 46475)
-- Name: CustomField; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."CustomField" (
    id integer NOT NULL,
    "colCode" character varying(255),
    label character varying(255),
    type character varying(255),
    input character varying(255),
    "textFormat" character varying(255),
    "applyTo" character varying(255),
    rows integer,
    required boolean,
    readonly boolean,
    "fullLine" boolean,
    ordr integer,
    descr character varying(255),
    "buildIn" boolean,
    "orgId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."CustomField" OWNER TO dbuser;

--
-- TOC entry 198 (class 1259 OID 46481)
-- Name: CustomFieldDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."CustomFieldDefine" (
    id integer NOT NULL,
    "colCode" character varying(255),
    label character varying(255),
    type character varying(255),
    input character varying(255),
    "textFormat" character varying(255),
    "applyTo" character varying(255),
    rows integer,
    required boolean,
    readonly boolean,
    "fullLine" boolean,
    ordr integer,
    descr character varying(255),
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."CustomFieldDefine" OWNER TO dbuser;

--
-- TOC entry 199 (class 1259 OID 46487)
-- Name: CustomFieldDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."CustomFieldDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4765 (class 0 OID 0)
-- Dependencies: 199
-- Name: CustomFieldDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."CustomFieldDefine_id_seq" OWNED BY public."CustomFieldDefine".id;


--
-- TOC entry 200 (class 1259 OID 46489)
-- Name: CustomFieldInputTypeRelationDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."CustomFieldInputTypeRelationDefine" (
    id integer NOT NULL,
    "inputValue" character varying(255),
    "typeValue" character varying(255)
);


ALTER TABLE public."CustomFieldInputTypeRelationDefine" OWNER TO dbuser;

--
-- TOC entry 201 (class 1259 OID 46495)
-- Name: CustomFieldInputTypeRelationDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."CustomFieldInputTypeRelationDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldInputTypeRelationDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4766 (class 0 OID 0)
-- Dependencies: 201
-- Name: CustomFieldInputTypeRelationDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."CustomFieldInputTypeRelationDefine_id_seq" OWNED BY public."CustomFieldInputTypeRelationDefine".id;


--
-- TOC entry 202 (class 1259 OID 46497)
-- Name: CustomFieldIputDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."CustomFieldIputDefine" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    ordr integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."CustomFieldIputDefine" OWNER TO dbuser;

--
-- TOC entry 203 (class 1259 OID 46503)
-- Name: CustomFieldIputDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."CustomFieldIputDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldIputDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4767 (class 0 OID 0)
-- Dependencies: 203
-- Name: CustomFieldIputDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."CustomFieldIputDefine_id_seq" OWNED BY public."CustomFieldIputDefine".id;


--
-- TOC entry 204 (class 1259 OID 46505)
-- Name: CustomFieldOption; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."CustomFieldOption" (
    id integer NOT NULL,
    label character varying(255),
    descr character varying(255),
    ordr integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    "fieldId" integer,
    "orgId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."CustomFieldOption" OWNER TO dbuser;

--
-- TOC entry 205 (class 1259 OID 46511)
-- Name: CustomFieldOptionDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."CustomFieldOptionDefine" (
    id integer NOT NULL,
    label character varying(255),
    descr character varying(255),
    ordr integer,
    "defaultVal" boolean,
    "fieldId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."CustomFieldOptionDefine" OWNER TO dbuser;

--
-- TOC entry 206 (class 1259 OID 46517)
-- Name: CustomFieldOptionDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."CustomFieldOptionDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldOptionDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4768 (class 0 OID 0)
-- Dependencies: 206
-- Name: CustomFieldOptionDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."CustomFieldOptionDefine_id_seq" OWNED BY public."CustomFieldOptionDefine".id;


--
-- TOC entry 207 (class 1259 OID 46519)
-- Name: CustomFieldOption_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."CustomFieldOption_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldOption_id_seq" OWNER TO dbuser;

--
-- TOC entry 4769 (class 0 OID 0)
-- Dependencies: 207
-- Name: CustomFieldOption_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."CustomFieldOption_id_seq" OWNED BY public."CustomFieldOption".id;


--
-- TOC entry 208 (class 1259 OID 46521)
-- Name: CustomFieldTypeDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."CustomFieldTypeDefine" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(500),
    ordr integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."CustomFieldTypeDefine" OWNER TO dbuser;

--
-- TOC entry 209 (class 1259 OID 46527)
-- Name: CustomFieldTypeDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."CustomFieldTypeDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldTypeDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4770 (class 0 OID 0)
-- Dependencies: 209
-- Name: CustomFieldTypeDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."CustomFieldTypeDefine_id_seq" OWNED BY public."CustomFieldTypeDefine".id;


--
-- TOC entry 210 (class 1259 OID 46529)
-- Name: CustomField_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."CustomField_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomField_id_seq" OWNER TO dbuser;

--
-- TOC entry 4771 (class 0 OID 0)
-- Dependencies: 210
-- Name: CustomField_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."CustomField_id_seq" OWNED BY public."CustomField".id;


--
-- TOC entry 211 (class 1259 OID 46531)
-- Name: IsuAttachment; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuAttachment" (
    id integer NOT NULL,
    name character varying(255),
    title character varying(255),
    uri character varying(255),
    descr character varying(10000),
    "docType" character varying(255),
    "issueId" integer,
    "userId" integer,
    deleted boolean,
    disabled boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuAttachment" OWNER TO dbuser;

--
-- TOC entry 212 (class 1259 OID 46537)
-- Name: IsuAttachment_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuAttachment_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuAttachment_id_seq" OWNER TO dbuser;

--
-- TOC entry 4772 (class 0 OID 0)
-- Dependencies: 212
-- Name: IsuAttachment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuAttachment_id_seq" OWNED BY public."IsuAttachment".id;


--
-- TOC entry 213 (class 1259 OID 46539)
-- Name: IsuComments; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuComments" (
    id integer NOT NULL,
    summary character varying(255),
    content character varying(255),
    "issueId" integer,
    "userId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."IsuComments" OWNER TO dbuser;

--
-- TOC entry 214 (class 1259 OID 46545)
-- Name: IsuComments_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuComments_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuComments_id_seq" OWNER TO dbuser;

--
-- TOC entry 4773 (class 0 OID 0)
-- Dependencies: 214
-- Name: IsuComments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuComments_id_seq" OWNED BY public."IsuComments".id;


--
-- TOC entry 215 (class 1259 OID 46547)
-- Name: IsuCustomFieldSolution; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuCustomFieldSolution" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(255),
    "orgId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."IsuCustomFieldSolution" OWNER TO dbuser;

--
-- TOC entry 216 (class 1259 OID 46553)
-- Name: IsuCustomFieldSolutionFieldRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuCustomFieldSolutionFieldRelation" (
    "solutionId" integer,
    "fieldId" integer
);


ALTER TABLE public."IsuCustomFieldSolutionFieldRelation" OWNER TO dbuser;

--
-- TOC entry 217 (class 1259 OID 46556)
-- Name: IsuCustomFieldSolutionProjectRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuCustomFieldSolutionProjectRelation" (
    "solutionId" integer NOT NULL,
    "orgId" integer,
    "projectId" integer NOT NULL
);


ALTER TABLE public."IsuCustomFieldSolutionProjectRelation" OWNER TO dbuser;

--
-- TOC entry 218 (class 1259 OID 46559)
-- Name: IsuCustomFieldSolution_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuCustomFieldSolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuCustomFieldSolution_id_seq" OWNER TO dbuser;

--
-- TOC entry 4774 (class 0 OID 0)
-- Dependencies: 218
-- Name: IsuCustomFieldSolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuCustomFieldSolution_id_seq" OWNED BY public."IsuCustomFieldSolution".id;


--
-- TOC entry 219 (class 1259 OID 46561)
-- Name: IsuDocument; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuDocument" (
    id integer NOT NULL,
    "createTime" timestamp without time zone,
    deleted boolean,
    disabled boolean,
    "updateTime" timestamp without time zone,
    version integer,
    descr character varying(10000),
    "docType" character varying(255),
    "issueId" integer,
    title character varying(255),
    uri character varying(255),
    "userId" integer
);


ALTER TABLE public."IsuDocument" OWNER TO dbuser;

--
-- TOC entry 220 (class 1259 OID 46567)
-- Name: IsuDocument_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuDocument_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuDocument_id_seq" OWNER TO dbuser;

--
-- TOC entry 4775 (class 0 OID 0)
-- Dependencies: 220
-- Name: IsuDocument_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuDocument_id_seq" OWNED BY public."IsuDocument".id;


--
-- TOC entry 221 (class 1259 OID 46569)
-- Name: IsuField; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuField" (
    id integer NOT NULL,
    "colCode" character varying(255),
    label character varying(255),
    type character varying(255),
    input character varying(255),
    "fullLine" boolean,
    required boolean,
    "defaultShowInFilters" boolean,
    "filterOrdr" integer,
    "defaultShowInColumns" boolean,
    "columnOrdr" integer,
    "defaultShowInPage" boolean,
    "elemOrdr" integer,
    readonly boolean,
    "orgId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuField" OWNER TO dbuser;

--
-- TOC entry 222 (class 1259 OID 46575)
-- Name: IsuFieldCodeToTableDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuFieldCodeToTableDefine" (
    id integer NOT NULL,
    "colCode" character varying(255),
    "table" character varying(255),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuFieldCodeToTableDefine" OWNER TO dbuser;

--
-- TOC entry 223 (class 1259 OID 46581)
-- Name: IsuFieldCodeToTableDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuFieldCodeToTableDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuFieldCodeToTableDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4776 (class 0 OID 0)
-- Dependencies: 223
-- Name: IsuFieldCodeToTableDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuFieldCodeToTableDefine_id_seq" OWNED BY public."IsuFieldCodeToTableDefine".id;


--
-- TOC entry 224 (class 1259 OID 46583)
-- Name: IsuFieldDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuFieldDefine" (
    id integer NOT NULL,
    "colCode" character varying(255),
    label character varying(255),
    type character varying(255),
    input character varying(255),
    "defaultShowInFilters" boolean,
    "filterOrdr" integer,
    "defaultShowInColumns" boolean,
    "columnOrdr" integer,
    "defaultShowInPage" boolean,
    "elemOrdr" integer,
    readonly boolean,
    "fullLine" boolean,
    required boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuFieldDefine" OWNER TO dbuser;

--
-- TOC entry 225 (class 1259 OID 46589)
-- Name: IsuFieldDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuFieldDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuFieldDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4777 (class 0 OID 0)
-- Dependencies: 225
-- Name: IsuFieldDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuFieldDefine_id_seq" OWNED BY public."IsuFieldDefine".id;


--
-- TOC entry 226 (class 1259 OID 46591)
-- Name: IsuField_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuField_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuField_id_seq" OWNER TO dbuser;

--
-- TOC entry 4778 (class 0 OID 0)
-- Dependencies: 226
-- Name: IsuField_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuField_id_seq" OWNED BY public."IsuField".id;


--
-- TOC entry 227 (class 1259 OID 46593)
-- Name: IsuHistory; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuHistory" (
    id integer NOT NULL,
    title character varying(255),
    descr character varying(1000),
    "issueId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuHistory" OWNER TO dbuser;

--
-- TOC entry 228 (class 1259 OID 46599)
-- Name: IsuHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuHistory_id_seq" OWNER TO dbuser;

--
-- TOC entry 4779 (class 0 OID 0)
-- Dependencies: 228
-- Name: IsuHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuHistory_id_seq" OWNED BY public."IsuHistory".id;


--
-- TOC entry 229 (class 1259 OID 46601)
-- Name: IsuIssue; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuIssue" (
    id integer NOT NULL,
    title character varying(500),
    "orgId" integer,
    "projectId" integer,
    "projectName" character varying(255),
    "typeId" integer,
    "statusId" integer,
    "priorityId" integer,
    "assigneeId" integer,
    "creatorId" integer,
    "reporterId" integer,
    "resolutionId" integer,
    "resolutionDescr" character varying(5000),
    "verId" integer,
    "envId" integer,
    "dueTime" timestamp without time zone,
    "resolveTime" timestamp without time zone,
    "setFinalTime" timestamp without time zone,
    tag character varying(500),
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean,
    uuid character varying(32),
    "extProp" jsonb,
    tsv_content tsvector,
    descr character varying(10000)
);


ALTER TABLE public."IsuIssue" OWNER TO dbuser;

--
-- TOC entry 230 (class 1259 OID 46607)
-- Name: IsuIssueExt; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuIssueExt" (
    pid integer NOT NULL,
    prop01 character varying(255),
    prop02 character varying(255),
    prop03 character varying(255),
    prop04 character varying(255),
    prop05 character varying(255),
    prop06 character varying(255),
    prop07 character varying(255),
    prop08 character varying(255),
    prop09 character varying(255),
    prop10 character varying(255),
    prop11 character varying(255),
    prop12 character varying(255),
    prop13 character varying(255),
    prop14 character varying(255),
    prop15 character varying(255),
    prop16 character varying(255),
    prop17 character varying(255),
    prop18 character varying(255),
    prop19 character varying(255),
    prop20 character varying(255),
    prop21 character varying(255),
    prop22 character varying(255),
    prop23 character varying(255),
    prop24 character varying(255),
    prop25 character varying(255),
    prop26 character varying(255),
    prop27 character varying(255),
    prop28 character varying(255),
    prop29 character varying(255),
    prop30 character varying(255),
    prop31 character varying(255),
    prop32 character varying(255),
    prop33 character varying(255),
    prop34 character varying(255),
    prop35 character varying(255),
    prop36 character varying(255),
    prop37 character varying(255),
    prop38 character varying(255),
    prop39 character varying(255),
    prop40 character varying(255),
    prop41 character varying(255),
    prop42 character varying(255),
    prop43 character varying(255),
    prop44 character varying(255),
    prop45 character varying(255),
    prop46 character varying(255),
    prop47 character varying(255),
    prop48 character varying(255),
    prop49 character varying(255),
    prop50 character varying(255)
);


ALTER TABLE public."IsuIssueExt" OWNER TO dbuser;

--
-- TOC entry 231 (class 1259 OID 46613)
-- Name: IsuIssueExt_pid_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuIssueExt_pid_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuIssueExt_pid_seq" OWNER TO dbuser;

--
-- TOC entry 4780 (class 0 OID 0)
-- Dependencies: 231
-- Name: IsuIssueExt_pid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuIssueExt_pid_seq" OWNED BY public."IsuIssueExt".pid;


--
-- TOC entry 232 (class 1259 OID 46615)
-- Name: IsuIssue_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuIssue_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuIssue_id_seq" OWNER TO dbuser;

--
-- TOC entry 4781 (class 0 OID 0)
-- Dependencies: 232
-- Name: IsuIssue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuIssue_id_seq" OWNED BY public."IsuIssue".id;


--
-- TOC entry 233 (class 1259 OID 46617)
-- Name: IsuLink; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuLink" (
    id integer NOT NULL,
    "reasonId" integer,
    "reasonName" character varying(255),
    "srcIssueId" integer,
    "dictIssueId" integer,
    disabled integer,
    deleted integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuLink" OWNER TO dbuser;

--
-- TOC entry 234 (class 1259 OID 46620)
-- Name: IsuLinkReasonDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuLinkReasonDefine" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuLinkReasonDefine" OWNER TO dbuser;

--
-- TOC entry 235 (class 1259 OID 46626)
-- Name: IsuLinkReasonDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuLinkReasonDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuLinkReasonDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4782 (class 0 OID 0)
-- Dependencies: 235
-- Name: IsuLinkReasonDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuLinkReasonDefine_id_seq" OWNED BY public."IsuLinkReasonDefine".id;


--
-- TOC entry 236 (class 1259 OID 46628)
-- Name: IsuLink_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuLink_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuLink_id_seq" OWNER TO dbuser;

--
-- TOC entry 4783 (class 0 OID 0)
-- Dependencies: 236
-- Name: IsuLink_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuLink_id_seq" OWNED BY public."IsuLink".id;


--
-- TOC entry 237 (class 1259 OID 46630)
-- Name: IsuNotification; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuNotification" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    "orgId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuNotification" OWNER TO dbuser;

--
-- TOC entry 238 (class 1259 OID 46636)
-- Name: IsuNotificationDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuNotificationDefine" (
    id integer NOT NULL,
    name character varying(255),
    code character varying(255),
    descr character varying(1000),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuNotificationDefine" OWNER TO dbuser;

--
-- TOC entry 239 (class 1259 OID 46642)
-- Name: IsuNotificationDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuNotificationDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuNotificationDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4784 (class 0 OID 0)
-- Dependencies: 239
-- Name: IsuNotificationDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuNotificationDefine_id_seq" OWNED BY public."IsuNotificationDefine".id;


--
-- TOC entry 240 (class 1259 OID 46644)
-- Name: IsuNotification_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuNotification_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuNotification_id_seq" OWNER TO dbuser;

--
-- TOC entry 4785 (class 0 OID 0)
-- Dependencies: 240
-- Name: IsuNotification_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuNotification_id_seq" OWNED BY public."IsuNotification".id;


--
-- TOC entry 241 (class 1259 OID 46646)
-- Name: IsuPage; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuPage" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    "orgId" integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuPage" OWNER TO dbuser;

--
-- TOC entry 242 (class 1259 OID 46652)
-- Name: IsuPageElement; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuPageElement" (
    id integer NOT NULL,
    "colCode" character varying(255),
    label character varying(255),
    type character varying(255),
    input character varying(255),
    "fullLine" boolean,
    required boolean,
    "buildIn" boolean,
    key character varying(255),
    "fieldId" integer,
    "pageId" integer,
    "orgId" integer,
    ordr integer,
    readonly boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuPageElement" OWNER TO dbuser;

--
-- TOC entry 243 (class 1259 OID 46658)
-- Name: IsuPageElement_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuPageElement_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPageElement_id_seq" OWNER TO dbuser;

--
-- TOC entry 4786 (class 0 OID 0)
-- Dependencies: 243
-- Name: IsuPageElement_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuPageElement_id_seq" OWNED BY public."IsuPageElement".id;


--
-- TOC entry 244 (class 1259 OID 46660)
-- Name: IsuPageSolution; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuPageSolution" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    "orgId" integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuPageSolution" OWNER TO dbuser;

--
-- TOC entry 245 (class 1259 OID 46666)
-- Name: IsuPageSolutionItem; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuPageSolutionItem" (
    id integer NOT NULL,
    "typeId" integer,
    opt character varying(255),
    "pageId" integer,
    "solutionId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuPageSolutionItem" OWNER TO dbuser;

--
-- TOC entry 246 (class 1259 OID 46669)
-- Name: IsuPageSolutionItem_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuPageSolutionItem_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPageSolutionItem_id_seq" OWNER TO dbuser;

--
-- TOC entry 4787 (class 0 OID 0)
-- Dependencies: 246
-- Name: IsuPageSolutionItem_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuPageSolutionItem_id_seq" OWNED BY public."IsuPageSolutionItem".id;


--
-- TOC entry 247 (class 1259 OID 46671)
-- Name: IsuPageSolution_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuPageSolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPageSolution_id_seq" OWNER TO dbuser;

--
-- TOC entry 4788 (class 0 OID 0)
-- Dependencies: 247
-- Name: IsuPageSolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuPageSolution_id_seq" OWNED BY public."IsuPageSolution".id;


--
-- TOC entry 248 (class 1259 OID 46673)
-- Name: IsuPage_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuPage_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPage_id_seq" OWNER TO dbuser;

--
-- TOC entry 4789 (class 0 OID 0)
-- Dependencies: 248
-- Name: IsuPage_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuPage_id_seq" OWNED BY public."IsuPage".id;


--
-- TOC entry 249 (class 1259 OID 46675)
-- Name: IsuPriority; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuPriority" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    descr character varying(1000),
    "defaultVal" boolean,
    "buildIn" boolean,
    ordr integer,
    "orgId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuPriority" OWNER TO dbuser;

--
-- TOC entry 250 (class 1259 OID 46681)
-- Name: IsuPriorityDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuPriorityDefine" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    descr character varying(1000),
    ordr integer,
    "defaultVal" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuPriorityDefine" OWNER TO dbuser;

--
-- TOC entry 251 (class 1259 OID 46687)
-- Name: IsuPriorityDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuPriorityDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPriorityDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4790 (class 0 OID 0)
-- Dependencies: 251
-- Name: IsuPriorityDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuPriorityDefine_id_seq" OWNED BY public."IsuPriorityDefine".id;


--
-- TOC entry 252 (class 1259 OID 46689)
-- Name: IsuPrioritySolution; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuPrioritySolution" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    "orgId" integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuPrioritySolution" OWNER TO dbuser;

--
-- TOC entry 253 (class 1259 OID 46695)
-- Name: IsuPrioritySolutionItem; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuPrioritySolutionItem" (
    "priorityId" integer,
    "solutionId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuPrioritySolutionItem" OWNER TO dbuser;

--
-- TOC entry 254 (class 1259 OID 46698)
-- Name: IsuPrioritySolution_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuPrioritySolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPrioritySolution_id_seq" OWNER TO dbuser;

--
-- TOC entry 4791 (class 0 OID 0)
-- Dependencies: 254
-- Name: IsuPrioritySolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuPrioritySolution_id_seq" OWNED BY public."IsuPrioritySolution".id;


--
-- TOC entry 255 (class 1259 OID 46700)
-- Name: IsuPriority_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuPriority_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPriority_id_seq" OWNER TO dbuser;

--
-- TOC entry 4792 (class 0 OID 0)
-- Dependencies: 255
-- Name: IsuPriority_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuPriority_id_seq" OWNED BY public."IsuPriority".id;


--
-- TOC entry 256 (class 1259 OID 46702)
-- Name: IsuQuery; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuQuery" (
    id integer NOT NULL,
    name character varying(255),
    rule character varying(1000),
    "orderBy" character varying(500),
    descr character varying(1000),
    "useTime" timestamp without time zone,
    "projectId" integer,
    "userId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuQuery" OWNER TO dbuser;

--
-- TOC entry 257 (class 1259 OID 46708)
-- Name: IsuQuery_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuQuery_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuQuery_id_seq" OWNER TO dbuser;

--
-- TOC entry 4793 (class 0 OID 0)
-- Dependencies: 257
-- Name: IsuQuery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuQuery_id_seq" OWNED BY public."IsuQuery".id;


--
-- TOC entry 258 (class 1259 OID 46710)
-- Name: IsuResolution; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuResolution" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    descr character varying(1000),
    ordr integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    "orgId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuResolution" OWNER TO dbuser;

--
-- TOC entry 259 (class 1259 OID 46716)
-- Name: IsuResolutionDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuResolutionDefine" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    "defaultVal" boolean,
    descr character varying(1000),
    ordr integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuResolutionDefine" OWNER TO dbuser;

--
-- TOC entry 260 (class 1259 OID 46722)
-- Name: IsuResolutionDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuResolutionDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuResolutionDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4794 (class 0 OID 0)
-- Dependencies: 260
-- Name: IsuResolutionDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuResolutionDefine_id_seq" OWNED BY public."IsuResolutionDefine".id;


--
-- TOC entry 261 (class 1259 OID 46724)
-- Name: IsuResolution_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuResolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuResolution_id_seq" OWNER TO dbuser;

--
-- TOC entry 4795 (class 0 OID 0)
-- Dependencies: 261
-- Name: IsuResolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuResolution_id_seq" OWNED BY public."IsuResolution".id;


--
-- TOC entry 262 (class 1259 OID 46726)
-- Name: IsuSeverity; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuSeverity" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    descr character varying(1000),
    "defaultVal" boolean,
    "buildIn" boolean,
    ordr integer,
    "orgId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuSeverity" OWNER TO dbuser;

--
-- TOC entry 263 (class 1259 OID 46732)
-- Name: IsuSeverityDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuSeverityDefine" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    descr character varying(1000),
    ordr integer,
    "defaultVal" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuSeverityDefine" OWNER TO dbuser;

--
-- TOC entry 264 (class 1259 OID 46738)
-- Name: IsuSeverityDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuSeverityDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuSeverityDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4796 (class 0 OID 0)
-- Dependencies: 264
-- Name: IsuSeverityDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuSeverityDefine_id_seq" OWNED BY public."IsuSeverityDefine".id;


--
-- TOC entry 265 (class 1259 OID 46740)
-- Name: IsuSeveritySolution; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuSeveritySolution" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    "orgId" integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuSeveritySolution" OWNER TO dbuser;

--
-- TOC entry 266 (class 1259 OID 46746)
-- Name: IsuSeveritySolutionItem; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuSeveritySolutionItem" (
    "severityId" integer,
    "solutionId" integer
);


ALTER TABLE public."IsuSeveritySolutionItem" OWNER TO dbuser;

--
-- TOC entry 267 (class 1259 OID 46749)
-- Name: IsuSeveritySolution_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuSeveritySolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuSeveritySolution_id_seq" OWNER TO dbuser;

--
-- TOC entry 4797 (class 0 OID 0)
-- Dependencies: 267
-- Name: IsuSeveritySolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuSeveritySolution_id_seq" OWNED BY public."IsuSeveritySolution".id;


--
-- TOC entry 268 (class 1259 OID 46751)
-- Name: IsuSeverity_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuSeverity_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuSeverity_id_seq" OWNER TO dbuser;

--
-- TOC entry 4798 (class 0 OID 0)
-- Dependencies: 268
-- Name: IsuSeverity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuSeverity_id_seq" OWNED BY public."IsuSeverity".id;


--
-- TOC entry 269 (class 1259 OID 46753)
-- Name: IsuStatus; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuStatus" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    descr character varying(1000),
    ordr integer,
    "orgId" integer,
    "categoryId" integer,
    "defaultVal" boolean,
    "finalVal" boolean,
    "buildIn" boolean,
    "startTime" timestamp without time zone,
    "endTime" timestamp without time zone,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuStatus" OWNER TO dbuser;

--
-- TOC entry 270 (class 1259 OID 46759)
-- Name: IsuStatusCategoryDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuStatusCategoryDefine" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    ordr integer,
    disabled boolean,
    deleted boolean,
    "finalVal" boolean
);


ALTER TABLE public."IsuStatusCategoryDefine" OWNER TO dbuser;

--
-- TOC entry 271 (class 1259 OID 46765)
-- Name: IsuStatusCategoryDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuStatusCategoryDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuStatusCategoryDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4799 (class 0 OID 0)
-- Dependencies: 271
-- Name: IsuStatusCategoryDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuStatusCategoryDefine_id_seq" OWNED BY public."IsuStatusCategoryDefine".id;


--
-- TOC entry 272 (class 1259 OID 46767)
-- Name: IsuStatusDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuStatusDefine" (
    id integer NOT NULL,
    label character varying(255),
    value character varying(255),
    descr character varying(1000),
    "defaultVal" boolean,
    "finalVal" boolean,
    "categoryId" integer,
    ordr integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuStatusDefine" OWNER TO dbuser;

--
-- TOC entry 273 (class 1259 OID 46773)
-- Name: IsuStatusDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuStatusDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuStatusDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4800 (class 0 OID 0)
-- Dependencies: 273
-- Name: IsuStatusDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuStatusDefine_id_seq" OWNED BY public."IsuStatusDefine".id;


--
-- TOC entry 274 (class 1259 OID 46775)
-- Name: IsuStatus_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuStatus_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuStatus_id_seq" OWNER TO dbuser;

--
-- TOC entry 4801 (class 0 OID 0)
-- Dependencies: 274
-- Name: IsuStatus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuStatus_id_seq" OWNED BY public."IsuStatus".id;


--
-- TOC entry 275 (class 1259 OID 46777)
-- Name: IsuTag; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuTag" (
    id integer NOT NULL,
    name character varying(255),
    "orgId" integer,
    "userId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."IsuTag" OWNER TO dbuser;

--
-- TOC entry 276 (class 1259 OID 46780)
-- Name: IsuTagRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuTagRelation" (
    id integer NOT NULL,
    "issueId" integer,
    "tagId" integer
);


ALTER TABLE public."IsuTagRelation" OWNER TO dbuser;

--
-- TOC entry 277 (class 1259 OID 46783)
-- Name: IsuTagRelation_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuTagRelation_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuTagRelation_id_seq" OWNER TO dbuser;

--
-- TOC entry 4802 (class 0 OID 0)
-- Dependencies: 277
-- Name: IsuTagRelation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuTagRelation_id_seq" OWNED BY public."IsuTagRelation".id;


--
-- TOC entry 278 (class 1259 OID 46785)
-- Name: IsuTag_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuTag_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuTag_id_seq" OWNER TO dbuser;

--
-- TOC entry 4803 (class 0 OID 0)
-- Dependencies: 278
-- Name: IsuTag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuTag_id_seq" OWNED BY public."IsuTag".id;


--
-- TOC entry 279 (class 1259 OID 46787)
-- Name: IsuType; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuType" (
    id integer NOT NULL,
    value character varying(255),
    label character varying(255),
    descr character varying(1000),
    ordr integer,
    "orgId" integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuType" OWNER TO dbuser;

--
-- TOC entry 280 (class 1259 OID 46793)
-- Name: IsuTypeDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuTypeDefine" (
    id integer NOT NULL,
    value character varying(255),
    label character varying(255),
    descr character varying(1000),
    ordr integer,
    "defaultVal" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuTypeDefine" OWNER TO dbuser;

--
-- TOC entry 281 (class 1259 OID 46799)
-- Name: IsuTypeDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuTypeDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuTypeDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4804 (class 0 OID 0)
-- Dependencies: 281
-- Name: IsuTypeDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuTypeDefine_id_seq" OWNED BY public."IsuTypeDefine".id;


--
-- TOC entry 282 (class 1259 OID 46801)
-- Name: IsuTypeSolution; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuTypeSolution" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    "orgId" integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuTypeSolution" OWNER TO dbuser;

--
-- TOC entry 283 (class 1259 OID 46807)
-- Name: IsuTypeSolutionItem; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuTypeSolutionItem" (
    "typeId" integer,
    "solutionId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuTypeSolutionItem" OWNER TO dbuser;

--
-- TOC entry 284 (class 1259 OID 46810)
-- Name: IsuTypeSolution_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuTypeSolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuTypeSolution_id_seq" OWNER TO dbuser;

--
-- TOC entry 4805 (class 0 OID 0)
-- Dependencies: 284
-- Name: IsuTypeSolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuTypeSolution_id_seq" OWNED BY public."IsuTypeSolution".id;


--
-- TOC entry 285 (class 1259 OID 46812)
-- Name: IsuType_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuType_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuType_id_seq" OWNER TO dbuser;

--
-- TOC entry 4806 (class 0 OID 0)
-- Dependencies: 285
-- Name: IsuType_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuType_id_seq" OWNED BY public."IsuType".id;


--
-- TOC entry 286 (class 1259 OID 46814)
-- Name: IsuWatch; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuWatch" (
    id integer NOT NULL,
    "userId" integer,
    "issueId" integer
);


ALTER TABLE public."IsuWatch" OWNER TO dbuser;

--
-- TOC entry 287 (class 1259 OID 46817)
-- Name: IsuWatch_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuWatch_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWatch_id_seq" OWNER TO dbuser;

--
-- TOC entry 4807 (class 0 OID 0)
-- Dependencies: 287
-- Name: IsuWatch_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuWatch_id_seq" OWNED BY public."IsuWatch".id;


--
-- TOC entry 288 (class 1259 OID 46819)
-- Name: IsuWorkflow; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuWorkflow" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    "buildIn" boolean,
    "orgId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean,
    "defaultVal" boolean
);


ALTER TABLE public."IsuWorkflow" OWNER TO dbuser;

--
-- TOC entry 289 (class 1259 OID 46825)
-- Name: IsuWorkflowSolution; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuWorkflowSolution" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    "orgId" integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuWorkflowSolution" OWNER TO dbuser;

--
-- TOC entry 290 (class 1259 OID 46831)
-- Name: IsuWorkflowSolutionItem; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuWorkflowSolutionItem" (
    id integer NOT NULL,
    "typeId" integer,
    "workflowId" integer,
    "solutionId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuWorkflowSolutionItem" OWNER TO dbuser;

--
-- TOC entry 291 (class 1259 OID 46834)
-- Name: IsuWorkflowSolutionItem_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuWorkflowSolutionItem_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowSolutionItem_id_seq" OWNER TO dbuser;

--
-- TOC entry 4808 (class 0 OID 0)
-- Dependencies: 291
-- Name: IsuWorkflowSolutionItem_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuWorkflowSolutionItem_id_seq" OWNED BY public."IsuWorkflowSolutionItem".id;


--
-- TOC entry 292 (class 1259 OID 46836)
-- Name: IsuWorkflowSolution_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuWorkflowSolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowSolution_id_seq" OWNER TO dbuser;

--
-- TOC entry 4809 (class 0 OID 0)
-- Dependencies: 292
-- Name: IsuWorkflowSolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuWorkflowSolution_id_seq" OWNED BY public."IsuWorkflowSolution".id;


--
-- TOC entry 293 (class 1259 OID 46838)
-- Name: IsuWorkflowStatusRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuWorkflowStatusRelation" (
    id integer NOT NULL,
    "workflowId" integer,
    "statusId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuWorkflowStatusRelation" OWNER TO dbuser;

--
-- TOC entry 294 (class 1259 OID 46841)
-- Name: IsuWorkflowStatusRelationDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuWorkflowStatusRelationDefine" (
    id integer NOT NULL,
    "workflowId" integer,
    "statusId" integer
);


ALTER TABLE public."IsuWorkflowStatusRelationDefine" OWNER TO dbuser;

--
-- TOC entry 295 (class 1259 OID 46844)
-- Name: IsuWorkflowStatusRelationDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuWorkflowStatusRelationDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowStatusRelationDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4810 (class 0 OID 0)
-- Dependencies: 295
-- Name: IsuWorkflowStatusRelationDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuWorkflowStatusRelationDefine_id_seq" OWNED BY public."IsuWorkflowStatusRelationDefine".id;


--
-- TOC entry 296 (class 1259 OID 46846)
-- Name: IsuWorkflowStatusRelation_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuWorkflowStatusRelation_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowStatusRelation_id_seq" OWNER TO dbuser;

--
-- TOC entry 4811 (class 0 OID 0)
-- Dependencies: 296
-- Name: IsuWorkflowStatusRelation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuWorkflowStatusRelation_id_seq" OWNED BY public."IsuWorkflowStatusRelation".id;


--
-- TOC entry 297 (class 1259 OID 46848)
-- Name: IsuWorkflowTransition; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuWorkflowTransition" (
    id integer NOT NULL,
    name character varying(255),
    "actionPageId" integer,
    "srcStatusId" integer,
    "dictStatusId" integer,
    "orgId" integer,
    ordr integer,
    "workflowId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuWorkflowTransition" OWNER TO dbuser;

--
-- TOC entry 298 (class 1259 OID 46851)
-- Name: IsuWorkflowTransitionDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuWorkflowTransitionDefine" (
    id integer NOT NULL,
    name character varying(255),
    "actionPageId" integer,
    "srcStatusId" integer,
    "dictStatusId" integer,
    "isSolveIssue" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."IsuWorkflowTransitionDefine" OWNER TO dbuser;

--
-- TOC entry 299 (class 1259 OID 46854)
-- Name: IsuWorkflowTransitionDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuWorkflowTransitionDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowTransitionDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4812 (class 0 OID 0)
-- Dependencies: 299
-- Name: IsuWorkflowTransitionDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuWorkflowTransitionDefine_id_seq" OWNED BY public."IsuWorkflowTransitionDefine".id;


--
-- TOC entry 300 (class 1259 OID 46856)
-- Name: IsuWorkflowTransitionProjectRoleRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."IsuWorkflowTransitionProjectRoleRelation" (
    id integer NOT NULL,
    "workflowId" integer,
    "workflowTransitionId" integer,
    "projectRoleId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuWorkflowTransitionProjectRoleRelation" OWNER TO dbuser;

--
-- TOC entry 301 (class 1259 OID 46859)
-- Name: IsuWorkflowTransitionProjectRoleRelation_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuWorkflowTransitionProjectRoleRelation_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowTransitionProjectRoleRelation_id_seq" OWNER TO dbuser;

--
-- TOC entry 4813 (class 0 OID 0)
-- Dependencies: 301
-- Name: IsuWorkflowTransitionProjectRoleRelation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuWorkflowTransitionProjectRoleRelation_id_seq" OWNED BY public."IsuWorkflowTransitionProjectRoleRelation".id;


--
-- TOC entry 302 (class 1259 OID 46861)
-- Name: IsuWorkflowTransition_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuWorkflowTransition_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowTransition_id_seq" OWNER TO dbuser;

--
-- TOC entry 4814 (class 0 OID 0)
-- Dependencies: 302
-- Name: IsuWorkflowTransition_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuWorkflowTransition_id_seq" OWNED BY public."IsuWorkflowTransition".id;


--
-- TOC entry 303 (class 1259 OID 46863)
-- Name: IsuWorkflow_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."IsuWorkflow_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflow_id_seq" OWNER TO dbuser;

--
-- TOC entry 4815 (class 0 OID 0)
-- Dependencies: 303
-- Name: IsuWorkflow_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."IsuWorkflow_id_seq" OWNED BY public."IsuWorkflow".id;


--
-- TOC entry 304 (class 1259 OID 46865)
-- Name: SysPrivilege; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."SysPrivilege" (
    id integer NOT NULL,
    code character varying(255),
    name character varying(255),
    descr character varying(255),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."SysPrivilege" OWNER TO dbuser;

--
-- TOC entry 305 (class 1259 OID 46871)
-- Name: SysPrivilege_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."SysPrivilege_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SysPrivilege_id_seq" OWNER TO dbuser;

--
-- TOC entry 4816 (class 0 OID 0)
-- Dependencies: 305
-- Name: SysPrivilege_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."SysPrivilege_id_seq" OWNED BY public."SysPrivilege".id;


--
-- TOC entry 306 (class 1259 OID 46873)
-- Name: SysRole; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."SysRole" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(255),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."SysRole" OWNER TO dbuser;

--
-- TOC entry 307 (class 1259 OID 46879)
-- Name: SysRolePrivilegeRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."SysRolePrivilegeRelation" (
    "roleId" integer NOT NULL,
    "privilegeId" integer NOT NULL
);


ALTER TABLE public."SysRolePrivilegeRelation" OWNER TO dbuser;

--
-- TOC entry 308 (class 1259 OID 46882)
-- Name: SysRoleUserRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."SysRoleUserRelation" (
    "roleId" integer NOT NULL,
    "userId" integer NOT NULL
);


ALTER TABLE public."SysRoleUserRelation" OWNER TO dbuser;

--
-- TOC entry 309 (class 1259 OID 46885)
-- Name: SysRole_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."SysRole_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SysRole_id_seq" OWNER TO dbuser;

--
-- TOC entry 4817 (class 0 OID 0)
-- Dependencies: 309
-- Name: SysRole_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."SysRole_id_seq" OWNED BY public."SysRole".id;


--
-- TOC entry 310 (class 1259 OID 46887)
-- Name: SysUser; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."SysUser" (
    id integer NOT NULL,
    name character varying(255),
    email character varying(255),
    password character varying(255),
    token character varying(255),
    avatar character varying(255),
    "verifyCode" character varying(255),
    "lastLoginTime" timestamp without time zone,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."SysUser" OWNER TO dbuser;

--
-- TOC entry 311 (class 1259 OID 46893)
-- Name: SysUser_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."SysUser_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SysUser_id_seq" OWNER TO dbuser;

--
-- TOC entry 4818 (class 0 OID 0)
-- Dependencies: 311
-- Name: SysUser_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."SysUser_id_seq" OWNED BY public."SysUser".id;


--
-- TOC entry 312 (class 1259 OID 46895)
-- Name: Test; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."Test" (
    id bigint NOT NULL,
    name character varying(255),
    "extProp" jsonb
);


ALTER TABLE public."Test" OWNER TO dbuser;

--
-- TOC entry 313 (class 1259 OID 46901)
-- Name: Test_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."Test_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Test_id_seq" OWNER TO dbuser;

--
-- TOC entry 4819 (class 0 OID 0)
-- Dependencies: 313
-- Name: Test_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."Test_id_seq" OWNED BY public."Test".id;


--
-- TOC entry 314 (class 1259 OID 46903)
-- Name: TstAlert; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstAlert" (
    id integer NOT NULL,
    title character varying(10000),
    uri character varying(255),
    type character varying(255),
    status character varying(255),
    "startTime" timestamp without time zone,
    "endTime" timestamp without time zone,
    "entityId" integer,
    "entityName" character varying(255),
    "isRead" boolean,
    "isSent" boolean,
    "assigneeId" integer,
    "userId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstAlert" OWNER TO dbuser;

--
-- TOC entry 315 (class 1259 OID 46909)
-- Name: TstAlert_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstAlert_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstAlert_id_seq" OWNER TO dbuser;

--
-- TOC entry 4820 (class 0 OID 0)
-- Dependencies: 315
-- Name: TstAlert_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstAlert_id_seq" OWNED BY public."TstAlert".id;


--
-- TOC entry 316 (class 1259 OID 46911)
-- Name: TstCase; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCase" (
    id integer NOT NULL,
    name character varying(255),
    content character varying(10000),
    objective character varying(1000),
    "contentType" character varying(255),
    estimate integer,
    "pId" integer,
    "isParent" boolean,
    ordr integer,
    "priorityId" integer,
    "typeId" integer,
    "reviewResult" boolean,
    "projectId" integer,
    "createById" integer,
    "updateById" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    "extProp" jsonb
);


ALTER TABLE public."TstCase" OWNER TO dbuser;

--
-- TOC entry 317 (class 1259 OID 46917)
-- Name: TstCaseAttachment; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseAttachment" (
    id integer NOT NULL,
    name character varying(255),
    title character varying(255),
    uri character varying(255),
    descr character varying(10000),
    "docType" character varying(255),
    "caseId" integer,
    "userId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."TstCaseAttachment" OWNER TO dbuser;

--
-- TOC entry 318 (class 1259 OID 46923)
-- Name: TstCaseAttachment_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseAttachment_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseAttachment_id_seq" OWNER TO dbuser;

--
-- TOC entry 4821 (class 0 OID 0)
-- Dependencies: 318
-- Name: TstCaseAttachment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseAttachment_id_seq" OWNED BY public."TstCaseAttachment".id;


--
-- TOC entry 319 (class 1259 OID 46925)
-- Name: TstCaseComments; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseComments" (
    id integer NOT NULL,
    summary character varying(255),
    content character varying(255),
    "caseId" integer,
    "userId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."TstCaseComments" OWNER TO dbuser;

--
-- TOC entry 320 (class 1259 OID 46931)
-- Name: TstCaseComments_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseComments_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseComments_id_seq" OWNER TO dbuser;

--
-- TOC entry 4822 (class 0 OID 0)
-- Dependencies: 320
-- Name: TstCaseComments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseComments_id_seq" OWNED BY public."TstCaseComments".id;


--
-- TOC entry 321 (class 1259 OID 46933)
-- Name: TstCaseExeStatusDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseExeStatusDefine" (
    id integer NOT NULL,
    value character varying(255),
    label character varying(255),
    descr character varying(255),
    ordr integer,
    "finalVal" boolean,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."TstCaseExeStatusDefine" OWNER TO dbuser;

--
-- TOC entry 322 (class 1259 OID 46939)
-- Name: TstCaseExeStatus_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseExeStatus_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseExeStatus_id_seq" OWNER TO dbuser;

--
-- TOC entry 4823 (class 0 OID 0)
-- Dependencies: 322
-- Name: TstCaseExeStatus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseExeStatus_id_seq" OWNED BY public."TstCaseExeStatusDefine".id;


--
-- TOC entry 323 (class 1259 OID 46941)
-- Name: TstCaseExeStatus; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseExeStatus" (
    id integer DEFAULT nextval('public."TstCaseExeStatus_id_seq"'::regclass) NOT NULL,
    value character varying(255),
    label character varying(255),
    descr character varying(255),
    ordr integer,
    "finalVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "orgId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstCaseExeStatus" OWNER TO dbuser;

--
-- TOC entry 324 (class 1259 OID 46948)
-- Name: TstCaseHistory; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseHistory" (
    id integer NOT NULL,
    title character varying(255),
    descr character varying(1000),
    "caseId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstCaseHistory" OWNER TO dbuser;

--
-- TOC entry 325 (class 1259 OID 46954)
-- Name: TstCaseHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseHistory_id_seq" OWNER TO dbuser;

--
-- TOC entry 4824 (class 0 OID 0)
-- Dependencies: 325
-- Name: TstCaseHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseHistory_id_seq" OWNED BY public."TstCaseHistory".id;


--
-- TOC entry 326 (class 1259 OID 46956)
-- Name: TstCaseInSuite; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseInSuite" (
    id integer NOT NULL,
    "caseId" integer,
    "isParent" boolean,
    ordr integer,
    "pId" integer,
    "projectId" integer,
    "suiteId" integer,
    deleted boolean,
    disabled boolean,
    "createBy" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstCaseInSuite" OWNER TO dbuser;

--
-- TOC entry 327 (class 1259 OID 46959)
-- Name: TstCaseInSuite_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseInSuite_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInSuite_id_seq" OWNER TO dbuser;

--
-- TOC entry 4825 (class 0 OID 0)
-- Dependencies: 327
-- Name: TstCaseInSuite_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseInSuite_id_seq" OWNED BY public."TstCaseInSuite".id;


--
-- TOC entry 328 (class 1259 OID 46961)
-- Name: TstCaseInTask; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseInTask" (
    id integer NOT NULL,
    "caseId" integer,
    "isParent" boolean,
    "pId" integer,
    ordr integer,
    "exeBy" integer,
    "exeTime" timestamp without time zone,
    status character varying(255),
    result character varying(255),
    "planId" integer,
    "projectId" integer,
    "taskId" integer,
    disabled boolean,
    deleted boolean,
    "createBy" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstCaseInTask" OWNER TO dbuser;

--
-- TOC entry 329 (class 1259 OID 46967)
-- Name: TstCaseInTaskAttachment; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseInTaskAttachment" (
    id integer NOT NULL,
    name character varying(255),
    title character varying(255),
    uri character varying(255),
    descr character varying(10000),
    "docType" character varying(255),
    "caseInTaskId" integer,
    "userId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."TstCaseInTaskAttachment" OWNER TO dbuser;

--
-- TOC entry 330 (class 1259 OID 46973)
-- Name: TstCaseInTaskAttachment_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseInTaskAttachment_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTaskAttachment_id_seq" OWNER TO dbuser;

--
-- TOC entry 4826 (class 0 OID 0)
-- Dependencies: 330
-- Name: TstCaseInTaskAttachment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseInTaskAttachment_id_seq" OWNED BY public."TstCaseInTaskAttachment".id;


--
-- TOC entry 331 (class 1259 OID 46975)
-- Name: TstCaseInTaskComments; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseInTaskComments" (
    id integer NOT NULL,
    summary character varying(255),
    content character varying(255),
    "caseInTaskId" integer,
    "userId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."TstCaseInTaskComments" OWNER TO dbuser;

--
-- TOC entry 332 (class 1259 OID 46981)
-- Name: TstCaseInTaskComments_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseInTaskComments_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTaskComments_id_seq" OWNER TO dbuser;

--
-- TOC entry 4827 (class 0 OID 0)
-- Dependencies: 332
-- Name: TstCaseInTaskComments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseInTaskComments_id_seq" OWNED BY public."TstCaseInTaskComments".id;


--
-- TOC entry 333 (class 1259 OID 46983)
-- Name: TstCaseInTaskHistory; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseInTaskHistory" (
    id integer NOT NULL,
    title character varying(255),
    descr character varying(1000),
    "caseId" integer,
    "caseInTaskId" integer,
    deleted boolean,
    disabled boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstCaseInTaskHistory" OWNER TO dbuser;

--
-- TOC entry 334 (class 1259 OID 46989)
-- Name: TstCaseInTaskHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseInTaskHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTaskHistory_id_seq" OWNER TO dbuser;

--
-- TOC entry 4828 (class 0 OID 0)
-- Dependencies: 334
-- Name: TstCaseInTaskHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseInTaskHistory_id_seq" OWNED BY public."TstCaseInTaskHistory".id;


--
-- TOC entry 335 (class 1259 OID 46991)
-- Name: TstCaseInTaskIssue; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseInTaskIssue" (
    id integer NOT NULL,
    "issueId" integer,
    "caseInTaskId" integer,
    "userId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."TstCaseInTaskIssue" OWNER TO dbuser;

--
-- TOC entry 336 (class 1259 OID 46994)
-- Name: TstCaseInTaskIssue_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseInTaskIssue_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTaskIssue_id_seq" OWNER TO dbuser;

--
-- TOC entry 4829 (class 0 OID 0)
-- Dependencies: 336
-- Name: TstCaseInTaskIssue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseInTaskIssue_id_seq" OWNED BY public."TstCaseInTaskIssue".id;


--
-- TOC entry 337 (class 1259 OID 46996)
-- Name: TstCaseInTask_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseInTask_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTask_id_seq" OWNER TO dbuser;

--
-- TOC entry 4830 (class 0 OID 0)
-- Dependencies: 337
-- Name: TstCaseInTask_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseInTask_id_seq" OWNED BY public."TstCaseInTask".id;


--
-- TOC entry 338 (class 1259 OID 46998)
-- Name: TstCasePriorityDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCasePriorityDefine" (
    id integer NOT NULL,
    value character varying(255),
    label character varying(255),
    descr character varying(255),
    ordr integer,
    "defaultVal" boolean,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."TstCasePriorityDefine" OWNER TO dbuser;

--
-- TOC entry 339 (class 1259 OID 47004)
-- Name: TstCasePriority_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCasePriority_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCasePriority_id_seq" OWNER TO dbuser;

--
-- TOC entry 4831 (class 0 OID 0)
-- Dependencies: 339
-- Name: TstCasePriority_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCasePriority_id_seq" OWNED BY public."TstCasePriorityDefine".id;


--
-- TOC entry 340 (class 1259 OID 47006)
-- Name: TstCasePriority; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCasePriority" (
    id integer DEFAULT nextval('public."TstCasePriority_id_seq"'::regclass) NOT NULL,
    value character varying(255),
    label character varying(255),
    descr character varying(255),
    ordr integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "orgId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstCasePriority" OWNER TO dbuser;

--
-- TOC entry 341 (class 1259 OID 47013)
-- Name: TstCaseStep; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseStep" (
    id integer NOT NULL,
    opt character varying(10000),
    expect character varying(10000),
    ordr integer,
    "caseId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstCaseStep" OWNER TO dbuser;

--
-- TOC entry 342 (class 1259 OID 47019)
-- Name: TstCaseStep_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseStep_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseStep_id_seq" OWNER TO dbuser;

--
-- TOC entry 4832 (class 0 OID 0)
-- Dependencies: 342
-- Name: TstCaseStep_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseStep_id_seq" OWNED BY public."TstCaseStep".id;


--
-- TOC entry 343 (class 1259 OID 47021)
-- Name: TstCaseTypeDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseTypeDefine" (
    id integer NOT NULL,
    value character varying(255),
    label character varying(255),
    descr character varying(255),
    ordr integer,
    "defaultVal" boolean,
    disabled boolean,
    deleted boolean
);


ALTER TABLE public."TstCaseTypeDefine" OWNER TO dbuser;

--
-- TOC entry 344 (class 1259 OID 47027)
-- Name: TstCaseType_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCaseType_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseType_id_seq" OWNER TO dbuser;

--
-- TOC entry 4833 (class 0 OID 0)
-- Dependencies: 344
-- Name: TstCaseType_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCaseType_id_seq" OWNED BY public."TstCaseTypeDefine".id;


--
-- TOC entry 345 (class 1259 OID 47029)
-- Name: TstCaseType; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstCaseType" (
    id integer DEFAULT nextval('public."TstCaseType_id_seq"'::regclass) NOT NULL,
    value character varying(255),
    label character varying(255),
    descr character varying(255),
    ordr integer,
    "defaultVal" boolean,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "orgId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstCaseType" OWNER TO dbuser;

--
-- TOC entry 346 (class 1259 OID 47036)
-- Name: TstCase_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstCase_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCase_id_seq" OWNER TO dbuser;

--
-- TOC entry 4834 (class 0 OID 0)
-- Dependencies: 346
-- Name: TstCase_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstCase_id_seq" OWNED BY public."TstCase".id;


--
-- TOC entry 347 (class 1259 OID 47038)
-- Name: TstDocument; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstDocument" (
    id integer NOT NULL,
    title character varying(255),
    version integer,
    descr character varying(10000),
    uri character varying(255),
    doc_type character varying(255),
    "userId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstDocument" OWNER TO dbuser;

--
-- TOC entry 348 (class 1259 OID 47044)
-- Name: TstDocument_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstDocument_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstDocument_id_seq" OWNER TO dbuser;

--
-- TOC entry 4835 (class 0 OID 0)
-- Dependencies: 348
-- Name: TstDocument_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstDocument_id_seq" OWNED BY public."TstDocument".id;


--
-- TOC entry 349 (class 1259 OID 47046)
-- Name: TstEmail; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstEmail" (
    id integer NOT NULL,
    subject character varying(255),
    content character varying(10000),
    "mailTo" character varying(255),
    "mailCc" character varying(255),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstEmail" OWNER TO dbuser;

--
-- TOC entry 350 (class 1259 OID 47052)
-- Name: TstEmail_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstEmail_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstEmail_id_seq" OWNER TO dbuser;

--
-- TOC entry 4836 (class 0 OID 0)
-- Dependencies: 350
-- Name: TstEmail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstEmail_id_seq" OWNED BY public."TstEmail".id;


--
-- TOC entry 351 (class 1259 OID 47054)
-- Name: TstEnv; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstEnv" (
    id integer NOT NULL,
    label character varying(255),
    descr character varying(1000),
    ordr integer,
    "defaultVal" boolean,
    "projectId" integer,
    "orgId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstEnv" OWNER TO dbuser;

--
-- TOC entry 352 (class 1259 OID 47060)
-- Name: TstEnv_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstEnv_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstEnv_id_seq" OWNER TO dbuser;

--
-- TOC entry 4837 (class 0 OID 0)
-- Dependencies: 352
-- Name: TstEnv_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstEnv_id_seq" OWNED BY public."TstEnv".id;


--
-- TOC entry 353 (class 1259 OID 47062)
-- Name: TstHistory; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstHistory" (
    id integer NOT NULL,
    title character varying(255),
    msg character varying(10000),
    descr character varying(1000),
    uri character varying(255),
    "entityType" character varying(255),
    "entityId" integer,
    "projectId" integer,
    "userId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstHistory" OWNER TO dbuser;

--
-- TOC entry 354 (class 1259 OID 47068)
-- Name: TstHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstHistory_id_seq" OWNER TO dbuser;

--
-- TOC entry 4838 (class 0 OID 0)
-- Dependencies: 354
-- Name: TstHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstHistory_id_seq" OWNED BY public."TstHistory".id;


--
-- TOC entry 355 (class 1259 OID 47070)
-- Name: TstModule; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstModule" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    ordr integer,
    "projectId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstModule" OWNER TO dbuser;

--
-- TOC entry 356 (class 1259 OID 47076)
-- Name: TstModule_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstModule_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstModule_id_seq" OWNER TO dbuser;

--
-- TOC entry 4839 (class 0 OID 0)
-- Dependencies: 356
-- Name: TstModule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstModule_id_seq" OWNED BY public."TstModule".id;


--
-- TOC entry 357 (class 1259 OID 47078)
-- Name: TstMsg; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstMsg" (
    id integer NOT NULL,
    title character varying(255),
    "isRead" boolean,
    "userId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstMsg" OWNER TO dbuser;

--
-- TOC entry 358 (class 1259 OID 47081)
-- Name: TstMsg_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstMsg_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstMsg_id_seq" OWNER TO dbuser;

--
-- TOC entry 4840 (class 0 OID 0)
-- Dependencies: 358
-- Name: TstMsg_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstMsg_id_seq" OWNED BY public."TstMsg".id;


--
-- TOC entry 359 (class 1259 OID 47083)
-- Name: TstOrg; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstOrg" (
    id integer NOT NULL,
    name character varying(255),
    website character varying(255),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstOrg" OWNER TO dbuser;

--
-- TOC entry 360 (class 1259 OID 47089)
-- Name: TstOrgGroup; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstOrgGroup" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(255),
    "orgId" integer,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstOrgGroup" OWNER TO dbuser;

--
-- TOC entry 361 (class 1259 OID 47095)
-- Name: TstOrgGroupUserRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstOrgGroupUserRelation" (
    "orgId" integer,
    "orgGroupId" integer,
    "userId" integer
);


ALTER TABLE public."TstOrgGroupUserRelation" OWNER TO dbuser;

--
-- TOC entry 362 (class 1259 OID 47098)
-- Name: TstOrgGroup_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstOrgGroup_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstOrgGroup_id_seq" OWNER TO dbuser;

--
-- TOC entry 4841 (class 0 OID 0)
-- Dependencies: 362
-- Name: TstOrgGroup_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstOrgGroup_id_seq" OWNED BY public."TstOrgGroup".id;


--
-- TOC entry 363 (class 1259 OID 47100)
-- Name: TstOrgPrivilegeDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstOrgPrivilegeDefine" (
    id integer NOT NULL,
    code character varying(255),
    name character varying(255),
    descr character varying(255),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    action character varying(255)
);


ALTER TABLE public."TstOrgPrivilegeDefine" OWNER TO dbuser;

--
-- TOC entry 364 (class 1259 OID 47106)
-- Name: TstOrgPrivilegeDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstOrgPrivilegeDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstOrgPrivilegeDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4842 (class 0 OID 0)
-- Dependencies: 364
-- Name: TstOrgPrivilegeDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstOrgPrivilegeDefine_id_seq" OWNED BY public."TstOrgPrivilegeDefine".id;


--
-- TOC entry 365 (class 1259 OID 47108)
-- Name: TstOrgRole; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstOrgRole" (
    id integer NOT NULL,
    name character varying(255),
    code character varying(255),
    descr character varying(255),
    "orgId" integer,
    "buildIn" boolean,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstOrgRole" OWNER TO dbuser;

--
-- TOC entry 366 (class 1259 OID 47114)
-- Name: TstOrgRoleGroupRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstOrgRoleGroupRelation" (
    "orgRoleId" integer NOT NULL,
    "orgGroupId" integer NOT NULL,
    "orgId" integer NOT NULL
);


ALTER TABLE public."TstOrgRoleGroupRelation" OWNER TO dbuser;

--
-- TOC entry 367 (class 1259 OID 47117)
-- Name: TstOrgRolePrivilegeRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstOrgRolePrivilegeRelation" (
    "orgId" integer NOT NULL,
    "orgRoleId" integer NOT NULL,
    "orgPrivilegeId" integer NOT NULL
);


ALTER TABLE public."TstOrgRolePrivilegeRelation" OWNER TO dbuser;

--
-- TOC entry 368 (class 1259 OID 47120)
-- Name: TstOrgRoleUserRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstOrgRoleUserRelation" (
    "orgRoleId" integer NOT NULL,
    "userId" integer NOT NULL,
    "orgId" integer NOT NULL
);


ALTER TABLE public."TstOrgRoleUserRelation" OWNER TO dbuser;

--
-- TOC entry 369 (class 1259 OID 47123)
-- Name: TstOrgRole_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstOrgRole_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstOrgRole_id_seq" OWNER TO dbuser;

--
-- TOC entry 4843 (class 0 OID 0)
-- Dependencies: 369
-- Name: TstOrgRole_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstOrgRole_id_seq" OWNED BY public."TstOrgRole".id;


--
-- TOC entry 370 (class 1259 OID 47125)
-- Name: TstOrgUserRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstOrgUserRelation" (
    "orgId" integer NOT NULL,
    "userId" integer NOT NULL
);


ALTER TABLE public."TstOrgUserRelation" OWNER TO dbuser;

--
-- TOC entry 371 (class 1259 OID 47128)
-- Name: TstOrg_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstOrg_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstOrg_id_seq" OWNER TO dbuser;

--
-- TOC entry 4844 (class 0 OID 0)
-- Dependencies: 371
-- Name: TstOrg_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstOrg_id_seq" OWNED BY public."TstOrg".id;


--
-- TOC entry 372 (class 1259 OID 47130)
-- Name: TstPlan; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstPlan" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    estimate integer,
    "startTime" timestamp without time zone,
    "endTime" timestamp without time zone,
    status character varying(255),
    "projectId" integer,
    "verId" integer,
    "userId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstPlan" OWNER TO dbuser;

--
-- TOC entry 373 (class 1259 OID 47136)
-- Name: TstPlan_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstPlan_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstPlan_id_seq" OWNER TO dbuser;

--
-- TOC entry 4845 (class 0 OID 0)
-- Dependencies: 373
-- Name: TstPlan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstPlan_id_seq" OWNED BY public."TstPlan".id;


--
-- TOC entry 374 (class 1259 OID 47138)
-- Name: TstProject; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstProject" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    type character varying(255),
    "issueTypeSolutionId" integer,
    "issuePrioritySolutionId" integer,
    "issuePageSolutionId" integer,
    "issueWorkflowSolutionId" integer,
    "lastAccessTime" timestamp without time zone,
    "orgId" integer,
    "parentId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstProject" OWNER TO dbuser;

--
-- TOC entry 375 (class 1259 OID 47144)
-- Name: TstProjectAccessHistory; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstProjectAccessHistory" (
    id integer NOT NULL,
    "lastAccessTime" timestamp without time zone,
    "orgId" integer,
    "prjId" integer,
    "prjName" character varying(255),
    "userId" integer,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstProjectAccessHistory" OWNER TO dbuser;

--
-- TOC entry 376 (class 1259 OID 47147)
-- Name: TstProjectAccessHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstProjectAccessHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstProjectAccessHistory_id_seq" OWNER TO dbuser;

--
-- TOC entry 4846 (class 0 OID 0)
-- Dependencies: 376
-- Name: TstProjectAccessHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstProjectAccessHistory_id_seq" OWNED BY public."TstProjectAccessHistory".id;


--
-- TOC entry 377 (class 1259 OID 47149)
-- Name: TstProjectPrivilegeDefine; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstProjectPrivilegeDefine" (
    id integer NOT NULL,
    code character varying(255),
    name character varying(255),
    action character varying(255),
    "actionName" character varying(255),
    descr character varying(255),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstProjectPrivilegeDefine" OWNER TO dbuser;

--
-- TOC entry 378 (class 1259 OID 47155)
-- Name: TstProjectPrivilegeDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstProjectPrivilegeDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstProjectPrivilegeDefine_id_seq" OWNER TO dbuser;

--
-- TOC entry 4847 (class 0 OID 0)
-- Dependencies: 378
-- Name: TstProjectPrivilegeDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstProjectPrivilegeDefine_id_seq" OWNED BY public."TstProjectPrivilegeDefine".id;


--
-- TOC entry 379 (class 1259 OID 47157)
-- Name: TstProjectRole; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstProjectRole" (
    id integer NOT NULL,
    code character varying(255),
    name character varying(255),
    descr character varying(255),
    "buildIn" boolean,
    "orgId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstProjectRole" OWNER TO dbuser;

--
-- TOC entry 380 (class 1259 OID 47163)
-- Name: TstProjectRoleEntityRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstProjectRoleEntityRelation" (
    "entityId" integer,
    "orgId" integer,
    "projectId" integer,
    "projectRoleId" integer,
    type character varying(255)
);


ALTER TABLE public."TstProjectRoleEntityRelation" OWNER TO dbuser;

--
-- TOC entry 381 (class 1259 OID 47166)
-- Name: TstProjectRolePriviledgeRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstProjectRolePriviledgeRelation" (
    "projectPrivilegeDefineId" integer,
    "projectRoleId" integer,
    "orgId" integer
);


ALTER TABLE public."TstProjectRolePriviledgeRelation" OWNER TO dbuser;

--
-- TOC entry 382 (class 1259 OID 47169)
-- Name: TstProjectRole_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstProjectRole_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstProjectRole_id_seq" OWNER TO dbuser;

--
-- TOC entry 4848 (class 0 OID 0)
-- Dependencies: 382
-- Name: TstProjectRole_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstProjectRole_id_seq" OWNED BY public."TstProjectRole".id;


--
-- TOC entry 383 (class 1259 OID 47171)
-- Name: TstProject_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstProject_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstProject_id_seq" OWNER TO dbuser;

--
-- TOC entry 4849 (class 0 OID 0)
-- Dependencies: 383
-- Name: TstProject_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstProject_id_seq" OWNED BY public."TstProject".id;


--
-- TOC entry 384 (class 1259 OID 47173)
-- Name: TstSuite; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstSuite" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    estimate integer,
    ordr integer,
    "projectId" integer,
    "userId" integer,
    "caseProjectId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstSuite" OWNER TO dbuser;

--
-- TOC entry 385 (class 1259 OID 47179)
-- Name: TstSuite_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstSuite_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstSuite_id_seq" OWNER TO dbuser;

--
-- TOC entry 4850 (class 0 OID 0)
-- Dependencies: 385
-- Name: TstSuite_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstSuite_id_seq" OWNED BY public."TstSuite".id;


--
-- TOC entry 386 (class 1259 OID 47181)
-- Name: TstTask; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstTask" (
    id integer NOT NULL,
    name character varying(255),
    descr character varying(1000),
    estimate integer,
    status character varying(255),
    "projectId" integer,
    "caseProjectId" integer,
    "planId" integer,
    "userId" integer,
    "envId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstTask" OWNER TO dbuser;

--
-- TOC entry 387 (class 1259 OID 47187)
-- Name: TstTaskAssigneeRelation; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstTaskAssigneeRelation" (
    "taskId" integer NOT NULL,
    "assigneeId" integer NOT NULL
);


ALTER TABLE public."TstTaskAssigneeRelation" OWNER TO dbuser;

--
-- TOC entry 388 (class 1259 OID 47190)
-- Name: TstTask_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstTask_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstTask_id_seq" OWNER TO dbuser;

--
-- TOC entry 4851 (class 0 OID 0)
-- Dependencies: 388
-- Name: TstTask_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstTask_id_seq" OWNED BY public."TstTask".id;


--
-- TOC entry 389 (class 1259 OID 47192)
-- Name: TstThread; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstThread" (
    id integer NOT NULL,
    content character varying(10000),
    "authorId" integer,
    "parentId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstThread" OWNER TO dbuser;

--
-- TOC entry 390 (class 1259 OID 47198)
-- Name: TstThread_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstThread_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstThread_id_seq" OWNER TO dbuser;

--
-- TOC entry 4852 (class 0 OID 0)
-- Dependencies: 390
-- Name: TstThread_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstThread_id_seq" OWNED BY public."TstThread".id;


--
-- TOC entry 391 (class 1259 OID 47200)
-- Name: TstUser; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstUser" (
    id integer NOT NULL,
    email character varying(255),
    nickname character varying(255),
    password character varying(255),
    phone character varying(255),
    avatar character varying(255),
    "defaultOrgId" integer,
    "defaultOrgName" character varying(255),
    "defaultPrjId" integer,
    "defaultPrjName" character varying(255),
    salt character varying(255),
    "verifyCode" character varying(255),
    "lastLoginTime" timestamp without time zone,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone,
    locked boolean
);


ALTER TABLE public."TstUser" OWNER TO dbuser;

--
-- TOC entry 392 (class 1259 OID 47206)
-- Name: TstUserSettings; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstUserSettings" (
    "leftSizeDesign" integer,
    "leftSizeExe" integer,
    "leftSizeIssue" integer,
    "issueView" character varying(255),
    "issueColumns" character varying(1000),
    "issueFields" character varying(1000),
    tql character varying(5000),
    "userId" integer NOT NULL
);


ALTER TABLE public."TstUserSettings" OWNER TO dbuser;

--
-- TOC entry 393 (class 1259 OID 47212)
-- Name: TstUserVerifyCode; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstUserVerifyCode" (
    id integer NOT NULL,
    code character varying(255),
    "expireTime" timestamp without time zone,
    "userId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstUserVerifyCode" OWNER TO dbuser;

--
-- TOC entry 394 (class 1259 OID 47215)
-- Name: TstUserVerifyCode_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstUserVerifyCode_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstUserVerifyCode_id_seq" OWNER TO dbuser;

--
-- TOC entry 4853 (class 0 OID 0)
-- Dependencies: 394
-- Name: TstUserVerifyCode_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstUserVerifyCode_id_seq" OWNED BY public."TstUserVerifyCode".id;


--
-- TOC entry 395 (class 1259 OID 47217)
-- Name: TstUser_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstUser_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstUser_id_seq" OWNER TO dbuser;

--
-- TOC entry 4854 (class 0 OID 0)
-- Dependencies: 395
-- Name: TstUser_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstUser_id_seq" OWNED BY public."TstUser".id;


--
-- TOC entry 396 (class 1259 OID 47219)
-- Name: TstVer; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public."TstVer" (
    id integer NOT NULL,
    label character varying(255),
    descr character varying(1000),
    status character varying(255),
    "startTime" timestamp without time zone,
    "endTime" timestamp without time zone,
    "defaultVal" boolean,
    ordr integer,
    "projectId" integer,
    "orgId" integer,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstVer" OWNER TO dbuser;

--
-- TOC entry 397 (class 1259 OID 47225)
-- Name: TstVer_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public."TstVer_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstVer_id_seq" OWNER TO dbuser;

--
-- TOC entry 4855 (class 0 OID 0)
-- Dependencies: 397
-- Name: TstVer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public."TstVer_id_seq" OWNED BY public."TstVer".id;


--
-- TOC entry 3780 (class 2604 OID 47230)
-- Name: CustomField id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomField" ALTER COLUMN id SET DEFAULT nextval('public."CustomField_id_seq"'::regclass);


--
-- TOC entry 3781 (class 2604 OID 47231)
-- Name: CustomFieldDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldDefine_id_seq"'::regclass);


--
-- TOC entry 3782 (class 2604 OID 47232)
-- Name: CustomFieldInputTypeRelationDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldInputTypeRelationDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldInputTypeRelationDefine_id_seq"'::regclass);


--
-- TOC entry 3783 (class 2604 OID 47233)
-- Name: CustomFieldIputDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldIputDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldIputDefine_id_seq"'::regclass);


--
-- TOC entry 3784 (class 2604 OID 47234)
-- Name: CustomFieldOption id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldOption" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldOption_id_seq"'::regclass);


--
-- TOC entry 3785 (class 2604 OID 47235)
-- Name: CustomFieldOptionDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldOptionDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldOptionDefine_id_seq"'::regclass);


--
-- TOC entry 3786 (class 2604 OID 47236)
-- Name: CustomFieldTypeDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldTypeDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldTypeDefine_id_seq"'::regclass);


--
-- TOC entry 3787 (class 2604 OID 47237)
-- Name: IsuAttachment id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuAttachment" ALTER COLUMN id SET DEFAULT nextval('public."IsuAttachment_id_seq"'::regclass);


--
-- TOC entry 3788 (class 2604 OID 47238)
-- Name: IsuComments id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuComments" ALTER COLUMN id SET DEFAULT nextval('public."IsuComments_id_seq"'::regclass);


--
-- TOC entry 3789 (class 2604 OID 47239)
-- Name: IsuCustomFieldSolution id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuCustomFieldSolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuCustomFieldSolution_id_seq"'::regclass);


--
-- TOC entry 3790 (class 2604 OID 47240)
-- Name: IsuDocument id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuDocument" ALTER COLUMN id SET DEFAULT nextval('public."IsuDocument_id_seq"'::regclass);


--
-- TOC entry 3791 (class 2604 OID 47241)
-- Name: IsuField id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuField" ALTER COLUMN id SET DEFAULT nextval('public."IsuField_id_seq"'::regclass);


--
-- TOC entry 3792 (class 2604 OID 47242)
-- Name: IsuFieldCodeToTableDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuFieldCodeToTableDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuFieldCodeToTableDefine_id_seq"'::regclass);


--
-- TOC entry 3793 (class 2604 OID 47243)
-- Name: IsuFieldDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuFieldDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuFieldDefine_id_seq"'::regclass);


--
-- TOC entry 3794 (class 2604 OID 47244)
-- Name: IsuHistory id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuHistory" ALTER COLUMN id SET DEFAULT nextval('public."IsuHistory_id_seq"'::regclass);


--
-- TOC entry 3795 (class 2604 OID 47245)
-- Name: IsuIssue id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue" ALTER COLUMN id SET DEFAULT nextval('public."IsuIssue_id_seq"'::regclass);


--
-- TOC entry 3796 (class 2604 OID 47246)
-- Name: IsuIssueExt pid; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssueExt" ALTER COLUMN pid SET DEFAULT nextval('public."IsuIssueExt_pid_seq"'::regclass);


--
-- TOC entry 3797 (class 2604 OID 47247)
-- Name: IsuLink id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuLink" ALTER COLUMN id SET DEFAULT nextval('public."IsuLink_id_seq"'::regclass);


--
-- TOC entry 3798 (class 2604 OID 47248)
-- Name: IsuLinkReasonDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuLinkReasonDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuLinkReasonDefine_id_seq"'::regclass);


--
-- TOC entry 3799 (class 2604 OID 47249)
-- Name: IsuNotification id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuNotification" ALTER COLUMN id SET DEFAULT nextval('public."IsuNotification_id_seq"'::regclass);


--
-- TOC entry 3800 (class 2604 OID 47250)
-- Name: IsuNotificationDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuNotificationDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuNotificationDefine_id_seq"'::regclass);


--
-- TOC entry 3801 (class 2604 OID 47251)
-- Name: IsuPage id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPage" ALTER COLUMN id SET DEFAULT nextval('public."IsuPage_id_seq"'::regclass);


--
-- TOC entry 3802 (class 2604 OID 47252)
-- Name: IsuPageElement id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageElement" ALTER COLUMN id SET DEFAULT nextval('public."IsuPageElement_id_seq"'::regclass);


--
-- TOC entry 3803 (class 2604 OID 47253)
-- Name: IsuPageSolution id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageSolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuPageSolution_id_seq"'::regclass);


--
-- TOC entry 3804 (class 2604 OID 47254)
-- Name: IsuPageSolutionItem id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageSolutionItem" ALTER COLUMN id SET DEFAULT nextval('public."IsuPageSolutionItem_id_seq"'::regclass);


--
-- TOC entry 3805 (class 2604 OID 47255)
-- Name: IsuPriority id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPriority" ALTER COLUMN id SET DEFAULT nextval('public."IsuPriority_id_seq"'::regclass);


--
-- TOC entry 3806 (class 2604 OID 47256)
-- Name: IsuPriorityDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPriorityDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuPriorityDefine_id_seq"'::regclass);


--
-- TOC entry 3807 (class 2604 OID 47257)
-- Name: IsuPrioritySolution id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPrioritySolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuPrioritySolution_id_seq"'::regclass);


--
-- TOC entry 3808 (class 2604 OID 47258)
-- Name: IsuQuery id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuQuery" ALTER COLUMN id SET DEFAULT nextval('public."IsuQuery_id_seq"'::regclass);


--
-- TOC entry 3809 (class 2604 OID 47259)
-- Name: IsuResolution id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuResolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuResolution_id_seq"'::regclass);


--
-- TOC entry 3810 (class 2604 OID 47260)
-- Name: IsuResolutionDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuResolutionDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuResolutionDefine_id_seq"'::regclass);


--
-- TOC entry 3811 (class 2604 OID 47261)
-- Name: IsuSeverity id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeverity" ALTER COLUMN id SET DEFAULT nextval('public."IsuSeverity_id_seq"'::regclass);


--
-- TOC entry 3812 (class 2604 OID 47262)
-- Name: IsuSeverityDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeverityDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuSeverityDefine_id_seq"'::regclass);


--
-- TOC entry 3813 (class 2604 OID 47263)
-- Name: IsuSeveritySolution id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeveritySolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuSeveritySolution_id_seq"'::regclass);


--
-- TOC entry 3814 (class 2604 OID 47264)
-- Name: IsuStatus id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuStatus" ALTER COLUMN id SET DEFAULT nextval('public."IsuStatus_id_seq"'::regclass);


--
-- TOC entry 3815 (class 2604 OID 47265)
-- Name: IsuStatusCategoryDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuStatusCategoryDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuStatusCategoryDefine_id_seq"'::regclass);


--
-- TOC entry 3816 (class 2604 OID 47266)
-- Name: IsuStatusDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuStatusDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuStatusDefine_id_seq"'::regclass);


--
-- TOC entry 3817 (class 2604 OID 47267)
-- Name: IsuTag id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTag" ALTER COLUMN id SET DEFAULT nextval('public."IsuTag_id_seq"'::regclass);


--
-- TOC entry 3818 (class 2604 OID 47268)
-- Name: IsuTagRelation id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTagRelation" ALTER COLUMN id SET DEFAULT nextval('public."IsuTagRelation_id_seq"'::regclass);


--
-- TOC entry 3819 (class 2604 OID 47269)
-- Name: IsuType id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuType" ALTER COLUMN id SET DEFAULT nextval('public."IsuType_id_seq"'::regclass);


--
-- TOC entry 3820 (class 2604 OID 47270)
-- Name: IsuTypeDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTypeDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuTypeDefine_id_seq"'::regclass);


--
-- TOC entry 3821 (class 2604 OID 47271)
-- Name: IsuTypeSolution id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTypeSolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuTypeSolution_id_seq"'::regclass);


--
-- TOC entry 3822 (class 2604 OID 47272)
-- Name: IsuWatch id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWatch" ALTER COLUMN id SET DEFAULT nextval('public."IsuWatch_id_seq"'::regclass);


--
-- TOC entry 3823 (class 2604 OID 47273)
-- Name: IsuWorkflow id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflow" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflow_id_seq"'::regclass);


--
-- TOC entry 3824 (class 2604 OID 47274)
-- Name: IsuWorkflowSolution id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowSolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowSolution_id_seq"'::regclass);


--
-- TOC entry 3825 (class 2604 OID 47275)
-- Name: IsuWorkflowSolutionItem id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowSolutionItem_id_seq"'::regclass);


--
-- TOC entry 3826 (class 2604 OID 47276)
-- Name: IsuWorkflowStatusRelation id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowStatusRelation_id_seq"'::regclass);


--
-- TOC entry 3827 (class 2604 OID 47277)
-- Name: IsuWorkflowStatusRelationDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelationDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowStatusRelationDefine_id_seq"'::regclass);


--
-- TOC entry 3828 (class 2604 OID 47278)
-- Name: IsuWorkflowTransition id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransition" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowTransition_id_seq"'::regclass);


--
-- TOC entry 3829 (class 2604 OID 47279)
-- Name: IsuWorkflowTransitionDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowTransitionDefine_id_seq"'::regclass);


--
-- TOC entry 3830 (class 2604 OID 47280)
-- Name: IsuWorkflowTransitionProjectRoleRelation id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowTransitionProjectRoleRelation_id_seq"'::regclass);


--
-- TOC entry 3831 (class 2604 OID 47281)
-- Name: SysPrivilege id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysPrivilege" ALTER COLUMN id SET DEFAULT nextval('public."SysPrivilege_id_seq"'::regclass);


--
-- TOC entry 3832 (class 2604 OID 47282)
-- Name: SysRole id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysRole" ALTER COLUMN id SET DEFAULT nextval('public."SysRole_id_seq"'::regclass);


--
-- TOC entry 3833 (class 2604 OID 47283)
-- Name: SysUser id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysUser" ALTER COLUMN id SET DEFAULT nextval('public."SysUser_id_seq"'::regclass);


--
-- TOC entry 3834 (class 2604 OID 47284)
-- Name: Test id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."Test" ALTER COLUMN id SET DEFAULT nextval('public."Test_id_seq"'::regclass);


--
-- TOC entry 3835 (class 2604 OID 47285)
-- Name: TstAlert id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstAlert" ALTER COLUMN id SET DEFAULT nextval('public."TstAlert_id_seq"'::regclass);


--
-- TOC entry 3836 (class 2604 OID 47286)
-- Name: TstCase id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCase" ALTER COLUMN id SET DEFAULT nextval('public."TstCase_id_seq"'::regclass);


--
-- TOC entry 3837 (class 2604 OID 47287)
-- Name: TstCaseAttachment id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseAttachment" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseAttachment_id_seq"'::regclass);


--
-- TOC entry 3838 (class 2604 OID 47288)
-- Name: TstCaseComments id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseComments" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseComments_id_seq"'::regclass);


--
-- TOC entry 3839 (class 2604 OID 47289)
-- Name: TstCaseExeStatusDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseExeStatusDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseExeStatus_id_seq"'::regclass);


--
-- TOC entry 3841 (class 2604 OID 47290)
-- Name: TstCaseHistory id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseHistory" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseHistory_id_seq"'::regclass);


--
-- TOC entry 3842 (class 2604 OID 47291)
-- Name: TstCaseInSuite id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInSuite" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInSuite_id_seq"'::regclass);


--
-- TOC entry 3843 (class 2604 OID 47292)
-- Name: TstCaseInTask id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTask" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTask_id_seq"'::regclass);


--
-- TOC entry 3844 (class 2604 OID 47293)
-- Name: TstCaseInTaskAttachment id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskAttachment" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTaskAttachment_id_seq"'::regclass);


--
-- TOC entry 3845 (class 2604 OID 47294)
-- Name: TstCaseInTaskComments id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskComments" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTaskComments_id_seq"'::regclass);


--
-- TOC entry 3846 (class 2604 OID 47295)
-- Name: TstCaseInTaskHistory id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskHistory" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTaskHistory_id_seq"'::regclass);


--
-- TOC entry 3847 (class 2604 OID 47296)
-- Name: TstCaseInTaskIssue id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskIssue" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTaskIssue_id_seq"'::regclass);


--
-- TOC entry 3848 (class 2604 OID 47297)
-- Name: TstCasePriorityDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCasePriorityDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstCasePriority_id_seq"'::regclass);


--
-- TOC entry 3850 (class 2604 OID 47298)
-- Name: TstCaseStep id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseStep" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseStep_id_seq"'::regclass);


--
-- TOC entry 3851 (class 2604 OID 47299)
-- Name: TstCaseTypeDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseTypeDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseType_id_seq"'::regclass);


--
-- TOC entry 3853 (class 2604 OID 47300)
-- Name: TstDocument id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstDocument" ALTER COLUMN id SET DEFAULT nextval('public."TstDocument_id_seq"'::regclass);


--
-- TOC entry 3854 (class 2604 OID 47301)
-- Name: TstEmail id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstEmail" ALTER COLUMN id SET DEFAULT nextval('public."TstEmail_id_seq"'::regclass);


--
-- TOC entry 3855 (class 2604 OID 47302)
-- Name: TstEnv id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstEnv" ALTER COLUMN id SET DEFAULT nextval('public."TstEnv_id_seq"'::regclass);


--
-- TOC entry 3856 (class 2604 OID 47303)
-- Name: TstHistory id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstHistory" ALTER COLUMN id SET DEFAULT nextval('public."TstHistory_id_seq"'::regclass);


--
-- TOC entry 3857 (class 2604 OID 47304)
-- Name: TstModule id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstModule" ALTER COLUMN id SET DEFAULT nextval('public."TstModule_id_seq"'::regclass);


--
-- TOC entry 3858 (class 2604 OID 47305)
-- Name: TstMsg id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstMsg" ALTER COLUMN id SET DEFAULT nextval('public."TstMsg_id_seq"'::regclass);


--
-- TOC entry 3859 (class 2604 OID 47306)
-- Name: TstOrg id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrg" ALTER COLUMN id SET DEFAULT nextval('public."TstOrg_id_seq"'::regclass);


--
-- TOC entry 3860 (class 2604 OID 47307)
-- Name: TstOrgGroup id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgGroup" ALTER COLUMN id SET DEFAULT nextval('public."TstOrgGroup_id_seq"'::regclass);


--
-- TOC entry 3861 (class 2604 OID 47308)
-- Name: TstOrgPrivilegeDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgPrivilegeDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstOrgPrivilegeDefine_id_seq"'::regclass);


--
-- TOC entry 3862 (class 2604 OID 47309)
-- Name: TstOrgRole id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRole" ALTER COLUMN id SET DEFAULT nextval('public."TstOrgRole_id_seq"'::regclass);


--
-- TOC entry 3863 (class 2604 OID 47310)
-- Name: TstPlan id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstPlan" ALTER COLUMN id SET DEFAULT nextval('public."TstPlan_id_seq"'::regclass);


--
-- TOC entry 3864 (class 2604 OID 47311)
-- Name: TstProject id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProject" ALTER COLUMN id SET DEFAULT nextval('public."TstProject_id_seq"'::regclass);


--
-- TOC entry 3865 (class 2604 OID 47312)
-- Name: TstProjectAccessHistory id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectAccessHistory" ALTER COLUMN id SET DEFAULT nextval('public."TstProjectAccessHistory_id_seq"'::regclass);


--
-- TOC entry 3866 (class 2604 OID 47313)
-- Name: TstProjectPrivilegeDefine id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectPrivilegeDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstProjectPrivilegeDefine_id_seq"'::regclass);


--
-- TOC entry 3867 (class 2604 OID 47314)
-- Name: TstProjectRole id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectRole" ALTER COLUMN id SET DEFAULT nextval('public."TstProjectRole_id_seq"'::regclass);


--
-- TOC entry 3868 (class 2604 OID 47315)
-- Name: TstSuite id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstSuite" ALTER COLUMN id SET DEFAULT nextval('public."TstSuite_id_seq"'::regclass);


--
-- TOC entry 3869 (class 2604 OID 47316)
-- Name: TstTask id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstTask" ALTER COLUMN id SET DEFAULT nextval('public."TstTask_id_seq"'::regclass);


--
-- TOC entry 3870 (class 2604 OID 47317)
-- Name: TstThread id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstThread" ALTER COLUMN id SET DEFAULT nextval('public."TstThread_id_seq"'::regclass);


--
-- TOC entry 3871 (class 2604 OID 47318)
-- Name: TstUser id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstUser" ALTER COLUMN id SET DEFAULT nextval('public."TstUser_id_seq"'::regclass);


--
-- TOC entry 3872 (class 2604 OID 47319)
-- Name: TstUserVerifyCode id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstUserVerifyCode" ALTER COLUMN id SET DEFAULT nextval('public."TstUserVerifyCode_id_seq"'::regclass);


--
-- TOC entry 3873 (class 2604 OID 47320)
-- Name: TstVer id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstVer" ALTER COLUMN id SET DEFAULT nextval('public."TstVer_id_seq"'::regclass);


--
-- TOC entry 4558 (class 0 OID 46475)
-- Dependencies: 197
-- Data for Name: CustomField; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."CustomField" (id, "colCode", label, type, input, "textFormat", "applyTo", rows, required, readonly, "fullLine", ordr, descr, "buildIn", "orgId", "createTime", "updateTime", disabled, deleted) FROM stdin;
1	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	1	2019-02-17 14:22:02.073523	\N	f	f
2	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	2	2019-02-18 21:51:44.67068	\N	f	f
3	prop001	字符串字段01	string	text	plain_text	issue	3	f	f	f	11	\N	f	2	2019-02-19 09:22:02.333544	\N	f	f
5	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	4	2019-02-19 11:38:28.481573	\N	f	f
6	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	5	2019-02-22 08:33:31.929668	\N	f	f
7	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	6	2019-02-23 13:16:23.128256	\N	f	f
8	prop001	文本字段01	string	text	plain_text	test_case	3	f	f	f	20	地方	f	5	2019-02-23 16:59:37.326243	2019-02-23 17:01:42.096496	f	f
9	prop002	是打发	string	text	plain_text	test_case	3	f	f	f	10	\N	f	5	2019-02-23 17:01:38.332385	2019-02-23 17:01:49.563576	f	t
10	prop002	下来菜单01	integer	dropdown	plain_text	test_case	3	f	f	f	30	\N	f	5	2019-02-23 17:23:00.621915	2019-02-23 17:23:16.986804	f	f
11	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	7	2019-03-11 00:56:22.895443	\N	f	f
12	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	8	2019-03-12 21:42:50.71871	\N	f	f
13	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	9	2019-03-12 21:44:18.566135	\N	f	f
14	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	10	2019-03-12 22:44:53.460728	\N	f	f
15	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	11	2019-03-12 22:52:14.395588	\N	f	f
16	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	12	2019-03-12 23:06:01.367716	\N	f	f
17	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	13	2019-03-12 23:11:45.58636	\N	f	f
18	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	\N	14	2019-03-12 23:27:21.303261	\N	f	f
\.


--
-- TOC entry 4559 (class 0 OID 46481)
-- Dependencies: 198
-- Data for Name: CustomFieldDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."CustomFieldDefine" (id, "colCode", label, type, input, "textFormat", "applyTo", rows, required, readonly, "fullLine", ordr, descr, "createTime", "updateTime", disabled, deleted) FROM stdin;
1	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	2018-11-09 12:06:02	\N	f	f
\.


--
-- TOC entry 4561 (class 0 OID 46489)
-- Dependencies: 200
-- Data for Name: CustomFieldInputTypeRelationDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."CustomFieldInputTypeRelationDefine" (id, "inputValue", "typeValue") FROM stdin;
1	text	string
2	number	integer
3	number	double
4	textarea	string
9	date	date
10	time	time
11	datetime	datetime
12	richtext	string
5	dropdown	integer
6	multi_select	integer
7	radio	integer
8	checkbox	integer
\.


--
-- TOC entry 4563 (class 0 OID 46497)
-- Dependencies: 202
-- Data for Name: CustomFieldIputDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."CustomFieldIputDefine" (id, label, value, ordr, disabled, deleted, "createTime", "updateTime") FROM stdin;
1	文本	text	1	f	f	2018-11-28 09:24:07	\N
2	数字	number	2	f	f	2018-11-28 09:24:07	\N
3	多行文本	textarea	3	f	f	2018-11-28 09:24:07	\N
4	下拉菜单	dropdown	4	f	f	2018-11-28 09:24:07	\N
5	下拉菜单(多选)	multi_select	5	f	f	2018-11-28 09:24:07	\N
6	单选按钮	radio	6	f	f	2018-11-28 09:24:07	\N
7	多选框	checkbox	7	f	f	2018-11-28 09:24:07	\N
8	日期	date	8	f	f	2018-11-28 09:24:07	\N
9	时间	time	9	f	f	2018-11-28 09:24:07	\N
10	日期时间	datetime	10	f	f	2018-12-07 13:55:03	\N
11	富文本	richtext	11	f	f	2018-12-26 12:39:44	\N
\.


--
-- TOC entry 4565 (class 0 OID 46505)
-- Dependencies: 204
-- Data for Name: CustomFieldOption; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."CustomFieldOption" (id, label, descr, ordr, "defaultVal", "buildIn", "fieldId", "orgId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	细微	\N	5	f	t	1	1	f	f	2019-02-17 14:22:02.073523	\N
2	一般	\N	4	t	t	1	1	f	f	2019-02-17 14:22:02.073523	\N
3	重要	\N	3	f	t	1	1	f	f	2019-02-17 14:22:02.073523	\N
4	紧急	\N	2	f	t	1	1	f	f	2019-02-17 14:22:02.073523	\N
5	阻塞	\N	1	f	t	1	1	f	f	2019-02-17 14:22:02.073523	\N
6	细微	\N	5	f	t	2	2	f	f	2019-02-18 21:51:44.67068	\N
7	一般	\N	4	t	t	2	2	f	f	2019-02-18 21:51:44.67068	\N
8	重要	\N	3	f	t	2	2	f	f	2019-02-18 21:51:44.67068	\N
9	紧急	\N	2	f	t	2	2	f	f	2019-02-18 21:51:44.67068	\N
10	阻塞	\N	1	f	t	2	2	f	f	2019-02-18 21:51:44.67068	\N
16	细微	\N	5	f	t	5	4	f	f	2019-02-19 11:38:28.481573	\N
17	一般	\N	4	t	t	5	4	f	f	2019-02-19 11:38:28.481573	\N
18	重要	\N	3	f	t	5	4	f	f	2019-02-19 11:38:28.481573	\N
19	紧急	\N	2	f	t	5	4	f	f	2019-02-19 11:38:28.481573	\N
20	阻塞	\N	1	f	t	5	4	f	f	2019-02-19 11:38:28.481573	\N
21	细微	\N	5	f	t	6	5	f	f	2019-02-22 08:33:31.929668	\N
22	一般	\N	4	t	t	6	5	f	f	2019-02-22 08:33:31.929668	\N
23	重要	\N	3	f	t	6	5	f	f	2019-02-22 08:33:31.929668	\N
24	紧急	\N	2	f	t	6	5	f	f	2019-02-22 08:33:31.929668	\N
25	阻塞	\N	1	f	t	6	5	f	f	2019-02-22 08:33:31.929668	\N
26	细微	\N	5	f	t	7	6	f	f	2019-02-23 13:16:23.128256	\N
27	一般	\N	4	t	t	7	6	f	f	2019-02-23 13:16:23.128256	\N
28	重要	\N	3	f	t	7	6	f	f	2019-02-23 13:16:23.128256	\N
29	紧急	\N	2	f	t	7	6	f	f	2019-02-23 13:16:23.128256	\N
30	阻塞	\N	1	f	t	7	6	f	f	2019-02-23 13:16:23.128256	\N
31	1	\N	10	f	f	10	5	f	f	2019-02-23 17:23:00.745879	2019-02-23 17:23:05.414625
32	2	\N	20	t	f	10	5	f	f	2019-02-23 17:23:02.151868	2019-02-23 17:23:05.414625
33	3	\N	30	f	f	10	5	f	f	2019-02-23 17:23:08.95248	\N
34	4	\N	40	f	f	10	5	f	t	2019-02-23 17:23:10.313641	\N
35	细微	\N	5	f	t	11	7	f	f	2019-03-11 00:56:22.895443	\N
36	一般	\N	4	t	t	11	7	f	f	2019-03-11 00:56:22.895443	\N
37	重要	\N	3	f	t	11	7	f	f	2019-03-11 00:56:22.895443	\N
38	紧急	\N	2	f	t	11	7	f	f	2019-03-11 00:56:22.895443	\N
39	阻塞	\N	1	f	t	11	7	f	f	2019-03-11 00:56:22.895443	\N
40	细微	\N	5	f	t	12	8	f	f	2019-03-12 21:42:50.71871	\N
41	一般	\N	4	t	t	12	8	f	f	2019-03-12 21:42:50.71871	\N
42	重要	\N	3	f	t	12	8	f	f	2019-03-12 21:42:50.71871	\N
43	紧急	\N	2	f	t	12	8	f	f	2019-03-12 21:42:50.71871	\N
44	阻塞	\N	1	f	t	12	8	f	f	2019-03-12 21:42:50.71871	\N
45	细微	\N	5	f	t	13	9	f	f	2019-03-12 21:44:18.566135	\N
46	一般	\N	4	t	t	13	9	f	f	2019-03-12 21:44:18.566135	\N
47	重要	\N	3	f	t	13	9	f	f	2019-03-12 21:44:18.566135	\N
48	紧急	\N	2	f	t	13	9	f	f	2019-03-12 21:44:18.566135	\N
49	阻塞	\N	1	f	t	13	9	f	f	2019-03-12 21:44:18.566135	\N
50	细微	\N	5	f	t	14	10	f	f	2019-03-12 22:44:53.460728	\N
51	一般	\N	4	t	t	14	10	f	f	2019-03-12 22:44:53.460728	\N
52	重要	\N	3	f	t	14	10	f	f	2019-03-12 22:44:53.460728	\N
53	紧急	\N	2	f	t	14	10	f	f	2019-03-12 22:44:53.460728	\N
54	阻塞	\N	1	f	t	14	10	f	f	2019-03-12 22:44:53.460728	\N
55	细微	\N	5	f	t	15	11	f	f	2019-03-12 22:52:14.395588	\N
56	一般	\N	4	t	t	15	11	f	f	2019-03-12 22:52:14.395588	\N
57	重要	\N	3	f	t	15	11	f	f	2019-03-12 22:52:14.395588	\N
58	紧急	\N	2	f	t	15	11	f	f	2019-03-12 22:52:14.395588	\N
59	阻塞	\N	1	f	t	15	11	f	f	2019-03-12 22:52:14.395588	\N
60	细微	\N	5	f	t	16	12	f	f	2019-03-12 23:06:01.367716	\N
61	一般	\N	4	t	t	16	12	f	f	2019-03-12 23:06:01.367716	\N
62	重要	\N	3	f	t	16	12	f	f	2019-03-12 23:06:01.367716	\N
63	紧急	\N	2	f	t	16	12	f	f	2019-03-12 23:06:01.367716	\N
64	阻塞	\N	1	f	t	16	12	f	f	2019-03-12 23:06:01.367716	\N
65	细微	\N	5	f	t	17	13	f	f	2019-03-12 23:11:45.58636	\N
66	一般	\N	4	t	t	17	13	f	f	2019-03-12 23:11:45.58636	\N
67	重要	\N	3	f	t	17	13	f	f	2019-03-12 23:11:45.58636	\N
68	紧急	\N	2	f	t	17	13	f	f	2019-03-12 23:11:45.58636	\N
69	阻塞	\N	1	f	t	17	13	f	f	2019-03-12 23:11:45.58636	\N
70	细微	\N	5	f	t	18	14	f	f	2019-03-12 23:27:21.303261	\N
71	一般	\N	4	t	t	18	14	f	f	2019-03-12 23:27:21.303261	\N
72	重要	\N	3	f	t	18	14	f	f	2019-03-12 23:27:21.303261	\N
73	紧急	\N	2	f	t	18	14	f	f	2019-03-12 23:27:21.303261	\N
74	阻塞	\N	1	f	t	18	14	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4566 (class 0 OID 46511)
-- Dependencies: 205
-- Data for Name: CustomFieldOptionDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."CustomFieldOptionDefine" (id, label, descr, ordr, "defaultVal", "fieldId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	阻塞	\N	1	f	1	f	f	2018-11-09 12:49:25	\N
2	紧急	\N	2	f	1	f	f	2018-11-09 12:49:28	\N
3	重要	\N	3	f	1	f	f	2018-11-09 12:49:31	\N
4	一般	\N	4	t	1	f	f	2018-11-09 12:49:33	\N
5	细微	\N	5	f	1	f	f	2018-11-09 12:49:36	\N
\.


--
-- TOC entry 4569 (class 0 OID 46521)
-- Dependencies: 208
-- Data for Name: CustomFieldTypeDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."CustomFieldTypeDefine" (id, label, value, ordr, disabled, deleted, "createTime", "updateTime") FROM stdin;
10	字符串	string	1	f	f	2018-11-28 09:24:07	\N
20	整数	integer	2	f	f	2018-11-28 09:24:07	\N
30	浮点数	double	3	f	f	2018-11-28 09:49:06	\N
40	日期	date	4	f	f	2018-11-28 09:24:07	\N
50	时间	time	5	f	f	2018-11-28 09:24:07	\N
60	日期时间	datetime	6	f	f	2018-12-07 13:57:41	\N
\.


--
-- TOC entry 4572 (class 0 OID 46531)
-- Dependencies: 211
-- Data for Name: IsuAttachment; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuAttachment" (id, name, title, uri, descr, "docType", "issueId", "userId", deleted, disabled, "createTime", "updateTime") FROM stdin;
2	樊登读书.txt	\N	upload/data/20190225/8e6af499-746f-4aeb-b730-b087130134b3.txt	\N	\N	5	5	f	f	2019-02-25 16:22:16.485	\N
1	樊登读书.txt	\N	upload/data/20190225/ee2ee6d6-5645-4d7a-9a84-2ab55196b936.txt	\N	\N	5	5	t	f	2019-02-25 16:20:44.99	\N
3	Dockerfile	\N	upload/data/20190225/99db6436-027b-4c43-b7ee-f056bc7a7d0d.	\N	\N	5	5	f	f	2019-02-25 16:22:53.055	\N
\.


--
-- TOC entry 4574 (class 0 OID 46539)
-- Dependencies: 213
-- Data for Name: IsuComments; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuComments" (id, summary, content, "issueId", "userId", "createTime", "updateTime", disabled, deleted) FROM stdin;
1	修改经办人	\N	8	5	2019-02-25 16:02:26.515123	2019-02-25 16:02:26.515123	f	f
2	修改经办人	\N	8	5	2019-02-25 16:03:44.447129	2019-02-25 16:03:44.447129	f	f
3	修改经办人	尽快搞定	8	5	2019-02-25 16:03:58.248304	2019-02-25 16:03:58.248304	f	f
5	添加备注	是的范德萨发	5	5	2019-02-25 16:23:13.089547	2019-02-25 16:23:13.089547	f	f
4	修改备注	你好2	5	5	2019-02-25 16:23:01.108938	2019-02-25 16:23:06.298133	t	f
\.


--
-- TOC entry 4576 (class 0 OID 46547)
-- Dependencies: 215
-- Data for Name: IsuCustomFieldSolution; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuCustomFieldSolution" (id, name, descr, "orgId", "createTime", "updateTime", disabled, deleted) FROM stdin;
\.


--
-- TOC entry 4577 (class 0 OID 46553)
-- Dependencies: 216
-- Data for Name: IsuCustomFieldSolutionFieldRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuCustomFieldSolutionFieldRelation" ("solutionId", "fieldId") FROM stdin;
\.


--
-- TOC entry 4578 (class 0 OID 46556)
-- Dependencies: 217
-- Data for Name: IsuCustomFieldSolutionProjectRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuCustomFieldSolutionProjectRelation" ("solutionId", "orgId", "projectId") FROM stdin;
\.


--
-- TOC entry 4580 (class 0 OID 46561)
-- Dependencies: 219
-- Data for Name: IsuDocument; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuDocument" (id, "createTime", deleted, disabled, "updateTime", version, descr, "docType", "issueId", title, uri, "userId") FROM stdin;
\.


--
-- TOC entry 4582 (class 0 OID 46569)
-- Dependencies: 221
-- Data for Name: IsuField; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuField" (id, "colCode", label, type, input, "fullLine", required, "defaultShowInFilters", "filterOrdr", "defaultShowInColumns", "columnOrdr", "defaultShowInPage", "elemOrdr", readonly, "orgId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	1	f	f	2019-02-17 14:22:02.073523	\N
2	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	1	f	f	2019-02-17 14:22:02.073523	\N
3	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	1	f	f	2019-02-17 14:22:02.073523	\N
4	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	1	f	f	2019-02-17 14:22:02.073523	\N
5	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	1	f	f	2019-02-17 14:22:02.073523	\N
6	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	1	f	f	2019-02-17 14:22:02.073523	\N
7	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	1	f	f	2019-02-17 14:22:02.073523	\N
8	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	1	f	f	2019-02-17 14:22:02.073523	\N
9	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	1	f	f	2019-02-17 14:22:02.073523	\N
10	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	1	f	f	2019-02-17 14:22:02.073523	\N
11	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	1	f	f	2019-02-17 14:22:02.073523	\N
12	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	1	f	f	2019-02-17 14:22:02.073523	\N
13	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	1	f	f	2019-02-17 14:22:02.073523	\N
15	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	2	f	f	2019-02-18 21:51:44.67068	\N
16	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	2	f	f	2019-02-18 21:51:44.67068	\N
17	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	2	f	f	2019-02-18 21:51:44.67068	\N
18	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	2	f	f	2019-02-18 21:51:44.67068	\N
19	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	2	f	f	2019-02-18 21:51:44.67068	\N
20	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	2	f	f	2019-02-18 21:51:44.67068	\N
21	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	2	f	f	2019-02-18 21:51:44.67068	\N
22	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	2	f	f	2019-02-18 21:51:44.67068	\N
23	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	2	f	f	2019-02-18 21:51:44.67068	\N
24	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	2	f	f	2019-02-18 21:51:44.67068	\N
25	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	2	f	f	2019-02-18 21:51:44.67068	\N
26	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	2	f	f	2019-02-18 21:51:44.67068	\N
27	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	2	f	f	2019-02-18 21:51:44.67068	\N
28	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	2	f	f	2019-02-18 21:51:44.67068	\N
43	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	4	f	f	2019-02-19 11:38:28.481573	\N
44	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	4	f	f	2019-02-19 11:38:28.481573	\N
45	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	4	f	f	2019-02-19 11:38:28.481573	\N
46	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	4	f	f	2019-02-19 11:38:28.481573	\N
47	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	4	f	f	2019-02-19 11:38:28.481573	\N
48	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	4	f	f	2019-02-19 11:38:28.481573	\N
49	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	4	f	f	2019-02-19 11:38:28.481573	\N
50	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	4	f	f	2019-02-19 11:38:28.481573	\N
51	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	4	f	f	2019-02-19 11:38:28.481573	\N
52	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	4	f	f	2019-02-19 11:38:28.481573	\N
53	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	4	f	f	2019-02-19 11:38:28.481573	\N
54	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	4	f	f	2019-02-19 11:38:28.481573	\N
55	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	4	f	f	2019-02-19 11:38:28.481573	\N
56	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	4	f	f	2019-02-19 11:38:28.481573	\N
57	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	5	f	f	2019-02-22 08:33:31.929668	\N
58	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	5	f	f	2019-02-22 08:33:31.929668	\N
59	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	5	f	f	2019-02-22 08:33:31.929668	\N
60	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	5	f	f	2019-02-22 08:33:31.929668	\N
61	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	5	f	f	2019-02-22 08:33:31.929668	\N
62	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	5	f	f	2019-02-22 08:33:31.929668	\N
63	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	5	f	f	2019-02-22 08:33:31.929668	\N
64	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	5	f	f	2019-02-22 08:33:31.929668	\N
65	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	5	f	f	2019-02-22 08:33:31.929668	\N
66	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	5	f	f	2019-02-22 08:33:31.929668	\N
67	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	5	f	f	2019-02-22 08:33:31.929668	\N
68	resolutionDescr	解决详情	string	textarea	t	f	\N	\N	\N	\N	f	20000	f	5	f	f	2019-02-22 08:33:31.929668	\N
69	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	5	f	f	2019-02-22 08:33:31.929668	\N
70	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	5	f	f	2019-02-22 08:33:31.929668	\N
71	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	6	f	f	2019-02-23 13:16:23.128256	\N
72	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	6	f	f	2019-02-23 13:16:23.128256	\N
73	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	6	f	f	2019-02-23 13:16:23.128256	\N
74	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	6	f	f	2019-02-23 13:16:23.128256	\N
75	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	6	f	f	2019-02-23 13:16:23.128256	\N
76	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	6	f	f	2019-02-23 13:16:23.128256	\N
77	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	6	f	f	2019-02-23 13:16:23.128256	\N
78	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	6	f	f	2019-02-23 13:16:23.128256	\N
79	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	6	f	f	2019-02-23 13:16:23.128256	\N
80	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	6	f	f	2019-02-23 13:16:23.128256	\N
81	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	6	f	f	2019-02-23 13:16:23.128256	\N
83	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	6	f	f	2019-02-23 13:16:23.128256	\N
84	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	6	f	f	2019-02-23 13:16:23.128256	\N
82	resolutionDescr	解决详情	string	textarea	t	f	\N	\N	\N	\N	f	20000	f	6	f	f	2019-02-23 13:16:23.128256	\N
85	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	7	f	f	2019-03-11 00:56:22.895443	\N
86	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	7	f	f	2019-03-11 00:56:22.895443	\N
87	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	7	f	f	2019-03-11 00:56:22.895443	\N
88	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	7	f	f	2019-03-11 00:56:22.895443	\N
89	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	7	f	f	2019-03-11 00:56:22.895443	\N
90	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	7	f	f	2019-03-11 00:56:22.895443	\N
91	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	7	f	f	2019-03-11 00:56:22.895443	\N
92	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	7	f	f	2019-03-11 00:56:22.895443	\N
93	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	7	f	f	2019-03-11 00:56:22.895443	\N
94	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	7	f	f	2019-03-11 00:56:22.895443	\N
95	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	7	f	f	2019-03-11 00:56:22.895443	\N
96	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	7	f	f	2019-03-11 00:56:22.895443	\N
97	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	7	f	f	2019-03-11 00:56:22.895443	\N
98	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	7	f	f	2019-03-11 00:56:22.895443	\N
99	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	8	f	f	2019-03-12 21:42:50.71871	\N
100	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	8	f	f	2019-03-12 21:42:50.71871	\N
101	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	8	f	f	2019-03-12 21:42:50.71871	\N
102	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	8	f	f	2019-03-12 21:42:50.71871	\N
103	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	8	f	f	2019-03-12 21:42:50.71871	\N
104	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	8	f	f	2019-03-12 21:42:50.71871	\N
105	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	8	f	f	2019-03-12 21:42:50.71871	\N
106	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	8	f	f	2019-03-12 21:42:50.71871	\N
107	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	8	f	f	2019-03-12 21:42:50.71871	\N
108	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	8	f	f	2019-03-12 21:42:50.71871	\N
109	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	8	f	f	2019-03-12 21:42:50.71871	\N
110	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	8	f	f	2019-03-12 21:42:50.71871	\N
111	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	8	f	f	2019-03-12 21:42:50.71871	\N
112	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	8	f	f	2019-03-12 21:42:50.71871	\N
113	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	9	f	f	2019-03-12 21:44:18.566135	\N
114	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	9	f	f	2019-03-12 21:44:18.566135	\N
115	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	9	f	f	2019-03-12 21:44:18.566135	\N
116	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	9	f	f	2019-03-12 21:44:18.566135	\N
117	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	9	f	f	2019-03-12 21:44:18.566135	\N
118	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	9	f	f	2019-03-12 21:44:18.566135	\N
119	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	9	f	f	2019-03-12 21:44:18.566135	\N
120	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	9	f	f	2019-03-12 21:44:18.566135	\N
121	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	9	f	f	2019-03-12 21:44:18.566135	\N
122	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	9	f	f	2019-03-12 21:44:18.566135	\N
123	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	9	f	f	2019-03-12 21:44:18.566135	\N
124	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	9	f	f	2019-03-12 21:44:18.566135	\N
125	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	9	f	f	2019-03-12 21:44:18.566135	\N
126	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	9	f	f	2019-03-12 21:44:18.566135	\N
127	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	10	f	f	2019-03-12 22:44:53.460728	\N
128	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	10	f	f	2019-03-12 22:44:53.460728	\N
129	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	10	f	f	2019-03-12 22:44:53.460728	\N
130	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	10	f	f	2019-03-12 22:44:53.460728	\N
131	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	10	f	f	2019-03-12 22:44:53.460728	\N
132	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	10	f	f	2019-03-12 22:44:53.460728	\N
133	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	10	f	f	2019-03-12 22:44:53.460728	\N
134	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	10	f	f	2019-03-12 22:44:53.460728	\N
135	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	10	f	f	2019-03-12 22:44:53.460728	\N
136	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	10	f	f	2019-03-12 22:44:53.460728	\N
137	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	10	f	f	2019-03-12 22:44:53.460728	\N
138	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	10	f	f	2019-03-12 22:44:53.460728	\N
139	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	10	f	f	2019-03-12 22:44:53.460728	\N
140	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	10	f	f	2019-03-12 22:44:53.460728	\N
141	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	11	f	f	2019-03-12 22:52:14.395588	\N
142	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	11	f	f	2019-03-12 22:52:14.395588	\N
143	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	11	f	f	2019-03-12 22:52:14.395588	\N
144	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	11	f	f	2019-03-12 22:52:14.395588	\N
145	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	11	f	f	2019-03-12 22:52:14.395588	\N
146	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	11	f	f	2019-03-12 22:52:14.395588	\N
147	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	11	f	f	2019-03-12 22:52:14.395588	\N
148	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	11	f	f	2019-03-12 22:52:14.395588	\N
149	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	11	f	f	2019-03-12 22:52:14.395588	\N
150	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	11	f	f	2019-03-12 22:52:14.395588	\N
151	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	11	f	f	2019-03-12 22:52:14.395588	\N
152	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	11	f	f	2019-03-12 22:52:14.395588	\N
153	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	11	f	f	2019-03-12 22:52:14.395588	\N
154	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	11	f	f	2019-03-12 22:52:14.395588	\N
155	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	12	f	f	2019-03-12 23:06:01.367716	\N
156	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	12	f	f	2019-03-12 23:06:01.367716	\N
157	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	12	f	f	2019-03-12 23:06:01.367716	\N
158	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	12	f	f	2019-03-12 23:06:01.367716	\N
159	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	12	f	f	2019-03-12 23:06:01.367716	\N
160	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	12	f	f	2019-03-12 23:06:01.367716	\N
161	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	12	f	f	2019-03-12 23:06:01.367716	\N
162	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	12	f	f	2019-03-12 23:06:01.367716	\N
163	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	12	f	f	2019-03-12 23:06:01.367716	\N
164	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	12	f	f	2019-03-12 23:06:01.367716	\N
165	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	12	f	f	2019-03-12 23:06:01.367716	\N
166	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	12	f	f	2019-03-12 23:06:01.367716	\N
167	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	12	f	f	2019-03-12 23:06:01.367716	\N
168	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	12	f	f	2019-03-12 23:06:01.367716	\N
169	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	13	f	f	2019-03-12 23:11:45.58636	\N
170	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	13	f	f	2019-03-12 23:11:45.58636	\N
171	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	13	f	f	2019-03-12 23:11:45.58636	\N
172	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	13	f	f	2019-03-12 23:11:45.58636	\N
173	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	13	f	f	2019-03-12 23:11:45.58636	\N
174	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	13	f	f	2019-03-12 23:11:45.58636	\N
175	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	13	f	f	2019-03-12 23:11:45.58636	\N
176	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	13	f	f	2019-03-12 23:11:45.58636	\N
177	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	13	f	f	2019-03-12 23:11:45.58636	\N
178	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	13	f	f	2019-03-12 23:11:45.58636	\N
179	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	13	f	f	2019-03-12 23:11:45.58636	\N
180	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	13	f	f	2019-03-12 23:11:45.58636	\N
181	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	13	f	f	2019-03-12 23:11:45.58636	\N
182	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	13	f	f	2019-03-12 23:11:45.58636	\N
183	typeId	类型	integer	dropdown	f	f	t	10200	t	10200	t	10200	f	14	f	f	2019-03-12 23:27:21.303261	\N
184	statusId	状态	integer	dropdown	f	f	t	10300	t	10300	t	10150	f	14	f	f	2019-03-12 23:27:21.303261	\N
185	priorityId	优先级	integer	dropdown	f	f	t	10400	t	10400	t	10400	f	14	f	f	2019-03-12 23:27:21.303261	\N
186	assigneeId	经办人	integer	dropdown	f	f	t	10500	t	10500	t	10500	f	14	f	f	2019-03-12 23:27:21.303261	\N
187	creatorId	创建人	integer	dropdown	f	f	f	10600	f	10600	f	11200	t	14	f	f	2019-03-12 23:27:21.303261	\N
188	reporterId	报告人	integer	dropdown	f	f	f	10700	f	10700	t	10550	f	14	f	f	2019-03-12 23:27:21.303261	\N
189	verId	版本	integer	dropdown	f	f	f	10800	f	10800	t	10600	f	14	f	f	2019-03-12 23:27:21.303261	\N
190	envId	环境	integer	dropdown	f	f	f	10900	f	10900	t	10700	f	14	f	f	2019-03-12 23:27:21.303261	\N
191	resolutionId	解决结果	integer	dropdown	f	f	f	11000	f	11000	f	11000	f	14	f	f	2019-03-12 23:27:21.303261	\N
192	dueTime	截止时间	date	date	f	f	f	11100	f	11100	f	10900	f	14	f	f	2019-03-12 23:27:21.303261	\N
193	resolveTime	解决时间	date	date	f	f	f	11200	f	11200	f	11100	f	14	f	f	2019-03-12 23:27:21.303261	\N
194	resolutionDescr	解决详情	string	textarea	f	f	\N	\N	\N	\N	f	20000	f	14	f	f	2019-03-12 23:27:21.303261	\N
195	title	标题	string	text	t	t	\N	\N	t	10100	t	10100	f	14	f	f	2019-03-12 23:27:21.303261	\N
196	descr	描述	string	textarea	t	f	f	11250	\N	\N	t	10800	f	14	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4583 (class 0 OID 46575)
-- Dependencies: 222
-- Data for Name: IsuFieldCodeToTableDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuFieldCodeToTableDefine" (id, "colCode", "table", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	typeId	IsuType	f	f	2018-11-23 11:44:15	\N
2	statusId	IsuStatus	f	f	2018-11-23 11:44:15	\N
3	priorityId	IsuPriority	f	f	2018-11-23 11:44:15	\N
4	verId	TstVer	f	f	2018-11-23 11:44:15	\N
5	envId	TstEnv	f	f	2018-11-23 11:44:15	\N
6	resolutionId	IsuResolution	f	f	2018-11-23 11:44:15	\N
7	assigneeId	TstUser	f	f	2018-11-23 11:44:15	\N
8	creatorId	TstUser	f	f	2018-11-23 11:44:15	\N
9	reporterId	TstUser	f	f	2018-11-23 11:44:15	\N
10	projectId	TstProject	f	f	2018-11-23 11:55:40	\N
\.


--
-- TOC entry 4585 (class 0 OID 46583)
-- Dependencies: 224
-- Data for Name: IsuFieldDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuFieldDefine" (id, "colCode", label, type, input, "defaultShowInFilters", "filterOrdr", "defaultShowInColumns", "columnOrdr", "defaultShowInPage", "elemOrdr", readonly, "fullLine", required, disabled, deleted, "createTime", "updateTime") FROM stdin;
2	projectId	项目	integer	dropdown	\N	\N	f	11300	\N	\N	f	\N	\N	f	f	2018-11-09 13:18:24	\N
3	typeId	类型	integer	dropdown	t	10200	t	10200	t	10200	f	f	f	f	f	2018-11-09 13:18:24	\N
4	statusId	状态	integer	dropdown	t	10300	t	10300	t	10150	f	f	f	f	f	2018-11-09 13:18:24	\N
5	priorityId	优先级	integer	dropdown	t	10400	t	10400	t	10400	f	f	f	f	f	2018-11-09 13:18:24	\N
6	assigneeId	经办人	integer	dropdown	t	10500	t	10500	t	10500	f	f	f	f	f	2018-11-09 13:18:24	\N
7	creatorId	创建人	integer	dropdown	f	10600	f	10600	f	11200	t	f	f	f	f	2018-11-09 13:18:24	\N
8	reporterId	报告人	integer	dropdown	f	10700	f	10700	t	10550	f	f	f	f	f	2018-11-09 13:18:24	\N
9	verId	版本	integer	dropdown	f	10800	f	10800	t	10600	f	f	f	f	f	2018-11-09 13:18:24	\N
10	envId	环境	integer	dropdown	f	10900	f	10900	t	10700	f	f	f	f	f	2018-11-09 13:18:24	\N
11	resolutionId	解决结果	integer	dropdown	f	11000	f	11000	f	11000	f	f	f	f	f	2018-11-09 13:18:24	\N
12	dueTime	截止时间	date	date	f	11100	f	11100	f	10900	f	f	f	f	f	2018-11-09 13:18:24	\N
13	resolveTime	解决时间	date	date	f	11200	f	11200	f	11100	f	f	f	f	f	2018-11-09 13:18:24	\N
14	comments	备注	string	textarea	\N	\N	\N	\N	\N	\N	\N	\N	\N	f	f	2018-11-09 13:18:24	\N
15	resolutionDescr	解决详情	string	textarea	\N	\N	\N	\N	f	20000	f	f	f	f	f	2018-11-09 13:18:24	\N
16	tag	标签	string	text	f	11400	\N	\N	\N	\N	\N	\N	\N	f	f	2018-12-18 08:38:44	\N
1	title	标题	string	text	\N	\N	t	10100	t	10100	f	t	t	f	f	2018-11-09 13:18:24	\N
17	descr	描述	string	textarea	f	11250	\N	\N	t	10800	f	t	f	f	f	2019-02-18 21:49:26.756654	\N
\.


--
-- TOC entry 4588 (class 0 OID 46593)
-- Dependencies: 227
-- Data for Name: IsuHistory; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuHistory" (id, title, descr, "issueId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	用户<span class="dict">Aaron Chen</span>更新	\N	1	f	f	2019-02-18 09:42:11.203498	\N
2	用户<span class="dict">Aaron Chen</span>更新	\N	2	f	f	2019-02-18 21:55:51.674958	\N
3	用户<span class="dict">Aaron Chen</span>更新 描述	\N	2	f	f	2019-02-18 21:58:49.162236	\N
4	用户<span class="dict">Aaron Chen</span>更新	\N	2	f	f	2019-02-18 22:07:40.143484	\N
5	用户<span class="dict">Aaron Chen</span>更新	\N	2	f	f	2019-02-18 22:09:39.938646	\N
6	用户<span class="dict">Aaron Chen</span>更新	\N	2	f	f	2019-02-19 09:22:57.63425	\N
7	用户<span class="dict">Aaron Chen</span>更新	\N	2	f	f	2019-02-19 09:23:31.599651	\N
8	用户<span class="dict">Aaron Chen</span>更新	\N	2	f	f	2019-02-19 10:45:51.586991	\N
9	用户<span class="dict">Aaron Chen</span>更新 标题	\N	2	f	f	2019-02-19 10:47:10.382176	\N
10	用户<span class="dict">Aaron Chen</span>更新 严重级别	\N	2	f	f	2019-02-19 10:49:12.747313	\N
11	用户<span class="dict">Aaron Chen</span>更新	\N	4	f	f	2019-02-19 18:20:50.316727	\N
12	用户Aaron Chen修改状态字段 挂起	\N	4	f	f	2019-02-20 16:57:00.822027	\N
13	用户Aaron Chen修改状态字段 重新打开	\N	4	f	f	2019-02-20 16:57:02.505129	\N
14	用户Aaron Chen修改状态字段 解决	\N	4	f	f	2019-02-20 16:57:03.959051	\N
15	用户Aaron Chen修改状态字段 挂起	\N	4	f	f	2019-02-20 16:57:05.895218	\N
16	用户Aaron Chen修改状态字段 重新打开	\N	4	f	f	2019-02-20 16:57:07.332384	\N
17	用户Aaron Chen更新字段 版本	\N	5	f	f	2019-02-25 10:03:46.621126	\N
18	用户Aaron Chen更新字段 环境	\N	5	f	f	2019-02-25 10:03:51.250531	\N
19	用户Aaron Chen更新字段 描述	\N	5	f	f	2019-02-25 10:04:13.525289	\N
20	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 10:04:19.590098	\N
21	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 10:40:11.711747	\N
22	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 10:40:13.959404	\N
23	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 10:44:37.151655	\N
24	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 10:57:19.838088	\N
25	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 10:57:20.893129	\N
26	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 10:59:01.108444	\N
27	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 11:00:45.553868	\N
28	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 11:03:29.256233	\N
29	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 11:03:52.87926	\N
30	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 11:09:48.046717	\N
31	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 11:09:55.354609	\N
32	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 11:10:03.688743	\N
33	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 11:11:55.807215	\N
34	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 11:13:02.57939	\N
35	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 11:13:05.175434	\N
36	用户Aaron Chen修改状态字段 关闭	\N	5	f	f	2019-02-25 12:24:01.113429	\N
37	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 12:25:51.729474	\N
38	用户Aaron Chen更新	\N	5	f	f	2019-02-25 13:06:05.078284	\N
39	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 13:06:05.245587	\N
40	用户Aaron Chen更新	\N	5	f	f	2019-02-25 13:08:54.198457	\N
41	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 13:08:54.343103	\N
42	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 13:11:32.706694	\N
43	用户Aaron Chen更新	\N	5	f	f	2019-02-25 13:21:02.23394	\N
44	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 13:21:21.690781	\N
45	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 13:21:30.68514	\N
46	用户Aaron Chen更新	\N	5	f	f	2019-02-25 13:26:09.221831	\N
47	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 13:26:09.422242	\N
48	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 13:26:11.761388	\N
49	用户Aaron Chen更新	\N	5	f	f	2019-02-25 13:26:15.883111	\N
50	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 13:26:15.935829	\N
51	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 13:26:55.97745	\N
52	用户Aaron Chen更新	\N	5	f	f	2019-02-25 13:26:58.350877	\N
53	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 13:30:59.113794	\N
54	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 13:31:32.118705	\N
55	用户Aaron Chen更新	\N	5	f	f	2019-02-25 13:31:35.187606	\N
56	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 13:31:41.93643	\N
57	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 13:31:44.90071	\N
58	用户Aaron Chen更新	\N	5	f	f	2019-02-25 13:31:49.958238	\N
59	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 13:31:49.972541	\N
60	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 13:31:51.908531	\N
61	用户Aaron Chen更新	\N	5	f	f	2019-02-25 13:38:25.469908	\N
62	用户Aaron Chen创建	\N	7	f	f	2019-02-25 13:39:07.265504	\N
63	用户Aaron Chen更新	\N	7	f	f	2019-02-25 13:40:39.499438	\N
64	用户Aaron Chen更新	\N	7	f	f	2019-02-25 13:40:59.528285	\N
65	用户Aaron Chen修改状态字段 解决	\N	7	f	f	2019-02-25 13:40:59.580187	\N
66	用户Aaron Chen修改状态字段 重新打开	\N	7	f	f	2019-02-25 13:41:14.858222	\N
67	用户Aaron Chen更新	\N	7	f	f	2019-02-25 15:59:45.287526	\N
68	用户Aaron Chen更新字段 严重级别	\N	7	f	f	2019-02-25 15:59:57.40001	\N
69	用户Aaron Chen创建	\N	8	f	f	2019-02-25 16:01:53.957952	\N
70	用户Aaron Chen更新	\N	8	f	f	2019-02-25 16:02:08.238896	\N
71	用户Aaron Chen分配经办人字段 test01	\N	8	f	f	2019-02-25 16:02:26.534066	\N
72	用户Aaron Chen分配经办人字段 Aaron Chen	\N	8	f	f	2019-02-25 16:03:44.464206	\N
73	用户Aaron Chen分配经办人字段 test01	\N	8	f	f	2019-02-25 16:03:58.256384	\N
74	用户Aaron Chen修改状态字段 关闭	\N	8	f	f	2019-02-25 16:04:13.867137	\N
75	用户Aaron Chen修改状态字段 重新打开	\N	8	f	f	2019-02-25 16:04:14.828468	\N
76	用户Aaron Chen修改状态字段 关闭	\N	8	f	f	2019-02-25 16:04:16.567835	\N
77	用户Aaron Chen修改状态字段 重新打开	\N	8	f	f	2019-02-25 16:04:17.358643	\N
78	用户Aaron Chen更新	\N	8	f	f	2019-02-25 16:04:22.460222	\N
79	用户Aaron Chen修改状态字段 解决	\N	8	f	f	2019-02-25 16:04:22.477622	\N
80	用户Aaron Chen修改状态字段 挂起	\N	8	f	f	2019-02-25 16:04:35.888733	\N
81	用户Aaron Chen修改状态字段 重新打开	\N	8	f	f	2019-02-25 16:04:36.681917	\N
82	用户Aaron Chen更新	\N	8	f	f	2019-02-25 16:04:38.893255	\N
83	用户Aaron Chen修改状态字段 解决	\N	8	f	f	2019-02-25 16:04:38.944509	\N
84	用户Aaron Chen关注问题	\N	8	f	f	2019-02-25 16:05:12.72554	\N
85	用户Aaron Chen修改关注列表	\N	8	f	f	2019-02-25 16:05:24.641566	\N
86	用户Aaron Chen建立链接字段 IS-5	\N	8	f	f	2019-02-25 16:10:00.712275	\N
87	用户Aaron Chen建立链接字段 IS-5	\N	8	f	f	2019-02-25 16:16:12.994654	\N
88	用户Aaron Chen上传附件字段 樊登读书.txt	\N	5	f	f	2019-02-25 16:20:44.988852	\N
89	用户Aaron Chen上传附件字段 樊登读书.txt	\N	5	f	f	2019-02-25 16:22:16.484225	\N
90	用户Aaron Chen删除附件字段 樊登读书.txt	\N	5	f	f	2019-02-25 16:22:32.772315	\N
91	用户Aaron Chen上传附件字段 Dockerfile	\N	5	f	f	2019-02-25 16:22:53.054389	\N
92	用户Aaron Chen创建	\N	9	f	f	2019-02-25 16:39:06.225398	\N
93	用户Aaron Chen更新	\N	5	f	f	2019-02-25 16:54:38.408307	\N
94	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 16:54:38.607807	\N
95	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 16:54:41.528152	\N
96	用户Aaron Chen更新	\N	5	f	f	2019-02-25 16:54:49.94607	\N
97	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 16:54:50.011179	\N
98	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 16:54:51.397873	\N
99	用户Aaron Chen更新	\N	5	f	f	2019-02-25 17:18:21.914195	\N
100	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 17:18:22.019031	\N
101	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 17:18:30.420345	\N
102	用户Aaron Chen更新	\N	5	f	f	2019-02-25 17:18:36.960988	\N
103	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 17:18:36.999675	\N
104	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 18:32:24.52925	\N
105	用户Aaron Chen更新	\N	5	f	f	2019-02-25 18:33:52.656293	\N
106	用户Aaron Chen修改状态字段 解决	\N	5	f	f	2019-02-25 18:33:52.786377	\N
107	用户Aaron Chen修改状态字段 重新打开	\N	5	f	f	2019-02-25 18:33:55.389531	\N
\.


--
-- TOC entry 4590 (class 0 OID 46601)
-- Dependencies: 229
-- Data for Name: IsuIssue; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuIssue" (id, title, "orgId", "projectId", "projectName", "typeId", "statusId", "priorityId", "assigneeId", "creatorId", "reporterId", "resolutionId", "resolutionDescr", "verId", "envId", "dueTime", "resolveTime", "setFinalTime", tag, "createTime", "updateTime", disabled, deleted, uuid, "extProp", tsv_content, descr) FROM stdin;
8	df3	5	10	\N	9	22	19	6	5	5	9	是打发	2	2	\N	\N	2019-02-25 16:04:35.886053	测试,哈哈	2019-02-25 16:01:53.957952	2019-02-25 16:20:07.376153	f	t	4932f8c48a0f497799f5408d865032f3	{}	'df3':2B 'sadsf':3 '测试':1A	sadsf
7	sdf	5	10	\N	9	24	19	\N	5	\N	9	111	\N	\N	\N	\N	\N	\N	2019-02-25 13:39:07.265504	2019-02-25 16:20:24.816079	f	t	2a1f4176780f4ec99fa2b865e7465d7d	{"prop01": 25}	'sdf':1B	\N
9	是的第三方	5	10	\N	13	31	25	\N	5	5	\N	\N	\N	\N	\N	\N	\N	\N	2019-02-25 16:39:06.225398	\N	f	f	5b96b52d82fc463ba21beb8ed2ba8a8c	{"prop01": 22}	'是的':1B '第三方':2B	\N
5	大富大贵	5	10	\N	9	24	19	5	5	5	13	SDFDS	2	2	\N	\N	2019-02-25 12:24:01.087588	\N	2019-02-22 08:33:31.929668	2019-02-25 18:33:55.365602	f	f	uuid	{}	'上海':5 '大富':1B '大贵':2B '江苏':3 '纽约':6 '苏州':4	江苏苏州上海纽约
4	北京天安门	4	8	\N	7	19	15	4	4	4	8	\N	\N	\N	\N	\N	2019-02-20 16:57:05.892954	苏州市工业园区	2019-02-19 11:38:28.481573	2019-02-20 16:57:07.290431	f	f	uuid	{}	'北京':3B '天安门':4B '工业园区':2A '苏州市':1A	\N
6	示例缺陷	6	15	\N	11	26	23	5	5	5	12	\N	\N	\N	\N	\N	\N	\N	2019-02-23 13:16:23.128256	\N	f	f	uuid	{}	'示例':1B '缺陷':2B	\N
1	北京朝阳区	1	2	是的范德萨发	1	1	3	1	1	1	2	\N	\N	\N	\N	\N	\N	东西,南京,闪电,第三方的	2019-02-17 14:22:02.073523	\N	f	f	uuid	{"prop01": 2}	'东西':1A '北京':5B '南京':2A '朝阳区':6B '第三方':4A '闪电':3A	\N
2	北京朝阳区	2	4	\N	3	6	7	2	2	2	4	\N	\N	\N	\N	\N	\N	\N	2019-02-18 21:51:44.67068	2019-02-19 10:49:12.747313	f	f	uuid	{"prop01": 10, "prop001": "江苏苏州"}	'上海':7 '北京':1B '南京':8 '朝阳区':2B '江苏':4C '苏州':5C,6 '阻塞':3C	苏州上海南京
10	示例缺陷	7	17	\N	14	34	28	11	11	11	15	\N	\N	\N	\N	\N	\N	\N	2019-03-11 00:56:22.895443	\N	f	f	uuid	{}	'示例':1B '缺陷':2B	\N
11	示例缺陷	8	19	\N	16	39	32	12	12	12	17	\N	\N	\N	\N	\N	\N	\N	2019-03-12 21:42:50.71871	\N	f	f	uuid	{}	'示例':1B '缺陷':2B	\N
12	示例缺陷	9	21	\N	18	44	36	13	13	13	19	\N	\N	\N	\N	\N	\N	\N	2019-03-12 21:44:18.566135	\N	f	f	uuid	{}	'示例':1B '缺陷':2B	\N
13	示例缺陷	10	23	\N	20	49	40	14	14	14	21	\N	\N	\N	\N	\N	\N	\N	2019-03-12 22:44:53.460728	\N	f	f	uuid	{}	'示例':1B '缺陷':2B	\N
14	示例缺陷	11	25	\N	22	54	44	15	15	15	23	\N	\N	\N	\N	\N	\N	\N	2019-03-12 22:52:14.395588	\N	f	f	uuid	{}	'示例':1B '缺陷':2B	\N
15	示例缺陷	12	27	\N	24	59	48	16	16	16	25	\N	\N	\N	\N	\N	\N	\N	2019-03-12 23:06:01.367716	\N	f	f	uuid	{}	'示例':1B '缺陷':2B	\N
16	示例缺陷	13	29	\N	26	64	52	17	17	17	27	\N	\N	\N	\N	\N	\N	\N	2019-03-12 23:11:45.58636	\N	f	f	uuid	{}	'示例':1B '缺陷':2B	\N
17	示例缺陷	14	31	\N	28	69	56	18	18	18	29	\N	\N	\N	\N	\N	\N	\N	2019-03-12 23:27:21.303261	\N	f	f	uuid	{}	'示例':1B '缺陷':2B	\N
\.


--
-- TOC entry 4591 (class 0 OID 46607)
-- Dependencies: 230
-- Data for Name: IsuIssueExt; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuIssueExt" (pid, prop01, prop02, prop03, prop04, prop05, prop06, prop07, prop08, prop09, prop10, prop11, prop12, prop13, prop14, prop15, prop16, prop17, prop18, prop19, prop20, prop21, prop22, prop23, prop24, prop25, prop26, prop27, prop28, prop29, prop30, prop31, prop32, prop33, prop34, prop35, prop36, prop37, prop38, prop39, prop40, prop41, prop42, prop43, prop44, prop45, prop46, prop47, prop48, prop49, prop50) FROM stdin;
\.


--
-- TOC entry 4594 (class 0 OID 46617)
-- Dependencies: 233
-- Data for Name: IsuLink; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuLink" (id, "reasonId", "reasonName", "srcIssueId", "dictIssueId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	10	重复	8	5	\N	\N	\N	\N
2	20	重复于	8	5	\N	\N	\N	\N
\.


--
-- TOC entry 4595 (class 0 OID 46620)
-- Dependencies: 234
-- Data for Name: IsuLinkReasonDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuLinkReasonDefine" (id, label, value, disabled, deleted, "createTime", "updateTime") FROM stdin;
20	重复于	\N	f	f	2018-12-18 08:59:57	\N
30	阻塞	\N	f	f	2018-12-18 09:03:19	\N
40	阻塞于	\N	f	f	2018-12-18 09:00:19	\N
50	相关于	\N	f	f	2018-12-18 09:03:22	\N
10	重复	\N	t	f	2018-12-18 09:03:16	\N
\.


--
-- TOC entry 4598 (class 0 OID 46630)
-- Dependencies: 237
-- Data for Name: IsuNotification; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuNotification" (id, name, descr, "orgId", disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4599 (class 0 OID 46636)
-- Dependencies: 238
-- Data for Name: IsuNotificationDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuNotificationDefine" (id, name, code, descr, disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4602 (class 0 OID 46646)
-- Dependencies: 241
-- Data for Name: IsuPage; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuPage" (id, name, descr, "orgId", "defaultVal", "buildIn", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	默认界面	\N	1	t	t	f	f	2019-02-17 14:22:02.073523	\N
2	默认问题解决界面	\N	1	f	t	f	f	2019-02-17 14:22:02.073523	\N
3	默认界面	\N	2	t	t	f	f	2019-02-18 21:51:44.67068	\N
4	默认问题解决界面	\N	2	f	t	f	f	2019-02-18 21:51:44.67068	\N
5	默认界面	\N	4	t	t	f	f	2019-02-19 11:38:28.481573	\N
6	默认问题解决界面	\N	4	f	t	f	f	2019-02-19 11:38:28.481573	\N
7	默认界面	\N	5	t	t	f	f	2019-02-22 08:33:31.929668	\N
8	默认问题解决界面	\N	5	f	t	f	f	2019-02-22 08:33:31.929668	\N
9	默认界面	\N	6	t	t	f	f	2019-02-23 13:16:23.128256	\N
10	默认问题解决界面	\N	6	f	t	f	f	2019-02-23 13:16:23.128256	\N
11	默认界面	\N	7	t	t	f	f	2019-03-11 00:56:22.895443	\N
12	默认问题解决界面	\N	7	f	t	f	f	2019-03-11 00:56:22.895443	\N
13	默认界面	\N	8	t	t	f	f	2019-03-12 21:42:50.71871	\N
14	默认问题解决界面	\N	8	f	t	f	f	2019-03-12 21:42:50.71871	\N
15	默认界面	\N	9	t	t	f	f	2019-03-12 21:44:18.566135	\N
16	默认问题解决界面	\N	9	f	t	f	f	2019-03-12 21:44:18.566135	\N
17	默认界面	\N	10	t	t	f	f	2019-03-12 22:44:53.460728	\N
18	默认问题解决界面	\N	10	f	t	f	f	2019-03-12 22:44:53.460728	\N
19	默认界面	\N	11	t	t	f	f	2019-03-12 22:52:14.395588	\N
20	默认问题解决界面	\N	11	f	t	f	f	2019-03-12 22:52:14.395588	\N
21	默认界面	\N	12	t	t	f	f	2019-03-12 23:06:01.367716	\N
22	默认问题解决界面	\N	12	f	t	f	f	2019-03-12 23:06:01.367716	\N
23	默认界面	\N	13	t	t	f	f	2019-03-12 23:11:45.58636	\N
24	默认问题解决界面	\N	13	f	t	f	f	2019-03-12 23:11:45.58636	\N
25	默认界面	\N	14	t	t	f	f	2019-03-12 23:27:21.303261	\N
26	默认问题解决界面	\N	14	f	t	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4603 (class 0 OID 46652)
-- Dependencies: 242
-- Data for Name: IsuPageElement; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuPageElement" (id, "colCode", label, type, input, "fullLine", required, "buildIn", key, "fieldId", "pageId", "orgId", ordr, readonly, disabled, deleted, "createTime", "updateTime") FROM stdin;
10	resolutionId	解决结果	integer	dropdown	f	f	t	sys-10	10	2	1	11000	f	f	f	2019-02-17 14:22:02.073523	\N
11	resolutionDescr	解决详情	string	textarea	f	f	t	sys-13	13	2	1	20000	f	f	f	2019-02-17 14:22:02.073523	\N
1	title	标题	string	text	t	t	t	sys-1	1	1	1	1	f	f	f	2019-02-17 14:22:02.073523	\N
2	statusId	状态	integer	dropdown	f	f	t	sys-3	3	1	1	2	f	f	f	2019-02-17 14:22:02.073523	\N
3	typeId	类型	integer	dropdown	f	f	t	sys-2	2	1	1	3	f	f	f	2019-02-17 14:22:02.073523	\N
4	priorityId	优先级	integer	dropdown	f	f	t	sys-4	4	1	1	4	f	f	f	2019-02-17 14:22:02.073523	\N
5	assigneeId	经办人	integer	dropdown	f	f	t	sys-5	5	1	1	5	f	f	f	2019-02-17 14:22:02.073523	\N
6	reporterId	报告人	integer	dropdown	f	f	t	sys-7	7	1	1	6	f	f	f	2019-02-17 14:22:02.073523	\N
7	verId	版本	integer	dropdown	f	f	t	sys-8	8	1	1	7	f	f	f	2019-02-17 14:22:02.073523	\N
8	envId	环境	integer	dropdown	f	f	t	sys-9	9	1	1	8	f	f	f	2019-02-17 14:22:02.073523	\N
9	descr	描述	string	textarea	t	f	t	sys-14	14	1	1	10	f	f	f	2019-02-17 14:22:02.073523	\N
12	prop01	严重级别	integer	dropdown	f	f	f	cust-1	1	1	1	9	f	f	f	2019-02-18 09:41:58.732147	\N
22	resolutionId	解决结果	integer	dropdown	f	f	t	sys-23	23	4	2	11000	f	f	f	2019-02-18 21:51:44.67068	\N
23	resolutionDescr	解决详情	string	textarea	f	f	t	sys-26	26	4	2	20000	f	f	f	2019-02-18 21:51:44.67068	\N
13	title	标题	string	text	t	t	t	sys-27	27	3	2	1	f	f	f	2019-02-18 21:51:44.67068	\N
14	statusId	状态	integer	dropdown	f	f	t	sys-16	16	3	2	2	f	f	f	2019-02-18 21:51:44.67068	\N
15	typeId	类型	integer	dropdown	f	f	t	sys-15	15	3	2	3	f	f	f	2019-02-18 21:51:44.67068	\N
16	priorityId	优先级	integer	dropdown	f	f	t	sys-17	17	3	2	4	f	f	f	2019-02-18 21:51:44.67068	\N
17	assigneeId	经办人	integer	dropdown	f	f	t	sys-18	18	3	2	5	f	f	f	2019-02-18 21:51:44.67068	\N
18	reporterId	报告人	integer	dropdown	f	f	t	sys-20	20	3	2	6	f	f	f	2019-02-18 21:51:44.67068	\N
19	verId	版本	integer	dropdown	f	f	t	sys-21	21	3	2	7	f	f	f	2019-02-18 21:51:44.67068	\N
20	envId	环境	integer	dropdown	f	f	t	sys-22	22	3	2	10	f	f	f	2019-02-18 21:51:44.67068	\N
21	descr	描述	string	textarea	t	f	t	sys-28	28	3	2	11	f	f	f	2019-02-18 21:51:44.67068	\N
24	prop01	严重级别	integer	dropdown	f	f	f	cust-2	2	3	2	8	f	f	f	2019-02-19 09:22:19.778804	\N
25	prop001	字符串字段01	string	text	f	f	f	cust-3	3	3	2	9	f	f	f	2019-02-19 09:22:21.561513	\N
35	resolutionId	解决结果	integer	dropdown	f	f	t	sys-51	51	6	4	11000	f	f	f	2019-02-19 11:38:28.481573	\N
47	resolutionId	解决结果	integer	dropdown	f	f	t	sys-65	65	8	5	2	f	f	f	2019-02-22 08:33:31.929668	\N
48	resolutionDescr	解决详情	string	textarea	t	f	t	sys-68	68	8	5	3	f	f	f	2019-02-22 08:33:31.929668	\N
36	resolutionDescr	解决详情	string	textarea	f	f	t	sys-54	54	6	4	20000	f	f	f	2019-02-19 11:38:28.481573	\N
26	title	标题	string	text	t	t	t	sys-55	55	5	4	1	f	f	f	2019-02-19 11:38:28.481573	\N
27	statusId	状态	integer	dropdown	f	f	t	sys-44	44	5	4	2	f	f	f	2019-02-19 11:38:28.481573	\N
28	typeId	类型	integer	dropdown	f	f	t	sys-43	43	5	4	3	f	f	f	2019-02-19 11:38:28.481573	\N
29	priorityId	优先级	integer	dropdown	f	f	t	sys-45	45	5	4	4	f	f	f	2019-02-19 11:38:28.481573	\N
30	assigneeId	经办人	integer	dropdown	f	f	t	sys-46	46	5	4	5	f	f	f	2019-02-19 11:38:28.481573	\N
31	reporterId	报告人	integer	dropdown	f	f	t	sys-48	48	5	4	6	f	f	f	2019-02-19 11:38:28.481573	\N
32	verId	版本	integer	dropdown	f	f	t	sys-49	49	5	4	7	f	f	f	2019-02-19 11:38:28.481573	\N
33	envId	环境	integer	dropdown	f	f	t	sys-50	50	5	4	8	f	f	f	2019-02-19 11:38:28.481573	\N
34	descr	描述	string	textarea	t	f	t	sys-56	56	5	4	10	f	f	f	2019-02-19 11:38:28.481573	\N
37	prop01	严重级别	integer	dropdown	f	f	f	cust-5	5	5	4	9	f	f	f	2019-02-20 17:05:13.556831	\N
49	title	标题	string	text	t	t	t	sys-83	83	9	6	10100	f	f	f	2019-02-23 13:16:23.128256	\N
38	title	标题	string	text	t	t	t	sys-69	69	7	5	1	f	f	f	2019-02-22 08:33:31.929668	\N
39	statusId	状态	integer	dropdown	f	f	t	sys-58	58	7	5	2	f	f	f	2019-02-22 08:33:31.929668	\N
40	typeId	类型	integer	dropdown	f	f	t	sys-57	57	7	5	3	f	f	f	2019-02-22 08:33:31.929668	\N
41	priorityId	优先级	integer	dropdown	f	f	t	sys-59	59	7	5	4	f	f	f	2019-02-22 08:33:31.929668	\N
42	assigneeId	经办人	integer	dropdown	f	f	t	sys-60	60	7	5	5	f	f	f	2019-02-22 08:33:31.929668	\N
43	reporterId	报告人	integer	dropdown	f	f	t	sys-62	62	7	5	6	f	f	f	2019-02-22 08:33:31.929668	\N
44	verId	版本	integer	dropdown	f	f	t	sys-63	63	7	5	7	f	f	f	2019-02-22 08:33:31.929668	\N
50	statusId	状态	integer	dropdown	f	f	t	sys-72	72	9	6	10150	f	f	f	2019-02-23 13:16:23.128256	\N
51	typeId	类型	integer	dropdown	f	f	t	sys-71	71	9	6	10200	f	f	f	2019-02-23 13:16:23.128256	\N
52	priorityId	优先级	integer	dropdown	f	f	t	sys-73	73	9	6	10400	f	f	f	2019-02-23 13:16:23.128256	\N
53	assigneeId	经办人	integer	dropdown	f	f	t	sys-74	74	9	6	10500	f	f	f	2019-02-23 13:16:23.128256	\N
54	reporterId	报告人	integer	dropdown	f	f	t	sys-76	76	9	6	10550	f	f	f	2019-02-23 13:16:23.128256	\N
55	verId	版本	integer	dropdown	f	f	t	sys-77	77	9	6	10600	f	f	f	2019-02-23 13:16:23.128256	\N
56	envId	环境	integer	dropdown	f	f	t	sys-78	78	9	6	10700	f	f	f	2019-02-23 13:16:23.128256	\N
57	descr	描述	string	textarea	t	f	t	sys-84	84	9	6	10800	f	f	f	2019-02-23 13:16:23.128256	\N
58	resolutionId	解决结果	integer	dropdown	f	f	t	sys-79	79	10	6	11000	f	f	f	2019-02-23 13:16:23.128256	\N
59	resolutionDescr	解决详情	string	textarea	f	f	t	sys-82	82	10	6	20000	f	f	f	2019-02-23 13:16:23.128256	\N
45	envId	环境	integer	dropdown	f	f	t	sys-64	64	7	5	8	f	f	f	2019-02-22 08:33:31.929668	\N
46	descr	描述	string	textarea	t	f	t	sys-70	70	7	5	10	f	f	f	2019-02-22 08:33:31.929668	\N
60	prop01	严重级别	integer	dropdown	f	f	f	cust-6	6	7	5	9	f	f	f	2019-02-25 13:38:01.434077	\N
61	title	标题	string	text	t	t	t	sys-69	69	8	5	1	f	f	f	2019-02-25 18:33:10.652212	\N
62	title	标题	string	text	t	t	t	sys-97	97	11	7	10100	f	f	f	2019-03-11 00:56:22.895443	\N
63	statusId	状态	integer	dropdown	f	f	t	sys-86	86	11	7	10150	f	f	f	2019-03-11 00:56:22.895443	\N
64	typeId	类型	integer	dropdown	f	f	t	sys-85	85	11	7	10200	f	f	f	2019-03-11 00:56:22.895443	\N
65	priorityId	优先级	integer	dropdown	f	f	t	sys-87	87	11	7	10400	f	f	f	2019-03-11 00:56:22.895443	\N
66	assigneeId	经办人	integer	dropdown	f	f	t	sys-88	88	11	7	10500	f	f	f	2019-03-11 00:56:22.895443	\N
67	reporterId	报告人	integer	dropdown	f	f	t	sys-90	90	11	7	10550	f	f	f	2019-03-11 00:56:22.895443	\N
68	verId	版本	integer	dropdown	f	f	t	sys-91	91	11	7	10600	f	f	f	2019-03-11 00:56:22.895443	\N
69	envId	环境	integer	dropdown	f	f	t	sys-92	92	11	7	10700	f	f	f	2019-03-11 00:56:22.895443	\N
70	descr	描述	string	textarea	t	f	t	sys-98	98	11	7	10800	f	f	f	2019-03-11 00:56:22.895443	\N
71	resolutionId	解决结果	integer	dropdown	f	f	t	sys-93	93	12	7	11000	f	f	f	2019-03-11 00:56:22.895443	\N
72	resolutionDescr	解决详情	string	textarea	f	f	t	sys-96	96	12	7	20000	f	f	f	2019-03-11 00:56:22.895443	\N
73	title	标题	string	text	t	t	t	sys-111	111	13	8	10100	f	f	f	2019-03-12 21:42:50.71871	\N
74	statusId	状态	integer	dropdown	f	f	t	sys-100	100	13	8	10150	f	f	f	2019-03-12 21:42:50.71871	\N
75	typeId	类型	integer	dropdown	f	f	t	sys-99	99	13	8	10200	f	f	f	2019-03-12 21:42:50.71871	\N
76	priorityId	优先级	integer	dropdown	f	f	t	sys-101	101	13	8	10400	f	f	f	2019-03-12 21:42:50.71871	\N
77	assigneeId	经办人	integer	dropdown	f	f	t	sys-102	102	13	8	10500	f	f	f	2019-03-12 21:42:50.71871	\N
78	reporterId	报告人	integer	dropdown	f	f	t	sys-104	104	13	8	10550	f	f	f	2019-03-12 21:42:50.71871	\N
79	verId	版本	integer	dropdown	f	f	t	sys-105	105	13	8	10600	f	f	f	2019-03-12 21:42:50.71871	\N
80	envId	环境	integer	dropdown	f	f	t	sys-106	106	13	8	10700	f	f	f	2019-03-12 21:42:50.71871	\N
81	descr	描述	string	textarea	t	f	t	sys-112	112	13	8	10800	f	f	f	2019-03-12 21:42:50.71871	\N
82	resolutionId	解决结果	integer	dropdown	f	f	t	sys-107	107	14	8	11000	f	f	f	2019-03-12 21:42:50.71871	\N
83	resolutionDescr	解决详情	string	textarea	f	f	t	sys-110	110	14	8	20000	f	f	f	2019-03-12 21:42:50.71871	\N
84	title	标题	string	text	t	t	t	sys-125	125	15	9	10100	f	f	f	2019-03-12 21:44:18.566135	\N
85	statusId	状态	integer	dropdown	f	f	t	sys-114	114	15	9	10150	f	f	f	2019-03-12 21:44:18.566135	\N
86	typeId	类型	integer	dropdown	f	f	t	sys-113	113	15	9	10200	f	f	f	2019-03-12 21:44:18.566135	\N
87	priorityId	优先级	integer	dropdown	f	f	t	sys-115	115	15	9	10400	f	f	f	2019-03-12 21:44:18.566135	\N
88	assigneeId	经办人	integer	dropdown	f	f	t	sys-116	116	15	9	10500	f	f	f	2019-03-12 21:44:18.566135	\N
89	reporterId	报告人	integer	dropdown	f	f	t	sys-118	118	15	9	10550	f	f	f	2019-03-12 21:44:18.566135	\N
90	verId	版本	integer	dropdown	f	f	t	sys-119	119	15	9	10600	f	f	f	2019-03-12 21:44:18.566135	\N
91	envId	环境	integer	dropdown	f	f	t	sys-120	120	15	9	10700	f	f	f	2019-03-12 21:44:18.566135	\N
92	descr	描述	string	textarea	t	f	t	sys-126	126	15	9	10800	f	f	f	2019-03-12 21:44:18.566135	\N
93	resolutionId	解决结果	integer	dropdown	f	f	t	sys-121	121	16	9	11000	f	f	f	2019-03-12 21:44:18.566135	\N
94	resolutionDescr	解决详情	string	textarea	f	f	t	sys-124	124	16	9	20000	f	f	f	2019-03-12 21:44:18.566135	\N
95	title	标题	string	text	t	t	t	sys-139	139	17	10	10100	f	f	f	2019-03-12 22:44:53.460728	\N
96	statusId	状态	integer	dropdown	f	f	t	sys-128	128	17	10	10150	f	f	f	2019-03-12 22:44:53.460728	\N
97	typeId	类型	integer	dropdown	f	f	t	sys-127	127	17	10	10200	f	f	f	2019-03-12 22:44:53.460728	\N
98	priorityId	优先级	integer	dropdown	f	f	t	sys-129	129	17	10	10400	f	f	f	2019-03-12 22:44:53.460728	\N
99	assigneeId	经办人	integer	dropdown	f	f	t	sys-130	130	17	10	10500	f	f	f	2019-03-12 22:44:53.460728	\N
100	reporterId	报告人	integer	dropdown	f	f	t	sys-132	132	17	10	10550	f	f	f	2019-03-12 22:44:53.460728	\N
101	verId	版本	integer	dropdown	f	f	t	sys-133	133	17	10	10600	f	f	f	2019-03-12 22:44:53.460728	\N
102	envId	环境	integer	dropdown	f	f	t	sys-134	134	17	10	10700	f	f	f	2019-03-12 22:44:53.460728	\N
103	descr	描述	string	textarea	t	f	t	sys-140	140	17	10	10800	f	f	f	2019-03-12 22:44:53.460728	\N
104	resolutionId	解决结果	integer	dropdown	f	f	t	sys-135	135	18	10	11000	f	f	f	2019-03-12 22:44:53.460728	\N
105	resolutionDescr	解决详情	string	textarea	f	f	t	sys-138	138	18	10	20000	f	f	f	2019-03-12 22:44:53.460728	\N
106	title	标题	string	text	t	t	t	sys-153	153	19	11	10100	f	f	f	2019-03-12 22:52:14.395588	\N
107	statusId	状态	integer	dropdown	f	f	t	sys-142	142	19	11	10150	f	f	f	2019-03-12 22:52:14.395588	\N
108	typeId	类型	integer	dropdown	f	f	t	sys-141	141	19	11	10200	f	f	f	2019-03-12 22:52:14.395588	\N
109	priorityId	优先级	integer	dropdown	f	f	t	sys-143	143	19	11	10400	f	f	f	2019-03-12 22:52:14.395588	\N
110	assigneeId	经办人	integer	dropdown	f	f	t	sys-144	144	19	11	10500	f	f	f	2019-03-12 22:52:14.395588	\N
111	reporterId	报告人	integer	dropdown	f	f	t	sys-146	146	19	11	10550	f	f	f	2019-03-12 22:52:14.395588	\N
112	verId	版本	integer	dropdown	f	f	t	sys-147	147	19	11	10600	f	f	f	2019-03-12 22:52:14.395588	\N
113	envId	环境	integer	dropdown	f	f	t	sys-148	148	19	11	10700	f	f	f	2019-03-12 22:52:14.395588	\N
114	descr	描述	string	textarea	t	f	t	sys-154	154	19	11	10800	f	f	f	2019-03-12 22:52:14.395588	\N
115	resolutionId	解决结果	integer	dropdown	f	f	t	sys-149	149	20	11	11000	f	f	f	2019-03-12 22:52:14.395588	\N
116	resolutionDescr	解决详情	string	textarea	f	f	t	sys-152	152	20	11	20000	f	f	f	2019-03-12 22:52:14.395588	\N
117	title	标题	string	text	t	t	t	sys-167	167	21	12	10100	f	f	f	2019-03-12 23:06:01.367716	\N
118	statusId	状态	integer	dropdown	f	f	t	sys-156	156	21	12	10150	f	f	f	2019-03-12 23:06:01.367716	\N
119	typeId	类型	integer	dropdown	f	f	t	sys-155	155	21	12	10200	f	f	f	2019-03-12 23:06:01.367716	\N
120	priorityId	优先级	integer	dropdown	f	f	t	sys-157	157	21	12	10400	f	f	f	2019-03-12 23:06:01.367716	\N
121	assigneeId	经办人	integer	dropdown	f	f	t	sys-158	158	21	12	10500	f	f	f	2019-03-12 23:06:01.367716	\N
122	reporterId	报告人	integer	dropdown	f	f	t	sys-160	160	21	12	10550	f	f	f	2019-03-12 23:06:01.367716	\N
123	verId	版本	integer	dropdown	f	f	t	sys-161	161	21	12	10600	f	f	f	2019-03-12 23:06:01.367716	\N
124	envId	环境	integer	dropdown	f	f	t	sys-162	162	21	12	10700	f	f	f	2019-03-12 23:06:01.367716	\N
125	descr	描述	string	textarea	t	f	t	sys-168	168	21	12	10800	f	f	f	2019-03-12 23:06:01.367716	\N
126	resolutionId	解决结果	integer	dropdown	f	f	t	sys-163	163	22	12	11000	f	f	f	2019-03-12 23:06:01.367716	\N
127	resolutionDescr	解决详情	string	textarea	f	f	t	sys-166	166	22	12	20000	f	f	f	2019-03-12 23:06:01.367716	\N
128	title	标题	string	text	t	t	t	sys-181	181	23	13	10100	f	f	f	2019-03-12 23:11:45.58636	\N
129	statusId	状态	integer	dropdown	f	f	t	sys-170	170	23	13	10150	f	f	f	2019-03-12 23:11:45.58636	\N
130	typeId	类型	integer	dropdown	f	f	t	sys-169	169	23	13	10200	f	f	f	2019-03-12 23:11:45.58636	\N
131	priorityId	优先级	integer	dropdown	f	f	t	sys-171	171	23	13	10400	f	f	f	2019-03-12 23:11:45.58636	\N
132	assigneeId	经办人	integer	dropdown	f	f	t	sys-172	172	23	13	10500	f	f	f	2019-03-12 23:11:45.58636	\N
133	reporterId	报告人	integer	dropdown	f	f	t	sys-174	174	23	13	10550	f	f	f	2019-03-12 23:11:45.58636	\N
134	verId	版本	integer	dropdown	f	f	t	sys-175	175	23	13	10600	f	f	f	2019-03-12 23:11:45.58636	\N
135	envId	环境	integer	dropdown	f	f	t	sys-176	176	23	13	10700	f	f	f	2019-03-12 23:11:45.58636	\N
136	descr	描述	string	textarea	t	f	t	sys-182	182	23	13	10800	f	f	f	2019-03-12 23:11:45.58636	\N
137	resolutionId	解决结果	integer	dropdown	f	f	t	sys-177	177	24	13	11000	f	f	f	2019-03-12 23:11:45.58636	\N
138	resolutionDescr	解决详情	string	textarea	f	f	t	sys-180	180	24	13	20000	f	f	f	2019-03-12 23:11:45.58636	\N
139	title	标题	string	text	t	t	t	sys-195	195	25	14	10100	f	f	f	2019-03-12 23:27:21.303261	\N
140	statusId	状态	integer	dropdown	f	f	t	sys-184	184	25	14	10150	f	f	f	2019-03-12 23:27:21.303261	\N
141	typeId	类型	integer	dropdown	f	f	t	sys-183	183	25	14	10200	f	f	f	2019-03-12 23:27:21.303261	\N
142	priorityId	优先级	integer	dropdown	f	f	t	sys-185	185	25	14	10400	f	f	f	2019-03-12 23:27:21.303261	\N
143	assigneeId	经办人	integer	dropdown	f	f	t	sys-186	186	25	14	10500	f	f	f	2019-03-12 23:27:21.303261	\N
144	reporterId	报告人	integer	dropdown	f	f	t	sys-188	188	25	14	10550	f	f	f	2019-03-12 23:27:21.303261	\N
145	verId	版本	integer	dropdown	f	f	t	sys-189	189	25	14	10600	f	f	f	2019-03-12 23:27:21.303261	\N
146	envId	环境	integer	dropdown	f	f	t	sys-190	190	25	14	10700	f	f	f	2019-03-12 23:27:21.303261	\N
147	descr	描述	string	textarea	t	f	t	sys-196	196	25	14	10800	f	f	f	2019-03-12 23:27:21.303261	\N
148	resolutionId	解决结果	integer	dropdown	f	f	t	sys-191	191	26	14	11000	f	f	f	2019-03-12 23:27:21.303261	\N
149	resolutionDescr	解决详情	string	textarea	f	f	t	sys-194	194	26	14	20000	f	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4605 (class 0 OID 46660)
-- Dependencies: 244
-- Data for Name: IsuPageSolution; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuPageSolution" (id, name, descr, "orgId", "defaultVal", "buildIn", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	默认界面方案	\N	1	t	\N	f	f	2019-02-17 14:22:02.073523	\N
2	默认界面方案	\N	2	t	\N	f	f	2019-02-18 21:51:44.67068	\N
3	默认界面方案	\N	4	t	\N	f	f	2019-02-19 11:38:28.481573	\N
4	默认界面方案	\N	5	t	\N	f	f	2019-02-22 08:33:31.929668	\N
5	默认界面方案	\N	6	t	\N	f	f	2019-02-23 13:16:23.128256	\N
6	界面方案01	\N	5	f	f	f	f	2019-02-25 17:23:49.924244	\N
7	xcvxc	\N	5	f	f	f	f	2019-02-25 18:16:33.617275	\N
8	sdfdsf	\N	5	f	f	f	f	2019-02-25 18:21:14.014288	\N
9	333	\N	5	f	f	f	f	2019-02-25 18:22:27.91503	\N
10	默认界面方案	\N	7	t	\N	f	f	2019-03-11 00:56:22.895443	\N
11	默认界面方案	\N	8	t	\N	f	f	2019-03-12 21:42:50.71871	\N
12	默认界面方案	\N	9	t	\N	f	f	2019-03-12 21:44:18.566135	\N
13	默认界面方案	\N	10	t	\N	f	f	2019-03-12 22:44:53.460728	\N
14	默认界面方案	\N	11	t	\N	f	f	2019-03-12 22:52:14.395588	\N
15	默认界面方案	\N	12	t	\N	f	f	2019-03-12 23:06:01.367716	\N
16	默认界面方案	\N	13	t	\N	f	f	2019-03-12 23:11:45.58636	\N
17	默认界面方案	\N	14	t	\N	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4606 (class 0 OID 46666)
-- Dependencies: 245
-- Data for Name: IsuPageSolutionItem; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuPageSolutionItem" (id, "typeId", opt, "pageId", "solutionId", "orgId") FROM stdin;
1	2	create	1	1	1
2	2	edit	1	1	1
3	2	view	1	1	1
4	1	create	1	1	1
5	1	edit	1	1	1
6	1	view	1	1	1
7	4	create	3	2	2
8	4	edit	3	2	2
9	4	view	3	2	2
10	3	create	3	2	2
11	3	edit	3	2	2
12	3	view	3	2	2
13	8	create	5	3	4
14	8	edit	5	3	4
15	8	view	5	3	4
16	7	create	5	3	4
17	7	edit	5	3	4
18	7	view	5	3	4
19	10	create	7	4	5
20	10	edit	7	4	5
21	10	view	7	4	5
23	9	edit	7	4	5
24	9	view	7	4	5
25	12	create	9	5	6
26	12	edit	9	5	6
27	12	view	9	5	6
28	11	create	9	5	6
29	11	edit	9	5	6
30	11	view	9	5	6
22	9	create	7	4	5
31	9	create	7	6	5
32	10	create	7	6	5
33	13	create	7	9	5
34	13	edit	7	9	5
35	13	view	7	9	5
36	10	create	7	9	5
37	10	edit	7	9	5
38	10	view	7	9	5
39	9	create	7	9	5
40	9	edit	7	9	5
41	9	view	7	9	5
42	15	create	11	10	7
43	15	edit	11	10	7
44	15	view	11	10	7
45	14	create	11	10	7
46	14	edit	11	10	7
47	14	view	11	10	7
48	17	create	13	11	8
49	17	edit	13	11	8
50	17	view	13	11	8
51	16	create	13	11	8
52	16	edit	13	11	8
53	16	view	13	11	8
54	19	create	15	12	9
55	19	edit	15	12	9
56	19	view	15	12	9
57	18	create	15	12	9
58	18	edit	15	12	9
59	18	view	15	12	9
60	21	create	17	13	10
61	21	edit	17	13	10
62	21	view	17	13	10
63	20	create	17	13	10
64	20	edit	17	13	10
65	20	view	17	13	10
66	23	create	19	14	11
67	23	edit	19	14	11
68	23	view	19	14	11
69	22	create	19	14	11
70	22	edit	19	14	11
71	22	view	19	14	11
72	25	create	21	15	12
73	25	edit	21	15	12
74	25	view	21	15	12
75	24	create	21	15	12
76	24	edit	21	15	12
77	24	view	21	15	12
78	27	create	23	16	13
79	27	edit	23	16	13
80	27	view	23	16	13
81	26	create	23	16	13
82	26	edit	23	16	13
83	26	view	23	16	13
84	29	create	25	17	14
85	29	edit	25	17	14
86	29	view	25	17	14
87	28	create	25	17	14
88	28	edit	25	17	14
89	28	view	25	17	14
\.


--
-- TOC entry 4610 (class 0 OID 46675)
-- Dependencies: 249
-- Data for Name: IsuPriority; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuPriority" (id, label, value, descr, "defaultVal", "buildIn", ordr, "orgId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	紧急	urgent	\N	f	t	1	1	f	f	2019-02-17 14:22:02.073523	\N
2	高	high	\N	f	t	2	1	f	f	2019-02-17 14:22:02.073523	\N
3	中	medium	\N	t	t	3	1	f	f	2019-02-17 14:22:02.073523	\N
4	低	low	\N	f	t	4	1	f	f	2019-02-17 14:22:02.073523	\N
5	紧急	urgent	\N	f	t	1	2	f	f	2019-02-18 21:51:44.67068	\N
6	高	high	\N	f	t	2	2	f	f	2019-02-18 21:51:44.67068	\N
7	中	medium	\N	t	t	3	2	f	f	2019-02-18 21:51:44.67068	\N
8	低	low	\N	f	t	4	2	f	f	2019-02-18 21:51:44.67068	\N
13	紧急	urgent	\N	f	t	1	4	f	f	2019-02-19 11:38:28.481573	\N
14	高	high	\N	f	t	2	4	f	f	2019-02-19 11:38:28.481573	\N
15	中	medium	\N	t	t	3	4	f	f	2019-02-19 11:38:28.481573	\N
16	低	low	\N	f	t	4	4	f	f	2019-02-19 11:38:28.481573	\N
21	紧急	urgent	\N	f	t	1	6	f	f	2019-02-23 13:16:23.128256	\N
22	高	high	\N	f	t	2	6	f	f	2019-02-23 13:16:23.128256	\N
23	中	medium	\N	t	t	3	6	f	f	2019-02-23 13:16:23.128256	\N
24	低	low	\N	f	t	4	6	f	f	2019-02-23 13:16:23.128256	\N
25	优先级01	\N	\N	f	f	14	5	f	f	2019-02-25 16:37:27.002507	2019-02-25 16:37:33.783984
20	低	low	\N	f	t	4	5	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:37:33.783984
18	高	high	\N	f	t	2	5	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:37:33.783984
17	紧急	urgent	\N	f	t	1	5	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:37:33.783984
19	中	medium	\N	t	t	3	5	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:37:33.783984
26	紧急	urgent	\N	f	t	1	7	f	f	2019-03-11 00:56:22.895443	\N
27	高	high	\N	f	t	2	7	f	f	2019-03-11 00:56:22.895443	\N
28	中	medium	\N	t	t	3	7	f	f	2019-03-11 00:56:22.895443	\N
29	低	low	\N	f	t	4	7	f	f	2019-03-11 00:56:22.895443	\N
30	紧急	urgent	\N	f	t	1	8	f	f	2019-03-12 21:42:50.71871	\N
31	高	high	\N	f	t	2	8	f	f	2019-03-12 21:42:50.71871	\N
32	中	medium	\N	t	t	3	8	f	f	2019-03-12 21:42:50.71871	\N
33	低	low	\N	f	t	4	8	f	f	2019-03-12 21:42:50.71871	\N
34	紧急	urgent	\N	f	t	1	9	f	f	2019-03-12 21:44:18.566135	\N
35	高	high	\N	f	t	2	9	f	f	2019-03-12 21:44:18.566135	\N
36	中	medium	\N	t	t	3	9	f	f	2019-03-12 21:44:18.566135	\N
37	低	low	\N	f	t	4	9	f	f	2019-03-12 21:44:18.566135	\N
38	紧急	urgent	\N	f	t	1	10	f	f	2019-03-12 22:44:53.460728	\N
39	高	high	\N	f	t	2	10	f	f	2019-03-12 22:44:53.460728	\N
40	中	medium	\N	t	t	3	10	f	f	2019-03-12 22:44:53.460728	\N
41	低	low	\N	f	t	4	10	f	f	2019-03-12 22:44:53.460728	\N
42	紧急	urgent	\N	f	t	1	11	f	f	2019-03-12 22:52:14.395588	\N
43	高	high	\N	f	t	2	11	f	f	2019-03-12 22:52:14.395588	\N
44	中	medium	\N	t	t	3	11	f	f	2019-03-12 22:52:14.395588	\N
45	低	low	\N	f	t	4	11	f	f	2019-03-12 22:52:14.395588	\N
46	紧急	urgent	\N	f	t	1	12	f	f	2019-03-12 23:06:01.367716	\N
47	高	high	\N	f	t	2	12	f	f	2019-03-12 23:06:01.367716	\N
48	中	medium	\N	t	t	3	12	f	f	2019-03-12 23:06:01.367716	\N
49	低	low	\N	f	t	4	12	f	f	2019-03-12 23:06:01.367716	\N
50	紧急	urgent	\N	f	t	1	13	f	f	2019-03-12 23:11:45.58636	\N
51	高	high	\N	f	t	2	13	f	f	2019-03-12 23:11:45.58636	\N
52	中	medium	\N	t	t	3	13	f	f	2019-03-12 23:11:45.58636	\N
53	低	low	\N	f	t	4	13	f	f	2019-03-12 23:11:45.58636	\N
54	紧急	urgent	\N	f	t	1	14	f	f	2019-03-12 23:27:21.303261	\N
55	高	high	\N	f	t	2	14	f	f	2019-03-12 23:27:21.303261	\N
56	中	medium	\N	t	t	3	14	f	f	2019-03-12 23:27:21.303261	\N
57	低	low	\N	f	t	4	14	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4611 (class 0 OID 46681)
-- Dependencies: 250
-- Data for Name: IsuPriorityDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuPriorityDefine" (id, label, value, descr, ordr, "defaultVal", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	紧急	urgent	\N	1	f	f	f	2018-11-09 11:28:35	\N
2	高	high	\N	2	f	f	f	2018-11-09 11:28:39	\N
3	中	medium	\N	3	t	f	f	2018-11-09 11:28:42	\N
4	低	low	\N	4	f	f	f	2018-11-09 11:28:45	\N
\.


--
-- TOC entry 4613 (class 0 OID 46689)
-- Dependencies: 252
-- Data for Name: IsuPrioritySolution; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuPrioritySolution" (id, name, descr, "orgId", "defaultVal", "buildIn", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	默认问题优先级方案	\N	1	t	t	f	f	2019-02-17 14:22:02.073523	\N
2	默认问题优先级方案	\N	2	t	t	f	f	2019-02-18 21:51:44.67068	\N
4	默认问题优先级方案	\N	4	t	t	f	f	2019-02-19 11:38:28.481573	\N
6	默认问题优先级方案	\N	6	t	t	f	f	2019-02-23 13:16:23.128256	\N
5	默认问题优先级方案	\N	5	t	t	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:37:45.579144
7	默认问题优先级方案	\N	7	t	t	f	f	2019-03-11 00:56:22.895443	\N
8	默认问题优先级方案	\N	8	t	t	f	f	2019-03-12 21:42:50.71871	\N
9	默认问题优先级方案	\N	9	t	t	f	f	2019-03-12 21:44:18.566135	\N
10	默认问题优先级方案	\N	10	t	t	f	f	2019-03-12 22:44:53.460728	\N
11	默认问题优先级方案	\N	11	t	t	f	f	2019-03-12 22:52:14.395588	\N
12	默认问题优先级方案	\N	12	t	t	f	f	2019-03-12 23:06:01.367716	\N
13	默认问题优先级方案	\N	13	t	t	f	f	2019-03-12 23:11:45.58636	\N
14	默认问题优先级方案	\N	14	t	t	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4614 (class 0 OID 46695)
-- Dependencies: 253
-- Data for Name: IsuPrioritySolutionItem; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuPrioritySolutionItem" ("priorityId", "solutionId", "orgId") FROM stdin;
4	1	1
3	1	1
2	1	1
1	1	1
8	2	2
7	2	2
6	2	2
5	2	2
16	4	4
15	4	4
14	4	4
13	4	4
20	5	5
19	5	5
18	5	5
17	5	5
24	6	6
23	6	6
22	6	6
21	6	6
25	5	5
29	7	7
28	7	7
27	7	7
26	7	7
33	8	8
32	8	8
31	8	8
30	8	8
37	9	9
36	9	9
35	9	9
34	9	9
41	10	10
40	10	10
39	10	10
38	10	10
45	11	11
44	11	11
43	11	11
42	11	11
49	12	12
48	12	12
47	12	12
46	12	12
53	13	13
52	13	13
51	13	13
50	13	13
57	14	14
56	14	14
55	14	14
54	14	14
\.


--
-- TOC entry 4617 (class 0 OID 46702)
-- Dependencies: 256
-- Data for Name: IsuQuery; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuQuery" (id, name, rule, "orderBy", descr, "useTime", "projectId", "userId", disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4619 (class 0 OID 46710)
-- Dependencies: 258
-- Data for Name: IsuResolution; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuResolution" (id, label, value, descr, ordr, "defaultVal", "buildIn", "orgId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	不是缺陷	not_defect	\N	20	\N	t	1	f	f	2019-02-17 14:22:02.073523	\N
2	修复	fixed	\N	10	t	t	1	f	f	2019-02-17 14:22:02.073523	\N
3	不是缺陷	not_defect	\N	20	\N	t	2	f	f	2019-02-18 21:51:44.67068	\N
4	修复	fixed	\N	10	t	t	2	f	f	2019-02-18 21:51:44.67068	\N
7	不是缺陷	not_defect	\N	20	\N	t	4	f	f	2019-02-19 11:38:28.481573	\N
8	修复	fixed	\N	10	t	t	4	f	f	2019-02-19 11:38:28.481573	\N
11	不是缺陷	not_defect	\N	20	\N	t	6	f	f	2019-02-23 13:16:23.128256	\N
12	修复	fixed	\N	10	t	t	6	f	f	2019-02-23 13:16:23.128256	\N
13	措施01	\N	\N	30	f	\N	5	f	f	2019-02-25 16:38:03.929515	2019-02-25 16:38:10.458772
9	不是缺陷	not_defect	\N	20	f	t	5	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:38:10.458772
10	修复	fixed	\N	10	t	t	5	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:38:10.458772
14	不是缺陷	not_defect	\N	20	\N	t	7	f	f	2019-03-11 00:56:22.895443	\N
15	修复	fixed	\N	10	t	t	7	f	f	2019-03-11 00:56:22.895443	\N
16	不是缺陷	not_defect	\N	20	\N	t	8	f	f	2019-03-12 21:42:50.71871	\N
17	修复	fixed	\N	10	t	t	8	f	f	2019-03-12 21:42:50.71871	\N
18	不是缺陷	not_defect	\N	20	\N	t	9	f	f	2019-03-12 21:44:18.566135	\N
19	修复	fixed	\N	10	t	t	9	f	f	2019-03-12 21:44:18.566135	\N
20	不是缺陷	not_defect	\N	20	\N	t	10	f	f	2019-03-12 22:44:53.460728	\N
21	修复	fixed	\N	10	t	t	10	f	f	2019-03-12 22:44:53.460728	\N
22	不是缺陷	not_defect	\N	20	\N	t	11	f	f	2019-03-12 22:52:14.395588	\N
23	修复	fixed	\N	10	t	t	11	f	f	2019-03-12 22:52:14.395588	\N
24	不是缺陷	not_defect	\N	20	\N	t	12	f	f	2019-03-12 23:06:01.367716	\N
25	修复	fixed	\N	10	t	t	12	f	f	2019-03-12 23:06:01.367716	\N
26	不是缺陷	not_defect	\N	20	\N	t	13	f	f	2019-03-12 23:11:45.58636	\N
27	修复	fixed	\N	10	t	t	13	f	f	2019-03-12 23:11:45.58636	\N
28	不是缺陷	not_defect	\N	20	\N	t	14	f	f	2019-03-12 23:27:21.303261	\N
29	修复	fixed	\N	10	t	t	14	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4620 (class 0 OID 46716)
-- Dependencies: 259
-- Data for Name: IsuResolutionDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuResolutionDefine" (id, label, value, "defaultVal", descr, ordr, disabled, deleted, "createTime", "updateTime") FROM stdin;
2	不是缺陷	not_defect	\N	\N	20	f	f	2018-11-23 15:26:22	\N
1	修复	fixed	t	\N	10	f	f	2018-11-23 15:25:52	\N
\.


--
-- TOC entry 4623 (class 0 OID 46726)
-- Dependencies: 262
-- Data for Name: IsuSeverity; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuSeverity" (id, label, value, descr, "defaultVal", "buildIn", ordr, "orgId", disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4624 (class 0 OID 46732)
-- Dependencies: 263
-- Data for Name: IsuSeverityDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuSeverityDefine" (id, label, value, descr, ordr, "defaultVal", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	阻塞	block	\N	1	f	f	f	2018-11-09 11:28:35	\N
2	紧急	critical	\N	2	f	f	f	2018-11-09 11:28:39	\N
3	重要	major	\N	3	f	f	f	2018-11-09 11:28:42	\N
4	一般	normal	\N	4	t	f	f	2018-11-09 11:42:21	\N
5	细微	minor	\N	5	f	f	f	2018-11-09 11:28:45	\N
\.


--
-- TOC entry 4626 (class 0 OID 46740)
-- Dependencies: 265
-- Data for Name: IsuSeveritySolution; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuSeveritySolution" (id, name, descr, "orgId", "defaultVal", "buildIn", disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4627 (class 0 OID 46746)
-- Dependencies: 266
-- Data for Name: IsuSeveritySolutionItem; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuSeveritySolutionItem" ("severityId", "solutionId") FROM stdin;
\.


--
-- TOC entry 4630 (class 0 OID 46753)
-- Dependencies: 269
-- Data for Name: IsuStatus; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuStatus" (id, label, value, descr, ordr, "orgId", "categoryId", "defaultVal", "finalVal", "buildIn", "startTime", "endTime", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	打开	open	\N	1	1	1	t	f	t	\N	\N	f	f	2019-02-17 14:22:02.073523	\N
2	解决	resolved	\N	2	1	2	f	f	t	\N	\N	f	f	2019-02-17 14:22:02.073523	\N
3	关闭	closed	\N	3	1	3	f	t	t	\N	\N	f	f	2019-02-17 14:22:02.073523	\N
4	重新打开	reopen	\N	4	1	1	f	f	t	\N	\N	f	f	2019-02-17 14:22:02.073523	\N
5	挂起	suspend	\N	5	1	3	f	t	t	\N	\N	f	f	2019-02-17 14:22:02.073523	\N
6	打开	open	\N	1	2	1	t	f	t	\N	\N	f	f	2019-02-18 21:51:44.67068	\N
7	解决	resolved	\N	2	2	2	f	f	t	\N	\N	f	f	2019-02-18 21:51:44.67068	\N
8	关闭	closed	\N	3	2	3	f	t	t	\N	\N	f	f	2019-02-18 21:51:44.67068	\N
9	重新打开	reopen	\N	4	2	1	f	f	t	\N	\N	f	f	2019-02-18 21:51:44.67068	\N
10	挂起	suspend	\N	5	2	3	f	t	t	\N	\N	f	f	2019-02-18 21:51:44.67068	\N
16	打开	open	\N	1	4	1	t	f	t	\N	\N	f	f	2019-02-19 11:38:28.481573	\N
17	解决	resolved	\N	2	4	2	f	f	t	\N	\N	f	f	2019-02-19 11:38:28.481573	\N
18	关闭	closed	\N	3	4	3	f	t	t	\N	\N	f	f	2019-02-19 11:38:28.481573	\N
19	重新打开	reopen	\N	4	4	1	f	f	t	\N	\N	f	f	2019-02-19 11:38:28.481573	\N
20	挂起	suspend	\N	5	4	3	f	t	t	\N	\N	f	f	2019-02-19 11:38:28.481573	\N
26	打开	open	\N	1	6	1	t	f	t	\N	\N	f	f	2019-02-23 13:16:23.128256	\N
27	解决	resolved	\N	2	6	2	f	f	t	\N	\N	f	f	2019-02-23 13:16:23.128256	\N
28	关闭	closed	\N	3	6	3	f	t	t	\N	\N	f	f	2019-02-23 13:16:23.128256	\N
29	重新打开	reopen	\N	4	6	1	f	f	t	\N	\N	f	f	2019-02-23 13:16:23.128256	\N
30	挂起	suspend	\N	5	6	3	f	t	t	\N	\N	f	f	2019-02-23 13:16:23.128256	\N
69	打开	open	\N	1	14	1	t	f	t	\N	\N	f	f	2019-03-12 23:27:21.303261	\N
70	解决	resolved	\N	2	14	2	f	f	t	\N	\N	f	f	2019-03-12 23:27:21.303261	\N
71	关闭	closed	\N	3	14	3	f	t	t	\N	\N	f	f	2019-03-12 23:27:21.303261	\N
72	重新打开	reopen	\N	4	14	1	f	f	t	\N	\N	f	f	2019-03-12 23:27:21.303261	\N
73	挂起	suspend	\N	5	14	3	f	t	t	\N	\N	f	f	2019-03-12 23:27:21.303261	\N
24	重新打开	reopen	\N	4	5	1	f	f	t	\N	\N	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:25:18.087559
23	关闭	closed	\N	3	5	3	f	t	t	\N	\N	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:25:18.087559
22	解决	resolved	\N	2	5	2	f	f	t	\N	\N	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:25:18.087559
21	打开	open	\N	1	5	1	t	f	t	\N	\N	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:25:18.087559
25	挂起	suspend	\N	5	5	3	f	t	t	\N	\N	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:25:25.066729
32	是的范德萨	\N	\N	25	5	2	f	f	\N	\N	\N	t	f	2019-02-25 16:25:45.587178	2019-02-25 16:25:50.630007
33	SDF	\N	\N	25	5	1	f	f	\N	\N	\N	t	f	2019-02-25 16:26:32.359469	2019-02-25 16:26:35.810633
31	状态01	\N	是打发	15	5	3	f	t	\N	\N	\N	f	f	2019-02-25 16:25:09.208569	2019-02-25 16:37:01.576481
34	打开	open	\N	1	7	1	t	f	t	\N	\N	f	f	2019-03-11 00:56:22.895443	\N
35	解决	resolved	\N	2	7	2	f	f	t	\N	\N	f	f	2019-03-11 00:56:22.895443	\N
36	关闭	closed	\N	3	7	3	f	t	t	\N	\N	f	f	2019-03-11 00:56:22.895443	\N
37	重新打开	reopen	\N	4	7	1	f	f	t	\N	\N	f	f	2019-03-11 00:56:22.895443	\N
38	挂起	suspend	\N	5	7	3	f	t	t	\N	\N	f	f	2019-03-11 00:56:22.895443	\N
39	打开	open	\N	1	8	1	t	f	t	\N	\N	f	f	2019-03-12 21:42:50.71871	\N
40	解决	resolved	\N	2	8	2	f	f	t	\N	\N	f	f	2019-03-12 21:42:50.71871	\N
41	关闭	closed	\N	3	8	3	f	t	t	\N	\N	f	f	2019-03-12 21:42:50.71871	\N
42	重新打开	reopen	\N	4	8	1	f	f	t	\N	\N	f	f	2019-03-12 21:42:50.71871	\N
43	挂起	suspend	\N	5	8	3	f	t	t	\N	\N	f	f	2019-03-12 21:42:50.71871	\N
44	打开	open	\N	1	9	1	t	f	t	\N	\N	f	f	2019-03-12 21:44:18.566135	\N
45	解决	resolved	\N	2	9	2	f	f	t	\N	\N	f	f	2019-03-12 21:44:18.566135	\N
46	关闭	closed	\N	3	9	3	f	t	t	\N	\N	f	f	2019-03-12 21:44:18.566135	\N
47	重新打开	reopen	\N	4	9	1	f	f	t	\N	\N	f	f	2019-03-12 21:44:18.566135	\N
48	挂起	suspend	\N	5	9	3	f	t	t	\N	\N	f	f	2019-03-12 21:44:18.566135	\N
49	打开	open	\N	1	10	1	t	f	t	\N	\N	f	f	2019-03-12 22:44:53.460728	\N
50	解决	resolved	\N	2	10	2	f	f	t	\N	\N	f	f	2019-03-12 22:44:53.460728	\N
51	关闭	closed	\N	3	10	3	f	t	t	\N	\N	f	f	2019-03-12 22:44:53.460728	\N
52	重新打开	reopen	\N	4	10	1	f	f	t	\N	\N	f	f	2019-03-12 22:44:53.460728	\N
53	挂起	suspend	\N	5	10	3	f	t	t	\N	\N	f	f	2019-03-12 22:44:53.460728	\N
54	打开	open	\N	1	11	1	t	f	t	\N	\N	f	f	2019-03-12 22:52:14.395588	\N
55	解决	resolved	\N	2	11	2	f	f	t	\N	\N	f	f	2019-03-12 22:52:14.395588	\N
56	关闭	closed	\N	3	11	3	f	t	t	\N	\N	f	f	2019-03-12 22:52:14.395588	\N
57	重新打开	reopen	\N	4	11	1	f	f	t	\N	\N	f	f	2019-03-12 22:52:14.395588	\N
58	挂起	suspend	\N	5	11	3	f	t	t	\N	\N	f	f	2019-03-12 22:52:14.395588	\N
59	打开	open	\N	1	12	1	t	f	t	\N	\N	f	f	2019-03-12 23:06:01.367716	\N
60	解决	resolved	\N	2	12	2	f	f	t	\N	\N	f	f	2019-03-12 23:06:01.367716	\N
61	关闭	closed	\N	3	12	3	f	t	t	\N	\N	f	f	2019-03-12 23:06:01.367716	\N
62	重新打开	reopen	\N	4	12	1	f	f	t	\N	\N	f	f	2019-03-12 23:06:01.367716	\N
63	挂起	suspend	\N	5	12	3	f	t	t	\N	\N	f	f	2019-03-12 23:06:01.367716	\N
64	打开	open	\N	1	13	1	t	f	t	\N	\N	f	f	2019-03-12 23:11:45.58636	\N
65	解决	resolved	\N	2	13	2	f	f	t	\N	\N	f	f	2019-03-12 23:11:45.58636	\N
66	关闭	closed	\N	3	13	3	f	t	t	\N	\N	f	f	2019-03-12 23:11:45.58636	\N
67	重新打开	reopen	\N	4	13	1	f	f	t	\N	\N	f	f	2019-03-12 23:11:45.58636	\N
68	挂起	suspend	\N	5	13	3	f	t	t	\N	\N	f	f	2019-03-12 23:11:45.58636	\N
\.


--
-- TOC entry 4631 (class 0 OID 46759)
-- Dependencies: 270
-- Data for Name: IsuStatusCategoryDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuStatusCategoryDefine" (id, label, value, ordr, disabled, deleted, "finalVal") FROM stdin;
1	待办	todo	1	f	f	f
2	处理中	in_progress	2	f	f	f
3	完成	completed	3	f	f	t
\.


--
-- TOC entry 4633 (class 0 OID 46767)
-- Dependencies: 272
-- Data for Name: IsuStatusDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuStatusDefine" (id, label, value, descr, "defaultVal", "finalVal", "categoryId", ordr, disabled, deleted, "createTime", "updateTime") FROM stdin;
1	打开	open	\N	t	f	1	1	f	f	2018-11-09 11:13:04	\N
2	解决	resolved	\N	f	f	2	2	f	f	2018-11-09 11:16:37	\N
3	关闭	closed	\N	f	t	3	3	f	f	2018-11-09 11:16:40	\N
4	重新打开	reopen	\N	f	f	1	4	f	f	2018-11-09 11:16:43	\N
5	挂起	suspend	\N	f	t	3	5	f	f	2018-11-09 11:16:46	\N
\.


--
-- TOC entry 4636 (class 0 OID 46777)
-- Dependencies: 275
-- Data for Name: IsuTag; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuTag" (id, name, "orgId", "userId", "createTime", "updateTime", disabled, deleted) FROM stdin;
2	苏州市工业园区	4	4	2019-02-19 19:08:59.322071	\N	f	f
3	测试	5	5	2019-02-25 16:04:48.318399	\N	f	f
4	哈哈	5	5	2019-02-25 16:05:05.405418	\N	f	f
\.


--
-- TOC entry 4637 (class 0 OID 46780)
-- Dependencies: 276
-- Data for Name: IsuTagRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuTagRelation" (id, "issueId", "tagId") FROM stdin;
1	4	2
3	8	3
4	8	4
\.


--
-- TOC entry 4640 (class 0 OID 46787)
-- Dependencies: 279
-- Data for Name: IsuType; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuType" (id, value, label, descr, ordr, "orgId", "defaultVal", "buildIn", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	defect	缺陷	\N	1	1	t	t	f	f	2019-02-17 14:22:02.073523	\N
2	todo	待办事项	\N	2	1	f	t	f	f	2019-02-17 14:22:02.073523	\N
3	defect	缺陷	\N	1	2	t	t	f	f	2019-02-18 21:51:44.67068	\N
4	todo	待办事项	\N	2	2	f	t	f	f	2019-02-18 21:51:44.67068	\N
7	defect	缺陷	\N	1	4	t	t	f	f	2019-02-19 11:38:28.481573	\N
8	todo	待办事项	\N	2	4	f	t	f	f	2019-02-19 11:38:28.481573	\N
11	defect	缺陷	\N	1	6	t	t	f	f	2019-02-23 13:16:23.128256	\N
12	todo	待办事项	\N	2	6	f	t	f	f	2019-02-23 13:16:23.128256	\N
13	\N	类型01	是打发	12	5	f	f	f	f	2019-02-25 16:36:22.066484	2019-02-25 16:37:37.541246
10	todo	待办事项	\N	2	5	f	t	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:37:37.541246
9	defect	缺陷	\N	1	5	t	t	f	f	2019-02-22 08:33:31.929668	2019-02-25 16:37:37.541246
14	defect	缺陷	\N	1	7	t	t	f	f	2019-03-11 00:56:22.895443	\N
15	todo	待办事项	\N	2	7	f	t	f	f	2019-03-11 00:56:22.895443	\N
16	defect	缺陷	\N	1	8	t	t	f	f	2019-03-12 21:42:50.71871	\N
17	todo	待办事项	\N	2	8	f	t	f	f	2019-03-12 21:42:50.71871	\N
18	defect	缺陷	\N	1	9	t	t	f	f	2019-03-12 21:44:18.566135	\N
19	todo	待办事项	\N	2	9	f	t	f	f	2019-03-12 21:44:18.566135	\N
20	defect	缺陷	\N	1	10	t	t	f	f	2019-03-12 22:44:53.460728	\N
21	todo	待办事项	\N	2	10	f	t	f	f	2019-03-12 22:44:53.460728	\N
22	defect	缺陷	\N	1	11	t	t	f	f	2019-03-12 22:52:14.395588	\N
23	todo	待办事项	\N	2	11	f	t	f	f	2019-03-12 22:52:14.395588	\N
24	defect	缺陷	\N	1	12	t	t	f	f	2019-03-12 23:06:01.367716	\N
25	todo	待办事项	\N	2	12	f	t	f	f	2019-03-12 23:06:01.367716	\N
26	defect	缺陷	\N	1	13	t	t	f	f	2019-03-12 23:11:45.58636	\N
27	todo	待办事项	\N	2	13	f	t	f	f	2019-03-12 23:11:45.58636	\N
28	defect	缺陷	\N	1	14	t	t	f	f	2019-03-12 23:27:21.303261	\N
29	todo	待办事项	\N	2	14	f	t	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4641 (class 0 OID 46793)
-- Dependencies: 280
-- Data for Name: IsuTypeDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuTypeDefine" (id, value, label, descr, ordr, "defaultVal", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	defect	缺陷	\N	1	t	f	f	2018-11-08 17:50:39	\N
2	todo	待办事项	\N	2	f	f	f	2018-11-08 17:54:24	\N
\.


--
-- TOC entry 4643 (class 0 OID 46801)
-- Dependencies: 282
-- Data for Name: IsuTypeSolution; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuTypeSolution" (id, name, descr, "orgId", "defaultVal", "buildIn", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	默认问题类型方案	\N	1	t	t	f	f	2019-02-17 14:22:02.073523	\N
2	默认问题类型方案	\N	2	t	t	f	f	2019-02-18 21:51:44.67068	\N
4	默认问题类型方案	\N	4	t	t	f	f	2019-02-19 11:38:28.481573	\N
5	默认问题类型方案	\N	5	t	t	f	f	2019-02-22 08:33:31.929668	\N
6	默认问题类型方案	\N	6	t	t	f	f	2019-02-23 13:16:23.128256	\N
7	默认问题类型方案	\N	7	t	t	f	f	2019-03-11 00:56:22.895443	\N
8	默认问题类型方案	\N	8	t	t	f	f	2019-03-12 21:42:50.71871	\N
9	默认问题类型方案	\N	9	t	t	f	f	2019-03-12 21:44:18.566135	\N
10	默认问题类型方案	\N	10	t	t	f	f	2019-03-12 22:44:53.460728	\N
11	默认问题类型方案	\N	11	t	t	f	f	2019-03-12 22:52:14.395588	\N
12	默认问题类型方案	\N	12	t	t	f	f	2019-03-12 23:06:01.367716	\N
13	默认问题类型方案	\N	13	t	t	f	f	2019-03-12 23:11:45.58636	\N
14	默认问题类型方案	\N	14	t	t	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4644 (class 0 OID 46807)
-- Dependencies: 283
-- Data for Name: IsuTypeSolutionItem; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuTypeSolutionItem" ("typeId", "solutionId", "orgId") FROM stdin;
2	1	1
1	1	1
4	2	2
3	2	2
8	4	4
7	4	4
10	5	5
9	5	5
12	6	6
11	6	6
13	5	5
15	7	7
14	7	7
17	8	8
16	8	8
19	9	9
18	9	9
21	10	10
20	10	10
23	11	11
22	11	11
25	12	12
24	12	12
27	13	13
26	13	13
29	14	14
28	14	14
\.


--
-- TOC entry 4647 (class 0 OID 46814)
-- Dependencies: 286
-- Data for Name: IsuWatch; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuWatch" (id, "userId", "issueId") FROM stdin;
1	5	8
2	6	8
\.


--
-- TOC entry 4649 (class 0 OID 46819)
-- Dependencies: 288
-- Data for Name: IsuWorkflow; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuWorkflow" (id, name, descr, "buildIn", "orgId", "createTime", "updateTime", disabled, deleted, "defaultVal") FROM stdin;
1	默认工作流	\N	t	1	2019-02-17 14:22:02.073523	\N	f	f	t
2	默认工作流	\N	t	2	2019-02-18 21:51:44.67068	\N	f	f	t
3	默认工作流	\N	t	4	2019-02-19 11:38:28.481573	\N	f	f	t
5	默认工作流	\N	t	6	2019-02-23 13:16:23.128256	\N	f	f	t
6	工作流01	\N	f	5	2019-02-25 17:22:57.974991	2019-02-25 18:16:48.826977	f	f	f
4	默认工作流	\N	t	5	2019-02-22 08:33:31.929668	2019-02-25 18:16:48.826977	f	f	t
7	默认工作流	\N	t	7	2019-03-11 00:56:22.895443	\N	f	f	t
8	默认工作流	\N	t	8	2019-03-12 21:42:50.71871	\N	f	f	t
9	默认工作流	\N	t	9	2019-03-12 21:44:18.566135	\N	f	f	t
10	默认工作流	\N	t	10	2019-03-12 22:44:53.460728	\N	f	f	t
11	默认工作流	\N	t	11	2019-03-12 22:52:14.395588	\N	f	f	t
12	默认工作流	\N	t	12	2019-03-12 23:06:01.367716	\N	f	f	t
13	默认工作流	\N	t	13	2019-03-12 23:11:45.58636	\N	f	f	t
14	默认工作流	\N	t	14	2019-03-12 23:27:21.303261	\N	f	f	t
\.


--
-- TOC entry 4650 (class 0 OID 46825)
-- Dependencies: 289
-- Data for Name: IsuWorkflowSolution; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuWorkflowSolution" (id, name, descr, "orgId", "defaultVal", "buildIn", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	默认工作流方案	\N	1	t	t	f	f	2019-02-17 14:22:02.073523	\N
2	默认工作流方案	\N	2	t	t	f	f	2019-02-18 21:51:44.67068	\N
3	默认工作流方案	\N	4	t	t	f	f	2019-02-19 11:38:28.481573	\N
4	默认工作流方案	\N	5	t	t	f	f	2019-02-22 08:33:31.929668	\N
5	默认工作流方案	\N	6	t	t	f	f	2019-02-23 13:16:23.128256	\N
6	工作流方案01	\N	5	f	f	f	f	2019-02-25 17:23:16.968294	\N
7	111	\N	5	f	f	f	f	2019-02-25 18:05:11.071463	\N
8	111	\N	5	f	f	f	f	2019-02-25 18:15:39.255885	\N
9	22	\N	5	f	f	f	f	2019-02-25 18:16:54.701015	\N
10	默认工作流方案	\N	7	t	t	f	f	2019-03-11 00:56:22.895443	\N
11	默认工作流方案	\N	8	t	t	f	f	2019-03-12 21:42:50.71871	\N
12	默认工作流方案	\N	9	t	t	f	f	2019-03-12 21:44:18.566135	\N
13	默认工作流方案	\N	10	t	t	f	f	2019-03-12 22:44:53.460728	\N
14	默认工作流方案	\N	11	t	t	f	f	2019-03-12 22:52:14.395588	\N
15	默认工作流方案	\N	12	t	t	f	f	2019-03-12 23:06:01.367716	\N
16	默认工作流方案	\N	13	t	t	f	f	2019-03-12 23:11:45.58636	\N
17	默认工作流方案	\N	14	t	t	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4651 (class 0 OID 46831)
-- Dependencies: 290
-- Data for Name: IsuWorkflowSolutionItem; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuWorkflowSolutionItem" (id, "typeId", "workflowId", "solutionId", "orgId") FROM stdin;
1	2	1	1	1
2	1	1	1	1
3	4	2	2	2
4	3	2	2	2
5	8	3	3	4
6	7	3	3	4
7	10	4	4	5
8	9	4	4	5
9	12	5	5	6
10	11	5	5	6
11	9	6	8	5
12	9	6	8	5
13	10	6	8	5
14	10	6	8	5
15	13	6	8	5
16	13	6	8	5
17	9	4	9	5
18	9	4	9	5
19	10	4	9	5
20	10	4	9	5
21	13	4	9	5
22	13	4	9	5
23	14	7	10	7
24	15	7	10	7
25	16	8	11	8
26	17	8	11	8
27	18	9	12	9
28	19	9	12	9
29	20	10	13	10
30	21	10	13	10
31	22	11	14	11
32	23	11	14	11
33	24	12	15	12
34	25	12	15	12
35	26	13	16	13
36	27	13	16	13
37	28	14	17	14
38	29	14	17	14
\.


--
-- TOC entry 4654 (class 0 OID 46838)
-- Dependencies: 293
-- Data for Name: IsuWorkflowStatusRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuWorkflowStatusRelation" (id, "workflowId", "statusId", "orgId") FROM stdin;
1	1	1	1
2	1	2	1
3	1	3	1
4	1	4	1
5	1	5	1
6	2	6	2
7	2	7	2
8	2	8	2
9	2	9	2
10	2	10	2
11	3	16	4
12	3	17	4
13	3	18	4
14	3	19	4
15	3	20	4
16	4	21	5
17	4	22	5
18	4	23	5
19	4	24	5
20	4	25	5
21	5	26	6
22	5	27	6
23	5	28	6
24	5	29	6
25	5	30	6
26	6	21	5
27	6	22	5
28	6	23	5
29	6	24	5
30	6	25	5
31	6	31	5
32	7	34	7
33	7	35	7
34	7	36	7
35	7	37	7
36	7	38	7
37	8	39	8
38	8	40	8
39	8	41	8
40	8	42	8
41	8	43	8
42	9	44	9
43	9	45	9
44	9	46	9
45	9	47	9
46	9	48	9
47	10	49	10
48	10	50	10
49	10	51	10
50	10	52	10
51	10	53	10
52	11	54	11
53	11	55	11
54	11	56	11
55	11	57	11
56	11	58	11
57	12	59	12
58	12	60	12
59	12	61	12
60	12	62	12
61	12	63	12
62	13	64	13
63	13	65	13
64	13	66	13
65	13	67	13
66	13	68	13
67	14	69	14
68	14	70	14
69	14	71	14
70	14	72	14
71	14	73	14
\.


--
-- TOC entry 4655 (class 0 OID 46841)
-- Dependencies: 294
-- Data for Name: IsuWorkflowStatusRelationDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuWorkflowStatusRelationDefine" (id, "workflowId", "statusId") FROM stdin;
21	\N	1
22	\N	2
23	\N	3
24	\N	4
25	\N	5
\.


--
-- TOC entry 4658 (class 0 OID 46848)
-- Dependencies: 297
-- Data for Name: IsuWorkflowTransition; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuWorkflowTransition" (id, name, "actionPageId", "srcStatusId", "dictStatusId", "orgId", ordr, "workflowId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	解决	2	1	2	1	\N	1	f	f	2019-02-17 14:22:02.073523	\N
2	挂起	\N	1	5	1	\N	1	f	f	2019-02-17 14:22:02.073523	\N
3	关闭	\N	1	3	1	\N	1	f	f	2019-02-17 14:22:02.073523	\N
4	关闭	\N	2	3	1	\N	1	f	f	2019-02-17 14:22:02.073523	\N
5	重新打开	\N	2	4	1	\N	1	f	f	2019-02-17 14:22:02.073523	\N
6	挂起	\N	2	5	1	\N	1	f	f	2019-02-17 14:22:02.073523	\N
7	解决	2	4	2	1	\N	1	f	f	2019-02-17 14:22:02.073523	\N
8	关闭	\N	4	3	1	\N	1	f	f	2019-02-17 14:22:02.073523	\N
9	重新打开	\N	5	4	1	\N	1	f	f	2019-02-17 14:22:02.073523	\N
10	解决	4	6	7	2	\N	2	f	f	2019-02-18 21:51:44.67068	\N
11	挂起	\N	6	10	2	\N	2	f	f	2019-02-18 21:51:44.67068	\N
12	关闭	\N	6	8	2	\N	2	f	f	2019-02-18 21:51:44.67068	\N
13	关闭	\N	7	8	2	\N	2	f	f	2019-02-18 21:51:44.67068	\N
14	重新打开	\N	7	9	2	\N	2	f	f	2019-02-18 21:51:44.67068	\N
15	挂起	\N	7	10	2	\N	2	f	f	2019-02-18 21:51:44.67068	\N
16	解决	4	9	7	2	\N	2	f	f	2019-02-18 21:51:44.67068	\N
17	关闭	\N	9	8	2	\N	2	f	f	2019-02-18 21:51:44.67068	\N
18	重新打开	\N	10	9	2	\N	2	f	f	2019-02-18 21:51:44.67068	\N
19	解决	6	16	17	4	\N	3	f	f	2019-02-19 11:38:28.481573	\N
20	挂起	\N	16	20	4	\N	3	f	f	2019-02-19 11:38:28.481573	\N
21	关闭	\N	16	18	4	\N	3	f	f	2019-02-19 11:38:28.481573	\N
22	关闭	\N	17	18	4	\N	3	f	f	2019-02-19 11:38:28.481573	\N
23	重新打开	\N	17	19	4	\N	3	f	f	2019-02-19 11:38:28.481573	\N
24	挂起	\N	17	20	4	\N	3	f	f	2019-02-19 11:38:28.481573	\N
25	解决	6	19	17	4	\N	3	f	f	2019-02-19 11:38:28.481573	\N
26	关闭	\N	19	18	4	\N	3	f	f	2019-02-19 11:38:28.481573	\N
27	重新打开	\N	20	19	4	\N	3	f	f	2019-02-19 11:38:28.481573	\N
29	挂起	\N	21	25	5	\N	4	f	f	2019-02-22 08:33:31.929668	\N
30	关闭	\N	21	23	5	\N	4	f	f	2019-02-22 08:33:31.929668	\N
31	关闭	\N	22	23	5	\N	4	f	f	2019-02-22 08:33:31.929668	\N
32	重新打开	\N	22	24	5	\N	4	f	f	2019-02-22 08:33:31.929668	\N
33	挂起	\N	22	25	5	\N	4	f	f	2019-02-22 08:33:31.929668	\N
34	解决	8	24	22	5	\N	4	f	f	2019-02-22 08:33:31.929668	\N
35	关闭	\N	24	23	5	\N	4	f	f	2019-02-22 08:33:31.929668	\N
36	重新打开	\N	25	24	5	\N	4	f	f	2019-02-22 08:33:31.929668	\N
37	解决	10	26	27	6	\N	5	f	f	2019-02-23 13:16:23.128256	\N
38	挂起	\N	26	30	6	\N	5	f	f	2019-02-23 13:16:23.128256	\N
39	关闭	\N	26	28	6	\N	5	f	f	2019-02-23 13:16:23.128256	\N
40	关闭	\N	27	28	6	\N	5	f	f	2019-02-23 13:16:23.128256	\N
41	重新打开	\N	27	29	6	\N	5	f	f	2019-02-23 13:16:23.128256	\N
42	挂起	\N	27	30	6	\N	5	f	f	2019-02-23 13:16:23.128256	\N
43	解决	10	29	27	6	\N	5	f	f	2019-02-23 13:16:23.128256	\N
44	关闭	\N	29	28	6	\N	5	f	f	2019-02-23 13:16:23.128256	\N
45	重新打开	\N	30	29	6	\N	5	f	f	2019-02-23 13:16:23.128256	\N
28	解决	8	21	22	5	\N	4	f	f	2019-02-22 08:33:31.929668	2019-02-25 10:39:54.58544
46	重新打开	\N	23	24	5	\N	4	f	f	2019-02-25 12:25:28.137122	\N
47	解决	12	34	35	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
48	挂起	\N	34	38	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
49	关闭	\N	34	36	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
50	关闭	\N	35	36	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
51	重新打开	\N	35	37	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
52	挂起	\N	35	38	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
53	解决	12	37	35	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
54	关闭	\N	37	36	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
55	重新打开	\N	38	37	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
56	重新打开	\N	36	37	7	\N	7	f	f	2019-03-11 00:56:22.895443	\N
57	解决	14	39	40	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
58	挂起	\N	39	43	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
59	关闭	\N	39	41	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
60	关闭	\N	40	41	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
61	重新打开	\N	40	42	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
62	挂起	\N	40	43	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
63	解决	14	42	40	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
64	关闭	\N	42	41	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
65	重新打开	\N	43	42	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
66	重新打开	\N	41	42	8	\N	8	f	f	2019-03-12 21:42:50.71871	\N
67	解决	16	44	45	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
68	挂起	\N	44	48	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
69	关闭	\N	44	46	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
70	关闭	\N	45	46	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
71	重新打开	\N	45	47	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
72	挂起	\N	45	48	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
73	解决	16	47	45	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
74	关闭	\N	47	46	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
75	重新打开	\N	48	47	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
76	重新打开	\N	46	47	9	\N	9	f	f	2019-03-12 21:44:18.566135	\N
77	解决	18	49	50	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
78	挂起	\N	49	53	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
79	关闭	\N	49	51	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
80	关闭	\N	50	51	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
81	重新打开	\N	50	52	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
82	挂起	\N	50	53	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
83	解决	18	52	50	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
84	关闭	\N	52	51	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
85	重新打开	\N	53	52	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
86	重新打开	\N	51	52	10	\N	10	f	f	2019-03-12 22:44:53.460728	\N
87	解决	20	54	55	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
88	挂起	\N	54	58	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
89	关闭	\N	54	56	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
90	关闭	\N	55	56	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
91	重新打开	\N	55	57	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
92	挂起	\N	55	58	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
93	解决	20	57	55	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
94	关闭	\N	57	56	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
95	重新打开	\N	58	57	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
96	重新打开	\N	56	57	11	\N	11	f	f	2019-03-12 22:52:14.395588	\N
97	解决	22	59	60	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
98	挂起	\N	59	63	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
99	关闭	\N	59	61	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
100	关闭	\N	60	61	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
101	重新打开	\N	60	62	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
102	挂起	\N	60	63	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
103	解决	22	62	60	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
104	关闭	\N	62	61	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
105	重新打开	\N	63	62	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
106	重新打开	\N	61	62	12	\N	12	f	f	2019-03-12 23:06:01.367716	\N
107	解决	24	64	65	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
108	挂起	\N	64	68	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
109	关闭	\N	64	66	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
110	关闭	\N	65	66	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
111	重新打开	\N	65	67	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
112	挂起	\N	65	68	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
113	解决	24	67	65	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
114	关闭	\N	67	66	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
115	重新打开	\N	68	67	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
116	重新打开	\N	66	67	13	\N	13	f	f	2019-03-12 23:11:45.58636	\N
117	解决	26	69	70	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
118	挂起	\N	69	73	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
119	关闭	\N	69	71	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
120	关闭	\N	70	71	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
121	重新打开	\N	70	72	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
122	挂起	\N	70	73	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
123	解决	26	72	70	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
124	关闭	\N	72	71	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
125	重新打开	\N	73	72	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
126	重新打开	\N	71	72	14	\N	14	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4659 (class 0 OID 46851)
-- Dependencies: 298
-- Data for Name: IsuWorkflowTransitionDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuWorkflowTransitionDefine" (id, name, "actionPageId", "srcStatusId", "dictStatusId", "isSolveIssue", disabled, deleted, "createTime", "updateTime") FROM stdin;
100	解决	\N	1	2	t	f	f	2018-11-15 16:54:22	\N
110	挂起	\N	1	5	\N	f	f	2018-11-15 16:54:32	\N
120	关闭	\N	1	3	\N	f	f	2018-11-15 17:09:13	\N
200	关闭	\N	2	3	\N	f	f	2018-11-15 16:54:26	\N
210	重新打开	\N	2	4	\N	f	f	2018-11-15 16:54:29	\N
220	挂起	\N	2	5	\N	f	f	2018-11-15 17:11:02	\N
300	解决	\N	4	2	t	f	f	2018-11-15 17:16:24	\N
310	关闭	\N	4	3	\N	f	f	2018-11-15 17:16:31	\N
320	重新打开	\N	5	4	\N	f	f	2018-11-15 17:16:34	\N
330	重新打开	\N	3	4	\N	f	f	2018-11-15 17:16:34	\N
\.


--
-- TOC entry 4661 (class 0 OID 46856)
-- Dependencies: 300
-- Data for Name: IsuWorkflowTransitionProjectRoleRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."IsuWorkflowTransitionProjectRoleRelation" (id, "workflowId", "workflowTransitionId", "projectRoleId", "orgId") FROM stdin;
1	1	9	4	1
2	1	8	4	1
3	1	7	4	1
4	1	6	4	1
5	1	5	4	1
6	1	4	4	1
7	1	3	4	1
8	1	2	4	1
9	1	1	4	1
10	1	9	3	1
11	1	8	3	1
12	1	7	3	1
13	1	6	3	1
14	1	5	3	1
15	1	4	3	1
16	1	3	3	1
17	1	2	3	1
18	1	1	3	1
19	1	9	2	1
20	1	8	2	1
21	1	7	2	1
22	1	6	2	1
23	1	5	2	1
24	1	4	2	1
25	1	3	2	1
26	1	2	2	1
27	1	1	2	1
28	1	9	1	1
29	1	8	1	1
30	1	7	1	1
31	1	6	1	1
32	1	5	1	1
33	1	4	1	1
34	1	3	1	1
35	1	2	1	1
36	1	1	1	1
37	2	18	8	2
38	2	17	8	2
39	2	16	8	2
40	2	15	8	2
41	2	14	8	2
42	2	13	8	2
43	2	12	8	2
44	2	11	8	2
45	2	10	8	2
46	2	18	7	2
47	2	17	7	2
48	2	16	7	2
49	2	15	7	2
50	2	14	7	2
51	2	13	7	2
52	2	12	7	2
53	2	11	7	2
54	2	10	7	2
55	2	18	6	2
56	2	17	6	2
57	2	16	6	2
58	2	15	6	2
59	2	14	6	2
60	2	13	6	2
61	2	12	6	2
62	2	11	6	2
63	2	10	6	2
64	2	18	5	2
65	2	17	5	2
66	2	16	5	2
67	2	15	5	2
68	2	14	5	2
69	2	13	5	2
70	2	12	5	2
71	2	11	5	2
72	2	10	5	2
73	3	27	16	4
74	3	26	16	4
75	3	25	16	4
76	3	24	16	4
77	3	23	16	4
78	3	22	16	4
79	3	21	16	4
80	3	20	16	4
81	3	19	16	4
82	3	27	15	4
83	3	26	15	4
84	3	25	15	4
85	3	24	15	4
86	3	23	15	4
87	3	22	15	4
88	3	21	15	4
89	3	20	15	4
90	3	19	15	4
91	3	27	14	4
92	3	26	14	4
93	3	25	14	4
94	3	24	14	4
95	3	23	14	4
96	3	22	14	4
97	3	21	14	4
98	3	20	14	4
99	3	19	14	4
100	3	27	13	4
101	3	26	13	4
102	3	25	13	4
103	3	24	13	4
104	3	23	13	4
105	3	22	13	4
106	3	21	13	4
107	3	20	13	4
108	3	19	13	4
109	4	36	20	5
110	4	35	20	5
111	4	34	20	5
112	4	33	20	5
113	4	32	20	5
114	4	31	20	5
115	4	30	20	5
116	4	29	20	5
118	4	36	19	5
119	4	35	19	5
120	4	34	19	5
121	4	33	19	5
122	4	32	19	5
123	4	31	19	5
124	4	30	19	5
125	4	29	19	5
127	4	36	18	5
128	4	35	18	5
129	4	34	18	5
130	4	33	18	5
131	4	32	18	5
132	4	31	18	5
133	4	30	18	5
134	4	29	18	5
136	4	36	17	5
137	4	35	17	5
138	4	34	17	5
139	4	33	17	5
140	4	32	17	5
141	4	31	17	5
142	4	30	17	5
143	4	29	17	5
145	5	45	24	6
146	5	44	24	6
147	5	43	24	6
148	5	42	24	6
149	5	41	24	6
150	5	40	24	6
151	5	39	24	6
152	5	38	24	6
153	5	37	24	6
154	5	45	23	6
155	5	44	23	6
156	5	43	23	6
157	5	42	23	6
158	5	41	23	6
159	5	40	23	6
160	5	39	23	6
161	5	38	23	6
162	5	37	23	6
163	5	45	22	6
164	5	44	22	6
165	5	43	22	6
166	5	42	22	6
167	5	41	22	6
168	5	40	22	6
169	5	39	22	6
170	5	38	22	6
171	5	37	22	6
172	5	45	21	6
173	5	44	21	6
174	5	43	21	6
175	5	42	21	6
176	5	41	21	6
177	5	40	21	6
178	5	39	21	6
179	5	38	21	6
180	5	37	21	6
181	4	28	17	5
182	4	28	18	5
183	4	28	19	5
184	4	28	20	5
185	4	46	17	5
186	4	46	18	5
187	4	46	19	5
188	4	46	20	5
189	7	56	28	7
190	7	55	28	7
191	7	54	28	7
192	7	53	28	7
193	7	52	28	7
194	7	51	28	7
195	7	50	28	7
196	7	49	28	7
197	7	48	28	7
198	7	47	28	7
199	7	56	27	7
200	7	55	27	7
201	7	54	27	7
202	7	53	27	7
203	7	52	27	7
204	7	51	27	7
205	7	50	27	7
206	7	49	27	7
207	7	48	27	7
208	7	47	27	7
209	7	56	26	7
210	7	55	26	7
211	7	54	26	7
212	7	53	26	7
213	7	52	26	7
214	7	51	26	7
215	7	50	26	7
216	7	49	26	7
217	7	48	26	7
218	7	47	26	7
219	7	56	25	7
220	7	55	25	7
221	7	54	25	7
222	7	53	25	7
223	7	52	25	7
224	7	51	25	7
225	7	50	25	7
226	7	49	25	7
227	7	48	25	7
228	7	47	25	7
229	8	66	32	8
230	8	65	32	8
231	8	64	32	8
232	8	63	32	8
233	8	62	32	8
234	8	61	32	8
235	8	60	32	8
236	8	59	32	8
237	8	58	32	8
238	8	57	32	8
239	8	66	31	8
240	8	65	31	8
241	8	64	31	8
242	8	63	31	8
243	8	62	31	8
244	8	61	31	8
245	8	60	31	8
246	8	59	31	8
247	8	58	31	8
248	8	57	31	8
249	8	66	30	8
250	8	65	30	8
251	8	64	30	8
252	8	63	30	8
253	8	62	30	8
254	8	61	30	8
255	8	60	30	8
256	8	59	30	8
257	8	58	30	8
258	8	57	30	8
259	8	66	29	8
260	8	65	29	8
261	8	64	29	8
262	8	63	29	8
263	8	62	29	8
264	8	61	29	8
265	8	60	29	8
266	8	59	29	8
267	8	58	29	8
268	8	57	29	8
269	9	76	36	9
270	9	75	36	9
271	9	74	36	9
272	9	73	36	9
273	9	72	36	9
274	9	71	36	9
275	9	70	36	9
276	9	69	36	9
277	9	68	36	9
278	9	67	36	9
279	9	76	35	9
280	9	75	35	9
281	9	74	35	9
282	9	73	35	9
283	9	72	35	9
284	9	71	35	9
285	9	70	35	9
286	9	69	35	9
287	9	68	35	9
288	9	67	35	9
289	9	76	34	9
290	9	75	34	9
291	9	74	34	9
292	9	73	34	9
293	9	72	34	9
294	9	71	34	9
295	9	70	34	9
296	9	69	34	9
297	9	68	34	9
298	9	67	34	9
299	9	76	33	9
300	9	75	33	9
301	9	74	33	9
302	9	73	33	9
303	9	72	33	9
304	9	71	33	9
305	9	70	33	9
306	9	69	33	9
307	9	68	33	9
308	9	67	33	9
309	10	86	40	10
310	10	85	40	10
311	10	84	40	10
312	10	83	40	10
313	10	82	40	10
314	10	81	40	10
315	10	80	40	10
316	10	79	40	10
317	10	78	40	10
318	10	77	40	10
319	10	86	39	10
320	10	85	39	10
321	10	84	39	10
322	10	83	39	10
323	10	82	39	10
324	10	81	39	10
325	10	80	39	10
326	10	79	39	10
327	10	78	39	10
328	10	77	39	10
329	10	86	38	10
330	10	85	38	10
331	10	84	38	10
332	10	83	38	10
333	10	82	38	10
334	10	81	38	10
335	10	80	38	10
336	10	79	38	10
337	10	78	38	10
338	10	77	38	10
339	10	86	37	10
340	10	85	37	10
341	10	84	37	10
342	10	83	37	10
343	10	82	37	10
344	10	81	37	10
345	10	80	37	10
346	10	79	37	10
347	10	78	37	10
348	10	77	37	10
349	11	96	44	11
350	11	95	44	11
351	11	94	44	11
352	11	93	44	11
353	11	92	44	11
354	11	91	44	11
355	11	90	44	11
356	11	89	44	11
357	11	88	44	11
358	11	87	44	11
359	11	96	43	11
360	11	95	43	11
361	11	94	43	11
362	11	93	43	11
363	11	92	43	11
364	11	91	43	11
365	11	90	43	11
366	11	89	43	11
367	11	88	43	11
368	11	87	43	11
369	11	96	42	11
370	11	95	42	11
371	11	94	42	11
372	11	93	42	11
373	11	92	42	11
374	11	91	42	11
375	11	90	42	11
376	11	89	42	11
377	11	88	42	11
378	11	87	42	11
379	11	96	41	11
380	11	95	41	11
381	11	94	41	11
382	11	93	41	11
383	11	92	41	11
384	11	91	41	11
385	11	90	41	11
386	11	89	41	11
387	11	88	41	11
388	11	87	41	11
389	12	106	48	12
390	12	105	48	12
391	12	104	48	12
392	12	103	48	12
393	12	102	48	12
394	12	101	48	12
395	12	100	48	12
396	12	99	48	12
397	12	98	48	12
398	12	97	48	12
399	12	106	47	12
400	12	105	47	12
401	12	104	47	12
402	12	103	47	12
403	12	102	47	12
404	12	101	47	12
405	12	100	47	12
406	12	99	47	12
407	12	98	47	12
408	12	97	47	12
409	12	106	46	12
410	12	105	46	12
411	12	104	46	12
412	12	103	46	12
413	12	102	46	12
414	12	101	46	12
415	12	100	46	12
416	12	99	46	12
417	12	98	46	12
418	12	97	46	12
419	12	106	45	12
420	12	105	45	12
421	12	104	45	12
422	12	103	45	12
423	12	102	45	12
424	12	101	45	12
425	12	100	45	12
426	12	99	45	12
427	12	98	45	12
428	12	97	45	12
429	13	107	52	13
430	13	108	52	13
431	13	109	52	13
432	13	110	52	13
433	13	111	52	13
434	13	112	52	13
435	13	113	52	13
436	13	114	52	13
437	13	115	52	13
438	13	116	52	13
439	13	107	51	13
440	13	108	51	13
441	13	109	51	13
442	13	110	51	13
443	13	111	51	13
444	13	112	51	13
445	13	113	51	13
446	13	114	51	13
447	13	115	51	13
448	13	116	51	13
449	13	107	50	13
450	13	108	50	13
451	13	109	50	13
452	13	110	50	13
453	13	111	50	13
454	13	112	50	13
455	13	113	50	13
456	13	114	50	13
457	13	115	50	13
458	13	116	50	13
459	13	107	49	13
460	13	108	49	13
461	13	109	49	13
462	13	110	49	13
463	13	111	49	13
464	13	112	49	13
465	13	113	49	13
466	13	114	49	13
467	13	115	49	13
468	13	116	49	13
469	14	117	56	14
470	14	118	56	14
471	14	119	56	14
472	14	120	56	14
473	14	121	56	14
474	14	122	56	14
475	14	123	56	14
476	14	124	56	14
477	14	125	56	14
478	14	126	56	14
479	14	117	55	14
480	14	118	55	14
481	14	119	55	14
482	14	120	55	14
483	14	121	55	14
484	14	122	55	14
485	14	123	55	14
486	14	124	55	14
487	14	125	55	14
488	14	126	55	14
489	14	117	54	14
490	14	118	54	14
491	14	119	54	14
492	14	120	54	14
493	14	121	54	14
494	14	122	54	14
495	14	123	54	14
496	14	124	54	14
497	14	125	54	14
498	14	126	54	14
499	14	117	53	14
500	14	118	53	14
501	14	119	53	14
502	14	120	53	14
503	14	121	53	14
504	14	122	53	14
505	14	123	53	14
506	14	124	53	14
507	14	125	53	14
508	14	126	53	14
\.


--
-- TOC entry 4665 (class 0 OID 46865)
-- Dependencies: 304
-- Data for Name: SysPrivilege; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."SysPrivilege" (id, code, name, descr, disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4667 (class 0 OID 46873)
-- Dependencies: 306
-- Data for Name: SysRole; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."SysRole" (id, name, descr, disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4668 (class 0 OID 46879)
-- Dependencies: 307
-- Data for Name: SysRolePrivilegeRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."SysRolePrivilegeRelation" ("roleId", "privilegeId") FROM stdin;
\.


--
-- TOC entry 4669 (class 0 OID 46882)
-- Dependencies: 308
-- Data for Name: SysRoleUserRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."SysRoleUserRelation" ("roleId", "userId") FROM stdin;
\.


--
-- TOC entry 4671 (class 0 OID 46887)
-- Dependencies: 310
-- Data for Name: SysUser; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."SysUser" (id, name, email, password, token, avatar, "verifyCode", "lastLoginTime", disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4673 (class 0 OID 46895)
-- Dependencies: 312
-- Data for Name: Test; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."Test" (id, name, "extProp") FROM stdin;
\.


--
-- TOC entry 4675 (class 0 OID 46903)
-- Dependencies: 314
-- Data for Name: TstAlert; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstAlert" (id, title, uri, type, status, "startTime", "endTime", "entityId", "entityName", "isRead", "isSent", "assigneeId", "userId", disabled, deleted, "createTime", "updateTime") FROM stdin;
5	是打发	\N	task	\N	\N	\N	7	\N	f	f	5	5	f	f	2019-02-23 11:33:57.013599	\N
7	是打发	\N	task	\N	\N	\N	8	\N	f	f	5	5	f	f	2019-02-23 11:44:52.306595	\N
16	任务01	\N	task	\N	\N	\N	6	\N	f	f	5	5	f	f	2019-02-23 16:21:52.260911	\N
\.


--
-- TOC entry 4677 (class 0 OID 46911)
-- Dependencies: 316
-- Data for Name: TstCase; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCase" (id, name, content, objective, "contentType", estimate, "pId", "isParent", ordr, "priorityId", "typeId", "reviewResult", "projectId", "createById", "updateById", disabled, deleted, "createTime", "updateTime", "extProp") FROM stdin;
1	测试用例	\N	\N	steps	10	\N	t	0	373	870	\N	2	1	\N	f	f	2019-02-17 14:22:02.073523	\N	\N
2	新特性	\N	\N	steps	10	1	t	0	373	870	\N	2	1	\N	f	f	2019-02-17 14:22:02.073523	\N	\N
3	新用例	\N	\N	steps	10	2	f	0	373	870	\N	2	1	\N	f	f	2019-02-17 14:22:02.073523	\N	\N
4	测试用例	\N	\N	steps	10	\N	t	0	376	877	\N	4	2	\N	f	f	2019-02-18 21:51:44.67068	\N	\N
5	新特性	\N	\N	steps	10	4	t	0	376	877	\N	4	2	\N	f	f	2019-02-18 21:51:44.67068	\N	\N
6	新用例	\N	\N	steps	10	5	f	0	376	877	\N	4	2	\N	f	f	2019-02-18 21:51:44.67068	\N	\N
10	测试用例	\N	\N	steps	10	\N	t	0	382	891	\N	8	4	\N	f	f	2019-02-19 11:38:28.481573	\N	\N
11	新特性	\N	\N	steps	10	10	t	0	382	891	\N	8	4	\N	f	f	2019-02-19 11:38:28.481573	\N	\N
12	新用例	\N	\N	steps	10	11	f	0	382	891	\N	8	4	4	f	f	2019-02-19 11:38:28.481573	2019-02-20 18:20:06.050665	{}
13	测试用例	\N	\N	steps	10	\N	t	0	385	898	\N	10	5	\N	f	f	2019-02-22 08:33:31.929668	\N	\N
14	新特性	\N	\N	steps	10	13	t	0	385	898	\N	10	5	\N	f	f	2019-02-22 08:33:31.929668	\N	\N
16	新特性	\N	\N	steps	10	14	t	1	385	898	\N	10	5	\N	f	f	2019-02-22 10:17:22.576221	\N	\N
18	新特性	\N	\N	steps	10	16	t	2	385	898	\N	10	5	\N	f	f	2019-02-22 10:19:17.596273	\N	\N
22	新特性	\N	\N	steps	10	14	t	2	385	898	\N	10	5	\N	f	f	2019-02-22 10:33:31.356267	\N	\N
23	新特性	\N	\N	steps	10	14	t	3	385	898	\N	10	5	\N	f	f	2019-02-22 10:33:43.008075	\N	\N
24	新特性	\N	\N	steps	10	14	t	4	385	898	\N	10	5	\N	f	f	2019-02-22 10:35:27.712233	\N	\N
20	新用例	\N	\N	steps	10	18	f	1	385	898	t	10	5	5	f	f	2019-02-22 10:19:38.463587	2019-02-22 10:52:24.115126	\N
21	新用例	\N	\N	steps	10	18	f	2	385	898	t	10	5	5	f	f	2019-02-22 10:19:40.696984	2019-02-22 10:52:25.118505	\N
19	新用例	\N	\N	steps	10	16	f	3	385	898	t	10	5	5	f	f	2019-02-22 10:19:24.660497	2019-02-22 10:52:26.116132	\N
17	富文本格式用例	<p>富文本格式SDF</p>	目的1234	richText	100	16	f	1	384	899	t	10	5	5	f	f	2019-02-22 10:19:14.108733	2019-02-23 12:56:34.476996	{}
25	新用例2	\N	\N	steps	10	24	f	1	385	898	f	10	5	5	f	f	2019-02-22 10:35:31.971659	2019-02-22 10:52:33.188182	\N
26	新用例	\N	\N	steps	10	24	f	2	385	898	f	10	5	5	f	f	2019-02-22 10:37:40.965236	2019-02-22 10:52:57.112139	{}
33	新用例	\N	\N	steps	10	32	f	1	385	899	\N	13	5	\N	f	f	2019-02-23 12:18:51.314119	2019-02-23 13:15:05.621163	\N
34	测试用例	\N	\N	steps	10	\N	t	0	\N	\N	\N	15	5	\N	f	f	2019-02-23 13:16:23.128256	\N	\N
35	新特性	\N	\N	steps	10	34	t	0	\N	\N	\N	15	5	\N	f	f	2019-02-23 13:16:23.128256	\N	\N
36	新用例	\N	\N	steps	10	35	f	0	388	905	\N	15	5	\N	f	f	2019-02-23 13:16:23.128256	\N	\N
42	新用例	\N	\N	steps	10	41	f	0	396	921	\N	19	12	\N	f	f	2019-03-12 21:42:50.71871	\N	\N
43	测试用例	\N	\N	steps	10	\N	t	0	\N	\N	\N	21	13	\N	f	f	2019-03-12 21:44:18.566135	\N	\N
44	新特性	\N	\N	steps	10	43	t	0	\N	\N	\N	21	13	\N	f	f	2019-03-12 21:44:18.566135	\N	\N
45	新用例	\N	\N	steps	10	44	f	0	399	928	\N	21	13	\N	f	f	2019-03-12 21:44:18.566135	\N	\N
46	测试用例	\N	\N	steps	10	\N	t	0	\N	\N	\N	23	14	\N	f	f	2019-03-12 22:44:53.460728	\N	\N
15	新用例2	<p>sdf</p>	目的123	steps	10	14	f	0	385	902	f	10	5	5	f	f	2019-02-22 08:33:31.929668	2019-02-23 17:24:35.799548	{"prop001": 33, "prop002": 31}
47	新特性	\N	\N	steps	10	46	t	0	\N	\N	\N	23	14	\N	f	f	2019-03-12 22:44:53.460728	\N	\N
37	测试用例	\N	\N	steps	10	\N	t	0	\N	\N	\N	17	11	\N	f	f	2019-03-11 00:56:22.895443	\N	\N
38	新特性	\N	\N	steps	10	37	t	0	\N	\N	\N	17	11	\N	f	f	2019-03-11 00:56:22.895443	\N	\N
48	新用例	\N	\N	steps	10	47	f	0	402	935	\N	23	14	\N	f	f	2019-03-12 22:44:53.460728	\N	\N
27	测试用例	\N	\N	steps	10	\N	t	1	385	898	\N	12	5	\N	f	f	2019-02-23 10:46:02.008954	\N	\N
29	新用例	\N	\N	steps	10	28	f	1	\N	\N	\N	12	5	\N	f	f	2019-02-23 10:46:02.008954	\N	\N
28	新特性	\N	\N	steps	10	27	t	1	385	898	\N	12	5	\N	f	f	2019-02-23 10:46:02.008954	\N	\N
30	新用例	\N	\N	steps	10	28	f	2	385	898	\N	12	5	\N	f	f	2019-02-23 12:00:50.910544	\N	\N
31	测试用例	\N	\N	steps	10	\N	t	1	385	898	\N	13	5	\N	f	f	2019-02-23 12:18:51.314119	\N	\N
32	新特性	\N	\N	steps	10	31	t	1	385	898	\N	13	5	\N	f	f	2019-02-23 12:18:51.314119	\N	\N
39	新用例	\N	\N	steps	10	38	f	0	393	914	\N	17	11	\N	f	f	2019-03-11 00:56:22.895443	\N	\N
40	测试用例	\N	\N	steps	10	\N	t	0	\N	\N	\N	19	12	\N	f	f	2019-03-12 21:42:50.71871	\N	\N
41	新特性	\N	\N	steps	10	40	t	0	\N	\N	\N	19	12	\N	f	f	2019-03-12 21:42:50.71871	\N	\N
49	测试用例	\N	\N	steps	10	\N	t	0	\N	\N	\N	25	15	\N	f	f	2019-03-12 22:52:14.395588	\N	\N
50	新特性	\N	\N	steps	10	49	t	0	\N	\N	\N	25	15	\N	f	f	2019-03-12 22:52:14.395588	\N	\N
51	新用例	\N	\N	steps	10	50	f	0	405	942	\N	25	15	\N	f	f	2019-03-12 22:52:14.395588	\N	\N
52	测试用例	\N	\N	steps	10	\N	t	0	\N	\N	\N	27	16	\N	f	f	2019-03-12 23:06:01.367716	\N	\N
53	新特性	\N	\N	steps	10	52	t	0	\N	\N	\N	27	16	\N	f	f	2019-03-12 23:06:01.367716	\N	\N
54	新用例	\N	\N	steps	10	53	f	0	408	949	\N	27	16	\N	f	f	2019-03-12 23:06:01.367716	\N	\N
55	测试用例	\N	\N	steps	10	\N	t	0	\N	\N	\N	29	17	\N	f	f	2019-03-12 23:11:45.58636	\N	\N
56	新特性	\N	\N	steps	10	55	t	0	\N	\N	\N	29	17	\N	f	f	2019-03-12 23:11:45.58636	\N	\N
57	新用例	\N	\N	steps	10	56	f	0	411	956	\N	29	17	\N	f	f	2019-03-12 23:11:45.58636	\N	\N
58	测试用例	\N	\N	steps	10	\N	t	0	\N	\N	\N	31	18	\N	f	f	2019-03-12 23:27:21.303261	\N	\N
59	新特性	\N	\N	steps	10	58	t	0	\N	\N	\N	31	18	\N	f	f	2019-03-12 23:27:21.303261	\N	\N
60	新用例	\N	\N	steps	10	59	f	0	414	963	\N	31	18	\N	f	f	2019-03-12 23:27:21.303261	\N	\N
62	新用例	\N	\N	steps	10	59	f	2	414	963	\N	31	18	\N	f	f	2019-03-15 07:57:16.553631	\N	\N
61	新用例	\N	\N	steps	10	59	f	1	414	963	\N	31	18	\N	f	t	2019-03-15 07:57:07.208629	\N	\N
\.


--
-- TOC entry 4678 (class 0 OID 46917)
-- Dependencies: 317
-- Data for Name: TstCaseAttachment; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseAttachment" (id, name, title, uri, descr, "docType", "caseId", "userId", "createTime", "updateTime", disabled, deleted) FROM stdin;
2	aliyun2.sql	\N	upload/data/20190222/eefd18e3-3fe6-4ca6-ade6-ed03521f82eb.sql	\N	\N	15	5	2019-02-22 11:27:45.078	\N	t	f
3	aliyun.sql	\N	upload/data/20190222/9f19e73c-b042-489d-ae56-bb4e24587622.sql	\N	\N	15	5	2019-02-22 11:32:51.954	\N	t	f
5	aliyun.sql	\N	upload/data/20190222/aliyun.sql-355898e30f6e4aa69ffdf2ee31852023.sql	\N	\N	15	5	2019-02-22 11:49:41.366	\N	f	f
6	樊登读书.txt	\N	upload/data/20190222/樊登读书.txt-86cfbec77cb447e0a132a81314316f99.txt	\N	\N	15	5	2019-02-22 11:51:29.653	\N	f	f
8	aliyun.sql	\N	upload/data/20190222/aliyun-4cd52b8b5e144de3a0332a0c760d543a.sql	\N	\N	15	5	2019-02-22 11:51:57.334	\N	t	f
7	Dockerfile	\N	upload/data/20190222/Dockerfile-bb176fb291174dad97fb941590d759a3.	\N	\N	15	5	2019-02-22 11:51:46.859	\N	t	f
1	aliyun.sql	\N	upload/data/20190222/b1590b0e-9f26-444d-a22a-cde9e931d0b8.sql	\N	\N	15	5	2019-02-22 11:27:38.06	\N	t	f
4	aliyun.sql	\N	upload/data/20190222/aliyun.sql-a7f726020c494265b8a2329c8bfd0b8a.sql	\N	\N	15	5	2019-02-22 11:44:56.277	\N	t	f
9	樊登读书 (1).txt	\N	upload/data/20190222/c6e091c1-b092-40f0-a34d-ba5a655e1dc9.txt	\N	\N	15	5	2019-02-22 15:50:46.627	\N	f	f
\.


--
-- TOC entry 4680 (class 0 OID 46925)
-- Dependencies: 319
-- Data for Name: TstCaseComments; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseComments" (id, summary, content, "caseId", "userId", "createTime", "updateTime", disabled, deleted) FROM stdin;
2	评审通过	\N	17	5	2019-02-22 10:52:23.352554	2019-02-22 10:52:23.352554	f	f
3	评审通过	\N	20	5	2019-02-22 10:52:24.115126	2019-02-22 10:52:24.115126	f	f
4	评审通过	\N	21	5	2019-02-22 10:52:25.118505	2019-02-22 10:52:25.118505	f	f
5	评审通过	\N	19	5	2019-02-22 10:52:26.116132	2019-02-22 10:52:26.116132	f	f
6	评审通过	\N	25	5	2019-02-22 10:52:27.172173	2019-02-22 10:52:27.172173	f	f
7	评审通过	\N	26	5	2019-02-22 10:52:27.960131	2019-02-22 10:52:27.960131	f	f
8	评审失败	\N	26	5	2019-02-22 10:52:30.176309	2019-02-22 10:52:30.176309	f	f
9	评审失败	\N	25	5	2019-02-22 10:52:33.188182	2019-02-22 10:52:33.188182	f	f
12	添加备注	的说法	26	5	2019-02-22 11:02:56.201424	2019-02-22 11:02:56.201424	f	f
11	添加备注	发放	26	5	2019-02-22 11:02:41.898198	2019-02-22 11:02:41.898198	t	f
10	评审失败	\N	26	5	2019-02-22 10:52:33.813193	2019-02-22 10:52:33.813193	t	f
13	评审失败	\N	15	5	2019-02-22 11:06:43.566223	2019-02-22 11:06:43.566223	t	f
14	添加备注	是打发	15	5	2019-02-22 11:26:57.944947	2019-02-22 11:26:57.944947	f	f
15	评审失败	\N	15	5	2019-02-22 11:27:03.258626	2019-02-22 11:27:03.258626	f	f
16	评审通过	\N	17	5	2019-02-22 11:27:10.533823	2019-02-22 11:27:10.533823	f	f
1	修改备注	sdf222333	15	5	2019-02-22 10:52:21.314567	2019-02-22 11:26:53.847968	t	f
17	评审通过	\N	15	5	2019-02-22 12:15:00.039277	2019-02-22 12:15:00.039277	f	f
18	评审通过	\N	15	5	2019-02-22 12:15:04.1551	2019-02-22 12:15:04.1551	f	f
19	评审通过	\N	15	5	2019-02-22 13:20:14.461617	2019-02-22 13:20:14.461617	f	f
20	评审失败	\N	15	5	2019-02-22 13:20:25.282674	2019-02-22 13:20:25.282674	f	f
\.


--
-- TOC entry 4684 (class 0 OID 46941)
-- Dependencies: 323
-- Data for Name: TstCaseExeStatus; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseExeStatus" (id, value, label, descr, ordr, "finalVal", "buildIn", disabled, deleted, "orgId", "createTime", "updateTime") FROM stdin;
519	untest	未执行	\N	10	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
520	pass	成功	\N	20	t	t	f	f	1	2019-02-17 14:22:02.073523	\N
521	fail	失败	\N	30	t	t	f	f	1	2019-02-17 14:22:02.073523	\N
522	block	阻塞	\N	40	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
523	untest	未执行	\N	10	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
524	pass	成功	\N	20	t	t	f	f	2	2019-02-18 21:51:44.67068	\N
525	fail	失败	\N	30	t	t	f	f	2	2019-02-18 21:51:44.67068	\N
526	block	阻塞	\N	40	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
531	untest	未执行	\N	10	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
532	pass	成功	\N	20	t	t	f	f	4	2019-02-19 11:38:28.481573	\N
533	fail	失败	\N	30	t	t	f	f	4	2019-02-19 11:38:28.481573	\N
534	block	阻塞	\N	40	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
535	untest	未执行	\N	10	f	t	f	f	5	2019-02-22 08:33:31.929668	\N
536	pass	成功	\N	20	t	t	f	f	5	2019-02-22 08:33:31.929668	\N
537	fail	失败	\N	30	t	t	f	f	5	2019-02-22 08:33:31.929668	\N
538	block	阻塞	\N	40	f	t	f	f	5	2019-02-22 08:33:31.929668	\N
539	untest	未执行	\N	10	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
540	pass	成功	\N	20	t	t	f	f	6	2019-02-23 13:16:23.128256	\N
541	fail	失败	\N	30	t	t	f	f	6	2019-02-23 13:16:23.128256	\N
542	block	阻塞	\N	40	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
543	untest	未执行	\N	10	f	t	f	f	7	2019-03-11 00:56:22.895443	\N
544	pass	成功	\N	20	t	t	f	f	7	2019-03-11 00:56:22.895443	\N
545	fail	失败	\N	30	t	t	f	f	7	2019-03-11 00:56:22.895443	\N
546	block	阻塞	\N	40	f	t	f	f	7	2019-03-11 00:56:22.895443	\N
547	untest	未执行	\N	10	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
548	pass	成功	\N	20	t	t	f	f	8	2019-03-12 21:42:50.71871	\N
549	fail	失败	\N	30	t	t	f	f	8	2019-03-12 21:42:50.71871	\N
550	block	阻塞	\N	40	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
551	untest	未执行	\N	10	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
552	pass	成功	\N	20	t	t	f	f	9	2019-03-12 21:44:18.566135	\N
553	fail	失败	\N	30	t	t	f	f	9	2019-03-12 21:44:18.566135	\N
554	block	阻塞	\N	40	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
555	untest	未执行	\N	10	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
556	pass	成功	\N	20	t	t	f	f	10	2019-03-12 22:44:53.460728	\N
557	fail	失败	\N	30	t	t	f	f	10	2019-03-12 22:44:53.460728	\N
558	block	阻塞	\N	40	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
559	untest	未执行	\N	10	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
560	pass	成功	\N	20	t	t	f	f	11	2019-03-12 22:52:14.395588	\N
561	fail	失败	\N	30	t	t	f	f	11	2019-03-12 22:52:14.395588	\N
562	block	阻塞	\N	40	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
563	untest	未执行	\N	10	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
564	pass	成功	\N	20	t	t	f	f	12	2019-03-12 23:06:01.367716	\N
565	fail	失败	\N	30	t	t	f	f	12	2019-03-12 23:06:01.367716	\N
566	block	阻塞	\N	40	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
567	untest	未执行	\N	10	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
568	pass	成功	\N	20	t	t	f	f	13	2019-03-12 23:11:45.58636	\N
569	fail	失败	\N	30	t	t	f	f	13	2019-03-12 23:11:45.58636	\N
570	block	阻塞	\N	40	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
571	untest	未执行	\N	10	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
572	pass	成功	\N	20	t	t	f	f	14	2019-03-12 23:27:21.303261	\N
573	fail	失败	\N	30	t	t	f	f	14	2019-03-12 23:27:21.303261	\N
574	block	阻塞	\N	40	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4682 (class 0 OID 46933)
-- Dependencies: 321
-- Data for Name: TstCaseExeStatusDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseExeStatusDefine" (id, value, label, descr, ordr, "finalVal", disabled, deleted) FROM stdin;
1	untest	未执行	\N	10	f	f	f
2	pass	成功	\N	20	t	f	f
3	fail	失败	\N	30	t	f	f
4	block	阻塞	\N	40	f	f	f
\.


--
-- TOC entry 4685 (class 0 OID 46948)
-- Dependencies: 324
-- Data for Name: TstCaseHistory; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseHistory" (id, title, descr, "caseId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	用户Aaron Chen更新	\N	12	f	f	2019-02-20 18:20:06.050665	\N
2	用户Aaron Chen创建	\N	16	f	f	2019-02-22 10:17:22.576221	\N
3	用户Aaron Chen创建	\N	17	f	f	2019-02-22 10:19:14.108733	\N
4	用户Aaron Chen创建	\N	18	f	f	2019-02-22 10:19:17.596273	\N
5	用户Aaron Chen创建	\N	19	f	f	2019-02-22 10:19:24.660497	\N
6	用户Aaron Chen创建	\N	20	f	f	2019-02-22 10:19:38.463587	\N
7	用户Aaron Chen创建	\N	21	f	f	2019-02-22 10:19:40.696984	\N
8	用户Aaron Chen创建	\N	22	f	f	2019-02-22 10:33:31.356267	\N
9	用户Aaron Chen创建	\N	23	f	f	2019-02-22 10:33:43.008075	\N
10	用户Aaron Chen创建	\N	24	f	f	2019-02-22 10:35:27.712233	\N
11	用户Aaron Chen创建	\N	25	f	f	2019-02-22 10:35:31.971659	\N
12	用户Aaron Chen改名	\N	25	f	f	2019-02-22 10:35:35.972064	\N
13	用户Aaron Chen创建	\N	26	f	f	2019-02-22 10:37:40.965236	\N
14	用户Aaron Chen更新	\N	26	f	f	2019-02-22 10:52:57.112139	\N
15	用户Aaron Chen上传附件字段 aliyun.sql	\N	15	f	f	2019-02-22 11:27:38.084241	\N
16	用户Aaron Chen上传附件字段 aliyun2.sql	\N	15	f	f	2019-02-22 11:27:45.079545	\N
17	用户Aaron Chen上传附件字段 aliyun.sql	\N	15	f	f	2019-02-22 11:32:51.964013	\N
18	用户Aaron Chen删除附件字段 aliyun2.sql	\N	15	f	f	2019-02-22 11:32:54.9075	\N
19	用户Aaron Chen删除附件字段 aliyun.sql	\N	15	f	f	2019-02-22 11:34:48.336776	\N
20	用户Aaron Chen上传附件字段 aliyun.sql	\N	15	f	f	2019-02-22 11:44:56.280979	\N
21	用户Aaron Chen上传附件字段 aliyun.sql	\N	15	f	f	2019-02-22 11:49:41.367695	\N
22	用户Aaron Chen上传附件字段 樊登读书.txt	\N	15	f	f	2019-02-22 11:51:29.663043	\N
23	用户Aaron Chen上传附件字段 Dockerfile	\N	15	f	f	2019-02-22 11:51:46.89287	\N
24	用户Aaron Chen上传附件字段 aliyun.sql	\N	15	f	f	2019-02-22 11:51:57.335498	\N
25	用户Aaron Chen删除附件字段 aliyun.sql	\N	15	f	f	2019-02-22 11:52:46.204616	\N
26	用户Aaron Chen删除附件字段 Dockerfile	\N	15	f	f	2019-02-22 11:52:48.476897	\N
27	用户Aaron Chen删除附件字段 aliyun.sql	\N	15	f	f	2019-02-22 11:52:51.177766	\N
28	用户Aaron Chen删除附件字段 aliyun.sql	\N	15	f	f	2019-02-22 11:52:52.994995	\N
29	用户Aaron Chen更新	\N	15	f	f	2019-02-22 11:53:04.853287	\N
30	用户Aaron Chen更新	\N	15	f	f	2019-02-22 11:55:21.067192	\N
31	用户Aaron Chen更新	\N	15	f	f	2019-02-22 11:56:08.532895	\N
32	用户Aaron Chen更新	\N	15	f	f	2019-02-22 11:59:27.024687	\N
33	用户Aaron Chen更新	\N	15	f	f	2019-02-22 12:00:50.582699	\N
34	用户Aaron Chen更新	\N	15	f	f	2019-02-22 12:05:22.316483	\N
35	用户Aaron Chen更新	\N	15	f	f	2019-02-22 12:06:04.993827	\N
36	用户Aaron Chen更新	\N	15	f	f	2019-02-22 12:06:54.427106	\N
37	用户Aaron Chen更新	\N	15	f	f	2019-02-22 12:08:30.561632	\N
38	用户Aaron Chen更新	\N	15	f	f	2019-02-22 12:14:27.470894	\N
39	用户Aaron Chen更新	\N	17	f	f	2019-02-22 12:20:44.132926	\N
40	用户Aaron Chen更新	\N	17	f	f	2019-02-22 12:20:55.097598	\N
41	用户Aaron Chen更新	\N	15	f	f	2019-02-22 12:27:20.524576	\N
42	用户Aaron Chen更新	\N	15	f	f	2019-02-22 12:29:35.114632	\N
43	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:07:19.644056	\N
44	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:07:22.435757	\N
45	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:07:39.623773	\N
46	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:07:50.161083	\N
47	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:13:19.952464	\N
48	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:13:26.407162	\N
49	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:14:12.178325	\N
50	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:15:06.600248	\N
51	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:15:23.0548	\N
52	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:15:39.537937	\N
53	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:15:50.954035	\N
54	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:19:44.343014	\N
55	用户Aaron Chen更新字段 内容类型	\N	15	f	f	2019-02-22 13:21:08.046665	\N
56	用户Aaron Chen更新	\N	15	f	f	2019-02-22 13:21:11.698504	\N
57	用户Aaron Chen更新字段 内容类型	\N	15	f	f	2019-02-22 13:21:24.410514	\N
58	用户Aaron Chen更新字段 内容类型	\N	15	f	f	2019-02-22 13:21:30.173346	\N
59	用户Aaron Chen更新字段 内容类型	\N	15	f	f	2019-02-22 13:21:34.213603	\N
60	用户Aaron Chen上传附件字段 樊登读书 (1).txt	\N	15	f	f	2019-02-22 15:50:46.66733	\N
61	用户Aaron Chen更新字段 内容类型	\N	17	f	f	2019-02-22 21:47:28.107446	\N
62	用户Aaron Chen更新	\N	17	f	f	2019-02-22 21:47:45.782851	\N
63	用户Aaron Chen更新	\N	17	f	f	2019-02-22 21:48:07.507091	\N
64	用户Aaron Chen创建	\N	29	f	f	2019-02-23 10:46:02.008954	\N
65	用户Aaron Chen创建	\N	30	f	f	2019-02-23 12:00:50.910544	\N
66	用户Aaron Chen创建	\N	33	f	f	2019-02-23 12:18:51.314119	\N
67	用户Aaron Chen更新	\N	15	f	f	2019-02-23 12:50:52.55069	\N
68	用户Aaron Chen更新	\N	17	f	f	2019-02-23 12:56:34.476996	\N
69	用户Aaron Chen更新	\N	33	f	f	2019-02-23 13:15:05.621163	\N
70	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:09:08.368262	\N
71	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:09:41.227422	\N
72	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:09:54.75326	\N
73	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:11:30.735478	\N
74	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:12:02.857926	\N
75	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:12:12.030332	\N
76	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:12:54.279058	\N
77	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:13:32.234085	\N
78	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:13:42.044597	\N
79	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:16:17.713308	\N
80	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:19:18.372484	\N
81	用户Aaron Chen改名	\N	15	f	f	2019-02-23 17:19:22.727572	\N
82	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:23:30.133124	\N
83	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:24:19.280994	\N
84	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:24:25.410646	\N
85	用户Aaron Chen更新	\N	15	f	f	2019-02-23 17:24:35.799548	\N
86	用户Aaron Chen2创建	\N	61	f	f	2019-03-15 07:57:07.311234	\N
87	用户Aaron Chen2创建	\N	62	f	f	2019-03-15 07:57:16.571297	\N
88	用户Aaron Chen2删除	\N	61	f	f	2019-03-15 08:18:32.189429	\N
\.


--
-- TOC entry 4687 (class 0 OID 46956)
-- Dependencies: 326
-- Data for Name: TstCaseInSuite; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseInSuite" (id, "caseId", "isParent", ordr, "pId", "projectId", "suiteId", deleted, disabled, "createBy", "createTime", "updateTime") FROM stdin;
15	13	t	0	\N	10	2	f	f	\N	2019-02-23 10:44:35.198816	\N
16	14	t	0	13	10	2	f	f	\N	2019-02-23 10:44:35.198816	\N
17	16	t	1	14	10	2	f	f	\N	2019-02-23 10:44:35.198816	\N
18	17	f	1	16	10	2	f	f	\N	2019-02-23 10:44:35.198816	\N
42	13	t	0	\N	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
43	14	t	0	13	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
44	15	f	0	14	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
45	16	t	1	14	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
46	17	f	1	16	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
47	18	t	2	16	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
48	20	f	1	18	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
49	21	f	2	18	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
50	19	f	3	16	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
51	22	t	2	14	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
52	23	t	3	14	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
53	24	t	4	14	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
54	25	f	1	24	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
55	26	f	2	24	10	1	f	f	\N	2019-02-23 12:50:03.695033	\N
\.


--
-- TOC entry 4689 (class 0 OID 46961)
-- Dependencies: 328
-- Data for Name: TstCaseInTask; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseInTask" (id, "caseId", "isParent", "pId", ordr, "exeBy", "exeTime", status, result, "planId", "projectId", "taskId", disabled, deleted, "createBy", "createTime", "updateTime") FROM stdin;
1	1	t	\N	1	\N	\N	untest	\N	1	2	1	f	f	1	2019-02-17 14:22:02.073523	\N
2	2	t	1	1	\N	\N	untest	\N	1	2	1	f	f	1	2019-02-17 14:22:02.073523	\N
3	3	f	2	1	\N	\N	untest	\N	1	2	1	f	f	1	2019-02-17 14:22:02.073523	\N
4	4	t	\N	1	\N	\N	untest	\N	2	4	2	f	f	2	2019-02-18 21:51:44.67068	\N
5	5	t	4	1	\N	\N	untest	\N	2	4	2	f	f	2	2019-02-18 21:51:44.67068	\N
6	6	f	5	1	\N	\N	untest	\N	2	4	2	f	f	2	2019-02-18 21:51:44.67068	\N
10	10	t	\N	1	\N	\N	untest	\N	4	8	4	f	f	4	2019-02-19 11:38:28.481573	\N
11	11	t	10	1	\N	\N	untest	\N	4	8	4	f	f	4	2019-02-19 11:38:28.481573	\N
12	12	f	11	1	\N	\N	untest	\N	4	8	4	f	f	4	2019-02-19 11:38:28.481573	\N
13	13	t	\N	1	\N	\N	untest	\N	5	10	5	f	f	5	2019-02-22 08:33:31.929668	\N
14	14	t	13	1	\N	\N	untest	\N	5	10	5	f	f	5	2019-02-22 08:33:31.929668	\N
15	15	f	14	1	\N	\N	untest	\N	5	10	5	f	f	5	2019-02-22 08:33:31.929668	\N
40	27	t	\N	1	\N	\N	untest	\N	7	10	7	f	f	\N	2019-02-23 11:39:38.628367	\N
41	28	t	27	1	\N	\N	untest	\N	7	10	7	f	f	\N	2019-02-23 11:39:38.628367	\N
42	29	f	28	1	\N	\N	untest	\N	7	10	7	f	f	\N	2019-02-23 11:39:38.628367	\N
50	27	t	\N	1	\N	\N	untest	\N	8	10	8	f	f	\N	2019-02-23 11:44:52.306595	\N
51	28	t	27	1	\N	\N	untest	\N	8	10	8	f	f	\N	2019-02-23 11:44:52.306595	\N
52	29	f	28	1	\N	\N	untest	\N	8	10	8	f	f	\N	2019-02-23 11:44:52.306595	\N
71	15	f	14	0	5	2019-02-23 12:50:23.510283	fail	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
23	27	t	\N	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 11:31:15.349667	\N
24	28	t	27	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 11:31:15.349667	\N
25	29	f	28	1	5	2019-02-23 11:59:00.006548	fail	\N	6	10	6	f	t	\N	2019-02-23 11:31:15.349667	\N
53	27	t	\N	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:27:18.614152	\N
54	28	t	27	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:27:18.614152	\N
55	29	f	28	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:27:18.614152	\N
56	30	f	28	2	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:27:18.614152	\N
57	27	t	\N	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:44:33.975449	\N
58	28	t	27	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:44:33.975449	\N
59	29	f	28	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:44:33.975449	\N
66	27	t	\N	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:48:33.861344	\N
67	28	t	27	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:48:33.861344	\N
68	29	f	28	1	5	2019-02-23 12:48:59.847042	block	\N	6	10	6	f	t	\N	2019-02-23 12:48:33.861344	\N
69	13	t	\N	0	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
70	14	t	13	0	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
72	16	t	14	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
73	17	f	16	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
74	18	t	16	2	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
75	20	f	18	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
76	21	f	18	2	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
77	19	f	16	3	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
78	22	t	14	2	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
86	34	t	\N	1	\N	\N	untest	\N	9	15	9	f	f	5	2019-02-23 13:16:23.128256	\N
87	35	t	34	1	\N	\N	untest	\N	9	15	9	f	f	5	2019-02-23 13:16:23.128256	\N
88	36	f	35	1	\N	\N	untest	\N	9	15	9	f	f	5	2019-02-23 13:16:23.128256	\N
79	23	t	14	3	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
80	24	t	14	4	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
81	25	f	24	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
82	26	f	24	2	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:50:13.638419	\N
83	31	t	\N	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:59:51.954835	\N
84	32	t	31	1	\N	\N	untest	\N	6	10	6	f	t	\N	2019-02-23 12:59:51.954835	\N
85	33	f	32	1	5	2019-02-23 13:44:01.796698	fail	\N	6	10	6	f	t	\N	2019-02-23 12:59:51.954835	\N
89	13	t	\N	0	\N	\N	untest	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
90	14	t	13	0	\N	\N	untest	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
92	16	t	14	1	\N	\N	untest	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
94	18	t	16	2	\N	\N	untest	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
98	22	t	14	2	\N	\N	untest	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
99	23	t	14	3	\N	\N	untest	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
100	24	t	14	4	\N	\N	untest	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
91	15	f	14	0	5	2019-02-23 16:22:03.345597	fail	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
93	17	f	16	1	5	2019-02-23 16:26:25.222351	block	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
95	20	f	18	1	5	2019-02-23 16:26:27.31754	block	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
96	21	f	18	2	5	2019-02-23 16:26:28.264029	pass	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
97	19	f	16	3	5	2019-02-23 16:26:28.530164	pass	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
101	25	f	24	1	5	2019-02-23 16:26:28.721485	pass	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
102	26	f	24	2	5	2019-02-23 16:26:29.586554	pass	\N	6	10	6	f	f	\N	2019-02-23 16:21:41.984651	\N
103	37	t	\N	1	\N	\N	untest	\N	10	17	10	f	f	11	2019-03-11 00:56:22.895443	\N
104	38	t	37	1	\N	\N	untest	\N	10	17	10	f	f	11	2019-03-11 00:56:22.895443	\N
105	39	f	38	1	\N	\N	untest	\N	10	17	10	f	f	11	2019-03-11 00:56:22.895443	\N
106	40	t	\N	1	\N	\N	untest	\N	11	19	11	f	f	12	2019-03-12 21:42:50.71871	\N
107	41	t	40	1	\N	\N	untest	\N	11	19	11	f	f	12	2019-03-12 21:42:50.71871	\N
108	42	f	41	1	\N	\N	untest	\N	11	19	11	f	f	12	2019-03-12 21:42:50.71871	\N
109	43	t	\N	1	\N	\N	untest	\N	12	21	12	f	f	13	2019-03-12 21:44:18.566135	\N
110	44	t	43	1	\N	\N	untest	\N	12	21	12	f	f	13	2019-03-12 21:44:18.566135	\N
111	45	f	44	1	\N	\N	untest	\N	12	21	12	f	f	13	2019-03-12 21:44:18.566135	\N
112	46	t	\N	1	\N	\N	untest	\N	13	23	13	f	f	14	2019-03-12 22:44:53.460728	\N
113	47	t	46	1	\N	\N	untest	\N	13	23	13	f	f	14	2019-03-12 22:44:53.460728	\N
114	48	f	47	1	\N	\N	untest	\N	13	23	13	f	f	14	2019-03-12 22:44:53.460728	\N
115	49	t	\N	1	\N	\N	untest	\N	14	25	14	f	f	15	2019-03-12 22:52:14.395588	\N
116	50	t	49	1	\N	\N	untest	\N	14	25	14	f	f	15	2019-03-12 22:52:14.395588	\N
117	51	f	50	1	\N	\N	untest	\N	14	25	14	f	f	15	2019-03-12 22:52:14.395588	\N
118	52	t	\N	1	\N	\N	untest	\N	15	27	15	f	f	16	2019-03-12 23:06:01.367716	\N
119	53	t	52	1	\N	\N	untest	\N	15	27	15	f	f	16	2019-03-12 23:06:01.367716	\N
120	54	f	53	1	\N	\N	untest	\N	15	27	15	f	f	16	2019-03-12 23:06:01.367716	\N
121	55	t	\N	1	\N	\N	untest	\N	16	29	16	f	f	17	2019-03-12 23:11:45.58636	\N
122	56	t	55	1	\N	\N	untest	\N	16	29	16	f	f	17	2019-03-12 23:11:45.58636	\N
123	57	f	56	1	\N	\N	untest	\N	16	29	16	f	f	17	2019-03-12 23:11:45.58636	\N
124	58	t	\N	1	\N	\N	untest	\N	17	31	17	f	f	18	2019-03-12 23:27:21.303261	\N
125	59	t	58	1	\N	\N	untest	\N	17	31	17	f	f	18	2019-03-12 23:27:21.303261	\N
126	60	f	59	1	\N	\N	untest	\N	17	31	17	f	f	18	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4690 (class 0 OID 46967)
-- Dependencies: 329
-- Data for Name: TstCaseInTaskAttachment; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseInTaskAttachment" (id, name, title, uri, descr, "docType", "caseInTaskId", "userId", "createTime", "updateTime", disabled, deleted) FROM stdin;
\.


--
-- TOC entry 4692 (class 0 OID 46975)
-- Dependencies: 331
-- Data for Name: TstCaseInTaskComments; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseInTaskComments" (id, summary, content, "caseInTaskId", "userId", "createTime", "updateTime", disabled, deleted) FROM stdin;
\.


--
-- TOC entry 4694 (class 0 OID 46983)
-- Dependencies: 333
-- Data for Name: TstCaseInTaskHistory; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseInTaskHistory" (id, title, descr, "caseId", "caseInTaskId", deleted, disabled, "createTime", "updateTime") FROM stdin;
1	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	25	f	f	2019-02-23 11:59:00.006548	\N
2	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	68	f	f	2019-02-23 12:48:59.847042	\N
3	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	71	f	f	2019-02-23 12:50:23.510283	\N
4	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	85	f	f	2019-02-23 13:44:01.796698	\N
5	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	91	f	f	2019-02-23 16:22:03.345597	\N
6	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	93	f	f	2019-02-23 16:26:25.222351	\N
7	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	95	f	f	2019-02-23 16:26:27.31754	\N
8	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	96	f	f	2019-02-23 16:26:28.264029	\N
9	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	97	f	f	2019-02-23 16:26:28.530164	\N
10	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	101	f	f	2019-02-23 16:26:28.721485	\N
11	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	102	f	f	2019-02-23 16:26:28.967924	\N
12	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	102	f	f	2019-02-23 16:26:29.218313	\N
13	用户Aaron Chen标记执行状态为标注执行结果	\N	\N	102	f	f	2019-02-23 16:26:29.586554	\N
\.


--
-- TOC entry 4696 (class 0 OID 46991)
-- Dependencies: 335
-- Data for Name: TstCaseInTaskIssue; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseInTaskIssue" (id, "issueId", "caseInTaskId", "userId", "createTime", "updateTime", disabled, deleted) FROM stdin;
\.


--
-- TOC entry 4701 (class 0 OID 47006)
-- Dependencies: 340
-- Data for Name: TstCasePriority; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCasePriority" (id, value, label, descr, ordr, "defaultVal", "buildIn", disabled, deleted, "orgId", "createTime", "updateTime") FROM stdin;
372	high	高	\N	10	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
373	medium	中	\N	20	t	t	f	f	1	2019-02-17 14:22:02.073523	\N
374	low	低	\N	30	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
375	high	高	\N	10	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
376	medium	中	\N	20	t	t	f	f	2	2019-02-18 21:51:44.67068	\N
377	low	低	\N	30	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
381	high	高	\N	10	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
382	medium	中	\N	20	t	t	f	f	4	2019-02-19 11:38:28.481573	\N
383	low	低	\N	30	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
387	high	高	\N	10	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
388	medium	中	\N	20	t	t	f	f	6	2019-02-23 13:16:23.128256	\N
389	low	低	\N	30	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
386	low	低	\N	30	f	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:07:44.029464
384	high	高	\N	10	f	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:07:44.029464
385	medium	中	\N	20	t	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:07:44.029464
390	\N	优先级01		40	f	\N	t	f	5	2019-02-23 17:07:31.352766	2019-02-23 17:09:24.704489
391	\N	优先级01	\N	40	\N	\N	t	f	5	2019-02-23 17:11:06.913514	2019-02-23 17:11:48.224994
392	high	高	\N	10	f	t	f	f	7	2019-03-11 00:56:22.895443	\N
393	medium	中	\N	20	t	t	f	f	7	2019-03-11 00:56:22.895443	\N
394	low	低	\N	30	f	t	f	f	7	2019-03-11 00:56:22.895443	\N
395	high	高	\N	10	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
396	medium	中	\N	20	t	t	f	f	8	2019-03-12 21:42:50.71871	\N
397	low	低	\N	30	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
398	high	高	\N	10	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
399	medium	中	\N	20	t	t	f	f	9	2019-03-12 21:44:18.566135	\N
400	low	低	\N	30	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
401	high	高	\N	10	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
402	medium	中	\N	20	t	t	f	f	10	2019-03-12 22:44:53.460728	\N
403	low	低	\N	30	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
404	high	高	\N	10	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
405	medium	中	\N	20	t	t	f	f	11	2019-03-12 22:52:14.395588	\N
406	low	低	\N	30	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
407	high	高	\N	10	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
408	medium	中	\N	20	t	t	f	f	12	2019-03-12 23:06:01.367716	\N
409	low	低	\N	30	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
410	high	高	\N	10	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
411	medium	中	\N	20	t	t	f	f	13	2019-03-12 23:11:45.58636	\N
412	low	低	\N	30	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
413	high	高	\N	10	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
414	medium	中	\N	20	t	t	f	f	14	2019-03-12 23:27:21.303261	\N
415	low	低	\N	30	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4699 (class 0 OID 46998)
-- Dependencies: 338
-- Data for Name: TstCasePriorityDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCasePriorityDefine" (id, value, label, descr, ordr, "defaultVal", disabled, deleted) FROM stdin;
1	high	高	\N	10	f	f	f
2	medium	中	\N	20	t	f	f
3	low	低	\N	30	f	f	f
\.


--
-- TOC entry 4702 (class 0 OID 47013)
-- Dependencies: 341
-- Data for Name: TstCaseStep; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseStep" (id, opt, expect, ordr, "caseId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	操作步骤1	期待结果1	1	3	f	f	2019-02-17 14:22:02.073523	\N
2	操作步骤2	期待结果2	2	3	f	f	2019-02-17 14:22:02.073523	\N
3	操作步骤3	期待结果3	3	3	f	f	2019-02-17 14:22:02.073523	\N
4	操作步骤1	期待结果1	1	6	f	f	2019-02-18 21:51:44.67068	\N
5	操作步骤2	期待结果2	2	6	f	f	2019-02-18 21:51:44.67068	\N
6	操作步骤3	期待结果3	3	6	f	f	2019-02-18 21:51:44.67068	\N
10	操作步骤1	期待结果1	1	12	f	f	2019-02-19 11:38:28.481573	\N
11	操作步骤2	期待结果2	2	12	f	f	2019-02-19 11:38:28.481573	\N
12	操作步骤3	期待结果3	3	12	f	f	2019-02-19 11:38:28.481573	\N
16	11	11	1	26	f	f	2019-02-22 10:38:11.919518	2019-02-22 10:38:45.618488
18	11		2	26	t	f	2019-02-22 10:38:30.412078	\N
17	2	2	2	26	t	f	2019-02-22 10:38:15.742597	\N
13	操作步骤1	期待结果1	1	15	t	f	2019-02-22 08:33:31.929668	\N
14	操作步骤2	期待结果2	1	15	f	f	2019-02-22 08:33:31.929668	\N
15	操作步骤3	期待结果3	2	15	t	f	2019-02-22 08:33:31.929668	\N
19	AA	BB	1	17	f	f	2019-02-22 21:47:26.480158	\N
20	操作步骤1	期待结果1	1	29	f	f	2019-02-23 10:46:02.008954	\N
21	操作步骤2	期待结果2	2	29	f	f	2019-02-23 10:46:02.008954	\N
22	操作步骤3	期待结果3	3	29	f	f	2019-02-23 10:46:02.008954	\N
23	操作步骤1	期待结果1	1	33	f	f	2019-02-23 12:18:51.314119	\N
24	操作步骤2	期待结果2	2	33	f	f	2019-02-23 12:18:51.314119	\N
25	操作步骤3	期待结果3	3	33	f	f	2019-02-23 12:18:51.314119	\N
26	操作步骤1	期待结果1	1	36	f	f	2019-02-23 13:16:23.128256	\N
27	操作步骤2	期待结果2	2	36	f	f	2019-02-23 13:16:23.128256	\N
28	操作步骤3	期待结果3	3	36	f	f	2019-02-23 13:16:23.128256	\N
29	操作步骤1	期待结果1	1	39	f	f	2019-03-11 00:56:22.895443	\N
30	操作步骤2	期待结果2	2	39	f	f	2019-03-11 00:56:22.895443	\N
31	操作步骤3	期待结果3	3	39	f	f	2019-03-11 00:56:22.895443	\N
32	操作步骤1	期待结果1	1	42	f	f	2019-03-12 21:42:50.71871	\N
33	操作步骤2	期待结果2	2	42	f	f	2019-03-12 21:42:50.71871	\N
34	操作步骤3	期待结果3	3	42	f	f	2019-03-12 21:42:50.71871	\N
35	操作步骤1	期待结果1	1	45	f	f	2019-03-12 21:44:18.566135	\N
36	操作步骤2	期待结果2	2	45	f	f	2019-03-12 21:44:18.566135	\N
37	操作步骤3	期待结果3	3	45	f	f	2019-03-12 21:44:18.566135	\N
38	操作步骤1	期待结果1	1	48	f	f	2019-03-12 22:44:53.460728	\N
39	操作步骤2	期待结果2	2	48	f	f	2019-03-12 22:44:53.460728	\N
40	操作步骤3	期待结果3	3	48	f	f	2019-03-12 22:44:53.460728	\N
41	操作步骤1	期待结果1	1	51	f	f	2019-03-12 22:52:14.395588	\N
42	操作步骤2	期待结果2	2	51	f	f	2019-03-12 22:52:14.395588	\N
43	操作步骤3	期待结果3	3	51	f	f	2019-03-12 22:52:14.395588	\N
44	操作步骤1	期待结果1	1	54	f	f	2019-03-12 23:06:01.367716	\N
45	操作步骤2	期待结果2	2	54	f	f	2019-03-12 23:06:01.367716	\N
46	操作步骤3	期待结果3	3	54	f	f	2019-03-12 23:06:01.367716	\N
47	操作步骤1	期待结果1	1	57	f	f	2019-03-12 23:11:45.58636	\N
48	操作步骤2	期待结果2	2	57	f	f	2019-03-12 23:11:45.58636	\N
49	操作步骤3	期待结果3	3	57	f	f	2019-03-12 23:11:45.58636	\N
50	操作步骤1	期待结果1	1	60	f	f	2019-03-12 23:27:21.303261	\N
51	操作步骤2	期待结果2	2	60	f	f	2019-03-12 23:27:21.303261	\N
52	操作步骤3	期待结果3	3	60	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4706 (class 0 OID 47029)
-- Dependencies: 345
-- Data for Name: TstCaseType; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseType" (id, value, label, descr, ordr, "defaultVal", "buildIn", disabled, deleted, "orgId", "createTime", "updateTime") FROM stdin;
870	functional	功能	\N	10	t	t	f	f	1	2019-02-17 14:22:02.073523	\N
871	performance	性能	\N	20	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
872	ui	界面	\N	30	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
873	compatibility	兼容性	\N	40	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
874	security	安全	\N	50	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
875	automation	自动化	\N	60	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
876	other	其它	\N	70	f	t	f	f	1	2019-02-17 14:22:02.073523	\N
877	functional	功能	\N	10	t	t	f	f	2	2019-02-18 21:51:44.67068	\N
878	performance	性能	\N	20	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
879	ui	界面	\N	30	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
880	compatibility	兼容性	\N	40	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
881	security	安全	\N	50	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
882	automation	自动化	\N	60	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
883	other	其它	\N	70	f	t	f	f	2	2019-02-18 21:51:44.67068	\N
891	functional	功能	\N	10	t	t	f	f	4	2019-02-19 11:38:28.481573	\N
892	performance	性能	\N	20	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
893	ui	界面	\N	30	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
894	compatibility	兼容性	\N	40	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
895	security	安全	\N	50	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
896	automation	自动化	\N	60	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
897	other	其它	\N	70	f	t	f	f	4	2019-02-19 11:38:28.481573	\N
905	functional	功能	\N	10	t	t	f	f	6	2019-02-23 13:16:23.128256	\N
906	performance	性能	\N	20	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
907	ui	界面	\N	30	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
908	compatibility	兼容性	\N	40	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
909	security	安全	\N	50	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
910	automation	自动化	\N	60	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
911	other	其它	\N	70	f	t	f	f	6	2019-02-23 13:16:23.128256	\N
935	functional	功能	\N	10	t	t	f	f	10	2019-03-12 22:44:53.460728	\N
936	performance	性能	\N	20	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
937	ui	界面	\N	30	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
938	compatibility	兼容性	\N	40	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
939	security	安全	\N	50	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
940	automation	自动化	\N	60	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
941	other	其它	\N	70	f	t	f	f	10	2019-03-12 22:44:53.460728	\N
942	functional	功能	\N	10	t	t	f	f	11	2019-03-12 22:52:14.395588	\N
904	other	其它	\N	70	f	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:02:05.352618
903	automation	自动化	\N	60	f	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:02:05.352618
902	security	安全	\N	50	f	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:02:05.352618
901	compatibility	兼容性	\N	40	f	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:02:05.352618
900	ui	界面	\N	30	f	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:02:05.352618
943	performance	性能	\N	20	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
944	ui	界面	\N	30	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
945	compatibility	兼容性	\N	40	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
946	security	安全	\N	50	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
947	automation	自动化	\N	60	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
899	performance	性能	\N	20	f	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:02:10.44992
898	functional	功能	\N	10	t	t	f	f	5	2019-02-22 08:33:31.929668	2019-02-23 17:02:10.44992
948	other	其它	\N	70	f	t	f	f	11	2019-03-12 22:52:14.395588	\N
949	functional	功能	\N	10	t	t	f	f	12	2019-03-12 23:06:01.367716	\N
912	发送到	类型01		80	f	\N	t	f	5	2019-02-23 17:02:26.8026	2019-02-23 17:09:18.189242
913	\N	类型01	\N	80	f	\N	t	f	5	2019-02-23 17:11:00.28105	2019-02-23 17:11:42.150505
914	functional	功能	\N	10	t	t	f	f	7	2019-03-11 00:56:22.895443	\N
915	performance	性能	\N	20	f	t	f	f	7	2019-03-11 00:56:22.895443	\N
916	ui	界面	\N	30	f	t	f	f	7	2019-03-11 00:56:22.895443	\N
917	compatibility	兼容性	\N	40	f	t	f	f	7	2019-03-11 00:56:22.895443	\N
918	security	安全	\N	50	f	t	f	f	7	2019-03-11 00:56:22.895443	\N
950	performance	性能	\N	20	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
951	ui	界面	\N	30	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
920	other	其它	\N	70	f	t	f	f	7	2019-03-11 00:56:22.895443	2019-03-12 17:24:47.631471
919	automation	自动化	\N	60	f	t	f	f	7	2019-03-11 00:56:22.895443	2019-03-12 17:24:47.673491
921	functional	功能	\N	10	t	t	f	f	8	2019-03-12 21:42:50.71871	\N
922	performance	性能	\N	20	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
923	ui	界面	\N	30	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
924	compatibility	兼容性	\N	40	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
925	security	安全	\N	50	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
926	automation	自动化	\N	60	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
927	other	其它	\N	70	f	t	f	f	8	2019-03-12 21:42:50.71871	\N
928	functional	功能	\N	10	t	t	f	f	9	2019-03-12 21:44:18.566135	\N
929	performance	性能	\N	20	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
930	ui	界面	\N	30	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
931	compatibility	兼容性	\N	40	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
932	security	安全	\N	50	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
933	automation	自动化	\N	60	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
934	other	其它	\N	70	f	t	f	f	9	2019-03-12 21:44:18.566135	\N
952	compatibility	兼容性	\N	40	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
953	security	安全	\N	50	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
954	automation	自动化	\N	60	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
955	other	其它	\N	70	f	t	f	f	12	2019-03-12 23:06:01.367716	\N
956	functional	功能	\N	10	t	t	f	f	13	2019-03-12 23:11:45.58636	\N
957	performance	性能	\N	20	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
958	ui	界面	\N	30	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
959	compatibility	兼容性	\N	40	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
960	security	安全	\N	50	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
961	automation	自动化	\N	60	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
962	other	其它	\N	70	f	t	f	f	13	2019-03-12 23:11:45.58636	\N
963	functional	功能	\N	10	t	t	f	f	14	2019-03-12 23:27:21.303261	\N
964	performance	性能	\N	20	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
965	ui	界面	\N	30	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
966	compatibility	兼容性	\N	40	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
967	security	安全	\N	50	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
968	automation	自动化	\N	60	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
969	other	其它	\N	70	f	t	f	f	14	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4704 (class 0 OID 47021)
-- Dependencies: 343
-- Data for Name: TstCaseTypeDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstCaseTypeDefine" (id, value, label, descr, ordr, "defaultVal", disabled, deleted) FROM stdin;
1	functional	功能	\N	10	t	f	f
2	performance	性能	\N	20	f	f	f
3	ui	界面	\N	30	f	f	f
4	compatibility	兼容性	\N	40	f	f	f
5	security	安全	\N	50	f	f	f
6	automation	自动化	\N	60	f	f	f
7	other	其它	\N	70	f	f	f
\.


--
-- TOC entry 4708 (class 0 OID 47038)
-- Dependencies: 347
-- Data for Name: TstDocument; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstDocument" (id, title, version, descr, uri, doc_type, "userId", disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4710 (class 0 OID 47046)
-- Dependencies: 349
-- Data for Name: TstEmail; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstEmail" (id, subject, content, "mailTo", "mailCc", disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4712 (class 0 OID 47054)
-- Dependencies: 351
-- Data for Name: TstEnv; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstEnv" (id, label, descr, ordr, "defaultVal", "projectId", "orgId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	环境01	描述	10	\N	10	5	f	t	2019-02-22 09:34:04.26843	2019-02-22 10:09:12.342152
2	环境01	\N	10	\N	10	5	f	f	2019-02-23 11:03:44.218639	\N
3	是打发	\N	20	\N	10	5	f	f	2019-02-23 11:14:43.672574	2019-02-23 11:31:00.587152
\.


--
-- TOC entry 4714 (class 0 OID 47062)
-- Dependencies: 353
-- Data for Name: TstHistory; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstHistory" (id, title, msg, descr, uri, "entityType", "entityId", "projectId", "userId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	2	2	1	f	f	2019-02-17 14:22:02.073523	\N
2	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	4	4	2	f	f	2019-02-18 21:51:44.67068	\N
4	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	8	8	4	f	f	2019-02-19 11:38:28.481573	\N
5	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	10	10	5	f	f	2019-02-22 08:33:31.929668	\N
6	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:36:01.728883	\N
7	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:36:10.74012	\N
8	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:36:34.858604	\N
9	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:36:43.131618	\N
10	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:45:10.246749	\N
11	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:47:09.980521	\N
12	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:50:20.277651	\N
13	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:50:27.936381	\N
14	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:50:45.676942	\N
15	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:51:14.929939	\N
16	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:51:38.69794	\N
17	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:56:17.996619	\N
18	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 09:56:26.654783	\N
19	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 10:06:20.167304	\N
20	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 10:06:59.586538	\N
21	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-22 10:07:03.20573	\N
22	用户Aaron Chen创建测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:44:06.987778	\N
23	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:44:07.194576	\N
24	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:44:12.057875	\N
25	用户Aaron Chen创建待发生{4}	\N	\N	\N	suite	2	10	5	f	f	2019-02-23 10:44:17.598463	\N
26	用户Aaron Chen更新待发生{4}	\N	\N	\N	suite	2	10	5	f	f	2019-02-23 10:44:22.000797	\N
27	用户Aaron Chen更新待发生{4}	\N	\N	\N	suite	2	10	5	f	f	2019-02-23 10:44:35.057926	\N
28	用户Aaron Chen更新待发生{4}	\N	\N	\N	suite	2	10	5	f	f	2019-02-23 10:44:35.207229	\N
29	用户Aaron Chen更新待发生{4}	\N	\N	\N	suite	2	10	5	f	f	2019-02-23 10:44:36.862286	\N
30	用户Aaron Chen更新待发生{4}	\N	\N	\N	suite	2	10	5	f	f	2019-02-23 10:45:23.064279	\N
31	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:46:15.610455	\N
32	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:46:15.701594	\N
33	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:46:18.22228	\N
34	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:56:40.083961	\N
35	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:57:16.550157	\N
36	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:57:16.618827	\N
37	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:57:26.003082	\N
38	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:57:26.027914	\N
39	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:57:28.315929	\N
40	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:57:36.554673	\N
41	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:57:36.63093	\N
42	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 10:57:38.349408	\N
43	用户Aaron Chen创建测试计划01{4}	\N	\N	\N	plan	6	10	5	f	f	2019-02-23 11:02:32.895061	\N
44	用户Aaron Chen更新示例计划{4}	\N	\N	\N	plan	5	10	5	f	f	2019-02-23 11:03:05.432896	\N
45	用户Aaron Chen更新测试计划01{4}	\N	\N	\N	plan	6	10	5	f	f	2019-02-23 11:08:37.504854	\N
46	用户Aaron Chen创建任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 11:08:56.63426	\N
47	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 11:16:50.651484	\N
48	用户Aaron Chen更新测试计划01{4}	\N	\N	\N	plan	6	10	5	f	f	2019-02-23 11:17:04.293603	\N
49	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 11:31:15.349667	\N
50	用户Aaron Chen更新测试计划01{4}	\N	\N	\N	plan	6	10	5	f	f	2019-02-23 11:31:17.357247	\N
51	用户Aaron Chen创建是的范德萨{4}	\N	\N	\N	plan	7	10	5	f	f	2019-02-23 11:32:19.951883	\N
52	用户Aaron Chen创建是打发{4}	\N	\N	\N	task	7	10	5	f	f	2019-02-23 11:32:20.135302	\N
53	用户Aaron Chen更新是打发{4}	\N	\N	\N	task	7	10	5	f	f	2019-02-23 11:33:45.685384	\N
54	用户Aaron Chen更新是打发{4}	\N	\N	\N	task	7	10	5	f	f	2019-02-23 11:33:57.013599	\N
55	用户Aaron Chen更新是打发{4}	\N	\N	\N	task	7	10	5	f	f	2019-02-23 11:39:38.628367	\N
56	用户Aaron Chen创建是打发{4}	\N	\N	\N	plan	8	10	5	f	f	2019-02-23 11:39:56.662062	\N
57	用户Aaron Chen创建是打发{4}	\N	\N	\N	task	8	10	5	f	f	2019-02-23 11:39:56.808437	\N
58	用户Aaron Chen更新是打发{4}	\N	\N	\N	task	8	10	5	f	f	2019-02-23 11:44:52.306595	\N
59	用户Aaron Chen更新是打发{4}	\N	\N	\N	plan	8	10	5	f	f	2019-02-23 11:45:07.784523	\N
60	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 12:27:18.614152	\N
61	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 12:29:08.15225	\N
62	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 12:44:33.975449	\N
65	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 12:48:33.861344	\N
66	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 12:48:39.126763	\N
67	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 12:48:44.441922	\N
68	用户Aaron Chen更新测试计划01{4}	\N	\N	\N	plan	6	10	5	f	f	2019-02-23 12:48:50.815222	\N
69	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 12:50:03.632165	\N
70	用户Aaron Chen更新测试集01{4}	\N	\N	\N	suite	1	10	5	f	f	2019-02-23 12:50:03.695033	\N
71	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 12:50:13.638419	\N
72	用户Aaron Chen更新测试计划01{4}	\N	\N	\N	plan	6	10	5	f	f	2019-02-23 12:50:15.558914	\N
73	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 12:59:51.954835	\N
74	用户Aaron Chen更新测试计划01{4}	\N	\N	\N	plan	6	10	5	f	f	2019-02-23 12:59:59.489512	\N
75	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	15	15	5	f	f	2019-02-23 13:16:23.128256	\N
76	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 16:21:41.984651	\N
77	用户Aaron Chen更新任务01{4}	\N	\N	\N	task	6	10	5	f	f	2019-02-23 16:21:52.260911	\N
78	用户Aaron Chen更新测试计划01{4}	\N	\N	\N	plan	6	10	5	f	f	2019-02-23 16:21:54.989147	\N
79	用户Aaron Chen更新默认项目{4}	\N	\N	\N	project_member	10	10	5	f	f	2019-02-25 10:01:53.253284	\N
80	用户<span class="dict">sdfds</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	17	17	11	f	f	2019-03-11 00:56:22.895443	\N
81	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	19	19	12	f	f	2019-03-12 21:42:50.71871	\N
82	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	21	21	13	f	f	2019-03-12 21:44:18.566135	\N
83	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	23	23	14	f	f	2019-03-12 22:44:53.460728	\N
84	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	25	25	15	f	f	2019-03-12 22:52:14.395588	\N
85	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	27	27	16	f	f	2019-03-12 23:06:01.367716	\N
86	用户<span class="dict">Aaron Chen</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	29	29	17	f	f	2019-03-12 23:11:45.58636	\N
87	用户<span class="dict">Aaron Chen2</span>初始化项目<span class="dict">默认项目</span>	\N	\N	\N	project	31	31	18	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4716 (class 0 OID 47070)
-- Dependencies: 355
-- Data for Name: TstModule; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstModule" (id, name, descr, ordr, "projectId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	模块01	描述	10	10	f	t	2019-02-22 09:33:00.948537	2019-02-22 10:07:12.45612
2	模块01	\N	10	10	f	f	2019-02-23 11:03:32.408275	\N
3	是打发	\N	10	17	f	f	2019-03-12 12:25:40.602231	\N
\.


--
-- TOC entry 4718 (class 0 OID 47078)
-- Dependencies: 357
-- Data for Name: TstMsg; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstMsg" (id, title, "isRead", "userId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	create_task	f	5	f	f	2019-02-23 11:08:56.63426	\N
2	update_case_for_task	f	5	f	f	2019-02-23 11:16:50.651484	\N
3	update_task	f	5	f	f	2019-02-23 11:16:50.651484	\N
4	update_case_for_task	f	5	f	f	2019-02-23 11:31:15.349667	\N
5	update_case_for_task	f	5	f	f	2019-02-23 11:32:20.135302	\N
6	create_task	f	5	f	f	2019-02-23 11:32:20.135302	\N
7	update_case_for_task	f	5	f	f	2019-02-23 11:33:45.685384	\N
8	update_task	f	5	f	f	2019-02-23 11:33:45.685384	\N
9	update_task	f	5	f	f	2019-02-23 11:33:57.013599	\N
10	update_case_for_task	f	5	f	f	2019-02-23 11:39:38.628367	\N
11	update_case_for_task	f	5	f	f	2019-02-23 11:39:56.808437	\N
12	create_task	f	5	f	f	2019-02-23 11:39:56.808437	\N
13	update_case_for_task	f	5	f	f	2019-02-23 11:44:52.306595	\N
14	update_task	f	5	f	f	2019-02-23 11:44:52.306595	\N
15	update_case_for_task	f	5	f	f	2019-02-23 12:27:18.614152	\N
16	update_case_for_task	f	5	f	f	2019-02-23 12:44:33.975449	\N
17	update_task	f	5	f	f	2019-02-23 12:44:33.975449	\N
22	update_case_for_task	f	5	f	f	2019-02-23 12:48:33.861344	\N
23	update_task	f	5	f	f	2019-02-23 12:48:33.861344	\N
24	update_task	f	5	f	f	2019-02-23 12:48:39.126763	\N
25	update_task	f	5	f	f	2019-02-23 12:48:44.441922	\N
26	update_case_for_task	f	5	f	f	2019-02-23 12:50:13.638419	\N
27	update_task	f	5	f	f	2019-02-23 12:50:13.638419	\N
28	update_case_for_task	f	5	f	f	2019-02-23 12:59:51.954835	\N
29	update_case_for_task	f	5	f	f	2019-02-23 16:21:41.984651	\N
30	update_task	f	5	f	f	2019-02-23 16:21:41.984651	\N
31	update_task	f	5	f	f	2019-02-23 16:21:52.260911	\N
32	update_issue_field	f	5	f	f	2019-02-25 10:03:46.621126	\N
33	update_issue_field	f	5	f	f	2019-02-25 10:03:51.250531	\N
34	update_issue_field	f	5	f	f	2019-02-25 10:04:13.525289	\N
35	update_issue	f	5	f	f	2019-02-25 13:06:05.078284	\N
36	update_issue	f	5	f	f	2019-02-25 13:08:54.198457	\N
37	update_issue	f	5	f	f	2019-02-25 13:21:02.23394	\N
38	update_issue	f	5	f	f	2019-02-25 13:26:09.221831	\N
39	update_issue	f	5	f	f	2019-02-25 13:26:15.883111	\N
40	update_issue	f	5	f	f	2019-02-25 13:26:58.350877	\N
41	update_issue	f	5	f	f	2019-02-25 13:31:35.187606	\N
42	update_issue	f	5	f	f	2019-02-25 13:31:49.958238	\N
43	update_issue	f	5	f	f	2019-02-25 13:38:25.469908	\N
44	create_issue	f	5	f	f	2019-02-25 13:39:07.265504	\N
45	update_issue	f	5	f	f	2019-02-25 13:40:39.499438	\N
46	update_issue	f	5	f	f	2019-02-25 13:40:59.528285	\N
47	update_issue	f	5	f	f	2019-02-25 15:59:45.287526	\N
48	update_issue_field	f	5	f	f	2019-02-25 15:59:57.40001	\N
49	create_issue	f	5	f	f	2019-02-25 16:01:53.957952	\N
50	update_issue	f	5	f	f	2019-02-25 16:02:08.238896	\N
51	create_comments_for_issue	f	5	f	f	2019-02-25 16:02:26.515123	\N
52	create_comments_for_issue	f	5	f	f	2019-02-25 16:03:44.447129	\N
53	create_comments_for_issue	f	5	f	f	2019-02-25 16:03:58.248304	\N
54	update_issue	f	5	f	f	2019-02-25 16:04:22.460222	\N
55	update_issue	f	5	f	f	2019-02-25 16:04:38.893255	\N
56	create_attament_for_issue	f	5	f	f	2019-02-25 16:20:44.988852	\N
57	create_attament_for_issue	f	5	f	f	2019-02-25 16:22:16.484225	\N
58	remove_attament_for_issue	f	5	f	f	2019-02-25 16:22:32.772315	\N
59	create_attament_for_issue	f	5	f	f	2019-02-25 16:22:53.054389	\N
60	create_comments_for_issue	f	5	f	f	2019-02-25 16:23:01.108938	\N
61	create_comments_for_issue	f	5	f	f	2019-02-25 16:23:06.298133	\N
62	create_comments_for_issue	f	5	f	f	2019-02-25 16:23:13.089547	\N
63	remove_comments_for_issue	f	5	f	f	2019-02-25 16:23:16.764895	\N
64	create_issue	f	5	f	f	2019-02-25 16:39:06.225398	\N
65	update_issue	f	5	f	f	2019-02-25 16:54:38.408307	\N
66	update_issue	f	5	f	f	2019-02-25 16:54:49.94607	\N
67	update_issue	f	5	f	f	2019-02-25 17:18:21.914195	\N
68	update_issue	f	5	f	f	2019-02-25 17:18:36.960988	\N
69	update_issue	f	5	f	f	2019-02-25 18:33:52.656293	\N
\.


--
-- TOC entry 4720 (class 0 OID 47083)
-- Dependencies: 359
-- Data for Name: TstOrg; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstOrg" (id, name, website, disabled, deleted, "createTime", "updateTime") FROM stdin;
1	Aaron Chen的组织	\N	f	f	2019-02-17 14:22:02.073523	\N
2	Aaron Chen的组织	\N	f	f	2019-02-18 21:51:44.67068	\N
4	Aaron Chen的组织	\N	f	f	2019-02-19 11:38:28.481573	\N
5	Aaron Chen的组织	\N	f	f	2019-02-22 08:33:31.929668	\N
6	组织	\N	f	f	2019-02-23 13:16:23.116	\N
7	sdfds的组织2sdf	\N	f	f	2019-03-11 00:56:22.895443	\N
8	Aaron Chen的组织	\N	f	f	2019-03-12 21:42:50.71871	\N
9	Aaron Chen的组织	\N	f	f	2019-03-12 21:44:18.566135	\N
10	Aaron Chen的组织	\N	f	f	2019-03-12 22:44:53.460728	\N
11	Aaron Chen的组织	\N	f	f	2019-03-12 22:52:14.395588	\N
12	Aaron Chen的组织	\N	f	f	2019-03-12 23:06:01.367716	\N
13	Aaron Chen的组织	\N	f	f	2019-03-12 23:11:45.58636	\N
14	Aaron Chen2的组织	\N	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4721 (class 0 OID 47089)
-- Dependencies: 360
-- Data for Name: TstOrgGroup; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstOrgGroup" (id, name, descr, "orgId", "buildIn", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	所有人	\N	1	t	f	f	2019-02-17 14:22:02.073523	\N
2	所有人	\N	2	t	f	f	2019-02-18 21:51:44.67068	\N
4	所有人	\N	4	t	f	f	2019-02-19 11:38:28.481573	\N
5	所有人	\N	5	t	f	f	2019-02-22 08:33:31.929668	\N
6	所有人	\N	6	t	f	f	2019-02-23 13:16:23.128256	\N
7	组织群组01	\N	5	f	f	f	2019-02-23 16:50:18.41403	\N
8	所有人	\N	7	t	f	f	2019-03-11 00:56:22.895443	\N
9	所有人	\N	8	t	f	f	2019-03-12 21:42:50.71871	\N
10	所有人	\N	9	t	f	f	2019-03-12 21:44:18.566135	\N
11	所有人	\N	10	t	f	f	2019-03-12 22:44:53.460728	\N
12	所有人	\N	11	t	f	f	2019-03-12 22:52:14.395588	\N
13	所有人	\N	12	t	f	f	2019-03-12 23:06:01.367716	\N
14	所有人	\N	13	t	f	f	2019-03-12 23:11:45.58636	\N
15	所有人	\N	14	t	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4722 (class 0 OID 47095)
-- Dependencies: 361
-- Data for Name: TstOrgGroupUserRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstOrgGroupUserRelation" ("orgId", "orgGroupId", "userId") FROM stdin;
1	1	1
2	2	2
4	4	4
5	5	5
5	5	6
6	6	5
5	7	5
5	5	7
5	5	8
5	5	9
5	5	10
7	8	11
8	9	12
9	10	13
10	11	14
11	12	15
12	13	16
13	14	17
14	15	18
\.


--
-- TOC entry 4724 (class 0 OID 47100)
-- Dependencies: 363
-- Data for Name: TstOrgPrivilegeDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstOrgPrivilegeDefine" (id, code, name, descr, disabled, deleted, "createTime", "updateTime", action) FROM stdin;
1	org_org	管理组织	\N	f	f	2017-04-05 09:39:15	2017-04-05 09:39:20	*
2	org_site	管理站点	\N	t	t	2017-04-05 09:39:15	2017-04-05 09:39:20	*
3	org_project	管理项目	\N	f	f	2017-04-05 09:39:15	2017-04-05 09:39:20	*
\.


--
-- TOC entry 4726 (class 0 OID 47108)
-- Dependencies: 365
-- Data for Name: TstOrgRole; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstOrgRole" (id, name, code, descr, "orgId", "buildIn", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	组织管理员	org_admin	\N	1	t	f	f	2019-02-17 14:22:02.073523	\N
2	项目管理员	project_admin	\N	1	t	f	f	2019-02-17 14:22:02.073523	\N
3	组织管理员	org_admin	\N	2	t	f	f	2019-02-18 21:51:44.67068	\N
4	项目管理员	project_admin	\N	2	t	f	f	2019-02-18 21:51:44.67068	\N
7	组织管理员	org_admin	\N	4	t	f	f	2019-02-19 11:38:28.481573	\N
8	项目管理员	project_admin	\N	4	t	f	f	2019-02-19 11:38:28.481573	\N
9	组织管理员	org_admin	\N	5	t	f	f	2019-02-22 08:33:31.929668	\N
10	项目管理员	project_admin	\N	5	t	f	f	2019-02-22 08:33:31.929668	\N
11	组织管理员	org_admin	\N	6	t	f	f	2019-02-23 13:16:23.128256	\N
12	项目管理员	project_admin	\N	6	t	f	f	2019-02-23 13:16:23.128256	\N
13	组织角色01	fc1bb687-7a9f-4a87-bf34-c594dcee36b0	是打发	5	f	f	f	2019-02-23 16:41:20.205903	2019-02-23 16:46:46.3617
14	组织管理员	org_admin	\N	7	t	f	f	2019-03-11 00:56:22.895443	\N
15	项目管理员	project_admin	\N	7	t	f	f	2019-03-11 00:56:22.895443	\N
16	组织管理员	org_admin	\N	8	t	f	f	2019-03-12 21:42:50.71871	\N
17	项目管理员	project_admin	\N	8	t	f	f	2019-03-12 21:42:50.71871	\N
18	组织管理员	org_admin	\N	9	t	f	f	2019-03-12 21:44:18.566135	\N
19	项目管理员	project_admin	\N	9	t	f	f	2019-03-12 21:44:18.566135	\N
20	组织管理员	org_admin	\N	10	t	f	f	2019-03-12 22:44:53.460728	\N
21	项目管理员	project_admin	\N	10	t	f	f	2019-03-12 22:44:53.460728	\N
22	组织管理员	org_admin	\N	11	t	f	f	2019-03-12 22:52:14.395588	\N
23	项目管理员	project_admin	\N	11	t	f	f	2019-03-12 22:52:14.395588	\N
24	组织管理员	org_admin	\N	12	t	f	f	2019-03-12 23:06:01.367716	\N
25	项目管理员	project_admin	\N	12	t	f	f	2019-03-12 23:06:01.367716	\N
26	组织管理员	org_admin	\N	13	t	f	f	2019-03-12 23:11:45.58636	\N
27	项目管理员	project_admin	\N	13	t	f	f	2019-03-12 23:11:45.58636	\N
28	组织管理员	org_admin	\N	14	t	f	f	2019-03-12 23:27:21.303261	\N
29	项目管理员	project_admin	\N	14	t	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4727 (class 0 OID 47114)
-- Dependencies: 366
-- Data for Name: TstOrgRoleGroupRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstOrgRoleGroupRelation" ("orgRoleId", "orgGroupId", "orgId") FROM stdin;
13	5	5
\.


--
-- TOC entry 4728 (class 0 OID 47117)
-- Dependencies: 367
-- Data for Name: TstOrgRolePrivilegeRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstOrgRolePrivilegeRelation" ("orgId", "orgRoleId", "orgPrivilegeId") FROM stdin;
1	1	1
1	1	3
1	2	3
2	3	1
2	3	3
2	4	3
4	7	1
4	7	3
4	8	3
5	9	1
5	9	3
5	10	3
6	11	1
6	11	3
6	12	3
5	13	1
5	13	3
7	14	1
7	14	3
7	15	3
8	16	1
8	16	3
8	17	3
9	18	1
9	18	3
9	19	3
10	20	1
10	20	3
10	21	3
11	22	1
11	22	3
11	23	3
12	24	1
12	24	3
12	25	3
13	26	1
13	26	3
13	27	3
14	28	1
14	28	3
14	29	3
\.


--
-- TOC entry 4729 (class 0 OID 47120)
-- Dependencies: 368
-- Data for Name: TstOrgRoleUserRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstOrgRoleUserRelation" ("orgRoleId", "userId", "orgId") FROM stdin;
1	1	1
3	2	2
7	4	4
9	5	5
11	5	6
13	6	5
14	11	7
16	12	8
18	13	9
20	14	10
22	15	11
24	16	12
26	17	13
28	18	14
\.


--
-- TOC entry 4731 (class 0 OID 47125)
-- Dependencies: 370
-- Data for Name: TstOrgUserRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstOrgUserRelation" ("orgId", "userId") FROM stdin;
1	1
2	2
4	4
5	5
5	6
6	5
5	7
5	8
5	9
5	10
7	11
8	12
9	13
10	14
11	15
12	16
13	17
14	18
\.


--
-- TOC entry 4733 (class 0 OID 47130)
-- Dependencies: 372
-- Data for Name: TstPlan; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstPlan" (id, name, descr, estimate, "startTime", "endTime", status, "projectId", "verId", "userId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	示例计划	\N	\N	\N	\N	not_start	2	\N	1	f	f	2019-02-17 14:22:02.073523	\N
2	示例计划	\N	\N	\N	\N	not_start	4	\N	2	f	f	2019-02-18 21:51:44.67068	\N
4	示例计划	\N	\N	\N	\N	not_start	8	\N	4	f	f	2019-02-19 11:38:28.481573	\N
5	示例计划	地方	11	2019-02-23 00:00:00	2019-02-25 00:00:00	not_start	10	\N	5	f	f	2019-02-22 08:33:31.929668	\N
8	是打发	\N	\N	\N	\N	not_start	10	\N	5	f	t	2019-02-23 11:39:56.596	\N
7	是的范德萨	\N	\N	\N	\N	not_start	10	\N	5	f	t	2019-02-23 11:32:19.708	\N
9	示例计划	\N	\N	\N	\N	not_start	15	\N	5	f	f	2019-02-23 13:16:23.128256	\N
6	测试计划01	\N	\N	\N	\N	in_progress	10	2	5	f	f	2019-02-23 11:02:32.815	\N
10	示例计划	\N	\N	\N	\N	not_start	17	\N	11	f	f	2019-03-11 00:56:22.895443	\N
11	示例计划	\N	\N	\N	\N	not_start	19	\N	12	f	f	2019-03-12 21:42:50.71871	\N
12	示例计划	\N	\N	\N	\N	not_start	21	\N	13	f	f	2019-03-12 21:44:18.566135	\N
13	示例计划	\N	\N	\N	\N	not_start	23	\N	14	f	f	2019-03-12 22:44:53.460728	\N
14	示例计划	\N	\N	\N	\N	not_start	25	\N	15	f	f	2019-03-12 22:52:14.395588	\N
15	示例计划	\N	\N	\N	\N	not_start	27	\N	16	f	f	2019-03-12 23:06:01.367716	\N
16	示例计划	\N	\N	\N	\N	not_start	29	\N	17	f	f	2019-03-12 23:11:45.58636	\N
17	示例计划	\N	\N	\N	\N	not_start	31	\N	18	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4735 (class 0 OID 47138)
-- Dependencies: 374
-- Data for Name: TstProject; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstProject" (id, name, descr, type, "issueTypeSolutionId", "issuePrioritySolutionId", "issuePageSolutionId", "issueWorkflowSolutionId", "lastAccessTime", "orgId", "parentId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	默认项目组	\N	group	\N	\N	\N	\N	\N	1	\N	f	f	2019-02-17 14:22:02.073523	\N
2	默认项目	\N	project	1	1	1	1	\N	1	1	f	f	2019-02-17 14:22:02.073523	\N
3	默认项目组	\N	group	\N	\N	\N	\N	\N	2	\N	f	f	2019-02-18 21:51:44.67068	\N
4	默认项目	\N	project	2	2	2	2	\N	2	3	f	f	2019-02-18 21:51:44.67068	\N
7	默认项目组	\N	group	\N	\N	\N	\N	\N	4	\N	f	f	2019-02-19 11:38:28.481573	\N
8	默认项目	\N	project	4	4	3	3	\N	4	7	f	f	2019-02-19 11:38:28.481573	\N
11	新项目组01	\N	group	5	5	4	4	\N	5	\N	f	f	2019-02-22 09:27:36.187	\N
9	默认项目组		group	\N	\N	\N	\N	\N	5	\N	f	f	2019-02-22 08:33:31.929668	\N
10	默认项目	\N	project	5	5	4	4	\N	5	9	f	f	2019-02-22 08:33:31.929668	\N
12	新项目01	\N	project	5	5	4	4	\N	5	9	f	t	2019-02-23 10:46:01.996	2019-02-23 12:18:41.660243
13	项目01	\N	project	5	5	4	4	\N	5	9	f	f	2019-02-23 12:18:51.287	\N
14	默认项目组	\N	group	\N	\N	\N	\N	\N	6	\N	f	f	2019-02-23 13:16:23.128256	\N
15	默认项目	\N	project	6	6	5	5	\N	6	14	f	f	2019-02-23 13:16:23.128256	\N
16	默认项目组	\N	group	\N	\N	\N	\N	\N	7	\N	f	f	2019-03-11 00:56:22.895443	\N
17	默认项目	\N	project	7	7	10	10	\N	7	16	f	f	2019-03-11 00:56:22.895443	\N
18	默认项目组	\N	group	\N	\N	\N	\N	\N	8	\N	f	f	2019-03-12 21:42:50.71871	\N
19	默认项目	\N	project	8	8	11	11	\N	8	18	f	f	2019-03-12 21:42:50.71871	\N
20	默认项目组	\N	group	\N	\N	\N	\N	\N	9	\N	f	f	2019-03-12 21:44:18.566135	\N
21	默认项目	\N	project	9	9	12	12	\N	9	20	f	f	2019-03-12 21:44:18.566135	\N
22	默认项目组	\N	group	\N	\N	\N	\N	\N	10	\N	f	f	2019-03-12 22:44:53.460728	\N
23	默认项目	\N	project	10	10	13	13	\N	10	22	f	f	2019-03-12 22:44:53.460728	\N
24	默认项目组	\N	group	\N	\N	\N	\N	\N	11	\N	f	f	2019-03-12 22:52:14.395588	\N
25	默认项目	\N	project	11	11	14	14	\N	11	24	f	f	2019-03-12 22:52:14.395588	\N
26	默认项目组	\N	group	\N	\N	\N	\N	\N	12	\N	f	f	2019-03-12 23:06:01.367716	\N
27	默认项目	\N	project	12	12	15	15	\N	12	26	f	f	2019-03-12 23:06:01.367716	\N
28	默认项目组	\N	group	\N	\N	\N	\N	\N	13	\N	f	f	2019-03-12 23:11:45.58636	\N
29	默认项目	\N	project	13	13	16	16	\N	13	28	f	f	2019-03-12 23:11:45.58636	\N
30	默认项目组	\N	group	\N	\N	\N	\N	\N	14	\N	f	f	2019-03-12 23:27:21.303261	\N
31	默认项目	\N	project	14	14	17	17	\N	14	30	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4736 (class 0 OID 47144)
-- Dependencies: 375
-- Data for Name: TstProjectAccessHistory; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstProjectAccessHistory" (id, "lastAccessTime", "orgId", "prjId", "prjName", "userId", "createTime", "updateTime") FROM stdin;
1	2019-02-17 14:22:02.073523	1	2	默认项目	1	2019-02-17 14:22:02.073523	\N
2	2019-02-18 21:51:44.67068	2	4	默认项目	2	2019-02-18 21:51:44.67068	\N
4	2019-02-19 11:38:28.481573	4	8	默认项目	4	2019-02-19 11:38:28.481573	\N
5	2019-02-22 08:33:31.929668	5	10	默认项目	5	2019-02-22 08:33:31.929668	\N
6	2019-02-22 09:35:14.358807	5	10	默认项目	6	\N	\N
7	2019-02-23 10:46:02.008954	5	12	新项目01	5	\N	\N
8	2019-02-23 12:18:51.314119	5	13	项目01	5	\N	\N
9	2019-02-23 13:16:23.128256	6	15	默认项目	5	2019-02-23 13:16:23.128256	\N
10	2019-03-10 18:16:08.037903	5	10	默认项目	7	\N	\N
11	2019-03-10 18:17:49.090395	5	10	默认项目	8	\N	\N
12	2019-03-10 18:23:46.242772	5	10	默认项目	9	\N	\N
13	2019-03-10 18:24:59.108467	5	10	默认项目	10	\N	\N
14	2019-03-11 00:56:22.895443	7	17	默认项目	11	2019-03-11 00:56:22.895443	\N
15	2019-03-12 21:42:50.71871	8	19	默认项目	12	2019-03-12 21:42:50.71871	\N
16	2019-03-12 21:44:18.566135	9	21	默认项目	13	2019-03-12 21:44:18.566135	\N
17	2019-03-12 22:44:53.460728	10	23	默认项目	14	2019-03-12 22:44:53.460728	\N
18	2019-03-12 22:52:14.395588	11	25	默认项目	15	2019-03-12 22:52:14.395588	\N
19	2019-03-12 23:06:01.367716	12	27	默认项目	16	2019-03-12 23:06:01.367716	\N
20	2019-03-12 23:11:45.58636	13	29	默认项目	17	2019-03-12 23:11:45.58636	\N
21	2019-03-12 23:27:21.303261	14	31	默认项目	18	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4738 (class 0 OID 47149)
-- Dependencies: 377
-- Data for Name: TstProjectPrivilegeDefine; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstProjectPrivilegeDefine" (id, code, name, action, "actionName", descr, disabled, deleted, "createTime", "updateTime") FROM stdin;
12100	test_case	测试用例	view	查看	\N	f	f	2017-12-26 10:11:16	2017-12-26 10:11:18
12200	test_case	测试用例	maintain	维护	\N	f	f	2017-04-05 11:52:26	2017-04-05 11:52:28
12300	test_case	测试用例	delete	删除	\N	f	f	2017-04-05 11:52:26	2017-04-05 11:52:28
12400	test_case	测试用例	review	评审	\N	f	f	2018-09-16 08:15:23	2018-09-16 08:15:26
13100	test_suite	测试集	view	查看	\N	f	f	2017-12-26 10:18:29	2017-12-26 10:18:38
13200	test_suite	测试集	maintain	维护	\N	f	f	2017-04-05 11:52:26	2017-04-05 11:52:28
13300	test_suite	测试集	delete	删除	\N	f	f	2017-04-05 11:52:26	2017-04-05 11:52:28
14100	test_plan	执行计划	view	查看	\N	f	f	2017-12-26 10:13:08	2017-12-26 10:13:11
14200	test_plan	执行计划	maintain	维护	\N	f	f	2017-04-05 11:52:26	2017-04-05 11:52:28
14300	test_plan	执行计划	delete	删除	\N	f	f	2017-04-05 11:52:26	2017-04-05 11:52:28
15100	test_task	测试任务	view	查看	\N	f	f	2017-04-05 11:52:26	2017-04-05 11:52:28
15200	test_task	测试任务	exe	执行	\N	f	f	2017-04-05 11:52:26	2017-04-05 11:52:28
15300	test_task	测试任务	close	关闭	\N	f	f	2017-04-05 11:52:26	2017-04-05 11:52:28
17100	issue	问题	view	查看	\N	f	f	2018-05-03 17:03:01	2018-05-03 17:03:08
17200	issue	问题	maintain	维护	\N	f	f	2018-05-03 17:03:01	2018-05-03 17:03:08
17300	issue	问题	delete	删除	\N	f	f	2018-05-03 17:03:01	2018-05-03 17:03:08
12000	project	项目	maintain	维护	\N	f	f	2017-12-26 10:11:16	2017-12-26 10:11:18
\.


--
-- TOC entry 4740 (class 0 OID 47157)
-- Dependencies: 379
-- Data for Name: TstProjectRole; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstProjectRole" (id, code, name, descr, "buildIn", "orgId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	test_leader	测试主管	\N	f	1	f	f	2019-02-17 14:22:02.073523	\N
2	test_designer	测试设计	\N	f	1	f	f	2019-02-17 14:22:02.073523	\N
3	tester	测试执行	\N	f	1	f	f	2019-02-17 14:22:02.073523	\N
4	readonly	只读用户	\N	f	1	f	f	2019-02-17 14:22:02.073523	\N
5	test_leader	测试主管	\N	f	2	f	f	2019-02-18 21:51:44.67068	\N
6	test_designer	测试设计	\N	f	2	f	f	2019-02-18 21:51:44.67068	\N
7	tester	测试执行	\N	f	2	f	f	2019-02-18 21:51:44.67068	\N
8	readonly	只读用户	\N	f	2	f	f	2019-02-18 21:51:44.67068	\N
13	test_leader	测试主管	\N	f	4	f	f	2019-02-19 11:38:28.481573	\N
14	test_designer	测试设计	\N	f	4	f	f	2019-02-19 11:38:28.481573	\N
15	tester	测试执行	\N	f	4	f	f	2019-02-19 11:38:28.481573	\N
16	readonly	只读用户	\N	f	4	f	f	2019-02-19 11:38:28.481573	\N
17	test_leader	测试主管	\N	f	5	f	f	2019-02-22 08:33:31.929668	\N
18	test_designer	测试设计	\N	f	5	f	f	2019-02-22 08:33:31.929668	\N
19	tester	测试执行	\N	f	5	f	f	2019-02-22 08:33:31.929668	\N
21	test_leader	测试主管	\N	f	6	f	f	2019-02-23 13:16:23.128256	\N
22	test_designer	测试设计	\N	f	6	f	f	2019-02-23 13:16:23.128256	\N
23	tester	测试执行	\N	f	6	f	f	2019-02-23 13:16:23.128256	\N
24	readonly	只读用户	\N	f	6	f	f	2019-02-23 13:16:23.128256	\N
20	readonly	只读用户	\N	f	5	f	f	2019-02-22 08:33:31.929668	2019-02-23 16:49:13.066816
26	test_designer	测试设计	\N	f	7	f	f	2019-03-11 00:56:22.895443	\N
27	tester	测试执行	\N	f	7	f	f	2019-03-11 00:56:22.895443	\N
28	readonly	只读用户	\N	f	7	f	f	2019-03-11 00:56:22.895443	\N
25	test_leader	测试主管	\N	f	7	f	f	2019-03-11 00:56:22.895443	2019-03-12 17:18:35.285133
29	test_leader	测试主管	\N	f	8	f	f	2019-03-12 21:42:50.71871	\N
30	test_designer	测试设计	\N	f	8	f	f	2019-03-12 21:42:50.71871	\N
31	tester	测试执行	\N	f	8	f	f	2019-03-12 21:42:50.71871	\N
32	readonly	只读用户	\N	f	8	f	f	2019-03-12 21:42:50.71871	\N
33	test_leader	测试主管	\N	f	9	f	f	2019-03-12 21:44:18.566135	\N
34	test_designer	测试设计	\N	f	9	f	f	2019-03-12 21:44:18.566135	\N
35	tester	测试执行	\N	f	9	f	f	2019-03-12 21:44:18.566135	\N
36	readonly	只读用户	\N	f	9	f	f	2019-03-12 21:44:18.566135	\N
37	test_leader	测试主管	\N	f	10	f	f	2019-03-12 22:44:53.460728	\N
38	test_designer	测试设计	\N	f	10	f	f	2019-03-12 22:44:53.460728	\N
39	tester	测试执行	\N	f	10	f	f	2019-03-12 22:44:53.460728	\N
40	readonly	只读用户	\N	f	10	f	f	2019-03-12 22:44:53.460728	\N
41	test_leader	测试主管	\N	f	11	f	f	2019-03-12 22:52:14.395588	\N
42	test_designer	测试设计	\N	f	11	f	f	2019-03-12 22:52:14.395588	\N
43	tester	测试执行	\N	f	11	f	f	2019-03-12 22:52:14.395588	\N
44	readonly	只读用户	\N	f	11	f	f	2019-03-12 22:52:14.395588	\N
45	test_leader	测试主管	\N	f	12	f	f	2019-03-12 23:06:01.367716	\N
46	test_designer	测试设计	\N	f	12	f	f	2019-03-12 23:06:01.367716	\N
47	tester	测试执行	\N	f	12	f	f	2019-03-12 23:06:01.367716	\N
48	readonly	只读用户	\N	f	12	f	f	2019-03-12 23:06:01.367716	\N
49	test_leader	测试主管	\N	f	13	f	f	2019-03-12 23:11:45.58636	\N
50	test_designer	测试设计	\N	f	13	f	f	2019-03-12 23:11:45.58636	\N
51	tester	测试执行	\N	f	13	f	f	2019-03-12 23:11:45.58636	\N
52	readonly	只读用户	\N	f	13	f	f	2019-03-12 23:11:45.58636	\N
53	test_leader	测试主管	\N	f	14	f	f	2019-03-12 23:27:21.303261	\N
54	test_designer	测试设计	\N	f	14	f	f	2019-03-12 23:27:21.303261	\N
55	tester	测试执行	\N	f	14	f	f	2019-03-12 23:27:21.303261	\N
56	readonly	只读用户	\N	f	14	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4741 (class 0 OID 47163)
-- Dependencies: 380
-- Data for Name: TstProjectRoleEntityRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstProjectRoleEntityRelation" ("entityId", "orgId", "projectId", "projectRoleId", type) FROM stdin;
1	1	2	1	user
2	2	4	5	user
4	4	8	13	user
5	5	10	17	user
5	5	12	17	user
5	5	13	17	user
5	6	15	21	user
6	5	10	18	user
7	5	10	18	user
8	5	10	18	user
9	5	10	18	user
10	5	10	18	user
11	7	17	25	user
12	8	19	29	user
13	9	21	33	user
14	10	23	37	user
15	11	25	41	user
16	12	27	45	user
17	13	29	49	user
18	14	31	53	user
\.


--
-- TOC entry 4742 (class 0 OID 47166)
-- Dependencies: 381
-- Data for Name: TstProjectRolePriviledgeRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstProjectRolePriviledgeRelation" ("projectPrivilegeDefineId", "projectRoleId", "orgId") FROM stdin;
12100	1	1
12200	1	1
12300	1	1
12400	1	1
13100	1	1
13200	1	1
13300	1	1
14100	1	1
14200	1	1
14300	1	1
15100	1	1
15200	1	1
15300	1	1
17100	1	1
17200	1	1
17300	1	1
12100	2	\N
12200	2	\N
12300	2	\N
13100	2	\N
13200	2	\N
13300	2	\N
14100	2	\N
14200	2	\N
14300	2	\N
15100	2	\N
15200	2	\N
15300	2	\N
17100	2	\N
17200	2	\N
17300	2	\N
12100	3	\N
12300	3	\N
13100	3	\N
13200	3	\N
13300	3	\N
14100	3	\N
14200	3	\N
14300	3	\N
15100	3	\N
15200	3	\N
15300	3	\N
17100	3	\N
17200	3	\N
17300	3	\N
12100	4	\N
13100	4	\N
14100	4	\N
15100	4	\N
17100	4	\N
12100	5	2
12200	5	2
12300	5	2
12400	5	2
13100	5	2
13200	5	2
13300	5	2
14100	5	2
14200	5	2
14300	5	2
15100	5	2
15200	5	2
15300	5	2
17100	5	2
17200	5	2
17300	5	2
12100	6	\N
12200	6	\N
12300	6	\N
13100	6	\N
13200	6	\N
13300	6	\N
14100	6	\N
14200	6	\N
14300	6	\N
15100	6	\N
15200	6	\N
15300	6	\N
17100	6	\N
17200	6	\N
17300	6	\N
12100	7	\N
12300	7	\N
13100	7	\N
13200	7	\N
13300	7	\N
14100	7	\N
14200	7	\N
14300	7	\N
15100	7	\N
15200	7	\N
15300	7	\N
17100	7	\N
17200	7	\N
17300	7	\N
12100	8	\N
13100	8	\N
14100	8	\N
15100	8	\N
17100	8	\N
12100	13	4
12200	13	4
12300	13	4
12400	13	4
13100	13	4
13200	13	4
13300	13	4
14100	13	4
14200	13	4
14300	13	4
15100	13	4
15200	13	4
15300	13	4
17100	13	4
17200	13	4
17300	13	4
12100	14	\N
12200	14	\N
12300	14	\N
13100	14	\N
13200	14	\N
13300	14	\N
14100	14	\N
14200	14	\N
14300	14	\N
15100	14	\N
15200	14	\N
15300	14	\N
17100	14	\N
17200	14	\N
17300	14	\N
12100	15	\N
12300	15	\N
13100	15	\N
13200	15	\N
13300	15	\N
14100	15	\N
14200	15	\N
14300	15	\N
15100	15	\N
15200	15	\N
15300	15	\N
17100	15	\N
17200	15	\N
17300	15	\N
12100	16	\N
13100	16	\N
14100	16	\N
15100	16	\N
17100	16	\N
12100	17	5
12200	17	5
12300	17	5
12400	17	5
13100	17	5
13200	17	5
13300	17	5
14100	17	5
14200	17	5
14300	17	5
15100	17	5
15200	17	5
15300	17	5
17100	17	5
17200	17	5
17300	17	5
12100	18	\N
12200	18	\N
12300	18	\N
13100	18	\N
13200	18	\N
13300	18	\N
14100	18	\N
14200	18	\N
14300	18	\N
15100	18	\N
15200	18	\N
15300	18	\N
17100	18	\N
17200	18	\N
17300	18	\N
12100	19	\N
12300	19	\N
13100	19	\N
13200	19	\N
13300	19	\N
14100	19	\N
14200	19	\N
14300	19	\N
15100	19	\N
15200	19	\N
15300	19	\N
17100	19	\N
17200	19	\N
17300	19	\N
12100	21	6
12200	21	6
12300	21	6
12400	21	6
13100	21	6
13200	21	6
13300	21	6
14100	21	6
14200	21	6
14300	21	6
15100	21	6
15200	21	6
15300	21	6
17100	21	6
17200	21	6
17300	21	6
12100	22	\N
12200	22	\N
12300	22	\N
13100	22	\N
13200	22	\N
13300	22	\N
14100	22	\N
14200	22	\N
14300	22	\N
15100	22	\N
15200	22	\N
15300	22	\N
17100	22	\N
17200	22	\N
17300	22	\N
12100	23	\N
12300	23	\N
13100	23	\N
13200	23	\N
13300	23	\N
14100	23	\N
14200	23	\N
14300	23	\N
15100	23	\N
15200	23	\N
15300	23	\N
17100	23	\N
17200	23	\N
17300	23	\N
12100	24	\N
13100	24	\N
14100	24	\N
15100	24	\N
17100	24	\N
12100	20	5
13100	20	5
14100	20	5
15100	20	5
17100	20	5
12100	26	\N
12200	26	\N
12300	26	\N
13100	26	\N
13200	26	\N
13300	26	\N
14100	26	\N
14200	26	\N
14300	26	\N
15100	26	\N
15200	26	\N
15300	26	\N
17100	26	\N
17200	26	\N
17300	26	\N
12100	27	\N
12300	27	\N
13100	27	\N
13200	27	\N
13300	27	\N
14100	27	\N
14200	27	\N
14300	27	\N
15100	27	\N
15200	27	\N
15300	27	\N
17100	27	\N
17200	27	\N
17300	27	\N
12100	28	\N
13100	28	\N
14100	28	\N
15100	28	\N
17100	28	\N
12000	25	7
12100	25	7
12200	25	7
12400	25	7
12300	25	7
13100	25	7
13200	25	7
13300	25	7
14100	25	7
14200	25	7
14300	25	7
15100	25	7
15200	25	7
15300	25	7
17100	25	7
17200	25	7
17300	25	7
12100	29	8
12200	29	8
12300	29	8
12400	29	8
13100	29	8
13200	29	8
13300	29	8
14100	29	8
14200	29	8
14300	29	8
15100	29	8
15200	29	8
15300	29	8
17100	29	8
17200	29	8
17300	29	8
12000	29	8
12100	30	\N
12200	30	\N
12300	30	\N
13100	30	\N
13200	30	\N
13300	30	\N
14100	30	\N
14200	30	\N
14300	30	\N
15100	30	\N
15200	30	\N
15300	30	\N
17100	30	\N
17200	30	\N
17300	30	\N
12100	31	\N
12300	31	\N
13100	31	\N
13200	31	\N
13300	31	\N
14100	31	\N
14200	31	\N
14300	31	\N
15100	31	\N
15200	31	\N
15300	31	\N
17100	31	\N
17200	31	\N
17300	31	\N
12100	32	\N
13100	32	\N
14100	32	\N
15100	32	\N
17100	32	\N
12100	33	9
12200	33	9
12300	33	9
12400	33	9
13100	33	9
13200	33	9
13300	33	9
14100	33	9
14200	33	9
14300	33	9
15100	33	9
15200	33	9
15300	33	9
17100	33	9
17200	33	9
17300	33	9
12000	33	9
12100	34	\N
12200	34	\N
12300	34	\N
13100	34	\N
13200	34	\N
13300	34	\N
14100	34	\N
14200	34	\N
14300	34	\N
15100	34	\N
15200	34	\N
15300	34	\N
17100	34	\N
17200	34	\N
17300	34	\N
12100	35	\N
12300	35	\N
13100	35	\N
13200	35	\N
13300	35	\N
14100	35	\N
14200	35	\N
14300	35	\N
15100	35	\N
15200	35	\N
15300	35	\N
17100	35	\N
17200	35	\N
17300	35	\N
12100	36	\N
13100	36	\N
14100	36	\N
15100	36	\N
17100	36	\N
12100	37	10
12200	37	10
12300	37	10
12400	37	10
13100	37	10
13200	37	10
13300	37	10
14100	37	10
14200	37	10
14300	37	10
15100	37	10
15200	37	10
15300	37	10
17100	37	10
17200	37	10
17300	37	10
12000	37	10
12100	38	\N
12200	38	\N
12300	38	\N
13100	38	\N
13200	38	\N
13300	38	\N
14100	38	\N
14200	38	\N
14300	38	\N
15100	38	\N
15200	38	\N
15300	38	\N
17100	38	\N
17200	38	\N
17300	38	\N
12100	39	\N
12300	39	\N
13100	39	\N
13200	39	\N
13300	39	\N
14100	39	\N
14200	39	\N
14300	39	\N
15100	39	\N
15200	39	\N
15300	39	\N
17100	39	\N
17200	39	\N
17300	39	\N
12100	40	\N
13100	40	\N
14100	40	\N
15100	40	\N
17100	40	\N
12100	41	11
12200	41	11
12300	41	11
12400	41	11
13100	41	11
13200	41	11
13300	41	11
14100	41	11
14200	41	11
14300	41	11
15100	41	11
15200	41	11
15300	41	11
17100	41	11
17200	41	11
17300	41	11
12000	41	11
12100	42	\N
12200	42	\N
12300	42	\N
13100	42	\N
13200	42	\N
13300	42	\N
14100	42	\N
14200	42	\N
14300	42	\N
15100	42	\N
15200	42	\N
15300	42	\N
17100	42	\N
17200	42	\N
17300	42	\N
12100	43	\N
12300	43	\N
13100	43	\N
13200	43	\N
13300	43	\N
14100	43	\N
14200	43	\N
14300	43	\N
15100	43	\N
15200	43	\N
15300	43	\N
17100	43	\N
17200	43	\N
17300	43	\N
12100	44	\N
13100	44	\N
14100	44	\N
15100	44	\N
17100	44	\N
12100	45	12
12200	45	12
12300	45	12
12400	45	12
13100	45	12
13200	45	12
13300	45	12
14100	45	12
14200	45	12
14300	45	12
15100	45	12
15200	45	12
15300	45	12
17100	45	12
17200	45	12
17300	45	12
12000	45	12
12100	46	\N
12200	46	\N
12300	46	\N
13100	46	\N
13200	46	\N
13300	46	\N
14100	46	\N
14200	46	\N
14300	46	\N
15100	46	\N
15200	46	\N
15300	46	\N
17100	46	\N
17200	46	\N
17300	46	\N
12100	47	\N
12300	47	\N
13100	47	\N
13200	47	\N
13300	47	\N
14100	47	\N
14200	47	\N
14300	47	\N
15100	47	\N
15200	47	\N
15300	47	\N
17100	47	\N
17200	47	\N
17300	47	\N
12100	48	\N
13100	48	\N
14100	48	\N
15100	48	\N
17100	48	\N
12100	49	13
12200	49	13
12300	49	13
12400	49	13
13100	49	13
13200	49	13
13300	49	13
14100	49	13
14200	49	13
14300	49	13
15100	49	13
15200	49	13
15300	49	13
17100	49	13
17200	49	13
17300	49	13
12000	49	13
12100	50	\N
12200	50	\N
12300	50	\N
13100	50	\N
13200	50	\N
13300	50	\N
14100	50	\N
14200	50	\N
14300	50	\N
15100	50	\N
15200	50	\N
15300	50	\N
17100	50	\N
17200	50	\N
17300	50	\N
12100	51	\N
12300	51	\N
13100	51	\N
13200	51	\N
13300	51	\N
14100	51	\N
14200	51	\N
14300	51	\N
15100	51	\N
15200	51	\N
15300	51	\N
17100	51	\N
17200	51	\N
17300	51	\N
12100	52	\N
13100	52	\N
14100	52	\N
15100	52	\N
17100	52	\N
12100	53	14
12200	53	14
12300	53	14
12400	53	14
13100	53	14
13200	53	14
13300	53	14
14100	53	14
14200	53	14
14300	53	14
15100	53	14
15200	53	14
15300	53	14
17100	53	14
17200	53	14
17300	53	14
12000	53	14
12100	54	\N
12200	54	\N
12300	54	\N
13100	54	\N
13200	54	\N
13300	54	\N
14100	54	\N
14200	54	\N
14300	54	\N
15100	54	\N
15200	54	\N
15300	54	\N
17100	54	\N
17200	54	\N
17300	54	\N
12100	55	\N
12300	55	\N
13100	55	\N
13200	55	\N
13300	55	\N
14100	55	\N
14200	55	\N
14300	55	\N
15100	55	\N
15200	55	\N
15300	55	\N
17100	55	\N
17200	55	\N
17300	55	\N
12100	56	\N
13100	56	\N
14100	56	\N
15100	56	\N
17100	56	\N
\.


--
-- TOC entry 4745 (class 0 OID 47173)
-- Dependencies: 384
-- Data for Name: TstSuite; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstSuite" (id, name, descr, estimate, ordr, "projectId", "userId", "caseProjectId", disabled, deleted, "createTime", "updateTime") FROM stdin;
2	待发生	是打发	\N	\N	10	5	10	f	t	2019-02-23 10:44:17.555433	2019-02-23 10:45:23.022087
1	测试集01	\N	1	\N	10	5	10	f	f	2019-02-23 10:44:06.909744	2019-02-23 12:50:03.695033
\.


--
-- TOC entry 4747 (class 0 OID 47181)
-- Dependencies: 386
-- Data for Name: TstTask; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstTask" (id, name, descr, estimate, status, "projectId", "caseProjectId", "planId", "userId", "envId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	示例任务	\N	\N	not_start	2	2	1	1	\N	f	f	2019-02-17 14:22:02.073523	\N
2	示例任务	\N	\N	not_start	4	4	2	2	\N	f	f	2019-02-18 21:51:44.67068	\N
4	示例任务	\N	\N	not_start	8	8	4	4	\N	f	f	2019-02-19 11:38:28.481573	\N
5	示例任务	\N	\N	not_start	10	10	5	5	\N	f	f	2019-02-22 08:33:31.929668	\N
7	是打发	\N	\N	not_start	10	12	7	5	3	f	f	2019-02-23 11:32:20.135302	2019-02-23 11:39:38.628367
8	是打发	\N	\N	not_start	10	12	8	5	\N	f	f	2019-02-23 11:39:56.808437	2019-02-23 11:44:52.306595
9	示例任务	\N	\N	not_start	15	15	9	5	\N	f	f	2019-02-23 13:16:23.128256	\N
6	任务01	\N	\N	in_progress	10	13	6	5	2	f	f	2019-02-23 11:08:56.63426	2019-02-23 16:21:52.260911
10	示例任务	\N	\N	not_start	17	17	10	11	\N	f	f	2019-03-11 00:56:22.895443	\N
11	示例任务	\N	\N	not_start	19	19	11	12	\N	f	f	2019-03-12 21:42:50.71871	\N
12	示例任务	\N	\N	not_start	21	21	12	13	\N	f	f	2019-03-12 21:44:18.566135	\N
13	示例任务	\N	\N	not_start	23	23	13	14	\N	f	f	2019-03-12 22:44:53.460728	\N
14	示例任务	\N	\N	not_start	25	25	14	15	\N	f	f	2019-03-12 22:52:14.395588	\N
15	示例任务	\N	\N	not_start	27	27	15	16	\N	f	f	2019-03-12 23:06:01.367716	\N
16	示例任务	\N	\N	not_start	29	29	16	17	\N	f	f	2019-03-12 23:11:45.58636	\N
17	示例任务	\N	\N	not_start	31	31	17	18	\N	f	f	2019-03-12 23:27:21.303261	\N
\.


--
-- TOC entry 4748 (class 0 OID 47187)
-- Dependencies: 387
-- Data for Name: TstTaskAssigneeRelation; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstTaskAssigneeRelation" ("taskId", "assigneeId") FROM stdin;
1	1
2	2
4	4
5	5
7	5
8	5
9	5
6	5
10	11
11	12
12	13
13	14
14	15
15	16
16	17
17	18
\.


--
-- TOC entry 4750 (class 0 OID 47192)
-- Dependencies: 389
-- Data for Name: TstThread; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstThread" (id, content, "authorId", "parentId", disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- TOC entry 4752 (class 0 OID 47200)
-- Dependencies: 391
-- Data for Name: TstUser; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstUser" (id, email, nickname, password, phone, avatar, "defaultOrgId", "defaultOrgName", "defaultPrjId", "defaultPrjName", salt, "verifyCode", "lastLoginTime", disabled, deleted, "createTime", "updateTime", locked) FROM stdin;
13	46282第三方6@qq.com	Aaron Chen	c9884f1e0852f4e51b74942f5c3370e1	11111111111	upload/sample/user/avatar.png	9	Aaron Chen的组织	21	默认项目	[B@1781f14f	\N	2019-03-12 22:35:06.713564	f	f	2019-03-12 21:44:18.536	2019-03-12 22:35:06.713564	f
4	462d826@qq.com	Aaron Chen	5983640513ecd316958e5f3ac86473bc	11111111111	upload/sample/user/avatar.png	4	Aaron Chen的组织	8	默认项目	[B@11516ca5	\N	2019-02-20 19:05:02.667	f	f	2019-02-19 11:38:28.462	2019-02-22 08:33:17.045108	\N
8	test02@ngtesting.com	test02@ngtesting.com	\N	\N	upload/sample/user/avatar.png	5	Aaron Chen的组织	10	默认项目	\N	\N	\N	f	f	2019-03-10 18:17:49.088	\N	\N
2	4628是打发26@qq.com	Aaron Chen	3d37bfc1aa25edf66235f99041a1f3d7	11111111111	upload/sample/user/avatar.png	2	Aaron Chen的组织	4	默认项目	[B@5a409e02	\N	2019-02-19 09:15:23.818	f	f	2019-02-18 21:51:44.662	2019-02-19 11:35:07.446743	\N
7	qi.chens@ngtesting.com	qi.chen@ngtesting.com	\N	\N	upload/sample/user/avatar.png	5	Aaron Chen的组织	10	默认项目	\N	\N	\N	f	f	2019-03-10 18:16:08.035	\N	\N
9	是打发	qi.chen@ngtesting.com	\N	\N	upload/sample/user/avatar.png	5	Aaron Chen的组织	10	默认项目	\N	\N	\N	f	f	2019-03-10 18:23:46.24	\N	\N
10	qi.chen@ngtesting.com	qi.chen@ngtesting.com	1fee41c2e76d98fc2aa48996e349a379	\N	upload/sample/user/avatar.png	5	Aaron Chen的组织	10	默认项目	[B@148329ae	\N	\N	f	f	2019-03-10 18:24:59.102	\N	\N
1	4628SDF26@qq.com	Aaron Chen	2da7561c70ac0054fcda1d90279d637e	11111111111	upload/sample/user/avatar.png	1	Aaron Chen的组织	2	默认项目	[B@ad7115e	\N	2019-02-17 14:22:12.707	f	f	2019-02-17 14:22:01.898	2019-02-18 21:51:06.314518	\N
6	test01@ngtesting.com	test01	\N	\N	upload/sample/user/avatar.png	5	Aaron Chen的组织	10	默认项目	\N	\N	\N	f	f	2019-02-22 09:35:14.353	\N	\N
14	46282226@qq.com	Aaron Chen	53681491922ad22c70562e6280409043	11111111111	upload/sample/user/avatar.png	10	Aaron Chen的组织	23	默认项目	[B@7c84f682	\N	2019-03-12 22:45:30.399464	f	f	2019-03-12 22:44:53.279	2019-03-12 22:45:30.399464	f
11	sdf@ss.com	sdfds	fe8b5b0580455511dbe4a23184017b8a	11111111111	upload/sample/user/avatar.png	7	sdfds的组织	17	默认项目	[B@39efaed0	\N	\N	f	f	2019-03-11 00:56:22.839	\N	f
5	46282uu6@qq.com	Aaron Chen	8f124d1bdbb487e55d6a7bec82cf8709	11111111111	upload/sample/user/avatar.png	5	Aaron Chen的组织	10	默认项目	[B@dbc0dab	\N	2019-03-10 21:44:03.341	f	f	2019-02-22 08:33:31.904	\N	\N
12	46282sdf6@qq.com	Aaron Chen	478dc269ece1f030befc8619bbea2236	11111111111	upload/sample/user/avatar.png	8	Aaron Chen的组织	19	默认项目	[B@3fda4c11	\N	\N	f	f	2019-03-12 21:42:50.581	\N	f
15	462826dsf@qq.com	Aaron Chen	aa876b9e704dcd6c9e3e676e02d8c766	11111111111	upload/sample/user/avatar.png	11	Aaron Chen的组织	25	默认项目	[B@40abaed3	\N	\N	f	f	2019-03-12 22:51:15.105	\N	f
16	55@qq.com	Aaron Chen	e901e89f356de1d393d42138185cf043	11111111111	upload/sample/user/avatar.png	12	Aaron Chen的组织	27	默认项目	[B@79fe03f0	\N	\N	f	f	2019-03-12 23:06:01.021	\N	f
17	46282dfsd6@qq.com	Aaron Chen	8a3a2940b686d288ecc03efe5368e2a3	11111111111	upload/sample/user/avatar.png	13	Aaron Chen的组织	29	默认项目	[B@4205e27f	\N	\N	f	f	2019-03-12 23:09:47.21	\N	f
18	462826@qq.com	Aaron Chen2	92ea2a3602949e9953c284cdebf70750	11111111111	upload/sample/user/avatar.png	14	Aaron Chen2的组织	31	默认项目	[B@4ed89c0b	\N	2019-03-12 23:27:52.581438	f	f	2019-03-12 23:27:13.495	2019-03-12 23:27:52.581438	f
\.


--
-- TOC entry 4753 (class 0 OID 47206)
-- Dependencies: 392
-- Data for Name: TstUserSettings; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstUserSettings" ("leftSizeDesign", "leftSizeExe", "leftSizeIssue", "issueView", "issueColumns", "issueFields", tql, "userId") FROM stdin;
300	200	\N	\N	title,typeId,statusId,priorityId,assigneeId	\N	\N	1
300	200	\N	\N	title,typeId,statusId,priorityId,assigneeId	\N	\N	2
300	200	\N	\N	title,typeId,statusId,priorityId,assigneeId	\N	\N	4
300	200	\N	\N	\N	\N	\N	6
300	200	248	table	title,typeId,statusId,priorityId,assigneeId	\N	\N	5
300	200	\N	\N	\N	\N	\N	7
300	200	\N	\N	\N	\N	\N	8
300	200	\N	\N	\N	\N	\N	9
300	200	\N	\N	\N	\N	\N	10
300	200	\N	\N	\N	\N	\N	11
300	200	\N	\N	\N	\N	\N	12
300	200	\N	\N	\N	\N	\N	13
300	200	\N	\N	\N	\N	\N	14
300	200	\N	\N	\N	\N	\N	15
300	200	\N	\N	\N	\N	\N	16
300	200	\N	\N	\N	\N	\N	17
300	200	\N	\N	\N	\N	\N	18
\.


--
-- TOC entry 4754 (class 0 OID 47212)
-- Dependencies: 393
-- Data for Name: TstUserVerifyCode; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstUserVerifyCode" (id, code, "expireTime", "userId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	c11cbf8ff70e47c19d7f24f527d8926f	2019-02-17 14:32:02.736	1	f	f	2019-02-17 14:22:02.736	\N
2	c2f8c7a338414cffa47093995fd461a6	2019-02-18 22:01:45.061	2	f	f	2019-02-18 21:51:45.061	\N
3	ad16669deaec4f08ac8109c30128b879	2019-02-19 11:48:28.81	4	f	f	2019-02-19 11:38:28.81	\N
4	bcf3c083abeb4df488c7fee5ce36b707	2019-02-22 08:43:32.374	5	f	f	2019-02-22 08:33:32.374	\N
5	fd02f83530a94eb8ae040737cf241636	2019-02-22 09:45:14.436	6	f	f	2019-02-22 09:35:14.436	\N
6	e66d0a7c3a7a40928aba440fa92c2d62	2019-03-10 18:26:08.15	7	f	f	2019-03-10 18:16:08.15	\N
7	5a822e25a88e43fab4ddfdd78c30741b	2019-03-10 18:27:49.118	8	f	f	2019-03-10 18:17:49.118	\N
8	95718ef243af409d8f7cf0ce66286a4f	2019-03-10 18:34:22.464	9	f	f	2019-03-10 18:24:22.464	\N
9	4e2afdb94fa54eb1b96e8ae297d458ba	2019-03-10 18:34:59.158	10	f	f	2019-03-10 18:24:59.158	\N
10	b8f660f3194643cdbf9afe792c55f5c3	2019-03-11 01:06:23.369	11	f	f	2019-03-11 00:56:23.369	\N
11	5283e783d36c4116ab4126d94f714963	2019-03-12 21:52:51.12	12	f	f	2019-03-12 21:42:51.12	\N
12	482dcff073134bc5906f42599e669ede	2019-03-12 21:54:18.61	13	t	t	2019-03-12 21:44:18.61	2019-03-12 22:35:06.708221
13	aa36d8404fad46fe89d66a3ecf7546a1	2019-03-12 22:54:54.191	14	t	t	2019-03-12 22:44:54.191	2019-03-12 22:45:30.338278
14	1b54c8a3392b496ba6333716fb91b75e	2019-03-12 23:02:14.492	15	f	f	2019-03-12 22:52:14.492	\N
15	14d1312ee155405782ed593e3c31e0d0	2019-03-12 23:16:01.734	16	f	f	2019-03-12 23:06:01.734	\N
16	360d9f23639041b89bba0cd7196097b4	2019-03-12 23:21:45.923	17	f	f	2019-03-12 23:11:45.923	\N
17	e1d5d6a2efcc48c4b91dd68d0805e7df	2019-03-12 23:37:22.094	18	t	t	2019-03-12 23:27:22.094	2019-03-12 23:27:52.577408
\.


--
-- TOC entry 4757 (class 0 OID 47219)
-- Dependencies: 396
-- Data for Name: TstVer; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public."TstVer" (id, label, descr, status, "startTime", "endTime", "defaultVal", ordr, "projectId", "orgId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	版本01	描述	\N	\N	\N	\N	10	10	5	f	t	2019-02-22 09:33:31.605157	2019-02-22 10:08:39.861038
2	版本01	\N	\N	\N	\N	\N	10	10	5	f	f	2019-02-23 11:03:38.124808	\N
\.


--
-- TOC entry 4856 (class 0 OID 0)
-- Dependencies: 199
-- Name: CustomFieldDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."CustomFieldDefine_id_seq"', 1, false);


--
-- TOC entry 4857 (class 0 OID 0)
-- Dependencies: 201
-- Name: CustomFieldInputTypeRelationDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."CustomFieldInputTypeRelationDefine_id_seq"', 1, false);


--
-- TOC entry 4858 (class 0 OID 0)
-- Dependencies: 203
-- Name: CustomFieldIputDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."CustomFieldIputDefine_id_seq"', 1, false);


--
-- TOC entry 4859 (class 0 OID 0)
-- Dependencies: 206
-- Name: CustomFieldOptionDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."CustomFieldOptionDefine_id_seq"', 1, false);


--
-- TOC entry 4860 (class 0 OID 0)
-- Dependencies: 207
-- Name: CustomFieldOption_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."CustomFieldOption_id_seq"', 74, true);


--
-- TOC entry 4861 (class 0 OID 0)
-- Dependencies: 209
-- Name: CustomFieldTypeDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."CustomFieldTypeDefine_id_seq"', 1, false);


--
-- TOC entry 4862 (class 0 OID 0)
-- Dependencies: 210
-- Name: CustomField_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."CustomField_id_seq"', 18, true);


--
-- TOC entry 4863 (class 0 OID 0)
-- Dependencies: 212
-- Name: IsuAttachment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuAttachment_id_seq"', 3, true);


--
-- TOC entry 4864 (class 0 OID 0)
-- Dependencies: 214
-- Name: IsuComments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuComments_id_seq"', 5, true);


--
-- TOC entry 4865 (class 0 OID 0)
-- Dependencies: 218
-- Name: IsuCustomFieldSolution_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuCustomFieldSolution_id_seq"', 1, false);


--
-- TOC entry 4866 (class 0 OID 0)
-- Dependencies: 220
-- Name: IsuDocument_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuDocument_id_seq"', 1, false);


--
-- TOC entry 4867 (class 0 OID 0)
-- Dependencies: 223
-- Name: IsuFieldCodeToTableDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuFieldCodeToTableDefine_id_seq"', 1, false);


--
-- TOC entry 4868 (class 0 OID 0)
-- Dependencies: 225
-- Name: IsuFieldDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuFieldDefine_id_seq"', 1, true);


--
-- TOC entry 4869 (class 0 OID 0)
-- Dependencies: 226
-- Name: IsuField_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuField_id_seq"', 196, true);


--
-- TOC entry 4870 (class 0 OID 0)
-- Dependencies: 228
-- Name: IsuHistory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuHistory_id_seq"', 107, true);


--
-- TOC entry 4871 (class 0 OID 0)
-- Dependencies: 231
-- Name: IsuIssueExt_pid_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuIssueExt_pid_seq"', 1, false);


--
-- TOC entry 4872 (class 0 OID 0)
-- Dependencies: 232
-- Name: IsuIssue_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuIssue_id_seq"', 17, true);


--
-- TOC entry 4873 (class 0 OID 0)
-- Dependencies: 235
-- Name: IsuLinkReasonDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuLinkReasonDefine_id_seq"', 1, false);


--
-- TOC entry 4874 (class 0 OID 0)
-- Dependencies: 236
-- Name: IsuLink_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuLink_id_seq"', 2, true);


--
-- TOC entry 4875 (class 0 OID 0)
-- Dependencies: 239
-- Name: IsuNotificationDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuNotificationDefine_id_seq"', 1, false);


--
-- TOC entry 4876 (class 0 OID 0)
-- Dependencies: 240
-- Name: IsuNotification_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuNotification_id_seq"', 1, false);


--
-- TOC entry 4877 (class 0 OID 0)
-- Dependencies: 243
-- Name: IsuPageElement_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuPageElement_id_seq"', 149, true);


--
-- TOC entry 4878 (class 0 OID 0)
-- Dependencies: 246
-- Name: IsuPageSolutionItem_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuPageSolutionItem_id_seq"', 89, true);


--
-- TOC entry 4879 (class 0 OID 0)
-- Dependencies: 247
-- Name: IsuPageSolution_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuPageSolution_id_seq"', 17, true);


--
-- TOC entry 4880 (class 0 OID 0)
-- Dependencies: 248
-- Name: IsuPage_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuPage_id_seq"', 26, true);


--
-- TOC entry 4881 (class 0 OID 0)
-- Dependencies: 251
-- Name: IsuPriorityDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuPriorityDefine_id_seq"', 1, false);


--
-- TOC entry 4882 (class 0 OID 0)
-- Dependencies: 254
-- Name: IsuPrioritySolution_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuPrioritySolution_id_seq"', 14, true);


--
-- TOC entry 4883 (class 0 OID 0)
-- Dependencies: 255
-- Name: IsuPriority_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuPriority_id_seq"', 57, true);


--
-- TOC entry 4884 (class 0 OID 0)
-- Dependencies: 257
-- Name: IsuQuery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuQuery_id_seq"', 1, false);


--
-- TOC entry 4885 (class 0 OID 0)
-- Dependencies: 260
-- Name: IsuResolutionDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuResolutionDefine_id_seq"', 1, false);


--
-- TOC entry 4886 (class 0 OID 0)
-- Dependencies: 261
-- Name: IsuResolution_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuResolution_id_seq"', 29, true);


--
-- TOC entry 4887 (class 0 OID 0)
-- Dependencies: 264
-- Name: IsuSeverityDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuSeverityDefine_id_seq"', 1, false);


--
-- TOC entry 4888 (class 0 OID 0)
-- Dependencies: 267
-- Name: IsuSeveritySolution_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuSeveritySolution_id_seq"', 1, false);


--
-- TOC entry 4889 (class 0 OID 0)
-- Dependencies: 268
-- Name: IsuSeverity_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuSeverity_id_seq"', 1, false);


--
-- TOC entry 4890 (class 0 OID 0)
-- Dependencies: 271
-- Name: IsuStatusCategoryDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuStatusCategoryDefine_id_seq"', 1, false);


--
-- TOC entry 4891 (class 0 OID 0)
-- Dependencies: 273
-- Name: IsuStatusDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuStatusDefine_id_seq"', 1, false);


--
-- TOC entry 4892 (class 0 OID 0)
-- Dependencies: 274
-- Name: IsuStatus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuStatus_id_seq"', 73, true);


--
-- TOC entry 4893 (class 0 OID 0)
-- Dependencies: 277
-- Name: IsuTagRelation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuTagRelation_id_seq"', 4, true);


--
-- TOC entry 4894 (class 0 OID 0)
-- Dependencies: 278
-- Name: IsuTag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuTag_id_seq"', 4, true);


--
-- TOC entry 4895 (class 0 OID 0)
-- Dependencies: 281
-- Name: IsuTypeDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuTypeDefine_id_seq"', 1, false);


--
-- TOC entry 4896 (class 0 OID 0)
-- Dependencies: 284
-- Name: IsuTypeSolution_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuTypeSolution_id_seq"', 14, true);


--
-- TOC entry 4897 (class 0 OID 0)
-- Dependencies: 285
-- Name: IsuType_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuType_id_seq"', 29, true);


--
-- TOC entry 4898 (class 0 OID 0)
-- Dependencies: 287
-- Name: IsuWatch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuWatch_id_seq"', 2, true);


--
-- TOC entry 4899 (class 0 OID 0)
-- Dependencies: 291
-- Name: IsuWorkflowSolutionItem_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuWorkflowSolutionItem_id_seq"', 38, true);


--
-- TOC entry 4900 (class 0 OID 0)
-- Dependencies: 292
-- Name: IsuWorkflowSolution_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuWorkflowSolution_id_seq"', 17, true);


--
-- TOC entry 4901 (class 0 OID 0)
-- Dependencies: 295
-- Name: IsuWorkflowStatusRelationDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuWorkflowStatusRelationDefine_id_seq"', 1, false);


--
-- TOC entry 4902 (class 0 OID 0)
-- Dependencies: 296
-- Name: IsuWorkflowStatusRelation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuWorkflowStatusRelation_id_seq"', 71, true);


--
-- TOC entry 4903 (class 0 OID 0)
-- Dependencies: 299
-- Name: IsuWorkflowTransitionDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuWorkflowTransitionDefine_id_seq"', 1, false);


--
-- TOC entry 4904 (class 0 OID 0)
-- Dependencies: 301
-- Name: IsuWorkflowTransitionProjectRoleRelation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuWorkflowTransitionProjectRoleRelation_id_seq"', 508, true);


--
-- TOC entry 4905 (class 0 OID 0)
-- Dependencies: 302
-- Name: IsuWorkflowTransition_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuWorkflowTransition_id_seq"', 126, true);


--
-- TOC entry 4906 (class 0 OID 0)
-- Dependencies: 303
-- Name: IsuWorkflow_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."IsuWorkflow_id_seq"', 14, true);


--
-- TOC entry 4907 (class 0 OID 0)
-- Dependencies: 305
-- Name: SysPrivilege_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."SysPrivilege_id_seq"', 1, false);


--
-- TOC entry 4908 (class 0 OID 0)
-- Dependencies: 309
-- Name: SysRole_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."SysRole_id_seq"', 1, false);


--
-- TOC entry 4909 (class 0 OID 0)
-- Dependencies: 311
-- Name: SysUser_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."SysUser_id_seq"', 1, false);


--
-- TOC entry 4910 (class 0 OID 0)
-- Dependencies: 313
-- Name: Test_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."Test_id_seq"', 1, false);


--
-- TOC entry 4911 (class 0 OID 0)
-- Dependencies: 315
-- Name: TstAlert_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstAlert_id_seq"', 16, true);


--
-- TOC entry 4912 (class 0 OID 0)
-- Dependencies: 318
-- Name: TstCaseAttachment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseAttachment_id_seq"', 9, true);


--
-- TOC entry 4913 (class 0 OID 0)
-- Dependencies: 320
-- Name: TstCaseComments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseComments_id_seq"', 20, true);


--
-- TOC entry 4914 (class 0 OID 0)
-- Dependencies: 322
-- Name: TstCaseExeStatus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseExeStatus_id_seq"', 574, true);


--
-- TOC entry 4915 (class 0 OID 0)
-- Dependencies: 325
-- Name: TstCaseHistory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseHistory_id_seq"', 88, true);


--
-- TOC entry 4916 (class 0 OID 0)
-- Dependencies: 327
-- Name: TstCaseInSuite_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseInSuite_id_seq"', 55, true);


--
-- TOC entry 4917 (class 0 OID 0)
-- Dependencies: 330
-- Name: TstCaseInTaskAttachment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseInTaskAttachment_id_seq"', 1, false);


--
-- TOC entry 4918 (class 0 OID 0)
-- Dependencies: 332
-- Name: TstCaseInTaskComments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseInTaskComments_id_seq"', 1, false);


--
-- TOC entry 4919 (class 0 OID 0)
-- Dependencies: 334
-- Name: TstCaseInTaskHistory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseInTaskHistory_id_seq"', 13, true);


--
-- TOC entry 4920 (class 0 OID 0)
-- Dependencies: 336
-- Name: TstCaseInTaskIssue_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseInTaskIssue_id_seq"', 1, false);


--
-- TOC entry 4921 (class 0 OID 0)
-- Dependencies: 337
-- Name: TstCaseInTask_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseInTask_id_seq"', 126, true);


--
-- TOC entry 4922 (class 0 OID 0)
-- Dependencies: 339
-- Name: TstCasePriority_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCasePriority_id_seq"', 415, true);


--
-- TOC entry 4923 (class 0 OID 0)
-- Dependencies: 342
-- Name: TstCaseStep_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseStep_id_seq"', 52, true);


--
-- TOC entry 4924 (class 0 OID 0)
-- Dependencies: 344
-- Name: TstCaseType_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCaseType_id_seq"', 969, true);


--
-- TOC entry 4925 (class 0 OID 0)
-- Dependencies: 346
-- Name: TstCase_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstCase_id_seq"', 62, true);


--
-- TOC entry 4926 (class 0 OID 0)
-- Dependencies: 348
-- Name: TstDocument_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstDocument_id_seq"', 1, false);


--
-- TOC entry 4927 (class 0 OID 0)
-- Dependencies: 350
-- Name: TstEmail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstEmail_id_seq"', 1, false);


--
-- TOC entry 4928 (class 0 OID 0)
-- Dependencies: 352
-- Name: TstEnv_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstEnv_id_seq"', 3, true);


--
-- TOC entry 4929 (class 0 OID 0)
-- Dependencies: 354
-- Name: TstHistory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstHistory_id_seq"', 87, true);


--
-- TOC entry 4930 (class 0 OID 0)
-- Dependencies: 356
-- Name: TstModule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstModule_id_seq"', 3, true);


--
-- TOC entry 4931 (class 0 OID 0)
-- Dependencies: 358
-- Name: TstMsg_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstMsg_id_seq"', 69, true);


--
-- TOC entry 4932 (class 0 OID 0)
-- Dependencies: 362
-- Name: TstOrgGroup_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstOrgGroup_id_seq"', 15, true);


--
-- TOC entry 4933 (class 0 OID 0)
-- Dependencies: 364
-- Name: TstOrgPrivilegeDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstOrgPrivilegeDefine_id_seq"', 1, false);


--
-- TOC entry 4934 (class 0 OID 0)
-- Dependencies: 369
-- Name: TstOrgRole_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstOrgRole_id_seq"', 29, true);


--
-- TOC entry 4935 (class 0 OID 0)
-- Dependencies: 371
-- Name: TstOrg_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstOrg_id_seq"', 14, true);


--
-- TOC entry 4936 (class 0 OID 0)
-- Dependencies: 373
-- Name: TstPlan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstPlan_id_seq"', 17, true);


--
-- TOC entry 4937 (class 0 OID 0)
-- Dependencies: 376
-- Name: TstProjectAccessHistory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstProjectAccessHistory_id_seq"', 21, true);


--
-- TOC entry 4938 (class 0 OID 0)
-- Dependencies: 378
-- Name: TstProjectPrivilegeDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstProjectPrivilegeDefine_id_seq"', 1, false);


--
-- TOC entry 4939 (class 0 OID 0)
-- Dependencies: 382
-- Name: TstProjectRole_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstProjectRole_id_seq"', 56, true);


--
-- TOC entry 4940 (class 0 OID 0)
-- Dependencies: 383
-- Name: TstProject_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstProject_id_seq"', 31, true);


--
-- TOC entry 4941 (class 0 OID 0)
-- Dependencies: 385
-- Name: TstSuite_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstSuite_id_seq"', 2, true);


--
-- TOC entry 4942 (class 0 OID 0)
-- Dependencies: 388
-- Name: TstTask_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstTask_id_seq"', 17, true);


--
-- TOC entry 4943 (class 0 OID 0)
-- Dependencies: 390
-- Name: TstThread_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstThread_id_seq"', 1, false);


--
-- TOC entry 4944 (class 0 OID 0)
-- Dependencies: 394
-- Name: TstUserVerifyCode_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstUserVerifyCode_id_seq"', 17, true);


--
-- TOC entry 4945 (class 0 OID 0)
-- Dependencies: 395
-- Name: TstUser_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstUser_id_seq"', 18, true);


--
-- TOC entry 4946 (class 0 OID 0)
-- Dependencies: 397
-- Name: TstVer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public."TstVer_id_seq"', 2, true);


--
-- TOC entry 3878 (class 2606 OID 47322)
-- Name: CustomFieldDefine CustomFieldDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldDefine"
    ADD CONSTRAINT "CustomFieldDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3880 (class 2606 OID 47324)
-- Name: CustomFieldInputTypeRelationDefine CustomFieldInputTypeRelationDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldInputTypeRelationDefine"
    ADD CONSTRAINT "CustomFieldInputTypeRelationDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3882 (class 2606 OID 47326)
-- Name: CustomFieldIputDefine CustomFieldIputDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldIputDefine"
    ADD CONSTRAINT "CustomFieldIputDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3888 (class 2606 OID 47328)
-- Name: CustomFieldOptionDefine CustomFieldOptionDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldOptionDefine"
    ADD CONSTRAINT "CustomFieldOptionDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3884 (class 2606 OID 47330)
-- Name: CustomFieldOption CustomFieldOption_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldOption"
    ADD CONSTRAINT "CustomFieldOption_pkey" PRIMARY KEY (id);


--
-- TOC entry 3891 (class 2606 OID 47332)
-- Name: CustomFieldTypeDefine CustomFieldTypeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldTypeDefine"
    ADD CONSTRAINT "CustomFieldTypeDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3875 (class 2606 OID 47334)
-- Name: CustomField CustomField_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomField"
    ADD CONSTRAINT "CustomField_pkey" PRIMARY KEY (id);


--
-- TOC entry 3893 (class 2606 OID 47336)
-- Name: IsuAttachment IsuAttachment_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuAttachment"
    ADD CONSTRAINT "IsuAttachment_pkey" PRIMARY KEY (id);


--
-- TOC entry 3897 (class 2606 OID 47338)
-- Name: IsuComments IsuComments_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuComments"
    ADD CONSTRAINT "IsuComments_pkey" PRIMARY KEY (id);


--
-- TOC entry 3901 (class 2606 OID 47340)
-- Name: IsuCustomFieldSolution IsuCustomFieldSolution_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuCustomFieldSolution"
    ADD CONSTRAINT "IsuCustomFieldSolution_pkey" PRIMARY KEY (id);


--
-- TOC entry 3909 (class 2606 OID 47342)
-- Name: IsuDocument IsuDocument_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuDocument"
    ADD CONSTRAINT "IsuDocument_pkey" PRIMARY KEY (id);


--
-- TOC entry 3916 (class 2606 OID 47344)
-- Name: IsuFieldCodeToTableDefine IsuFieldCodeToTableDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuFieldCodeToTableDefine"
    ADD CONSTRAINT "IsuFieldCodeToTableDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3918 (class 2606 OID 47346)
-- Name: IsuFieldDefine IsuFieldDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuFieldDefine"
    ADD CONSTRAINT "IsuFieldDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3913 (class 2606 OID 47348)
-- Name: IsuField IsuField_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuField"
    ADD CONSTRAINT "IsuField_pkey" PRIMARY KEY (id);


--
-- TOC entry 3920 (class 2606 OID 47350)
-- Name: IsuHistory IsuHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuHistory"
    ADD CONSTRAINT "IsuHistory_pkey" PRIMARY KEY (id);


--
-- TOC entry 3937 (class 2606 OID 47352)
-- Name: IsuIssueExt IsuIssueExt_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssueExt"
    ADD CONSTRAINT "IsuIssueExt_pkey" PRIMARY KEY (pid);


--
-- TOC entry 3923 (class 2606 OID 47354)
-- Name: IsuIssue IsuIssue_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_pkey" PRIMARY KEY (id);


--
-- TOC entry 3945 (class 2606 OID 47356)
-- Name: IsuLinkReasonDefine IsuLinkReasonDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuLinkReasonDefine"
    ADD CONSTRAINT "IsuLinkReasonDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3940 (class 2606 OID 47358)
-- Name: IsuLink IsuLink_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuLink"
    ADD CONSTRAINT "IsuLink_pkey" PRIMARY KEY (id);


--
-- TOC entry 3950 (class 2606 OID 47360)
-- Name: IsuNotificationDefine IsuNotificationDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuNotificationDefine"
    ADD CONSTRAINT "IsuNotificationDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3947 (class 2606 OID 47362)
-- Name: IsuNotification IsuNotification_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuNotification"
    ADD CONSTRAINT "IsuNotification_pkey" PRIMARY KEY (id);


--
-- TOC entry 3955 (class 2606 OID 47364)
-- Name: IsuPageElement IsuPageElement_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageElement"
    ADD CONSTRAINT "IsuPageElement_pkey" PRIMARY KEY (id);


--
-- TOC entry 3960 (class 2606 OID 47366)
-- Name: IsuPageSolutionItem IsuPageSolutionItem_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_pkey" PRIMARY KEY (id);


--
-- TOC entry 3957 (class 2606 OID 47368)
-- Name: IsuPageSolution IsuPageSolution_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageSolution"
    ADD CONSTRAINT "IsuPageSolution_pkey" PRIMARY KEY (id);


--
-- TOC entry 3952 (class 2606 OID 47370)
-- Name: IsuPage IsuPage_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPage"
    ADD CONSTRAINT "IsuPage_pkey" PRIMARY KEY (id);


--
-- TOC entry 3969 (class 2606 OID 47372)
-- Name: IsuPriorityDefine IsuPriorityDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPriorityDefine"
    ADD CONSTRAINT "IsuPriorityDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3971 (class 2606 OID 47374)
-- Name: IsuPrioritySolution IsuPrioritySolution_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPrioritySolution"
    ADD CONSTRAINT "IsuPrioritySolution_pkey" PRIMARY KEY (id);


--
-- TOC entry 3966 (class 2606 OID 47376)
-- Name: IsuPriority IsuPriority_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPriority"
    ADD CONSTRAINT "IsuPriority_pkey" PRIMARY KEY (id);


--
-- TOC entry 3977 (class 2606 OID 47378)
-- Name: IsuQuery IsuQuery_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuQuery"
    ADD CONSTRAINT "IsuQuery_pkey" PRIMARY KEY (id);


--
-- TOC entry 3984 (class 2606 OID 47380)
-- Name: IsuResolutionDefine IsuResolutionDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuResolutionDefine"
    ADD CONSTRAINT "IsuResolutionDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3981 (class 2606 OID 47382)
-- Name: IsuResolution IsuResolution_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuResolution"
    ADD CONSTRAINT "IsuResolution_pkey" PRIMARY KEY (id);


--
-- TOC entry 3989 (class 2606 OID 47384)
-- Name: IsuSeverityDefine IsuSeverityDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeverityDefine"
    ADD CONSTRAINT "IsuSeverityDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3991 (class 2606 OID 47386)
-- Name: IsuSeveritySolution IsuSeveritySolution_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeveritySolution"
    ADD CONSTRAINT "IsuSeveritySolution_pkey" PRIMARY KEY (id);


--
-- TOC entry 3986 (class 2606 OID 47388)
-- Name: IsuSeverity IsuSeverity_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeverity"
    ADD CONSTRAINT "IsuSeverity_pkey" PRIMARY KEY (id);


--
-- TOC entry 4000 (class 2606 OID 47390)
-- Name: IsuStatusCategoryDefine IsuStatusCategoryDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuStatusCategoryDefine"
    ADD CONSTRAINT "IsuStatusCategoryDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 4002 (class 2606 OID 47392)
-- Name: IsuStatusDefine IsuStatusDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuStatusDefine"
    ADD CONSTRAINT "IsuStatusDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 3996 (class 2606 OID 47394)
-- Name: IsuStatus IsuStatus_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuStatus"
    ADD CONSTRAINT "IsuStatus_pkey" PRIMARY KEY (id);


--
-- TOC entry 4009 (class 2606 OID 47396)
-- Name: IsuTagRelation IsuTagRelation_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTagRelation"
    ADD CONSTRAINT "IsuTagRelation_pkey" PRIMARY KEY (id);


--
-- TOC entry 4005 (class 2606 OID 47398)
-- Name: IsuTag IsuTag_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTag"
    ADD CONSTRAINT "IsuTag_pkey" PRIMARY KEY (id);


--
-- TOC entry 4016 (class 2606 OID 47400)
-- Name: IsuTypeDefine IsuTypeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTypeDefine"
    ADD CONSTRAINT "IsuTypeDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 4018 (class 2606 OID 47402)
-- Name: IsuTypeSolution IsuTypeSolution_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTypeSolution"
    ADD CONSTRAINT "IsuTypeSolution_pkey" PRIMARY KEY (id);


--
-- TOC entry 4013 (class 2606 OID 47404)
-- Name: IsuType IsuType_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuType"
    ADD CONSTRAINT "IsuType_pkey" PRIMARY KEY (id);


--
-- TOC entry 4024 (class 2606 OID 47406)
-- Name: IsuWatch IsuWatch_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWatch"
    ADD CONSTRAINT "IsuWatch_pkey" PRIMARY KEY (id);


--
-- TOC entry 4034 (class 2606 OID 47408)
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_pkey" PRIMARY KEY (id);


--
-- TOC entry 4031 (class 2606 OID 47410)
-- Name: IsuWorkflowSolution IsuWorkflowSolution_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowSolution"
    ADD CONSTRAINT "IsuWorkflowSolution_pkey" PRIMARY KEY (id);


--
-- TOC entry 4045 (class 2606 OID 47412)
-- Name: IsuWorkflowStatusRelationDefine IsuWorkflowStatusRelationDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelationDefine"
    ADD CONSTRAINT "IsuWorkflowStatusRelationDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 4040 (class 2606 OID 47414)
-- Name: IsuWorkflowStatusRelation IsuWorkflowStatusRelation_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation"
    ADD CONSTRAINT "IsuWorkflowStatusRelation_pkey" PRIMARY KEY (id);


--
-- TOC entry 4053 (class 2606 OID 47416)
-- Name: IsuWorkflowTransitionDefine IsuWorkflowTransitionDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionDefine"
    ADD CONSTRAINT "IsuWorkflowTransitionDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 4057 (class 2606 OID 47418)
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelation_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_pkey" PRIMARY KEY (id);


--
-- TOC entry 4048 (class 2606 OID 47420)
-- Name: IsuWorkflowTransition IsuWorkflowTransition_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransition"
    ADD CONSTRAINT "IsuWorkflowTransition_pkey" PRIMARY KEY (id);


--
-- TOC entry 4028 (class 2606 OID 47422)
-- Name: IsuWorkflow IsuWorkflow_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflow"
    ADD CONSTRAINT "IsuWorkflow_pkey" PRIMARY KEY (id);


--
-- TOC entry 4063 (class 2606 OID 47424)
-- Name: SysPrivilege SysPrivilege_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysPrivilege"
    ADD CONSTRAINT "SysPrivilege_pkey" PRIMARY KEY (id);


--
-- TOC entry 4065 (class 2606 OID 47426)
-- Name: SysRole SysRole_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysRole"
    ADD CONSTRAINT "SysRole_pkey" PRIMARY KEY (id);


--
-- TOC entry 4071 (class 2606 OID 47428)
-- Name: SysUser SysUser_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysUser"
    ADD CONSTRAINT "SysUser_pkey" PRIMARY KEY (id);


--
-- TOC entry 4073 (class 2606 OID 47430)
-- Name: Test Test_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."Test"
    ADD CONSTRAINT "Test_pkey" PRIMARY KEY (id);


--
-- TOC entry 4075 (class 2606 OID 47432)
-- Name: TstAlert TstAlert_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstAlert"
    ADD CONSTRAINT "TstAlert_pkey" PRIMARY KEY (id);


--
-- TOC entry 4088 (class 2606 OID 47434)
-- Name: TstCaseAttachment TstCaseAttachment_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseAttachment"
    ADD CONSTRAINT "TstCaseAttachment_pkey" PRIMARY KEY (id);


--
-- TOC entry 4092 (class 2606 OID 47436)
-- Name: TstCaseComments TstCaseComments_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseComments"
    ADD CONSTRAINT "TstCaseComments_pkey" PRIMARY KEY (id);


--
-- TOC entry 4096 (class 2606 OID 47438)
-- Name: TstCaseExeStatusDefine TstCaseExeStatusDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseExeStatusDefine"
    ADD CONSTRAINT "TstCaseExeStatusDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 4098 (class 2606 OID 47440)
-- Name: TstCaseExeStatus TstCaseExeStatus_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseExeStatus"
    ADD CONSTRAINT "TstCaseExeStatus_pkey" PRIMARY KEY (id);


--
-- TOC entry 4101 (class 2606 OID 47442)
-- Name: TstCaseHistory TstCaseHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseHistory"
    ADD CONSTRAINT "TstCaseHistory_pkey" PRIMARY KEY (id);


--
-- TOC entry 4104 (class 2606 OID 47444)
-- Name: TstCaseInSuite TstCaseInSuite_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_pkey" PRIMARY KEY (id);


--
-- TOC entry 4119 (class 2606 OID 47446)
-- Name: TstCaseInTaskAttachment TstCaseInTaskAttachment_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskAttachment"
    ADD CONSTRAINT "TstCaseInTaskAttachment_pkey" PRIMARY KEY (id);


--
-- TOC entry 4123 (class 2606 OID 47448)
-- Name: TstCaseInTaskComments TstCaseInTaskComments_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskComments"
    ADD CONSTRAINT "TstCaseInTaskComments_pkey" PRIMARY KEY (id);


--
-- TOC entry 4127 (class 2606 OID 47450)
-- Name: TstCaseInTaskHistory TstCaseInTaskHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskHistory"
    ADD CONSTRAINT "TstCaseInTaskHistory_pkey" PRIMARY KEY (id);


--
-- TOC entry 4131 (class 2606 OID 47452)
-- Name: TstCaseInTaskIssue TstCaseInTaskIssue_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskIssue"
    ADD CONSTRAINT "TstCaseInTaskIssue_pkey" PRIMARY KEY (id);


--
-- TOC entry 4110 (class 2606 OID 47454)
-- Name: TstCaseInTask TstCaseInTask_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_pkey" PRIMARY KEY (id);


--
-- TOC entry 4136 (class 2606 OID 47456)
-- Name: TstCasePriorityDefine TstCasePriorityDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCasePriorityDefine"
    ADD CONSTRAINT "TstCasePriorityDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 4138 (class 2606 OID 47458)
-- Name: TstCasePriority TstCasePriority_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCasePriority"
    ADD CONSTRAINT "TstCasePriority_pkey" PRIMARY KEY (id);


--
-- TOC entry 4141 (class 2606 OID 47460)
-- Name: TstCaseStep TstCaseStep_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseStep"
    ADD CONSTRAINT "TstCaseStep_pkey" PRIMARY KEY (id);


--
-- TOC entry 4144 (class 2606 OID 47462)
-- Name: TstCaseTypeDefine TstCaseTypeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseTypeDefine"
    ADD CONSTRAINT "TstCaseTypeDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 4146 (class 2606 OID 47464)
-- Name: TstCaseType TstCaseType_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseType"
    ADD CONSTRAINT "TstCaseType_pkey" PRIMARY KEY (id);


--
-- TOC entry 4079 (class 2606 OID 47466)
-- Name: TstCase TstCase_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_pkey" PRIMARY KEY (id);


--
-- TOC entry 4149 (class 2606 OID 47468)
-- Name: TstDocument TstDocument_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstDocument"
    ADD CONSTRAINT "TstDocument_pkey" PRIMARY KEY (id);


--
-- TOC entry 4152 (class 2606 OID 47470)
-- Name: TstEmail TstEmail_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstEmail"
    ADD CONSTRAINT "TstEmail_pkey" PRIMARY KEY (id);


--
-- TOC entry 4154 (class 2606 OID 47472)
-- Name: TstEnv TstEnv_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstEnv"
    ADD CONSTRAINT "TstEnv_pkey" PRIMARY KEY (id);


--
-- TOC entry 4158 (class 2606 OID 47474)
-- Name: TstHistory TstHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstHistory"
    ADD CONSTRAINT "TstHistory_pkey" PRIMARY KEY (id);


--
-- TOC entry 4162 (class 2606 OID 47476)
-- Name: TstModule TstModule_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstModule"
    ADD CONSTRAINT "TstModule_pkey" PRIMARY KEY (id);


--
-- TOC entry 4165 (class 2606 OID 47478)
-- Name: TstMsg TstMsg_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstMsg"
    ADD CONSTRAINT "TstMsg_pkey" PRIMARY KEY (id);


--
-- TOC entry 4170 (class 2606 OID 47480)
-- Name: TstOrgGroup TstOrgGroup_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgGroup"
    ADD CONSTRAINT "TstOrgGroup_pkey" PRIMARY KEY (id);


--
-- TOC entry 4176 (class 2606 OID 47482)
-- Name: TstOrgPrivilegeDefine TstOrgPrivilegeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgPrivilegeDefine"
    ADD CONSTRAINT "TstOrgPrivilegeDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 4178 (class 2606 OID 47484)
-- Name: TstOrgRole TstOrgRole_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRole"
    ADD CONSTRAINT "TstOrgRole_pkey" PRIMARY KEY (id);


--
-- TOC entry 4168 (class 2606 OID 47486)
-- Name: TstOrg TstOrg_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrg"
    ADD CONSTRAINT "TstOrg_pkey" PRIMARY KEY (id);


--
-- TOC entry 4192 (class 2606 OID 47488)
-- Name: TstPlan TstPlan_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstPlan"
    ADD CONSTRAINT "TstPlan_pkey" PRIMARY KEY (id);


--
-- TOC entry 4205 (class 2606 OID 47490)
-- Name: TstProjectAccessHistory TstProjectAccessHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectAccessHistory"
    ADD CONSTRAINT "TstProjectAccessHistory_pkey" PRIMARY KEY (id);


--
-- TOC entry 4210 (class 2606 OID 47492)
-- Name: TstProjectPrivilegeDefine TstProjectPrivilegeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectPrivilegeDefine"
    ADD CONSTRAINT "TstProjectPrivilegeDefine_pkey" PRIMARY KEY (id);


--
-- TOC entry 4212 (class 2606 OID 47494)
-- Name: TstProjectRole TstProjectRole_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectRole"
    ADD CONSTRAINT "TstProjectRole_pkey" PRIMARY KEY (id);


--
-- TOC entry 4197 (class 2606 OID 47496)
-- Name: TstProject TstProject_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_pkey" PRIMARY KEY (id);


--
-- TOC entry 4221 (class 2606 OID 47498)
-- Name: TstSuite TstSuite_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstSuite"
    ADD CONSTRAINT "TstSuite_pkey" PRIMARY KEY (id);


--
-- TOC entry 4226 (class 2606 OID 47500)
-- Name: TstTask TstTask_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_pkey" PRIMARY KEY (id);


--
-- TOC entry 4235 (class 2606 OID 47502)
-- Name: TstThread TstThread_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstThread"
    ADD CONSTRAINT "TstThread_pkey" PRIMARY KEY (id);


--
-- TOC entry 4244 (class 2606 OID 47504)
-- Name: TstUserVerifyCode TstUserVerifyCode_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstUserVerifyCode"
    ADD CONSTRAINT "TstUserVerifyCode_pkey" PRIMARY KEY (id);


--
-- TOC entry 4239 (class 2606 OID 47506)
-- Name: TstUser TstUser_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstUser"
    ADD CONSTRAINT "TstUser_pkey" PRIMARY KEY (id);


--
-- TOC entry 4247 (class 2606 OID 47508)
-- Name: TstVer TstVer_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstVer"
    ADD CONSTRAINT "TstVer_pkey" PRIMARY KEY (id);


--
-- TOC entry 3889 (class 1259 OID 47509)
-- Name: fki_CustomFieldOptionDefine_fieldId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_CustomFieldOptionDefine_fieldId_fkey" ON public."CustomFieldOptionDefine" USING btree ("fieldId");


--
-- TOC entry 3885 (class 1259 OID 47510)
-- Name: fki_CustomFieldOption_fieldId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_CustomFieldOption_fieldId_fkey" ON public."CustomFieldOption" USING btree ("fieldId");


--
-- TOC entry 3886 (class 1259 OID 47511)
-- Name: fki_CustomFieldOption_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_CustomFieldOption_orgId_fkey" ON public."CustomFieldOption" USING btree ("orgId");


--
-- TOC entry 3876 (class 1259 OID 47512)
-- Name: fki_CustomField_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_CustomField_orgId_fkey" ON public."CustomField" USING btree ("orgId");


--
-- TOC entry 3894 (class 1259 OID 47513)
-- Name: fki_IsuAttachment_issueId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuAttachment_issueId_fkey" ON public."IsuAttachment" USING btree ("issueId");


--
-- TOC entry 3895 (class 1259 OID 47514)
-- Name: fki_IsuAttachment_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuAttachment_userId_fkey" ON public."IsuAttachment" USING btree ("userId");


--
-- TOC entry 3898 (class 1259 OID 47515)
-- Name: fki_IsuComments_issueId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuComments_issueId_fkey" ON public."IsuComments" USING btree ("issueId");


--
-- TOC entry 3899 (class 1259 OID 47516)
-- Name: fki_IsuComments_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuComments_userId_fkey" ON public."IsuComments" USING btree ("userId");


--
-- TOC entry 3903 (class 1259 OID 47517)
-- Name: fki_IsuCustomFieldSolutionFieldRelation_fieldId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuCustomFieldSolutionFieldRelation_fieldId_fkey" ON public."IsuCustomFieldSolutionFieldRelation" USING btree ("fieldId");


--
-- TOC entry 3904 (class 1259 OID 47518)
-- Name: fki_IsuCustomFieldSolutionFieldRelation_solutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuCustomFieldSolutionFieldRelation_solutionId_fkey" ON public."IsuCustomFieldSolutionFieldRelation" USING btree ("solutionId");


--
-- TOC entry 3905 (class 1259 OID 47519)
-- Name: fki_IsuCustomFieldSolutionProjectRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuCustomFieldSolutionProjectRelation_orgId_fkey" ON public."IsuCustomFieldSolutionProjectRelation" USING btree ("orgId");


--
-- TOC entry 3906 (class 1259 OID 47520)
-- Name: fki_IsuCustomFieldSolutionProjectRelation_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuCustomFieldSolutionProjectRelation_projectId_fkey" ON public."IsuCustomFieldSolutionProjectRelation" USING btree ("projectId");


--
-- TOC entry 3907 (class 1259 OID 47521)
-- Name: fki_IsuCustomFieldSolutionProjectRelation_solutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuCustomFieldSolutionProjectRelation_solutionId_fkey" ON public."IsuCustomFieldSolutionProjectRelation" USING btree ("solutionId");


--
-- TOC entry 3902 (class 1259 OID 47522)
-- Name: fki_IsuCustomFieldSolution_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuCustomFieldSolution_orgId_fkey" ON public."IsuCustomFieldSolution" USING btree ("orgId");


--
-- TOC entry 3910 (class 1259 OID 47523)
-- Name: fki_IsuDocument_issueId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuDocument_issueId_fkey" ON public."IsuDocument" USING btree ("issueId");


--
-- TOC entry 3911 (class 1259 OID 47524)
-- Name: fki_IsuDocument_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuDocument_userId_fkey" ON public."IsuDocument" USING btree ("userId");


--
-- TOC entry 3914 (class 1259 OID 47525)
-- Name: fki_IsuField_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuField_orgId_fkey" ON public."IsuField" USING btree ("orgId");


--
-- TOC entry 3921 (class 1259 OID 47526)
-- Name: fki_IsuHistory_issueId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuHistory_issueId_fkey" ON public."IsuHistory" USING btree ("issueId");


--
-- TOC entry 3938 (class 1259 OID 47527)
-- Name: fki_IsuIssueExt_pid_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssueExt_pid_fkey" ON public."IsuIssueExt" USING btree (pid);


--
-- TOC entry 3924 (class 1259 OID 47528)
-- Name: fki_IsuIssue_assigneeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_assigneeId_fkey" ON public."IsuIssue" USING btree ("assigneeId");


--
-- TOC entry 3925 (class 1259 OID 47529)
-- Name: fki_IsuIssue_creatorId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_creatorId_fkey" ON public."IsuIssue" USING btree ("creatorId");


--
-- TOC entry 3926 (class 1259 OID 47530)
-- Name: fki_IsuIssue_envId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_envId_fkey" ON public."IsuIssue" USING btree ("envId");


--
-- TOC entry 3927 (class 1259 OID 47531)
-- Name: fki_IsuIssue_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_orgId_fkey" ON public."IsuIssue" USING btree ("orgId");


--
-- TOC entry 3928 (class 1259 OID 47532)
-- Name: fki_IsuIssue_priorityId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_priorityId_fkey" ON public."IsuIssue" USING btree ("priorityId");


--
-- TOC entry 3929 (class 1259 OID 47533)
-- Name: fki_IsuIssue_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_projectId_fkey" ON public."IsuIssue" USING btree ("projectId");


--
-- TOC entry 3930 (class 1259 OID 47534)
-- Name: fki_IsuIssue_reporterId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_reporterId_fkey" ON public."IsuIssue" USING btree ("reporterId");


--
-- TOC entry 3931 (class 1259 OID 47535)
-- Name: fki_IsuIssue_resolutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_resolutionId_fkey" ON public."IsuIssue" USING btree ("resolutionId");


--
-- TOC entry 3932 (class 1259 OID 47536)
-- Name: fki_IsuIssue_statusId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_statusId_fkey" ON public."IsuIssue" USING btree ("statusId");


--
-- TOC entry 3933 (class 1259 OID 47537)
-- Name: fki_IsuIssue_typeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_typeId_fkey" ON public."IsuIssue" USING btree ("typeId");


--
-- TOC entry 3934 (class 1259 OID 47538)
-- Name: fki_IsuIssue_verId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuIssue_verId_fkey" ON public."IsuIssue" USING btree ("verId");


--
-- TOC entry 3941 (class 1259 OID 47539)
-- Name: fki_IsuLink_dictIssueId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuLink_dictIssueId_fkey" ON public."IsuLink" USING btree ("dictIssueId");


--
-- TOC entry 3942 (class 1259 OID 47540)
-- Name: fki_IsuLink_reasonId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuLink_reasonId_fkey" ON public."IsuLink" USING btree ("reasonId");


--
-- TOC entry 3943 (class 1259 OID 47541)
-- Name: fki_IsuLink_srcIssueId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuLink_srcIssueId_fkey" ON public."IsuLink" USING btree ("srcIssueId");


--
-- TOC entry 3948 (class 1259 OID 47542)
-- Name: fki_IsuNotification_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuNotification_orgId_fkey" ON public."IsuNotification" USING btree ("orgId");


--
-- TOC entry 3961 (class 1259 OID 47543)
-- Name: fki_IsuPageSolutionItem_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPageSolutionItem_orgId_fkey" ON public."IsuPageSolutionItem" USING btree ("orgId");


--
-- TOC entry 3962 (class 1259 OID 47544)
-- Name: fki_IsuPageSolutionItem_pageId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPageSolutionItem_pageId_fkey" ON public."IsuPageSolutionItem" USING btree ("pageId");


--
-- TOC entry 3963 (class 1259 OID 47545)
-- Name: fki_IsuPageSolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPageSolutionItem_solutionId_fkey" ON public."IsuPageSolutionItem" USING btree ("solutionId");


--
-- TOC entry 3964 (class 1259 OID 47546)
-- Name: fki_IsuPageSolutionItem_typeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPageSolutionItem_typeId_fkey" ON public."IsuPageSolutionItem" USING btree ("typeId");


--
-- TOC entry 3958 (class 1259 OID 47547)
-- Name: fki_IsuPageSolution_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPageSolution_orgId_fkey" ON public."IsuPageSolution" USING btree ("orgId");


--
-- TOC entry 3953 (class 1259 OID 47548)
-- Name: fki_IsuPage_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPage_orgId_fkey" ON public."IsuPage" USING btree ("orgId");


--
-- TOC entry 3973 (class 1259 OID 47549)
-- Name: fki_IsuPrioritySolutionItem_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPrioritySolutionItem_orgId_fkey" ON public."IsuPrioritySolutionItem" USING btree ("orgId");


--
-- TOC entry 3974 (class 1259 OID 47550)
-- Name: fki_IsuPrioritySolutionItem_priorityId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPrioritySolutionItem_priorityId_fkey" ON public."IsuPrioritySolutionItem" USING btree ("priorityId");


--
-- TOC entry 3975 (class 1259 OID 47551)
-- Name: fki_IsuPrioritySolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPrioritySolutionItem_solutionId_fkey" ON public."IsuPrioritySolutionItem" USING btree ("solutionId");


--
-- TOC entry 3972 (class 1259 OID 47552)
-- Name: fki_IsuPrioritySolution_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPrioritySolution_orgId_fkey" ON public."IsuPrioritySolution" USING btree ("orgId");


--
-- TOC entry 3967 (class 1259 OID 47553)
-- Name: fki_IsuPriority_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuPriority_orgId_fkey" ON public."IsuPriority" USING btree ("orgId");


--
-- TOC entry 3978 (class 1259 OID 47554)
-- Name: fki_IsuQuery_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuQuery_projectId_fkey" ON public."IsuQuery" USING btree ("projectId");


--
-- TOC entry 3979 (class 1259 OID 47555)
-- Name: fki_IsuQuery_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuQuery_userId_fkey" ON public."IsuQuery" USING btree ("userId");


--
-- TOC entry 3982 (class 1259 OID 47556)
-- Name: fki_IsuResolution_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuResolution_orgId_fkey" ON public."IsuResolution" USING btree ("orgId");


--
-- TOC entry 3993 (class 1259 OID 47557)
-- Name: fki_IsuSeveritySolutionItem_severityId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuSeveritySolutionItem_severityId_fkey" ON public."IsuSeveritySolutionItem" USING btree ("severityId");


--
-- TOC entry 3994 (class 1259 OID 47558)
-- Name: fki_IsuSeveritySolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuSeveritySolutionItem_solutionId_fkey" ON public."IsuSeveritySolutionItem" USING btree ("solutionId");


--
-- TOC entry 3992 (class 1259 OID 47559)
-- Name: fki_IsuSeveritySolution_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuSeveritySolution_orgId_fkey" ON public."IsuSeveritySolution" USING btree ("orgId");


--
-- TOC entry 3987 (class 1259 OID 47560)
-- Name: fki_IsuSeverity_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuSeverity_orgId_fkey" ON public."IsuSeverity" USING btree ("orgId");


--
-- TOC entry 4003 (class 1259 OID 47561)
-- Name: fki_IsuStatusDefine_categoryId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuStatusDefine_categoryId_fkey" ON public."IsuStatusDefine" USING btree ("categoryId");


--
-- TOC entry 3997 (class 1259 OID 47562)
-- Name: fki_IsuStatus_categoryId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuStatus_categoryId_fkey" ON public."IsuStatus" USING btree ("categoryId");


--
-- TOC entry 3998 (class 1259 OID 47563)
-- Name: fki_IsuStatus_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuStatus_orgId_fkey" ON public."IsuStatus" USING btree ("orgId");


--
-- TOC entry 4010 (class 1259 OID 47564)
-- Name: fki_IsuTagRelation_issueId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuTagRelation_issueId_fkey" ON public."IsuTagRelation" USING btree ("issueId");


--
-- TOC entry 4011 (class 1259 OID 47565)
-- Name: fki_IsuTagRelation_tagId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuTagRelation_tagId_fkey" ON public."IsuTagRelation" USING btree ("tagId");


--
-- TOC entry 4006 (class 1259 OID 47566)
-- Name: fki_IsuTag_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuTag_orgId_fkey" ON public."IsuTag" USING btree ("orgId");


--
-- TOC entry 4007 (class 1259 OID 47567)
-- Name: fki_IsuTag_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuTag_userId_fkey" ON public."IsuTag" USING btree ("userId");


--
-- TOC entry 4020 (class 1259 OID 47568)
-- Name: fki_IsuTypeSolutionItem_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuTypeSolutionItem_orgId_fkey" ON public."IsuTypeSolutionItem" USING btree ("orgId");


--
-- TOC entry 4021 (class 1259 OID 47569)
-- Name: fki_IsuTypeSolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuTypeSolutionItem_solutionId_fkey" ON public."IsuTypeSolutionItem" USING btree ("solutionId");


--
-- TOC entry 4022 (class 1259 OID 47570)
-- Name: fki_IsuTypeSolutionItem_typeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuTypeSolutionItem_typeId_fkey" ON public."IsuTypeSolutionItem" USING btree ("typeId");


--
-- TOC entry 4019 (class 1259 OID 47571)
-- Name: fki_IsuTypeSolution_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuTypeSolution_orgId_fkey" ON public."IsuTypeSolution" USING btree ("orgId");


--
-- TOC entry 4014 (class 1259 OID 47572)
-- Name: fki_IsuType_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuType_orgId_fkey" ON public."IsuType" USING btree ("orgId");


--
-- TOC entry 4025 (class 1259 OID 47573)
-- Name: fki_IsuWatch_issueId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWatch_issueId_fkey" ON public."IsuWatch" USING btree ("issueId");


--
-- TOC entry 4026 (class 1259 OID 47574)
-- Name: fki_IsuWatch_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWatch_userId_fkey" ON public."IsuWatch" USING btree ("userId");


--
-- TOC entry 4035 (class 1259 OID 47575)
-- Name: fki_IsuWorkflowSolutionItem_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowSolutionItem_orgId_fkey" ON public."IsuWorkflowSolutionItem" USING btree ("orgId");


--
-- TOC entry 4036 (class 1259 OID 47576)
-- Name: fki_IsuWorkflowSolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowSolutionItem_solutionId_fkey" ON public."IsuWorkflowSolutionItem" USING btree ("solutionId");


--
-- TOC entry 4037 (class 1259 OID 47577)
-- Name: fki_IsuWorkflowSolutionItem_typeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowSolutionItem_typeId_fkey" ON public."IsuWorkflowSolutionItem" USING btree ("typeId");


--
-- TOC entry 4038 (class 1259 OID 47578)
-- Name: fki_IsuWorkflowSolutionItem_workflowId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowSolutionItem_workflowId_fkey" ON public."IsuWorkflowSolutionItem" USING btree ("workflowId");


--
-- TOC entry 4032 (class 1259 OID 47579)
-- Name: fki_IsuWorkflowSolution_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowSolution_orgId_fkey" ON public."IsuWorkflowSolution" USING btree ("orgId");


--
-- TOC entry 4046 (class 1259 OID 47580)
-- Name: fki_IsuWorkflowStatusRelationDefine_statusId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowStatusRelationDefine_statusId_fkey" ON public."IsuWorkflowStatusRelationDefine" USING btree ("statusId");


--
-- TOC entry 4041 (class 1259 OID 47581)
-- Name: fki_IsuWorkflowStatusRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowStatusRelation_orgId_fkey" ON public."IsuWorkflowStatusRelation" USING btree ("orgId");


--
-- TOC entry 4042 (class 1259 OID 47582)
-- Name: fki_IsuWorkflowStatusRelation_statusId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowStatusRelation_statusId_fkey" ON public."IsuWorkflowStatusRelation" USING btree ("statusId");


--
-- TOC entry 4043 (class 1259 OID 47583)
-- Name: fki_IsuWorkflowStatusRelation_workflowId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowStatusRelation_workflowId_fkey" ON public."IsuWorkflowStatusRelation" USING btree ("workflowId");


--
-- TOC entry 4054 (class 1259 OID 47584)
-- Name: fki_IsuWorkflowTransitionDefine_dictStatusId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowTransitionDefine_dictStatusId_fkey" ON public."IsuWorkflowTransitionDefine" USING btree ("dictStatusId");


--
-- TOC entry 4055 (class 1259 OID 47585)
-- Name: fki_IsuWorkflowTransitionDefine_srcStatusId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowTransitionDefine_srcStatusId_fkey" ON public."IsuWorkflowTransitionDefine" USING btree ("srcStatusId");


--
-- TOC entry 4058 (class 1259 OID 47586)
-- Name: fki_IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_" ON public."IsuWorkflowTransitionProjectRoleRelation" USING btree ("workflowTransitionId");


--
-- TOC entry 4059 (class 1259 OID 47587)
-- Name: fki_IsuWorkflowTransitionProjectRoleRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowTransitionProjectRoleRelation_orgId_fkey" ON public."IsuWorkflowTransitionProjectRoleRelation" USING btree ("orgId");


--
-- TOC entry 4060 (class 1259 OID 47588)
-- Name: fki_IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey" ON public."IsuWorkflowTransitionProjectRoleRelation" USING btree ("projectRoleId");


--
-- TOC entry 4061 (class 1259 OID 47589)
-- Name: fki_IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey" ON public."IsuWorkflowTransitionProjectRoleRelation" USING btree ("workflowId");


--
-- TOC entry 4049 (class 1259 OID 47590)
-- Name: fki_IsuWorkflowTransition_actionPageId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowTransition_actionPageId_fkey" ON public."IsuWorkflowTransition" USING btree ("actionPageId");


--
-- TOC entry 4050 (class 1259 OID 47591)
-- Name: fki_IsuWorkflowTransition_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowTransition_orgId_fkey" ON public."IsuWorkflowTransition" USING btree ("orgId");


--
-- TOC entry 4051 (class 1259 OID 47592)
-- Name: fki_IsuWorkflowTransition_workflowId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflowTransition_workflowId_fkey" ON public."IsuWorkflowTransition" USING btree ("workflowId");


--
-- TOC entry 4029 (class 1259 OID 47593)
-- Name: fki_IsuWorkflow_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_IsuWorkflow_orgId_fkey" ON public."IsuWorkflow" USING btree ("orgId");


--
-- TOC entry 4066 (class 1259 OID 47594)
-- Name: fki_SysRolePrivilegeRelation_privilegeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_SysRolePrivilegeRelation_privilegeId_fkey" ON public."SysRolePrivilegeRelation" USING btree ("privilegeId");


--
-- TOC entry 4067 (class 1259 OID 47595)
-- Name: fki_SysRolePrivilegeRelation_roleId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_SysRolePrivilegeRelation_roleId_fkey" ON public."SysRolePrivilegeRelation" USING btree ("roleId");


--
-- TOC entry 4068 (class 1259 OID 47596)
-- Name: fki_SysRoleUserRelation_roleId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_SysRoleUserRelation_roleId_fkey" ON public."SysRoleUserRelation" USING btree ("roleId");


--
-- TOC entry 4069 (class 1259 OID 47597)
-- Name: fki_SysRoleUserRelation_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_SysRoleUserRelation_userId_fkey" ON public."SysRoleUserRelation" USING btree ("userId");


--
-- TOC entry 4076 (class 1259 OID 47598)
-- Name: fki_TstAlert_assigneeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstAlert_assigneeId_fkey" ON public."TstAlert" USING btree ("assigneeId");


--
-- TOC entry 4077 (class 1259 OID 47599)
-- Name: fki_TstAlert_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstAlert_userId_fkey" ON public."TstAlert" USING btree ("userId");


--
-- TOC entry 4089 (class 1259 OID 47600)
-- Name: fki_TstCaseAttachment_caseId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseAttachment_caseId_fkey" ON public."TstCaseAttachment" USING btree ("caseId");


--
-- TOC entry 4090 (class 1259 OID 47601)
-- Name: fki_TstCaseAttachment_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseAttachment_userId_fkey" ON public."TstCaseAttachment" USING btree ("userId");


--
-- TOC entry 4093 (class 1259 OID 47602)
-- Name: fki_TstCaseComments_caseId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseComments_caseId_fkey" ON public."TstCaseComments" USING btree ("caseId");


--
-- TOC entry 4094 (class 1259 OID 47603)
-- Name: fki_TstCaseComments_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseComments_userId_fkey" ON public."TstCaseComments" USING btree ("userId");


--
-- TOC entry 4099 (class 1259 OID 47604)
-- Name: fki_TstCaseExeStatus_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseExeStatus_orgId_fkey" ON public."TstCaseExeStatus" USING btree ("orgId");


--
-- TOC entry 4102 (class 1259 OID 47605)
-- Name: fki_TstCaseHistory_caseId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseHistory_caseId_fkey" ON public."TstCaseHistory" USING btree ("caseId");


--
-- TOC entry 4105 (class 1259 OID 47606)
-- Name: fki_TstCaseInSuite_caseId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInSuite_caseId_fkey" ON public."TstCaseInSuite" USING btree ("caseId");


--
-- TOC entry 4106 (class 1259 OID 47607)
-- Name: fki_TstCaseInSuite_pId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInSuite_pId_fkey" ON public."TstCaseInSuite" USING btree ("pId");


--
-- TOC entry 4107 (class 1259 OID 47608)
-- Name: fki_TstCaseInSuite_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInSuite_projectId_fkey" ON public."TstCaseInSuite" USING btree ("projectId");


--
-- TOC entry 4108 (class 1259 OID 47609)
-- Name: fki_TstCaseInSuite_suiteId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInSuite_suiteId_fkey" ON public."TstCaseInSuite" USING btree ("suiteId");


--
-- TOC entry 4120 (class 1259 OID 47610)
-- Name: fki_TstCaseInTaskAttachment_caseInTaskId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTaskAttachment_caseInTaskId_fkey" ON public."TstCaseInTaskAttachment" USING btree ("caseInTaskId");


--
-- TOC entry 4121 (class 1259 OID 47611)
-- Name: fki_TstCaseInTaskAttachment_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTaskAttachment_userId_fkey" ON public."TstCaseInTaskAttachment" USING btree ("userId");


--
-- TOC entry 4124 (class 1259 OID 47612)
-- Name: fki_TstCaseInTaskComments_caseInTaskId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTaskComments_caseInTaskId_fkey" ON public."TstCaseInTaskComments" USING btree ("caseInTaskId");


--
-- TOC entry 4125 (class 1259 OID 47613)
-- Name: fki_TstCaseInTaskComments_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTaskComments_userId_fkey" ON public."TstCaseInTaskComments" USING btree ("userId");


--
-- TOC entry 4128 (class 1259 OID 47614)
-- Name: fki_TstCaseInTaskHistory_caseId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTaskHistory_caseId_fkey" ON public."TstCaseInTaskHistory" USING btree ("caseId");


--
-- TOC entry 4129 (class 1259 OID 47615)
-- Name: fki_TstCaseInTaskHistory_caseInTaskId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTaskHistory_caseInTaskId_fkey" ON public."TstCaseInTaskHistory" USING btree ("caseInTaskId");


--
-- TOC entry 4132 (class 1259 OID 47616)
-- Name: fki_TstCaseInTaskIssue_caseInTaskId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTaskIssue_caseInTaskId_fkey" ON public."TstCaseInTaskIssue" USING btree ("caseInTaskId");


--
-- TOC entry 4133 (class 1259 OID 47617)
-- Name: fki_TstCaseInTaskIssue_issueId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTaskIssue_issueId_fkey" ON public."TstCaseInTaskIssue" USING btree ("issueId");


--
-- TOC entry 4134 (class 1259 OID 47618)
-- Name: fki_TstCaseInTaskIssue_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTaskIssue_userId_fkey" ON public."TstCaseInTaskIssue" USING btree ("userId");


--
-- TOC entry 4111 (class 1259 OID 47619)
-- Name: fki_TstCaseInTask_caseId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTask_caseId_fkey" ON public."TstCaseInTask" USING btree ("caseId");


--
-- TOC entry 4112 (class 1259 OID 47620)
-- Name: fki_TstCaseInTask_createBy_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTask_createBy_fkey" ON public."TstCaseInTask" USING btree ("createBy");


--
-- TOC entry 4113 (class 1259 OID 47621)
-- Name: fki_TstCaseInTask_exeBy_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTask_exeBy_fkey" ON public."TstCaseInTask" USING btree ("exeBy");


--
-- TOC entry 4114 (class 1259 OID 47622)
-- Name: fki_TstCaseInTask_pId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTask_pId_fkey" ON public."TstCaseInTask" USING btree ("pId");


--
-- TOC entry 4115 (class 1259 OID 47623)
-- Name: fki_TstCaseInTask_planId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTask_planId_fkey" ON public."TstCaseInTask" USING btree ("planId");


--
-- TOC entry 4116 (class 1259 OID 47624)
-- Name: fki_TstCaseInTask_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTask_projectId_fkey" ON public."TstCaseInTask" USING btree ("projectId");


--
-- TOC entry 4117 (class 1259 OID 47625)
-- Name: fki_TstCaseInTask_taskId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseInTask_taskId_fkey" ON public."TstCaseInTask" USING btree ("taskId");


--
-- TOC entry 4139 (class 1259 OID 47626)
-- Name: fki_TstCasePriority_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCasePriority_orgId_fkey" ON public."TstCasePriority" USING btree ("orgId");


--
-- TOC entry 4142 (class 1259 OID 47627)
-- Name: fki_TstCaseStep_caseId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseStep_caseId_fkey" ON public."TstCaseStep" USING btree ("caseId");


--
-- TOC entry 4147 (class 1259 OID 47628)
-- Name: fki_TstCaseType_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCaseType_orgId_fkey" ON public."TstCaseType" USING btree ("orgId");


--
-- TOC entry 4080 (class 1259 OID 47629)
-- Name: fki_TstCase_createById_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCase_createById_fkey" ON public."TstCase" USING btree ("createById");


--
-- TOC entry 4081 (class 1259 OID 47630)
-- Name: fki_TstCase_pId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCase_pId_fkey" ON public."TstCase" USING btree ("pId");


--
-- TOC entry 4082 (class 1259 OID 47631)
-- Name: fki_TstCase_priorityId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCase_priorityId_fkey" ON public."TstCase" USING btree ("priorityId");


--
-- TOC entry 4083 (class 1259 OID 47632)
-- Name: fki_TstCase_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCase_projectId_fkey" ON public."TstCase" USING btree ("projectId");


--
-- TOC entry 4084 (class 1259 OID 47633)
-- Name: fki_TstCase_typeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCase_typeId_fkey" ON public."TstCase" USING btree ("typeId");


--
-- TOC entry 4085 (class 1259 OID 47634)
-- Name: fki_TstCase_updateById_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstCase_updateById_fkey" ON public."TstCase" USING btree ("updateById");


--
-- TOC entry 4150 (class 1259 OID 47635)
-- Name: fki_TstDocument_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstDocument_userId_fkey" ON public."TstDocument" USING btree ("userId");


--
-- TOC entry 4155 (class 1259 OID 47636)
-- Name: fki_TstEnv_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstEnv_orgId_fkey" ON public."TstEnv" USING btree ("orgId");


--
-- TOC entry 4156 (class 1259 OID 47637)
-- Name: fki_TstEnv_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstEnv_projectId_fkey" ON public."TstEnv" USING btree ("projectId");


--
-- TOC entry 4159 (class 1259 OID 47638)
-- Name: fki_TstHistory_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstHistory_projectId_fkey" ON public."TstHistory" USING btree ("projectId");


--
-- TOC entry 4160 (class 1259 OID 47639)
-- Name: fki_TstHistory_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstHistory_userId_fkey" ON public."TstHistory" USING btree ("userId");


--
-- TOC entry 4163 (class 1259 OID 47640)
-- Name: fki_TstModule_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstModule_projectId_fkey" ON public."TstModule" USING btree ("projectId");


--
-- TOC entry 4166 (class 1259 OID 47641)
-- Name: fki_TstMsg_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstMsg_userId_fkey" ON public."TstMsg" USING btree ("userId");


--
-- TOC entry 4172 (class 1259 OID 47642)
-- Name: fki_TstOrgGroupUserRelation_orgGroupId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgGroupUserRelation_orgGroupId_fkey" ON public."TstOrgGroupUserRelation" USING btree ("orgGroupId");


--
-- TOC entry 4173 (class 1259 OID 47643)
-- Name: fki_TstOrgGroupUserRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgGroupUserRelation_orgId_fkey" ON public."TstOrgGroupUserRelation" USING btree ("orgId");


--
-- TOC entry 4174 (class 1259 OID 47644)
-- Name: fki_TstOrgGroupUserRelation_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgGroupUserRelation_userId_fkey" ON public."TstOrgGroupUserRelation" USING btree ("userId");


--
-- TOC entry 4171 (class 1259 OID 47645)
-- Name: fki_TstOrgGroup_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgGroup_orgId_fkey" ON public."TstOrgGroup" USING btree ("orgId");


--
-- TOC entry 4180 (class 1259 OID 47646)
-- Name: fki_TstOrgRoleGroupRelation_orgGroupId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRoleGroupRelation_orgGroupId_fkey" ON public."TstOrgRoleGroupRelation" USING btree ("orgGroupId");


--
-- TOC entry 4181 (class 1259 OID 47647)
-- Name: fki_TstOrgRoleGroupRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRoleGroupRelation_orgId_fkey" ON public."TstOrgRoleGroupRelation" USING btree ("orgId");


--
-- TOC entry 4182 (class 1259 OID 47648)
-- Name: fki_TstOrgRoleGroupRelation_orgRoleId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRoleGroupRelation_orgRoleId_fkey" ON public."TstOrgRoleGroupRelation" USING btree ("orgRoleId");


--
-- TOC entry 4183 (class 1259 OID 47649)
-- Name: fki_TstOrgRolePrivilegeRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRolePrivilegeRelation_orgId_fkey" ON public."TstOrgRolePrivilegeRelation" USING btree ("orgId");


--
-- TOC entry 4184 (class 1259 OID 47650)
-- Name: fki_TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey" ON public."TstOrgRolePrivilegeRelation" USING btree ("orgPrivilegeId");


--
-- TOC entry 4185 (class 1259 OID 47651)
-- Name: fki_TstOrgRolePrivilegeRelation_orgRoleId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRolePrivilegeRelation_orgRoleId_fkey" ON public."TstOrgRolePrivilegeRelation" USING btree ("orgRoleId");


--
-- TOC entry 4186 (class 1259 OID 47652)
-- Name: fki_TstOrgRoleUserRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRoleUserRelation_orgId_fkey" ON public."TstOrgRoleUserRelation" USING btree ("orgId");


--
-- TOC entry 4187 (class 1259 OID 47653)
-- Name: fki_TstOrgRoleUserRelation_orgRoleId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRoleUserRelation_orgRoleId_fkey" ON public."TstOrgRoleUserRelation" USING btree ("orgRoleId");


--
-- TOC entry 4188 (class 1259 OID 47654)
-- Name: fki_TstOrgRoleUserRelation_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRoleUserRelation_userId_fkey" ON public."TstOrgRoleUserRelation" USING btree ("userId");


--
-- TOC entry 4179 (class 1259 OID 47655)
-- Name: fki_TstOrgRole_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgRole_orgId_fkey" ON public."TstOrgRole" USING btree ("orgId");


--
-- TOC entry 4189 (class 1259 OID 47656)
-- Name: fki_TstOrgUserRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgUserRelation_orgId_fkey" ON public."TstOrgUserRelation" USING btree ("orgId");


--
-- TOC entry 4190 (class 1259 OID 47657)
-- Name: fki_TstOrgUserRelation_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstOrgUserRelation_userId_fkey" ON public."TstOrgUserRelation" USING btree ("userId");


--
-- TOC entry 4193 (class 1259 OID 47658)
-- Name: fki_TstPlan_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstPlan_projectId_fkey" ON public."TstPlan" USING btree ("projectId");


--
-- TOC entry 4194 (class 1259 OID 47659)
-- Name: fki_TstPlan_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstPlan_userId_fkey" ON public."TstPlan" USING btree ("userId");


--
-- TOC entry 4195 (class 1259 OID 47660)
-- Name: fki_TstPlan_verId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstPlan_verId_fkey" ON public."TstPlan" USING btree ("verId");


--
-- TOC entry 4206 (class 1259 OID 47661)
-- Name: fki_TstProjectAccessHistory_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectAccessHistory_orgId_fkey" ON public."TstProjectAccessHistory" USING btree ("orgId");


--
-- TOC entry 4207 (class 1259 OID 47662)
-- Name: fki_TstProjectAccessHistory_prjId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectAccessHistory_prjId_fkey" ON public."TstProjectAccessHistory" USING btree ("prjId");


--
-- TOC entry 4208 (class 1259 OID 47663)
-- Name: fki_TstProjectAccessHistory_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectAccessHistory_userId_fkey" ON public."TstProjectAccessHistory" USING btree ("userId");


--
-- TOC entry 4214 (class 1259 OID 47664)
-- Name: fki_TstProjectRoleEntityRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectRoleEntityRelation_orgId_fkey" ON public."TstProjectRoleEntityRelation" USING btree ("orgId");


--
-- TOC entry 4215 (class 1259 OID 47665)
-- Name: fki_TstProjectRoleEntityRelation_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectRoleEntityRelation_projectId_fkey" ON public."TstProjectRoleEntityRelation" USING btree ("projectId");


--
-- TOC entry 4216 (class 1259 OID 47666)
-- Name: fki_TstProjectRoleEntityRelation_projectRoleId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectRoleEntityRelation_projectRoleId_fkey" ON public."TstProjectRoleEntityRelation" USING btree ("projectRoleId");


--
-- TOC entry 4217 (class 1259 OID 47667)
-- Name: fki_TstProjectRolePriviledgeRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectRolePriviledgeRelation_orgId_fkey" ON public."TstProjectRolePriviledgeRelation" USING btree ("orgId");


--
-- TOC entry 4218 (class 1259 OID 47668)
-- Name: fki_TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_f; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_f" ON public."TstProjectRolePriviledgeRelation" USING btree ("projectPrivilegeDefineId");


--
-- TOC entry 4219 (class 1259 OID 47669)
-- Name: fki_TstProjectRolePriviledgeRelation_projectRoleId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectRolePriviledgeRelation_projectRoleId_fkey" ON public."TstProjectRolePriviledgeRelation" USING btree ("projectRoleId");


--
-- TOC entry 4213 (class 1259 OID 47670)
-- Name: fki_TstProjectRole_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProjectRole_orgId_fkey" ON public."TstProjectRole" USING btree ("orgId");


--
-- TOC entry 4198 (class 1259 OID 47671)
-- Name: fki_TstProject_issuePageSolutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProject_issuePageSolutionId_fkey" ON public."TstProject" USING btree ("issuePageSolutionId");


--
-- TOC entry 4199 (class 1259 OID 47672)
-- Name: fki_TstProject_issuePrioritySolutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProject_issuePrioritySolutionId_fkey" ON public."TstProject" USING btree ("issuePrioritySolutionId");


--
-- TOC entry 4200 (class 1259 OID 47673)
-- Name: fki_TstProject_issueTypeSolutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProject_issueTypeSolutionId_fkey" ON public."TstProject" USING btree ("issueTypeSolutionId");


--
-- TOC entry 4201 (class 1259 OID 47674)
-- Name: fki_TstProject_issueWorkflowSolutionId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProject_issueWorkflowSolutionId_fkey" ON public."TstProject" USING btree ("issueWorkflowSolutionId");


--
-- TOC entry 4202 (class 1259 OID 47675)
-- Name: fki_TstProject_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProject_orgId_fkey" ON public."TstProject" USING btree ("orgId");


--
-- TOC entry 4203 (class 1259 OID 47676)
-- Name: fki_TstProject_parentId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstProject_parentId_fkey" ON public."TstProject" USING btree ("parentId");


--
-- TOC entry 4222 (class 1259 OID 47677)
-- Name: fki_TstSuite_caseProjectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstSuite_caseProjectId_fkey" ON public."TstSuite" USING btree ("caseProjectId");


--
-- TOC entry 4223 (class 1259 OID 47678)
-- Name: fki_TstSuite_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstSuite_projectId_fkey" ON public."TstSuite" USING btree ("projectId");


--
-- TOC entry 4224 (class 1259 OID 47679)
-- Name: fki_TstSuite_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstSuite_userId_fkey" ON public."TstSuite" USING btree ("userId");


--
-- TOC entry 4232 (class 1259 OID 47680)
-- Name: fki_TstTaskAssigneeRelation_assigneeId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstTaskAssigneeRelation_assigneeId_fkey" ON public."TstTaskAssigneeRelation" USING btree ("assigneeId");


--
-- TOC entry 4233 (class 1259 OID 47681)
-- Name: fki_TstTaskAssigneeRelation_taskId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstTaskAssigneeRelation_taskId_fkey" ON public."TstTaskAssigneeRelation" USING btree ("taskId");


--
-- TOC entry 4227 (class 1259 OID 47682)
-- Name: fki_TstTask_caseProjectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstTask_caseProjectId_fkey" ON public."TstTask" USING btree ("caseProjectId");


--
-- TOC entry 4228 (class 1259 OID 47683)
-- Name: fki_TstTask_envId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstTask_envId_fkey" ON public."TstTask" USING btree ("envId");


--
-- TOC entry 4229 (class 1259 OID 47684)
-- Name: fki_TstTask_planId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstTask_planId_fkey" ON public."TstTask" USING btree ("planId");


--
-- TOC entry 4230 (class 1259 OID 47685)
-- Name: fki_TstTask_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstTask_projectId_fkey" ON public."TstTask" USING btree ("projectId");


--
-- TOC entry 4231 (class 1259 OID 47686)
-- Name: fki_TstTask_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstTask_userId_fkey" ON public."TstTask" USING btree ("userId");


--
-- TOC entry 4236 (class 1259 OID 47687)
-- Name: fki_TstThread_authorId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstThread_authorId_fkey" ON public."TstThread" USING btree ("authorId");


--
-- TOC entry 4237 (class 1259 OID 47688)
-- Name: fki_TstThread_parentId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstThread_parentId_fkey" ON public."TstThread" USING btree ("parentId");


--
-- TOC entry 4242 (class 1259 OID 47689)
-- Name: fki_TstUserSettings_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstUserSettings_userId_fkey" ON public."TstUserSettings" USING btree ("userId");


--
-- TOC entry 4245 (class 1259 OID 47690)
-- Name: fki_TstUserVerifyCode_userId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstUserVerifyCode_userId_fkey" ON public."TstUserVerifyCode" USING btree ("userId");


--
-- TOC entry 4240 (class 1259 OID 47691)
-- Name: fki_TstUser_defaultOrgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstUser_defaultOrgId_fkey" ON public."TstUser" USING btree ("defaultOrgId");


--
-- TOC entry 4241 (class 1259 OID 47692)
-- Name: fki_TstUser_defaultPrjId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstUser_defaultPrjId_fkey" ON public."TstUser" USING btree ("defaultPrjId");


--
-- TOC entry 4248 (class 1259 OID 47693)
-- Name: fki_TstVer_orgId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstVer_orgId_fkey" ON public."TstVer" USING btree ("orgId");


--
-- TOC entry 4249 (class 1259 OID 47694)
-- Name: fki_TstVer_projectId_fkey; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX "fki_TstVer_projectId_fkey" ON public."TstVer" USING btree ("projectId");


--
-- TOC entry 3935 (class 1259 OID 47695)
-- Name: idx_isu_issue_extprop; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_isu_issue_extprop ON public."IsuIssue" USING gin ("extProp" jsonb_path_ops);


--
-- TOC entry 4086 (class 1259 OID 47696)
-- Name: idx_test_case_extprop; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_test_case_extprop ON public."TstCase" USING gin ("extProp" jsonb_path_ops);


--
-- TOC entry 4436 (class 2620 OID 48645)
-- Name: IsuIssue issue_tsvector_update_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER issue_tsvector_update_trigger AFTER INSERT OR UPDATE OF title, tag, "extProp", descr ON public."IsuIssue" FOR EACH ROW EXECUTE PROCEDURE public.update_issue_tsv_content();


--
-- TOC entry 4253 (class 2606 OID 47697)
-- Name: CustomFieldOptionDefine CustomFieldOptionDefine_fieldId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldOptionDefine"
    ADD CONSTRAINT "CustomFieldOptionDefine_fieldId_fkey" FOREIGN KEY ("fieldId") REFERENCES public."CustomFieldDefine"(id);


--
-- TOC entry 4251 (class 2606 OID 47702)
-- Name: CustomFieldOption CustomFieldOption_fieldId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldOption"
    ADD CONSTRAINT "CustomFieldOption_fieldId_fkey" FOREIGN KEY ("fieldId") REFERENCES public."CustomField"(id);


--
-- TOC entry 4252 (class 2606 OID 47707)
-- Name: CustomFieldOption CustomFieldOption_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomFieldOption"
    ADD CONSTRAINT "CustomFieldOption_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4250 (class 2606 OID 47712)
-- Name: CustomField CustomField_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."CustomField"
    ADD CONSTRAINT "CustomField_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4254 (class 2606 OID 47717)
-- Name: IsuAttachment IsuAttachment_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuAttachment"
    ADD CONSTRAINT "IsuAttachment_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4255 (class 2606 OID 47722)
-- Name: IsuAttachment IsuAttachment_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuAttachment"
    ADD CONSTRAINT "IsuAttachment_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4256 (class 2606 OID 47727)
-- Name: IsuComments IsuComments_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuComments"
    ADD CONSTRAINT "IsuComments_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4257 (class 2606 OID 47732)
-- Name: IsuComments IsuComments_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuComments"
    ADD CONSTRAINT "IsuComments_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4259 (class 2606 OID 47737)
-- Name: IsuCustomFieldSolutionFieldRelation IsuCustomFieldSolutionFieldRelation_fieldId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionFieldRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionFieldRelation_fieldId_fkey" FOREIGN KEY ("fieldId") REFERENCES public."CustomField"(id);


--
-- TOC entry 4260 (class 2606 OID 47742)
-- Name: IsuCustomFieldSolutionFieldRelation IsuCustomFieldSolutionFieldRelation_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionFieldRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionFieldRelation_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuCustomFieldSolution"(id);


--
-- TOC entry 4261 (class 2606 OID 47747)
-- Name: IsuCustomFieldSolutionProjectRelation IsuCustomFieldSolutionProjectRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionProjectRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionProjectRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4262 (class 2606 OID 47752)
-- Name: IsuCustomFieldSolutionProjectRelation IsuCustomFieldSolutionProjectRelation_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionProjectRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionProjectRelation_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4263 (class 2606 OID 47757)
-- Name: IsuCustomFieldSolutionProjectRelation IsuCustomFieldSolutionProjectRelation_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionProjectRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionProjectRelation_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuCustomFieldSolution"(id);


--
-- TOC entry 4258 (class 2606 OID 47762)
-- Name: IsuCustomFieldSolution IsuCustomFieldSolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuCustomFieldSolution"
    ADD CONSTRAINT "IsuCustomFieldSolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4264 (class 2606 OID 47767)
-- Name: IsuDocument IsuDocument_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuDocument"
    ADD CONSTRAINT "IsuDocument_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4265 (class 2606 OID 47772)
-- Name: IsuDocument IsuDocument_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuDocument"
    ADD CONSTRAINT "IsuDocument_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4266 (class 2606 OID 47777)
-- Name: IsuField IsuField_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuField"
    ADD CONSTRAINT "IsuField_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4267 (class 2606 OID 47782)
-- Name: IsuHistory IsuHistory_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuHistory"
    ADD CONSTRAINT "IsuHistory_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4279 (class 2606 OID 47787)
-- Name: IsuIssueExt IsuIssueExt_pid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssueExt"
    ADD CONSTRAINT "IsuIssueExt_pid_fkey" FOREIGN KEY (pid) REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4268 (class 2606 OID 47792)
-- Name: IsuIssue IsuIssue_assigneeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_assigneeId_fkey" FOREIGN KEY ("assigneeId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4269 (class 2606 OID 47797)
-- Name: IsuIssue IsuIssue_creatorId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_creatorId_fkey" FOREIGN KEY ("creatorId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4270 (class 2606 OID 47802)
-- Name: IsuIssue IsuIssue_envId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_envId_fkey" FOREIGN KEY ("envId") REFERENCES public."TstEnv"(id);


--
-- TOC entry 4271 (class 2606 OID 47807)
-- Name: IsuIssue IsuIssue_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4272 (class 2606 OID 47812)
-- Name: IsuIssue IsuIssue_priorityId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_priorityId_fkey" FOREIGN KEY ("priorityId") REFERENCES public."IsuPriority"(id);


--
-- TOC entry 4273 (class 2606 OID 47817)
-- Name: IsuIssue IsuIssue_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4274 (class 2606 OID 47822)
-- Name: IsuIssue IsuIssue_reporterId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_reporterId_fkey" FOREIGN KEY ("reporterId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4275 (class 2606 OID 47827)
-- Name: IsuIssue IsuIssue_resolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_resolutionId_fkey" FOREIGN KEY ("resolutionId") REFERENCES public."IsuResolution"(id);


--
-- TOC entry 4276 (class 2606 OID 47832)
-- Name: IsuIssue IsuIssue_statusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_statusId_fkey" FOREIGN KEY ("statusId") REFERENCES public."IsuStatus"(id);


--
-- TOC entry 4277 (class 2606 OID 47837)
-- Name: IsuIssue IsuIssue_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."IsuType"(id);


--
-- TOC entry 4278 (class 2606 OID 47842)
-- Name: IsuIssue IsuIssue_verId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_verId_fkey" FOREIGN KEY ("verId") REFERENCES public."TstVer"(id);


--
-- TOC entry 4280 (class 2606 OID 47847)
-- Name: IsuLink IsuLink_dictIssueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuLink"
    ADD CONSTRAINT "IsuLink_dictIssueId_fkey" FOREIGN KEY ("dictIssueId") REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4281 (class 2606 OID 47852)
-- Name: IsuLink IsuLink_reasonId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuLink"
    ADD CONSTRAINT "IsuLink_reasonId_fkey" FOREIGN KEY ("reasonId") REFERENCES public."IsuLinkReasonDefine"(id);


--
-- TOC entry 4282 (class 2606 OID 47857)
-- Name: IsuLink IsuLink_srcIssueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuLink"
    ADD CONSTRAINT "IsuLink_srcIssueId_fkey" FOREIGN KEY ("srcIssueId") REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4283 (class 2606 OID 47862)
-- Name: IsuNotification IsuNotification_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuNotification"
    ADD CONSTRAINT "IsuNotification_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4286 (class 2606 OID 47867)
-- Name: IsuPageSolutionItem IsuPageSolutionItem_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4287 (class 2606 OID 47872)
-- Name: IsuPageSolutionItem IsuPageSolutionItem_pageId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_pageId_fkey" FOREIGN KEY ("pageId") REFERENCES public."IsuPage"(id);


--
-- TOC entry 4288 (class 2606 OID 47877)
-- Name: IsuPageSolutionItem IsuPageSolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuPageSolution"(id);


--
-- TOC entry 4289 (class 2606 OID 47882)
-- Name: IsuPageSolutionItem IsuPageSolutionItem_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."IsuType"(id);


--
-- TOC entry 4285 (class 2606 OID 47887)
-- Name: IsuPageSolution IsuPageSolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPageSolution"
    ADD CONSTRAINT "IsuPageSolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4284 (class 2606 OID 47892)
-- Name: IsuPage IsuPage_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPage"
    ADD CONSTRAINT "IsuPage_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4292 (class 2606 OID 47897)
-- Name: IsuPrioritySolutionItem IsuPrioritySolutionItem_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPrioritySolutionItem"
    ADD CONSTRAINT "IsuPrioritySolutionItem_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4293 (class 2606 OID 47902)
-- Name: IsuPrioritySolutionItem IsuPrioritySolutionItem_priorityId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPrioritySolutionItem"
    ADD CONSTRAINT "IsuPrioritySolutionItem_priorityId_fkey" FOREIGN KEY ("priorityId") REFERENCES public."IsuPriority"(id);


--
-- TOC entry 4294 (class 2606 OID 47907)
-- Name: IsuPrioritySolutionItem IsuPrioritySolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPrioritySolutionItem"
    ADD CONSTRAINT "IsuPrioritySolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuPrioritySolution"(id);


--
-- TOC entry 4291 (class 2606 OID 47912)
-- Name: IsuPrioritySolution IsuPrioritySolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPrioritySolution"
    ADD CONSTRAINT "IsuPrioritySolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4290 (class 2606 OID 47917)
-- Name: IsuPriority IsuPriority_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuPriority"
    ADD CONSTRAINT "IsuPriority_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4295 (class 2606 OID 47922)
-- Name: IsuQuery IsuQuery_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuQuery"
    ADD CONSTRAINT "IsuQuery_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4296 (class 2606 OID 47927)
-- Name: IsuQuery IsuQuery_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuQuery"
    ADD CONSTRAINT "IsuQuery_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4297 (class 2606 OID 47932)
-- Name: IsuResolution IsuResolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuResolution"
    ADD CONSTRAINT "IsuResolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4300 (class 2606 OID 47937)
-- Name: IsuSeveritySolutionItem IsuSeveritySolutionItem_severityId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeveritySolutionItem"
    ADD CONSTRAINT "IsuSeveritySolutionItem_severityId_fkey" FOREIGN KEY ("severityId") REFERENCES public."IsuSeverity"(id);


--
-- TOC entry 4301 (class 2606 OID 47942)
-- Name: IsuSeveritySolutionItem IsuSeveritySolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeveritySolutionItem"
    ADD CONSTRAINT "IsuSeveritySolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuSeveritySolution"(id);


--
-- TOC entry 4299 (class 2606 OID 47947)
-- Name: IsuSeveritySolution IsuSeveritySolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeveritySolution"
    ADD CONSTRAINT "IsuSeveritySolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4298 (class 2606 OID 47952)
-- Name: IsuSeverity IsuSeverity_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuSeverity"
    ADD CONSTRAINT "IsuSeverity_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4304 (class 2606 OID 47957)
-- Name: IsuStatusDefine IsuStatusDefine_categoryId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuStatusDefine"
    ADD CONSTRAINT "IsuStatusDefine_categoryId_fkey" FOREIGN KEY ("categoryId") REFERENCES public."IsuStatusCategoryDefine"(id);


--
-- TOC entry 4302 (class 2606 OID 47962)
-- Name: IsuStatus IsuStatus_categoryId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuStatus"
    ADD CONSTRAINT "IsuStatus_categoryId_fkey" FOREIGN KEY ("categoryId") REFERENCES public."IsuStatusCategoryDefine"(id);


--
-- TOC entry 4303 (class 2606 OID 47967)
-- Name: IsuStatus IsuStatus_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuStatus"
    ADD CONSTRAINT "IsuStatus_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4307 (class 2606 OID 47972)
-- Name: IsuTagRelation IsuTagRelation_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTagRelation"
    ADD CONSTRAINT "IsuTagRelation_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4308 (class 2606 OID 47977)
-- Name: IsuTagRelation IsuTagRelation_tagId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTagRelation"
    ADD CONSTRAINT "IsuTagRelation_tagId_fkey" FOREIGN KEY ("tagId") REFERENCES public."IsuTag"(id);


--
-- TOC entry 4305 (class 2606 OID 47982)
-- Name: IsuTag IsuTag_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTag"
    ADD CONSTRAINT "IsuTag_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4306 (class 2606 OID 47987)
-- Name: IsuTag IsuTag_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTag"
    ADD CONSTRAINT "IsuTag_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4311 (class 2606 OID 47992)
-- Name: IsuTypeSolutionItem IsuTypeSolutionItem_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTypeSolutionItem"
    ADD CONSTRAINT "IsuTypeSolutionItem_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4312 (class 2606 OID 47997)
-- Name: IsuTypeSolutionItem IsuTypeSolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTypeSolutionItem"
    ADD CONSTRAINT "IsuTypeSolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuTypeSolution"(id);


--
-- TOC entry 4313 (class 2606 OID 48002)
-- Name: IsuTypeSolutionItem IsuTypeSolutionItem_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTypeSolutionItem"
    ADD CONSTRAINT "IsuTypeSolutionItem_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."IsuType"(id);


--
-- TOC entry 4310 (class 2606 OID 48007)
-- Name: IsuTypeSolution IsuTypeSolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuTypeSolution"
    ADD CONSTRAINT "IsuTypeSolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4309 (class 2606 OID 48012)
-- Name: IsuType IsuType_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuType"
    ADD CONSTRAINT "IsuType_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4314 (class 2606 OID 48017)
-- Name: IsuWatch IsuWatch_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWatch"
    ADD CONSTRAINT "IsuWatch_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4315 (class 2606 OID 48022)
-- Name: IsuWatch IsuWatch_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWatch"
    ADD CONSTRAINT "IsuWatch_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4318 (class 2606 OID 48027)
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4319 (class 2606 OID 48032)
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuWorkflowSolution"(id);


--
-- TOC entry 4320 (class 2606 OID 48037)
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."IsuType"(id);


--
-- TOC entry 4321 (class 2606 OID 48042)
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_workflowId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES public."IsuWorkflow"(id);


--
-- TOC entry 4317 (class 2606 OID 48047)
-- Name: IsuWorkflowSolution IsuWorkflowSolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowSolution"
    ADD CONSTRAINT "IsuWorkflowSolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4325 (class 2606 OID 48052)
-- Name: IsuWorkflowStatusRelationDefine IsuWorkflowStatusRelationDefine_statusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelationDefine"
    ADD CONSTRAINT "IsuWorkflowStatusRelationDefine_statusId_fkey" FOREIGN KEY ("statusId") REFERENCES public."IsuStatusDefine"(id);


--
-- TOC entry 4322 (class 2606 OID 48057)
-- Name: IsuWorkflowStatusRelation IsuWorkflowStatusRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation"
    ADD CONSTRAINT "IsuWorkflowStatusRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4323 (class 2606 OID 48062)
-- Name: IsuWorkflowStatusRelation IsuWorkflowStatusRelation_statusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation"
    ADD CONSTRAINT "IsuWorkflowStatusRelation_statusId_fkey" FOREIGN KEY ("statusId") REFERENCES public."IsuStatus"(id);


--
-- TOC entry 4324 (class 2606 OID 48067)
-- Name: IsuWorkflowStatusRelation IsuWorkflowStatusRelation_workflowId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation"
    ADD CONSTRAINT "IsuWorkflowStatusRelation_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES public."IsuWorkflow"(id);


--
-- TOC entry 4329 (class 2606 OID 48072)
-- Name: IsuWorkflowTransitionDefine IsuWorkflowTransitionDefine_dictStatusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionDefine"
    ADD CONSTRAINT "IsuWorkflowTransitionDefine_dictStatusId_fkey" FOREIGN KEY ("dictStatusId") REFERENCES public."IsuStatusDefine"(id);


--
-- TOC entry 4330 (class 2606 OID 48077)
-- Name: IsuWorkflowTransitionDefine IsuWorkflowTransitionDefine_srcStatusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionDefine"
    ADD CONSTRAINT "IsuWorkflowTransitionDefine_srcStatusId_fkey" FOREIGN KEY ("srcStatusId") REFERENCES public."IsuStatusDefine"(id);


--
-- TOC entry 4331 (class 2606 OID 48082)
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_fkey" FOREIGN KEY ("workflowTransitionId") REFERENCES public."IsuWorkflowTransition"(id);


--
-- TOC entry 4332 (class 2606 OID 48087)
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4333 (class 2606 OID 48092)
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey" FOREIGN KEY ("projectRoleId") REFERENCES public."TstProjectRole"(id);


--
-- TOC entry 4334 (class 2606 OID 48097)
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES public."IsuWorkflow"(id);


--
-- TOC entry 4326 (class 2606 OID 48102)
-- Name: IsuWorkflowTransition IsuWorkflowTransition_actionPageId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransition"
    ADD CONSTRAINT "IsuWorkflowTransition_actionPageId_fkey" FOREIGN KEY ("actionPageId") REFERENCES public."IsuPage"(id);


--
-- TOC entry 4327 (class 2606 OID 48107)
-- Name: IsuWorkflowTransition IsuWorkflowTransition_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransition"
    ADD CONSTRAINT "IsuWorkflowTransition_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4328 (class 2606 OID 48112)
-- Name: IsuWorkflowTransition IsuWorkflowTransition_workflowId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflowTransition"
    ADD CONSTRAINT "IsuWorkflowTransition_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES public."IsuWorkflow"(id);


--
-- TOC entry 4316 (class 2606 OID 48117)
-- Name: IsuWorkflow IsuWorkflow_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."IsuWorkflow"
    ADD CONSTRAINT "IsuWorkflow_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4335 (class 2606 OID 48122)
-- Name: SysRolePrivilegeRelation SysRolePrivilegeRelation_privilegeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysRolePrivilegeRelation"
    ADD CONSTRAINT "SysRolePrivilegeRelation_privilegeId_fkey" FOREIGN KEY ("privilegeId") REFERENCES public."SysPrivilege"(id);


--
-- TOC entry 4336 (class 2606 OID 48127)
-- Name: SysRolePrivilegeRelation SysRolePrivilegeRelation_roleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysRolePrivilegeRelation"
    ADD CONSTRAINT "SysRolePrivilegeRelation_roleId_fkey" FOREIGN KEY ("roleId") REFERENCES public."SysRole"(id);


--
-- TOC entry 4337 (class 2606 OID 48132)
-- Name: SysRoleUserRelation SysRoleUserRelation_roleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysRoleUserRelation"
    ADD CONSTRAINT "SysRoleUserRelation_roleId_fkey" FOREIGN KEY ("roleId") REFERENCES public."SysRole"(id);


--
-- TOC entry 4338 (class 2606 OID 48137)
-- Name: SysRoleUserRelation SysRoleUserRelation_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."SysRoleUserRelation"
    ADD CONSTRAINT "SysRoleUserRelation_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4339 (class 2606 OID 48142)
-- Name: TstAlert TstAlert_assigneeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstAlert"
    ADD CONSTRAINT "TstAlert_assigneeId_fkey" FOREIGN KEY ("assigneeId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4340 (class 2606 OID 48147)
-- Name: TstAlert TstAlert_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstAlert"
    ADD CONSTRAINT "TstAlert_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4347 (class 2606 OID 48152)
-- Name: TstCaseAttachment TstCaseAttachment_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseAttachment"
    ADD CONSTRAINT "TstCaseAttachment_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4348 (class 2606 OID 48157)
-- Name: TstCaseAttachment TstCaseAttachment_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseAttachment"
    ADD CONSTRAINT "TstCaseAttachment_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4349 (class 2606 OID 48162)
-- Name: TstCaseComments TstCaseComments_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseComments"
    ADD CONSTRAINT "TstCaseComments_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4350 (class 2606 OID 48167)
-- Name: TstCaseComments TstCaseComments_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseComments"
    ADD CONSTRAINT "TstCaseComments_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4351 (class 2606 OID 48172)
-- Name: TstCaseExeStatus TstCaseExeStatus_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseExeStatus"
    ADD CONSTRAINT "TstCaseExeStatus_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4352 (class 2606 OID 48177)
-- Name: TstCaseHistory TstCaseHistory_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseHistory"
    ADD CONSTRAINT "TstCaseHistory_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4353 (class 2606 OID 48182)
-- Name: TstCaseInSuite TstCaseInSuite_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4354 (class 2606 OID 48187)
-- Name: TstCaseInSuite TstCaseInSuite_pId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_pId_fkey" FOREIGN KEY ("pId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4355 (class 2606 OID 48192)
-- Name: TstCaseInSuite TstCaseInSuite_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4356 (class 2606 OID 48197)
-- Name: TstCaseInSuite TstCaseInSuite_suiteId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_suiteId_fkey" FOREIGN KEY ("suiteId") REFERENCES public."TstSuite"(id);


--
-- TOC entry 4364 (class 2606 OID 48202)
-- Name: TstCaseInTaskAttachment TstCaseInTaskAttachment_caseInTaskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskAttachment"
    ADD CONSTRAINT "TstCaseInTaskAttachment_caseInTaskId_fkey" FOREIGN KEY ("caseInTaskId") REFERENCES public."TstCaseInTask"(id);


--
-- TOC entry 4365 (class 2606 OID 48207)
-- Name: TstCaseInTaskAttachment TstCaseInTaskAttachment_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskAttachment"
    ADD CONSTRAINT "TstCaseInTaskAttachment_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4366 (class 2606 OID 48212)
-- Name: TstCaseInTaskComments TstCaseInTaskComments_caseInTaskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskComments"
    ADD CONSTRAINT "TstCaseInTaskComments_caseInTaskId_fkey" FOREIGN KEY ("caseInTaskId") REFERENCES public."TstCaseInTask"(id);


--
-- TOC entry 4367 (class 2606 OID 48217)
-- Name: TstCaseInTaskComments TstCaseInTaskComments_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskComments"
    ADD CONSTRAINT "TstCaseInTaskComments_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4368 (class 2606 OID 48222)
-- Name: TstCaseInTaskHistory TstCaseInTaskHistory_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskHistory"
    ADD CONSTRAINT "TstCaseInTaskHistory_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4369 (class 2606 OID 48227)
-- Name: TstCaseInTaskHistory TstCaseInTaskHistory_caseInTaskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskHistory"
    ADD CONSTRAINT "TstCaseInTaskHistory_caseInTaskId_fkey" FOREIGN KEY ("caseInTaskId") REFERENCES public."TstCaseInTask"(id);


--
-- TOC entry 4370 (class 2606 OID 48232)
-- Name: TstCaseInTaskIssue TstCaseInTaskIssue_caseInTaskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskIssue"
    ADD CONSTRAINT "TstCaseInTaskIssue_caseInTaskId_fkey" FOREIGN KEY ("caseInTaskId") REFERENCES public."TstCaseInTask"(id);


--
-- TOC entry 4371 (class 2606 OID 48237)
-- Name: TstCaseInTaskIssue TstCaseInTaskIssue_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskIssue"
    ADD CONSTRAINT "TstCaseInTaskIssue_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- TOC entry 4372 (class 2606 OID 48242)
-- Name: TstCaseInTaskIssue TstCaseInTaskIssue_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTaskIssue"
    ADD CONSTRAINT "TstCaseInTaskIssue_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4357 (class 2606 OID 48247)
-- Name: TstCaseInTask TstCaseInTask_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4358 (class 2606 OID 48252)
-- Name: TstCaseInTask TstCaseInTask_createBy_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_createBy_fkey" FOREIGN KEY ("createBy") REFERENCES public."TstUser"(id);


--
-- TOC entry 4359 (class 2606 OID 48257)
-- Name: TstCaseInTask TstCaseInTask_exeBy_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_exeBy_fkey" FOREIGN KEY ("exeBy") REFERENCES public."TstUser"(id);


--
-- TOC entry 4360 (class 2606 OID 48262)
-- Name: TstCaseInTask TstCaseInTask_pId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_pId_fkey" FOREIGN KEY ("pId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4361 (class 2606 OID 48267)
-- Name: TstCaseInTask TstCaseInTask_planId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_planId_fkey" FOREIGN KEY ("planId") REFERENCES public."TstPlan"(id);


--
-- TOC entry 4362 (class 2606 OID 48272)
-- Name: TstCaseInTask TstCaseInTask_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4363 (class 2606 OID 48277)
-- Name: TstCaseInTask TstCaseInTask_taskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_taskId_fkey" FOREIGN KEY ("taskId") REFERENCES public."TstTask"(id);


--
-- TOC entry 4373 (class 2606 OID 48282)
-- Name: TstCasePriority TstCasePriority_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCasePriority"
    ADD CONSTRAINT "TstCasePriority_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4374 (class 2606 OID 48287)
-- Name: TstCaseStep TstCaseStep_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseStep"
    ADD CONSTRAINT "TstCaseStep_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4375 (class 2606 OID 48292)
-- Name: TstCaseType TstCaseType_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCaseType"
    ADD CONSTRAINT "TstCaseType_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4341 (class 2606 OID 48297)
-- Name: TstCase TstCase_createById_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_createById_fkey" FOREIGN KEY ("createById") REFERENCES public."TstUser"(id);


--
-- TOC entry 4342 (class 2606 OID 48302)
-- Name: TstCase TstCase_pId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_pId_fkey" FOREIGN KEY ("pId") REFERENCES public."TstCase"(id);


--
-- TOC entry 4343 (class 2606 OID 48307)
-- Name: TstCase TstCase_priorityId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_priorityId_fkey" FOREIGN KEY ("priorityId") REFERENCES public."TstCasePriority"(id);


--
-- TOC entry 4344 (class 2606 OID 48312)
-- Name: TstCase TstCase_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4345 (class 2606 OID 48317)
-- Name: TstCase TstCase_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."TstCaseType"(id);


--
-- TOC entry 4346 (class 2606 OID 48322)
-- Name: TstCase TstCase_updateById_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_updateById_fkey" FOREIGN KEY ("updateById") REFERENCES public."TstUser"(id);


--
-- TOC entry 4376 (class 2606 OID 48327)
-- Name: TstDocument TstDocument_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstDocument"
    ADD CONSTRAINT "TstDocument_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4377 (class 2606 OID 48332)
-- Name: TstEnv TstEnv_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstEnv"
    ADD CONSTRAINT "TstEnv_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4378 (class 2606 OID 48337)
-- Name: TstEnv TstEnv_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstEnv"
    ADD CONSTRAINT "TstEnv_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4379 (class 2606 OID 48342)
-- Name: TstHistory TstHistory_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstHistory"
    ADD CONSTRAINT "TstHistory_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4380 (class 2606 OID 48347)
-- Name: TstHistory TstHistory_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstHistory"
    ADD CONSTRAINT "TstHistory_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4381 (class 2606 OID 48352)
-- Name: TstModule TstModule_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstModule"
    ADD CONSTRAINT "TstModule_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4382 (class 2606 OID 48357)
-- Name: TstMsg TstMsg_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstMsg"
    ADD CONSTRAINT "TstMsg_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4384 (class 2606 OID 48362)
-- Name: TstOrgGroupUserRelation TstOrgGroupUserRelation_orgGroupId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgGroupUserRelation"
    ADD CONSTRAINT "TstOrgGroupUserRelation_orgGroupId_fkey" FOREIGN KEY ("orgGroupId") REFERENCES public."TstOrgGroup"(id);


--
-- TOC entry 4385 (class 2606 OID 48367)
-- Name: TstOrgGroupUserRelation TstOrgGroupUserRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgGroupUserRelation"
    ADD CONSTRAINT "TstOrgGroupUserRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4386 (class 2606 OID 48372)
-- Name: TstOrgGroupUserRelation TstOrgGroupUserRelation_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgGroupUserRelation"
    ADD CONSTRAINT "TstOrgGroupUserRelation_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4383 (class 2606 OID 48377)
-- Name: TstOrgGroup TstOrgGroup_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgGroup"
    ADD CONSTRAINT "TstOrgGroup_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4388 (class 2606 OID 48382)
-- Name: TstOrgRoleGroupRelation TstOrgRoleGroupRelation_orgGroupId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRoleGroupRelation"
    ADD CONSTRAINT "TstOrgRoleGroupRelation_orgGroupId_fkey" FOREIGN KEY ("orgGroupId") REFERENCES public."TstOrgGroup"(id);


--
-- TOC entry 4389 (class 2606 OID 48387)
-- Name: TstOrgRoleGroupRelation TstOrgRoleGroupRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRoleGroupRelation"
    ADD CONSTRAINT "TstOrgRoleGroupRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4390 (class 2606 OID 48392)
-- Name: TstOrgRoleGroupRelation TstOrgRoleGroupRelation_orgRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRoleGroupRelation"
    ADD CONSTRAINT "TstOrgRoleGroupRelation_orgRoleId_fkey" FOREIGN KEY ("orgRoleId") REFERENCES public."TstOrgRole"(id);


--
-- TOC entry 4391 (class 2606 OID 48397)
-- Name: TstOrgRolePrivilegeRelation TstOrgRolePrivilegeRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRolePrivilegeRelation"
    ADD CONSTRAINT "TstOrgRolePrivilegeRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4392 (class 2606 OID 48402)
-- Name: TstOrgRolePrivilegeRelation TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRolePrivilegeRelation"
    ADD CONSTRAINT "TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey" FOREIGN KEY ("orgPrivilegeId") REFERENCES public."TstOrgPrivilegeDefine"(id);


--
-- TOC entry 4393 (class 2606 OID 48407)
-- Name: TstOrgRolePrivilegeRelation TstOrgRolePrivilegeRelation_orgRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRolePrivilegeRelation"
    ADD CONSTRAINT "TstOrgRolePrivilegeRelation_orgRoleId_fkey" FOREIGN KEY ("orgRoleId") REFERENCES public."TstOrgRole"(id);


--
-- TOC entry 4394 (class 2606 OID 48412)
-- Name: TstOrgRoleUserRelation TstOrgRoleUserRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRoleUserRelation"
    ADD CONSTRAINT "TstOrgRoleUserRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4395 (class 2606 OID 48417)
-- Name: TstOrgRoleUserRelation TstOrgRoleUserRelation_orgRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRoleUserRelation"
    ADD CONSTRAINT "TstOrgRoleUserRelation_orgRoleId_fkey" FOREIGN KEY ("orgRoleId") REFERENCES public."TstOrgRole"(id);


--
-- TOC entry 4396 (class 2606 OID 48422)
-- Name: TstOrgRoleUserRelation TstOrgRoleUserRelation_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRoleUserRelation"
    ADD CONSTRAINT "TstOrgRoleUserRelation_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4387 (class 2606 OID 48427)
-- Name: TstOrgRole TstOrgRole_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgRole"
    ADD CONSTRAINT "TstOrgRole_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4397 (class 2606 OID 48432)
-- Name: TstOrgUserRelation TstOrgUserRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgUserRelation"
    ADD CONSTRAINT "TstOrgUserRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4398 (class 2606 OID 48437)
-- Name: TstOrgUserRelation TstOrgUserRelation_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstOrgUserRelation"
    ADD CONSTRAINT "TstOrgUserRelation_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4399 (class 2606 OID 48442)
-- Name: TstPlan TstPlan_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstPlan"
    ADD CONSTRAINT "TstPlan_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4400 (class 2606 OID 48447)
-- Name: TstPlan TstPlan_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstPlan"
    ADD CONSTRAINT "TstPlan_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4401 (class 2606 OID 48452)
-- Name: TstPlan TstPlan_verId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstPlan"
    ADD CONSTRAINT "TstPlan_verId_fkey" FOREIGN KEY ("verId") REFERENCES public."TstVer"(id);


--
-- TOC entry 4408 (class 2606 OID 48457)
-- Name: TstProjectAccessHistory TstProjectAccessHistory_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectAccessHistory"
    ADD CONSTRAINT "TstProjectAccessHistory_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4409 (class 2606 OID 48462)
-- Name: TstProjectAccessHistory TstProjectAccessHistory_prjId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectAccessHistory"
    ADD CONSTRAINT "TstProjectAccessHistory_prjId_fkey" FOREIGN KEY ("prjId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4410 (class 2606 OID 48467)
-- Name: TstProjectAccessHistory TstProjectAccessHistory_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectAccessHistory"
    ADD CONSTRAINT "TstProjectAccessHistory_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4412 (class 2606 OID 48472)
-- Name: TstProjectRoleEntityRelation TstProjectRoleEntityRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectRoleEntityRelation"
    ADD CONSTRAINT "TstProjectRoleEntityRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4413 (class 2606 OID 48477)
-- Name: TstProjectRoleEntityRelation TstProjectRoleEntityRelation_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectRoleEntityRelation"
    ADD CONSTRAINT "TstProjectRoleEntityRelation_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4414 (class 2606 OID 48482)
-- Name: TstProjectRoleEntityRelation TstProjectRoleEntityRelation_projectRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectRoleEntityRelation"
    ADD CONSTRAINT "TstProjectRoleEntityRelation_projectRoleId_fkey" FOREIGN KEY ("projectRoleId") REFERENCES public."TstProjectRole"(id);


--
-- TOC entry 4415 (class 2606 OID 48487)
-- Name: TstProjectRolePriviledgeRelation TstProjectRolePriviledgeRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectRolePriviledgeRelation"
    ADD CONSTRAINT "TstProjectRolePriviledgeRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4416 (class 2606 OID 48492)
-- Name: TstProjectRolePriviledgeRelation TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectRolePriviledgeRelation"
    ADD CONSTRAINT "TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_fkey" FOREIGN KEY ("projectPrivilegeDefineId") REFERENCES public."TstProjectPrivilegeDefine"(id);


--
-- TOC entry 4417 (class 2606 OID 48497)
-- Name: TstProjectRolePriviledgeRelation TstProjectRolePriviledgeRelation_projectRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectRolePriviledgeRelation"
    ADD CONSTRAINT "TstProjectRolePriviledgeRelation_projectRoleId_fkey" FOREIGN KEY ("projectRoleId") REFERENCES public."TstProjectRole"(id);


--
-- TOC entry 4411 (class 2606 OID 48502)
-- Name: TstProjectRole TstProjectRole_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProjectRole"
    ADD CONSTRAINT "TstProjectRole_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4402 (class 2606 OID 48507)
-- Name: TstProject TstProject_issuePageSolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_issuePageSolutionId_fkey" FOREIGN KEY ("issuePageSolutionId") REFERENCES public."IsuPageSolution"(id);


--
-- TOC entry 4403 (class 2606 OID 48512)
-- Name: TstProject TstProject_issuePrioritySolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_issuePrioritySolutionId_fkey" FOREIGN KEY ("issuePrioritySolutionId") REFERENCES public."IsuPrioritySolution"(id);


--
-- TOC entry 4404 (class 2606 OID 48517)
-- Name: TstProject TstProject_issueTypeSolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_issueTypeSolutionId_fkey" FOREIGN KEY ("issueTypeSolutionId") REFERENCES public."IsuTypeSolution"(id);


--
-- TOC entry 4405 (class 2606 OID 48522)
-- Name: TstProject TstProject_issueWorkflowSolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_issueWorkflowSolutionId_fkey" FOREIGN KEY ("issueWorkflowSolutionId") REFERENCES public."IsuWorkflowSolution"(id);


--
-- TOC entry 4406 (class 2606 OID 48527)
-- Name: TstProject TstProject_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4407 (class 2606 OID 48532)
-- Name: TstProject TstProject_parentId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_parentId_fkey" FOREIGN KEY ("parentId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4418 (class 2606 OID 48537)
-- Name: TstSuite TstSuite_caseProjectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstSuite"
    ADD CONSTRAINT "TstSuite_caseProjectId_fkey" FOREIGN KEY ("caseProjectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4419 (class 2606 OID 48542)
-- Name: TstSuite TstSuite_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstSuite"
    ADD CONSTRAINT "TstSuite_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4420 (class 2606 OID 48547)
-- Name: TstSuite TstSuite_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstSuite"
    ADD CONSTRAINT "TstSuite_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4426 (class 2606 OID 48552)
-- Name: TstTaskAssigneeRelation TstTaskAssigneeRelation_assigneeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstTaskAssigneeRelation"
    ADD CONSTRAINT "TstTaskAssigneeRelation_assigneeId_fkey" FOREIGN KEY ("assigneeId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4427 (class 2606 OID 48557)
-- Name: TstTaskAssigneeRelation TstTaskAssigneeRelation_taskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstTaskAssigneeRelation"
    ADD CONSTRAINT "TstTaskAssigneeRelation_taskId_fkey" FOREIGN KEY ("taskId") REFERENCES public."TstTask"(id);


--
-- TOC entry 4421 (class 2606 OID 48562)
-- Name: TstTask TstTask_caseProjectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_caseProjectId_fkey" FOREIGN KEY ("caseProjectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4422 (class 2606 OID 48567)
-- Name: TstTask TstTask_envId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_envId_fkey" FOREIGN KEY ("envId") REFERENCES public."TstEnv"(id);


--
-- TOC entry 4423 (class 2606 OID 48572)
-- Name: TstTask TstTask_planId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_planId_fkey" FOREIGN KEY ("planId") REFERENCES public."TstPlan"(id);


--
-- TOC entry 4424 (class 2606 OID 48577)
-- Name: TstTask TstTask_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4425 (class 2606 OID 48582)
-- Name: TstTask TstTask_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4428 (class 2606 OID 48587)
-- Name: TstThread TstThread_authorId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstThread"
    ADD CONSTRAINT "TstThread_authorId_fkey" FOREIGN KEY ("authorId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4429 (class 2606 OID 48592)
-- Name: TstThread TstThread_parentId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstThread"
    ADD CONSTRAINT "TstThread_parentId_fkey" FOREIGN KEY ("parentId") REFERENCES public."TstThread"(id);


--
-- TOC entry 4432 (class 2606 OID 48597)
-- Name: TstUserSettings TstUserSettings_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstUserSettings"
    ADD CONSTRAINT "TstUserSettings_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4433 (class 2606 OID 48602)
-- Name: TstUserVerifyCode TstUserVerifyCode_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstUserVerifyCode"
    ADD CONSTRAINT "TstUserVerifyCode_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- TOC entry 4430 (class 2606 OID 48607)
-- Name: TstUser TstUser_defaultOrgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstUser"
    ADD CONSTRAINT "TstUser_defaultOrgId_fkey" FOREIGN KEY ("defaultOrgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4431 (class 2606 OID 48612)
-- Name: TstUser TstUser_defaultPrjId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstUser"
    ADD CONSTRAINT "TstUser_defaultPrjId_fkey" FOREIGN KEY ("defaultPrjId") REFERENCES public."TstProject"(id);


--
-- TOC entry 4434 (class 2606 OID 48617)
-- Name: TstVer TstVer_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstVer"
    ADD CONSTRAINT "TstVer_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- TOC entry 4435 (class 2606 OID 48622)
-- Name: TstVer TstVer_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public."TstVer"
    ADD CONSTRAINT "TstVer_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


-- Completed on 2019-06-12 11:15:21 CST

--
-- PostgreSQL database dump complete
--

