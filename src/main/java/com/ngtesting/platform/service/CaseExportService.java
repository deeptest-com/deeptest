package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCase;
import org.apache.poi.ss.usermodel.Sheet;
import org.apache.poi.xssf.usermodel.XSSFCellStyle;

import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;

public interface CaseExportService extends BaseService {

    String export(Integer projectId);

	Integer writeHeader(Sheet sheet, Integer rowCount, XSSFCellStyle cellStyle);

	Integer writeTestCase(TstCase testCase, Sheet sheet, Integer topId, Integer rowCount,
                          AtomicInteger level, XSSFCellStyle cellStyle);

    List<TstCase> sortParentAndChild(List<TstCase> entities);

    void setLevel(List<TstCase> entities, Integer level);
}
