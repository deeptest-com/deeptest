package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseDao;
import com.ngtesting.platform.dao.CaseStepDao;
import com.ngtesting.platform.dao.TestSuiteDao;
import com.ngtesting.platform.dao.TestTaskDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseStep;
import com.ngtesting.platform.service.intf.CaseExportService;
import com.ngtesting.platform.service.intf.CaseHistoryService;
import com.ngtesting.platform.utils.FileUtil;
import org.apache.poi.ss.usermodel.*;
import org.apache.poi.ss.util.CellRangeAddress;
import org.apache.poi.xssf.usermodel.XSSFCellStyle;
import org.apache.poi.xssf.usermodel.XSSFFont;
import org.apache.poi.xssf.usermodel.XSSFWorkbook;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.FileOutputStream;
import java.io.OutputStream;
import java.util.*;

@Service
public class CaseExportServiceImpl extends BaseServiceImpl implements CaseExportService {

    @Autowired
    CaseDao caseDao;
    @Autowired
    CaseStepDao caseStepDao;

    @Autowired
    TestSuiteDao testSuiteDao;
    @Autowired
    TestTaskDao testTaskDao;

    public static List<String> ExtPropList;

    @Autowired
    CaseHistoryService caseHistoryService;

    @Override
    public String export(Integer projectId) {
        String fileName = UUID.randomUUID().toString() + ".xlsx";
        String fileDir = Constant.FTP_UPLOAD_DIR + "export/";
        String fileRelatPath = fileDir + fileName;
        String filePath = Constant.WORK_DIR + fileRelatPath;

        FileUtil.CreateDirIfNeeded(Constant.WORK_DIR + fileDir);

        XSSFWorkbook wb = new XSSFWorkbook();
        Sheet sheet = wb.createSheet();
        sheet.autoSizeColumn(1, true);
        sheet.setColumnWidth(0, 10 * 256);
        sheet.setColumnWidth(1, 50 * 256);
        sheet.setColumnWidth(2, 16 * 256);
        sheet.setColumnWidth(3, 16 * 256);
        sheet.setColumnWidth(4, 16 * 256);

        Integer topId = null;
        Integer rowCount = 0;

        XSSFCellStyle cellStyle = wb.createCellStyle();
        XSSFCellStyle stepStyle = wb.createCellStyle();

        Font fontStyle = wb.createFont();
//        fontStyle.setBold(true); // 加粗
        fontStyle.setFontName("黑体"); // 字体
        fontStyle.setFontHeightInPoints((short) 13); // 大小
        cellStyle.setFont(fontStyle);

        Font fontStyle2 = wb.createFont();
        ((XSSFFont) fontStyle2).setColor(IndexedColors.BLUE.getIndex());
        fontStyle2.setFontHeightInPoints((short) 13); // 大小
        stepStyle.setFont(fontStyle2);

//        cellStyle.setBorderBottom(BorderStyle.THIN);
//        cellStyle.setBorderLeft(BorderStyle.THIN);
//        cellStyle.setBorderRight(BorderStyle.THIN);
//        cellStyle.setBorderTop(BorderStyle.THIN);

        rowCount = writeHeader(sheet, rowCount, cellStyle);

        List<TstCase> ls = caseDao.queryCaseWithStepInfoByPrj(projectId);
        List<TstCase> pos = sortParentAndChild(ls);

        for (TstCase testCase : pos) {
            if (topId == null) {
                topId = testCase.getId();
            }
            rowCount = writeTestCase(testCase, sheet, rowCount, cellStyle, stepStyle);
        }

        try {
            OutputStream out = new FileOutputStream(filePath);
            wb.write(out);
        } catch (Exception ex) {
            ex.printStackTrace();
        }

        return fileRelatPath;
    }

    @Override
    public Integer writeHeader(Sheet sheet, Integer rowCount, XSSFCellStyle cellStyle) {
        Row titleRow = sheet.createRow(rowCount++);
        int cellCount = 0;
        Cell idCell = titleRow.createCell(cellCount++);
        Cell titleCell = titleRow.createCell(cellCount++);
        Cell typeCell = titleRow.createCell(cellCount++);
        Cell priorityCell = titleRow.createCell(cellCount++);
        Cell estimateCell = titleRow.createCell(cellCount++);
        Cell objectiveCell = titleRow.createCell(cellCount++);

        idCell.setCellValue("层级");
        titleCell.setCellValue("标题");
        typeCell.setCellValue("类型");
        priorityCell.setCellValue("优先级");
        estimateCell.setCellValue("耗时");
        objectiveCell.setCellValue("目的");

        idCell.setCellStyle(cellStyle);
        titleCell.setCellStyle(cellStyle);

        typeCell.setCellStyle(cellStyle);
        priorityCell.setCellStyle(cellStyle);
        estimateCell.setCellStyle(cellStyle);
        objectiveCell.setCellStyle(cellStyle);

        return rowCount;
    }

    @Override
    public Integer writeTestCase(TstCase testCase, Sheet sheet, Integer rowCount,
                                 XSSFCellStyle cellStyle, XSSFCellStyle stepStyle) {
        Integer ind = testCase.getLevel();
        XSSFCellStyle indentionStyle = (XSSFCellStyle)cellStyle.clone();
        indentionStyle.setIndention(ind.shortValue());

        stepStyle = (XSSFCellStyle)stepStyle.clone();
        stepStyle.setIndention(ind.shortValue());

        Row row = sheet.createRow(rowCount++);
        int cellCount = 0;
        Cell idCell = row.createCell(cellCount++);
        Cell titleCell = row.createCell(cellCount++);
        Cell typeCell = row.createCell(cellCount++);
        Cell priorityCell = row.createCell(cellCount++);
        Cell estimateCell = row.createCell(cellCount++);
        Cell objectiveCell = row.createCell(cellCount++);

        idCell.setCellValue(testCase.getLevel());
        titleCell.setCellValue(testCase.getName());

        if (testCase.getLeaf()) {
            typeCell.setCellValue(testCase.getTypeName());
            priorityCell.setCellValue(testCase.getPriorityName());
            estimateCell.setCellValue(testCase.getEstimate() == null ? "" : testCase.getEstimate().toString());
            objectiveCell.setCellValue(testCase.getObjective());
        }

        idCell.setCellStyle(cellStyle);
        titleCell.setCellStyle(indentionStyle);

        if (testCase.getLeaf()) {
            typeCell.setCellStyle(cellStyle);
            priorityCell.setCellStyle(cellStyle);
            estimateCell.setCellStyle(cellStyle);
            objectiveCell.setCellStyle(cellStyle);
        }

        if (testCase.getLeaf()) {
            for (TstCaseStep step : testCase.getSteps()) {
                sheet.addMergedRegion(new CellRangeAddress(rowCount, rowCount, 2, 5));

                Row stepRow = sheet.createRow(rowCount++);
                cellCount = 0;
                Cell ordrCell = stepRow.createCell(cellCount++);
                Cell optCell = stepRow.createCell(cellCount++);
                Cell resultCell = stepRow.createCell(cellCount++);

                ordrCell.setCellValue(step.getOrdr());

                optCell.setCellValue(step.getOpt());
                optCell.setCellStyle(stepStyle);

                resultCell.setCellValue(step.getExpect());
                resultCell.setCellStyle(stepStyle);
            }
        }

        return rowCount;
    }

    @Override
    public List<TstCase> sortParentAndChild(List<TstCase> entities) {
        // 1. 寻找集合中的所有根节点
        Map<Integer, List<TstCase>> pMap = new HashMap<>(); // 父节点为key
        Set<Integer> ids = new HashSet<Integer>(); // 存储节点id
        Set<Integer> pids = new HashSet<Integer>(); //存储父节点id

        for (TstCase entity : entities) {
            ids.add(entity.getId());
            Integer pid = entity.getpId();
            pids.add(pid);
            if(null == pMap.get(pid)) {
                pMap.put(pid, new ArrayList<>());
            }
            pMap.get(pid).add(entity);
        }
        pids.removeAll(ids); // 得到根节点

        // 2. 父子排序
        List<TstCase> sortedList= new ArrayList<>();
        for (Integer rootPid : pids) {
            List<TstCase> queue = pMap.remove(rootPid);
            setLevel(queue, 0);

            if (null != queue) {
                while(queue.size() > 0) {
                    TstCase entity = queue.remove(0);
                    sortedList.add(entity);
                    List<TstCase> tmpList = pMap.remove(entity.getId());
                    if (null != tmpList) {
                        setLevel(tmpList, entity.getLevel() + 1);
                        queue.addAll(0, tmpList); // 将子节点插在下一个兄弟节点前
                    }
                }
            }
        }
        return sortedList;
    }

    @Override
    public void setLevel(List<TstCase> entities, Integer level) {
        for (TstCase cs : entities) {
            cs.setLevel(level);
        }
    }

}

