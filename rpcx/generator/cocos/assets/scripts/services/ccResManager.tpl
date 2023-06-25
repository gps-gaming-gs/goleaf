import { _decorator, Component, AssetManager, assetManager, Asset } from 'cc';

export class ResManager extends Component {
    private static instance: ResManager | null = null;

    private totalAb: number = 0;
    private nowAb: number = 0;

    private now: number = 0;
    private total: number = 0;

    protected onLoad(): void {
        if (ResManager.instance === null) {
            ResManager.instance = this;
        } else {
            this.destroy();
            return;
        }
    }

    public static get Instance(): ResManager {
        return ResManager.instance!;
    }

    private loadAndRef(abBundle: AssetManager.Bundle, url: string, typeAsset: any, progress: Function, endFunc: Function): void {
        abBundle.load(url, typeAsset, (err: any, asset: Asset) => {
            if (err) {
                console.log("load assets: ", err);
                return;
            }

            console.log("load asset success: ", url);
            asset.addRef(); // 增加一个引用技术;

            this.now++;
            if (progress) {
                progress(this.now, this.total);
            }
            if (this.now >= this.total && endFunc) {
                endFunc();
            }
        });
    }

    private loadAssetsInUrls(abBundle: AssetManager.Bundle, typeAsset: any, urls: Array<string>, progress: Function, endFunc: Function): void {
        for (let i = 0; i < urls.length; i++) {
            this.loadAndRef(abBundle, urls[i], typeAsset, progress, endFunc);
        }
    }

    private releaseAssetsInUrls(abBundle: AssetManager.Bundle, typeAsset: any, urls: Array<string>): void {
        for (let i = 0; i < urls.length; i++) {
            let asset: Asset = abBundle.get(urls[i]) as Asset;
            if (!asset) {
                continue;
            }

            asset.decRef(true);
        }
    }

    private preloadAssetsInAssetsBundles(resPkg: any, progress: Function, endFunc: Function): void {
        for (let key in resPkg) {
            let abBundle: AssetManager.Bundle = assetManager.getBundle(key) as AssetManager.Bundle;
            if (!abBundle) {
                continue;
            }

            if (Array.isArray(resPkg[key])) {
                for (let i = 0; i < resPkg[key].length; i++) {
                    this.loadAssetsInUrls(abBundle, resPkg[key][i].typeAsset, resPkg[key][i].urls, progress, endFunc);
                }
            } else {
                let typeAsset = resPkg[key];
                let urls: Array<string> = abBundle.getDirWithPath("/").map((info: any) => info.path);

                this.loadAssetsInUrls(abBundle, typeAsset, urls, progress, endFunc);
            }
        }
    }

    /**
     * 加載資源包
     * @param resPkg
     * @param progress
     * @param endFunc
     */
    public preloadResPkg(resPkg: any, progress: Function, endFunc: Function): void {
        this.totalAb = 0;
        this.nowAb = 0;

        this.total = 0;
        this.now = 0;

        for (let key in resPkg) {
            this.totalAb++;

            if (Array.isArray(resPkg[key])) {
                for (let i = 0; i < resPkg[key].length; i++) {
                    this.total += resPkg[key][i].urls.length;
                }
            }
        }

        for (let key in resPkg) {
            assetManager.loadBundle(key, (err, bundle: AssetManager.Bundle) => {
                if (err) {
                    console.log("load bundle error: ", err);
                    return;
                }

                this.nowAb++;

                if (!(Array.isArray(resPkg[key]))) {
                    let infos = bundle.getDirWithPath("/");
                    this.total += infos.length;
                }

                if (this.nowAb >= this.totalAb) {
                    this.preloadAssetsInAssetsBundles(resPkg, progress, endFunc);
                }
            });
        }
    }

    /**
     * 釋放資源包
     * @param resPkg
     */
    public releaseResPkg(resPkg: any): void {
        for (let key in resPkg) {
            let abBundle: AssetManager.Bundle = assetManager.getBundle(key) as AssetManager.Bundle;
            if (!abBundle) {
                continue;
            }

            if (Array.isArray(resPkg[key])) {
                for (let i = 0; i < resPkg[key].length; i++) {
                    this.releaseAssetsInUrls(abBundle, resPkg[key][i].typeAsset, resPkg[key][i].urls);
                }
            } else {
                let typeAsset = resPkg[key];
                let urls: Array<string> = abBundle.getDirWithPath("/").map((info: any) => info.path);

                this.releaseAssetsInUrls(abBundle, typeAsset, urls);
            }
        }
    }

    /**
     * 單個資源加載 / 加載較大資源
     * @param abName
     * @param url
     * @param typeClass
     * @param endFunc
     */
    public preloadAsset(abName: string, url: string, typeClass: any, endFunc: Function): void {
        assetManager.loadBundle(abName, (err, abBundle: AssetManager.Bundle) => {
            if (err) {
                console.log(err);
                return;
            }

            abBundle.load(url, typeClass, (err, asset: Asset) => {
                if (err) {
                    console.log(err);
                    return;
                }

                if (endFunc) {
                    endFunc();
                }
            });
        });
    }

    /**
     * 單個資源釋放
     * @param abName
     * @param url
     * @returns
     */
    public releaseAsset(abName: string, url: string): void {
        let abBundle: AssetManager.Bundle = assetManager.getBundle(abName) as AssetManager.Bundle;
        if (!abBundle) {
            return;
        }

        abBundle.release(url);
    }

    /**
     * 取得資源
     * @param abName
     * @param url
     * @returns
     */
    public getAsset(abName: string, url: string): any {
        let abBundle: AssetManager.Bundle = assetManager.getBundle(abName) as AssetManager.Bundle;
        if (!abBundle) {
            return null;
        }

        return abBundle.get(url);
    }
}
