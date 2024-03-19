// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package script

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"

func ScrollTo(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_ScrollTo_cb26`,
		Function: `function __templ_ScrollTo_cb26(id){document.getElementById(id).scrollIntoView({ behavior: 'auto' });
}`,
		Call:       templ.SafeScript(`__templ_ScrollTo_cb26`, id),
		CallInline: templ.SafeScriptInline(`__templ_ScrollTo_cb26`, id),
	}
}

func HighlightPageNavigation() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_HighlightPageNavigation_91e3`,
		Function: `function __templ_HighlightPageNavigation_91e3(){let main = qs('#main'); // Assuming qs is a shorthand for querySelector
    let pagenavItems = qsa(".pagenav-item")
    let articles = qsa('.article')
    let debounceTimer;
    main.addEventListener('scroll', function() {
        clearTimeout(debounceTimer)// Clear the previous timeout if the event is fired again within the wait period
        debounceTimer = setTimeout(function() { // set a new timeout
            let scroll = main.scrollTop + qs('#header').offsetHeight + 100;
            let activeArticle = null;
            if (scroll < articles[0].offsetTop) {
                activeArticle = articles[0];
            } else {
                for (let articlePOS = 0; articlePOS < articles.length; articlePOS++) {
                    let article = articles[articlePOS];
                    let articleTop = article.offsetTop;
                    let articleBottom = articleTop + article.offsetHeight;
                    if (scroll >= articleTop && scroll <= articleBottom) {
                        activeArticle = articles[articlePOS];
                        break;
                    }
                }
            }
            if (activeArticle) {
                let activeArticleID = activeArticle.id;
                for (let i = 0; i < pagenavItems.length; i++) {
                    let pagenavItem = pagenavItems[i];
                    let articleref = pagenavItem.getAttribute('articleref');
                    if (articleref === activeArticleID) {
                        pagenavItem.classList.add('text-blue-500');
                        pagenavItem.classList.remove('text-black');
                    } else {
                        pagenavItem.classList.add('text-black');
                        pagenavItem.classList.remove('text-blue-500');
                    }
                }
            }
        }, 100);
    });
}`,
		Call:       templ.SafeScript(`__templ_HighlightPageNavigation_91e3`),
		CallInline: templ.SafeScriptInline(`__templ_HighlightPageNavigation_91e3`),
	}
}
