/**
 * AJAX long-polling
 *
 * 1. sends a request to the server (without a timestamp parameter)
 * 2. waits for an answer from server.php (which can take forever)
 * 3. if server.php responds (whenever), put data_from_file into #response
 * 4. and call the function again
 *
 * @param ScaleId
 */
function getContent(ScaleId)
{
    var queryString = {'ScaleId' : ScaleId};

    $.ajax(
        {
            type: 'GET',
            url: '/music/lyric',
            data: queryString,
            success: function(data){
                // put result data into "obj"
                var obj = jQuery.parseJSON(data);
                // put the data_from_file into #response
                $('#response').append("<p>"+obj.Text+"</p>");
                // call the function again, this time with the timestamp we just got from server.php
                getContent(obj.ScaleId);
            }
        }
    );
}

// initialize jQuery
$(function() {
    getContent(-1);
});
