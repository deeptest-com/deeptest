package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.dao.TstCustomFieldProjectRelationDao;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.model.TstCustomFieldProjectRelation;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.service.TestCustomFieldProjectRelationService;
import com.ngtesting.platform.service.ProjectService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class TestCustomFieldProjectRelationServiceImpl extends BaseServiceImpl implements TestCustomFieldProjectRelationService {
    @Autowired
    TstCustomFieldProjectRelationDao customFieldProjectRelationDao;
    @Autowired
    ProjectDao projectDao;
    @Autowired
    ProjectService projectService;

    @Override
    public List<TstCustomFieldProjectRelation> listRelationsByField(Integer orgId, Integer fieldId) {
        List<TstProject> allProjects = projectDao.query(orgId, null, null);

        List<TstCustomFieldProjectRelation> relations;
        if (fieldId == null) {
            relations = new LinkedList<>();
        } else {
            relations = customFieldProjectRelationDao.query(orgId, fieldId);
        }

        List<TstCustomFieldProjectRelation> vos = new LinkedList<>();
        for (TstProject project : allProjects) {
            TstCustomFieldProjectRelation vo = genVo(orgId, project, fieldId);

            vo.setSelected(false);
            vo.setSelecting(false);
            for (TstCustomFieldProjectRelation po : relations) {
                if (po.getProjectId().longValue() == project.getId().longValue()
                        && po.getCustomFieldId().longValue() == fieldId.longValue()) {
                    vo.setSelected(true);
                    vo.setSelecting(true);
                }
            }
            vos.add(vo);

            for (TstProject child : project.getChildren()) {
                TstCustomFieldProjectRelation childVo = genVo(orgId, child, fieldId);

                childVo.setSelected(false);
                childVo.setSelecting(false);
                for (TstCustomFieldProjectRelation po : relations) {
                    if (po.getProjectId().longValue() == child.getId().longValue()
                            && po.getCustomFieldId().longValue() == fieldId.longValue()) {
                        childVo.setSelected(true);
                        childVo.setSelecting(true);
                    }
                }
                vos.add(childVo);
            }
        }

        return vos;
    }

    @Override
    public boolean saveRelationsByField(Integer orgId, Integer fieldId, List<TstCustomFieldProjectRelation> relations) {
        if (relations == null) {
            return false;
        }

        List<TstCustomFieldProjectRelation> selectedList = new LinkedList<>();
        for (Object obj: relations) {
            TstCustomFieldProjectRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstCustomFieldProjectRelation.class);
            if (vo.getSelecting()) {
                vo.setCustomFieldId(fieldId);
                selectedList.add(vo);
            }
        }

        customFieldProjectRelationDao.removeAllProjectsForField(orgId, fieldId);
        if (selectedList.size() > 0) {
            customFieldProjectRelationDao.saveRelations(selectedList);
        }

        return true;
    }

    @Override
    public TstCustomFieldProjectRelation genVo(Integer orgId, TstProject project, Integer fieldId) {
        TstCustomFieldProjectRelation vo = new TstCustomFieldProjectRelation();
        vo.setOrgId(orgId);
        vo.setCustomFieldId(fieldId);
        vo.setProjectId(project.getId());
        vo.setProjectName(project.getName());
        vo.setProjectType(project.getType().toString());

        return vo;
    }

}
