import { AudioManager } from "./AudioManager"
import { MAINGAME_AUDIO } from "./AudioConstant";
import { ResManager } from "../ResManager";

export class MainGameAudioManager extends AudioManager {

    public static get Instance(): MainGameAudioManager {
        return MainGameAudioManager.instance!;
    }

    protected onLoad(): void {
        if(!MainGameAudioManager.instance) {
            MainGameAudioManager.instance = this;
        } else {
            this.destroy();
            return;
        }
        const audioClip = ResManager.Instance.getAsset("Audio", "MAINGAME_AUDIO")
        this.loadAudioClip(audioClip, MAINGAME_AUDIO);
    }

}