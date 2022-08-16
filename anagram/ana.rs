use std::{
    collections::HashMap,
    fs::File,
    io::{BufRead, BufReader},
};

fn main() {
    let mut the_dict: HashMap<String, Vec<String>> = HashMap::new();

    let file = File::open("wordlist.txt").unwrap();
    let mut words = BufReader::new(file).lines();

    while let Some(Ok(word)) = words.next() {
        let word = word.trim().replace("'", "");

        if word.len() < 2 {
            continue;
        }

        let mut word_vec: Vec<char> = word.chars().collect();
        word_vec.sort();
        let sorted_word: String = word_vec.into_iter().collect();

        if the_dict.contains_key(&sorted_word) {
            if !the_dict[&sorted_word].contains(&word) {
                the_dict.get_mut(&sorted_word).unwrap().push(word);
            }
        } else {
            the_dict.insert(sorted_word, vec![]);
        }
    }

    let mut counter = 0;
    for (k, v) in the_dict {
        if v.len() >= 10 {
            counter += 1;
            println!("[{}] {} has {} words: {:?}", counter, k, v.len(), v);
        }
    }
}
