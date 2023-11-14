export default function is_dark(): boolean {
    return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
}
