package cn.linkr.testspace.dao;

import java.util.List;

import cn.linkr.testspace.entity.TestProject;

public interface ProjectDao {
	 public List findProjectByProcedure(Long companyId, Boolean isActive);
	 public TestProject moveProject(Long projectId, Long newParentId);
}
