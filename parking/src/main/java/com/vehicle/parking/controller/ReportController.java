package com.vehicle.parking.controller;

import com.itextpdf.text.DocumentException;
import com.vehicle.parking.service.ReportService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

@RestController
@RequestMapping("/api/reports")
public class ReportController {

    @Autowired
    private ReportService reportService;

    @GetMapping("/empleados/pdf")
    public void generateEmpleadoPdfReport(HttpServletResponse response) throws IOException, DocumentException {
        reportService.generateEmpleadoPdfReport(response);
    }

    @GetMapping("/reservas/pdf")
    public void generateReservaPdfReport(HttpServletResponse response) throws IOException, DocumentException {
        reportService.generateReservaPdfReport(response);
    }

    @GetMapping("/empleados/excel")
    public void generateEmpleadoExcelReport(HttpServletResponse response) throws IOException {
        reportService.generateEmpleadoExcelReport(response);
    }

    @GetMapping("/reservas/excel")
    public void generateReservaExcelReport(HttpServletResponse response) throws IOException {
        reportService.generateReservaExcelReport(response);
    }
}
