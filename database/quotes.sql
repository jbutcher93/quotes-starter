drop table if exists quotes;

create table quotes 
(ID varchar(200) primary key not null unique, 
phrase varchar(500), 
author varchar(50));

insert into quotes(ID, phrase, author) values
    ('b513f4ec-ddd8-4d54-ae47-d78a2ab436', 'Don''t communicate by sharing memory, share memory by communicating.', 'Rob Pike'),
    ('dd50bc8d-17cc-4bf9-a5ce-674d9e501408', 'Concurrency is not parallelism.', 'Rob Pike'),
    ('fe4a522e-d668-4e51-85aa-65be75049618', 'Clear is better than clever.', 'Rob Pike'),
    ('20f41f99-159d-4539-9182-3c3e531ebbf9', 'When reviewing Go code, if I run into a situation where I see an unnecessary deviation from idiomatic Go style or best practice, I add an entry here complete with some rationale, and link to it.', 'Dmitri Shuralyov'),
    ('b0372040-94aa-4f1f-b905-05ee2e2efecd', 'I can do this for the smallest and most subtle of details, since I care about Go a lot. I can reuse this each time the same issue comes up, instead of having to re-write the rationale multiple times, or skip explaining why I make a given suggestion.', 'Dmitri Shuralyov');