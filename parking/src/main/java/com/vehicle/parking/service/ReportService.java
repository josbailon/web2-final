package com.vehicle.parking.service;

import com.itextpdf.text.*;
import com.itextpdf.text.pdf.PdfPCell;
import com.itextpdf.text.pdf.PdfPTable;
import com.itextpdf.text.pdf.PdfWriter;
import com.vehicle.parking.model.Empleado;
import com.vehicle.parking.model.Reserva;
import com.vehicle.parking.repository.EmpleadoRepository;
import com.vehicle.parking.repository.ReservaRepository;
import org.apache.poi.ss.usermodel.*;
import org.apache.poi.xssf.usermodel.XSSFWorkbook;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.List;

@Service
public class ReportService {

    @Autowired
    private EmpleadoRepository empleadoRepository;

    @Autowired
    private ReservaRepository reservaRepository;

    public void generateEmpleadoPdfReport(HttpServletResponse response) throws DocumentException, IOException {
        List<Empleado> empleados = empleadoRepository.findAll();

        Document document = new Document();
        PdfWriter.getInstance(document, response.getOutputStream());

        document.open();
        document.add(new Paragraph("Reporte de Empleados"));

        PdfPTable table = new PdfPTable(6);
        table.setWidthPercentage(100);
        addTableHeader(table, new String[]{"ID", "Nombre", "Apellido", "Fecha Ingreso", "Email", "Salario", "Estado"});

        for (Empleado empleado : empleados) {
            addTableRow(table, new String[]{
                empleado.getId().toString(),
                empleado.getNombre(),
                empleado.getApellido(),
                empleado.getFechaIngreso().toString(),
                empleado.getEmail(),
                empleado.getSalario().toString(),
                empleado.getEstado().toString()
            });
        }

        document.add(table);
        document.close();
    }

    public void generateReservaPdfReport(HttpServletResponse response) throws DocumentException, IOException {
        List<Reserva> reservas = reservaRepository.findAll();

        Document document = new Document();
        PdfWriter.getInstance(document, response.getOutputStream());

        document.open();
        document.add(new Paragraph("Reporte de Reservas"));

        PdfPTable table = new PdfPTable(8);
        table.setWidthPercentage(100);
        addTableHeader(table, new String[]{"ID", "ID Cliente", "ID Vehiculo", "Fecha Reserva", "Hora Entrada", "Hora Salida", "Estado"});

        for (Reserva reserva : reservas) {
            addTableRow(table, new String[]{
                reserva.getId().toString(),
                reserva.getIdCliente().toString(),
                reserva.getIdVehiculo().toString(),
                reserva.getFechaReserva().toString(),
                reserva.getHoraEntrada().toString(),
                reserva.getHoraSalida().toString(),
                reserva.getEstadoReserva().toString()
            });
        }

        document.add(table);
        document.close();
    }

    public void generateEmpleadoExcelReport(HttpServletResponse response) throws IOException {
        List<Empleado> empleados = empleadoRepository.findAll();

        Workbook workbook = new XSSFWorkbook();
        Sheet sheet = workbook.createSheet("Empleados");

        int rowNum = 0;
        Row headerRow = sheet.createRow(rowNum++);
        addExcelHeader(headerRow, new String[]{"ID", "Nombre", "Apellido", "Fecha Ingreso", "Email", "Salario", "Estado"});

        for (Empleado empleado : empleados) {
            Row row = sheet.createRow(rowNum++);
            addExcelRow(row, new String[]{
                empleado.getId().toString(),
                empleado.getNombre(),
                empleado.getApellido(),
                empleado.getFechaIngreso().toString(),
                empleado.getEmail(),
                empleado.getSalario().toString(),
                empleado.getEstado().toString()
            });
        }

        response.setHeader("Content-Disposition", "attachment;filename=empleados_report.xlsx");
        workbook.write(response.getOutputStream());
        workbook.close();
    }

    public void generateReservaExcelReport(HttpServletResponse response) throws IOException {
        List<Reserva> reservas = reservaRepository.findAll();

        Workbook workbook = new XSSFWorkbook();
        Sheet sheet = workbook.createSheet("Reservas");

        int rowNum = 0;
        Row headerRow = sheet.createRow(rowNum++);
        addExcelHeader(headerRow, new String[]{"ID", "ID Cliente", "ID Vehiculo", "Fecha Reserva", "Hora Entrada", "Hora Salida", "Estado"});

        for (Reserva reserva : reservas) {
            Row row = sheet.createRow(rowNum++);
            addExcelRow(row, new String[]{
                reserva.getId().toString(),
                reserva.getIdCliente().toString(),
                reserva.getIdVehiculo().toString(),
                reserva.getFechaReserva().toString(),
                reserva.getHoraEntrada().toString(),
                reserva.getHoraSalida().toString(),
                reserva.getEstadoReserva().toString()
            });
        }

        response.setHeader("Content-Disposition", "attachment;filename=reservas_report.xlsx");
        workbook.write(response.getOutputStream());
        workbook.close();
    }

    private void addTableHeader(PdfPTable table, String[] headers) {
        for (String header : headers) {
            PdfPCell cell = new PdfPCell();
            cell.setPhrase(new Phrase(header));
            table.addCell(cell);
        }
    }

    private void addTableRow(PdfPTable table, String[] values) {
        for (String value : values) {
            table.addCell(new PdfPCell(new Phrase(value)));
        }
    }

    private void addExcelHeader(Row row, String[] headers) {
        int colNum = 0;
        for (String header : headers) {
            Cell cell = row.createCell(colNum++);
            cell.setCellValue(header);
        }
    }

    private void addExcelRow(Row row, String[] values) {
        int colNum = 0;
        for (String value : values) {
            Cell cell = row.createCell(colNum++);
            cell.setCellValue(value);
        }
    }
}
