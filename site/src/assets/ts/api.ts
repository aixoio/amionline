
export interface Event {
    id?: number;
    success: boolean;
    time_ms: number;
    target_ip: string;
    time_of_request: number;
}

export interface Events_Responce {
    cached?: boolean;
    events?: Event[];
    cached_at?: number;
    success: boolean;
    error_msg: string;
}

export async function get_all_events(): Promise<Events_Responce> {
    const request = await fetch("/api/events")
    const json = await request.json()
    
    return json as Events_Responce
}

export async function get_last_20_events(): Promise<Events_Responce> {
    const request = await fetch("/api/events/last/20")
    const json = await request.json()
    
    return json as Events_Responce
}

export async function get_all_uncached_events(): Promise<Events_Responce> {
    const request = await fetch("/api/events/uncached")
    if (request.status == 404) {
        return await get_all_events()
    }
    const json = await request.json()
    
    return json as Events_Responce
}

export async function get_last_20_uncached_events(): Promise<Events_Responce> {
    const request = await fetch("/api/events/uncached/last/20")
    if (request.status == 404) {
        return await get_last_20_events()
    }
    const json = await request.json()
    
    return json as Events_Responce
}
