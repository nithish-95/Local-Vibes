import { createClient } from '@supabase/supabase-js';

// Replace with your actual Supabase Project URL and anon key
const supabaseUrl = SUPERBASE_URL;
const supabaseAnonKey = SUPERBASE_PUBLIC_KEY;

export const supabase = createClient(supabaseUrl, supabaseAnonKey);
