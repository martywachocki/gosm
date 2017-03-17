$(function() {


    function updateTable(services) {
        var table = $('#services-table');
        var tbody = $('<tbody></tbody>')
        for (var i in services) {
            var service = services[i];
            var row = $('<tr></tr>');
            row.append('<td>' + service.name + '</td>');
            row.append('<td>' + service.protocol + '</td>');
            row.append('<td>' + service.host + '</td>');
            row.append('<td>' + (service.port != 0 ? service.port : 'N/A') + '</td>');
            row.append('<td class="' + getStatusTextClass(service.status) + '">' + service.status + '</td>');1
            row.append('<td></td>');
            row.append('<td></td>');
            row.append('<td><button class="btn btn-sm btn-warning">Edit</button> <button class="btn btn-sm btn-danger">Delete</button></td>');
            tbody.append(row);
        }
        table.find('tbody').replaceWith(tbody);    
    }
    
    function getServices() {
        $.get('/services', function(services) {
            updateTable(services);
        });
    }

    function getStatusTextClass(status) {
        switch(status) {
            case 'ONLINE':
                return 'text-success';
            case 'PENDING':
                return 'text-warning';
            case 'OFFLINE':
                return 'text-danger';
            default:
                return '';
        }
    }

    getServices();
    setInterval(getServices, 2000);


});