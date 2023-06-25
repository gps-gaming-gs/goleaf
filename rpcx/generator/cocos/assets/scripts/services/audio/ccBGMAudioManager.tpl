import { AudioManager } from "./AudioManager"
import { ResManager } from "../ResManager";

export class BGMAudioManager extends AudioManager {

    public static get Instance(): BGMAudioManager {
        return BGMAudioManager.instance!;
    }

    protected onLoad(): void {
        if(!BGMAudioManager.instance) {
            BGMAudioManager.instance = this;
        } else {
            this.destroy();
            return;
        }
        // BGM
        const audioClip = ResManager.Instance.getAsset("Audio", "BGM");
        this.loadAudioClip(audioClip)
    }

}