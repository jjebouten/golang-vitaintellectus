
/**
 * Track downloads, share buttons + external links
 */
jQuery(function(){

    var filetypes = /\.(zip|exe|dmg|pdf|doc.*|xls.*|ppt.*|mp3|txt|rar|wma|mov|avi|wmv|flv|wav)$/i;
    var baseHref = '';
    if (jQuery('base').attr('href') != undefined) baseHref = jQuery('base').attr('href');
    jQuery('a').on('click', function(event) {

        // only proceed if ga is available
        if (typeof(ga) === 'undefined') return;

        var el = jQuery(this);
        var track = true;
        var href = (typeof(el.attr('href')) != 'undefined' ) ? el.attr('href') :"";
        var isThisDomain = href.match(document.domain.split('.').reverse()[1] + '.' + document.domain.split('.').reverse()[0]);
        if (!href.match(/^javascript:/i)) {
            var elEv = []; elEv.value=0, elEv.non_i=false;
            if (el.data('share-type')) {
                elEv.category = "share";
                elEv.action = el.data('share-type');
                elEv.label = el.data('share') ? el.data('share') : $('title').text();
                elEv.loc = href;
            }
            else if (href.match(/^mailto\:/i)) {
                elEv.category = "email";
                elEv.action = "click";
                elEv.label = href.replace(/^mailto\:/i, '');
                elEv.loc = href;
            }
            else if (href.match(filetypes)) {
                var extension = (/[.]/.exec(href)) ? /[^.]+$/.exec(href) : undefined;
                elEv.category = "download";
                elEv.action = "click-" + extension[0];
                elEv.label = href.replace(/ /g,"-");
                elEv.loc = baseHref + href;
            }
            // secure download
            else if (href.match(/download=/)) {
                var extension = (/[.]/.exec($(this).text())) ? /[^.]+$/.exec($(this).text()) : undefined;
                elEv.category = "download";
                elEv.action = "click-" + (extension ? extension[0] : 'download');
                elEv.label = $(this).text();
                elEv.loc = baseHref + href;
            }
            else if (href.match(/^https?\:/i) && !isThisDomain) {
                elEv.category = "external";
                elEv.action = "click";
                elEv.label = href.replace(/^https?\:\/\//i, '');
                elEv.non_i = true;
                elEv.loc = href;
            }
            else if (href.match(/^tel\:/i)) {
                elEv.category = "telephone";
                elEv.action = "click";
                elEv.label = href.replace(/^tel\:/i, '');
                elEv.loc = href;
            }
            else track = false;

            if (track) {
                ga('send', 'event', elEv.category.toLowerCase(), elEv.action.toLowerCase(), elEv.label.toLowerCase(), elEv.value);
                if ( el.attr('target') == undefined || el.attr('target').toLowerCase() != '_blank') {
                    setTimeout(function() { location.href = elEv.loc; }, 400);
                    event.preventDefault();
                }
            }
        }
    });
});
