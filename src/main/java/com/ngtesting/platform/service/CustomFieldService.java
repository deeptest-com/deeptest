package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCustomField;
import com.ngtesting.platform.model.TstProject;

import java.util.List;

public interface CustomFieldService extends BaseService {
	List<TstCustomField> list(Long orgId);
	List<TstCustomField> listForCaseByOrg(Long orgId);
    List<TstCustomField> listForCaseByProject(Long orgId, Long projectId);

	TstCustomField save(TstCustomField vo, Long orgId);
	boolean delete(Long id);

	List<TstCustomField> genVos(List<TstCustomField> pos);
	TstCustomField genVo(TstCustomField po);

	List<TstCustomField> listVos(Long orgId);

	List<String> listApplyTo();

	List<String> listType();

	List<String> listFormat();

	boolean changeOrderPers(Long id, String act);

	List<TstProject> listProjectsForField(Long orgId, Long fieldId);

	boolean saveRelationsProjects(Long id, List<TstProject> projects);

	void initPo(TstCustomField po, TstCustomField vo);

    String getLastUnusedColumn(Long orgId);
}
