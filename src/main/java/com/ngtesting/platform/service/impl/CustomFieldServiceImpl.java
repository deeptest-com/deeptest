package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.TestCustomField;
import com.ngtesting.platform.entity.TestCustomField.FieldApplyTo;
import com.ngtesting.platform.entity.TestCustomField.FieldFormat;
import com.ngtesting.platform.entity.TestCustomField.FieldType;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.service.CustomFieldOptionService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.TestProjectVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Set;

@Service
public class CustomFieldServiceImpl extends BaseServiceImpl implements CustomFieldService {

    @Autowired
    ProjectService projectService;
    @Autowired
    CustomFieldOptionService customFieldOptionService;

    @Override
    public List<TestCustomField> list(Long orgId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCustomField.class);

        dc.add(Restrictions.eq("orgId", orgId));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));

        dc.addOrder(Order.asc("ordr"));
        List ls = findAllByCriteria(dc);

        return ls;
    }

    @Override
    public List<TestCustomField> listForCaseByOrg(Long orgId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCustomField.class);

        dc.add(Restrictions.eq("orgId", orgId));
        dc.add(Restrictions.eq("applyTo", FieldApplyTo.test_case));

        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));

        dc.addOrder(Order.asc("ordr"));
        List ls = findAllByCriteria(dc);

        return ls;
    }

    @Override
    public List<CustomFieldVo> listForCaseByProject(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCustomField.class);

        dc.createAlias("projectSet", "p").add(Restrictions.eq("p.id", projectId));
        dc.add(Restrictions.eq("applyTo", FieldApplyTo.test_case));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.addOrder(Order.asc("ordr"));
        List<TestCustomField> ls1 = findAllByCriteria(dc);

        DetachedCriteria dc2 = DetachedCriteria.forClass(TestCustomField.class);
        dc2.add(Restrictions.eq("global", true));
        dc2.add(Restrictions.eq("applyTo", FieldApplyTo.test_case));
        dc2.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc2.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc2.addOrder(Order.asc("ordr"));
        List<TestCustomField> ls2 = findAllByCriteria(dc2);

        ls2.addAll(ls1);
        List<CustomFieldVo> vos = genVos(ls2);

        return vos;
    }

    @Override
    public List<CustomFieldVo> listVos(Long orgId) {
        List<TestCustomField> ls = list(orgId);

        List<CustomFieldVo> vos = genVos(ls);
        return vos;
    }

    @Override
    public TestCustomField save(CustomFieldVo vo, Long orgId) {
        if (vo == null) {
            return null;
        }

        TestCustomField po;
        if (vo.getId() != null) {
            po = (TestCustomField) get(TestCustomField.class, vo.getId());
        } else {
            po = new TestCustomField();
        }
        this.initPo(po, vo);

        po.setApplyTo(TestCustomField.FieldApplyTo.valueOf(vo.getApplyTo()));
        po.setType(TestCustomField.FieldType.valueOf(vo.getType()));

        if (StringUtil.isNotEmpty(vo.getFormat())) {
            po.setFormat(TestCustomField.FieldFormat.valueOf(vo.getFormat()));
        }

        po.setOrgId(orgId);

        if (vo.getId() == null) {
            String hql = "select max(ordr) from TestCustomField";
            Integer maxOrder = (Integer) getByHQL(hql);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            po.setOrdr(maxOrder + 10);
        }
        if (!po.getType().equals(FieldType.text)) {
            po.setRows(0);
            po.setFormat(null);
        }
        if (po.getGlobal() && po.getProjectSet().size() > 0) {
            po.setProjectSet(new HashSet<TestProject>(0));
        }

        saveOrUpdate(po);
        return po;
    }

    @Override
    public boolean delete(Long id) {
        TestCustomField po = (TestCustomField) get(TestCustomField.class, id);
        po.getProjectSet().clear();
        getDao().delete(po);

        return true;
    }

    @Override
    public List<String> listApplyTo() {
        List<String> ls = new LinkedList<String>();
        for (FieldApplyTo item : TestCustomField.FieldApplyTo.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

    @Override
    public List<String> listType() {
        List<String> ls = new LinkedList<String>();
        for (FieldType item : TestCustomField.FieldType.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

    @Override
    public List<String> listFormat() {
        List<String> ls = new LinkedList<String>();
        for (FieldFormat item : TestCustomField.FieldFormat.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

    @Override
    public boolean changeOrderPers(Long id, String act) {
        TestCustomField type = (TestCustomField) get(TestCustomField.class, id);

        String hql = "from TestCustomField tp where tp.deleted = false and tp.disabled = false ";
        if ("up".equals(act)) {
            hql += "and tp.ordr < ? order by ordr desc";
        } else if ("down".equals(act)) {
            hql += "and tp.ordr > ? order by ordr asc";
        } else {
            return false;
        }

        TestCustomField neighbor = (TestCustomField) getDao().findFirstByHQL(hql, type.getOrdr());

        Integer order = type.getOrdr();
        type.setOrdr(neighbor.getOrdr());
        neighbor.setOrdr(order);

        saveOrUpdate(type);
        saveOrUpdate(neighbor);

        return true;
    }

    @Override
    public List<TestProjectVo> listProjectsForField(Long orgId, Long fieldId) {
        List<TestProject> allProjects = projectService.list(orgId, null, null);

        Set<TestProject> projectsForField;
        if (fieldId == null) {
            projectsForField = new HashSet<TestProject>();
        } else {
            TestCustomField field = (TestCustomField) get(TestCustomField.class, fieldId);
            projectsForField = field.getProjectSet();
        }

        List<TestProjectVo> vos = new LinkedList<TestProjectVo>();
        for (TestProject po1 : allProjects) {
            TestProjectVo vo = projectService.genVo(po1);

            vo.setSelected(false);
            vo.setSelecting(false);
            for (TestProject item : projectsForField) {
                if (po1.getId().longValue() == item.getId().longValue()) {
                    vo.setSelected(true);
                    vo.setSelecting(true);
                }
            }
            vos.add(vo);

            for (TestProject child : po1.getChildren()) {
                TestProjectVo childVo = projectService.genVo(child);

                for (TestProject item : projectsForField) {
                    if (child.getId().longValue() == item.getId().longValue()) {
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
    public boolean saveRelationsProjects(Long fieldId, List<TestProjectVo> projects) {
        if (projects == null) {
            return false;
        }

        TestCustomField field = (TestCustomField) get(TestCustomField.class, fieldId);
        Set<TestProject> projectSet = field.getProjectSet();

        for (Object obj : projects) {
            TestProjectVo vo = JSON.parseObject(JSON.toJSONString(obj), TestProjectVo.class);
            if (vo.getSelecting() != vo.getSelected()) { // 变化了
                TestProject project = (TestProject) get(TestProject.class, vo.getId());

                if (vo.getSelecting() && !projectSet.contains(project)) { // 勾选
                    projectSet.add(project);
                } else if (project != null) { // 取消
                    projectSet.remove(project);
                }
            }
        }
        saveOrUpdate(field);

        return true;
    }

    @Override
    public String getLastUnusedColumn(Long orgId) {
        String hql = "select cf.myColumn from TestCustomField cf where cf.deleted = false and cf.disabled = false " +
                "and cf.orgId = ? order by cf.myColumn asc";

        String ret = null;
        List<String> ls = getDao().getListByHQL(hql, orgId);
        for (int i = 1; i <= 20; i++) {
            String prop = "prop" + String.format("%02d", i);
            if (!ls.contains(prop)) {
                ret = prop;
                break;
            }
        }

        return ret;
    }

    @Override
    public List<CustomFieldVo> genVos(List<TestCustomField> pos) {
        List<CustomFieldVo> vos = new LinkedList<CustomFieldVo>();

        for (TestCustomField po : pos) {
            CustomFieldVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }
    @Override
    public CustomFieldVo genVo(TestCustomField po) {
        if (po == null) {
            return null;
        }
        CustomFieldVo vo = new CustomFieldVo();
        BeanUtilEx.copyProperties(vo, po);

        vo.setOptionVos(this.customFieldOptionService.genVos(po.getOptions()));

        return vo;
    }

    @Override
    public void initPo(TestCustomField po, CustomFieldVo vo) {
        po.setCode(vo.getCode());
        po.setMyColumn(vo.getMyColumn());
        po.setLabel(vo.getLabel());
        po.setDescr(vo.getDescr());
        po.setRows(vo.getRows());
        po.setGlobal(vo.getGlobal());
        po.setRequired(vo.getRequired());
    }

}
