$(document).ready(function () {

    // ===========================================
    // Big target
    // ===========================================
    $('.bigTarget').on('click', function (e) {
        var $link = $('a[href]:first', this);
        if ($link.length) {
            e.preventDefault();
            document.location.href = $link.prop('href');
        }
    });
    $('.bigTarget a').on('click', function (e) {
        e.stopPropagation();
    });

    $("#existingcustomer").click(function () {
        $(".newcustomer").removeClass('open');
        $(".existingcustomer").addClass('open');


        //Button color
        $("#existingcustomer").removeClass('btn-warning');
        $("#existingcustomer").addClass('btn-success');

        $("#newcustomer").removeClass('btn-success');
        $("#newcustomer").addClass('btn-warning');


    });

    $("#newcustomer").click(function () {
        $(".existingcustomer").removeClass('open');
        $(".newcustomer").addClass('open');

        //Button color
        $("#newcustomer").removeClass('btn-warning');
        $("#newcustomer").addClass('btn-success');

        $("#existingcustomer").removeClass('btn-success');
        $("#existingcustomer").addClass('btn-warning');

    });

    $('.dataTable').DataTable({
        "scrollX": true
    });

});