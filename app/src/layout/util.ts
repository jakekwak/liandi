import {Layout} from "./index";
import {Wnd} from "./Wnd";
import {i18n} from "../i18n";
import {mountFile, mountWebDAV} from "../util/mount";
import {Tab} from "./Tab";
import {Model} from "./Model";
import {Graph} from "../graph";
import {Editor} from "../editor";
import {Files} from "../files";
import {setPadding, setTypewriterPosition} from "../../vditore/src/ts/ui/initUI";
import {newFile} from "../util/newFile";
import {Constants} from "../constants";
import {Outline} from "../outline";

export const resizeTabs = () => {
    const models = getAllModels();
    models.editor.forEach((item) => {
        if (item.vditore.vditor) {
            setPadding(item.vditore.vditor);
            setTypewriterPosition(item.vditore.vditor);
            item.element.style.height = item.element.parentElement.clientHeight + "px";
        }
    });
    models.graph.forEach((item) => {
      item.resize();
    });
};

export const getAllModels = () => {
    const models: IModels = {
        editor: [],
        graph: [],
        files: []
    };
    const getTabs = (layout: Layout) => {
        for (let i = 0; i < layout.children.length; i++) {
            const item = layout.children[i];
            if (item instanceof Tab) {
                const model = item.model;
                if (model instanceof Editor) {
                    models.editor.push(model);
                } else if (model instanceof Graph) {
                    models.graph.push(model);
                } else if (model instanceof Files) {
                    models.files.push(model);
                }
            } else {
                getTabs(item as Layout);
            }
        }
    };

    getTabs(window.liandi.layout);
    return models;
};

export const copyTab = (tab: Tab) => {
    let panel = "";
    if (tab.model instanceof Graph) {
        panel = '<div class="graph__input"><input class="input"></div><div class="fn__flex-1"></div>';
    }
    return new Tab({
        title: tab.headElement.innerHTML.replace('<svg class="item__svg item__svg--close"><use xlink:href="#iconClose"></use></svg>', ""),
        panel,
        callback(newTab: Tab) {
            let model: Model;
            if (tab.model instanceof Editor) {
                model = new Editor({
                    tab: newTab,
                    url: tab.model.url,
                    path: tab.model.path
                });
            } else if (tab.model instanceof Graph) {
                model = new Graph({
                    tab: newTab,
                    url: tab.model.url,
                    path: tab.model.path
                });
            } else if (tab.model instanceof Files) {
                model = new Files(newTab);
            } else if (tab.model instanceof Outline) {
                model = new Outline({
                    tab: newTab,
                    contentElement: tab.panelElement.innerHTML,
                    url: tab.model.url,
                    path: tab.model.path
                });
            }
            newTab.addModel(model);
        }
    });
};

export const getInstanceById = (id: string) => {
    const _getInstanceById = (item: Layout | Wnd, id: string) => {
        if (item.id === id) {
            return item;
        }
        if (!item.children) {
            return;
        }
        let ret: Tab | Layout | Wnd;
        for (let i = 0; i < item.children.length; i++) {
            ret = _getInstanceById(item.children[i] as Layout, id) as Tab;
            if (ret) {
                return ret;
            }
        }
    };
    return _getInstanceById(window.liandi.layout, id);
};

export const addResize = (obj: Layout | Wnd) => {
    if (obj.resize) {
        const resizeWnd = (resizeElement: HTMLElement, direction: string) => {
            const setSize = (item: HTMLElement, direction: string) => {
                if (item.classList.contains("fn__flex-1")) {
                    if (direction === "lr") {
                        item.style.width = item.clientWidth + "px";
                    } else {
                        item.style.height = item.clientHeight + "px";
                    }
                    item.classList.remove("fn__flex-1");
                }
            };

            resizeElement.addEventListener("mousedown", (event: MouseEvent) => {
                const documentSelf = document;
                const nextElement = resizeElement.nextElementSibling as HTMLElement;
                const previousElement = resizeElement.previousElementSibling as HTMLElement;
                setSize(nextElement, direction);
                setSize(previousElement, direction);
                const x = event[direction === "lr" ? "clientX" : "clientY"];
                const previousSize = direction === "lr" ? previousElement.clientWidth : previousElement.clientHeight;
                const nextSize = direction === "lr" ? nextElement.clientWidth : nextElement.clientHeight;
                document.body.style.userSelect = "none";

                documentSelf.ondragstart = () => false;

                documentSelf.onmousemove = (moveEvent: MouseEvent) => {
                    const previousNowSize = (previousSize + (moveEvent[direction === "lr" ? "clientX" : "clientY"] - x));
                    const nextNowSize = (nextSize - (moveEvent[direction === "lr" ? "clientX" : "clientY"] - x));
                    if (previousNowSize < 6 || nextNowSize < 6) {
                        return;
                    }
                    previousElement.style[direction === "lr" ? "width" : "height"] = previousNowSize + "px";
                    nextElement.style[direction === "lr" ? "width" : "height"] = nextNowSize + "px";
                };

                documentSelf.onmouseup = () => {
                    document.body.style.userSelect = "auto";
                    documentSelf.onmousemove = null;
                    documentSelf.onmouseup = null;
                    documentSelf.ondragstart = null;
                    documentSelf.onselectstart = null;
                    documentSelf.onselect = null;

                    if (!nextElement.nextElementSibling) {
                        nextElement.style[direction === "lr" ? "width" : "height"] = "";
                        nextElement.classList.add("fn__flex-1");
                    }
                    window.liandi.rightLayoutWidth = window.liandi.rightLayout.element.clientWidth;
                    window.liandi.bottomLayoutHeight = window.liandi.bottomLayout.element.clientHeight;
                    resizeTabs();
                };
            });
        };

        const resizeElement = document.createElement("div");
        if (obj.resize === "lr") {
            resizeElement.classList.add("layout__resize--lr");
        }
        resizeElement.classList.add("layout__resize");
        obj.element.insertAdjacentElement("beforebegin", resizeElement);
        resizeWnd(resizeElement, obj.resize);
    }
};

export const addInitWnd = () => {
    const wnd = new Wnd();
    wnd.addTab(new Tab({
        panel: `<div class="layout__empty">
                    <div class="item fn__flex-inline">${i18n[window.liandi.config.lang].search}/${i18n[window.liandi.config.lang].config} &lt;Double Shift></div>
                    <div class="item fn__a fn__pointer" id="editorEmptyFile">${i18n[window.liandi.config.lang].newFile}</div>
                    <div class="item fn__a fn__pointer" id="editorEmptyMount">${i18n[window.liandi.config.lang].mount}</div>
                    <div class="item fn__a fn__pointer" id="editorEmptyMountDAV">${i18n[window.liandi.config.lang].mountWebDAV}</div>
                </div>`,
        callback(tab: Tab) {
            tab.panelElement.querySelector("#editorEmptyMount").addEventListener("click", () => {
                mountFile();
            });
            tab.panelElement.querySelector("#editorEmptyMountDAV").addEventListener("click", () => {
                mountWebDAV();
            });
            tab.panelElement.querySelector("#editorEmptyFile").addEventListener("click", () => {
                newFile(undefined, Constants.CB_CREATE_HOTKEY);
            });
        }
    }));
    window.liandi.centerLayout.addWnd(wnd);

    const topWnd = new Wnd();
    topWnd.addTab(new Tab({}));
    window.liandi.topLayout.addWnd(topWnd);

    const leftWnd = new Wnd();
    leftWnd.addTab(new Tab({}));
    window.liandi.leftLayout.addWnd(leftWnd);

    const rightWnd = new Wnd();
    rightWnd.addTab(new Tab({}));
    window.liandi.rightLayout.addWnd(rightWnd);

    const bottomWnd = new Wnd();
    bottomWnd.addTab(new Tab({}));
    window.liandi.bottomLayout.addWnd(bottomWnd);
};
