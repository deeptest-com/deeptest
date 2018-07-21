package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCustomField;
import com.ngtesting.platform.model.TstProject;

import java.util.List;

public interface CustomFieldService extends BaseService {
	List<TstCustomField> list(Integer orgId);
	List<TstCustomField> listForCaseByOrg(Integer orgId);
    List<TstCustomField> listForCaseByProject(Integer orgId, Integer projectId);

	TstCustomField save(TstCustomField vo, Integer orgId);
	boolean delete(Integer id);

	List<TstCustomField> genVos(List<TstCustomField> pos);
	TstCustomField genVo(TstCustomField po);

	List<TstCustomField> listVos(Integer orgId);

	List<String> listApplyTo();

	List<String> listType();

	List<String> listFormat();

	boolean changeOrderPers(Integer id, String act);

	List<TstProject> listProjectsForField(Integer orgId, Integer fieldId);

	boolean saveRelationsProjects(Integer id, List<TstProject> projects);

	void initPo(TstCustomField po, TstCustomField vo);

    String getLastUnusedColumn(Integer orgId);
}
