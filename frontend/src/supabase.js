import { createClient } from '@supabase/supabase-js';

// Replace with your actual Supabase Project URL and anon key
const supabaseUrl = 'https://wljtyfpnrelgkmbnunvv.supabase.co';
const supabaseAnonKey = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6IndsanR5ZnBucmVsZ2ttYm51bnZ2Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTIyOTE1MDEsImV4cCI6MjA2Nzg2NzUwMX0.uc0wn3qqn7BktjviSjt9TzF0AQhYYs_oZGJkC8nv06Q';

export const supabase = createClient(supabaseUrl, supabaseAnonKey);
