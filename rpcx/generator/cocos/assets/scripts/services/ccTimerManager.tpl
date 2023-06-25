import { Component } from 'cc';

class TimerNode {
  public callback: Function;
  public duration: number;
  public delay: number;
  public repeat: number;
  public passedTime: number;
  public param: any;
  public isRemoved: boolean;
  public timerId: number;

  /**
   * Timer
   * @param callback
   * @param repeat
   * @param duration
   * @param delay
   * @param param
   */
  constructor(callback: Function, repeat: number, duration: number, delay: number, param: any) {
    this.callback = callback;
    this.repeat = repeat;
    this.duration = duration;
    this.delay = delay;
    this.param = param;
    this.passedTime = duration;
    this.isRemoved = false;
    this.timerId = 0;
  }
}

export class TimerManager extends Component {
  private static instance: TimerManager | null = null;

  private autoIncId: number = 1;
  private timers: Record<number, TimerNode> = {};
  private removeTimers: TimerNode[] = [];
  private newAddTimers: TimerNode[] = [];

  public static get Instance(): TimerManager {
    return TimerManager.instance!;
  }

  protected onLoad(): void {
    if (!TimerManager.instance) {
      TimerManager.instance = this;
    } else {
      this.destroy();
      return;
    }
  }

  update(dt: number): void {
    this.newAddTimers.forEach((timer) => {
      this.timers[timer.timerId] = timer;
    });
    this.newAddTimers.length = 0;

    for (const key in this.timers) {
      const timer = this.timers[key];
      if (timer.isRemoved) {
        this.removeTimers.push(timer);
        continue;
      }

      timer.passedTime += dt;

      if (timer.passedTime >= (timer.delay + timer.duration)) {
        timer.callback(timer.param);
        timer.repeat--;
        timer.passedTime -= (timer.delay + timer.duration);
        timer.delay = 0;

        if (timer.repeat === 0) {
          timer.isRemoved = true;
          this.removeTimers.push(timer);
        }
      }
    }

    this.removeTimers.forEach((timer) => {
      delete this.timers[timer.timerId];
    });
    this.removeTimers.length = 0;
  }

  /**
   * 排程無參數循環
   * @param func
   * @param repeat
   * @param duration
   * @param delay
   */
  public Schedule(func: Function, repeat: number, duration: number, delay: number): number;
  /**
   * 排程帶參數循環
   * @param func
   * @param repeat
   * @param duration
   * @param delay
   * @param param
   */
  public Schedule(func: Function, repeat: number, duration: number, delay: number, param: any): number;
  public Schedule(...args: any): number {
    const [func, repeat, duration, delay, param] = args;
    const timer = new TimerNode(func, repeat, duration, delay, param);
    const timerId = this.autoIncId++;
    timer.timerId = timerId;
    this.newAddTimers.push(timer);
    return timerId;
  }

  /**
   * 排程無參數執行1次
   * @param func
   * @param delay
   */
  public Once(func: Function, delay: number): number;
  /**
   * 排程帶參數執行1次
   * @param func
   * @param delay
   * @param param
   */
  public Once(func: Function, delay: number, param: any): number;
  public Once(...args: any): number {
    const [func, delay, param] = args;
    return this.Schedule(func, 1, 0, delay, param);
  }

  /**
   * 移除排程
   * @param timerId
   */
  public Unschedule(timerId: number): void {
    const timer = this.timers[timerId];
    if (timer) {
      timer.isRemoved = true;
    }
  }
}
