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
SET default_tablespace = '';

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

