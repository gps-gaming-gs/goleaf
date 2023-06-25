import { _decorator, AudioClip, Component, debug, error, resources, warn } from 'cc';
import gHowl from "../../lib/howler.js";

export abstract class AudioManager extends Component {

    protected static instance: any | null = null;

    protected howl: gHowl.Howl | null = null;

    protected soundIds: object = {};

    protected tmpId: any = null;

    /**
     * 全局 靜音
     * @param muted
     */
    public static Mute(muted: boolean = true) {
        gHowl.Howler.mute(muted);
    }

    /**
     * 全域 停止
     */
    public static Stop() {
        gHowl.Howler.stop();
    }

    /**
     * 載入音效
     * @param audio 音效檔
     * @param clips 切片
     * @param loop 巡迴播放
     * @param autoplay 自動播放
     */
    public loadAudioClip(audio:AudioClip, clips:object = {}, loop:boolean = false, autoplay:boolean = false){
        let HowlConf = {
            src: audio.nativeUrl,
            sprite: {},
            autoplay: autoplay,
            loop: loop,
            volume: 1.0
        };
        for (const clipKey in clips) {
            let clip  = clips[clipKey];
            let start = this.secToMills(clip.from);
            let duration = this.secToMills(clip.to - clip.from);
            HowlConf['sprite'][clipKey] = [start, duration, false];
        }
        this.howl = new gHowl.Howl(HowlConf);
    }

    public rate(rate: number, clipKey: string = "_"): AudioManager {
        let soundId = this.tmpId ?? this.soundIds[clipKey];
        this.howl.rate(rate, soundId);
        return this;
    }

    public fade(start: number, end: number, duration: number, clipKey: string = "_"): AudioManager {
        let soundId = this.tmpId ?? this.soundIds[clipKey];
        this.howl.fade(start, end, duration, soundId);
        return this;
    }

    public volume(vol: number, clipKey: string = "_"): AudioManager {
        let soundId = this.tmpId ?? this.soundIds[clipKey];
        this.howl.volume(vol, soundId);
        return this;
    }

    public stop(clipKey: string = "_"): AudioManager  {
        let soundId = this.tmpId ?? this.soundIds[clipKey];
        this.howl.stop(soundId);
        return this;
    }

    public mute(muted: boolean, clipKey: string = "_"): AudioManager  {
        let soundId = this.tmpId ?? this.soundIds[clipKey];
        this.howl.mute(muted, soundId);
        return this;
    }

    public loop(loop:boolean = true, clipKey: string = "_"): AudioManager  {
        let soundId = this.tmpId ?? this.soundIds[clipKey];
        this.howl.loop(loop, soundId);
        return this;
    }

    public play(clipKey: string = ""): AudioManager {
        let soundId = this.howl.play(clipKey);
        this.soundIds[clipKey==''?"_":clipKey] = soundId;
        this.tmpId = soundId;
        return this;
    }

    public muteAll(muted: boolean = true): AudioManager {
        for (const soundId in this.soundIds) {
            this.howl.mute(muted, this.soundIds[soundId]);
        }
        return this;
    }

    public stopAll(): AudioManager {
        for (const soundId in this.soundIds) {
            this.howl.stop(this.soundIds[soundId]);
        }
        return this;
    }

    private secToMills(sec) {
        return Math.round(sec * 1000)
    }
}

