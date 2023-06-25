import { AudioManager } from "./AudioManager"
import { GENERAL_AUDIO } from "./AudioConstant";
import { ResManager } from "../ResManager";

export class GeneralAudioManager extends AudioManager {

    public static get Instance(): GeneralAudioManager {
        return GeneralAudioManager.instance!;
    }

    protected onLoad(): void {
        if(!GeneralAudioManager.instance) {
            GeneralAudioManager.instance = this;
        } else {
            this.destroy();
            return;
        }
        const audioClip = ResManager.Instance.getAsset("Audio", "GENERAL_AUDIO")
        this.loadAudioClip(audioClip, GENERAL_AUDIO);
    }

}