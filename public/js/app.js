$(function() {


    function updateTable(services) {
        if (!services) {
            return;
        }
        services.sort(function (a, b) {
            if (a.status == b.status) {
                return a.name > b.name;
            }
            return getStatusPriority(a.status) < getStatusPriority(b.status);
        });
        var table = $('#services-table');
        var tbody = $('<tbody></tbody>');
        for (var i in services) {
            var service = services[i];
            var row = $('<tr data-id="' + service.id + '"></tr>');
            row.append('<td>' + service.name + '</td>');
            row.append('<td>' + service.protocol + '</td>');
            row.append('<td>' + service.host + '</td>');
            row.append('<td>' + (service.port ? service.port : 'N/A') + '</td>');
            row.append('<td class="' + getStatusTextClass(service.status) + '">' + service.status + '</td>');1
            row.append('<td></td>');
            row.append('<td></td>');
            row.append('<td><button class="btn btn-sm btn-warning edit-service">Edit</button> <button class="btn btn-sm btn-danger delete-service">Delete</button></td>');
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

    function getStatusPriority(status) {
        switch(status) {
            case 'ONLINE':
                return 0;
            case 'PENDING':
                return 1;
            case 'OFFLINE':
                return 2;
            default:
                return 0;
        }
    }

    $('select[name=protocol]').on('change', function() {
        var port = $(this).closest('form').find('input[name=port]').closest('.form-group');
        if ($(this).val() != 'http' && $(this).val() != 'https' && $(this).val() != 'icmp') {
            port.show();
        } else {
            port.hide();
        }
    });

    $('#add-service-button').click(function() {
        $('#add-service-modal input').val('');
        $('#add-service-modal').modal();
    });

    $('#services-table').on('click', '.edit-service', function() {
        var id = $(this).closest('tr').data('id');
        $.ajax({
            url: '/services/' + id,
            method: 'GET',
            success: function(service) {
                var port = $('#edit-service-modal').find('input[name=port]').closest('.form-group');
                if (service.protocol != 'http' && service.protocol != 'https' && service.protocol != 'icmp') {
                    port.show();
                } else {
                    port.hide();
                }

                $('#edit-service-modal input[name=id]').val(service.id);
                $('#edit-service-modal input[name=name]').val(service.name);
                $('#edit-service-modal select[name=protocol]').val(service.protocol);
                $('#edit-service-modal input[name=host]').val(service.host);
                $('#edit-service-modal input[name=port]').val(service.port);
                $('#edit-service-modal').modal();
            }
        });
    });

    $('#services-table').on('click', '.delete-service', function() {
        var serviceName = $(this).closest('tr').find('td:first-child').text();
        var id = $(this).closest('tr').data('id');
        swal({
            title: 'Are you sure?',
            text: 'Do you really want to delete the service "' + serviceName + '"?',
            type: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#DD6B55',
            confirmButtonText: 'Delete',
            closeOnConfirm: false
        },
        function(){
            $.ajax({
                url: '/services/' + id,
                method: 'DELETE',
                success: function() {
                    $('#services-table tr[data-id=1]').remove()
                    swal('Deleted', 'The service "' + serviceName + '" has been deleted', 'success');
                }
            });
        });
    });

    $('#add-service-form').submit(function(e) {
        e.preventDefault();
        var service = $(this).serialize();
        $.modal.close();
        var name = $(this).find('input[name=name]').val();
         $.ajax({
            url: '/services',
            method: 'POST',
            data: service,
            success: function() {
                swal({
                    title: 'Success',
                    text: 'The service "' + name + '" has been added successfully',
                    type: 'success'
                });
                getServices();     
            }
        });
    });


    $('#edit-service-form').submit(function(e) {
        e.preventDefault();
        var service = $(this).serialize();
        var id = $(this).find('input[name=id]').val();
        var name = $(this).find('input[name=name]').val();
        $.modal.close();
         $.ajax({
            url: '/services/' + id,
            method: 'PUT',
            data: service,
            success: function() {
                console.log('test');
                swal({
                    title: 'Success',
                    text: 'The service "' + name + '" has been updated successfully',
                    type: 'success'
                });
                getServices();     
            }
        });
    });

    getServices();
    setInterval(getServices, 2000);

});