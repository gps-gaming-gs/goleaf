// 用來讓 services/ResManager.ts 加載 assets/AssetsPackage 資料夾內的資源
// 記得將加載的資料夾***"配置為Bundle"***
// 使用：const uiPrefab = ResManager.Instance.getAsset("GUI", "UIPrefab/Background");
//      const audioClip = ResManager.Instance.getAsset("Audio", "BGM");
import { AudioClip, Prefab } from "cc";

const MainGameRes = {
    "Audio": AudioClip,
    "GUI": [
        {
            assetType: Prefab,
            urls: [
                "UIPrefab/Background",
                "UIPrefab/LoginUI"
            ]
        }
    ]
}

const FreeGameRes = {

}

export { MainGameRes, FreeGameRes }