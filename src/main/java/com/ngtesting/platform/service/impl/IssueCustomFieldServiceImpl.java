package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstCustomField;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.service.IssueCustomFieldService;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class IssueCustomFieldServiceImpl extends BaseServiceImpl implements IssueCustomFieldService {

//    @Autowired
//    ProjectService projectService;
//    @Autowired
//    CustomFieldOptionService customFieldOptionService;

    @Override
    public List<TstCustomField> list(Integer orgId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCustomField.class);
//
//        dc.add(Restrictions.eq("orgId", orgId));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("ordr"));
//        List ls = findAllByCriteria(dc);
//
//        return ls;

        return null;
    }

    @Override
    public List<TstCustomField> listForCaseByOrg(Integer orgId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCustomField.class);
//
//        dc.add(Restrictions.eq("orgId", orgId));
//        dc.add(Restrictions.eq("applyTo", FieldApplyTo.test_case));
//
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("ordr"));
//        List ls = findAllByCriteria(dc);
//
//        return ls;

        return null;
    }

    @Override
    public List<TstCustomField> listForCaseByProject(Integer orgId, Integer projectId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCustomField.class);
//
//        dc.createAlias("projectSet", "p").add(Restrictions.eq("p.id", projectId));
//        dc.add(Restrictions.eq("applyTo", FieldApplyTo.test_case));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.addOrder(Order.asc("ordr"));
//        List<TstCustomField> ls1 = findAllByCriteria(dc);
//
//        DetachedCriteria dc2 = DetachedCriteria.forClass(TstCustomField.class);
//        dc2.add(Restrictions.eq("orgId", orgId));
//        dc2.add(Restrictions.eq("global", true));
//        dc2.add(Restrictions.eq("applyTo", FieldApplyTo.test_case));
//        dc2.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc2.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc2.addOrder(Order.asc("ordr"));
//        List<TstCustomField> ls2 = findAllByCriteria(dc2);
//
//        ls2.addAll(ls1);
//        List<TstCustomField> vos = genVos(ls2);
//
//        return vos;

        return null;
    }

    @Override
    public List<TstCustomField> listVos(Integer orgId) {
        List<TstCustomField> ls = list(orgId);

        List<TstCustomField> vos = genVos(ls);
        return vos;
    }

    @Override
    public TstCustomField save(TstCustomField vo, Integer orgId) {
//        if (vo == null) {
//            return null;
//        }
//
//        TstCustomField po;
//        if (vo.getId() != null) {
//            po = (TstCustomField) getDetail(TstCustomField.class, vo.getId());
//        } else {
//            po = new TstCustomField();
//        }
//        this.initPo(po, vo);
//
//        po.setApplyTo(TstCustomField.FieldApplyTo.valueOf(vo.getApplyTo()));
//        po.setType(TstCustomField.FieldType.valueOf(vo.getType()));
//
//        if (StringUtil.isNotEmpty(vo.getFormat())) {
//            po.setFormat(TstCustomField.FieldFormat.valueOf(vo.getFormat()));
//        }
//
//        po.setOrgId(orgId);
//
//        if (vo.getId() == null) {
//            String hql = "select max(ordr) from TstCustomField";
//            Integer maxOrder = (Integer) getByHQL(hql);
//            if (maxOrder == null) {
//                maxOrder = 0;
//            }
//            po.setOrdr(maxOrder + 10);
//        }
//        if (!po.getType().equals(FieldType.text)) {
//            po.setRows(0);
//            po.setFormat(null);
//        }
//        if (po.getGlobal() && po.getProjectSet().size() > 0) {
//            po.setProjectSet(new HashSet<TstProject>(0));
//        }
//
//        saveOrUpdate(po);
//        return po;

        return null;
    }

    @Override
    public boolean delete(Integer id) {
//        TstCustomField po = (TstCustomField) getDetail(TstCustomField.class, id);
//        po.getProjectSet().clear();
//        getDao().delete(po);

        return true;
    }

    @Override
    public List<String> listApplyTo() {
        List<String> ls = new LinkedList<String>();
//        for (FieldApplyTo item : TstCustomField.FieldApplyTo.values()) {
//            ls.add(item.toString());
//        }
        return ls;
    }

    @Override
    public List<String> listType() {
        List<String> ls = new LinkedList<String>();
//        for (FieldType item : TstCustomField.FieldType.values()) {
//            ls.add(item.toString());
//        }
        return ls;
    }

    @Override
    public List<String> listFormat() {
        List<String> ls = new LinkedList<String>();
//        for (FieldFormat item : TstCustomField.FieldFormat.values()) {
//            ls.add(item.toString());
//        }
        return ls;
    }

    @Override
    public boolean changeOrderPers(Integer id, String act) {
//        TstCustomField type = (TstCustomField) getDetail(TstCustomField.class, id);
//
//        String hql = "from TstCustomField tp where tp.deleted = false and tp.disabled = false ";
//        if ("up".equals(act)) {
//            hql += "and tp.ordr < ? order by ordr desc";
//        } else if ("down".equals(act)) {
//            hql += "and tp.ordr > ? order by ordr asc";
//        } else {
//            return false;
//        }
//
//        TstCustomField neighbor = (TstCustomField) getDao().findFirstByHQL(hql, type.getOrdr());
//
//        Integer order = type.getOrdr();
//        type.setOrdr(neighbor.getOrdr());
//        neighbor.setOrdr(order);
//
//        saveOrUpdate(type);
//        saveOrUpdate(neighbor);

        return true;
    }

    @Override
    public List<TstProject> listProjectsForField(Integer orgId, Integer fieldId) {
//        List<TstProject> allProjects = projectService.list(orgId, null, null);
//
//        Set<TstProject> projectsForField;
//        if (fieldId == null) {
//            projectsForField = new HashSet<TstProject>();
//        } else {
//            TstCustomField field = (TstCustomField) getDetail(TstCustomField.class, fieldId);
//            projectsForField = field.getProjectSet();
//        }
//
//        List<TstProject> vos = new LinkedList<TstProject>();
//        for (TstProject po1 : allProjects) {
//            TstProject vo = projectService.genVo(po1, null);
//
//            vo.setSelected(false);
//            vo.setSelecting(false);
//            for (TstProject item : projectsForField) {
//                if (po1.getId().longValue() == item.getId().longValue()) {
//                    vo.setSelected(true);
//                    vo.setSelecting(true);
//                }
//            }
//            vos.add(vo);
//
//            for (TstProject child : po1.getChildren()) {
//                TstProject childVo = projectService.genVo(child, null);
//
//                for (TstProject item : projectsForField) {
//                    if (child.getId().longValue() == item.getId().longValue()) {
//                        childVo.setSelected(true);
//                        childVo.setSelecting(true);
//                    }
//                }
//
//                vos.add(childVo);
//            }
//        }
//
//        return vos;

        return null;
    }

    @Override
    public boolean saveRelationsProjects(Integer fieldId, List<TstProject> projects) {
//        if (projects == null) {
//            return false;
//        }
//
//        TstCustomField field = (TstCustomField) getDetail(TstCustomField.class, fieldId);
//        Set<TstProject> projectSet = field.getProjectSet();
//
//        for (Object obj : projects) {
//            TstProject vo = JSON.parseObject(JSON.toJSONString(obj), TstProject.class);
//            if (vo.getSelecting() != vo.getSelected()) { // 变化了
//                TstProject project = (TstProject) getDetail(TstProject.class, vo.getId());
//
//                if (vo.getSelecting() && !projectSet.contains(project)) { // 勾选
//                    projectSet.add(project);
//                } else if (project != null) { // 取消
//                    projectSet.remove(project);
//                }
//            }
//        }
//        saveOrUpdate(field);

        return true;
    }

    @Override
    public String getLastUnusedColumn(Integer orgId) {
//        String hql = "select cf.myColumn from TstCustomField cf where cf.deleted = false and cf.disabled = false " +
//                "and cf.orgId = ? order by cf.myColumn asc";
//
//        String ret = null;
//        List<String> ls = getDao().getListByHQL(hql, orgId);
//        for (int i = 1; i <= 20; i++) {
//            String prop = "prop" + String.format("%02d", i);
//            if (!ls.contains(prop)) {
//                ret = prop;
//                break;
//            }
//        }
//
//        return ret;

        return null;
    }

    @Override
    public TstCustomField get(Integer customFieldId) {
        return null;
    }

    @Override
    public List<TstCustomField> genVos(List<TstCustomField> pos) {
        List<TstCustomField> vos = new LinkedList<TstCustomField>();

        for (TstCustomField po : pos) {
            TstCustomField vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }
    @Override
    public TstCustomField genVo(TstCustomField po) {
//        if (po == null) {
//            return null;
//        }
//        TstCustomField vo = new TstCustomField();
//        BeanUtilEx.copyProperties(vo, po);
//
//        vo.setOptionVos(this.customFieldOptionService.genVos(po.getOptions()));
//
//        return vo;

        return null;
    }

    @Override
    public void initPo(TstCustomField po, TstCustomField vo) {
        po.setCode(vo.getCode());
        po.setMyColumn(vo.getMyColumn());
        po.setLabel(vo.getLabel());
        po.setDescr(vo.getDescr());
        po.setRows(vo.getRows());
        po.setGlobal(vo.getGlobal());
        po.setRequired(vo.getRequired());
    }

}
