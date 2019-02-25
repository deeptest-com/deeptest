--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

ALTER TABLE ONLY public."TstVer" DROP CONSTRAINT "TstVer_projectId_fkey";
ALTER TABLE ONLY public."TstVer" DROP CONSTRAINT "TstVer_orgId_fkey";
ALTER TABLE ONLY public."TstUser" DROP CONSTRAINT "TstUser_defaultPrjId_fkey";
ALTER TABLE ONLY public."TstUser" DROP CONSTRAINT "TstUser_defaultOrgId_fkey";
ALTER TABLE ONLY public."TstUserVerifyCode" DROP CONSTRAINT "TstUserVerifyCode_userId_fkey";
ALTER TABLE ONLY public."TstUserSettings" DROP CONSTRAINT "TstUserSettings_userId_fkey";
ALTER TABLE ONLY public."TstThread" DROP CONSTRAINT "TstThread_parentId_fkey";
ALTER TABLE ONLY public."TstThread" DROP CONSTRAINT "TstThread_authorId_fkey";
ALTER TABLE ONLY public."TstTask" DROP CONSTRAINT "TstTask_userId_fkey";
ALTER TABLE ONLY public."TstTask" DROP CONSTRAINT "TstTask_projectId_fkey";
ALTER TABLE ONLY public."TstTask" DROP CONSTRAINT "TstTask_planId_fkey";
ALTER TABLE ONLY public."TstTask" DROP CONSTRAINT "TstTask_envId_fkey";
ALTER TABLE ONLY public."TstTask" DROP CONSTRAINT "TstTask_caseProjectId_fkey";
ALTER TABLE ONLY public."TstTaskAssigneeRelation" DROP CONSTRAINT "TstTaskAssigneeRelation_taskId_fkey";
ALTER TABLE ONLY public."TstTaskAssigneeRelation" DROP CONSTRAINT "TstTaskAssigneeRelation_assigneeId_fkey";
ALTER TABLE ONLY public."TstSuite" DROP CONSTRAINT "TstSuite_userId_fkey";
ALTER TABLE ONLY public."TstSuite" DROP CONSTRAINT "TstSuite_projectId_fkey";
ALTER TABLE ONLY public."TstSuite" DROP CONSTRAINT "TstSuite_caseProjectId_fkey";
ALTER TABLE ONLY public."TstProject" DROP CONSTRAINT "TstProject_parentId_fkey";
ALTER TABLE ONLY public."TstProject" DROP CONSTRAINT "TstProject_orgId_fkey";
ALTER TABLE ONLY public."TstProject" DROP CONSTRAINT "TstProject_issueWorkflowSolutionId_fkey";
ALTER TABLE ONLY public."TstProject" DROP CONSTRAINT "TstProject_issueTypeSolutionId_fkey";
ALTER TABLE ONLY public."TstProject" DROP CONSTRAINT "TstProject_issuePrioritySolutionId_fkey";
ALTER TABLE ONLY public."TstProject" DROP CONSTRAINT "TstProject_issuePageSolutionId_fkey";
ALTER TABLE ONLY public."TstProjectRole" DROP CONSTRAINT "TstProjectRole_orgId_fkey";
ALTER TABLE ONLY public."TstProjectRolePriviledgeRelation" DROP CONSTRAINT "TstProjectRolePriviledgeRelation_projectRoleId_fkey";
ALTER TABLE ONLY public."TstProjectRolePriviledgeRelation" DROP CONSTRAINT "TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_fkey";
ALTER TABLE ONLY public."TstProjectRolePriviledgeRelation" DROP CONSTRAINT "TstProjectRolePriviledgeRelation_orgId_fkey";
ALTER TABLE ONLY public."TstProjectRoleEntityRelation" DROP CONSTRAINT "TstProjectRoleEntityRelation_projectRoleId_fkey";
ALTER TABLE ONLY public."TstProjectRoleEntityRelation" DROP CONSTRAINT "TstProjectRoleEntityRelation_projectId_fkey";
ALTER TABLE ONLY public."TstProjectRoleEntityRelation" DROP CONSTRAINT "TstProjectRoleEntityRelation_orgId_fkey";
ALTER TABLE ONLY public."TstProjectAccessHistory" DROP CONSTRAINT "TstProjectAccessHistory_userId_fkey";
ALTER TABLE ONLY public."TstProjectAccessHistory" DROP CONSTRAINT "TstProjectAccessHistory_prjId_fkey";
ALTER TABLE ONLY public."TstProjectAccessHistory" DROP CONSTRAINT "TstProjectAccessHistory_orgId_fkey";
ALTER TABLE ONLY public."TstPlan" DROP CONSTRAINT "TstPlan_verId_fkey";
ALTER TABLE ONLY public."TstPlan" DROP CONSTRAINT "TstPlan_userId_fkey";
ALTER TABLE ONLY public."TstPlan" DROP CONSTRAINT "TstPlan_projectId_fkey";
ALTER TABLE ONLY public."TstOrgUserRelation" DROP CONSTRAINT "TstOrgUserRelation_userId_fkey";
ALTER TABLE ONLY public."TstOrgUserRelation" DROP CONSTRAINT "TstOrgUserRelation_orgId_fkey";
ALTER TABLE ONLY public."TstOrgRole" DROP CONSTRAINT "TstOrgRole_orgId_fkey";
ALTER TABLE ONLY public."TstOrgRoleUserRelation" DROP CONSTRAINT "TstOrgRoleUserRelation_userId_fkey";
ALTER TABLE ONLY public."TstOrgRoleUserRelation" DROP CONSTRAINT "TstOrgRoleUserRelation_orgRoleId_fkey";
ALTER TABLE ONLY public."TstOrgRoleUserRelation" DROP CONSTRAINT "TstOrgRoleUserRelation_orgId_fkey";
ALTER TABLE ONLY public."TstOrgRolePrivilegeRelation" DROP CONSTRAINT "TstOrgRolePrivilegeRelation_orgRoleId_fkey";
ALTER TABLE ONLY public."TstOrgRolePrivilegeRelation" DROP CONSTRAINT "TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey";
ALTER TABLE ONLY public."TstOrgRolePrivilegeRelation" DROP CONSTRAINT "TstOrgRolePrivilegeRelation_orgId_fkey";
ALTER TABLE ONLY public."TstOrgRoleGroupRelation" DROP CONSTRAINT "TstOrgRoleGroupRelation_orgRoleId_fkey";
ALTER TABLE ONLY public."TstOrgRoleGroupRelation" DROP CONSTRAINT "TstOrgRoleGroupRelation_orgId_fkey";
ALTER TABLE ONLY public."TstOrgRoleGroupRelation" DROP CONSTRAINT "TstOrgRoleGroupRelation_orgGroupId_fkey";
ALTER TABLE ONLY public."TstOrgGroup" DROP CONSTRAINT "TstOrgGroup_orgId_fkey";
ALTER TABLE ONLY public."TstOrgGroupUserRelation" DROP CONSTRAINT "TstOrgGroupUserRelation_userId_fkey";
ALTER TABLE ONLY public."TstOrgGroupUserRelation" DROP CONSTRAINT "TstOrgGroupUserRelation_orgId_fkey";
ALTER TABLE ONLY public."TstOrgGroupUserRelation" DROP CONSTRAINT "TstOrgGroupUserRelation_orgGroupId_fkey";
ALTER TABLE ONLY public."TstMsg" DROP CONSTRAINT "TstMsg_userId_fkey";
ALTER TABLE ONLY public."TstModule" DROP CONSTRAINT "TstModule_projectId_fkey";
ALTER TABLE ONLY public."TstHistory" DROP CONSTRAINT "TstHistory_userId_fkey";
ALTER TABLE ONLY public."TstHistory" DROP CONSTRAINT "TstHistory_projectId_fkey";
ALTER TABLE ONLY public."TstEnv" DROP CONSTRAINT "TstEnv_projectId_fkey";
ALTER TABLE ONLY public."TstEnv" DROP CONSTRAINT "TstEnv_orgId_fkey";
ALTER TABLE ONLY public."TstDocument" DROP CONSTRAINT "TstDocument_userId_fkey";
ALTER TABLE ONLY public."TstCase" DROP CONSTRAINT "TstCase_updateById_fkey";
ALTER TABLE ONLY public."TstCase" DROP CONSTRAINT "TstCase_typeId_fkey";
ALTER TABLE ONLY public."TstCase" DROP CONSTRAINT "TstCase_projectId_fkey";
ALTER TABLE ONLY public."TstCase" DROP CONSTRAINT "TstCase_priorityId_fkey";
ALTER TABLE ONLY public."TstCase" DROP CONSTRAINT "TstCase_pId_fkey";
ALTER TABLE ONLY public."TstCase" DROP CONSTRAINT "TstCase_createById_fkey";
ALTER TABLE ONLY public."TstCaseType" DROP CONSTRAINT "TstCaseType_orgId_fkey";
ALTER TABLE ONLY public."TstCaseStep" DROP CONSTRAINT "TstCaseStep_caseId_fkey";
ALTER TABLE ONLY public."TstCasePriority" DROP CONSTRAINT "TstCasePriority_orgId_fkey";
ALTER TABLE ONLY public."TstCaseInTask" DROP CONSTRAINT "TstCaseInTask_taskId_fkey";
ALTER TABLE ONLY public."TstCaseInTask" DROP CONSTRAINT "TstCaseInTask_projectId_fkey";
ALTER TABLE ONLY public."TstCaseInTask" DROP CONSTRAINT "TstCaseInTask_planId_fkey";
ALTER TABLE ONLY public."TstCaseInTask" DROP CONSTRAINT "TstCaseInTask_pId_fkey";
ALTER TABLE ONLY public."TstCaseInTask" DROP CONSTRAINT "TstCaseInTask_exeBy_fkey";
ALTER TABLE ONLY public."TstCaseInTask" DROP CONSTRAINT "TstCaseInTask_createBy_fkey";
ALTER TABLE ONLY public."TstCaseInTask" DROP CONSTRAINT "TstCaseInTask_caseId_fkey";
ALTER TABLE ONLY public."TstCaseInTaskIssue" DROP CONSTRAINT "TstCaseInTaskIssue_userId_fkey";
ALTER TABLE ONLY public."TstCaseInTaskIssue" DROP CONSTRAINT "TstCaseInTaskIssue_issueId_fkey";
ALTER TABLE ONLY public."TstCaseInTaskIssue" DROP CONSTRAINT "TstCaseInTaskIssue_caseInTaskId_fkey";
ALTER TABLE ONLY public."TstCaseInTaskHistory" DROP CONSTRAINT "TstCaseInTaskHistory_caseInTaskId_fkey";
ALTER TABLE ONLY public."TstCaseInTaskHistory" DROP CONSTRAINT "TstCaseInTaskHistory_caseId_fkey";
ALTER TABLE ONLY public."TstCaseInTaskComments" DROP CONSTRAINT "TstCaseInTaskComments_userId_fkey";
ALTER TABLE ONLY public."TstCaseInTaskComments" DROP CONSTRAINT "TstCaseInTaskComments_caseInTaskId_fkey";
ALTER TABLE ONLY public."TstCaseInTaskAttachment" DROP CONSTRAINT "TstCaseInTaskAttachment_userId_fkey";
ALTER TABLE ONLY public."TstCaseInTaskAttachment" DROP CONSTRAINT "TstCaseInTaskAttachment_caseInTaskId_fkey";
ALTER TABLE ONLY public."TstCaseInSuite" DROP CONSTRAINT "TstCaseInSuite_suiteId_fkey";
ALTER TABLE ONLY public."TstCaseInSuite" DROP CONSTRAINT "TstCaseInSuite_projectId_fkey";
ALTER TABLE ONLY public."TstCaseInSuite" DROP CONSTRAINT "TstCaseInSuite_pId_fkey";
ALTER TABLE ONLY public."TstCaseInSuite" DROP CONSTRAINT "TstCaseInSuite_caseId_fkey";
ALTER TABLE ONLY public."TstCaseHistory" DROP CONSTRAINT "TstCaseHistory_caseId_fkey";
ALTER TABLE ONLY public."TstCaseExeStatus" DROP CONSTRAINT "TstCaseExeStatus_orgId_fkey";
ALTER TABLE ONLY public."TstCaseComments" DROP CONSTRAINT "TstCaseComments_userId_fkey";
ALTER TABLE ONLY public."TstCaseComments" DROP CONSTRAINT "TstCaseComments_caseId_fkey";
ALTER TABLE ONLY public."TstCaseAttachment" DROP CONSTRAINT "TstCaseAttachment_userId_fkey";
ALTER TABLE ONLY public."TstCaseAttachment" DROP CONSTRAINT "TstCaseAttachment_caseId_fkey";
ALTER TABLE ONLY public."TstAlert" DROP CONSTRAINT "TstAlert_userId_fkey";
ALTER TABLE ONLY public."TstAlert" DROP CONSTRAINT "TstAlert_assigneeId_fkey";
ALTER TABLE ONLY public."SysRoleUserRelation" DROP CONSTRAINT "SysRoleUserRelation_userId_fkey";
ALTER TABLE ONLY public."SysRoleUserRelation" DROP CONSTRAINT "SysRoleUserRelation_roleId_fkey";
ALTER TABLE ONLY public."SysRolePrivilegeRelation" DROP CONSTRAINT "SysRolePrivilegeRelation_roleId_fkey";
ALTER TABLE ONLY public."SysRolePrivilegeRelation" DROP CONSTRAINT "SysRolePrivilegeRelation_privilegeId_fkey";
ALTER TABLE ONLY public."IsuWorkflow" DROP CONSTRAINT "IsuWorkflow_orgId_fkey";
ALTER TABLE ONLY public."IsuWorkflowTransition" DROP CONSTRAINT "IsuWorkflowTransition_workflowId_fkey";
ALTER TABLE ONLY public."IsuWorkflowTransition" DROP CONSTRAINT "IsuWorkflowTransition_orgId_fkey";
ALTER TABLE ONLY public."IsuWorkflowTransition" DROP CONSTRAINT "IsuWorkflowTransition_actionPageId_fkey";
ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation" DROP CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey";
ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation" DROP CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey";
ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation" DROP CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_orgId_fkey";
ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation" DROP CONSTRAINT "IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_fkey";
ALTER TABLE ONLY public."IsuWorkflowTransitionDefine" DROP CONSTRAINT "IsuWorkflowTransitionDefine_srcStatusId_fkey";
ALTER TABLE ONLY public."IsuWorkflowTransitionDefine" DROP CONSTRAINT "IsuWorkflowTransitionDefine_dictStatusId_fkey";
ALTER TABLE ONLY public."IsuWorkflowStatusRelation" DROP CONSTRAINT "IsuWorkflowStatusRelation_workflowId_fkey";
ALTER TABLE ONLY public."IsuWorkflowStatusRelation" DROP CONSTRAINT "IsuWorkflowStatusRelation_statusId_fkey";
ALTER TABLE ONLY public."IsuWorkflowStatusRelation" DROP CONSTRAINT "IsuWorkflowStatusRelation_orgId_fkey";
ALTER TABLE ONLY public."IsuWorkflowStatusRelationDefine" DROP CONSTRAINT "IsuWorkflowStatusRelationDefine_statusId_fkey";
ALTER TABLE ONLY public."IsuWorkflowSolution" DROP CONSTRAINT "IsuWorkflowSolution_orgId_fkey";
ALTER TABLE ONLY public."IsuWorkflowSolutionItem" DROP CONSTRAINT "IsuWorkflowSolutionItem_workflowId_fkey";
ALTER TABLE ONLY public."IsuWorkflowSolutionItem" DROP CONSTRAINT "IsuWorkflowSolutionItem_typeId_fkey";
ALTER TABLE ONLY public."IsuWorkflowSolutionItem" DROP CONSTRAINT "IsuWorkflowSolutionItem_solutionId_fkey";
ALTER TABLE ONLY public."IsuWorkflowSolutionItem" DROP CONSTRAINT "IsuWorkflowSolutionItem_orgId_fkey";
ALTER TABLE ONLY public."IsuWatch" DROP CONSTRAINT "IsuWatch_userId_fkey";
ALTER TABLE ONLY public."IsuWatch" DROP CONSTRAINT "IsuWatch_issueId_fkey";
ALTER TABLE ONLY public."IsuType" DROP CONSTRAINT "IsuType_orgId_fkey";
ALTER TABLE ONLY public."IsuTypeSolution" DROP CONSTRAINT "IsuTypeSolution_orgId_fkey";
ALTER TABLE ONLY public."IsuTypeSolutionItem" DROP CONSTRAINT "IsuTypeSolutionItem_typeId_fkey";
ALTER TABLE ONLY public."IsuTypeSolutionItem" DROP CONSTRAINT "IsuTypeSolutionItem_solutionId_fkey";
ALTER TABLE ONLY public."IsuTypeSolutionItem" DROP CONSTRAINT "IsuTypeSolutionItem_orgId_fkey";
ALTER TABLE ONLY public."IsuTag" DROP CONSTRAINT "IsuTag_userId_fkey";
ALTER TABLE ONLY public."IsuTag" DROP CONSTRAINT "IsuTag_orgId_fkey";
ALTER TABLE ONLY public."IsuTagRelation" DROP CONSTRAINT "IsuTagRelation_tagId_fkey";
ALTER TABLE ONLY public."IsuTagRelation" DROP CONSTRAINT "IsuTagRelation_issueId_fkey";
ALTER TABLE ONLY public."IsuStatus" DROP CONSTRAINT "IsuStatus_orgId_fkey";
ALTER TABLE ONLY public."IsuStatus" DROP CONSTRAINT "IsuStatus_categoryId_fkey";
ALTER TABLE ONLY public."IsuStatusDefine" DROP CONSTRAINT "IsuStatusDefine_categoryId_fkey";
ALTER TABLE ONLY public."IsuSeverity" DROP CONSTRAINT "IsuSeverity_orgId_fkey";
ALTER TABLE ONLY public."IsuSeveritySolution" DROP CONSTRAINT "IsuSeveritySolution_orgId_fkey";
ALTER TABLE ONLY public."IsuSeveritySolutionItem" DROP CONSTRAINT "IsuSeveritySolutionItem_solutionId_fkey";
ALTER TABLE ONLY public."IsuSeveritySolutionItem" DROP CONSTRAINT "IsuSeveritySolutionItem_severityId_fkey";
ALTER TABLE ONLY public."IsuResolution" DROP CONSTRAINT "IsuResolution_orgId_fkey";
ALTER TABLE ONLY public."IsuQuery" DROP CONSTRAINT "IsuQuery_userId_fkey";
ALTER TABLE ONLY public."IsuQuery" DROP CONSTRAINT "IsuQuery_projectId_fkey";
ALTER TABLE ONLY public."IsuPriority" DROP CONSTRAINT "IsuPriority_orgId_fkey";
ALTER TABLE ONLY public."IsuPrioritySolution" DROP CONSTRAINT "IsuPrioritySolution_orgId_fkey";
ALTER TABLE ONLY public."IsuPrioritySolutionItem" DROP CONSTRAINT "IsuPrioritySolutionItem_solutionId_fkey";
ALTER TABLE ONLY public."IsuPrioritySolutionItem" DROP CONSTRAINT "IsuPrioritySolutionItem_priorityId_fkey";
ALTER TABLE ONLY public."IsuPrioritySolutionItem" DROP CONSTRAINT "IsuPrioritySolutionItem_orgId_fkey";
ALTER TABLE ONLY public."IsuPage" DROP CONSTRAINT "IsuPage_orgId_fkey";
ALTER TABLE ONLY public."IsuPageSolution" DROP CONSTRAINT "IsuPageSolution_orgId_fkey";
ALTER TABLE ONLY public."IsuPageSolutionItem" DROP CONSTRAINT "IsuPageSolutionItem_typeId_fkey";
ALTER TABLE ONLY public."IsuPageSolutionItem" DROP CONSTRAINT "IsuPageSolutionItem_solutionId_fkey";
ALTER TABLE ONLY public."IsuPageSolutionItem" DROP CONSTRAINT "IsuPageSolutionItem_pageId_fkey";
ALTER TABLE ONLY public."IsuPageSolutionItem" DROP CONSTRAINT "IsuPageSolutionItem_orgId_fkey";
ALTER TABLE ONLY public."IsuNotification" DROP CONSTRAINT "IsuNotification_orgId_fkey";
ALTER TABLE ONLY public."IsuLink" DROP CONSTRAINT "IsuLink_srcIssueId_fkey";
ALTER TABLE ONLY public."IsuLink" DROP CONSTRAINT "IsuLink_reasonId_fkey";
ALTER TABLE ONLY public."IsuLink" DROP CONSTRAINT "IsuLink_dictIssueId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_verId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_typeId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_statusId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_resolutionId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_reporterId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_projectId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_priorityId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_orgId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_envId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_creatorId_fkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_assigneeId_fkey";
ALTER TABLE ONLY public."IsuIssueExt" DROP CONSTRAINT "IsuIssueExt_pid_fkey";
ALTER TABLE ONLY public."IsuHistory" DROP CONSTRAINT "IsuHistory_issueId_fkey";
ALTER TABLE ONLY public."IsuField" DROP CONSTRAINT "IsuField_orgId_fkey";
ALTER TABLE ONLY public."IsuDocument" DROP CONSTRAINT "IsuDocument_userId_fkey";
ALTER TABLE ONLY public."IsuDocument" DROP CONSTRAINT "IsuDocument_issueId_fkey";
ALTER TABLE ONLY public."IsuCustomFieldSolution" DROP CONSTRAINT "IsuCustomFieldSolution_orgId_fkey";
ALTER TABLE ONLY public."IsuCustomFieldSolutionProjectRelation" DROP CONSTRAINT "IsuCustomFieldSolutionProjectRelation_solutionId_fkey";
ALTER TABLE ONLY public."IsuCustomFieldSolutionProjectRelation" DROP CONSTRAINT "IsuCustomFieldSolutionProjectRelation_projectId_fkey";
ALTER TABLE ONLY public."IsuCustomFieldSolutionProjectRelation" DROP CONSTRAINT "IsuCustomFieldSolutionProjectRelation_orgId_fkey";
ALTER TABLE ONLY public."IsuCustomFieldSolutionFieldRelation" DROP CONSTRAINT "IsuCustomFieldSolutionFieldRelation_solutionId_fkey";
ALTER TABLE ONLY public."IsuCustomFieldSolutionFieldRelation" DROP CONSTRAINT "IsuCustomFieldSolutionFieldRelation_fieldId_fkey";
ALTER TABLE ONLY public."IsuComments" DROP CONSTRAINT "IsuComments_userId_fkey";
ALTER TABLE ONLY public."IsuComments" DROP CONSTRAINT "IsuComments_issueId_fkey";
ALTER TABLE ONLY public."IsuAttachment" DROP CONSTRAINT "IsuAttachment_userId_fkey";
ALTER TABLE ONLY public."IsuAttachment" DROP CONSTRAINT "IsuAttachment_issueId_fkey";
ALTER TABLE ONLY public."CustomField" DROP CONSTRAINT "CustomField_orgId_fkey";
ALTER TABLE ONLY public."CustomFieldOption" DROP CONSTRAINT "CustomFieldOption_orgId_fkey";
ALTER TABLE ONLY public."CustomFieldOption" DROP CONSTRAINT "CustomFieldOption_fieldId_fkey";
ALTER TABLE ONLY public."CustomFieldOptionDefine" DROP CONSTRAINT "CustomFieldOptionDefine_fieldId_fkey";
DROP TRIGGER issue_tsvector_update_trigger ON public."IsuIssue";
DROP INDEX public.idx_test_case_extprop;
DROP INDEX public.idx_isu_issue_extprop;
DROP INDEX public."fki_TstVer_projectId_fkey";
DROP INDEX public."fki_TstVer_orgId_fkey";
DROP INDEX public."fki_TstUser_defaultPrjId_fkey";
DROP INDEX public."fki_TstUser_defaultOrgId_fkey";
DROP INDEX public."fki_TstUserVerifyCode_userId_fkey";
DROP INDEX public."fki_TstUserSettings_userId_fkey";
DROP INDEX public."fki_TstThread_parentId_fkey";
DROP INDEX public."fki_TstThread_authorId_fkey";
DROP INDEX public."fki_TstTask_userId_fkey";
DROP INDEX public."fki_TstTask_projectId_fkey";
DROP INDEX public."fki_TstTask_planId_fkey";
DROP INDEX public."fki_TstTask_envId_fkey";
DROP INDEX public."fki_TstTask_caseProjectId_fkey";
DROP INDEX public."fki_TstTaskAssigneeRelation_taskId_fkey";
DROP INDEX public."fki_TstTaskAssigneeRelation_assigneeId_fkey";
DROP INDEX public."fki_TstSuite_userId_fkey";
DROP INDEX public."fki_TstSuite_projectId_fkey";
DROP INDEX public."fki_TstSuite_caseProjectId_fkey";
DROP INDEX public."fki_TstProject_parentId_fkey";
DROP INDEX public."fki_TstProject_orgId_fkey";
DROP INDEX public."fki_TstProject_issueWorkflowSolutionId_fkey";
DROP INDEX public."fki_TstProject_issueTypeSolutionId_fkey";
DROP INDEX public."fki_TstProject_issuePrioritySolutionId_fkey";
DROP INDEX public."fki_TstProject_issuePageSolutionId_fkey";
DROP INDEX public."fki_TstProjectRole_orgId_fkey";
DROP INDEX public."fki_TstProjectRolePriviledgeRelation_projectRoleId_fkey";
DROP INDEX public."fki_TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_f";
DROP INDEX public."fki_TstProjectRolePriviledgeRelation_orgId_fkey";
DROP INDEX public."fki_TstProjectRoleEntityRelation_projectRoleId_fkey";
DROP INDEX public."fki_TstProjectRoleEntityRelation_projectId_fkey";
DROP INDEX public."fki_TstProjectRoleEntityRelation_orgId_fkey";
DROP INDEX public."fki_TstProjectAccessHistory_userId_fkey";
DROP INDEX public."fki_TstProjectAccessHistory_prjId_fkey";
DROP INDEX public."fki_TstProjectAccessHistory_orgId_fkey";
DROP INDEX public."fki_TstPlan_verId_fkey";
DROP INDEX public."fki_TstPlan_userId_fkey";
DROP INDEX public."fki_TstPlan_projectId_fkey";
DROP INDEX public."fki_TstOrgUserRelation_userId_fkey";
DROP INDEX public."fki_TstOrgUserRelation_orgId_fkey";
DROP INDEX public."fki_TstOrgRole_orgId_fkey";
DROP INDEX public."fki_TstOrgRoleUserRelation_userId_fkey";
DROP INDEX public."fki_TstOrgRoleUserRelation_orgRoleId_fkey";
DROP INDEX public."fki_TstOrgRoleUserRelation_orgId_fkey";
DROP INDEX public."fki_TstOrgRolePrivilegeRelation_orgRoleId_fkey";
DROP INDEX public."fki_TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey";
DROP INDEX public."fki_TstOrgRolePrivilegeRelation_orgId_fkey";
DROP INDEX public."fki_TstOrgRoleGroupRelation_orgRoleId_fkey";
DROP INDEX public."fki_TstOrgRoleGroupRelation_orgId_fkey";
DROP INDEX public."fki_TstOrgRoleGroupRelation_orgGroupId_fkey";
DROP INDEX public."fki_TstOrgGroup_orgId_fkey";
DROP INDEX public."fki_TstOrgGroupUserRelation_userId_fkey";
DROP INDEX public."fki_TstOrgGroupUserRelation_orgId_fkey";
DROP INDEX public."fki_TstOrgGroupUserRelation_orgGroupId_fkey";
DROP INDEX public."fki_TstMsg_userId_fkey";
DROP INDEX public."fki_TstModule_projectId_fkey";
DROP INDEX public."fki_TstHistory_userId_fkey";
DROP INDEX public."fki_TstHistory_projectId_fkey";
DROP INDEX public."fki_TstEnv_projectId_fkey";
DROP INDEX public."fki_TstEnv_orgId_fkey";
DROP INDEX public."fki_TstDocument_userId_fkey";
DROP INDEX public."fki_TstCase_updateById_fkey";
DROP INDEX public."fki_TstCase_typeId_fkey";
DROP INDEX public."fki_TstCase_projectId_fkey";
DROP INDEX public."fki_TstCase_priorityId_fkey";
DROP INDEX public."fki_TstCase_pId_fkey";
DROP INDEX public."fki_TstCase_createById_fkey";
DROP INDEX public."fki_TstCaseType_orgId_fkey";
DROP INDEX public."fki_TstCaseStep_caseId_fkey";
DROP INDEX public."fki_TstCasePriority_orgId_fkey";
DROP INDEX public."fki_TstCaseInTask_taskId_fkey";
DROP INDEX public."fki_TstCaseInTask_projectId_fkey";
DROP INDEX public."fki_TstCaseInTask_planId_fkey";
DROP INDEX public."fki_TstCaseInTask_pId_fkey";
DROP INDEX public."fki_TstCaseInTask_exeBy_fkey";
DROP INDEX public."fki_TstCaseInTask_createBy_fkey";
DROP INDEX public."fki_TstCaseInTask_caseId_fkey";
DROP INDEX public."fki_TstCaseInTaskIssue_userId_fkey";
DROP INDEX public."fki_TstCaseInTaskIssue_issueId_fkey";
DROP INDEX public."fki_TstCaseInTaskIssue_caseInTaskId_fkey";
DROP INDEX public."fki_TstCaseInTaskHistory_caseInTaskId_fkey";
DROP INDEX public."fki_TstCaseInTaskHistory_caseId_fkey";
DROP INDEX public."fki_TstCaseInTaskComments_userId_fkey";
DROP INDEX public."fki_TstCaseInTaskComments_caseInTaskId_fkey";
DROP INDEX public."fki_TstCaseInTaskAttachment_userId_fkey";
DROP INDEX public."fki_TstCaseInTaskAttachment_caseInTaskId_fkey";
DROP INDEX public."fki_TstCaseInSuite_suiteId_fkey";
DROP INDEX public."fki_TstCaseInSuite_projectId_fkey";
DROP INDEX public."fki_TstCaseInSuite_pId_fkey";
DROP INDEX public."fki_TstCaseInSuite_caseId_fkey";
DROP INDEX public."fki_TstCaseHistory_caseId_fkey";
DROP INDEX public."fki_TstCaseExeStatus_orgId_fkey";
DROP INDEX public."fki_TstCaseComments_userId_fkey";
DROP INDEX public."fki_TstCaseComments_caseId_fkey";
DROP INDEX public."fki_TstCaseAttachment_userId_fkey";
DROP INDEX public."fki_TstCaseAttachment_caseId_fkey";
DROP INDEX public."fki_TstAlert_userId_fkey";
DROP INDEX public."fki_TstAlert_assigneeId_fkey";
DROP INDEX public."fki_SysRoleUserRelation_userId_fkey";
DROP INDEX public."fki_SysRoleUserRelation_roleId_fkey";
DROP INDEX public."fki_SysRolePrivilegeRelation_roleId_fkey";
DROP INDEX public."fki_SysRolePrivilegeRelation_privilegeId_fkey";
DROP INDEX public."fki_IsuWorkflow_orgId_fkey";
DROP INDEX public."fki_IsuWorkflowTransition_workflowId_fkey";
DROP INDEX public."fki_IsuWorkflowTransition_orgId_fkey";
DROP INDEX public."fki_IsuWorkflowTransition_actionPageId_fkey";
DROP INDEX public."fki_IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey";
DROP INDEX public."fki_IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey";
DROP INDEX public."fki_IsuWorkflowTransitionProjectRoleRelation_orgId_fkey";
DROP INDEX public."fki_IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_";
DROP INDEX public."fki_IsuWorkflowTransitionDefine_srcStatusId_fkey";
DROP INDEX public."fki_IsuWorkflowTransitionDefine_dictStatusId_fkey";
DROP INDEX public."fki_IsuWorkflowStatusRelation_workflowId_fkey";
DROP INDEX public."fki_IsuWorkflowStatusRelation_statusId_fkey";
DROP INDEX public."fki_IsuWorkflowStatusRelation_orgId_fkey";
DROP INDEX public."fki_IsuWorkflowStatusRelationDefine_statusId_fkey";
DROP INDEX public."fki_IsuWorkflowSolution_orgId_fkey";
DROP INDEX public."fki_IsuWorkflowSolutionItem_workflowId_fkey";
DROP INDEX public."fki_IsuWorkflowSolutionItem_typeId_fkey";
DROP INDEX public."fki_IsuWorkflowSolutionItem_solutionId_fkey";
DROP INDEX public."fki_IsuWorkflowSolutionItem_orgId_fkey";
DROP INDEX public."fki_IsuWatch_userId_fkey";
DROP INDEX public."fki_IsuWatch_issueId_fkey";
DROP INDEX public."fki_IsuType_orgId_fkey";
DROP INDEX public."fki_IsuTypeSolution_orgId_fkey";
DROP INDEX public."fki_IsuTypeSolutionItem_typeId_fkey";
DROP INDEX public."fki_IsuTypeSolutionItem_solutionId_fkey";
DROP INDEX public."fki_IsuTypeSolutionItem_orgId_fkey";
DROP INDEX public."fki_IsuTag_userId_fkey";
DROP INDEX public."fki_IsuTag_orgId_fkey";
DROP INDEX public."fki_IsuTagRelation_tagId_fkey";
DROP INDEX public."fki_IsuTagRelation_issueId_fkey";
DROP INDEX public."fki_IsuStatus_orgId_fkey";
DROP INDEX public."fki_IsuStatus_categoryId_fkey";
DROP INDEX public."fki_IsuStatusDefine_categoryId_fkey";
DROP INDEX public."fki_IsuSeverity_orgId_fkey";
DROP INDEX public."fki_IsuSeveritySolution_orgId_fkey";
DROP INDEX public."fki_IsuSeveritySolutionItem_solutionId_fkey";
DROP INDEX public."fki_IsuSeveritySolutionItem_severityId_fkey";
DROP INDEX public."fki_IsuResolution_orgId_fkey";
DROP INDEX public."fki_IsuQuery_userId_fkey";
DROP INDEX public."fki_IsuQuery_projectId_fkey";
DROP INDEX public."fki_IsuPriority_orgId_fkey";
DROP INDEX public."fki_IsuPrioritySolution_orgId_fkey";
DROP INDEX public."fki_IsuPrioritySolutionItem_solutionId_fkey";
DROP INDEX public."fki_IsuPrioritySolutionItem_priorityId_fkey";
DROP INDEX public."fki_IsuPrioritySolutionItem_orgId_fkey";
DROP INDEX public."fki_IsuPage_orgId_fkey";
DROP INDEX public."fki_IsuPageSolution_orgId_fkey";
DROP INDEX public."fki_IsuPageSolutionItem_typeId_fkey";
DROP INDEX public."fki_IsuPageSolutionItem_solutionId_fkey";
DROP INDEX public."fki_IsuPageSolutionItem_pageId_fkey";
DROP INDEX public."fki_IsuPageSolutionItem_orgId_fkey";
DROP INDEX public."fki_IsuNotification_orgId_fkey";
DROP INDEX public."fki_IsuLink_srcIssueId_fkey";
DROP INDEX public."fki_IsuLink_reasonId_fkey";
DROP INDEX public."fki_IsuLink_dictIssueId_fkey";
DROP INDEX public."fki_IsuIssue_verId_fkey";
DROP INDEX public."fki_IsuIssue_typeId_fkey";
DROP INDEX public."fki_IsuIssue_statusId_fkey";
DROP INDEX public."fki_IsuIssue_resolutionId_fkey";
DROP INDEX public."fki_IsuIssue_reporterId_fkey";
DROP INDEX public."fki_IsuIssue_projectId_fkey";
DROP INDEX public."fki_IsuIssue_priorityId_fkey";
DROP INDEX public."fki_IsuIssue_orgId_fkey";
DROP INDEX public."fki_IsuIssue_envId_fkey";
DROP INDEX public."fki_IsuIssue_creatorId_fkey";
DROP INDEX public."fki_IsuIssue_assigneeId_fkey";
DROP INDEX public."fki_IsuIssueExt_pid_fkey";
DROP INDEX public."fki_IsuHistory_issueId_fkey";
DROP INDEX public."fki_IsuField_orgId_fkey";
DROP INDEX public."fki_IsuDocument_userId_fkey";
DROP INDEX public."fki_IsuDocument_issueId_fkey";
DROP INDEX public."fki_IsuCustomFieldSolution_orgId_fkey";
DROP INDEX public."fki_IsuCustomFieldSolutionProjectRelation_solutionId_fkey";
DROP INDEX public."fki_IsuCustomFieldSolutionProjectRelation_projectId_fkey";
DROP INDEX public."fki_IsuCustomFieldSolutionProjectRelation_orgId_fkey";
DROP INDEX public."fki_IsuCustomFieldSolutionFieldRelation_solutionId_fkey";
DROP INDEX public."fki_IsuCustomFieldSolutionFieldRelation_fieldId_fkey";
DROP INDEX public."fki_IsuComments_userId_fkey";
DROP INDEX public."fki_IsuComments_issueId_fkey";
DROP INDEX public."fki_IsuAttachment_userId_fkey";
DROP INDEX public."fki_IsuAttachment_issueId_fkey";
DROP INDEX public."fki_CustomField_orgId_fkey";
DROP INDEX public."fki_CustomFieldOption_orgId_fkey";
DROP INDEX public."fki_CustomFieldOption_fieldId_fkey";
DROP INDEX public."fki_CustomFieldOptionDefine_fieldId_fkey";
ALTER TABLE ONLY public."TstVer" DROP CONSTRAINT "TstVer_pkey";
ALTER TABLE ONLY public."TstUser" DROP CONSTRAINT "TstUser_pkey";
ALTER TABLE ONLY public."TstUserVerifyCode" DROP CONSTRAINT "TstUserVerifyCode_pkey";
ALTER TABLE ONLY public."TstThread" DROP CONSTRAINT "TstThread_pkey";
ALTER TABLE ONLY public."TstTask" DROP CONSTRAINT "TstTask_pkey";
ALTER TABLE ONLY public."TstSuite" DROP CONSTRAINT "TstSuite_pkey";
ALTER TABLE ONLY public."TstProject" DROP CONSTRAINT "TstProject_pkey";
ALTER TABLE ONLY public."TstProjectRole" DROP CONSTRAINT "TstProjectRole_pkey";
ALTER TABLE ONLY public."TstProjectPrivilegeDefine" DROP CONSTRAINT "TstProjectPrivilegeDefine_pkey";
ALTER TABLE ONLY public."TstProjectAccessHistory" DROP CONSTRAINT "TstProjectAccessHistory_pkey";
ALTER TABLE ONLY public."TstPlan" DROP CONSTRAINT "TstPlan_pkey";
ALTER TABLE ONLY public."TstOrg" DROP CONSTRAINT "TstOrg_pkey";
ALTER TABLE ONLY public."TstOrgRole" DROP CONSTRAINT "TstOrgRole_pkey";
ALTER TABLE ONLY public."TstOrgPrivilegeDefine" DROP CONSTRAINT "TstOrgPrivilegeDefine_pkey";
ALTER TABLE ONLY public."TstOrgGroup" DROP CONSTRAINT "TstOrgGroup_pkey";
ALTER TABLE ONLY public."TstMsg" DROP CONSTRAINT "TstMsg_pkey";
ALTER TABLE ONLY public."TstModule" DROP CONSTRAINT "TstModule_pkey";
ALTER TABLE ONLY public."TstHistory" DROP CONSTRAINT "TstHistory_pkey";
ALTER TABLE ONLY public."TstEnv" DROP CONSTRAINT "TstEnv_pkey";
ALTER TABLE ONLY public."TstEmail" DROP CONSTRAINT "TstEmail_pkey";
ALTER TABLE ONLY public."TstDocument" DROP CONSTRAINT "TstDocument_pkey";
ALTER TABLE ONLY public."TstCase" DROP CONSTRAINT "TstCase_pkey";
ALTER TABLE ONLY public."TstCaseType" DROP CONSTRAINT "TstCaseType_pkey";
ALTER TABLE ONLY public."TstCaseTypeDefine" DROP CONSTRAINT "TstCaseTypeDefine_pkey";
ALTER TABLE ONLY public."TstCaseStep" DROP CONSTRAINT "TstCaseStep_pkey";
ALTER TABLE ONLY public."TstCasePriority" DROP CONSTRAINT "TstCasePriority_pkey";
ALTER TABLE ONLY public."TstCasePriorityDefine" DROP CONSTRAINT "TstCasePriorityDefine_pkey";
ALTER TABLE ONLY public."TstCaseInTask" DROP CONSTRAINT "TstCaseInTask_pkey";
ALTER TABLE ONLY public."TstCaseInTaskIssue" DROP CONSTRAINT "TstCaseInTaskIssue_pkey";
ALTER TABLE ONLY public."TstCaseInTaskHistory" DROP CONSTRAINT "TstCaseInTaskHistory_pkey";
ALTER TABLE ONLY public."TstCaseInTaskComments" DROP CONSTRAINT "TstCaseInTaskComments_pkey";
ALTER TABLE ONLY public."TstCaseInTaskAttachment" DROP CONSTRAINT "TstCaseInTaskAttachment_pkey";
ALTER TABLE ONLY public."TstCaseInSuite" DROP CONSTRAINT "TstCaseInSuite_pkey";
ALTER TABLE ONLY public."TstCaseHistory" DROP CONSTRAINT "TstCaseHistory_pkey";
ALTER TABLE ONLY public."TstCaseExeStatus" DROP CONSTRAINT "TstCaseExeStatus_pkey";
ALTER TABLE ONLY public."TstCaseExeStatusDefine" DROP CONSTRAINT "TstCaseExeStatusDefine_pkey";
ALTER TABLE ONLY public."TstCaseComments" DROP CONSTRAINT "TstCaseComments_pkey";
ALTER TABLE ONLY public."TstCaseAttachment" DROP CONSTRAINT "TstCaseAttachment_pkey";
ALTER TABLE ONLY public."TstAlert" DROP CONSTRAINT "TstAlert_pkey";
ALTER TABLE ONLY public."Test" DROP CONSTRAINT "Test_pkey";
ALTER TABLE ONLY public."SysUser" DROP CONSTRAINT "SysUser_pkey";
ALTER TABLE ONLY public."SysRole" DROP CONSTRAINT "SysRole_pkey";
ALTER TABLE ONLY public."SysPrivilege" DROP CONSTRAINT "SysPrivilege_pkey";
ALTER TABLE ONLY public."IsuWorkflow" DROP CONSTRAINT "IsuWorkflow_pkey";
ALTER TABLE ONLY public."IsuWorkflowTransition" DROP CONSTRAINT "IsuWorkflowTransition_pkey";
ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation" DROP CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_pkey";
ALTER TABLE ONLY public."IsuWorkflowTransitionDefine" DROP CONSTRAINT "IsuWorkflowTransitionDefine_pkey";
ALTER TABLE ONLY public."IsuWorkflowStatusRelation" DROP CONSTRAINT "IsuWorkflowStatusRelation_pkey";
ALTER TABLE ONLY public."IsuWorkflowStatusRelationDefine" DROP CONSTRAINT "IsuWorkflowStatusRelationDefine_pkey";
ALTER TABLE ONLY public."IsuWorkflowSolution" DROP CONSTRAINT "IsuWorkflowSolution_pkey";
ALTER TABLE ONLY public."IsuWorkflowSolutionItem" DROP CONSTRAINT "IsuWorkflowSolutionItem_pkey";
ALTER TABLE ONLY public."IsuWatch" DROP CONSTRAINT "IsuWatch_pkey";
ALTER TABLE ONLY public."IsuType" DROP CONSTRAINT "IsuType_pkey";
ALTER TABLE ONLY public."IsuTypeSolution" DROP CONSTRAINT "IsuTypeSolution_pkey";
ALTER TABLE ONLY public."IsuTypeDefine" DROP CONSTRAINT "IsuTypeDefine_pkey";
ALTER TABLE ONLY public."IsuTag" DROP CONSTRAINT "IsuTag_pkey";
ALTER TABLE ONLY public."IsuTagRelation" DROP CONSTRAINT "IsuTagRelation_pkey";
ALTER TABLE ONLY public."IsuStatus" DROP CONSTRAINT "IsuStatus_pkey";
ALTER TABLE ONLY public."IsuStatusDefine" DROP CONSTRAINT "IsuStatusDefine_pkey";
ALTER TABLE ONLY public."IsuStatusCategoryDefine" DROP CONSTRAINT "IsuStatusCategoryDefine_pkey";
ALTER TABLE ONLY public."IsuSeverity" DROP CONSTRAINT "IsuSeverity_pkey";
ALTER TABLE ONLY public."IsuSeveritySolution" DROP CONSTRAINT "IsuSeveritySolution_pkey";
ALTER TABLE ONLY public."IsuSeverityDefine" DROP CONSTRAINT "IsuSeverityDefine_pkey";
ALTER TABLE ONLY public."IsuResolution" DROP CONSTRAINT "IsuResolution_pkey";
ALTER TABLE ONLY public."IsuResolutionDefine" DROP CONSTRAINT "IsuResolutionDefine_pkey";
ALTER TABLE ONLY public."IsuQuery" DROP CONSTRAINT "IsuQuery_pkey";
ALTER TABLE ONLY public."IsuPriority" DROP CONSTRAINT "IsuPriority_pkey";
ALTER TABLE ONLY public."IsuPrioritySolution" DROP CONSTRAINT "IsuPrioritySolution_pkey";
ALTER TABLE ONLY public."IsuPriorityDefine" DROP CONSTRAINT "IsuPriorityDefine_pkey";
ALTER TABLE ONLY public."IsuPage" DROP CONSTRAINT "IsuPage_pkey";
ALTER TABLE ONLY public."IsuPageSolution" DROP CONSTRAINT "IsuPageSolution_pkey";
ALTER TABLE ONLY public."IsuPageSolutionItem" DROP CONSTRAINT "IsuPageSolutionItem_pkey";
ALTER TABLE ONLY public."IsuPageElement" DROP CONSTRAINT "IsuPageElement_pkey";
ALTER TABLE ONLY public."IsuNotification" DROP CONSTRAINT "IsuNotification_pkey";
ALTER TABLE ONLY public."IsuNotificationDefine" DROP CONSTRAINT "IsuNotificationDefine_pkey";
ALTER TABLE ONLY public."IsuLink" DROP CONSTRAINT "IsuLink_pkey";
ALTER TABLE ONLY public."IsuLinkReasonDefine" DROP CONSTRAINT "IsuLinkReasonDefine_pkey";
ALTER TABLE ONLY public."IsuIssue" DROP CONSTRAINT "IsuIssue_pkey";
ALTER TABLE ONLY public."IsuIssueExt" DROP CONSTRAINT "IsuIssueExt_pkey";
ALTER TABLE ONLY public."IsuHistory" DROP CONSTRAINT "IsuHistory_pkey";
ALTER TABLE ONLY public."IsuField" DROP CONSTRAINT "IsuField_pkey";
ALTER TABLE ONLY public."IsuFieldDefine" DROP CONSTRAINT "IsuFieldDefine_pkey";
ALTER TABLE ONLY public."IsuFieldCodeToTableDefine" DROP CONSTRAINT "IsuFieldCodeToTableDefine_pkey";
ALTER TABLE ONLY public."IsuDocument" DROP CONSTRAINT "IsuDocument_pkey";
ALTER TABLE ONLY public."IsuCustomFieldSolution" DROP CONSTRAINT "IsuCustomFieldSolution_pkey";
ALTER TABLE ONLY public."IsuComments" DROP CONSTRAINT "IsuComments_pkey";
ALTER TABLE ONLY public."IsuAttachment" DROP CONSTRAINT "IsuAttachment_pkey";
ALTER TABLE ONLY public."CustomField" DROP CONSTRAINT "CustomField_pkey";
ALTER TABLE ONLY public."CustomFieldTypeDefine" DROP CONSTRAINT "CustomFieldTypeDefine_pkey";
ALTER TABLE ONLY public."CustomFieldOption" DROP CONSTRAINT "CustomFieldOption_pkey";
ALTER TABLE ONLY public."CustomFieldOptionDefine" DROP CONSTRAINT "CustomFieldOptionDefine_pkey";
ALTER TABLE ONLY public."CustomFieldIputDefine" DROP CONSTRAINT "CustomFieldIputDefine_pkey";
ALTER TABLE ONLY public."CustomFieldInputTypeRelationDefine" DROP CONSTRAINT "CustomFieldInputTypeRelationDefine_pkey";
ALTER TABLE ONLY public."CustomFieldDefine" DROP CONSTRAINT "CustomFieldDefine_pkey";
ALTER TABLE public."TstVer" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstUserVerifyCode" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstUser" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstThread" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstTask" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstSuite" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstProjectRole" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstProjectPrivilegeDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstProjectAccessHistory" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstProject" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstPlan" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstOrgRole" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstOrgPrivilegeDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstOrgGroup" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstOrg" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstMsg" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstModule" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstHistory" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstEnv" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstEmail" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstDocument" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseTypeDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseStep" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCasePriorityDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseInTaskIssue" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseInTaskHistory" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseInTaskComments" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseInTaskAttachment" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseInTask" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseInSuite" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseHistory" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseExeStatusDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseComments" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCaseAttachment" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstCase" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."TstAlert" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."Test" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."SysUser" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."SysRole" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."SysPrivilege" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuWorkflowTransitionProjectRoleRelation" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuWorkflowTransitionDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuWorkflowTransition" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuWorkflowStatusRelationDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuWorkflowStatusRelation" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuWorkflowSolutionItem" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuWorkflowSolution" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuWorkflow" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuWatch" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuTypeSolution" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuTypeDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuType" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuTagRelation" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuTag" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuStatusDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuStatusCategoryDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuStatus" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuSeveritySolution" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuSeverityDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuSeverity" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuResolutionDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuResolution" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuQuery" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuPrioritySolution" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuPriorityDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuPriority" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuPageSolutionItem" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuPageSolution" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuPageElement" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuPage" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuNotificationDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuNotification" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuLinkReasonDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuLink" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuIssueExt" ALTER COLUMN pid DROP DEFAULT;
ALTER TABLE public."IsuIssue" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuHistory" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuFieldDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuFieldCodeToTableDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuField" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuDocument" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuCustomFieldSolution" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuComments" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."IsuAttachment" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."CustomFieldTypeDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."CustomFieldOptionDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."CustomFieldOption" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."CustomFieldIputDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."CustomFieldInputTypeRelationDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."CustomFieldDefine" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public."CustomField" ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE public."TstVer_id_seq";
DROP TABLE public."TstVer";
DROP SEQUENCE public."TstUser_id_seq";
DROP SEQUENCE public."TstUserVerifyCode_id_seq";
DROP TABLE public."TstUserVerifyCode";
DROP TABLE public."TstUserSettings";
DROP TABLE public."TstUser";
DROP SEQUENCE public."TstThread_id_seq";
DROP TABLE public."TstThread";
DROP SEQUENCE public."TstTask_id_seq";
DROP TABLE public."TstTaskAssigneeRelation";
DROP TABLE public."TstTask";
DROP SEQUENCE public."TstSuite_id_seq";
DROP TABLE public."TstSuite";
DROP SEQUENCE public."TstProject_id_seq";
DROP SEQUENCE public."TstProjectRole_id_seq";
DROP TABLE public."TstProjectRolePriviledgeRelation";
DROP TABLE public."TstProjectRoleEntityRelation";
DROP TABLE public."TstProjectRole";
DROP SEQUENCE public."TstProjectPrivilegeDefine_id_seq";
DROP TABLE public."TstProjectPrivilegeDefine";
DROP SEQUENCE public."TstProjectAccessHistory_id_seq";
DROP TABLE public."TstProjectAccessHistory";
DROP TABLE public."TstProject";
DROP SEQUENCE public."TstPlan_id_seq";
DROP TABLE public."TstPlan";
DROP SEQUENCE public."TstOrg_id_seq";
DROP TABLE public."TstOrgUserRelation";
DROP SEQUENCE public."TstOrgRole_id_seq";
DROP TABLE public."TstOrgRoleUserRelation";
DROP TABLE public."TstOrgRolePrivilegeRelation";
DROP TABLE public."TstOrgRoleGroupRelation";
DROP TABLE public."TstOrgRole";
DROP SEQUENCE public."TstOrgPrivilegeDefine_id_seq";
DROP TABLE public."TstOrgPrivilegeDefine";
DROP SEQUENCE public."TstOrgGroup_id_seq";
DROP TABLE public."TstOrgGroupUserRelation";
DROP TABLE public."TstOrgGroup";
DROP TABLE public."TstOrg";
DROP SEQUENCE public."TstMsg_id_seq";
DROP TABLE public."TstMsg";
DROP SEQUENCE public."TstModule_id_seq";
DROP TABLE public."TstModule";
DROP SEQUENCE public."TstHistory_id_seq";
DROP TABLE public."TstHistory";
DROP SEQUENCE public."TstEnv_id_seq";
DROP TABLE public."TstEnv";
DROP SEQUENCE public."TstEmail_id_seq";
DROP TABLE public."TstEmail";
DROP SEQUENCE public."TstDocument_id_seq";
DROP TABLE public."TstDocument";
DROP SEQUENCE public."TstCase_id_seq";
DROP TABLE public."TstCaseType";
DROP SEQUENCE public."TstCaseType_id_seq";
DROP TABLE public."TstCaseTypeDefine";
DROP SEQUENCE public."TstCaseStep_id_seq";
DROP TABLE public."TstCaseStep";
DROP TABLE public."TstCasePriority";
DROP SEQUENCE public."TstCasePriority_id_seq";
DROP TABLE public."TstCasePriorityDefine";
DROP SEQUENCE public."TstCaseInTask_id_seq";
DROP SEQUENCE public."TstCaseInTaskIssue_id_seq";
DROP TABLE public."TstCaseInTaskIssue";
DROP SEQUENCE public."TstCaseInTaskHistory_id_seq";
DROP TABLE public."TstCaseInTaskHistory";
DROP SEQUENCE public."TstCaseInTaskComments_id_seq";
DROP TABLE public."TstCaseInTaskComments";
DROP SEQUENCE public."TstCaseInTaskAttachment_id_seq";
DROP TABLE public."TstCaseInTaskAttachment";
DROP TABLE public."TstCaseInTask";
DROP SEQUENCE public."TstCaseInSuite_id_seq";
DROP TABLE public."TstCaseInSuite";
DROP SEQUENCE public."TstCaseHistory_id_seq";
DROP TABLE public."TstCaseHistory";
DROP TABLE public."TstCaseExeStatus";
DROP SEQUENCE public."TstCaseExeStatus_id_seq";
DROP TABLE public."TstCaseExeStatusDefine";
DROP SEQUENCE public."TstCaseComments_id_seq";
DROP TABLE public."TstCaseComments";
DROP SEQUENCE public."TstCaseAttachment_id_seq";
DROP TABLE public."TstCaseAttachment";
DROP TABLE public."TstCase";
DROP SEQUENCE public."TstAlert_id_seq";
DROP TABLE public."TstAlert";
DROP SEQUENCE public."Test_id_seq";
DROP TABLE public."Test";
DROP SEQUENCE public."SysUser_id_seq";
DROP TABLE public."SysUser";
DROP SEQUENCE public."SysRole_id_seq";
DROP TABLE public."SysRoleUserRelation";
DROP TABLE public."SysRolePrivilegeRelation";
DROP TABLE public."SysRole";
DROP SEQUENCE public."SysPrivilege_id_seq";
DROP TABLE public."SysPrivilege";
DROP SEQUENCE public."IsuWorkflow_id_seq";
DROP SEQUENCE public."IsuWorkflowTransition_id_seq";
DROP SEQUENCE public."IsuWorkflowTransitionProjectRoleRelation_id_seq";
DROP TABLE public."IsuWorkflowTransitionProjectRoleRelation";
DROP SEQUENCE public."IsuWorkflowTransitionDefine_id_seq";
DROP TABLE public."IsuWorkflowTransitionDefine";
DROP TABLE public."IsuWorkflowTransition";
DROP SEQUENCE public."IsuWorkflowStatusRelation_id_seq";
DROP SEQUENCE public."IsuWorkflowStatusRelationDefine_id_seq";
DROP TABLE public."IsuWorkflowStatusRelationDefine";
DROP TABLE public."IsuWorkflowStatusRelation";
DROP SEQUENCE public."IsuWorkflowSolution_id_seq";
DROP SEQUENCE public."IsuWorkflowSolutionItem_id_seq";
DROP TABLE public."IsuWorkflowSolutionItem";
DROP TABLE public."IsuWorkflowSolution";
DROP TABLE public."IsuWorkflow";
DROP SEQUENCE public."IsuWatch_id_seq";
DROP TABLE public."IsuWatch";
DROP SEQUENCE public."IsuType_id_seq";
DROP SEQUENCE public."IsuTypeSolution_id_seq";
DROP TABLE public."IsuTypeSolutionItem";
DROP TABLE public."IsuTypeSolution";
DROP SEQUENCE public."IsuTypeDefine_id_seq";
DROP TABLE public."IsuTypeDefine";
DROP TABLE public."IsuType";
DROP SEQUENCE public."IsuTag_id_seq";
DROP SEQUENCE public."IsuTagRelation_id_seq";
DROP TABLE public."IsuTagRelation";
DROP TABLE public."IsuTag";
DROP SEQUENCE public."IsuStatus_id_seq";
DROP SEQUENCE public."IsuStatusDefine_id_seq";
DROP TABLE public."IsuStatusDefine";
DROP SEQUENCE public."IsuStatusCategoryDefine_id_seq";
DROP TABLE public."IsuStatusCategoryDefine";
DROP TABLE public."IsuStatus";
DROP SEQUENCE public."IsuSeverity_id_seq";
DROP SEQUENCE public."IsuSeveritySolution_id_seq";
DROP TABLE public."IsuSeveritySolutionItem";
DROP TABLE public."IsuSeveritySolution";
DROP SEQUENCE public."IsuSeverityDefine_id_seq";
DROP TABLE public."IsuSeverityDefine";
DROP TABLE public."IsuSeverity";
DROP SEQUENCE public."IsuResolution_id_seq";
DROP SEQUENCE public."IsuResolutionDefine_id_seq";
DROP TABLE public."IsuResolutionDefine";
DROP TABLE public."IsuResolution";
DROP SEQUENCE public."IsuQuery_id_seq";
DROP TABLE public."IsuQuery";
DROP SEQUENCE public."IsuPriority_id_seq";
DROP SEQUENCE public."IsuPrioritySolution_id_seq";
DROP TABLE public."IsuPrioritySolutionItem";
DROP TABLE public."IsuPrioritySolution";
DROP SEQUENCE public."IsuPriorityDefine_id_seq";
DROP TABLE public."IsuPriorityDefine";
DROP TABLE public."IsuPriority";
DROP SEQUENCE public."IsuPage_id_seq";
DROP SEQUENCE public."IsuPageSolution_id_seq";
DROP SEQUENCE public."IsuPageSolutionItem_id_seq";
DROP TABLE public."IsuPageSolutionItem";
DROP TABLE public."IsuPageSolution";
DROP SEQUENCE public."IsuPageElement_id_seq";
DROP TABLE public."IsuPageElement";
DROP TABLE public."IsuPage";
DROP SEQUENCE public."IsuNotification_id_seq";
DROP SEQUENCE public."IsuNotificationDefine_id_seq";
DROP TABLE public."IsuNotificationDefine";
DROP TABLE public."IsuNotification";
DROP SEQUENCE public."IsuLink_id_seq";
DROP SEQUENCE public."IsuLinkReasonDefine_id_seq";
DROP TABLE public."IsuLinkReasonDefine";
DROP TABLE public."IsuLink";
DROP SEQUENCE public."IsuIssue_id_seq";
DROP SEQUENCE public."IsuIssueExt_pid_seq";
DROP TABLE public."IsuIssueExt";
DROP TABLE public."IsuIssue";
DROP SEQUENCE public."IsuHistory_id_seq";
DROP TABLE public."IsuHistory";
DROP SEQUENCE public."IsuField_id_seq";
DROP SEQUENCE public."IsuFieldDefine_id_seq";
DROP TABLE public."IsuFieldDefine";
DROP SEQUENCE public."IsuFieldCodeToTableDefine_id_seq";
DROP TABLE public."IsuFieldCodeToTableDefine";
DROP TABLE public."IsuField";
DROP SEQUENCE public."IsuDocument_id_seq";
DROP TABLE public."IsuDocument";
DROP SEQUENCE public."IsuCustomFieldSolution_id_seq";
DROP TABLE public."IsuCustomFieldSolutionProjectRelation";
DROP TABLE public."IsuCustomFieldSolutionFieldRelation";
DROP TABLE public."IsuCustomFieldSolution";
DROP SEQUENCE public."IsuComments_id_seq";
DROP TABLE public."IsuComments";
DROP SEQUENCE public."IsuAttachment_id_seq";
DROP TABLE public."IsuAttachment";
DROP SEQUENCE public."CustomField_id_seq";
DROP SEQUENCE public."CustomFieldTypeDefine_id_seq";
DROP TABLE public."CustomFieldTypeDefine";
DROP SEQUENCE public."CustomFieldOption_id_seq";
DROP SEQUENCE public."CustomFieldOptionDefine_id_seq";
DROP TABLE public."CustomFieldOptionDefine";
DROP TABLE public."CustomFieldOption";
DROP SEQUENCE public."CustomFieldIputDefine_id_seq";
DROP TABLE public."CustomFieldIputDefine";
DROP SEQUENCE public."CustomFieldInputTypeRelationDefine_id_seq";
DROP TABLE public."CustomFieldInputTypeRelationDefine";
DROP SEQUENCE public."CustomFieldDefine_id_seq";
DROP TABLE public."CustomFieldDefine";
DROP TABLE public."CustomField";
DROP TEXT SEARCH CONFIGURATION public.chinese_zh;
DROP FUNCTION public.user_not_in_project(p_user_id integer, p_project_id integer);
DROP FUNCTION public.update_workflow_statuses(p_workflow_id integer, p_status_ids character varying, p_org_id integer);
DROP FUNCTION public.update_workflow_statuses(p_workflow_id integer, p_status_ids character varying);
DROP FUNCTION public.update_issue_tsv_content();
DROP FUNCTION public.test(_p integer);
DROP FUNCTION public.remove_user_from_org(p_user_id integer, p_org_id integer);
DROP FUNCTION public.remove_case_and_its_children(p_case_id integer);
DROP FUNCTION public.remove_all_tables();
DROP FUNCTION public.remove_all();
DROP FUNCTION public.init_user(p_user_id integer, p_org_name character varying);
DROP FUNCTION public.init_org_issue_page_solution_item(p_issue_page_id integer, p_issue_page_solution_id integer, p_org_id integer);
DROP FUNCTION public.init_org_custom_field_option(p_org_id integer);
DROP FUNCTION public.init_org(p_org_id integer, p_user_id integer);
DROP FUNCTION public.get_project_privilege_for_user(p_user_id integer, p_project_id integer, p_project_type character varying);
DROP FUNCTION public.get_org_privilege_for_user(p_user_id integer, p_org_id integer);
DROP FUNCTION public.gen_project_access_history(p_org_id integer, p_project_id integer, p_project_name character varying, p_user_id integer);
DROP FUNCTION public.close_plan_if_all_task_closed(p_plan_id integer);
DROP FUNCTION public.chart_test_execution_result_by_plan(p_plan_id integer);
DROP FUNCTION public.chart_test_execution_progress_by_plan(p_plan_id integer, p_day_numb integer);
DROP FUNCTION public.chart_test_execution_process_by_project(p_project_id integer, p_project_type character varying, p_day_numb integer);
DROP FUNCTION public.chart_test_execution_process_by_plan_user(p_plan_id integer, p_day_numb integer);
DROP FUNCTION public.chart_test_execution_process_by_plan(p_plan_id integer, p_day_numb integer);
DROP FUNCTION public.chart_test_design_progress_by_project(p_project_id integer, p_project_type character varying, p_day_numb integer);
DROP FUNCTION public.chart_issue_trend_final(p_project_id integer, p_project_type character varying, p_day_numb integer);
DROP FUNCTION public.chart_issue_trend_create(p_project_id integer, p_project_type character varying, p_day_numb integer);
DROP FUNCTION public.chart_issue_distrib_by_status(p_project_id integer, p_project_type character varying);
DROP FUNCTION public.chart_issue_distrib_by_priority(p_project_id integer, p_project_type character varying);
DROP FUNCTION public.chart_issue_age(p_project_id integer, p_project_type character varying, p_day_numb integer);
DROP FUNCTION public.add_cases_to_task_by_suites(p_suite_ids character varying, p_task_id integer);
DROP FUNCTION public.add_cases_to_task(p_case_ids character varying, p_task_id integer);
DROP FUNCTION public.add_cases_to_suite(p_case_ids character varying, p_suite_id integer);
DROP FUNCTION public.add_case_to_task(p_case_id integer, p_task_id integer, p_plan_id integer, p_project_id integer);
DROP FUNCTION public.add_case_to_suite(p_case_id integer, p_suite_id integer, p_project_id integer);
DROP FUNCTION public._user_project_role(p_user_id integer, p_project_id integer, p_project_type character varying);
DROP FUNCTION public._user_project_role(p_user_id integer);
DROP FUNCTION public._user_org_role(p_user_id integer, p_org_id integer);
DROP FUNCTION public._user_org_role(p_user_id integer);
DROP FUNCTION public._project_user(p_project_id integer);
DROP FUNCTION public._project_list(p_project_id integer, p_project_type character varying);
DROP FUNCTION public._date_list(p_time_before timestamp without time zone);
DROP FUNCTION public._date_before(p_day_numb integer);
DROP EXTENSION zhparser;
--
-- Name: zhparser; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS zhparser WITH SCHEMA public;


--
-- Name: EXTENSION zhparser; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION zhparser IS 'a parser for full-text search of Chinese';


--
-- Name: _date_before(integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public._date_before(p_day_numb integer) OWNER TO ngtesting;

--
-- Name: _date_list(timestamp without time zone); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public._date_list(p_time_before timestamp without time zone) OWNER TO ngtesting;

--
-- Name: _project_list(integer, character varying); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public._project_list(p_project_id integer, p_project_type character varying) OWNER TO ngtesting;

--
-- Name: _project_user(integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public._project_user(p_project_id integer) OWNER TO ngtesting;

--
-- Name: _user_org_role(integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public._user_org_role(p_user_id integer) OWNER TO ngtesting;

--
-- Name: _user_org_role(integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public._user_org_role(p_user_id integer, p_org_id integer) OWNER TO ngtesting;

--
-- Name: _user_project_role(integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public._user_project_role(p_user_id integer) OWNER TO ngtesting;

--
-- Name: _user_project_role(integer, integer, character varying); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public._user_project_role(p_user_id integer, p_project_id integer, p_project_type character varying) OWNER TO ngtesting;

--
-- Name: add_case_to_suite(integer, integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.add_case_to_suite(p_case_id integer, p_suite_id integer, p_project_id integer) OWNER TO ngtesting;

--
-- Name: add_case_to_task(integer, integer, integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.add_case_to_task(p_case_id integer, p_task_id integer, p_plan_id integer, p_project_id integer) OWNER TO ngtesting;

--
-- Name: add_cases_to_suite(character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.add_cases_to_suite(p_case_ids character varying, p_suite_id integer) OWNER TO ngtesting;

--
-- Name: add_cases_to_task(character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.add_cases_to_task(p_case_ids character varying, p_task_id integer) OWNER TO ngtesting;

--
-- Name: add_cases_to_task_by_suites(character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.add_cases_to_task_by_suites(p_suite_ids character varying, p_task_id integer) OWNER TO ngtesting;

--
-- Name: chart_issue_age(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_issue_age(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO ngtesting;

--
-- Name: chart_issue_distrib_by_priority(integer, character varying); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_issue_distrib_by_priority(p_project_id integer, p_project_type character varying) OWNER TO ngtesting;

--
-- Name: chart_issue_distrib_by_status(integer, character varying); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_issue_distrib_by_status(p_project_id integer, p_project_type character varying) OWNER TO ngtesting;

--
-- Name: chart_issue_trend_create(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_issue_trend_create(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO ngtesting;

--
-- Name: chart_issue_trend_final(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_issue_trend_final(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO ngtesting;

--
-- Name: chart_test_design_progress_by_project(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_test_design_progress_by_project(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO ngtesting;

--
-- Name: chart_test_execution_process_by_plan(integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_test_execution_process_by_plan(p_plan_id integer, p_day_numb integer) OWNER TO ngtesting;

--
-- Name: chart_test_execution_process_by_plan_user(integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_test_execution_process_by_plan_user(p_plan_id integer, p_day_numb integer) OWNER TO ngtesting;

--
-- Name: chart_test_execution_process_by_project(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_test_execution_process_by_project(p_project_id integer, p_project_type character varying, p_day_numb integer) OWNER TO ngtesting;

--
-- Name: chart_test_execution_progress_by_plan(integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_test_execution_progress_by_plan(p_plan_id integer, p_day_numb integer) OWNER TO ngtesting;

--
-- Name: chart_test_execution_result_by_plan(integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.chart_test_execution_result_by_plan(p_plan_id integer) OWNER TO ngtesting;

--
-- Name: close_plan_if_all_task_closed(integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.close_plan_if_all_task_closed(p_plan_id integer) OWNER TO ngtesting;

--
-- Name: gen_project_access_history(integer, integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.gen_project_access_history(p_org_id integer, p_project_id integer, p_project_name character varying, p_user_id integer) OWNER TO ngtesting;

--
-- Name: get_org_privilege_for_user(integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
--

CREATE FUNCTION public.get_org_privilege_for_user(p_user_id integer, p_org_id integer) RETURNS TABLE(code character varying, name character varying)
    LANGUAGE plpgsql
    AS $$  
declare  
	
BEGIN
	RETURN QUERY
	
	SELECT priv."code",priv.name 
	from "TstOrgPrivilegeDefine" priv
	join "TstOrgRolePrivilegeRelation" r_role_priv ON r_role_priv."orgPrivilegeId"=priv.id
	where r_role_priv."orgRoleId" = any (select "orgRoleId" from _user_org_role(p_user_id, p_org_id))
		AND NOT priv.deleted and NOT priv.disabled
	order by priv.id asc;

END;  
$$;


ALTER FUNCTION public.get_org_privilege_for_user(p_user_id integer, p_org_id integer) OWNER TO ngtesting;

--
-- Name: get_project_privilege_for_user(integer, integer, character varying); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.get_project_privilege_for_user(p_user_id integer, p_project_id integer, p_project_type character varying) OWNER TO ngtesting;

--
-- Name: init_org(integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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
    select d.id,project_role_id from "TstProjectPrivilegeDefine" d where d.id != 12400;

    insert into "TstProjectRole" (code, name, "buildIn", "orgId", disabled, deleted, "createTime")
    values('tester', '测试执行', false, p_org_id, false, false, NOW());
    select max(id) from "TstProjectRole" into project_role_id;

    insert into "TstProjectRolePriviledgeRelation" ( "projectPrivilegeDefineId",   "projectRoleId" )
    select d.id,project_role_id from "TstProjectPrivilegeDefine" d where d.id != 12200 and d.id != 12400;

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


ALTER FUNCTION public.init_org(p_org_id integer, p_user_id integer) OWNER TO ngtesting;

--
-- Name: init_org_custom_field_option(integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.init_org_custom_field_option(p_org_id integer) OWNER TO ngtesting;

--
-- Name: init_org_issue_page_solution_item(integer, integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.init_org_issue_page_solution_item(p_issue_page_id integer, p_issue_page_solution_id integer, p_org_id integer) OWNER TO ngtesting;

--
-- Name: init_user(integer, character varying); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.init_user(p_user_id integer, p_org_name character varying) OWNER TO ngtesting;

--
-- Name: remove_all(); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.remove_all() OWNER TO ngtesting;

--
-- Name: remove_all_tables(); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.remove_all_tables() OWNER TO ngtesting;

--
-- Name: remove_case_and_its_children(integer); Type: FUNCTION; Schema: public; Owner: ngtesting
--

CREATE FUNCTION public.remove_case_and_its_children(p_case_id integer) RETURNS integer
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
		SELECT cs.id FROM cs ORDER BY cs.id
	);
	
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	RAISE NOTICE 'v_cnt = %', v_cnt;
	
    RETURN v_cnt;  
END;  
$$;


ALTER FUNCTION public.remove_case_and_its_children(p_case_id integer) OWNER TO ngtesting;

--
-- Name: remove_user_from_org(integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.remove_user_from_org(p_user_id integer, p_org_id integer) OWNER TO ngtesting;

--
-- Name: test(integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.test(_p integer) OWNER TO ngtesting;

--
-- Name: update_issue_tsv_content(); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.update_issue_tsv_content() OWNER TO ngtesting;

--
-- Name: update_workflow_statuses(integer, character varying); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.update_workflow_statuses(p_workflow_id integer, p_status_ids character varying) OWNER TO ngtesting;

--
-- Name: update_workflow_statuses(integer, character varying, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.update_workflow_statuses(p_workflow_id integer, p_status_ids character varying, p_org_id integer) OWNER TO ngtesting;

--
-- Name: user_not_in_project(integer, integer); Type: FUNCTION; Schema: public; Owner: ngtesting
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


ALTER FUNCTION public.user_not_in_project(p_user_id integer, p_project_id integer) OWNER TO ngtesting;

--
-- Name: chinese_zh; Type: TEXT SEARCH CONFIGURATION; Schema: public; Owner: ngtesting
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


ALTER TEXT SEARCH CONFIGURATION public.chinese_zh OWNER TO ngtesting;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: CustomField; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."CustomField" OWNER TO ngtesting;

--
-- Name: CustomFieldDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."CustomFieldDefine" OWNER TO ngtesting;

--
-- Name: CustomFieldDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."CustomFieldDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldDefine_id_seq" OWNER TO ngtesting;

--
-- Name: CustomFieldDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."CustomFieldDefine_id_seq" OWNED BY public."CustomFieldDefine".id;


--
-- Name: CustomFieldInputTypeRelationDefine; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."CustomFieldInputTypeRelationDefine" (
    id integer NOT NULL,
    "inputValue" character varying(255),
    "typeValue" character varying(255)
);


ALTER TABLE public."CustomFieldInputTypeRelationDefine" OWNER TO ngtesting;

--
-- Name: CustomFieldInputTypeRelationDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."CustomFieldInputTypeRelationDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldInputTypeRelationDefine_id_seq" OWNER TO ngtesting;

--
-- Name: CustomFieldInputTypeRelationDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."CustomFieldInputTypeRelationDefine_id_seq" OWNED BY public."CustomFieldInputTypeRelationDefine".id;


--
-- Name: CustomFieldIputDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."CustomFieldIputDefine" OWNER TO ngtesting;

--
-- Name: CustomFieldIputDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."CustomFieldIputDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldIputDefine_id_seq" OWNER TO ngtesting;

--
-- Name: CustomFieldIputDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."CustomFieldIputDefine_id_seq" OWNED BY public."CustomFieldIputDefine".id;


--
-- Name: CustomFieldOption; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."CustomFieldOption" OWNER TO ngtesting;

--
-- Name: CustomFieldOptionDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."CustomFieldOptionDefine" OWNER TO ngtesting;

--
-- Name: CustomFieldOptionDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."CustomFieldOptionDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldOptionDefine_id_seq" OWNER TO ngtesting;

--
-- Name: CustomFieldOptionDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."CustomFieldOptionDefine_id_seq" OWNED BY public."CustomFieldOptionDefine".id;


--
-- Name: CustomFieldOption_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."CustomFieldOption_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldOption_id_seq" OWNER TO ngtesting;

--
-- Name: CustomFieldOption_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."CustomFieldOption_id_seq" OWNED BY public."CustomFieldOption".id;


--
-- Name: CustomFieldTypeDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."CustomFieldTypeDefine" OWNER TO ngtesting;

--
-- Name: CustomFieldTypeDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."CustomFieldTypeDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomFieldTypeDefine_id_seq" OWNER TO ngtesting;

--
-- Name: CustomFieldTypeDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."CustomFieldTypeDefine_id_seq" OWNED BY public."CustomFieldTypeDefine".id;


--
-- Name: CustomField_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."CustomField_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."CustomField_id_seq" OWNER TO ngtesting;

--
-- Name: CustomField_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."CustomField_id_seq" OWNED BY public."CustomField".id;


--
-- Name: IsuAttachment; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuAttachment" OWNER TO ngtesting;

--
-- Name: IsuAttachment_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuAttachment_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuAttachment_id_seq" OWNER TO ngtesting;

--
-- Name: IsuAttachment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuAttachment_id_seq" OWNED BY public."IsuAttachment".id;


--
-- Name: IsuComments; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuComments" OWNER TO ngtesting;

--
-- Name: IsuComments_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuComments_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuComments_id_seq" OWNER TO ngtesting;

--
-- Name: IsuComments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuComments_id_seq" OWNED BY public."IsuComments".id;


--
-- Name: IsuCustomFieldSolution; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuCustomFieldSolution" OWNER TO ngtesting;

--
-- Name: IsuCustomFieldSolutionFieldRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuCustomFieldSolutionFieldRelation" (
    "solutionId" integer,
    "fieldId" integer
);


ALTER TABLE public."IsuCustomFieldSolutionFieldRelation" OWNER TO ngtesting;

--
-- Name: IsuCustomFieldSolutionProjectRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuCustomFieldSolutionProjectRelation" (
    "solutionId" integer NOT NULL,
    "orgId" integer,
    "projectId" integer NOT NULL
);


ALTER TABLE public."IsuCustomFieldSolutionProjectRelation" OWNER TO ngtesting;

--
-- Name: IsuCustomFieldSolution_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuCustomFieldSolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuCustomFieldSolution_id_seq" OWNER TO ngtesting;

--
-- Name: IsuCustomFieldSolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuCustomFieldSolution_id_seq" OWNED BY public."IsuCustomFieldSolution".id;


--
-- Name: IsuDocument; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuDocument" OWNER TO ngtesting;

--
-- Name: IsuDocument_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuDocument_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuDocument_id_seq" OWNER TO ngtesting;

--
-- Name: IsuDocument_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuDocument_id_seq" OWNED BY public."IsuDocument".id;


--
-- Name: IsuField; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuField" OWNER TO ngtesting;

--
-- Name: IsuFieldCodeToTableDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuFieldCodeToTableDefine" OWNER TO ngtesting;

--
-- Name: IsuFieldCodeToTableDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuFieldCodeToTableDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuFieldCodeToTableDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuFieldCodeToTableDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuFieldCodeToTableDefine_id_seq" OWNED BY public."IsuFieldCodeToTableDefine".id;


--
-- Name: IsuFieldDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuFieldDefine" OWNER TO ngtesting;

--
-- Name: IsuFieldDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuFieldDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuFieldDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuFieldDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuFieldDefine_id_seq" OWNED BY public."IsuFieldDefine".id;


--
-- Name: IsuField_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuField_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuField_id_seq" OWNER TO ngtesting;

--
-- Name: IsuField_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuField_id_seq" OWNED BY public."IsuField".id;


--
-- Name: IsuHistory; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuHistory" OWNER TO ngtesting;

--
-- Name: IsuHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuHistory_id_seq" OWNER TO ngtesting;

--
-- Name: IsuHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuHistory_id_seq" OWNED BY public."IsuHistory".id;


--
-- Name: IsuIssue; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuIssue" OWNER TO ngtesting;

--
-- Name: IsuIssueExt; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuIssueExt" OWNER TO ngtesting;

--
-- Name: IsuIssueExt_pid_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuIssueExt_pid_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuIssueExt_pid_seq" OWNER TO ngtesting;

--
-- Name: IsuIssueExt_pid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuIssueExt_pid_seq" OWNED BY public."IsuIssueExt".pid;


--
-- Name: IsuIssue_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuIssue_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuIssue_id_seq" OWNER TO ngtesting;

--
-- Name: IsuIssue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuIssue_id_seq" OWNED BY public."IsuIssue".id;


--
-- Name: IsuLink; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuLink" OWNER TO ngtesting;

--
-- Name: IsuLinkReasonDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuLinkReasonDefine" OWNER TO ngtesting;

--
-- Name: IsuLinkReasonDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuLinkReasonDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuLinkReasonDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuLinkReasonDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuLinkReasonDefine_id_seq" OWNED BY public."IsuLinkReasonDefine".id;


--
-- Name: IsuLink_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuLink_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuLink_id_seq" OWNER TO ngtesting;

--
-- Name: IsuLink_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuLink_id_seq" OWNED BY public."IsuLink".id;


--
-- Name: IsuNotification; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuNotification" OWNER TO ngtesting;

--
-- Name: IsuNotificationDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuNotificationDefine" OWNER TO ngtesting;

--
-- Name: IsuNotificationDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuNotificationDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuNotificationDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuNotificationDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuNotificationDefine_id_seq" OWNED BY public."IsuNotificationDefine".id;


--
-- Name: IsuNotification_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuNotification_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuNotification_id_seq" OWNER TO ngtesting;

--
-- Name: IsuNotification_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuNotification_id_seq" OWNED BY public."IsuNotification".id;


--
-- Name: IsuPage; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuPage" OWNER TO ngtesting;

--
-- Name: IsuPageElement; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuPageElement" OWNER TO ngtesting;

--
-- Name: IsuPageElement_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuPageElement_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPageElement_id_seq" OWNER TO ngtesting;

--
-- Name: IsuPageElement_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuPageElement_id_seq" OWNED BY public."IsuPageElement".id;


--
-- Name: IsuPageSolution; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuPageSolution" OWNER TO ngtesting;

--
-- Name: IsuPageSolutionItem; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuPageSolutionItem" (
    id integer NOT NULL,
    "typeId" integer,
    opt character varying(255),
    "pageId" integer,
    "solutionId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuPageSolutionItem" OWNER TO ngtesting;

--
-- Name: IsuPageSolutionItem_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuPageSolutionItem_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPageSolutionItem_id_seq" OWNER TO ngtesting;

--
-- Name: IsuPageSolutionItem_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuPageSolutionItem_id_seq" OWNED BY public."IsuPageSolutionItem".id;


--
-- Name: IsuPageSolution_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuPageSolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPageSolution_id_seq" OWNER TO ngtesting;

--
-- Name: IsuPageSolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuPageSolution_id_seq" OWNED BY public."IsuPageSolution".id;


--
-- Name: IsuPage_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuPage_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPage_id_seq" OWNER TO ngtesting;

--
-- Name: IsuPage_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuPage_id_seq" OWNED BY public."IsuPage".id;


--
-- Name: IsuPriority; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuPriority" OWNER TO ngtesting;

--
-- Name: IsuPriorityDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuPriorityDefine" OWNER TO ngtesting;

--
-- Name: IsuPriorityDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuPriorityDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPriorityDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuPriorityDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuPriorityDefine_id_seq" OWNED BY public."IsuPriorityDefine".id;


--
-- Name: IsuPrioritySolution; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuPrioritySolution" OWNER TO ngtesting;

--
-- Name: IsuPrioritySolutionItem; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuPrioritySolutionItem" (
    "priorityId" integer,
    "solutionId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuPrioritySolutionItem" OWNER TO ngtesting;

--
-- Name: IsuPrioritySolution_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuPrioritySolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPrioritySolution_id_seq" OWNER TO ngtesting;

--
-- Name: IsuPrioritySolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuPrioritySolution_id_seq" OWNED BY public."IsuPrioritySolution".id;


--
-- Name: IsuPriority_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuPriority_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuPriority_id_seq" OWNER TO ngtesting;

--
-- Name: IsuPriority_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuPriority_id_seq" OWNED BY public."IsuPriority".id;


--
-- Name: IsuQuery; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuQuery" OWNER TO ngtesting;

--
-- Name: IsuQuery_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuQuery_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuQuery_id_seq" OWNER TO ngtesting;

--
-- Name: IsuQuery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuQuery_id_seq" OWNED BY public."IsuQuery".id;


--
-- Name: IsuResolution; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuResolution" OWNER TO ngtesting;

--
-- Name: IsuResolutionDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuResolutionDefine" OWNER TO ngtesting;

--
-- Name: IsuResolutionDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuResolutionDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuResolutionDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuResolutionDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuResolutionDefine_id_seq" OWNED BY public."IsuResolutionDefine".id;


--
-- Name: IsuResolution_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuResolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuResolution_id_seq" OWNER TO ngtesting;

--
-- Name: IsuResolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuResolution_id_seq" OWNED BY public."IsuResolution".id;


--
-- Name: IsuSeverity; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuSeverity" OWNER TO ngtesting;

--
-- Name: IsuSeverityDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuSeverityDefine" OWNER TO ngtesting;

--
-- Name: IsuSeverityDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuSeverityDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuSeverityDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuSeverityDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuSeverityDefine_id_seq" OWNED BY public."IsuSeverityDefine".id;


--
-- Name: IsuSeveritySolution; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuSeveritySolution" OWNER TO ngtesting;

--
-- Name: IsuSeveritySolutionItem; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuSeveritySolutionItem" (
    "severityId" integer,
    "solutionId" integer
);


ALTER TABLE public."IsuSeveritySolutionItem" OWNER TO ngtesting;

--
-- Name: IsuSeveritySolution_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuSeveritySolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuSeveritySolution_id_seq" OWNER TO ngtesting;

--
-- Name: IsuSeveritySolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuSeveritySolution_id_seq" OWNED BY public."IsuSeveritySolution".id;


--
-- Name: IsuSeverity_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuSeverity_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuSeverity_id_seq" OWNER TO ngtesting;

--
-- Name: IsuSeverity_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuSeverity_id_seq" OWNED BY public."IsuSeverity".id;


--
-- Name: IsuStatus; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuStatus" OWNER TO ngtesting;

--
-- Name: IsuStatusCategoryDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuStatusCategoryDefine" OWNER TO ngtesting;

--
-- Name: IsuStatusCategoryDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuStatusCategoryDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuStatusCategoryDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuStatusCategoryDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuStatusCategoryDefine_id_seq" OWNED BY public."IsuStatusCategoryDefine".id;


--
-- Name: IsuStatusDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuStatusDefine" OWNER TO ngtesting;

--
-- Name: IsuStatusDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuStatusDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuStatusDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuStatusDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuStatusDefine_id_seq" OWNED BY public."IsuStatusDefine".id;


--
-- Name: IsuStatus_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuStatus_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuStatus_id_seq" OWNER TO ngtesting;

--
-- Name: IsuStatus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuStatus_id_seq" OWNED BY public."IsuStatus".id;


--
-- Name: IsuTag; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuTag" OWNER TO ngtesting;

--
-- Name: IsuTagRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuTagRelation" (
    id integer NOT NULL,
    "issueId" integer,
    "tagId" integer
);


ALTER TABLE public."IsuTagRelation" OWNER TO ngtesting;

--
-- Name: IsuTagRelation_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuTagRelation_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuTagRelation_id_seq" OWNER TO ngtesting;

--
-- Name: IsuTagRelation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuTagRelation_id_seq" OWNED BY public."IsuTagRelation".id;


--
-- Name: IsuTag_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuTag_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuTag_id_seq" OWNER TO ngtesting;

--
-- Name: IsuTag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuTag_id_seq" OWNED BY public."IsuTag".id;


--
-- Name: IsuType; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuType" OWNER TO ngtesting;

--
-- Name: IsuTypeDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuTypeDefine" OWNER TO ngtesting;

--
-- Name: IsuTypeDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuTypeDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuTypeDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuTypeDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuTypeDefine_id_seq" OWNED BY public."IsuTypeDefine".id;


--
-- Name: IsuTypeSolution; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuTypeSolution" OWNER TO ngtesting;

--
-- Name: IsuTypeSolutionItem; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuTypeSolutionItem" (
    "typeId" integer,
    "solutionId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuTypeSolutionItem" OWNER TO ngtesting;

--
-- Name: IsuTypeSolution_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuTypeSolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuTypeSolution_id_seq" OWNER TO ngtesting;

--
-- Name: IsuTypeSolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuTypeSolution_id_seq" OWNED BY public."IsuTypeSolution".id;


--
-- Name: IsuType_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuType_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuType_id_seq" OWNER TO ngtesting;

--
-- Name: IsuType_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuType_id_seq" OWNED BY public."IsuType".id;


--
-- Name: IsuWatch; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuWatch" (
    id integer NOT NULL,
    "userId" integer,
    "issueId" integer
);


ALTER TABLE public."IsuWatch" OWNER TO ngtesting;

--
-- Name: IsuWatch_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuWatch_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWatch_id_seq" OWNER TO ngtesting;

--
-- Name: IsuWatch_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuWatch_id_seq" OWNED BY public."IsuWatch".id;


--
-- Name: IsuWorkflow; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuWorkflow" OWNER TO ngtesting;

--
-- Name: IsuWorkflowSolution; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuWorkflowSolution" OWNER TO ngtesting;

--
-- Name: IsuWorkflowSolutionItem; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuWorkflowSolutionItem" (
    id integer NOT NULL,
    "typeId" integer,
    "workflowId" integer,
    "solutionId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuWorkflowSolutionItem" OWNER TO ngtesting;

--
-- Name: IsuWorkflowSolutionItem_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuWorkflowSolutionItem_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowSolutionItem_id_seq" OWNER TO ngtesting;

--
-- Name: IsuWorkflowSolutionItem_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuWorkflowSolutionItem_id_seq" OWNED BY public."IsuWorkflowSolutionItem".id;


--
-- Name: IsuWorkflowSolution_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuWorkflowSolution_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowSolution_id_seq" OWNER TO ngtesting;

--
-- Name: IsuWorkflowSolution_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuWorkflowSolution_id_seq" OWNED BY public."IsuWorkflowSolution".id;


--
-- Name: IsuWorkflowStatusRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuWorkflowStatusRelation" (
    id integer NOT NULL,
    "workflowId" integer,
    "statusId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuWorkflowStatusRelation" OWNER TO ngtesting;

--
-- Name: IsuWorkflowStatusRelationDefine; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuWorkflowStatusRelationDefine" (
    id integer NOT NULL,
    "workflowId" integer,
    "statusId" integer
);


ALTER TABLE public."IsuWorkflowStatusRelationDefine" OWNER TO ngtesting;

--
-- Name: IsuWorkflowStatusRelationDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuWorkflowStatusRelationDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowStatusRelationDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuWorkflowStatusRelationDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuWorkflowStatusRelationDefine_id_seq" OWNED BY public."IsuWorkflowStatusRelationDefine".id;


--
-- Name: IsuWorkflowStatusRelation_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuWorkflowStatusRelation_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowStatusRelation_id_seq" OWNER TO ngtesting;

--
-- Name: IsuWorkflowStatusRelation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuWorkflowStatusRelation_id_seq" OWNED BY public."IsuWorkflowStatusRelation".id;


--
-- Name: IsuWorkflowTransition; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuWorkflowTransition" OWNER TO ngtesting;

--
-- Name: IsuWorkflowTransitionDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."IsuWorkflowTransitionDefine" OWNER TO ngtesting;

--
-- Name: IsuWorkflowTransitionDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuWorkflowTransitionDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowTransitionDefine_id_seq" OWNER TO ngtesting;

--
-- Name: IsuWorkflowTransitionDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuWorkflowTransitionDefine_id_seq" OWNED BY public."IsuWorkflowTransitionDefine".id;


--
-- Name: IsuWorkflowTransitionProjectRoleRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."IsuWorkflowTransitionProjectRoleRelation" (
    id integer NOT NULL,
    "workflowId" integer,
    "workflowTransitionId" integer,
    "projectRoleId" integer,
    "orgId" integer
);


ALTER TABLE public."IsuWorkflowTransitionProjectRoleRelation" OWNER TO ngtesting;

--
-- Name: IsuWorkflowTransitionProjectRoleRelation_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuWorkflowTransitionProjectRoleRelation_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowTransitionProjectRoleRelation_id_seq" OWNER TO ngtesting;

--
-- Name: IsuWorkflowTransitionProjectRoleRelation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuWorkflowTransitionProjectRoleRelation_id_seq" OWNED BY public."IsuWorkflowTransitionProjectRoleRelation".id;


--
-- Name: IsuWorkflowTransition_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuWorkflowTransition_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflowTransition_id_seq" OWNER TO ngtesting;

--
-- Name: IsuWorkflowTransition_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuWorkflowTransition_id_seq" OWNED BY public."IsuWorkflowTransition".id;


--
-- Name: IsuWorkflow_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."IsuWorkflow_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."IsuWorkflow_id_seq" OWNER TO ngtesting;

--
-- Name: IsuWorkflow_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."IsuWorkflow_id_seq" OWNED BY public."IsuWorkflow".id;


--
-- Name: SysPrivilege; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."SysPrivilege" OWNER TO ngtesting;

--
-- Name: SysPrivilege_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."SysPrivilege_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SysPrivilege_id_seq" OWNER TO ngtesting;

--
-- Name: SysPrivilege_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."SysPrivilege_id_seq" OWNED BY public."SysPrivilege".id;


--
-- Name: SysRole; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."SysRole" OWNER TO ngtesting;

--
-- Name: SysRolePrivilegeRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."SysRolePrivilegeRelation" (
    "roleId" integer NOT NULL,
    "privilegeId" integer NOT NULL
);


ALTER TABLE public."SysRolePrivilegeRelation" OWNER TO ngtesting;

--
-- Name: SysRoleUserRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."SysRoleUserRelation" (
    "roleId" integer NOT NULL,
    "userId" integer NOT NULL
);


ALTER TABLE public."SysRoleUserRelation" OWNER TO ngtesting;

--
-- Name: SysRole_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."SysRole_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SysRole_id_seq" OWNER TO ngtesting;

--
-- Name: SysRole_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."SysRole_id_seq" OWNED BY public."SysRole".id;


--
-- Name: SysUser; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."SysUser" OWNER TO ngtesting;

--
-- Name: SysUser_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."SysUser_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SysUser_id_seq" OWNER TO ngtesting;

--
-- Name: SysUser_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."SysUser_id_seq" OWNED BY public."SysUser".id;


--
-- Name: Test; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."Test" (
    id bigint NOT NULL,
    name character varying(255),
    "extProp" jsonb
);


ALTER TABLE public."Test" OWNER TO ngtesting;

--
-- Name: Test_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."Test_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Test_id_seq" OWNER TO ngtesting;

--
-- Name: Test_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."Test_id_seq" OWNED BY public."Test".id;


--
-- Name: TstAlert; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstAlert" OWNER TO ngtesting;

--
-- Name: TstAlert_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstAlert_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstAlert_id_seq" OWNER TO ngtesting;

--
-- Name: TstAlert_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstAlert_id_seq" OWNED BY public."TstAlert".id;


--
-- Name: TstCase; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCase" OWNER TO ngtesting;

--
-- Name: TstCaseAttachment; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseAttachment" OWNER TO ngtesting;

--
-- Name: TstCaseAttachment_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseAttachment_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseAttachment_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseAttachment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseAttachment_id_seq" OWNED BY public."TstCaseAttachment".id;


--
-- Name: TstCaseComments; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseComments" OWNER TO ngtesting;

--
-- Name: TstCaseComments_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseComments_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseComments_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseComments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseComments_id_seq" OWNED BY public."TstCaseComments".id;


--
-- Name: TstCaseExeStatusDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseExeStatusDefine" OWNER TO ngtesting;

--
-- Name: TstCaseExeStatus_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseExeStatus_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseExeStatus_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseExeStatus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseExeStatus_id_seq" OWNED BY public."TstCaseExeStatusDefine".id;


--
-- Name: TstCaseExeStatus; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseExeStatus" OWNER TO ngtesting;

--
-- Name: TstCaseHistory; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseHistory" OWNER TO ngtesting;

--
-- Name: TstCaseHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseHistory_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseHistory_id_seq" OWNED BY public."TstCaseHistory".id;


--
-- Name: TstCaseInSuite; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseInSuite" OWNER TO ngtesting;

--
-- Name: TstCaseInSuite_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseInSuite_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInSuite_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseInSuite_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseInSuite_id_seq" OWNED BY public."TstCaseInSuite".id;


--
-- Name: TstCaseInTask; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseInTask" OWNER TO ngtesting;

--
-- Name: TstCaseInTaskAttachment; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseInTaskAttachment" OWNER TO ngtesting;

--
-- Name: TstCaseInTaskAttachment_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseInTaskAttachment_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTaskAttachment_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseInTaskAttachment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseInTaskAttachment_id_seq" OWNED BY public."TstCaseInTaskAttachment".id;


--
-- Name: TstCaseInTaskComments; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseInTaskComments" OWNER TO ngtesting;

--
-- Name: TstCaseInTaskComments_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseInTaskComments_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTaskComments_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseInTaskComments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseInTaskComments_id_seq" OWNED BY public."TstCaseInTaskComments".id;


--
-- Name: TstCaseInTaskHistory; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseInTaskHistory" OWNER TO ngtesting;

--
-- Name: TstCaseInTaskHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseInTaskHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTaskHistory_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseInTaskHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseInTaskHistory_id_seq" OWNED BY public."TstCaseInTaskHistory".id;


--
-- Name: TstCaseInTaskIssue; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseInTaskIssue" OWNER TO ngtesting;

--
-- Name: TstCaseInTaskIssue_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseInTaskIssue_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTaskIssue_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseInTaskIssue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseInTaskIssue_id_seq" OWNED BY public."TstCaseInTaskIssue".id;


--
-- Name: TstCaseInTask_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseInTask_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseInTask_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseInTask_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseInTask_id_seq" OWNED BY public."TstCaseInTask".id;


--
-- Name: TstCasePriorityDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCasePriorityDefine" OWNER TO ngtesting;

--
-- Name: TstCasePriority_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCasePriority_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCasePriority_id_seq" OWNER TO ngtesting;

--
-- Name: TstCasePriority_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCasePriority_id_seq" OWNED BY public."TstCasePriorityDefine".id;


--
-- Name: TstCasePriority; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCasePriority" OWNER TO ngtesting;

--
-- Name: TstCaseStep; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseStep" OWNER TO ngtesting;

--
-- Name: TstCaseStep_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseStep_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseStep_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseStep_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseStep_id_seq" OWNED BY public."TstCaseStep".id;


--
-- Name: TstCaseTypeDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseTypeDefine" OWNER TO ngtesting;

--
-- Name: TstCaseType_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCaseType_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCaseType_id_seq" OWNER TO ngtesting;

--
-- Name: TstCaseType_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCaseType_id_seq" OWNED BY public."TstCaseTypeDefine".id;


--
-- Name: TstCaseType; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstCaseType" OWNER TO ngtesting;

--
-- Name: TstCase_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstCase_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstCase_id_seq" OWNER TO ngtesting;

--
-- Name: TstCase_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstCase_id_seq" OWNED BY public."TstCase".id;


--
-- Name: TstDocument; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstDocument" OWNER TO ngtesting;

--
-- Name: TstDocument_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstDocument_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstDocument_id_seq" OWNER TO ngtesting;

--
-- Name: TstDocument_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstDocument_id_seq" OWNED BY public."TstDocument".id;


--
-- Name: TstEmail; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstEmail" OWNER TO ngtesting;

--
-- Name: TstEmail_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstEmail_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstEmail_id_seq" OWNER TO ngtesting;

--
-- Name: TstEmail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstEmail_id_seq" OWNED BY public."TstEmail".id;


--
-- Name: TstEnv; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstEnv" OWNER TO ngtesting;

--
-- Name: TstEnv_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstEnv_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstEnv_id_seq" OWNER TO ngtesting;

--
-- Name: TstEnv_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstEnv_id_seq" OWNED BY public."TstEnv".id;


--
-- Name: TstHistory; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstHistory" OWNER TO ngtesting;

--
-- Name: TstHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstHistory_id_seq" OWNER TO ngtesting;

--
-- Name: TstHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstHistory_id_seq" OWNED BY public."TstHistory".id;


--
-- Name: TstModule; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstModule" OWNER TO ngtesting;

--
-- Name: TstModule_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstModule_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstModule_id_seq" OWNER TO ngtesting;

--
-- Name: TstModule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstModule_id_seq" OWNED BY public."TstModule".id;


--
-- Name: TstMsg; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstMsg" OWNER TO ngtesting;

--
-- Name: TstMsg_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstMsg_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstMsg_id_seq" OWNER TO ngtesting;

--
-- Name: TstMsg_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstMsg_id_seq" OWNED BY public."TstMsg".id;


--
-- Name: TstOrg; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstOrg" OWNER TO ngtesting;

--
-- Name: TstOrgGroup; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstOrgGroup" OWNER TO ngtesting;

--
-- Name: TstOrgGroupUserRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."TstOrgGroupUserRelation" (
    "orgId" integer,
    "orgGroupId" integer,
    "userId" integer
);


ALTER TABLE public."TstOrgGroupUserRelation" OWNER TO ngtesting;

--
-- Name: TstOrgGroup_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstOrgGroup_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstOrgGroup_id_seq" OWNER TO ngtesting;

--
-- Name: TstOrgGroup_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstOrgGroup_id_seq" OWNED BY public."TstOrgGroup".id;


--
-- Name: TstOrgPrivilegeDefine; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."TstOrgPrivilegeDefine" (
    id integer NOT NULL,
    code character varying(255),
    name character varying(255),
    descr character varying(255),
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstOrgPrivilegeDefine" OWNER TO ngtesting;

--
-- Name: TstOrgPrivilegeDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstOrgPrivilegeDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstOrgPrivilegeDefine_id_seq" OWNER TO ngtesting;

--
-- Name: TstOrgPrivilegeDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstOrgPrivilegeDefine_id_seq" OWNED BY public."TstOrgPrivilegeDefine".id;


--
-- Name: TstOrgRole; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstOrgRole" OWNER TO ngtesting;

--
-- Name: TstOrgRoleGroupRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."TstOrgRoleGroupRelation" (
    "orgRoleId" integer NOT NULL,
    "orgGroupId" integer NOT NULL,
    "orgId" integer NOT NULL
);


ALTER TABLE public."TstOrgRoleGroupRelation" OWNER TO ngtesting;

--
-- Name: TstOrgRolePrivilegeRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."TstOrgRolePrivilegeRelation" (
    "orgId" integer NOT NULL,
    "orgRoleId" integer NOT NULL,
    "orgPrivilegeId" integer NOT NULL
);


ALTER TABLE public."TstOrgRolePrivilegeRelation" OWNER TO ngtesting;

--
-- Name: TstOrgRoleUserRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."TstOrgRoleUserRelation" (
    "orgRoleId" integer NOT NULL,
    "userId" integer NOT NULL,
    "orgId" integer NOT NULL
);


ALTER TABLE public."TstOrgRoleUserRelation" OWNER TO ngtesting;

--
-- Name: TstOrgRole_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstOrgRole_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstOrgRole_id_seq" OWNER TO ngtesting;

--
-- Name: TstOrgRole_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstOrgRole_id_seq" OWNED BY public."TstOrgRole".id;


--
-- Name: TstOrgUserRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."TstOrgUserRelation" (
    "orgId" integer NOT NULL,
    "userId" integer NOT NULL
);


ALTER TABLE public."TstOrgUserRelation" OWNER TO ngtesting;

--
-- Name: TstOrg_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstOrg_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstOrg_id_seq" OWNER TO ngtesting;

--
-- Name: TstOrg_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstOrg_id_seq" OWNED BY public."TstOrg".id;


--
-- Name: TstPlan; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstPlan" OWNER TO ngtesting;

--
-- Name: TstPlan_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstPlan_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstPlan_id_seq" OWNER TO ngtesting;

--
-- Name: TstPlan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstPlan_id_seq" OWNED BY public."TstPlan".id;


--
-- Name: TstProject; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstProject" OWNER TO ngtesting;

--
-- Name: TstProjectAccessHistory; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstProjectAccessHistory" OWNER TO ngtesting;

--
-- Name: TstProjectAccessHistory_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstProjectAccessHistory_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstProjectAccessHistory_id_seq" OWNER TO ngtesting;

--
-- Name: TstProjectAccessHistory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstProjectAccessHistory_id_seq" OWNED BY public."TstProjectAccessHistory".id;


--
-- Name: TstProjectPrivilegeDefine; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstProjectPrivilegeDefine" OWNER TO ngtesting;

--
-- Name: TstProjectPrivilegeDefine_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstProjectPrivilegeDefine_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstProjectPrivilegeDefine_id_seq" OWNER TO ngtesting;

--
-- Name: TstProjectPrivilegeDefine_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstProjectPrivilegeDefine_id_seq" OWNED BY public."TstProjectPrivilegeDefine".id;


--
-- Name: TstProjectRole; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstProjectRole" OWNER TO ngtesting;

--
-- Name: TstProjectRoleEntityRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."TstProjectRoleEntityRelation" (
    "entityId" integer,
    "orgId" integer,
    "projectId" integer,
    "projectRoleId" integer,
    type character varying(255)
);


ALTER TABLE public."TstProjectRoleEntityRelation" OWNER TO ngtesting;

--
-- Name: TstProjectRolePriviledgeRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."TstProjectRolePriviledgeRelation" (
    "projectPrivilegeDefineId" integer,
    "projectRoleId" integer,
    "orgId" integer
);


ALTER TABLE public."TstProjectRolePriviledgeRelation" OWNER TO ngtesting;

--
-- Name: TstProjectRole_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstProjectRole_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstProjectRole_id_seq" OWNER TO ngtesting;

--
-- Name: TstProjectRole_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstProjectRole_id_seq" OWNED BY public."TstProjectRole".id;


--
-- Name: TstProject_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstProject_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstProject_id_seq" OWNER TO ngtesting;

--
-- Name: TstProject_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstProject_id_seq" OWNED BY public."TstProject".id;


--
-- Name: TstSuite; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstSuite" OWNER TO ngtesting;

--
-- Name: TstSuite_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstSuite_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstSuite_id_seq" OWNER TO ngtesting;

--
-- Name: TstSuite_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstSuite_id_seq" OWNED BY public."TstSuite".id;


--
-- Name: TstTask; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstTask" OWNER TO ngtesting;

--
-- Name: TstTaskAssigneeRelation; Type: TABLE; Schema: public; Owner: ngtesting
--

CREATE TABLE public."TstTaskAssigneeRelation" (
    "taskId" integer NOT NULL,
    "assigneeId" integer NOT NULL
);


ALTER TABLE public."TstTaskAssigneeRelation" OWNER TO ngtesting;

--
-- Name: TstTask_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstTask_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstTask_id_seq" OWNER TO ngtesting;

--
-- Name: TstTask_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstTask_id_seq" OWNED BY public."TstTask".id;


--
-- Name: TstThread; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstThread" OWNER TO ngtesting;

--
-- Name: TstThread_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstThread_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstThread_id_seq" OWNER TO ngtesting;

--
-- Name: TstThread_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstThread_id_seq" OWNED BY public."TstThread".id;


--
-- Name: TstUser; Type: TABLE; Schema: public; Owner: ngtesting
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
    token character varying(255),
    "verifyCode" character varying(255),
    "lastLoginTime" timestamp without time zone,
    disabled boolean,
    deleted boolean,
    "createTime" timestamp without time zone,
    "updateTime" timestamp without time zone
);


ALTER TABLE public."TstUser" OWNER TO ngtesting;

--
-- Name: TstUserSettings; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstUserSettings" OWNER TO ngtesting;

--
-- Name: TstUserVerifyCode; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstUserVerifyCode" OWNER TO ngtesting;

--
-- Name: TstUserVerifyCode_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstUserVerifyCode_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstUserVerifyCode_id_seq" OWNER TO ngtesting;

--
-- Name: TstUserVerifyCode_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstUserVerifyCode_id_seq" OWNED BY public."TstUserVerifyCode".id;


--
-- Name: TstUser_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstUser_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstUser_id_seq" OWNER TO ngtesting;

--
-- Name: TstUser_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstUser_id_seq" OWNED BY public."TstUser".id;


--
-- Name: TstVer; Type: TABLE; Schema: public; Owner: ngtesting
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


ALTER TABLE public."TstVer" OWNER TO ngtesting;

--
-- Name: TstVer_id_seq; Type: SEQUENCE; Schema: public; Owner: ngtesting
--

CREATE SEQUENCE public."TstVer_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."TstVer_id_seq" OWNER TO ngtesting;

--
-- Name: TstVer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ngtesting
--

ALTER SEQUENCE public."TstVer_id_seq" OWNED BY public."TstVer".id;


--
-- Name: CustomField id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomField" ALTER COLUMN id SET DEFAULT nextval('public."CustomField_id_seq"'::regclass);


--
-- Name: CustomFieldDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldDefine_id_seq"'::regclass);


--
-- Name: CustomFieldInputTypeRelationDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldInputTypeRelationDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldInputTypeRelationDefine_id_seq"'::regclass);


--
-- Name: CustomFieldIputDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldIputDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldIputDefine_id_seq"'::regclass);


--
-- Name: CustomFieldOption id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldOption" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldOption_id_seq"'::regclass);


--
-- Name: CustomFieldOptionDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldOptionDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldOptionDefine_id_seq"'::regclass);


--
-- Name: CustomFieldTypeDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldTypeDefine" ALTER COLUMN id SET DEFAULT nextval('public."CustomFieldTypeDefine_id_seq"'::regclass);


--
-- Name: IsuAttachment id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuAttachment" ALTER COLUMN id SET DEFAULT nextval('public."IsuAttachment_id_seq"'::regclass);


--
-- Name: IsuComments id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuComments" ALTER COLUMN id SET DEFAULT nextval('public."IsuComments_id_seq"'::regclass);


--
-- Name: IsuCustomFieldSolution id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuCustomFieldSolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuCustomFieldSolution_id_seq"'::regclass);


--
-- Name: IsuDocument id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuDocument" ALTER COLUMN id SET DEFAULT nextval('public."IsuDocument_id_seq"'::regclass);


--
-- Name: IsuField id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuField" ALTER COLUMN id SET DEFAULT nextval('public."IsuField_id_seq"'::regclass);


--
-- Name: IsuFieldCodeToTableDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuFieldCodeToTableDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuFieldCodeToTableDefine_id_seq"'::regclass);


--
-- Name: IsuFieldDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuFieldDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuFieldDefine_id_seq"'::regclass);


--
-- Name: IsuHistory id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuHistory" ALTER COLUMN id SET DEFAULT nextval('public."IsuHistory_id_seq"'::regclass);


--
-- Name: IsuIssue id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue" ALTER COLUMN id SET DEFAULT nextval('public."IsuIssue_id_seq"'::regclass);


--
-- Name: IsuIssueExt pid; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssueExt" ALTER COLUMN pid SET DEFAULT nextval('public."IsuIssueExt_pid_seq"'::regclass);


--
-- Name: IsuLink id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuLink" ALTER COLUMN id SET DEFAULT nextval('public."IsuLink_id_seq"'::regclass);


--
-- Name: IsuLinkReasonDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuLinkReasonDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuLinkReasonDefine_id_seq"'::regclass);


--
-- Name: IsuNotification id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuNotification" ALTER COLUMN id SET DEFAULT nextval('public."IsuNotification_id_seq"'::regclass);


--
-- Name: IsuNotificationDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuNotificationDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuNotificationDefine_id_seq"'::regclass);


--
-- Name: IsuPage id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPage" ALTER COLUMN id SET DEFAULT nextval('public."IsuPage_id_seq"'::regclass);


--
-- Name: IsuPageElement id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageElement" ALTER COLUMN id SET DEFAULT nextval('public."IsuPageElement_id_seq"'::regclass);


--
-- Name: IsuPageSolution id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageSolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuPageSolution_id_seq"'::regclass);


--
-- Name: IsuPageSolutionItem id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageSolutionItem" ALTER COLUMN id SET DEFAULT nextval('public."IsuPageSolutionItem_id_seq"'::regclass);


--
-- Name: IsuPriority id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPriority" ALTER COLUMN id SET DEFAULT nextval('public."IsuPriority_id_seq"'::regclass);


--
-- Name: IsuPriorityDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPriorityDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuPriorityDefine_id_seq"'::regclass);


--
-- Name: IsuPrioritySolution id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPrioritySolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuPrioritySolution_id_seq"'::regclass);


--
-- Name: IsuQuery id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuQuery" ALTER COLUMN id SET DEFAULT nextval('public."IsuQuery_id_seq"'::regclass);


--
-- Name: IsuResolution id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuResolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuResolution_id_seq"'::regclass);


--
-- Name: IsuResolutionDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuResolutionDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuResolutionDefine_id_seq"'::regclass);


--
-- Name: IsuSeverity id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeverity" ALTER COLUMN id SET DEFAULT nextval('public."IsuSeverity_id_seq"'::regclass);


--
-- Name: IsuSeverityDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeverityDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuSeverityDefine_id_seq"'::regclass);


--
-- Name: IsuSeveritySolution id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeveritySolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuSeveritySolution_id_seq"'::regclass);


--
-- Name: IsuStatus id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuStatus" ALTER COLUMN id SET DEFAULT nextval('public."IsuStatus_id_seq"'::regclass);


--
-- Name: IsuStatusCategoryDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuStatusCategoryDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuStatusCategoryDefine_id_seq"'::regclass);


--
-- Name: IsuStatusDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuStatusDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuStatusDefine_id_seq"'::regclass);


--
-- Name: IsuTag id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTag" ALTER COLUMN id SET DEFAULT nextval('public."IsuTag_id_seq"'::regclass);


--
-- Name: IsuTagRelation id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTagRelation" ALTER COLUMN id SET DEFAULT nextval('public."IsuTagRelation_id_seq"'::regclass);


--
-- Name: IsuType id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuType" ALTER COLUMN id SET DEFAULT nextval('public."IsuType_id_seq"'::regclass);


--
-- Name: IsuTypeDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTypeDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuTypeDefine_id_seq"'::regclass);


--
-- Name: IsuTypeSolution id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTypeSolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuTypeSolution_id_seq"'::regclass);


--
-- Name: IsuWatch id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWatch" ALTER COLUMN id SET DEFAULT nextval('public."IsuWatch_id_seq"'::regclass);


--
-- Name: IsuWorkflow id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflow" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflow_id_seq"'::regclass);


--
-- Name: IsuWorkflowSolution id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowSolution" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowSolution_id_seq"'::regclass);


--
-- Name: IsuWorkflowSolutionItem id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowSolutionItem_id_seq"'::regclass);


--
-- Name: IsuWorkflowStatusRelation id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowStatusRelation_id_seq"'::regclass);


--
-- Name: IsuWorkflowStatusRelationDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelationDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowStatusRelationDefine_id_seq"'::regclass);


--
-- Name: IsuWorkflowTransition id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransition" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowTransition_id_seq"'::regclass);


--
-- Name: IsuWorkflowTransitionDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionDefine" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowTransitionDefine_id_seq"'::regclass);


--
-- Name: IsuWorkflowTransitionProjectRoleRelation id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation" ALTER COLUMN id SET DEFAULT nextval('public."IsuWorkflowTransitionProjectRoleRelation_id_seq"'::regclass);


--
-- Name: SysPrivilege id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysPrivilege" ALTER COLUMN id SET DEFAULT nextval('public."SysPrivilege_id_seq"'::regclass);


--
-- Name: SysRole id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysRole" ALTER COLUMN id SET DEFAULT nextval('public."SysRole_id_seq"'::regclass);


--
-- Name: SysUser id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysUser" ALTER COLUMN id SET DEFAULT nextval('public."SysUser_id_seq"'::regclass);


--
-- Name: Test id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."Test" ALTER COLUMN id SET DEFAULT nextval('public."Test_id_seq"'::regclass);


--
-- Name: TstAlert id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstAlert" ALTER COLUMN id SET DEFAULT nextval('public."TstAlert_id_seq"'::regclass);


--
-- Name: TstCase id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCase" ALTER COLUMN id SET DEFAULT nextval('public."TstCase_id_seq"'::regclass);


--
-- Name: TstCaseAttachment id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseAttachment" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseAttachment_id_seq"'::regclass);


--
-- Name: TstCaseComments id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseComments" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseComments_id_seq"'::regclass);


--
-- Name: TstCaseExeStatusDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseExeStatusDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseExeStatus_id_seq"'::regclass);


--
-- Name: TstCaseHistory id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseHistory" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseHistory_id_seq"'::regclass);


--
-- Name: TstCaseInSuite id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInSuite" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInSuite_id_seq"'::regclass);


--
-- Name: TstCaseInTask id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTask" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTask_id_seq"'::regclass);


--
-- Name: TstCaseInTaskAttachment id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskAttachment" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTaskAttachment_id_seq"'::regclass);


--
-- Name: TstCaseInTaskComments id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskComments" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTaskComments_id_seq"'::regclass);


--
-- Name: TstCaseInTaskHistory id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskHistory" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTaskHistory_id_seq"'::regclass);


--
-- Name: TstCaseInTaskIssue id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskIssue" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseInTaskIssue_id_seq"'::regclass);


--
-- Name: TstCasePriorityDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCasePriorityDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstCasePriority_id_seq"'::regclass);


--
-- Name: TstCaseStep id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseStep" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseStep_id_seq"'::regclass);


--
-- Name: TstCaseTypeDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseTypeDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstCaseType_id_seq"'::regclass);


--
-- Name: TstDocument id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstDocument" ALTER COLUMN id SET DEFAULT nextval('public."TstDocument_id_seq"'::regclass);


--
-- Name: TstEmail id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstEmail" ALTER COLUMN id SET DEFAULT nextval('public."TstEmail_id_seq"'::regclass);


--
-- Name: TstEnv id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstEnv" ALTER COLUMN id SET DEFAULT nextval('public."TstEnv_id_seq"'::regclass);


--
-- Name: TstHistory id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstHistory" ALTER COLUMN id SET DEFAULT nextval('public."TstHistory_id_seq"'::regclass);


--
-- Name: TstModule id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstModule" ALTER COLUMN id SET DEFAULT nextval('public."TstModule_id_seq"'::regclass);


--
-- Name: TstMsg id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstMsg" ALTER COLUMN id SET DEFAULT nextval('public."TstMsg_id_seq"'::regclass);


--
-- Name: TstOrg id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrg" ALTER COLUMN id SET DEFAULT nextval('public."TstOrg_id_seq"'::regclass);


--
-- Name: TstOrgGroup id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgGroup" ALTER COLUMN id SET DEFAULT nextval('public."TstOrgGroup_id_seq"'::regclass);


--
-- Name: TstOrgPrivilegeDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgPrivilegeDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstOrgPrivilegeDefine_id_seq"'::regclass);


--
-- Name: TstOrgRole id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRole" ALTER COLUMN id SET DEFAULT nextval('public."TstOrgRole_id_seq"'::regclass);


--
-- Name: TstPlan id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstPlan" ALTER COLUMN id SET DEFAULT nextval('public."TstPlan_id_seq"'::regclass);


--
-- Name: TstProject id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProject" ALTER COLUMN id SET DEFAULT nextval('public."TstProject_id_seq"'::regclass);


--
-- Name: TstProjectAccessHistory id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectAccessHistory" ALTER COLUMN id SET DEFAULT nextval('public."TstProjectAccessHistory_id_seq"'::regclass);


--
-- Name: TstProjectPrivilegeDefine id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectPrivilegeDefine" ALTER COLUMN id SET DEFAULT nextval('public."TstProjectPrivilegeDefine_id_seq"'::regclass);


--
-- Name: TstProjectRole id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectRole" ALTER COLUMN id SET DEFAULT nextval('public."TstProjectRole_id_seq"'::regclass);


--
-- Name: TstSuite id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstSuite" ALTER COLUMN id SET DEFAULT nextval('public."TstSuite_id_seq"'::regclass);


--
-- Name: TstTask id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstTask" ALTER COLUMN id SET DEFAULT nextval('public."TstTask_id_seq"'::regclass);


--
-- Name: TstThread id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstThread" ALTER COLUMN id SET DEFAULT nextval('public."TstThread_id_seq"'::regclass);


--
-- Name: TstUser id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstUser" ALTER COLUMN id SET DEFAULT nextval('public."TstUser_id_seq"'::regclass);


--
-- Name: TstUserVerifyCode id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstUserVerifyCode" ALTER COLUMN id SET DEFAULT nextval('public."TstUserVerifyCode_id_seq"'::regclass);


--
-- Name: TstVer id; Type: DEFAULT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstVer" ALTER COLUMN id SET DEFAULT nextval('public."TstVer_id_seq"'::regclass);


--
-- Name: CustomFieldDefine CustomFieldDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldDefine"
    ADD CONSTRAINT "CustomFieldDefine_pkey" PRIMARY KEY (id);


--
-- Name: CustomFieldInputTypeRelationDefine CustomFieldInputTypeRelationDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldInputTypeRelationDefine"
    ADD CONSTRAINT "CustomFieldInputTypeRelationDefine_pkey" PRIMARY KEY (id);


--
-- Name: CustomFieldIputDefine CustomFieldIputDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldIputDefine"
    ADD CONSTRAINT "CustomFieldIputDefine_pkey" PRIMARY KEY (id);


--
-- Name: CustomFieldOptionDefine CustomFieldOptionDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldOptionDefine"
    ADD CONSTRAINT "CustomFieldOptionDefine_pkey" PRIMARY KEY (id);


--
-- Name: CustomFieldOption CustomFieldOption_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldOption"
    ADD CONSTRAINT "CustomFieldOption_pkey" PRIMARY KEY (id);


--
-- Name: CustomFieldTypeDefine CustomFieldTypeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldTypeDefine"
    ADD CONSTRAINT "CustomFieldTypeDefine_pkey" PRIMARY KEY (id);


--
-- Name: CustomField CustomField_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomField"
    ADD CONSTRAINT "CustomField_pkey" PRIMARY KEY (id);


--
-- Name: IsuAttachment IsuAttachment_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuAttachment"
    ADD CONSTRAINT "IsuAttachment_pkey" PRIMARY KEY (id);


--
-- Name: IsuComments IsuComments_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuComments"
    ADD CONSTRAINT "IsuComments_pkey" PRIMARY KEY (id);


--
-- Name: IsuCustomFieldSolution IsuCustomFieldSolution_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuCustomFieldSolution"
    ADD CONSTRAINT "IsuCustomFieldSolution_pkey" PRIMARY KEY (id);


--
-- Name: IsuDocument IsuDocument_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuDocument"
    ADD CONSTRAINT "IsuDocument_pkey" PRIMARY KEY (id);


--
-- Name: IsuFieldCodeToTableDefine IsuFieldCodeToTableDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuFieldCodeToTableDefine"
    ADD CONSTRAINT "IsuFieldCodeToTableDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuFieldDefine IsuFieldDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuFieldDefine"
    ADD CONSTRAINT "IsuFieldDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuField IsuField_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuField"
    ADD CONSTRAINT "IsuField_pkey" PRIMARY KEY (id);


--
-- Name: IsuHistory IsuHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuHistory"
    ADD CONSTRAINT "IsuHistory_pkey" PRIMARY KEY (id);


--
-- Name: IsuIssueExt IsuIssueExt_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssueExt"
    ADD CONSTRAINT "IsuIssueExt_pkey" PRIMARY KEY (pid);


--
-- Name: IsuIssue IsuIssue_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_pkey" PRIMARY KEY (id);


--
-- Name: IsuLinkReasonDefine IsuLinkReasonDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuLinkReasonDefine"
    ADD CONSTRAINT "IsuLinkReasonDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuLink IsuLink_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuLink"
    ADD CONSTRAINT "IsuLink_pkey" PRIMARY KEY (id);


--
-- Name: IsuNotificationDefine IsuNotificationDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuNotificationDefine"
    ADD CONSTRAINT "IsuNotificationDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuNotification IsuNotification_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuNotification"
    ADD CONSTRAINT "IsuNotification_pkey" PRIMARY KEY (id);


--
-- Name: IsuPageElement IsuPageElement_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageElement"
    ADD CONSTRAINT "IsuPageElement_pkey" PRIMARY KEY (id);


--
-- Name: IsuPageSolutionItem IsuPageSolutionItem_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_pkey" PRIMARY KEY (id);


--
-- Name: IsuPageSolution IsuPageSolution_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageSolution"
    ADD CONSTRAINT "IsuPageSolution_pkey" PRIMARY KEY (id);


--
-- Name: IsuPage IsuPage_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPage"
    ADD CONSTRAINT "IsuPage_pkey" PRIMARY KEY (id);


--
-- Name: IsuPriorityDefine IsuPriorityDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPriorityDefine"
    ADD CONSTRAINT "IsuPriorityDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuPrioritySolution IsuPrioritySolution_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPrioritySolution"
    ADD CONSTRAINT "IsuPrioritySolution_pkey" PRIMARY KEY (id);


--
-- Name: IsuPriority IsuPriority_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPriority"
    ADD CONSTRAINT "IsuPriority_pkey" PRIMARY KEY (id);


--
-- Name: IsuQuery IsuQuery_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuQuery"
    ADD CONSTRAINT "IsuQuery_pkey" PRIMARY KEY (id);


--
-- Name: IsuResolutionDefine IsuResolutionDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuResolutionDefine"
    ADD CONSTRAINT "IsuResolutionDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuResolution IsuResolution_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuResolution"
    ADD CONSTRAINT "IsuResolution_pkey" PRIMARY KEY (id);


--
-- Name: IsuSeverityDefine IsuSeverityDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeverityDefine"
    ADD CONSTRAINT "IsuSeverityDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuSeveritySolution IsuSeveritySolution_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeveritySolution"
    ADD CONSTRAINT "IsuSeveritySolution_pkey" PRIMARY KEY (id);


--
-- Name: IsuSeverity IsuSeverity_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeverity"
    ADD CONSTRAINT "IsuSeverity_pkey" PRIMARY KEY (id);


--
-- Name: IsuStatusCategoryDefine IsuStatusCategoryDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuStatusCategoryDefine"
    ADD CONSTRAINT "IsuStatusCategoryDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuStatusDefine IsuStatusDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuStatusDefine"
    ADD CONSTRAINT "IsuStatusDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuStatus IsuStatus_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuStatus"
    ADD CONSTRAINT "IsuStatus_pkey" PRIMARY KEY (id);


--
-- Name: IsuTagRelation IsuTagRelation_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTagRelation"
    ADD CONSTRAINT "IsuTagRelation_pkey" PRIMARY KEY (id);


--
-- Name: IsuTag IsuTag_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTag"
    ADD CONSTRAINT "IsuTag_pkey" PRIMARY KEY (id);


--
-- Name: IsuTypeDefine IsuTypeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTypeDefine"
    ADD CONSTRAINT "IsuTypeDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuTypeSolution IsuTypeSolution_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTypeSolution"
    ADD CONSTRAINT "IsuTypeSolution_pkey" PRIMARY KEY (id);


--
-- Name: IsuType IsuType_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuType"
    ADD CONSTRAINT "IsuType_pkey" PRIMARY KEY (id);


--
-- Name: IsuWatch IsuWatch_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWatch"
    ADD CONSTRAINT "IsuWatch_pkey" PRIMARY KEY (id);


--
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_pkey" PRIMARY KEY (id);


--
-- Name: IsuWorkflowSolution IsuWorkflowSolution_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowSolution"
    ADD CONSTRAINT "IsuWorkflowSolution_pkey" PRIMARY KEY (id);


--
-- Name: IsuWorkflowStatusRelationDefine IsuWorkflowStatusRelationDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelationDefine"
    ADD CONSTRAINT "IsuWorkflowStatusRelationDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuWorkflowStatusRelation IsuWorkflowStatusRelation_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation"
    ADD CONSTRAINT "IsuWorkflowStatusRelation_pkey" PRIMARY KEY (id);


--
-- Name: IsuWorkflowTransitionDefine IsuWorkflowTransitionDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionDefine"
    ADD CONSTRAINT "IsuWorkflowTransitionDefine_pkey" PRIMARY KEY (id);


--
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelation_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_pkey" PRIMARY KEY (id);


--
-- Name: IsuWorkflowTransition IsuWorkflowTransition_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransition"
    ADD CONSTRAINT "IsuWorkflowTransition_pkey" PRIMARY KEY (id);


--
-- Name: IsuWorkflow IsuWorkflow_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflow"
    ADD CONSTRAINT "IsuWorkflow_pkey" PRIMARY KEY (id);


--
-- Name: SysPrivilege SysPrivilege_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysPrivilege"
    ADD CONSTRAINT "SysPrivilege_pkey" PRIMARY KEY (id);


--
-- Name: SysRole SysRole_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysRole"
    ADD CONSTRAINT "SysRole_pkey" PRIMARY KEY (id);


--
-- Name: SysUser SysUser_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysUser"
    ADD CONSTRAINT "SysUser_pkey" PRIMARY KEY (id);


--
-- Name: Test Test_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."Test"
    ADD CONSTRAINT "Test_pkey" PRIMARY KEY (id);


--
-- Name: TstAlert TstAlert_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstAlert"
    ADD CONSTRAINT "TstAlert_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseAttachment TstCaseAttachment_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseAttachment"
    ADD CONSTRAINT "TstCaseAttachment_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseComments TstCaseComments_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseComments"
    ADD CONSTRAINT "TstCaseComments_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseExeStatusDefine TstCaseExeStatusDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseExeStatusDefine"
    ADD CONSTRAINT "TstCaseExeStatusDefine_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseExeStatus TstCaseExeStatus_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseExeStatus"
    ADD CONSTRAINT "TstCaseExeStatus_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseHistory TstCaseHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseHistory"
    ADD CONSTRAINT "TstCaseHistory_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseInSuite TstCaseInSuite_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseInTaskAttachment TstCaseInTaskAttachment_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskAttachment"
    ADD CONSTRAINT "TstCaseInTaskAttachment_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseInTaskComments TstCaseInTaskComments_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskComments"
    ADD CONSTRAINT "TstCaseInTaskComments_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseInTaskHistory TstCaseInTaskHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskHistory"
    ADD CONSTRAINT "TstCaseInTaskHistory_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseInTaskIssue TstCaseInTaskIssue_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskIssue"
    ADD CONSTRAINT "TstCaseInTaskIssue_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseInTask TstCaseInTask_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_pkey" PRIMARY KEY (id);


--
-- Name: TstCasePriorityDefine TstCasePriorityDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCasePriorityDefine"
    ADD CONSTRAINT "TstCasePriorityDefine_pkey" PRIMARY KEY (id);


--
-- Name: TstCasePriority TstCasePriority_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCasePriority"
    ADD CONSTRAINT "TstCasePriority_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseStep TstCaseStep_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseStep"
    ADD CONSTRAINT "TstCaseStep_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseTypeDefine TstCaseTypeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseTypeDefine"
    ADD CONSTRAINT "TstCaseTypeDefine_pkey" PRIMARY KEY (id);


--
-- Name: TstCaseType TstCaseType_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseType"
    ADD CONSTRAINT "TstCaseType_pkey" PRIMARY KEY (id);


--
-- Name: TstCase TstCase_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_pkey" PRIMARY KEY (id);


--
-- Name: TstDocument TstDocument_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstDocument"
    ADD CONSTRAINT "TstDocument_pkey" PRIMARY KEY (id);


--
-- Name: TstEmail TstEmail_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstEmail"
    ADD CONSTRAINT "TstEmail_pkey" PRIMARY KEY (id);


--
-- Name: TstEnv TstEnv_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstEnv"
    ADD CONSTRAINT "TstEnv_pkey" PRIMARY KEY (id);


--
-- Name: TstHistory TstHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstHistory"
    ADD CONSTRAINT "TstHistory_pkey" PRIMARY KEY (id);


--
-- Name: TstModule TstModule_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstModule"
    ADD CONSTRAINT "TstModule_pkey" PRIMARY KEY (id);


--
-- Name: TstMsg TstMsg_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstMsg"
    ADD CONSTRAINT "TstMsg_pkey" PRIMARY KEY (id);


--
-- Name: TstOrgGroup TstOrgGroup_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgGroup"
    ADD CONSTRAINT "TstOrgGroup_pkey" PRIMARY KEY (id);


--
-- Name: TstOrgPrivilegeDefine TstOrgPrivilegeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgPrivilegeDefine"
    ADD CONSTRAINT "TstOrgPrivilegeDefine_pkey" PRIMARY KEY (id);


--
-- Name: TstOrgRole TstOrgRole_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRole"
    ADD CONSTRAINT "TstOrgRole_pkey" PRIMARY KEY (id);


--
-- Name: TstOrg TstOrg_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrg"
    ADD CONSTRAINT "TstOrg_pkey" PRIMARY KEY (id);


--
-- Name: TstPlan TstPlan_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstPlan"
    ADD CONSTRAINT "TstPlan_pkey" PRIMARY KEY (id);


--
-- Name: TstProjectAccessHistory TstProjectAccessHistory_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectAccessHistory"
    ADD CONSTRAINT "TstProjectAccessHistory_pkey" PRIMARY KEY (id);


--
-- Name: TstProjectPrivilegeDefine TstProjectPrivilegeDefine_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectPrivilegeDefine"
    ADD CONSTRAINT "TstProjectPrivilegeDefine_pkey" PRIMARY KEY (id);


--
-- Name: TstProjectRole TstProjectRole_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectRole"
    ADD CONSTRAINT "TstProjectRole_pkey" PRIMARY KEY (id);


--
-- Name: TstProject TstProject_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_pkey" PRIMARY KEY (id);


--
-- Name: TstSuite TstSuite_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstSuite"
    ADD CONSTRAINT "TstSuite_pkey" PRIMARY KEY (id);


--
-- Name: TstTask TstTask_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_pkey" PRIMARY KEY (id);


--
-- Name: TstThread TstThread_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstThread"
    ADD CONSTRAINT "TstThread_pkey" PRIMARY KEY (id);


--
-- Name: TstUserVerifyCode TstUserVerifyCode_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstUserVerifyCode"
    ADD CONSTRAINT "TstUserVerifyCode_pkey" PRIMARY KEY (id);


--
-- Name: TstUser TstUser_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstUser"
    ADD CONSTRAINT "TstUser_pkey" PRIMARY KEY (id);


--
-- Name: TstVer TstVer_pkey; Type: CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstVer"
    ADD CONSTRAINT "TstVer_pkey" PRIMARY KEY (id);


--
-- Name: fki_CustomFieldOptionDefine_fieldId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_CustomFieldOptionDefine_fieldId_fkey" ON public."CustomFieldOptionDefine" USING btree ("fieldId");


--
-- Name: fki_CustomFieldOption_fieldId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_CustomFieldOption_fieldId_fkey" ON public."CustomFieldOption" USING btree ("fieldId");


--
-- Name: fki_CustomFieldOption_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_CustomFieldOption_orgId_fkey" ON public."CustomFieldOption" USING btree ("orgId");


--
-- Name: fki_CustomField_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_CustomField_orgId_fkey" ON public."CustomField" USING btree ("orgId");


--
-- Name: fki_IsuAttachment_issueId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuAttachment_issueId_fkey" ON public."IsuAttachment" USING btree ("issueId");


--
-- Name: fki_IsuAttachment_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuAttachment_userId_fkey" ON public."IsuAttachment" USING btree ("userId");


--
-- Name: fki_IsuComments_issueId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuComments_issueId_fkey" ON public."IsuComments" USING btree ("issueId");


--
-- Name: fki_IsuComments_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuComments_userId_fkey" ON public."IsuComments" USING btree ("userId");


--
-- Name: fki_IsuCustomFieldSolutionFieldRelation_fieldId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuCustomFieldSolutionFieldRelation_fieldId_fkey" ON public."IsuCustomFieldSolutionFieldRelation" USING btree ("fieldId");


--
-- Name: fki_IsuCustomFieldSolutionFieldRelation_solutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuCustomFieldSolutionFieldRelation_solutionId_fkey" ON public."IsuCustomFieldSolutionFieldRelation" USING btree ("solutionId");


--
-- Name: fki_IsuCustomFieldSolutionProjectRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuCustomFieldSolutionProjectRelation_orgId_fkey" ON public."IsuCustomFieldSolutionProjectRelation" USING btree ("orgId");


--
-- Name: fki_IsuCustomFieldSolutionProjectRelation_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuCustomFieldSolutionProjectRelation_projectId_fkey" ON public."IsuCustomFieldSolutionProjectRelation" USING btree ("projectId");


--
-- Name: fki_IsuCustomFieldSolutionProjectRelation_solutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuCustomFieldSolutionProjectRelation_solutionId_fkey" ON public."IsuCustomFieldSolutionProjectRelation" USING btree ("solutionId");


--
-- Name: fki_IsuCustomFieldSolution_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuCustomFieldSolution_orgId_fkey" ON public."IsuCustomFieldSolution" USING btree ("orgId");


--
-- Name: fki_IsuDocument_issueId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuDocument_issueId_fkey" ON public."IsuDocument" USING btree ("issueId");


--
-- Name: fki_IsuDocument_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuDocument_userId_fkey" ON public."IsuDocument" USING btree ("userId");


--
-- Name: fki_IsuField_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuField_orgId_fkey" ON public."IsuField" USING btree ("orgId");


--
-- Name: fki_IsuHistory_issueId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuHistory_issueId_fkey" ON public."IsuHistory" USING btree ("issueId");


--
-- Name: fki_IsuIssueExt_pid_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssueExt_pid_fkey" ON public."IsuIssueExt" USING btree (pid);


--
-- Name: fki_IsuIssue_assigneeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_assigneeId_fkey" ON public."IsuIssue" USING btree ("assigneeId");


--
-- Name: fki_IsuIssue_creatorId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_creatorId_fkey" ON public."IsuIssue" USING btree ("creatorId");


--
-- Name: fki_IsuIssue_envId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_envId_fkey" ON public."IsuIssue" USING btree ("envId");


--
-- Name: fki_IsuIssue_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_orgId_fkey" ON public."IsuIssue" USING btree ("orgId");


--
-- Name: fki_IsuIssue_priorityId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_priorityId_fkey" ON public."IsuIssue" USING btree ("priorityId");


--
-- Name: fki_IsuIssue_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_projectId_fkey" ON public."IsuIssue" USING btree ("projectId");


--
-- Name: fki_IsuIssue_reporterId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_reporterId_fkey" ON public."IsuIssue" USING btree ("reporterId");


--
-- Name: fki_IsuIssue_resolutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_resolutionId_fkey" ON public."IsuIssue" USING btree ("resolutionId");


--
-- Name: fki_IsuIssue_statusId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_statusId_fkey" ON public."IsuIssue" USING btree ("statusId");


--
-- Name: fki_IsuIssue_typeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_typeId_fkey" ON public."IsuIssue" USING btree ("typeId");


--
-- Name: fki_IsuIssue_verId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuIssue_verId_fkey" ON public."IsuIssue" USING btree ("verId");


--
-- Name: fki_IsuLink_dictIssueId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuLink_dictIssueId_fkey" ON public."IsuLink" USING btree ("dictIssueId");


--
-- Name: fki_IsuLink_reasonId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuLink_reasonId_fkey" ON public."IsuLink" USING btree ("reasonId");


--
-- Name: fki_IsuLink_srcIssueId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuLink_srcIssueId_fkey" ON public."IsuLink" USING btree ("srcIssueId");


--
-- Name: fki_IsuNotification_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuNotification_orgId_fkey" ON public."IsuNotification" USING btree ("orgId");


--
-- Name: fki_IsuPageSolutionItem_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPageSolutionItem_orgId_fkey" ON public."IsuPageSolutionItem" USING btree ("orgId");


--
-- Name: fki_IsuPageSolutionItem_pageId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPageSolutionItem_pageId_fkey" ON public."IsuPageSolutionItem" USING btree ("pageId");


--
-- Name: fki_IsuPageSolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPageSolutionItem_solutionId_fkey" ON public."IsuPageSolutionItem" USING btree ("solutionId");


--
-- Name: fki_IsuPageSolutionItem_typeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPageSolutionItem_typeId_fkey" ON public."IsuPageSolutionItem" USING btree ("typeId");


--
-- Name: fki_IsuPageSolution_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPageSolution_orgId_fkey" ON public."IsuPageSolution" USING btree ("orgId");


--
-- Name: fki_IsuPage_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPage_orgId_fkey" ON public."IsuPage" USING btree ("orgId");


--
-- Name: fki_IsuPrioritySolutionItem_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPrioritySolutionItem_orgId_fkey" ON public."IsuPrioritySolutionItem" USING btree ("orgId");


--
-- Name: fki_IsuPrioritySolutionItem_priorityId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPrioritySolutionItem_priorityId_fkey" ON public."IsuPrioritySolutionItem" USING btree ("priorityId");


--
-- Name: fki_IsuPrioritySolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPrioritySolutionItem_solutionId_fkey" ON public."IsuPrioritySolutionItem" USING btree ("solutionId");


--
-- Name: fki_IsuPrioritySolution_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPrioritySolution_orgId_fkey" ON public."IsuPrioritySolution" USING btree ("orgId");


--
-- Name: fki_IsuPriority_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuPriority_orgId_fkey" ON public."IsuPriority" USING btree ("orgId");


--
-- Name: fki_IsuQuery_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuQuery_projectId_fkey" ON public."IsuQuery" USING btree ("projectId");


--
-- Name: fki_IsuQuery_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuQuery_userId_fkey" ON public."IsuQuery" USING btree ("userId");


--
-- Name: fki_IsuResolution_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuResolution_orgId_fkey" ON public."IsuResolution" USING btree ("orgId");


--
-- Name: fki_IsuSeveritySolutionItem_severityId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuSeveritySolutionItem_severityId_fkey" ON public."IsuSeveritySolutionItem" USING btree ("severityId");


--
-- Name: fki_IsuSeveritySolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuSeveritySolutionItem_solutionId_fkey" ON public."IsuSeveritySolutionItem" USING btree ("solutionId");


--
-- Name: fki_IsuSeveritySolution_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuSeveritySolution_orgId_fkey" ON public."IsuSeveritySolution" USING btree ("orgId");


--
-- Name: fki_IsuSeverity_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuSeverity_orgId_fkey" ON public."IsuSeverity" USING btree ("orgId");


--
-- Name: fki_IsuStatusDefine_categoryId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuStatusDefine_categoryId_fkey" ON public."IsuStatusDefine" USING btree ("categoryId");


--
-- Name: fki_IsuStatus_categoryId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuStatus_categoryId_fkey" ON public."IsuStatus" USING btree ("categoryId");


--
-- Name: fki_IsuStatus_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuStatus_orgId_fkey" ON public."IsuStatus" USING btree ("orgId");


--
-- Name: fki_IsuTagRelation_issueId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuTagRelation_issueId_fkey" ON public."IsuTagRelation" USING btree ("issueId");


--
-- Name: fki_IsuTagRelation_tagId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuTagRelation_tagId_fkey" ON public."IsuTagRelation" USING btree ("tagId");


--
-- Name: fki_IsuTag_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuTag_orgId_fkey" ON public."IsuTag" USING btree ("orgId");


--
-- Name: fki_IsuTag_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuTag_userId_fkey" ON public."IsuTag" USING btree ("userId");


--
-- Name: fki_IsuTypeSolutionItem_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuTypeSolutionItem_orgId_fkey" ON public."IsuTypeSolutionItem" USING btree ("orgId");


--
-- Name: fki_IsuTypeSolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuTypeSolutionItem_solutionId_fkey" ON public."IsuTypeSolutionItem" USING btree ("solutionId");


--
-- Name: fki_IsuTypeSolutionItem_typeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuTypeSolutionItem_typeId_fkey" ON public."IsuTypeSolutionItem" USING btree ("typeId");


--
-- Name: fki_IsuTypeSolution_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuTypeSolution_orgId_fkey" ON public."IsuTypeSolution" USING btree ("orgId");


--
-- Name: fki_IsuType_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuType_orgId_fkey" ON public."IsuType" USING btree ("orgId");


--
-- Name: fki_IsuWatch_issueId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWatch_issueId_fkey" ON public."IsuWatch" USING btree ("issueId");


--
-- Name: fki_IsuWatch_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWatch_userId_fkey" ON public."IsuWatch" USING btree ("userId");


--
-- Name: fki_IsuWorkflowSolutionItem_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowSolutionItem_orgId_fkey" ON public."IsuWorkflowSolutionItem" USING btree ("orgId");


--
-- Name: fki_IsuWorkflowSolutionItem_solutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowSolutionItem_solutionId_fkey" ON public."IsuWorkflowSolutionItem" USING btree ("solutionId");


--
-- Name: fki_IsuWorkflowSolutionItem_typeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowSolutionItem_typeId_fkey" ON public."IsuWorkflowSolutionItem" USING btree ("typeId");


--
-- Name: fki_IsuWorkflowSolutionItem_workflowId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowSolutionItem_workflowId_fkey" ON public."IsuWorkflowSolutionItem" USING btree ("workflowId");


--
-- Name: fki_IsuWorkflowSolution_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowSolution_orgId_fkey" ON public."IsuWorkflowSolution" USING btree ("orgId");


--
-- Name: fki_IsuWorkflowStatusRelationDefine_statusId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowStatusRelationDefine_statusId_fkey" ON public."IsuWorkflowStatusRelationDefine" USING btree ("statusId");


--
-- Name: fki_IsuWorkflowStatusRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowStatusRelation_orgId_fkey" ON public."IsuWorkflowStatusRelation" USING btree ("orgId");


--
-- Name: fki_IsuWorkflowStatusRelation_statusId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowStatusRelation_statusId_fkey" ON public."IsuWorkflowStatusRelation" USING btree ("statusId");


--
-- Name: fki_IsuWorkflowStatusRelation_workflowId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowStatusRelation_workflowId_fkey" ON public."IsuWorkflowStatusRelation" USING btree ("workflowId");


--
-- Name: fki_IsuWorkflowTransitionDefine_dictStatusId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowTransitionDefine_dictStatusId_fkey" ON public."IsuWorkflowTransitionDefine" USING btree ("dictStatusId");


--
-- Name: fki_IsuWorkflowTransitionDefine_srcStatusId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowTransitionDefine_srcStatusId_fkey" ON public."IsuWorkflowTransitionDefine" USING btree ("srcStatusId");


--
-- Name: fki_IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_" ON public."IsuWorkflowTransitionProjectRoleRelation" USING btree ("workflowTransitionId");


--
-- Name: fki_IsuWorkflowTransitionProjectRoleRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowTransitionProjectRoleRelation_orgId_fkey" ON public."IsuWorkflowTransitionProjectRoleRelation" USING btree ("orgId");


--
-- Name: fki_IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey" ON public."IsuWorkflowTransitionProjectRoleRelation" USING btree ("projectRoleId");


--
-- Name: fki_IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey" ON public."IsuWorkflowTransitionProjectRoleRelation" USING btree ("workflowId");


--
-- Name: fki_IsuWorkflowTransition_actionPageId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowTransition_actionPageId_fkey" ON public."IsuWorkflowTransition" USING btree ("actionPageId");


--
-- Name: fki_IsuWorkflowTransition_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowTransition_orgId_fkey" ON public."IsuWorkflowTransition" USING btree ("orgId");


--
-- Name: fki_IsuWorkflowTransition_workflowId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflowTransition_workflowId_fkey" ON public."IsuWorkflowTransition" USING btree ("workflowId");


--
-- Name: fki_IsuWorkflow_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_IsuWorkflow_orgId_fkey" ON public."IsuWorkflow" USING btree ("orgId");


--
-- Name: fki_SysRolePrivilegeRelation_privilegeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_SysRolePrivilegeRelation_privilegeId_fkey" ON public."SysRolePrivilegeRelation" USING btree ("privilegeId");


--
-- Name: fki_SysRolePrivilegeRelation_roleId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_SysRolePrivilegeRelation_roleId_fkey" ON public."SysRolePrivilegeRelation" USING btree ("roleId");


--
-- Name: fki_SysRoleUserRelation_roleId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_SysRoleUserRelation_roleId_fkey" ON public."SysRoleUserRelation" USING btree ("roleId");


--
-- Name: fki_SysRoleUserRelation_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_SysRoleUserRelation_userId_fkey" ON public."SysRoleUserRelation" USING btree ("userId");


--
-- Name: fki_TstAlert_assigneeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstAlert_assigneeId_fkey" ON public."TstAlert" USING btree ("assigneeId");


--
-- Name: fki_TstAlert_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstAlert_userId_fkey" ON public."TstAlert" USING btree ("userId");


--
-- Name: fki_TstCaseAttachment_caseId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseAttachment_caseId_fkey" ON public."TstCaseAttachment" USING btree ("caseId");


--
-- Name: fki_TstCaseAttachment_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseAttachment_userId_fkey" ON public."TstCaseAttachment" USING btree ("userId");


--
-- Name: fki_TstCaseComments_caseId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseComments_caseId_fkey" ON public."TstCaseComments" USING btree ("caseId");


--
-- Name: fki_TstCaseComments_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseComments_userId_fkey" ON public."TstCaseComments" USING btree ("userId");


--
-- Name: fki_TstCaseExeStatus_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseExeStatus_orgId_fkey" ON public."TstCaseExeStatus" USING btree ("orgId");


--
-- Name: fki_TstCaseHistory_caseId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseHistory_caseId_fkey" ON public."TstCaseHistory" USING btree ("caseId");


--
-- Name: fki_TstCaseInSuite_caseId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInSuite_caseId_fkey" ON public."TstCaseInSuite" USING btree ("caseId");


--
-- Name: fki_TstCaseInSuite_pId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInSuite_pId_fkey" ON public."TstCaseInSuite" USING btree ("pId");


--
-- Name: fki_TstCaseInSuite_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInSuite_projectId_fkey" ON public."TstCaseInSuite" USING btree ("projectId");


--
-- Name: fki_TstCaseInSuite_suiteId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInSuite_suiteId_fkey" ON public."TstCaseInSuite" USING btree ("suiteId");


--
-- Name: fki_TstCaseInTaskAttachment_caseInTaskId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTaskAttachment_caseInTaskId_fkey" ON public."TstCaseInTaskAttachment" USING btree ("caseInTaskId");


--
-- Name: fki_TstCaseInTaskAttachment_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTaskAttachment_userId_fkey" ON public."TstCaseInTaskAttachment" USING btree ("userId");


--
-- Name: fki_TstCaseInTaskComments_caseInTaskId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTaskComments_caseInTaskId_fkey" ON public."TstCaseInTaskComments" USING btree ("caseInTaskId");


--
-- Name: fki_TstCaseInTaskComments_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTaskComments_userId_fkey" ON public."TstCaseInTaskComments" USING btree ("userId");


--
-- Name: fki_TstCaseInTaskHistory_caseId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTaskHistory_caseId_fkey" ON public."TstCaseInTaskHistory" USING btree ("caseId");


--
-- Name: fki_TstCaseInTaskHistory_caseInTaskId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTaskHistory_caseInTaskId_fkey" ON public."TstCaseInTaskHistory" USING btree ("caseInTaskId");


--
-- Name: fki_TstCaseInTaskIssue_caseInTaskId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTaskIssue_caseInTaskId_fkey" ON public."TstCaseInTaskIssue" USING btree ("caseInTaskId");


--
-- Name: fki_TstCaseInTaskIssue_issueId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTaskIssue_issueId_fkey" ON public."TstCaseInTaskIssue" USING btree ("issueId");


--
-- Name: fki_TstCaseInTaskIssue_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTaskIssue_userId_fkey" ON public."TstCaseInTaskIssue" USING btree ("userId");


--
-- Name: fki_TstCaseInTask_caseId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTask_caseId_fkey" ON public."TstCaseInTask" USING btree ("caseId");


--
-- Name: fki_TstCaseInTask_createBy_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTask_createBy_fkey" ON public."TstCaseInTask" USING btree ("createBy");


--
-- Name: fki_TstCaseInTask_exeBy_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTask_exeBy_fkey" ON public."TstCaseInTask" USING btree ("exeBy");


--
-- Name: fki_TstCaseInTask_pId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTask_pId_fkey" ON public."TstCaseInTask" USING btree ("pId");


--
-- Name: fki_TstCaseInTask_planId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTask_planId_fkey" ON public."TstCaseInTask" USING btree ("planId");


--
-- Name: fki_TstCaseInTask_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTask_projectId_fkey" ON public."TstCaseInTask" USING btree ("projectId");


--
-- Name: fki_TstCaseInTask_taskId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseInTask_taskId_fkey" ON public."TstCaseInTask" USING btree ("taskId");


--
-- Name: fki_TstCasePriority_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCasePriority_orgId_fkey" ON public."TstCasePriority" USING btree ("orgId");


--
-- Name: fki_TstCaseStep_caseId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseStep_caseId_fkey" ON public."TstCaseStep" USING btree ("caseId");


--
-- Name: fki_TstCaseType_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCaseType_orgId_fkey" ON public."TstCaseType" USING btree ("orgId");


--
-- Name: fki_TstCase_createById_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCase_createById_fkey" ON public."TstCase" USING btree ("createById");


--
-- Name: fki_TstCase_pId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCase_pId_fkey" ON public."TstCase" USING btree ("pId");


--
-- Name: fki_TstCase_priorityId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCase_priorityId_fkey" ON public."TstCase" USING btree ("priorityId");


--
-- Name: fki_TstCase_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCase_projectId_fkey" ON public."TstCase" USING btree ("projectId");


--
-- Name: fki_TstCase_typeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCase_typeId_fkey" ON public."TstCase" USING btree ("typeId");


--
-- Name: fki_TstCase_updateById_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstCase_updateById_fkey" ON public."TstCase" USING btree ("updateById");


--
-- Name: fki_TstDocument_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstDocument_userId_fkey" ON public."TstDocument" USING btree ("userId");


--
-- Name: fki_TstEnv_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstEnv_orgId_fkey" ON public."TstEnv" USING btree ("orgId");


--
-- Name: fki_TstEnv_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstEnv_projectId_fkey" ON public."TstEnv" USING btree ("projectId");


--
-- Name: fki_TstHistory_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstHistory_projectId_fkey" ON public."TstHistory" USING btree ("projectId");


--
-- Name: fki_TstHistory_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstHistory_userId_fkey" ON public."TstHistory" USING btree ("userId");


--
-- Name: fki_TstModule_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstModule_projectId_fkey" ON public."TstModule" USING btree ("projectId");


--
-- Name: fki_TstMsg_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstMsg_userId_fkey" ON public."TstMsg" USING btree ("userId");


--
-- Name: fki_TstOrgGroupUserRelation_orgGroupId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgGroupUserRelation_orgGroupId_fkey" ON public."TstOrgGroupUserRelation" USING btree ("orgGroupId");


--
-- Name: fki_TstOrgGroupUserRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgGroupUserRelation_orgId_fkey" ON public."TstOrgGroupUserRelation" USING btree ("orgId");


--
-- Name: fki_TstOrgGroupUserRelation_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgGroupUserRelation_userId_fkey" ON public."TstOrgGroupUserRelation" USING btree ("userId");


--
-- Name: fki_TstOrgGroup_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgGroup_orgId_fkey" ON public."TstOrgGroup" USING btree ("orgId");


--
-- Name: fki_TstOrgRoleGroupRelation_orgGroupId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRoleGroupRelation_orgGroupId_fkey" ON public."TstOrgRoleGroupRelation" USING btree ("orgGroupId");


--
-- Name: fki_TstOrgRoleGroupRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRoleGroupRelation_orgId_fkey" ON public."TstOrgRoleGroupRelation" USING btree ("orgId");


--
-- Name: fki_TstOrgRoleGroupRelation_orgRoleId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRoleGroupRelation_orgRoleId_fkey" ON public."TstOrgRoleGroupRelation" USING btree ("orgRoleId");


--
-- Name: fki_TstOrgRolePrivilegeRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRolePrivilegeRelation_orgId_fkey" ON public."TstOrgRolePrivilegeRelation" USING btree ("orgId");


--
-- Name: fki_TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey" ON public."TstOrgRolePrivilegeRelation" USING btree ("orgPrivilegeId");


--
-- Name: fki_TstOrgRolePrivilegeRelation_orgRoleId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRolePrivilegeRelation_orgRoleId_fkey" ON public."TstOrgRolePrivilegeRelation" USING btree ("orgRoleId");


--
-- Name: fki_TstOrgRoleUserRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRoleUserRelation_orgId_fkey" ON public."TstOrgRoleUserRelation" USING btree ("orgId");


--
-- Name: fki_TstOrgRoleUserRelation_orgRoleId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRoleUserRelation_orgRoleId_fkey" ON public."TstOrgRoleUserRelation" USING btree ("orgRoleId");


--
-- Name: fki_TstOrgRoleUserRelation_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRoleUserRelation_userId_fkey" ON public."TstOrgRoleUserRelation" USING btree ("userId");


--
-- Name: fki_TstOrgRole_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgRole_orgId_fkey" ON public."TstOrgRole" USING btree ("orgId");


--
-- Name: fki_TstOrgUserRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgUserRelation_orgId_fkey" ON public."TstOrgUserRelation" USING btree ("orgId");


--
-- Name: fki_TstOrgUserRelation_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstOrgUserRelation_userId_fkey" ON public."TstOrgUserRelation" USING btree ("userId");


--
-- Name: fki_TstPlan_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstPlan_projectId_fkey" ON public."TstPlan" USING btree ("projectId");


--
-- Name: fki_TstPlan_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstPlan_userId_fkey" ON public."TstPlan" USING btree ("userId");


--
-- Name: fki_TstPlan_verId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstPlan_verId_fkey" ON public."TstPlan" USING btree ("verId");


--
-- Name: fki_TstProjectAccessHistory_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectAccessHistory_orgId_fkey" ON public."TstProjectAccessHistory" USING btree ("orgId");


--
-- Name: fki_TstProjectAccessHistory_prjId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectAccessHistory_prjId_fkey" ON public."TstProjectAccessHistory" USING btree ("prjId");


--
-- Name: fki_TstProjectAccessHistory_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectAccessHistory_userId_fkey" ON public."TstProjectAccessHistory" USING btree ("userId");


--
-- Name: fki_TstProjectRoleEntityRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectRoleEntityRelation_orgId_fkey" ON public."TstProjectRoleEntityRelation" USING btree ("orgId");


--
-- Name: fki_TstProjectRoleEntityRelation_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectRoleEntityRelation_projectId_fkey" ON public."TstProjectRoleEntityRelation" USING btree ("projectId");


--
-- Name: fki_TstProjectRoleEntityRelation_projectRoleId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectRoleEntityRelation_projectRoleId_fkey" ON public."TstProjectRoleEntityRelation" USING btree ("projectRoleId");


--
-- Name: fki_TstProjectRolePriviledgeRelation_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectRolePriviledgeRelation_orgId_fkey" ON public."TstProjectRolePriviledgeRelation" USING btree ("orgId");


--
-- Name: fki_TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_f; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_f" ON public."TstProjectRolePriviledgeRelation" USING btree ("projectPrivilegeDefineId");


--
-- Name: fki_TstProjectRolePriviledgeRelation_projectRoleId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectRolePriviledgeRelation_projectRoleId_fkey" ON public."TstProjectRolePriviledgeRelation" USING btree ("projectRoleId");


--
-- Name: fki_TstProjectRole_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProjectRole_orgId_fkey" ON public."TstProjectRole" USING btree ("orgId");


--
-- Name: fki_TstProject_issuePageSolutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProject_issuePageSolutionId_fkey" ON public."TstProject" USING btree ("issuePageSolutionId");


--
-- Name: fki_TstProject_issuePrioritySolutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProject_issuePrioritySolutionId_fkey" ON public."TstProject" USING btree ("issuePrioritySolutionId");


--
-- Name: fki_TstProject_issueTypeSolutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProject_issueTypeSolutionId_fkey" ON public."TstProject" USING btree ("issueTypeSolutionId");


--
-- Name: fki_TstProject_issueWorkflowSolutionId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProject_issueWorkflowSolutionId_fkey" ON public."TstProject" USING btree ("issueWorkflowSolutionId");


--
-- Name: fki_TstProject_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProject_orgId_fkey" ON public."TstProject" USING btree ("orgId");


--
-- Name: fki_TstProject_parentId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstProject_parentId_fkey" ON public."TstProject" USING btree ("parentId");


--
-- Name: fki_TstSuite_caseProjectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstSuite_caseProjectId_fkey" ON public."TstSuite" USING btree ("caseProjectId");


--
-- Name: fki_TstSuite_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstSuite_projectId_fkey" ON public."TstSuite" USING btree ("projectId");


--
-- Name: fki_TstSuite_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstSuite_userId_fkey" ON public."TstSuite" USING btree ("userId");


--
-- Name: fki_TstTaskAssigneeRelation_assigneeId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstTaskAssigneeRelation_assigneeId_fkey" ON public."TstTaskAssigneeRelation" USING btree ("assigneeId");


--
-- Name: fki_TstTaskAssigneeRelation_taskId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstTaskAssigneeRelation_taskId_fkey" ON public."TstTaskAssigneeRelation" USING btree ("taskId");


--
-- Name: fki_TstTask_caseProjectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstTask_caseProjectId_fkey" ON public."TstTask" USING btree ("caseProjectId");


--
-- Name: fki_TstTask_envId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstTask_envId_fkey" ON public."TstTask" USING btree ("envId");


--
-- Name: fki_TstTask_planId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstTask_planId_fkey" ON public."TstTask" USING btree ("planId");


--
-- Name: fki_TstTask_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstTask_projectId_fkey" ON public."TstTask" USING btree ("projectId");


--
-- Name: fki_TstTask_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstTask_userId_fkey" ON public."TstTask" USING btree ("userId");


--
-- Name: fki_TstThread_authorId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstThread_authorId_fkey" ON public."TstThread" USING btree ("authorId");


--
-- Name: fki_TstThread_parentId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstThread_parentId_fkey" ON public."TstThread" USING btree ("parentId");


--
-- Name: fki_TstUserSettings_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstUserSettings_userId_fkey" ON public."TstUserSettings" USING btree ("userId");


--
-- Name: fki_TstUserVerifyCode_userId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstUserVerifyCode_userId_fkey" ON public."TstUserVerifyCode" USING btree ("userId");


--
-- Name: fki_TstUser_defaultOrgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstUser_defaultOrgId_fkey" ON public."TstUser" USING btree ("defaultOrgId");


--
-- Name: fki_TstUser_defaultPrjId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstUser_defaultPrjId_fkey" ON public."TstUser" USING btree ("defaultPrjId");


--
-- Name: fki_TstVer_orgId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstVer_orgId_fkey" ON public."TstVer" USING btree ("orgId");


--
-- Name: fki_TstVer_projectId_fkey; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX "fki_TstVer_projectId_fkey" ON public."TstVer" USING btree ("projectId");


--
-- Name: idx_isu_issue_extprop; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX idx_isu_issue_extprop ON public."IsuIssue" USING gin ("extProp" jsonb_path_ops);


--
-- Name: idx_test_case_extprop; Type: INDEX; Schema: public; Owner: ngtesting
--

CREATE INDEX idx_test_case_extprop ON public."TstCase" USING gin ("extProp" jsonb_path_ops);


--
-- Name: IsuIssue issue_tsvector_update_trigger; Type: TRIGGER; Schema: public; Owner: ngtesting
--

CREATE TRIGGER issue_tsvector_update_trigger AFTER INSERT OR UPDATE OF title, tag, "extProp", descr ON public."IsuIssue" FOR EACH ROW EXECUTE PROCEDURE public.update_issue_tsv_content();


--
-- Name: CustomFieldOptionDefine CustomFieldOptionDefine_fieldId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldOptionDefine"
    ADD CONSTRAINT "CustomFieldOptionDefine_fieldId_fkey" FOREIGN KEY ("fieldId") REFERENCES public."CustomFieldDefine"(id);


--
-- Name: CustomFieldOption CustomFieldOption_fieldId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldOption"
    ADD CONSTRAINT "CustomFieldOption_fieldId_fkey" FOREIGN KEY ("fieldId") REFERENCES public."CustomField"(id);


--
-- Name: CustomFieldOption CustomFieldOption_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomFieldOption"
    ADD CONSTRAINT "CustomFieldOption_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: CustomField CustomField_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."CustomField"
    ADD CONSTRAINT "CustomField_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuAttachment IsuAttachment_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuAttachment"
    ADD CONSTRAINT "IsuAttachment_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- Name: IsuAttachment IsuAttachment_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuAttachment"
    ADD CONSTRAINT "IsuAttachment_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: IsuComments IsuComments_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuComments"
    ADD CONSTRAINT "IsuComments_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- Name: IsuComments IsuComments_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuComments"
    ADD CONSTRAINT "IsuComments_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: IsuCustomFieldSolutionFieldRelation IsuCustomFieldSolutionFieldRelation_fieldId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionFieldRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionFieldRelation_fieldId_fkey" FOREIGN KEY ("fieldId") REFERENCES public."CustomField"(id);


--
-- Name: IsuCustomFieldSolutionFieldRelation IsuCustomFieldSolutionFieldRelation_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionFieldRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionFieldRelation_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuCustomFieldSolution"(id);


--
-- Name: IsuCustomFieldSolutionProjectRelation IsuCustomFieldSolutionProjectRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionProjectRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionProjectRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuCustomFieldSolutionProjectRelation IsuCustomFieldSolutionProjectRelation_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionProjectRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionProjectRelation_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: IsuCustomFieldSolutionProjectRelation IsuCustomFieldSolutionProjectRelation_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuCustomFieldSolutionProjectRelation"
    ADD CONSTRAINT "IsuCustomFieldSolutionProjectRelation_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuCustomFieldSolution"(id);


--
-- Name: IsuCustomFieldSolution IsuCustomFieldSolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuCustomFieldSolution"
    ADD CONSTRAINT "IsuCustomFieldSolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuDocument IsuDocument_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuDocument"
    ADD CONSTRAINT "IsuDocument_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- Name: IsuDocument IsuDocument_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuDocument"
    ADD CONSTRAINT "IsuDocument_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: IsuField IsuField_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuField"
    ADD CONSTRAINT "IsuField_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuHistory IsuHistory_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuHistory"
    ADD CONSTRAINT "IsuHistory_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- Name: IsuIssueExt IsuIssueExt_pid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssueExt"
    ADD CONSTRAINT "IsuIssueExt_pid_fkey" FOREIGN KEY (pid) REFERENCES public."IsuIssue"(id);


--
-- Name: IsuIssue IsuIssue_assigneeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_assigneeId_fkey" FOREIGN KEY ("assigneeId") REFERENCES public."TstUser"(id);


--
-- Name: IsuIssue IsuIssue_creatorId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_creatorId_fkey" FOREIGN KEY ("creatorId") REFERENCES public."TstUser"(id);


--
-- Name: IsuIssue IsuIssue_envId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_envId_fkey" FOREIGN KEY ("envId") REFERENCES public."TstEnv"(id);


--
-- Name: IsuIssue IsuIssue_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuIssue IsuIssue_priorityId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_priorityId_fkey" FOREIGN KEY ("priorityId") REFERENCES public."IsuPriority"(id);


--
-- Name: IsuIssue IsuIssue_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: IsuIssue IsuIssue_reporterId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_reporterId_fkey" FOREIGN KEY ("reporterId") REFERENCES public."TstUser"(id);


--
-- Name: IsuIssue IsuIssue_resolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_resolutionId_fkey" FOREIGN KEY ("resolutionId") REFERENCES public."IsuResolution"(id);


--
-- Name: IsuIssue IsuIssue_statusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_statusId_fkey" FOREIGN KEY ("statusId") REFERENCES public."IsuStatus"(id);


--
-- Name: IsuIssue IsuIssue_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."IsuType"(id);


--
-- Name: IsuIssue IsuIssue_verId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuIssue"
    ADD CONSTRAINT "IsuIssue_verId_fkey" FOREIGN KEY ("verId") REFERENCES public."TstVer"(id);


--
-- Name: IsuLink IsuLink_dictIssueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuLink"
    ADD CONSTRAINT "IsuLink_dictIssueId_fkey" FOREIGN KEY ("dictIssueId") REFERENCES public."IsuIssue"(id);


--
-- Name: IsuLink IsuLink_reasonId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuLink"
    ADD CONSTRAINT "IsuLink_reasonId_fkey" FOREIGN KEY ("reasonId") REFERENCES public."IsuLinkReasonDefine"(id);


--
-- Name: IsuLink IsuLink_srcIssueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuLink"
    ADD CONSTRAINT "IsuLink_srcIssueId_fkey" FOREIGN KEY ("srcIssueId") REFERENCES public."IsuIssue"(id);


--
-- Name: IsuNotification IsuNotification_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuNotification"
    ADD CONSTRAINT "IsuNotification_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuPageSolutionItem IsuPageSolutionItem_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuPageSolutionItem IsuPageSolutionItem_pageId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_pageId_fkey" FOREIGN KEY ("pageId") REFERENCES public."IsuPage"(id);


--
-- Name: IsuPageSolutionItem IsuPageSolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuPageSolution"(id);


--
-- Name: IsuPageSolutionItem IsuPageSolutionItem_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageSolutionItem"
    ADD CONSTRAINT "IsuPageSolutionItem_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."IsuType"(id);


--
-- Name: IsuPageSolution IsuPageSolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPageSolution"
    ADD CONSTRAINT "IsuPageSolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuPage IsuPage_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPage"
    ADD CONSTRAINT "IsuPage_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuPrioritySolutionItem IsuPrioritySolutionItem_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPrioritySolutionItem"
    ADD CONSTRAINT "IsuPrioritySolutionItem_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuPrioritySolutionItem IsuPrioritySolutionItem_priorityId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPrioritySolutionItem"
    ADD CONSTRAINT "IsuPrioritySolutionItem_priorityId_fkey" FOREIGN KEY ("priorityId") REFERENCES public."IsuPriority"(id);


--
-- Name: IsuPrioritySolutionItem IsuPrioritySolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPrioritySolutionItem"
    ADD CONSTRAINT "IsuPrioritySolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuPrioritySolution"(id);


--
-- Name: IsuPrioritySolution IsuPrioritySolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPrioritySolution"
    ADD CONSTRAINT "IsuPrioritySolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuPriority IsuPriority_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuPriority"
    ADD CONSTRAINT "IsuPriority_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuQuery IsuQuery_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuQuery"
    ADD CONSTRAINT "IsuQuery_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: IsuQuery IsuQuery_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuQuery"
    ADD CONSTRAINT "IsuQuery_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: IsuResolution IsuResolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuResolution"
    ADD CONSTRAINT "IsuResolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuSeveritySolutionItem IsuSeveritySolutionItem_severityId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeveritySolutionItem"
    ADD CONSTRAINT "IsuSeveritySolutionItem_severityId_fkey" FOREIGN KEY ("severityId") REFERENCES public."IsuSeverity"(id);


--
-- Name: IsuSeveritySolutionItem IsuSeveritySolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeveritySolutionItem"
    ADD CONSTRAINT "IsuSeveritySolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuSeveritySolution"(id);


--
-- Name: IsuSeveritySolution IsuSeveritySolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeveritySolution"
    ADD CONSTRAINT "IsuSeveritySolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuSeverity IsuSeverity_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuSeverity"
    ADD CONSTRAINT "IsuSeverity_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuStatusDefine IsuStatusDefine_categoryId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuStatusDefine"
    ADD CONSTRAINT "IsuStatusDefine_categoryId_fkey" FOREIGN KEY ("categoryId") REFERENCES public."IsuStatusCategoryDefine"(id);


--
-- Name: IsuStatus IsuStatus_categoryId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuStatus"
    ADD CONSTRAINT "IsuStatus_categoryId_fkey" FOREIGN KEY ("categoryId") REFERENCES public."IsuStatusCategoryDefine"(id);


--
-- Name: IsuStatus IsuStatus_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuStatus"
    ADD CONSTRAINT "IsuStatus_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuTagRelation IsuTagRelation_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTagRelation"
    ADD CONSTRAINT "IsuTagRelation_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- Name: IsuTagRelation IsuTagRelation_tagId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTagRelation"
    ADD CONSTRAINT "IsuTagRelation_tagId_fkey" FOREIGN KEY ("tagId") REFERENCES public."IsuTag"(id);


--
-- Name: IsuTag IsuTag_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTag"
    ADD CONSTRAINT "IsuTag_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuTag IsuTag_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTag"
    ADD CONSTRAINT "IsuTag_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: IsuTypeSolutionItem IsuTypeSolutionItem_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTypeSolutionItem"
    ADD CONSTRAINT "IsuTypeSolutionItem_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuTypeSolutionItem IsuTypeSolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTypeSolutionItem"
    ADD CONSTRAINT "IsuTypeSolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuTypeSolution"(id);


--
-- Name: IsuTypeSolutionItem IsuTypeSolutionItem_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTypeSolutionItem"
    ADD CONSTRAINT "IsuTypeSolutionItem_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."IsuType"(id);


--
-- Name: IsuTypeSolution IsuTypeSolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuTypeSolution"
    ADD CONSTRAINT "IsuTypeSolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuType IsuType_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuType"
    ADD CONSTRAINT "IsuType_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuWatch IsuWatch_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWatch"
    ADD CONSTRAINT "IsuWatch_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- Name: IsuWatch IsuWatch_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWatch"
    ADD CONSTRAINT "IsuWatch_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_solutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_solutionId_fkey" FOREIGN KEY ("solutionId") REFERENCES public."IsuWorkflowSolution"(id);


--
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."IsuType"(id);


--
-- Name: IsuWorkflowSolutionItem IsuWorkflowSolutionItem_workflowId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowSolutionItem"
    ADD CONSTRAINT "IsuWorkflowSolutionItem_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES public."IsuWorkflow"(id);


--
-- Name: IsuWorkflowSolution IsuWorkflowSolution_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowSolution"
    ADD CONSTRAINT "IsuWorkflowSolution_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuWorkflowStatusRelationDefine IsuWorkflowStatusRelationDefine_statusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelationDefine"
    ADD CONSTRAINT "IsuWorkflowStatusRelationDefine_statusId_fkey" FOREIGN KEY ("statusId") REFERENCES public."IsuStatusDefine"(id);


--
-- Name: IsuWorkflowStatusRelation IsuWorkflowStatusRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation"
    ADD CONSTRAINT "IsuWorkflowStatusRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuWorkflowStatusRelation IsuWorkflowStatusRelation_statusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation"
    ADD CONSTRAINT "IsuWorkflowStatusRelation_statusId_fkey" FOREIGN KEY ("statusId") REFERENCES public."IsuStatus"(id);


--
-- Name: IsuWorkflowStatusRelation IsuWorkflowStatusRelation_workflowId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowStatusRelation"
    ADD CONSTRAINT "IsuWorkflowStatusRelation_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES public."IsuWorkflow"(id);


--
-- Name: IsuWorkflowTransitionDefine IsuWorkflowTransitionDefine_dictStatusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionDefine"
    ADD CONSTRAINT "IsuWorkflowTransitionDefine_dictStatusId_fkey" FOREIGN KEY ("dictStatusId") REFERENCES public."IsuStatusDefine"(id);


--
-- Name: IsuWorkflowTransitionDefine IsuWorkflowTransitionDefine_srcStatusId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionDefine"
    ADD CONSTRAINT "IsuWorkflowTransitionDefine_srcStatusId_fkey" FOREIGN KEY ("srcStatusId") REFERENCES public."IsuStatusDefine"(id);


--
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelat_workflowTransitionId_fkey" FOREIGN KEY ("workflowTransitionId") REFERENCES public."IsuWorkflowTransition"(id);


--
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_projectRoleId_fkey" FOREIGN KEY ("projectRoleId") REFERENCES public."TstProjectRole"(id);


--
-- Name: IsuWorkflowTransitionProjectRoleRelation IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransitionProjectRoleRelation"
    ADD CONSTRAINT "IsuWorkflowTransitionProjectRoleRelation_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES public."IsuWorkflow"(id);


--
-- Name: IsuWorkflowTransition IsuWorkflowTransition_actionPageId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransition"
    ADD CONSTRAINT "IsuWorkflowTransition_actionPageId_fkey" FOREIGN KEY ("actionPageId") REFERENCES public."IsuPage"(id);


--
-- Name: IsuWorkflowTransition IsuWorkflowTransition_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransition"
    ADD CONSTRAINT "IsuWorkflowTransition_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: IsuWorkflowTransition IsuWorkflowTransition_workflowId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflowTransition"
    ADD CONSTRAINT "IsuWorkflowTransition_workflowId_fkey" FOREIGN KEY ("workflowId") REFERENCES public."IsuWorkflow"(id);


--
-- Name: IsuWorkflow IsuWorkflow_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."IsuWorkflow"
    ADD CONSTRAINT "IsuWorkflow_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: SysRolePrivilegeRelation SysRolePrivilegeRelation_privilegeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysRolePrivilegeRelation"
    ADD CONSTRAINT "SysRolePrivilegeRelation_privilegeId_fkey" FOREIGN KEY ("privilegeId") REFERENCES public."SysPrivilege"(id);


--
-- Name: SysRolePrivilegeRelation SysRolePrivilegeRelation_roleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysRolePrivilegeRelation"
    ADD CONSTRAINT "SysRolePrivilegeRelation_roleId_fkey" FOREIGN KEY ("roleId") REFERENCES public."SysRole"(id);


--
-- Name: SysRoleUserRelation SysRoleUserRelation_roleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysRoleUserRelation"
    ADD CONSTRAINT "SysRoleUserRelation_roleId_fkey" FOREIGN KEY ("roleId") REFERENCES public."SysRole"(id);


--
-- Name: SysRoleUserRelation SysRoleUserRelation_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."SysRoleUserRelation"
    ADD CONSTRAINT "SysRoleUserRelation_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstAlert TstAlert_assigneeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstAlert"
    ADD CONSTRAINT "TstAlert_assigneeId_fkey" FOREIGN KEY ("assigneeId") REFERENCES public."TstUser"(id);


--
-- Name: TstAlert TstAlert_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstAlert"
    ADD CONSTRAINT "TstAlert_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstCaseAttachment TstCaseAttachment_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseAttachment"
    ADD CONSTRAINT "TstCaseAttachment_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- Name: TstCaseAttachment TstCaseAttachment_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseAttachment"
    ADD CONSTRAINT "TstCaseAttachment_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstCaseComments TstCaseComments_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseComments"
    ADD CONSTRAINT "TstCaseComments_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- Name: TstCaseComments TstCaseComments_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseComments"
    ADD CONSTRAINT "TstCaseComments_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstCaseExeStatus TstCaseExeStatus_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseExeStatus"
    ADD CONSTRAINT "TstCaseExeStatus_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstCaseHistory TstCaseHistory_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseHistory"
    ADD CONSTRAINT "TstCaseHistory_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- Name: TstCaseInSuite TstCaseInSuite_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- Name: TstCaseInSuite TstCaseInSuite_pId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_pId_fkey" FOREIGN KEY ("pId") REFERENCES public."TstCase"(id);


--
-- Name: TstCaseInSuite TstCaseInSuite_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstCaseInSuite TstCaseInSuite_suiteId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInSuite"
    ADD CONSTRAINT "TstCaseInSuite_suiteId_fkey" FOREIGN KEY ("suiteId") REFERENCES public."TstSuite"(id);


--
-- Name: TstCaseInTaskAttachment TstCaseInTaskAttachment_caseInTaskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskAttachment"
    ADD CONSTRAINT "TstCaseInTaskAttachment_caseInTaskId_fkey" FOREIGN KEY ("caseInTaskId") REFERENCES public."TstCaseInTask"(id);


--
-- Name: TstCaseInTaskAttachment TstCaseInTaskAttachment_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskAttachment"
    ADD CONSTRAINT "TstCaseInTaskAttachment_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstCaseInTaskComments TstCaseInTaskComments_caseInTaskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskComments"
    ADD CONSTRAINT "TstCaseInTaskComments_caseInTaskId_fkey" FOREIGN KEY ("caseInTaskId") REFERENCES public."TstCaseInTask"(id);


--
-- Name: TstCaseInTaskComments TstCaseInTaskComments_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskComments"
    ADD CONSTRAINT "TstCaseInTaskComments_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstCaseInTaskHistory TstCaseInTaskHistory_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskHistory"
    ADD CONSTRAINT "TstCaseInTaskHistory_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- Name: TstCaseInTaskHistory TstCaseInTaskHistory_caseInTaskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskHistory"
    ADD CONSTRAINT "TstCaseInTaskHistory_caseInTaskId_fkey" FOREIGN KEY ("caseInTaskId") REFERENCES public."TstCaseInTask"(id);


--
-- Name: TstCaseInTaskIssue TstCaseInTaskIssue_caseInTaskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskIssue"
    ADD CONSTRAINT "TstCaseInTaskIssue_caseInTaskId_fkey" FOREIGN KEY ("caseInTaskId") REFERENCES public."TstCaseInTask"(id);


--
-- Name: TstCaseInTaskIssue TstCaseInTaskIssue_issueId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskIssue"
    ADD CONSTRAINT "TstCaseInTaskIssue_issueId_fkey" FOREIGN KEY ("issueId") REFERENCES public."IsuIssue"(id);


--
-- Name: TstCaseInTaskIssue TstCaseInTaskIssue_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTaskIssue"
    ADD CONSTRAINT "TstCaseInTaskIssue_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstCaseInTask TstCaseInTask_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- Name: TstCaseInTask TstCaseInTask_createBy_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_createBy_fkey" FOREIGN KEY ("createBy") REFERENCES public."TstUser"(id);


--
-- Name: TstCaseInTask TstCaseInTask_exeBy_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_exeBy_fkey" FOREIGN KEY ("exeBy") REFERENCES public."TstUser"(id);


--
-- Name: TstCaseInTask TstCaseInTask_pId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_pId_fkey" FOREIGN KEY ("pId") REFERENCES public."TstCase"(id);


--
-- Name: TstCaseInTask TstCaseInTask_planId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_planId_fkey" FOREIGN KEY ("planId") REFERENCES public."TstPlan"(id);


--
-- Name: TstCaseInTask TstCaseInTask_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstCaseInTask TstCaseInTask_taskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseInTask"
    ADD CONSTRAINT "TstCaseInTask_taskId_fkey" FOREIGN KEY ("taskId") REFERENCES public."TstTask"(id);


--
-- Name: TstCasePriority TstCasePriority_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCasePriority"
    ADD CONSTRAINT "TstCasePriority_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstCaseStep TstCaseStep_caseId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseStep"
    ADD CONSTRAINT "TstCaseStep_caseId_fkey" FOREIGN KEY ("caseId") REFERENCES public."TstCase"(id);


--
-- Name: TstCaseType TstCaseType_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCaseType"
    ADD CONSTRAINT "TstCaseType_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstCase TstCase_createById_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_createById_fkey" FOREIGN KEY ("createById") REFERENCES public."TstUser"(id);


--
-- Name: TstCase TstCase_pId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_pId_fkey" FOREIGN KEY ("pId") REFERENCES public."TstCase"(id);


--
-- Name: TstCase TstCase_priorityId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_priorityId_fkey" FOREIGN KEY ("priorityId") REFERENCES public."TstCasePriority"(id);


--
-- Name: TstCase TstCase_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstCase TstCase_typeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_typeId_fkey" FOREIGN KEY ("typeId") REFERENCES public."TstCaseType"(id);


--
-- Name: TstCase TstCase_updateById_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstCase"
    ADD CONSTRAINT "TstCase_updateById_fkey" FOREIGN KEY ("updateById") REFERENCES public."TstUser"(id);


--
-- Name: TstDocument TstDocument_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstDocument"
    ADD CONSTRAINT "TstDocument_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstEnv TstEnv_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstEnv"
    ADD CONSTRAINT "TstEnv_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstEnv TstEnv_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstEnv"
    ADD CONSTRAINT "TstEnv_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstHistory TstHistory_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstHistory"
    ADD CONSTRAINT "TstHistory_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstHistory TstHistory_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstHistory"
    ADD CONSTRAINT "TstHistory_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstModule TstModule_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstModule"
    ADD CONSTRAINT "TstModule_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstMsg TstMsg_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstMsg"
    ADD CONSTRAINT "TstMsg_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstOrgGroupUserRelation TstOrgGroupUserRelation_orgGroupId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgGroupUserRelation"
    ADD CONSTRAINT "TstOrgGroupUserRelation_orgGroupId_fkey" FOREIGN KEY ("orgGroupId") REFERENCES public."TstOrgGroup"(id);


--
-- Name: TstOrgGroupUserRelation TstOrgGroupUserRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgGroupUserRelation"
    ADD CONSTRAINT "TstOrgGroupUserRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstOrgGroupUserRelation TstOrgGroupUserRelation_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgGroupUserRelation"
    ADD CONSTRAINT "TstOrgGroupUserRelation_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstOrgGroup TstOrgGroup_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgGroup"
    ADD CONSTRAINT "TstOrgGroup_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstOrgRoleGroupRelation TstOrgRoleGroupRelation_orgGroupId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRoleGroupRelation"
    ADD CONSTRAINT "TstOrgRoleGroupRelation_orgGroupId_fkey" FOREIGN KEY ("orgGroupId") REFERENCES public."TstOrgGroup"(id);


--
-- Name: TstOrgRoleGroupRelation TstOrgRoleGroupRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRoleGroupRelation"
    ADD CONSTRAINT "TstOrgRoleGroupRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstOrgRoleGroupRelation TstOrgRoleGroupRelation_orgRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRoleGroupRelation"
    ADD CONSTRAINT "TstOrgRoleGroupRelation_orgRoleId_fkey" FOREIGN KEY ("orgRoleId") REFERENCES public."TstOrgRole"(id);


--
-- Name: TstOrgRolePrivilegeRelation TstOrgRolePrivilegeRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRolePrivilegeRelation"
    ADD CONSTRAINT "TstOrgRolePrivilegeRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstOrgRolePrivilegeRelation TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRolePrivilegeRelation"
    ADD CONSTRAINT "TstOrgRolePrivilegeRelation_orgPrivilegeId_fkey" FOREIGN KEY ("orgPrivilegeId") REFERENCES public."TstOrgPrivilegeDefine"(id);


--
-- Name: TstOrgRolePrivilegeRelation TstOrgRolePrivilegeRelation_orgRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRolePrivilegeRelation"
    ADD CONSTRAINT "TstOrgRolePrivilegeRelation_orgRoleId_fkey" FOREIGN KEY ("orgRoleId") REFERENCES public."TstOrgRole"(id);


--
-- Name: TstOrgRoleUserRelation TstOrgRoleUserRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRoleUserRelation"
    ADD CONSTRAINT "TstOrgRoleUserRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstOrgRoleUserRelation TstOrgRoleUserRelation_orgRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRoleUserRelation"
    ADD CONSTRAINT "TstOrgRoleUserRelation_orgRoleId_fkey" FOREIGN KEY ("orgRoleId") REFERENCES public."TstOrgRole"(id);


--
-- Name: TstOrgRoleUserRelation TstOrgRoleUserRelation_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRoleUserRelation"
    ADD CONSTRAINT "TstOrgRoleUserRelation_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstOrgRole TstOrgRole_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgRole"
    ADD CONSTRAINT "TstOrgRole_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstOrgUserRelation TstOrgUserRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgUserRelation"
    ADD CONSTRAINT "TstOrgUserRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstOrgUserRelation TstOrgUserRelation_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstOrgUserRelation"
    ADD CONSTRAINT "TstOrgUserRelation_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstPlan TstPlan_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstPlan"
    ADD CONSTRAINT "TstPlan_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstPlan TstPlan_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstPlan"
    ADD CONSTRAINT "TstPlan_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstPlan TstPlan_verId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstPlan"
    ADD CONSTRAINT "TstPlan_verId_fkey" FOREIGN KEY ("verId") REFERENCES public."TstVer"(id);


--
-- Name: TstProjectAccessHistory TstProjectAccessHistory_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectAccessHistory"
    ADD CONSTRAINT "TstProjectAccessHistory_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstProjectAccessHistory TstProjectAccessHistory_prjId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectAccessHistory"
    ADD CONSTRAINT "TstProjectAccessHistory_prjId_fkey" FOREIGN KEY ("prjId") REFERENCES public."TstProject"(id);


--
-- Name: TstProjectAccessHistory TstProjectAccessHistory_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectAccessHistory"
    ADD CONSTRAINT "TstProjectAccessHistory_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstProjectRoleEntityRelation TstProjectRoleEntityRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectRoleEntityRelation"
    ADD CONSTRAINT "TstProjectRoleEntityRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstProjectRoleEntityRelation TstProjectRoleEntityRelation_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectRoleEntityRelation"
    ADD CONSTRAINT "TstProjectRoleEntityRelation_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstProjectRoleEntityRelation TstProjectRoleEntityRelation_projectRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectRoleEntityRelation"
    ADD CONSTRAINT "TstProjectRoleEntityRelation_projectRoleId_fkey" FOREIGN KEY ("projectRoleId") REFERENCES public."TstProjectRole"(id);


--
-- Name: TstProjectRolePriviledgeRelation TstProjectRolePriviledgeRelation_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectRolePriviledgeRelation"
    ADD CONSTRAINT "TstProjectRolePriviledgeRelation_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstProjectRolePriviledgeRelation TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectRolePriviledgeRelation"
    ADD CONSTRAINT "TstProjectRolePriviledgeRelation_projectPrivilegeDefineId_fkey" FOREIGN KEY ("projectPrivilegeDefineId") REFERENCES public."TstProjectPrivilegeDefine"(id);


--
-- Name: TstProjectRolePriviledgeRelation TstProjectRolePriviledgeRelation_projectRoleId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectRolePriviledgeRelation"
    ADD CONSTRAINT "TstProjectRolePriviledgeRelation_projectRoleId_fkey" FOREIGN KEY ("projectRoleId") REFERENCES public."TstProjectRole"(id);


--
-- Name: TstProjectRole TstProjectRole_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProjectRole"
    ADD CONSTRAINT "TstProjectRole_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstProject TstProject_issuePageSolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_issuePageSolutionId_fkey" FOREIGN KEY ("issuePageSolutionId") REFERENCES public."IsuPageSolution"(id);


--
-- Name: TstProject TstProject_issuePrioritySolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_issuePrioritySolutionId_fkey" FOREIGN KEY ("issuePrioritySolutionId") REFERENCES public."IsuPrioritySolution"(id);


--
-- Name: TstProject TstProject_issueTypeSolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_issueTypeSolutionId_fkey" FOREIGN KEY ("issueTypeSolutionId") REFERENCES public."IsuTypeSolution"(id);


--
-- Name: TstProject TstProject_issueWorkflowSolutionId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_issueWorkflowSolutionId_fkey" FOREIGN KEY ("issueWorkflowSolutionId") REFERENCES public."IsuWorkflowSolution"(id);


--
-- Name: TstProject TstProject_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstProject TstProject_parentId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstProject"
    ADD CONSTRAINT "TstProject_parentId_fkey" FOREIGN KEY ("parentId") REFERENCES public."TstProject"(id);


--
-- Name: TstSuite TstSuite_caseProjectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstSuite"
    ADD CONSTRAINT "TstSuite_caseProjectId_fkey" FOREIGN KEY ("caseProjectId") REFERENCES public."TstProject"(id);


--
-- Name: TstSuite TstSuite_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstSuite"
    ADD CONSTRAINT "TstSuite_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstSuite TstSuite_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstSuite"
    ADD CONSTRAINT "TstSuite_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstTaskAssigneeRelation TstTaskAssigneeRelation_assigneeId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstTaskAssigneeRelation"
    ADD CONSTRAINT "TstTaskAssigneeRelation_assigneeId_fkey" FOREIGN KEY ("assigneeId") REFERENCES public."TstUser"(id);


--
-- Name: TstTaskAssigneeRelation TstTaskAssigneeRelation_taskId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstTaskAssigneeRelation"
    ADD CONSTRAINT "TstTaskAssigneeRelation_taskId_fkey" FOREIGN KEY ("taskId") REFERENCES public."TstTask"(id);


--
-- Name: TstTask TstTask_caseProjectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_caseProjectId_fkey" FOREIGN KEY ("caseProjectId") REFERENCES public."TstProject"(id);


--
-- Name: TstTask TstTask_envId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_envId_fkey" FOREIGN KEY ("envId") REFERENCES public."TstEnv"(id);


--
-- Name: TstTask TstTask_planId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_planId_fkey" FOREIGN KEY ("planId") REFERENCES public."TstPlan"(id);


--
-- Name: TstTask TstTask_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- Name: TstTask TstTask_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstTask"
    ADD CONSTRAINT "TstTask_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstThread TstThread_authorId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstThread"
    ADD CONSTRAINT "TstThread_authorId_fkey" FOREIGN KEY ("authorId") REFERENCES public."TstUser"(id);


--
-- Name: TstThread TstThread_parentId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstThread"
    ADD CONSTRAINT "TstThread_parentId_fkey" FOREIGN KEY ("parentId") REFERENCES public."TstThread"(id);


--
-- Name: TstUserSettings TstUserSettings_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstUserSettings"
    ADD CONSTRAINT "TstUserSettings_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstUserVerifyCode TstUserVerifyCode_userId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstUserVerifyCode"
    ADD CONSTRAINT "TstUserVerifyCode_userId_fkey" FOREIGN KEY ("userId") REFERENCES public."TstUser"(id);


--
-- Name: TstUser TstUser_defaultOrgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstUser"
    ADD CONSTRAINT "TstUser_defaultOrgId_fkey" FOREIGN KEY ("defaultOrgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstUser TstUser_defaultPrjId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstUser"
    ADD CONSTRAINT "TstUser_defaultPrjId_fkey" FOREIGN KEY ("defaultPrjId") REFERENCES public."TstProject"(id);


--
-- Name: TstVer TstVer_orgId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstVer"
    ADD CONSTRAINT "TstVer_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES public."TstOrg"(id);


--
-- Name: TstVer TstVer_projectId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ngtesting
--

ALTER TABLE ONLY public."TstVer"
    ADD CONSTRAINT "TstVer_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES public."TstProject"(id);


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

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
-- Data for Name: CustomFieldDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."CustomFieldDefine" (id, "colCode", label, type, input, "textFormat", "applyTo", rows, required, readonly, "fullLine", ordr, descr, "createTime", "updateTime", disabled, deleted) FROM stdin;
1	prop01	严重级别	integer	dropdown	\N	issue	\N	f	f	f	1	\N	2018-11-09 12:06:02	\N	f	f
\.


--
-- Data for Name: CustomFieldInputTypeRelationDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
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
-- Data for Name: CustomFieldIputDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
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
-- Data for Name: CustomFieldOptionDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."CustomFieldOptionDefine" (id, label, descr, ordr, "defaultVal", "fieldId", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	阻塞	\N	1	f	1	f	f	2018-11-09 12:49:25	\N
2	紧急	\N	2	f	1	f	f	2018-11-09 12:49:28	\N
3	重要	\N	3	f	1	f	f	2018-11-09 12:49:31	\N
4	一般	\N	4	t	1	f	f	2018-11-09 12:49:33	\N
5	细微	\N	5	f	1	f	f	2018-11-09 12:49:36	\N
\.


--
-- Data for Name: CustomFieldTypeDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
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
-- Data for Name: IsuFieldCodeToTableDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
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
-- Data for Name: IsuFieldDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
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
-- Data for Name: IsuLinkReasonDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."IsuLinkReasonDefine" (id, label, value, disabled, deleted, "createTime", "updateTime") FROM stdin;
20	重复于	\N	f	f	2018-12-18 08:59:57	\N
30	阻塞	\N	f	f	2018-12-18 09:03:19	\N
40	阻塞于	\N	f	f	2018-12-18 09:00:19	\N
50	相关于	\N	f	f	2018-12-18 09:03:22	\N
10	重复	\N	t	f	2018-12-18 09:03:16	\N
\.


--
-- Data for Name: IsuNotificationDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."IsuNotificationDefine" (id, name, code, descr, disabled, deleted, "createTime", "updateTime") FROM stdin;
\.


--
-- Data for Name: IsuPriorityDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."IsuPriorityDefine" (id, label, value, descr, ordr, "defaultVal", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	紧急	urgent	\N	1	f	f	f	2018-11-09 11:28:35	\N
2	高	high	\N	2	f	f	f	2018-11-09 11:28:39	\N
3	中	medium	\N	3	t	f	f	2018-11-09 11:28:42	\N
4	低	low	\N	4	f	f	f	2018-11-09 11:28:45	\N
\.


--
-- Data for Name: IsuResolutionDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."IsuResolutionDefine" (id, label, value, "defaultVal", descr, ordr, disabled, deleted, "createTime", "updateTime") FROM stdin;
2	不是缺陷	not_defect	\N	\N	20	f	f	2018-11-23 15:26:22	\N
1	修复	fixed	t	\N	10	f	f	2018-11-23 15:25:52	\N
\.


--
-- Data for Name: IsuSeverityDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."IsuSeverityDefine" (id, label, value, descr, ordr, "defaultVal", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	阻塞	block	\N	1	f	f	f	2018-11-09 11:28:35	\N
2	紧急	critical	\N	2	f	f	f	2018-11-09 11:28:39	\N
3	重要	major	\N	3	f	f	f	2018-11-09 11:28:42	\N
4	一般	normal	\N	4	t	f	f	2018-11-09 11:42:21	\N
5	细微	minor	\N	5	f	f	f	2018-11-09 11:28:45	\N
\.


--
-- Data for Name: IsuStatusCategoryDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."IsuStatusCategoryDefine" (id, label, value, ordr, disabled, deleted, "finalVal") FROM stdin;
1	待办	todo	1	f	f	f
2	处理中	in_progress	2	f	f	f
3	完成	completed	3	f	f	t
\.


--
-- Data for Name: IsuStatusDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."IsuStatusDefine" (id, label, value, descr, "defaultVal", "finalVal", "categoryId", ordr, disabled, deleted, "createTime", "updateTime") FROM stdin;
1	打开	open	\N	t	f	1	1	f	f	2018-11-09 11:13:04	\N
2	解决	resolved	\N	f	f	2	2	f	f	2018-11-09 11:16:37	\N
3	关闭	closed	\N	f	t	3	3	f	f	2018-11-09 11:16:40	\N
4	重新打开	reopen	\N	f	f	1	4	f	f	2018-11-09 11:16:43	\N
5	挂起	suspend	\N	f	t	3	5	f	f	2018-11-09 11:16:46	\N
\.


--
-- Data for Name: IsuTypeDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."IsuTypeDefine" (id, value, label, descr, ordr, "defaultVal", disabled, deleted, "createTime", "updateTime") FROM stdin;
1	defect	缺陷	\N	1	t	f	f	2018-11-08 17:50:39	\N
2	todo	待办事项	\N	2	f	f	f	2018-11-08 17:54:24	\N
\.


--
-- Data for Name: IsuWorkflowStatusRelationDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."IsuWorkflowStatusRelationDefine" (id, "workflowId", "statusId") FROM stdin;
21	\N	1
22	\N	2
23	\N	3
24	\N	4
25	\N	5
\.


--
-- Data for Name: IsuWorkflowTransitionDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
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
-- Data for Name: TstCaseExeStatusDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."TstCaseExeStatusDefine" (id, value, label, descr, ordr, "finalVal", disabled, deleted) FROM stdin;
1	untest	未执行	\N	10	f	f	f
2	pass	成功	\N	20	t	f	f
3	fail	失败	\N	30	t	f	f
4	block	阻塞	\N	40	f	f	f
\.


--
-- Data for Name: TstCasePriorityDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."TstCasePriorityDefine" (id, value, label, descr, ordr, "defaultVal", disabled, deleted) FROM stdin;
1	high	高	\N	10	f	f	f
2	medium	中	\N	20	t	f	f
3	low	低	\N	30	f	f	f
\.


--
-- Data for Name: TstCaseTypeDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
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
-- Data for Name: TstOrgPrivilegeDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
--

COPY public."TstOrgPrivilegeDefine" (id, code, name, descr, disabled, deleted, "createTime", "updateTime") FROM stdin;
1	org-admin	管理组织	\N	f	f	2017-04-05 09:39:15	2017-04-05 09:39:20
2	site-admin	管理站点	\N	t	t	2017-04-05 09:39:15	2017-04-05 09:39:20
3	project-admin	管理项目	\N	f	f	2017-04-05 09:39:15	2017-04-05 09:39:20
\.


--
-- Data for Name: TstProjectPrivilegeDefine; Type: TABLE DATA; Schema: public; Owner: ngtesting
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
\.


--
-- Name: CustomFieldDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."CustomFieldDefine_id_seq"', 1, false);


--
-- Name: CustomFieldInputTypeRelationDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."CustomFieldInputTypeRelationDefine_id_seq"', 1, false);


--
-- Name: CustomFieldIputDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."CustomFieldIputDefine_id_seq"', 1, false);


--
-- Name: CustomFieldOptionDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."CustomFieldOptionDefine_id_seq"', 1, false);


--
-- Name: CustomFieldTypeDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."CustomFieldTypeDefine_id_seq"', 1, false);


--
-- Name: IsuFieldCodeToTableDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuFieldCodeToTableDefine_id_seq"', 1, false);


--
-- Name: IsuFieldDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuFieldDefine_id_seq"', 1, true);


--
-- Name: IsuLinkReasonDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuLinkReasonDefine_id_seq"', 1, false);


--
-- Name: IsuNotificationDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuNotificationDefine_id_seq"', 1, false);


--
-- Name: IsuPriorityDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuPriorityDefine_id_seq"', 1, false);


--
-- Name: IsuResolutionDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuResolutionDefine_id_seq"', 1, false);


--
-- Name: IsuSeverityDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuSeverityDefine_id_seq"', 1, false);


--
-- Name: IsuStatusCategoryDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuStatusCategoryDefine_id_seq"', 1, false);


--
-- Name: IsuStatusDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuStatusDefine_id_seq"', 1, false);


--
-- Name: IsuTypeDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuTypeDefine_id_seq"', 1, false);


--
-- Name: IsuWorkflowStatusRelationDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuWorkflowStatusRelationDefine_id_seq"', 1, false);


--
-- Name: IsuWorkflowTransitionDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."IsuWorkflowTransitionDefine_id_seq"', 1, false);


--
-- Name: TstCaseExeStatus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."TstCaseExeStatus_id_seq"', 542, true);


--
-- Name: TstCasePriority_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."TstCasePriority_id_seq"', 391, true);


--
-- Name: TstCaseType_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."TstCaseType_id_seq"', 913, true);


--
-- Name: TstOrgPrivilegeDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."TstOrgPrivilegeDefine_id_seq"', 1, false);


--
-- Name: TstProjectPrivilegeDefine_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ngtesting
--

SELECT pg_catalog.setval('public."TstProjectPrivilegeDefine_id_seq"', 1, false);


--
-- PostgreSQL database dump complete
--

